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
		reflect.ValueOf(dst).Elem().Set(valueOfSrc)
	} else if valueOfDst.Elem().Kind() == valueOfSrc.Kind() {
		reflect.ValueOf(dst).Elem().Set(valueOfSrc.Convert(valueOfDst.Elem().Type()))
	} else {
		err = errors.New(unsupportedTransformation)
	}
	return err
}
