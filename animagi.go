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

	if reflect.ValueOf(dst).Kind() != reflect.Ptr || !reflect.ValueOf(dst).Elem().CanSet() {
		return errors.New(dstError)
	}

	if reflect.PtrTo(reflect.TypeOf(src)) == reflect.TypeOf(dst) {
		reflect.ValueOf(dst).Elem().Set(reflect.ValueOf(src))
	} else {
		err = errors.New(unsupportedTransformation)
	}
	return err
}
