package utils

import (
	"strings"

	"github.com/licat233/genzero/sql"
)

// 过滤需要的表，以及需要忽略的表，返回需要的表
func FilterTables(source sql.TableCollection, checkTables []string, ignoreTables []string) sql.TableCollection {
	if len(checkTables) == 0 {
		return filterIgnoreTables(source, ignoreTables)
	}
	if checkTables[0] == "*" {
		return filterIgnoreTables(source, ignoreTables)
	}
	tablesMap := make(map[string]bool)
	for _, t := range checkTables {
		tablesMap[t] = true
	}
	var result sql.TableCollection
	for _, t := range source {
		if _, ok := tablesMap[t.Name]; ok {
			result = append(result, t)
		}
	}
	return filterIgnoreTables(result, ignoreTables)
}

// 过滤掉不需要的表，返回剩余的表
func filterIgnoreTables(source sql.TableCollection, ignoreTables []string) sql.TableCollection {
	if len(ignoreTables) == 0 {
		return source
	}
	if ignoreTables[0] == "*" {
		return []sql.Table{}
	}
	tablesMap := make(map[string]bool)
	for _, t := range ignoreTables {
		tablesMap[t] = true
	}
	var result sql.TableCollection
	for _, t := range source {
		if _, ok := tablesMap[t.Name]; !ok {
			result = append(result, t)
		}
	}
	return result
}

// 过滤掉不需要的字段，返回剩余的字段
func FilterIgnoreFields(source sql.FieldCollection, ignoreFields []string) sql.FieldCollection {
	if len(ignoreFields) == 0 {
		return source
	}
	if ignoreFields[0] == "*" {
		return []sql.Field{}
	}
	fieldsMap := make(map[string]bool)
	for _, t := range ignoreFields {
		fieldsMap[t] = true
	}
	var result sql.FieldCollection
	for _, t := range source {
		if _, ok := fieldsMap[t.Name]; !ok {
			result = append(result, t)
		}
	}
	return result
}

// 合并两个slice，返回合并后的slice
func MergeSlice(slice1, slice2 []string) []string {
	m := map[string]bool{}
	for _, v := range slice1 {
		m[v] = true
	}
	for _, v := range slice2 {
		m[v] = true
	}
	result := make([]string, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

// slice是否包含v
func SliceContains(slice []string, v string) bool {
	for _, s := range slice {
		if s == v {
			return true
		}
	}
	return false
}

func HasUuid(fields sql.FieldCollection) bool {
	for _, field := range fields {
		if strings.ToLower(field.Name) == "uuid" {
			return true
		}
	}
	return false
}

func HasName(fields sql.FieldCollection) bool {
	for _, field := range fields {
		if strings.ToLower(field.Name) == "name" {
			return true
		}
	}
	return false
}

func HasIsDeleted(fields sql.FieldCollection) bool {
	for _, field := range fields {
		if strings.ToLower(field.Name) == "is_deleted" {
			return true
		}
	}
	return false
}
