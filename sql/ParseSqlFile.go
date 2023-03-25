package sql

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"github.com/chuckpreslar/inflect"
	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/tools"
)

func ParseSqlFile(filename string) (*Schema, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	schema := &Schema{}

	var table *Table

	scanner := bufio.NewScanner(file)

	mustFilter, checkColumnStr := TableFilterCondition()
	var currentTableName string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "--") {
			// Ignore comments
			continue
		}
		if strings.Contains(line, "CREATE TABLE") {
			tableName := strings.Split(line, "`")[1]
			currentTableName = tableName
			if tools.HasInSlice(config.C.DB.IgnoreTables, tableName) {
				continue
			}
			if mustFilter {
				if !tools.HasInSlice(checkColumnStr, tableName) {
					continue
				}
			}

			table = &Table{}
			table.Name = tableName
			continue
		}
		if table == nil {
			continue
		}
		if strings.Contains(line, "ENGINE") && strings.Contains(line, "COMMENT=") {
			comment := strings.Trim(strings.Split(line, "COMMENT='")[1], "'")
			comment = strings.TrimSuffix(comment, "';")
			table.Comment = comment
			schema.Tables = append(schema.Tables, *table)
			table = nil
			currentTableName = ""
			continue
		}
		if line[0] == '`' {
			fieldName := strings.Split(line, "`")[1]

			has := tools.ToSnake(fieldName) == "is_deleted"
			if has {
				table.HasDeleteFiled = has
			}

			if tools.HasInSlice(config.C.DB.IgnoreColumns, fieldName) {
				continue
			}

			if !strings.Contains(line, " ") {
				//不存在类型，忽略改行
				continue
			}

			fieldComment := ""
			if strings.Contains(line, "COMMENT") {
				fieldComment = strings.Trim(strings.Split(line, "COMMENT '")[1], "',")
			}

			fieldType := strings.TrimRightFunc(strings.Split(line, " ")[1], func(r rune) bool {
				return r < 'a' || r > 'z'
			})
			if fieldType == "_enum" || fieldType == "set" {
				enumList := regexp.MustCompile(`[_enum|set]\((.+?)\)`).FindStringSubmatch(fieldType)
				if len(enumList) < 2 {
					continue
				}
				enums := strings.FieldsFunc(enumList[1], func(c rune) bool {
					cs := string(c)
					return cs == "," || cs == "'"
				})

				enumName := inflect.Singularize(tools.ToCamel(currentTableName)) + tools.ToCamel(fieldName)
				enum, err := NewEnumFromStrings(enumName, fieldComment, enums)
				if nil != err {
					return nil, err
				}
				table.Enums = append(table.Enums, *enum)
				continue
			}

			field := Field{
				Primary:            false,
				Name:               currentTableName,
				UpperCamelCaseName: "",
				Type:               fieldType,
				Comment:            fieldComment,
				DefaultValue:       "",
				Tag:                "db",
				Nullable:           false,
			}

			field.Name = fieldName
			field.UpperCamelCaseName = tools.ToCamel(field.Name)

			field.Type = goType(fieldType)

			if strings.Contains(line, "DEFAULT'") {
				field.DefaultValue = strings.Split(line, "DEFAULT ")[1]
			}
			field.Primary = strings.Contains(line, "PRIMARY KEY")
			field.Nullable = !strings.Contains(line, "NOT NULL")

			table.Fields = append(table.Fields, field)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return schema, nil
}
