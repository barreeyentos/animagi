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
			srcDescription := describeStructure(src)
			mapToDestination("", src, dst, srcDescription)
		}
	} else {
		err = errors.New(unsupportedTransformation)
	}
	return err
}

func describeStructure(structure interface{}) map[string]reflect.Type {
	structureDescription := make(map[string]reflect.Type)
	structureValue := findValueOf(structure)

	for i := 0; i < structureValue.NumField(); i++ {
		field := structureValue.Field(i)
		fieldName := structureValue.Type().Field(i).Name
		switch reflect.Indirect(field).Kind() {
		case reflect.Struct:
			subDescription := describeStructure(field)
			for k, v := range subDescription {
				structureDescription[fieldName+"."+k] = v
			}
		default:
			structureDescription[fieldName] = field.Type()
		}
	}
	return structureDescription
}

func mapToDestination(currentLevel string, src, dst interface{}, srcDescription map[string]reflect.Type) {
	srcValue := findValueOf(src)
	dstValue := findValueOf(dst)

	for i := 0; i < dstValue.NumField(); i++ {
		field := dstValue.Field(i)
		fieldName := dstValue.Type().Field(i).Name
		fullPathName := appendFieldName(currentLevel, fieldName)

		switch reflect.Indirect(field).Kind() {
		case reflect.Struct:
			if srcValue.FieldByName(fieldName).IsValid() {
				mapToDestination(fullPathName, srcValue.FieldByName(fieldName), field, srcDescription)
			}
		default:
			if srcDescription[fullPathName] != nil && field.CanSet() {
				srcFieldValue := srcValue.FieldByName(fieldName)
				if reflect.Indirect(field).Type() == srcFieldValue.Type() {
					field.Set(srcFieldValue)
				} else if srcFieldValue.Type().ConvertibleTo(reflect.Indirect(field).Type()) {
					field.Set(srcFieldValue.Convert(reflect.Indirect(field).Type()))
				}
			}
		}
	}
}

func findValueOf(val interface{}) (valueOf reflect.Value) {
	if reflect.TypeOf(val) != reflect.TypeOf(valueOf) {
		valueOf = reflect.Indirect(reflect.ValueOf(val))
	} else {
		valueOf = val.(reflect.Value)
	}
	return valueOf
}

func appendFieldName(prefix, fieldName string) (fullName string) {
	if len(prefix) != 0 {
		fullName = prefix + "." + fieldName
	} else {
		fullName = fieldName
	}
	return fullName
}
