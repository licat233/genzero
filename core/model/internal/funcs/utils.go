package funcs

import (
	"fmt"

	"github.com/licat233/genzero/tools"
)

func modelName(tableName string) string {
	return fmt.Sprintf("default%sModel", tools.ToCamel(tableName))
}
