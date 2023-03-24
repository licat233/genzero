package sql

import (
	"fmt"
	"regexp"
	"strings"
)

type Enum struct {
	Name    string
	Comment string
	Fields  EnumFieldCollection
}

type EnumCollection []Enum

type EnumField struct {
	Name string
	Tag  int
}

type EnumFieldCollection []EnumField

func NewEnumFromStrings(name, comment string, ss []string) (*Enum, error) {
	enum := &Enum{}
	enum.Name = name
	enum.Comment = comment

	for i, s := range ss {
		err := enum.AppendField(NewEnumField(s, i))
		if nil != err {
			return nil, err
		}
	}

	return enum, nil
}

func NewEnumField(name string, tag int) *EnumField {
	name = strings.ToUpper(name)

	re := regexp.MustCompile(`([^\w]+)`)
	name = re.ReplaceAllString(name, "_")

	return &EnumField{name, tag}
}

func (e *Enum) Copy() *Enum {
	return &Enum{
		Name:    e.Name,
		Comment: e.Comment,
		Fields:  e.Fields.Copy(),
	}
}

func (e *Enum) AppendField(ef *EnumField) error {
	for _, f := range e.Fields {
		if f.Tag == ef.Tag {
			return fmt.Errorf("tag `%d` is already in use by field `%s`", ef.Tag, f.Name)
		}
	}

	e.Fields = append(e.Fields, *ef)

	return nil
}

func (e EnumCollection) Copy() EnumCollection {
	res := make(EnumCollection, len(e))
	for i, v := range e {
		res[i] = *v.Copy()
	}
	return res
}

func (e *EnumField) Copy() *EnumField {
	return &EnumField{
		Name: e.Name,
		Tag:  e.Tag,
	}
}

func (e EnumFieldCollection) Copy() EnumFieldCollection {
	res := make(EnumFieldCollection, len(e))
	for i, f := range e {
		res[i] = *f.Copy()
	}
	return res
}
