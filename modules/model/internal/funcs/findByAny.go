package funcs

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/licat233/genzero/sql"
	"github.com/licat233/genzero/tools"
)

type FindByAny struct {
	field       sql.Field
	modelName   string
	name        string
	req         string
	resp        string
	fullName    string
	IsCacheMode bool
	Table       *sql.Table
}

var _ ModelFunc = (*FindByAny)(nil)

type FindByAnyCollection []*FindByAny

var _ ModelFunc = (FindByAnyCollection)(nil)

func NewFindByAnyCollection(table *sql.Table, isCacheMode bool) FindByAnyCollection {
	var res FindByAnyCollection = make([]*FindByAny, 0)
	for _, field := range table.GetFields() {
		// if strings.ToLower(field.Name) == "id" {
		// 	continue
		// }
		res = append(res, NewFindByAny(table, isCacheMode, field))
	}
	return res
}

func (f FindByAnyCollection) String() string {
	var buf = new(bytes.Buffer)
	for _, findByAny := range f {
		buf.WriteString(findByAny.String() + "\n")
	}
	return buf.String()
}

func (f FindByAnyCollection) FullName() string {
	var list []string
	for _, findByAny := range f {
		list = append(list, findByAny.FullName())
	}
	return strings.Join(list, "\n")
}

func NewFindByAny(t *sql.Table, isCacheMode bool, field sql.Field) *FindByAny {
	modelName := modelName(t.Name)
	name := fmt.Sprintf("FindBy%s", field.UpperCamelCaseName)
	req := fmt.Sprintf("ctx context.Context, %s %s", tools.ToLowerCamel(field.Name), field.Type)
	resp := fmt.Sprintf("(*%s, error)", tools.ToCamel(t.Name))
	fullName := fmt.Sprintf("%s(%s) %s", name, req, resp)
	return &FindByAny{
		field:       field,
		modelName:   modelName,
		name:        name,
		req:         req,
		resp:        resp,
		fullName:    fullName,
		IsCacheMode: isCacheMode,
		Table:       t,
	}
}

func (s *FindByAny) String() string {
	var buf = new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("\nfunc (m *%s) %s {\n", s.modelName, s.fullName))
	buf.WriteString(fmt.Sprintf("var resp %s\n", tools.ToCamel(s.Table.Name)))
	if delField := s.Table.GetIsDeletedField(); delField != nil {
		buf.WriteString("query := fmt.Sprintf(\"select %s from %s where `" + s.field.Name + "` = ? and `" + delField.Name + "` = '0' limit 1\", " + tools.ToLowerCamel(s.Table.Name) + "Rows, m.table)\n")
	} else {
		buf.WriteString("query := fmt.Sprintf(\"select %s from %s where `" + s.field.Name + "` = ? limit 1\", " + tools.ToLowerCamel(s.Table.Name) + "Rows, m.table)\n")
	}

	fieldV := tools.ToLowerCamel(s.field.Name)
	if s.field.Type == "time.Time" {
		fieldV = fieldV + ".Format(\"2006-01-02 15:04:05\")"
	}

	if s.IsCacheMode {
		if strings.ToLower(s.field.Name) == "id" {
			buf.WriteString(fmt.Sprintf("%sIdKey := fmt.Sprintf(\"%%s%%v\", cache%sIdPrefix, id)\n", tools.ToLowerCamel(s.Table.Name), tools.ToCamel(s.Table.Name)))
			buf.WriteString(fmt.Sprintf("err := m.QueryRowCtx(ctx, &resp, %sIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {\n", tools.ToLowerCamel(s.Table.Name)))
			buf.WriteString("return conn.QueryRowCtx(ctx, v, query, id)\n")
			buf.WriteString("})\n")
		} else {
			buf.WriteString("err := m.QueryRowNoCacheCtx(ctx, &resp, query, " + fieldV + ")\n")
		}
	} else {
		buf.WriteString("err := m.conn.QueryRowCtx(ctx, &resp, query, " + fieldV + ")\n")
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

func (t *FindByAny) FullName() string {
	return t.fullName
}
