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
	var isBlockNote bool
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "/*") {
			isBlockNote = true
		}
		if isBlockNote && strings.HasSuffix(line, "*/") {
			isBlockNote = false
			continue
		}
		if isBlockNote {
			continue
		}
		if strings.HasPrefix(line, "--") {
			// Ignore comments
			continue
		}
		chips := strings.Split(line, " ")
		if len(chips) < 2 {
			continue
		}
		if strings.Contains(line, "CREATE TABLE") {
			if len(chips) < 3 {
				continue
			}
			tableName := strings.Trim(chips[2], "`")
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

		if ok, _ := MutipleStringSubmatch(line, `^\)?\sENGINE\s?=`, `^\)?\sengine\s?=`); ok {
			if err != nil {
				tools.Error(err.Error())
			}
			comment := PickTableComment(line)
			table.Comment = comment
			schema.Tables = append(schema.Tables, *table)
			table = nil
			currentTableName = ""
			continue
		}
		if fieldName := PickFieldName(line); fieldName != "" {
			if has := tools.ToSnake(fieldName) == "is_deleted"; has {
				table.HasDeleteFiled = true
			}

			if has := strings.ToLower(fieldName) == "uuid"; has {
				table.HasUuid = true
			}

			if tools.HasInSlice(config.C.DB.IgnoreColumns, fieldName) {
				continue
			}

			fieldComment := PickFieldComment(line)
			fieldType := PickFieldType(line)
			fieldType = strings.ToLower(fieldType)
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
				Name:               fieldName,
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

			field.DefaultValue = PickFieldDefaultValue(line)

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
