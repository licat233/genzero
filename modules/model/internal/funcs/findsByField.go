package funcs

import (
	"bytes"
	"fmt"

	"github.com/licat233/genzero/sql"
	"github.com/licat233/genzero/tools"
)

type findsByField struct {
	field        sql.Field
	modelName    string
	name         string
	req          string
	resp         string
	fullName     string
	IsCacheMode  bool
	Table        *sql.Table
	reqFieldName string
}

var _ ModelFunc = (*findsByField)(nil)

type FindsByFieldCollection []*findsByField

var _ ModelFunc = (FindsByFieldCollection)(nil)

func NewFindsByFieldCollection(table *sql.Table, isCacheMode bool) FindsByFieldCollection {
	var res FindsByFieldCollection = make([]*findsByField, 0)
	for _, field := range table.GetFields() {
		res = append(res, newFindsByField(table, isCacheMode, field))
	}
	return res
}

func (f FindsByFieldCollection) String() string {
	var buf = new(bytes.Buffer)
	for _, findByField := range f {
		buf.WriteString(findByField.String() + "\n")
	}
	return buf.String()
}

func (f FindsByFieldCollection) FullName() string {
	var buf = new(bytes.Buffer)
	for _, findByField := range f {
		buf.WriteString(findByField.FullName() + "\n")
	}
	return buf.String()
}

func newFindsByField(t *sql.Table, isCacheMode bool, field sql.Field) *findsByField {
	modelName := modelName(t.Name)
	reqFieldName := tools.ToLowerCamel(field.Name)
	name := fmt.Sprintf("FindsBy%s", tools.ToCamel(reqFieldName))
	req := fmt.Sprintf("ctx context.Context, %s %s", reqFieldName, field.Type)
	resp := fmt.Sprintf("([]*%s, error)", tools.ToCamel(t.Name))
	fullName := fmt.Sprintf("%s(%s) %s", name, req, resp)
	return &findsByField{
		field:        field,
		modelName:    modelName,
		name:         name,
		req:          req,
		resp:         resp,
		fullName:     fullName,
		IsCacheMode:  isCacheMode,
		Table:        t,
		reqFieldName: reqFieldName,
	}
}

func (s *findsByField) String() string {
	var buf = new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("\nfunc (m *%s) %s {\n", s.modelName, s.fullName))
	buf.WriteString(fmt.Sprintf("var resp = make([]*%s, 0)\n", tools.ToCamel(s.Table.Name)))
	if delField := s.Table.GetIsDeletedField(); delField != nil {
		buf.WriteString("query := fmt.Sprintf(\"select %s from %s where `" + s.field.Name + "` = ? and `" + delField.Name + "` = '0' \", " + tools.ToLowerCamel(s.Table.Name) + "Rows, m.table)\n")
	} else {
		buf.WriteString("query := fmt.Sprintf(\"select %s from %s where `" + s.field.Name + "` = ? \", " + tools.ToLowerCamel(s.Table.Name) + "Rows, m.table)\n")
	}

	if s.IsCacheMode {
		buf.WriteString("err := m.QueryRowsNoCacheCtx(ctx, &resp, query, " + s.reqFieldName + ")\n")
	} else {
		buf.WriteString("err := m.conn.QueryRowsCtx(ctx, &resp, query, " + s.reqFieldName + ")\n")
	}
	buf.WriteString("return resp, err")
	buf.WriteString("\n}\n")
	return buf.String()
}

func (t *findsByField) FullName() string {
	return t.fullName
}
