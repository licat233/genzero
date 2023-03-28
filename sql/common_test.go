package sql

import (
	"fmt"
	"testing"
)

func TestPick(t *testing.T) {
	line := "`id` INT NOT NULL AUTO_INCREMENT COMMENT '表主键',"
	res := PickFieldType(line)
	fmt.Println(res)
}
