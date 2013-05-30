package convert

import (
	"reflect"
)

func Convert(src interface{}, dstptr interface{}) error {

	t1 := reflect.ValueOf(src)
	t2 := reflect.ValueOf(dstptr)

	l := t1.NumField()

	for i := 0; i < l; i++ {
		f1 := t1.Field(i)
		f2 := t2.Elem().Field(i)
		convertField(f1, &f2)
	}

	return nil
}
func ConvertMap(src map[string]interface{}, dstrptr interface{}) error {
	l := len(src)
	dval := reflect.ValueOf(dstrptr)
	dtype := dval.Elem().Type()
	for i := 0; i < l; i++ {
		nm := dtype.Field(i).Name
		f1 := reflect.ValueOf(src[nm])
		f2 := dval.Elem().FieldByName(nm)
		convertField(f1, &f2)
	}
	return nil
}
func convertField(src reflect.Value, dstptr *reflect.Value) error {
	t1 := src.Type()
	t2 := dstptr.Type()
	if t1.ConvertibleTo(t2) {
		resp := src.Convert(t2)
		dstptr.Set(resp)
	}
	return nil
}
