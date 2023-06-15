package funcs

import (
	"bytes"
	"fmt"

	"github.com/licat233/genzero/sql"
	"github.com/licat233/genzero/tools"
)

type FormatUuidKey struct {
	modelName   string
	name        string
	req         string
	resp        string
	fullName    string
	IsCacheMode bool
	Table       *sql.Table
}

var _ ModelFunc = (*FormatUuidKey)(nil)

func NewFormatUuidKey(t *sql.Table, isCacheMode bool) *FormatUuidKey {
	modelName := modelName(t.Name)
	name := "formatUuidKey"
	req := "uuid string"
	resp := "string"
	fullName := fmt.Sprintf("%s(%s) %s", name, req, resp)
	return &FormatUuidKey{
		modelName:   modelName,
		name:        name,
		req:         req,
		resp:        resp,
		fullName:    fullName,
		IsCacheMode: isCacheMode,
		Table:       t,
	}
}

func (t *FormatUuidKey) String() string {
	if !t.Table.ExistUuidField() {
		return ""
	}
	var buf = new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("\nfunc (m *%s) %s {", t.modelName, t.fullName))
	buf.WriteString(fmt.Sprintf("\n\treturn fmt.Sprintf(\"cache:%s:uuid:%%v\", uuid)", tools.ToLowerCamel(t.Table.Name)))
	buf.WriteString("\n}\n")
	return buf.String()
}

func (t *FormatUuidKey) FullName() string {
	if !t.Table.ExistUuidField() {
		return ""
	}
	return t.fullName
}
