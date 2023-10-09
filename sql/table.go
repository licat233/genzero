package sql

import (
	"database/sql"
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
	Name    string          `sql:"name"`
	Comment string          `sql:"comment"`
	Fields  FieldCollection `sql:"fields"`
	Enums   EnumCollection  `sql:"enums"`
}

func (t *Table) Copy() *Table {
	return &Table{
		Name:    t.Name,
		Comment: t.Comment,
		Fields:  t.Fields.Copy(),
		Enums:   t.Enums.Copy(),
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
		if IsDeleteField(field.Name) {
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
		if IsUuidField(field.Name) {
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
		if IsDelAtField(field.Name) {
			return &field
		}
	}
	return nil
}

func (t *Table) ExistDelAtField() bool {
	return t.GetDelAtField() != nil
}

func (t *Table) GetFields() FieldCollection {
	res := []Field{}
	for _, field := range t.Fields {
		if field.Hide {
			continue
		}
		res = append(res, field)
	}
	return res
}

func (t *Table) ExistNameField() bool {
	return t.GetNameField() != nil
}

func (t *Table) GetNameField() *Field {
	for _, field := range t.Fields {
		if IsNameField(field.Name) {
			return &field
		}
	}
	return nil
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
	Hide               bool   `json:"hide"`
	Unique             bool   `json:"unique"`
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
		Hide:               f.Hide,
		Unique:             f.Unique,
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
	"text":               "sql.NullString",
	"longtext":           "sql.NullString",
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
