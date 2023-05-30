package internal

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

type Enum struct {
	Name    string
	Comment string
	Fields  EnumFieldCollection
}

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

func (e *Enum) AppendField(ef *EnumField) error {
	for _, f := range e.Fields {
		if f.Tag == ef.Tag {
			return fmt.Errorf("tag `%d` is already in use by field `%s`", ef.Tag, f.Name)
		}
	}

	e.Fields = append(e.Fields, ef)

	return nil
}

func (e *Enum) String() string {
	buf := new(bytes.Buffer)

	buf.WriteString(fmt.Sprintf("// %s \n", e.Comment))
	buf.WriteString(fmt.Sprintf("enum %s {\n", e.Name))
	buf.WriteString(fmt.Sprint(e.Fields))
	buf.WriteString("}\n")

	return buf.String()
}

type EnumField struct {
	Name string
	Tag  int
}

func NewEnumField(name string, tag int) *EnumField {
	name = strings.ToUpper(name)

	re := regexp.MustCompile(`([^\w]+)`)
	name = re.ReplaceAllString(name, "_")

	return &EnumField{name, tag}
}

func (ef EnumField) String() string {
	return fmt.Sprintf("%s = %d;\n", ef.Name, ef.Tag)
}

type EnumFieldCollection []*EnumField

func (efc EnumFieldCollection) Copy() EnumFieldCollection {
	var data = make(EnumFieldCollection, len(efc))
	for k, v := range efc {
		cp := *v
		data[k] = &cp
	}
	return data
}

func (efc EnumFieldCollection) String() string {
	buf := new(bytes.Buffer)
	for _, f := range efc {
		buf.WriteString(fmt.Sprint(f))
	}
	return buf.String()
}

type EnumCollection []*Enum

func (ec EnumCollection) Len() int {
	return len(ec)
}

func (ec EnumCollection) Less(i, j int) bool {
	return ec[i].Name < ec[j].Name
}

func (ec EnumCollection) Swap(i, j int) {
	ec[i], ec[j] = ec[j], ec[i]
}
