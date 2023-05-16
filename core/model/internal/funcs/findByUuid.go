package funcs

import (
	"bytes"
	"fmt"

	"github.com/licat233/genzero/sql"
	"github.com/licat233/genzero/tools"
)

//已被遗弃

type FindByUuid struct {
	modelName   string
	name        string
	req         string
	resp        string
	fullName    string
	IsCacheMode bool
	Table       *sql.Table
}

var _ ModelFunc = (*FindByUuid)(nil)

func NewFindByUuid(t *sql.Table, isCacheMode bool) *FindByUuid {
	modelName := modelName(t.Name)
	name := "FindByUuid"
	req := "ctx context.Context, uuid string"
	resp := fmt.Sprintf("(*%s, error)", tools.ToCamel(t.Name))
	fullName := fmt.Sprintf("%s(%s) %s", name, req, resp)
	return &FindByUuid{
		modelName:   modelName,
		name:        name,
		req:         req,
		resp:        resp,
		fullName:    fullName,
		IsCacheMode: isCacheMode,
		Table:       t,
	}
}

func (s *FindByUuid) String() string {
	var buf = new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("\nfunc (m *%s) %s {\n", s.modelName, s.fullName))
	// buf.WriteString("uuidKey := m.formatUuidKey(uuid)\n")
	buf.WriteString(fmt.Sprintf("var resp %s\n", tools.ToCamel(s.Table.Name)))
	if s.Table.HasDeleteFiled {
		buf.WriteString("query := fmt.Sprintf(\"select %s from %s where `uuid` = ? and `is_deleted` = '0' limit 1\", " + tools.ToLowerCamel(s.Table.Name) + "Rows, m.table)\n")
	} else {
		buf.WriteString("query := fmt.Sprintf(\"select %s from %s where `uuid` = ? limit 1\", " + tools.ToLowerCamel(s.Table.Name) + "Rows, m.table)\n")
	}

	if s.IsCacheMode {
		// buf.WriteString("err := m.QueryRowCtx(ctx, &resp, uuidKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {")
		// buf.WriteString("return conn.QueryRowCtx(ctx, v, query, uuid)")
		// buf.WriteString("})")
		buf.WriteString("err := m.QueryRowNoCacheCtx(ctx, &resp, query, uuid)\n")
	} else {
		buf.WriteString("err := m.conn.QueryRowCtx(ctx, &resp, query, uuid)\n")
	}
	buf.WriteString(`switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}`)
	buf.WriteString("\n}\n")
	return buf.String()
}

func (t *FindByUuid) FullName() string {
	return t.fullName
}
