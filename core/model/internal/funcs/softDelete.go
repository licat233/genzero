package funcs

import (
	"bytes"
	"fmt"

	"github.com/licat233/genzero/core/model/conf"
	"github.com/licat233/genzero/sql"
)

type SoftDelete struct {
	modelName string
	name      string
	req       string
	resp      string
	fullName  string
	Table     *sql.Table
}

var _ ModelFunc = (*SoftDelete)(nil)

func NewSoftDelete(t *sql.Table) *SoftDelete {
	modelName := modelName(t.Name)
	name := "SoftDelete"
	req := "ctx context.Context, id int64"
	resp := "error"
	fullName := fmt.Sprintf("%s(%s) %s", name, req, resp)
	if !t.HasDeleteFiled {
		fullName = ""
	}
	return &SoftDelete{
		modelName: modelName,
		name:      name,
		req:       req,
		resp:      resp,
		fullName:  fullName,
		Table:     t,
	}
}

func (s *SoftDelete) String() string {
	if !s.Table.HasDeleteFiled {
		return ""
	}
	var buf = new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("\nfunc (m *%s) %s {", s.modelName, s.fullName))
	buf.WriteString("\n//更新 is_deleted 状态，不走缓存")
	buf.WriteString("\nquery := fmt.Sprintf(\"update %s set `is_deleted` = '1', `delete_at` = now() where `id` = ?\", m.table)")
	if conf.IsCacheMode {
		buf.WriteString("\n_, err := m.ExecNoCacheCtx(ctx, query, id)")
	} else {
		buf.WriteString("\n_, err := m.conn.ExecCtx(ctx, query, id)")
	}
	if !conf.IsCacheMode {
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
	return s.fullName
}

func (s *SoftDelete) Req() string {
	return s.req
}

func (s *SoftDelete) Resp() string {
	return s.resp
}

func (s *SoftDelete) Name() string {
	return s.name
}

func (s *SoftDelete) ModelName() string {
	return s.modelName
}
