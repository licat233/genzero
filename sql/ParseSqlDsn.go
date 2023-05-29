package sql

import (
	"database/sql"
	"regexp"
	"strings"

	"github.com/chuckpreslar/inflect"
	_ "github.com/go-sql-driver/mysql"
	"github.com/licat233/genzero/config"
	"github.com/licat233/genzero/tools"
)

func ParseSqlDsn(dsn string) (*Schema, error) {
	Conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = Conn.Ping(); err != nil {
		return nil, err
	}

	var tables []string
	rows, err := Conn.Query("show tables")
	if err != nil {
		return nil, err
	}

	mustFilter, checkColumnStr := TableFilterCondition()
	for rows.Next() {
		var table string
		err := rows.Scan(&table)
		if err != nil {
			return nil, err
		}
		if tools.HasInSlice(config.C.DB.IgnoreTables, table) {
			continue
		}
		if mustFilter {
			if !tools.HasInSlice(checkColumnStr, table) {
				continue
			}
		}

		tables = append(tables, table)
	}

	schema := &Schema{}
	var schemaName string

	err = Conn.QueryRow("SELECT SCHEMA()").Scan(&schemaName)
	if err != nil {
		return nil, err
	}

	schema.Name = schemaName

	tablesStr := strings.Join(tables, ",")
	columns, err := DbColumns(Conn, schema.Name, tablesStr)
	if err != nil {
		return nil, err
	}

	tableMap := map[string]*Table{}

	for _, column := range columns {

		if tools.HasInSlice(config.C.DB.IgnoreColumns, column.ColumnName) {
			continue
		}

		tableName := column.TableName
		fieldType := column.ColumnType
		fieldName := column.ColumnName
		fieldComment := column.ColumnComment

		table, ok := tableMap[tableName]
		if !ok {
			table = &Table{
				Name:    tableName,
				Comment: column.TableComment,
				Fields:  []Field{},
				Enums:   []Enum{},
			}
			table.Name = tableName
			tableMap[tableName] = table
		}

		if fieldType == "_enum" || fieldType == "set" {
			enumList := regexp.MustCompile(`[_enum|set]\((.+?)\)`).FindStringSubmatch(fieldType)
			enums := strings.FieldsFunc(enumList[1], func(c rune) bool {
				cs := string(c)
				return cs == "," || cs == "'"
			})

			enumName := inflect.Singularize(tools.ToCamel(tableName)) + tools.ToCamel(fieldName)
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
			UpperCamelCaseName: tools.ToCamel(fieldName),
			Type:               goType(fieldType),
			Comment:            fieldComment,
			DefaultValue:       "",
			Tag:                "db",
			Nullable:           column.IsNullable == "YES",
		}
		table.Fields = append(table.Fields, field)
	}

	for _, column := range columns {
		if column.ColumnName == "is_deleted" {
			table, ok := tableMap[column.TableName]
			if ok {
				table.HasDeleteFiled = true
			}
		}

		if strings.ToLower(column.ColumnName) == "uuid" {
			table, ok := tableMap[column.TableName]
			if ok {
				table.HasUuid = true
			}
		}
	}

	for _, table := range tableMap {
		schema.Tables = append(schema.Tables, *table)
	}

	return schema, nil
}

func DbColumns(db *sql.DB, schema, table string) ([]*Column, error) {

	tableArr := strings.Split(strings.Trim(table, ","), ",")

	q := "SELECT c.TABLE_NAME, c.COLUMN_NAME, c.IS_NULLABLE, c.DATA_TYPE, " +
		"c.CHARACTER_MAXIMUM_LENGTH, c.NUMERIC_PRECISION, c.NUMERIC_SCALE, c.COLUMN_TYPE ,c.COLUMN_COMMENT,t.TABLE_COMMENT " +
		"FROM INFORMATION_SCHEMA.COLUMNS as c  LEFT JOIN  INFORMATION_SCHEMA.TABLES as t  on c.TABLE_NAME = t.TABLE_NAME and  c.TABLE_SCHEMA = t.TABLE_SCHEMA" +
		" WHERE c.TABLE_SCHEMA = ?"

	if table != "" && table != "*" {
		q += " AND c.TABLE_NAME IN('" + strings.TrimRight(strings.Join(tableArr, "' ,'"), ",") + "')"
	}

	q += " ORDER BY c.TABLE_NAME, c.ORDINAL_POSITION"

	rows, err := db.Query(q, schema)
	defer func() {
		rows.Close()
	}()
	if nil != err {
		return nil, err
	}

	cols := []*Column{}

	for rows.Next() {
		cs := &Column{}
		err := rows.Scan(&cs.TableName, &cs.ColumnName, &cs.IsNullable, &cs.DataType,
			&cs.CharacterMaximumLength, &cs.NumericPrecision, &cs.NumericScale, &cs.ColumnType, &cs.ColumnComment, &cs.TableComment)
		if err != nil {
			return nil, err
		}
		typName := cs.ColumnType
		re := regexp.MustCompile(`\(\d*\)`)
		typName = re.ReplaceAllString(typName, "")
		cs.ColumnType = typName
		cols = append(cols, cs)
	}
	if err := rows.Err(); nil != err {
		return nil, err
	}

	return cols, nil
}
