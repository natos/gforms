package gforms

import (
	"reflect"
)

type MultipleTextField struct {
	BaseField
}

func (f *MultipleTextField) New() FieldInterface {
	fi := new(MultipleTextFieldInstance)
	fi.Model = f
	fi.V = nilV("")
	return fi
}

type MultipleTextFieldInstance struct {
	FieldInstance
}

func NewMultipleTextField(name string, vs Validators, ws ...Widget) Field {
	f := new(MultipleTextField)
	f.name = name
	f.validators = vs
	if len(ws) > 0 {
		f.widget = ws[0]
	} else {
		f.widget = SelectMultipleWidget(map[string]string{}, nil)
	}
	return f
}

func (f *MultipleTextFieldInstance) Clean(data Data) error {
	m, hasField := data[f.Model.GetName()]
	if hasField {
		f.V = m
		m.Kind = reflect.Slice
		m.Value = m.rawValueAsStringArray()
		m.IsNil = false
		return nil
	}
	return nil
}

func (f *MultipleTextFieldInstance) html() string {
	return ""
}

func (f *MultipleTextFieldInstance) Html() string {
	return fieldToHtml(f)
}