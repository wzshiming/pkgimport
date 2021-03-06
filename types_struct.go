package gotype

import (
	"bytes"
)

type typeStruct struct {
	typeBase
	fields types // fields
}

func (t *typeStruct) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("struct{")
	for i, v := range t.fields {
		if i != 0 {
			buf.WriteString("; ")
		}
		buf.WriteString(v.String())
	}
	buf.WriteByte('}')
	return buf.String()
}

func (t *typeStruct) Kind() Kind {
	return Struct
}

func (t *typeStruct) NumField() int {
	return t.fields.Len()
}

func (t *typeStruct) Field(i int) Type {
	return t.fields.Index(i)
}

func (t *typeStruct) FieldByName(name string) (Type, bool) {
	anonymo := []int{}
	for i, v := range t.fields {
		fieldName := v.Name()
		if fieldName == name {
			return v, true
		}
		if v.IsAnonymous() {
			anonymo = append(anonymo, i)
		}
	}

	for _, i := range anonymo {
		v := t.fields[i]
		v = v.Elem()
		t, ok := v.FieldByName(name)
		if ok {
			return t, true
		}
	}
	return nil, false
}

func (t *typeStruct) MethodByName(name string) (Type, bool) {
	for _, v := range t.fields {
		if v.IsAnonymous() {
			v = v.Elem()
			b, ok := v.MethodByName(name)
			if ok {
				return b, true
			}
		}
	}
	return nil, false
}
