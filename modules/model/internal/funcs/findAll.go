package funcs

import (
	"bytes"
	"fmt"

	"github.com/licat233/genzero/sql"
	"github.com/licat233/genzero/tools"
)

type FindAll struct {
	modelName   string
	name        string
	req         string
	resp        string
	fullName    string
	IsCacheMode bool
	Table       *sql.Table
}

var _ ModelFunc = (*FindAll)(nil)

func NewFindAll(t *sql.Table, isCacheMode bool) *FindAll {
	modelName := modelName(t.Name)
	name := "FindAll"
	req := "ctx context.Context"
	resp := fmt.Sprintf("([]*%s, error)", tools.ToCamel(t.Name))
	fullName := fmt.Sprintf("%s(%s) %s", name, req, resp)
	return &FindAll{
		modelName:   modelName,
		name:        name,
		req:         req,
		resp:        resp,
		fullName:    fullName,
		IsCacheMode: isCacheMode,
		Table:       t,
	}
}

func (s *FindAll) String() string {
	var buf = new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("\nfunc (m *%s) %s {\n", s.modelName, s.fullName))
	// buf.WriteString("uuidKey := m.formatUuidKey(uuid)\n")
	buf.WriteString(fmt.Sprintf("var resp = make([]*%s, 0)\n", tools.ToCamel(s.Table.Name)))
	if delField := s.Table.GetIsDeletedField(); delField != nil {
		buf.WriteString("query := fmt.Sprintf(\"select %s from %s where `" + delField.Name + "` = '0' limit 99999\", " + tools.ToLowerCamel(s.Table.Name) + "Rows, m.table)\n")
	} else {
		buf.WriteString("query := fmt.Sprintf(\"select %s from %s limit 99999\", " + tools.ToLowerCamel(s.Table.Name) + "Rows, m.table)\n")
	}

	if s.IsCacheMode {
		buf.WriteString("err := m.QueryRowsNoCacheCtx(ctx, &resp, query)\n")
	} else {
		buf.WriteString("err := m.conn.QueryRowsCtx(ctx, &resp, query)\n")
	}
	buf.WriteString(`switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}`)
	buf.WriteString("\n}\n")
	return buf.String()
}

func (t *FindAll) FullName() string {
	return t.fullName
}
