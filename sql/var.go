package sql

import (
	"database/sql"

	"github.com/licat233/genzero/tools"
)

type Schema struct {
	Name    string
	Comment string
	Tables  TableCollection
}

func (s *Schema) Copy() *Schema {
	return &Schema{
		Name:    s.Name,
		Comment: s.Comment,
		Tables:  s.Tables.Copy(),
	}
}

type Table struct {
	Name    string `sql:"name"`
	Comment string `sql:"comment"`
	// HasDeleteField bool   `sql:"is_deleted"`
	// HasUuid        bool   `sql:"uuid"`

	Fields FieldCollection `sql:"fields"`
	Enums  EnumCollection  `sql:"enums"`
}

func (t *Table) Copy() *Table {
	return &Table{
		Name:    t.Name,
		Comment: t.Comment,
		// HasDeleteField: t.HasDeleteField,
		// HasUuid:        t.HasUuid,
		Fields: t.Fields.Copy(),
		Enums:  t.Enums.Copy(),
	}
}

func (t *Table) ExistField(fieldName string) bool {
	for _, field := range t.Fields {
		if field.Name == fieldName {
			return true
		}
	}
	return false
}

func (t *Table) GetIsDeletedField() *Field {
	for _, field := range t.Fields {
		snake_name := tools.ToSnake(field.Name)
		if tools.HasInSlice(DelFieldNames, snake_name) {
			return &field
		}
	}
	return nil
}

func (t *Table) ExistIsDelField() bool {
	return t.GetIsDeletedField() != nil
}

func (t *Table) GetUuidField() *Field {
	for _, field := range t.Fields {
		snake_name := tools.ToSnake(field.Name)
		if tools.HasInSlice(UuidFieldNames, snake_name) {
			return &field
		}
	}
	return nil
}

func (t *Table) ExistUuidField() bool {
	return t.GetUuidField() != nil
}

func (t *Table) GetDelAtField() *Field {
	for _, field := range t.Fields {
		snake_name := tools.ToSnake(field.Name)
		if tools.HasInSlice(DelAtFieldNames, snake_name) {
			return &field
		}
	}
	return nil
}

func (t *Table) ExistDelAtField() bool {
	return t.GetDelAtField() != nil
}

type TableCollection []Table

func (t TableCollection) Copy() TableCollection {
	res := make(TableCollection, len(t))
	for i, v := range t {
		res[i] = *v.Copy()
	}
	return res
}

type Field struct {
	Primary            bool   `json:"primary"`
	Name               string `json:"name"`
	UpperCamelCaseName string `json:"upper_camel_case_name"`
	Type               string `json:"type"`
	Comment            string `json:"comment"`
	DefaultValue       string `json:"default_value"`
	Tag                string `json:"tag"`
	Nullable           bool   `json:"nullable"`
}

func (f *Field) Copy() *Field {
	return &Field{
		Primary:            f.Primary,
		Name:               f.Name,
		UpperCamelCaseName: f.UpperCamelCaseName,
		Type:               f.Type,
		Comment:            f.Comment,
		DefaultValue:       f.DefaultValue,
		Tag:                f.Tag,
		Nullable:           f.Nullable,
	}
}

type FieldCollection []Field

func (f FieldCollection) Copy() FieldCollection {
	res := make(FieldCollection, len(f))
	for i, v := range f {
		res[i] = *v.Copy()
	}
	return res
}

// Column represents a database column.
type Column struct {
	Style                  string
	TableName              string
	TableComment           string
	ColumnName             string
	IsNullable             string
	DataType               string
	CharacterMaximumLength sql.NullInt64
	NumericPrecision       sql.NullInt64
	NumericScale           sql.NullInt64
	ColumnType             string
	ColumnComment          string
}

// map for converting mysql type to golang types
var TypeForMysqlToGo = map[string]string{
	"int":                "int64",
	"integer":            "int64",
	"tinyint":            "int64",
	"smallint":           "int64",
	"mediumint":          "int64",
	"bigint":             "int64",
	"int unsigned":       "int64",
	"integer unsigned":   "int64",
	"tinyint unsigned":   "int64",
	"smallint unsigned":  "int64",
	"mediumint unsigned": "int64",
	"bigint unsigned":    "int64",
	"bit":                "int64",
	"bool":               "bool",
	"enum":               "string",
	"set":                "string",
	"varchar":            "string",
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"longtext":           "string",
	"blob":               "string",
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"date":               "time.Time", // time.Time or string
	"datetime":           "time.Time", // time.Time or string
	"timestamp":          "time.Time", // time.Time or string
	"time":               "time.Time", // time.Time or string
	"float":              "float64",
	"double":             "float64",
	"decimal":            "float64",
	"binary":             "string",
	"varbinary":          "string",
}
