package funcs

import (
	"bytes"
	"fmt"

	"github.com/licat233/genzero/sql"
)

//已被遗弃

type FindCount struct {
	modelName   string
	name        string
	req         string
	resp        string
	fullName    string
	IsCacheMode bool
	Table       *sql.Table
}

var _ ModelFunc = (*FindCount)(nil)

func NewFindCount(t *sql.Table, isCacheMode bool) *FindCount {
	modelName := modelName(t.Name)
	name := "FindCount"
	req := "ctx context.Context"
	resp := "(int64, error)"
	fullName := fmt.Sprintf("%s(%s) %s", name, req, resp)
	return &FindCount{
		modelName:   modelName,
		name:        name,
		req:         req,
		resp:        resp,
		fullName:    fullName,
		IsCacheMode: isCacheMode,
		Table:       t,
	}
}

func (s *FindCount) String() string {
	var buf = new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("\nfunc (m *%s) %s {\n", s.modelName, s.fullName))
	buf.WriteString("var count int64\n")
	if delField := s.Table.GetIsDeletedField(); delField != nil {
		buf.WriteString("query := fmt.Sprintf(\"select count(*) as count from %s where `" + delField.Name + "` = '0'\", m.table)\n")
	} else {
		buf.WriteString("query := fmt.Sprintf(\"select count(*) as count from %s\", m.table)\n")
	}

	if s.IsCacheMode {
		buf.WriteString("err := m.QueryRowNoCacheCtx(ctx, &count, query)\n")
	} else {
		buf.WriteString("err := m.conn.QueryRowCtx(ctx, &count, query)\n")
	}
	buf.WriteString("return count, err\n")
	buf.WriteString("}\n")
	return buf.String()
}

func (t *FindCount) FullName() string {
	return t.fullName
}
