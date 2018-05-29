package gotype

func newTypeValueBind(typ, val Type) Type {
	if typ.Kind() == Struct {
		if v, ok := val.(*typeValuePairs); ok {
			nt := &typeStruct{}
			fl := typ.NumField()
			for i := 0; i != fl; i++ {
				f := typ.Field(i)
				name := f.Name()
				b := f
				if val, ok := v.li.Search(name); ok {
					b = newTypeValueBind(f, val)
				}
				nt.fields = append(nt.fields, b)
			}
			al := typ.NumAnonymo()
			for i := 0; i != al; i++ {
				f := typ.Field(i)
				name := f.Name()
				b := f
				if val, ok := v.li.Search(name); ok {
					b = newTypeValueBind(f, val)
				}
				nt.anonymo = append(nt.anonymo, b)
			}
			return nt
		}
	}

	return &typeValueBind{
		Type: typ,
		val:  val,
	}
}

type typeValueBind struct {
	Type
	val Type
}

func (t *typeValueBind) Value() string {
	return t.val.Value()
}
