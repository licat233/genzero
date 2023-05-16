package funcs

import (
	"bytes"
	"fmt"

	"github.com/licat233/genzero/sql"
	"github.com/licat233/genzero/tools"
)

//已被遗弃

type FindsByIds struct {
	modelName   string
	name        string
	req         string
	resp        string
	fullName    string
	IsCacheMode bool
	Table       *sql.Table
}

var _ ModelFunc = (*FindsByIds)(nil)

func NewFindsByIds(t *sql.Table, isCacheMode bool) *FindsByIds {
	modelName := modelName(t.Name)
	name := "FindsByIds"
	req := "ctx context.Context, ids []int64"
	resp := fmt.Sprintf("([]*%s, error)", tools.ToCamel(t.Name))
	fullName := fmt.Sprintf("%s(%s) %s", name, req, resp)
	return &FindsByIds{
		modelName:   modelName,
		name:        name,
		req:         req,
		resp:        resp,
		fullName:    fullName,
		IsCacheMode: isCacheMode,
		Table:       t,
	}
}

func (s *FindsByIds) String() string {
	var buf = new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("\nfunc (m *%s) %s {\n", s.modelName, s.fullName))
	buf.WriteString(fmt.Sprintf("var resp = make([]*%s, 0)\n", tools.ToCamel(s.Table.Name)))
	buf.WriteString("if len(ids) == 0 {\n")
	buf.WriteString("return resp, nil\n")
	buf.WriteString("}\n")
	if s.Table.HasDeleteFiled {
		buf.WriteString("query := fmt.Sprintf(\"select %s from %s where `id` in(?) and `is_deleted` = '0' \", " + tools.ToLowerCamel(s.Table.Name) + "Rows, m.table)\n")
	} else {
		buf.WriteString("query := fmt.Sprintf(\"select %s from %s where `id` in(?) \", " + tools.ToLowerCamel(s.Table.Name) + "Rows, m.table)\n")
	}

	if s.IsCacheMode {
		buf.WriteString("err := m.QueryRowsNoCacheCtx(ctx, &resp, query, ids)\n")
	} else {
		buf.WriteString("err := m.conn.QueryRowsCtx(ctx, &resp, query, ids)\n")
	}
	buf.WriteString("return resp, err")
	buf.WriteString("\n}\n")
	return buf.String()
}

func (t *FindsByIds) FullName() string {
	return t.fullName
}
