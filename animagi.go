package animagi

import (
	"errors"
	"reflect"
)

const (
	dstError                  = "dst must be settable"
	unsupportedTransformation = "could not transform to dst"
)

/*
Transform will map the data from src into
dst by calculating the fields most similar
counterpart and copying the values over.
If src and dst are of the same type then
Transform basically does a copy.

dst must be settable or an error will be returned
*/
func Transform(src, dst interface{}) (err error) {

	typeOfSrc := reflect.TypeOf(src)
	typeOfDst := reflect.TypeOf(dst)
	valueOfSrc := reflect.ValueOf(src)
	valueOfDst := reflect.ValueOf(dst)

	if valueOfDst.Kind() != reflect.Ptr || !valueOfDst.Elem().CanSet() {
		return errors.New(dstError)
	}

	if reflect.PtrTo(typeOfSrc) == typeOfDst {
		valueOfDst.Elem().Set(valueOfSrc)
	} else if valueOfDst.Elem().Kind() == valueOfSrc.Kind() {
		if valueOfSrc.Kind() != reflect.Struct {
			valueOfDst.Elem().Set(valueOfSrc.Convert(valueOfDst.Elem().Type()))
		} else {
			dstDescription := inspectDestination(dst)
			mapToDestination("", src, dst, dstDescription)
		}
	} else {
		err = errors.New(unsupportedTransformation)
	}
	return err
}

func inspectDestination(dst interface{}) map[string]reflect.Type {
	dstDescription := make(map[string]reflect.Type)
	var dstValue reflect.Value

	if reflect.TypeOf(dst) != reflect.TypeOf(dstValue) {
		dstValue = reflect.Indirect(reflect.ValueOf(dst))
	} else {
		dstValue = dst.(reflect.Value)
	}

	for i := 0; i < dstValue.NumField(); i++ {
		field := dstValue.Field(i)
		if field.CanSet() {
			fieldName := dstValue.Type().Field(i).Name
			switch reflect.Indirect(field).Kind() {
			case reflect.Struct:
				subDescription := inspectDestination(field)
				for k, v := range subDescription {
					dstDescription[fieldName+"."+k] = v
				}
			default:
				dstDescription[fieldName] = field.Type()
			}
		}
	}
	return dstDescription
}

func mapToDestination(currentLevel string, src, dst interface{}, dstDescription map[string]reflect.Type) {
	var srcValue reflect.Value

	var dstValue reflect.Value

	if reflect.TypeOf(dst) != reflect.TypeOf(dstValue) {
		dstValue = reflect.Indirect(reflect.ValueOf(dst))
	} else {
		dstValue = dst.(reflect.Value)
	}

	if reflect.TypeOf(src) != reflect.TypeOf(srcValue) {
		srcValue = reflect.Indirect(reflect.ValueOf(src))
	} else {
		srcValue = src.(reflect.Value)
	}

	for i := 0; i < srcValue.NumField(); i++ {
		field := srcValue.Field(i)
		fieldName := srcValue.Type().Field(i).Name
		var fullPathName string
		if len(currentLevel) != 0 {
			fullPathName = currentLevel + "." + fieldName
		} else {
			fullPathName = fieldName
		}
		switch reflect.Indirect(field).Kind() {
		case reflect.Struct:
			if dstValue.FieldByName(fieldName).IsValid() {
				mapToDestination(fullPathName, field, dstValue.FieldByName(fieldName), dstDescription)
			}
		default:
			if dstDescription[fullPathName] != nil && dstValue.CanSet() {
				if dstValue.FieldByName(fieldName).CanSet() {
					if reflect.Indirect(field).Type() == dstValue.FieldByName(fieldName).Type() {
						dstValue.FieldByName(fieldName).Set(field)
					} else {
						dstValue.FieldByName(fieldName).Set(field.Convert(dstValue.FieldByName(fieldName).Type()))
					}
				}
			}
		}
	}
}
