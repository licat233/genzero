package funcs

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/licat233/genzero/sql"
	"github.com/licat233/genzero/tools"
)

type FindsByAny struct {
	field       sql.Field
	modelName   string
	name        string
	req         string
	resp        string
	fullName    string
	IsCacheMode bool
	Table       *sql.Table
}

var _ ModelFunc = (*FindsByAny)(nil)

type FindsByAnyCollection []*FindsByAny

var _ ModelFunc = (FindsByAnyCollection)(nil)

func NewFindsByAnyCollection(table *sql.Table, isCacheMode bool) FindsByAnyCollection {
	var res FindsByAnyCollection = make([]*FindsByAny, 0)
	for _, field := range table.Fields {
		if strings.ToLower(field.Name) == "id" {
			continue
		}
		res = append(res, NewFindsByAny(table, isCacheMode, field))
	}
	return res
}

func (f FindsByAnyCollection) String() string {
	var buf = new(bytes.Buffer)
	for _, findByAny := range f {
		buf.WriteString(findByAny.String() + "\n")
	}
	return buf.String()
}

func (f FindsByAnyCollection) FullName() string {
	var buf = new(bytes.Buffer)
	for _, findByAny := range f {
		buf.WriteString(findByAny.FullName() + "\n")
	}
	return buf.String()
}

func NewFindsByAny(t *sql.Table, isCacheMode bool, field sql.Field) *FindsByAny {
	modelName := modelName(t.Name)
	name := fmt.Sprintf("FindsBy%s", field.UpperCamelCaseName)
	req := fmt.Sprintf("ctx context.Context, %s %s", tools.ToLowerCamel(field.Name), field.Type)
	resp := fmt.Sprintf("([]*%s, error)", tools.ToCamel(t.Name))
	fullName := fmt.Sprintf("%s(%s) %s", name, req, resp)
	return &FindsByAny{
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

func (s *FindsByAny) String() string {
	var buf = new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("\nfunc (m *%s) %s {\n", s.modelName, s.fullName))
	buf.WriteString(fmt.Sprintf("var resp = make([]*%s, 0)\n", tools.ToCamel(s.Table.Name)))
	if delField := s.Table.GetIsDeletedField(); delField != nil {
		buf.WriteString("query := fmt.Sprintf(\"select %s from %s where `" + s.field.Name + "` = ? and `" + delField.Name + "` = '0' \", " + tools.ToLowerCamel(s.Table.Name) + "Rows, m.table)\n")
	} else {
		buf.WriteString("query := fmt.Sprintf(\"select %s from %s where `" + s.field.Name + "` = ? \", " + tools.ToLowerCamel(s.Table.Name) + "Rows, m.table)\n")
	}

	fieldV := tools.ToLowerCamel(s.field.Name)
	if s.field.Type == "time.Time" {
		fieldV = fieldV + ".Format(\"2006-01-02 15:04:05\")"
	}

	if s.IsCacheMode {
		buf.WriteString("err := m.QueryRowNoCacheCtx(ctx, &resp, query, " + fieldV + ")\n")
	} else {
		buf.WriteString("err := m.conn.QueryRowCtx(ctx, &resp, query, " + fieldV + ")\n")
	}
	buf.WriteString("return resp, err")
	buf.WriteString("\n}\n")
	return buf.String()
}

func (t *FindsByAny) FullName() string {
	return t.fullName
}
