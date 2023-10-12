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
			if tools.SliceContain(config.C.DB.IgnoreTables, tableName) {
				continue
			}
			if mustFilter {
				if !tools.SliceContain(checkColumnStr, tableName) {
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
			// if tools.HasInSlice(config.C.DB.IgnoreColumns, fieldName) {
			// 	continue
			// }

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

			_type := goType(fieldType)
			if name := strings.ToLower(fieldType); name == "longtext" || name == "text" {
				if strings.Contains(strings.ToLower(line), " not null ") {
					_type = "string"
				}
			}

			field := Field{
				Primary:            strings.Contains(line, "PRIMARY KEY"),
				Name:               fieldName,
				UpperCamelCaseName: tools.ToCamel(fieldName),
				Type:               _type,
				RawType:            fieldType,
				Comment:            fieldComment,
				DefaultValue:       PickFieldDefaultValue(line),
				Tag:                "db",
				Nullable:           !strings.Contains(line, "NOT NULL"),
				Hide:               IsIgnoreField(fieldName),
			}

			table.Fields = append(table.Fields, field)
		}

		if fieldName := PickUniqueFieldName(line); fieldName != "" {
			for i := range table.Fields {
				if strings.EqualFold(table.Fields[i].Name, fieldName) {
					table.Fields[i].Unique = true
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return schema, nil
}
