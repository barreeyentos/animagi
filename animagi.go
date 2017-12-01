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

	if cannotModifyField(dst) {
		return errors.New(dstError)
	}

	valueOfSrc := findValueOf(src)
	valueOfDst := findValueOf(dst)
	if valueOfSrc.Kind() == valueOfDst.Kind() {
		switch valueOfDst.Kind() {
		case reflect.Struct:
			srcDescription := describeStructure(src)
			mapToDestination("", src, dst, srcDescription)
		default:
			setValueOfDst(valueOfDst, valueOfSrc)
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
		srcFieldValue := srcValue.FieldByName(fieldName)
		if srcFieldValue.IsValid() && field.CanSet() {
			switch field.Kind() {
			case reflect.Struct:
				mapToDestination(fullPathName, srcFieldValue, field, srcDescription)
			case reflect.Ptr:
				if srcDescription[fullPathName] != nil {
					field.Set(reflect.New(reflect.TypeOf(field.Interface()).Elem()))
					setValueOfDst(field.Elem(), srcFieldValue)
				}
			default:
				if srcDescription[fullPathName] != nil {
					setValueOfDst(field, srcFieldValue)
				}
			}
		}
	}
}

func setValueOfDst(dst, src reflect.Value) {
	if dst.Type() == reflect.Indirect(src).Type() {
		dst.Set(reflect.Indirect(src))
	} else if reflect.Indirect(src).Type().ConvertibleTo(reflect.Indirect(dst).Type()) {
		dst.Set(reflect.Indirect(src).Convert(reflect.Indirect(dst).Type()))
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

func cannotModifyField(field interface{}) bool {
	return reflect.ValueOf(field).Kind() != reflect.Ptr || !reflect.ValueOf(field).Elem().CanSet()
}
