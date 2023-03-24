package funcs

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/licat233/genzero/core/model/conf"
	"github.com/licat233/genzero/sql"
	"github.com/licat233/genzero/tools"
)

type FindList struct {
	modelName string
	name      string
	req       string
	resp      string
	fullname  string
	Table     *sql.Table
}

var _ ModelFunc = (*FindList)(nil)

func NewFindList(t *sql.Table) FindList {
	camelName := tools.ToCamel(t.Name)
	lowerName := tools.ToLowerCamel(t.Name)
	modelName := modelName(t.Name)
	name := "FindList"
	req := fmt.Sprintf("ctx context.Context, pageSize, page int64, keyword string, %s *%s", lowerName, camelName)
	resp := fmt.Sprintf("(resp []*%s, total int64, err error)", camelName)
	fullName := fmt.Sprintf("%s(%s) %s", name, req, resp)
	return FindList{
		modelName: modelName,
		name:      name,
		req:       req,
		resp:      resp,
		fullname:  fullName,
		Table:     t,
	}
}

func (s FindList) String() string {
	hasName := s.hasName()

	camelName := tools.ToCamel(s.Table.Name)
	lowerName := tools.ToLowerCamel(s.Table.Name)
	var buf = new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("\nfunc (m *%s) %s {", s.modelName, s.fullname))
	if hasName {
		buf.WriteString("\n\thasName := false")
	}
	if s.Table.HasDeleteFiled {
		buf.WriteString(fmt.Sprintf("\n\tsq := squirrel.Select(%sRows).From(m.table).Where(\"`is_deleted`= '0'\")", lowerName))
	} else {
		buf.WriteString(fmt.Sprintf("\n\tsq := squirrel.Select(%sRows).From(m.table)", lowerName))
	}

	s.thanString(buf)

	buf.WriteString("\n\tif pageSize > 0 && page > 0 {")
	buf.WriteString("\n\t\tsqCount := sq.RemoveLimit().RemoveOffset()")
	buf.WriteString("\n\t\tsq = sq.Limit(uint64(pageSize)).Offset(uint64((page - 1) * pageSize))")
	buf.WriteString("\n\t\tqueryCount, agrsCount, e := sqCount.ToSql()")
	buf.WriteString("\n\t\tif e != nil {\n\t\t\terr = e\n\t\t\treturn\n\t\t}")
	buf.WriteString(fmt.Sprintf("\n\t\tqueryCount = strings.ReplaceAll(queryCount, %sRows, \"COUNT(*)\")", lowerName))
	if conf.IsCacheMode {
		buf.WriteString("\n\t\tif err = m.QueryRowNoCacheCtx(ctx, &total, queryCount, agrsCount...); err != nil {\n\t\t\treturn\n\t\t}")
	} else {
		buf.WriteString("\n\t\tif err = m.conn.QueryRowCtx(ctx, &total, queryCount, agrsCount...); err != nil {\n\t\t\treturn\n\t\t}")
	}

	buf.WriteString("\n\t}")
	buf.WriteString("\n\tquery, agrs, err := sq.ToSql()\n\tif err != nil {\n\t\treturn\n\t}")
	buf.WriteString(fmt.Sprintf("\n\tresp = make([]*%s, 0)", camelName))
	if conf.IsCacheMode {
		buf.WriteString("\n\tif err = m.QueryRowsNoCacheCtx(ctx, &resp, query, agrs...); err != nil {\n\t\treturn\n\t}")
	} else {
		buf.WriteString("\n\tif err = m.conn.QueryRowsCtx(ctx, &resp, query, agrs...); err != nil {\n\t\treturn\n\t}")
	}

	buf.WriteString("\n\treturn")
	buf.WriteString("\n}\n")
	return buf.String()
}

func (s FindList) FullName() string {
	return s.fullname
}

func (s FindList) Req() string {
	return s.req
}

func (s FindList) Resp() string {
	return s.resp
}

func (s FindList) Name() string {
	return s.name
}

func (s FindList) ModelName() string {
	return s.modelName
}

func (s FindList) hasName() bool {
	for _, field := range s.Table.Fields {
		if isNameField(field) {
			return true
		}
	}
	return false
}

func (s FindList) thanString(buf *bytes.Buffer) {
	lowerName := tools.ToLowerCamel(s.Table.Name)

	hasName := false
	buf.WriteString(fmt.Sprintf("\n\tif %s != nil {", lowerName))
	for _, field := range s.Table.Fields {
		var condition string
		fieldString := fmt.Sprintf("%s.%s", lowerName, tools.ToCamel(field.Name))
		//判断是字符串，还是数字
		if field.Type == "int64" || field.Type == "float64" || field.Type == "int" {
			if tools.IsIdColumn(field.Name) {
				condition = fieldString + " > 0"
			} else {
				condition = fieldString + " >= 0"
			}
		} else if field.Type == "string" {
			condition = fieldString + " != \"\""
		} else if field.Type == "time.Time" {
			condition = fieldString + ".IsZero()"
		} else {
			tools.Warning("unknow column type:", s.Table.Name, "-", field.Name, "-", field.Type)
			continue
		}
		buf.WriteString(fmt.Sprintf("\n\t\tif %s {", condition))
		buf.WriteString(fmt.Sprintf("\n\t\t\tsq = sq.Where(\"%s = ?\", %s)", field.Name, fieldString))
		if isNameField(field) && !hasName {
			buf.WriteString("\n\t\t\thasName = true")
			hasName = true
		}
		buf.WriteString("\n\t\t}")
	}
	buf.WriteString("\n\t}")
	if hasName {
		buf.WriteString("\n\tif keyword != \"\" && !hasName {")
		buf.WriteString("\n\t\tsq = sq.Where(\"name LIKE ?\", fmt.Sprintf(\"%%%s%%\", keyword))")
		buf.WriteString("\n\t}")
	}
}

func isNameField(field sql.Field) bool {
	return strings.ToLower(field.Name) == "name" && field.Type == "string"
}
