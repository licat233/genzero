package funcs

import (
	"bytes"
	"fmt"

	"github.com/licat233/genzero/parser"
)

type TableName struct {
	modelName string
	name      string
	req       string
	resp      string
	fullName  string
	Table     *parser.Table
}

var _ ModelFunc = (*TableName)(nil)

func NewTableName(t *parser.Table) *TableName {
	modelName := modelName(t.Name)
	name := "TableName"
	req := ""
	resp := "string"
	fullName := fmt.Sprintf("%s(%s) %s", name, req, resp)
	return &TableName{
		modelName: modelName,
		name:      name,
		resp:      resp,
		fullName:  fullName,
		Table:     t,
	}
}

func (t *TableName) String() string {
	var buf = new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("\nfunc (m *%s) %s {", t.modelName, t.fullName))
	buf.WriteString("\n\treturn m.table")
	buf.WriteString("\n}\n")
	return buf.String()
}

func (s *TableName) FullName() string {
	return s.fullName
}

func (s *TableName) Req() string {
	return s.req
}

func (s *TableName) Resp() string {
	return s.resp
}

func (s *TableName) Name() string {
	return s.name
}

func (s *TableName) ModelName() string {
	return s.modelName
}
