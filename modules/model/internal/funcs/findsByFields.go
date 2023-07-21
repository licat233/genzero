package funcs

import (
	"bytes"
	"fmt"

	"github.com/licat233/genzero/sql"
	"github.com/licat233/genzero/tools"
)

type findsByFields struct {
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

var _ ModelFunc = (*findsByFields)(nil)

type FindsByFieldsCollection []*findsByFields

var _ ModelFunc = (FindsByFieldsCollection)(nil)

func NewFindsByFieldsCollection(table *sql.Table, isCacheMode bool) FindsByFieldsCollection {
	var res FindsByFieldsCollection = make([]*findsByFields, 0)
	for _, field := range table.GetFields() {
		res = append(res, newFindsByFields(table, isCacheMode, field))
	}
	return res
}

func (f FindsByFieldsCollection) String() string {
	var buf = new(bytes.Buffer)
	for _, findByFields := range f {
		buf.WriteString(findByFields.String() + "\n")
	}
	return buf.String()
}

func (f FindsByFieldsCollection) FullName() string {
	var buf = new(bytes.Buffer)
	for _, findByFields := range f {
		buf.WriteString(findByFields.FullName() + "\n")
	}
	return buf.String()
}

func newFindsByFields(t *sql.Table, isCacheMode bool, field sql.Field) *findsByFields {
	modelName := modelName(t.Name)
	reqFieldName := tools.ToLowerCamel(tools.PluralizedName(field.Name))
	name := fmt.Sprintf("FindsBy%s", tools.ToCamel(reqFieldName))
	req := fmt.Sprintf("ctx context.Context, %s []%s", reqFieldName, field.Type)
	resp := fmt.Sprintf("([]*%s, error)", tools.ToCamel(t.Name))
	fullName := fmt.Sprintf("%s(%s) %s", name, req, resp)
	return &findsByFields{
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

func (s *findsByFields) String() string {
	var buf = new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("\nfunc (m *%s) %s {\n", s.modelName, s.fullName))
	buf.WriteString(fmt.Sprintf("var resp = make([]*%s, 0)\n", tools.ToCamel(s.Table.Name)))
	buf.WriteString(fmt.Sprintf("if len(%s) == 0 {\n", s.reqFieldName))
	buf.WriteString("return resp, nil\n")
	buf.WriteString("}\n")
	buf.WriteString("strs := []string{}\n")
	buf.WriteString("for _, v := range " + s.reqFieldName + " {\n")
	switch s.field.Type {
	case "time.Time":
		buf.WriteString("strs = append(strs, v.Format(\"2006-01-02 15:04:05\"))\n")
	case "string":
		buf.WriteString("strs = append(strs, v)\n")
	case "int64":
		buf.WriteString("strs = append(strs, strconv.FormatInt(v,10))\n")
	case "int":
		buf.WriteString("strs = append(strs, strconv.Itoa(v))\n")
	default:
		buf.WriteString("strs = append(strs, fmt.Sprint(v))\n")
	}
	buf.WriteString("}\n")
	if delField := s.Table.GetIsDeletedField(); delField != nil {
		buf.WriteString("query := fmt.Sprintf(\"select %s from %s where `" + s.field.Name + "` in (?) and `" + delField.Name + "` = '0' \", " + tools.ToLowerCamel(s.Table.Name) + "Rows, m.table)\n")
	} else {
		buf.WriteString("query := fmt.Sprintf(\"select %s from %s where `" + s.field.Name + "` in (?)\", " + tools.ToLowerCamel(s.Table.Name) + "Rows, m.table)\n")
	}

	buf.WriteString("agr := strings.Join(strs, \",\")\n")
	if s.IsCacheMode {
		buf.WriteString("err := m.QueryRowsNoCacheCtx(ctx, &resp, query, agr)\n")
	} else {
		buf.WriteString("err := m.conn.QueryRowsCtx(ctx, &resp, query, agr)\n")
	}
	buf.WriteString("return resp, err")
	buf.WriteString("\n}\n")
	return buf.String()
}

func (t *findsByFields) FullName() string {
	return t.fullName
}
