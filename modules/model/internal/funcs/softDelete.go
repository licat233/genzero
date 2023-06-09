package funcs

import (
	"bytes"
	"fmt"

	"github.com/licat233/genzero/sql"
)

type SoftDelete struct {
	modelName   string
	name        string
	req         string
	resp        string
	fullName    string
	IsCacheMode bool
	Table       *sql.Table
}

var _ ModelFunc = (*SoftDelete)(nil)

func NewSoftDelete(t *sql.Table, isCacheMode bool) *SoftDelete {
	modelName := modelName(t.Name)
	name := "SoftDelete"
	req := "ctx context.Context, id int64"
	resp := "error"
	fullName := fmt.Sprintf("%s(%s) %s", name, req, resp)
	if !t.ExistIsDelField() {
		fullName = ""
	}
	return &SoftDelete{
		modelName:   modelName,
		name:        name,
		req:         req,
		resp:        resp,
		fullName:    fullName,
		IsCacheMode: isCacheMode,
		Table:       t,
	}
}

func (s *SoftDelete) String() string {
	delField := s.Table.GetIsDeletedField()
	if delField == nil {
		return ""
	}
	var buf = new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("\nfunc (m *%s) %s {", s.modelName, s.fullName))
	if delAtField := s.Table.GetDelAtField(); delAtField != nil {
		buf.WriteString("\nquery := fmt.Sprintf(\"update %s set `" + delField.Name + "` = '1', `" + delAtField.Name + "`= now() where `id` = ?\", m.table)")
	} else {
		buf.WriteString("\nquery := fmt.Sprintf(\"update %s set `" + delField.Name + "` = '1' where `id` = ?\", m.table)")
	}
	if s.IsCacheMode {
		buf.WriteString("\n_, err := m.ExecNoCacheCtx(ctx, query, id)")
	} else {
		buf.WriteString("\n_, err := m.conn.ExecCtx(ctx, query, id)")
	}
	if !s.IsCacheMode {
		buf.WriteString("\nreturn err")
		buf.WriteString("\n}\n")
		return buf.String()
	}
	buf.WriteString("\nif err != nil {")
	buf.WriteString("\nreturn err")
	buf.WriteString("\n}\n")
	buf.WriteString("\n//删除缓存")
	buf.WriteString("\nerr = m.DelCacheCtx(ctx, m.formatPrimary(id))")
	buf.WriteString("\nreturn err")
	buf.WriteString("\n}\n")
	return buf.String()
}

func (s *SoftDelete) FullName() string {
	if !s.Table.ExistIsDelField() {
		return ""
	}
	return s.fullName
}
