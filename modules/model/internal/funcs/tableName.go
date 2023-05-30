package funcs

import (
	"bytes"
	"fmt"

	"github.com/licat233/genzero/sql"
)

type TableName struct {
	modelName   string
	name        string
	req         string
	resp        string
	fullName    string
	IsCacheMode bool
	Table       *sql.Table
}

var _ ModelFunc = (*TableName)(nil)

func NewTableName(t *sql.Table, isCacheMode bool) *TableName {
	modelName := modelName(t.Name)
	name := "TableName"
	req := ""
	resp := "string"
	fullName := fmt.Sprintf("%s(%s) %s", name, req, resp)
	return &TableName{
		modelName:   modelName,
		name:        name,
		req:         req,
		resp:        resp,
		fullName:    fullName,
		IsCacheMode: isCacheMode,
		Table:       t,
	}
}

func (t *TableName) String() string {
	var buf = new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("\nfunc (m *%s) %s {", t.modelName, t.fullName))
	buf.WriteString("\n\treturn m.table")
	buf.WriteString("\n}\n")
	return buf.String()
}

func (t *TableName) FullName() string {
	return t.fullName
}
