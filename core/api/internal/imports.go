package internal

import (
	"fmt"
	"strings"
)

type Import struct {
	Filename string
}

func NewImport(name string) *Import {
	return &Import{
		Filename: name,
	}
}

func (i *Import) String() string {
	name := strings.TrimSpace(i.Filename)
	if name == "" {
		return ""
	}
	return fmt.Sprintf("import \"%s\";\n", name)
}

type ImportCollection []*Import

func (ic ImportCollection) Len() int {
	return len(ic)
}

func (ic ImportCollection) Less(i, j int) bool {
	return ic[i].Filename < ic[j].Filename
}

func (ic ImportCollection) Swap(i, j int) {
	ic[i], ic[j] = ic[j], ic[i]
}

func (ic ImportCollection) Search(x string) int {
	for k, v := range ic {
		if v.Filename == x {
			return k
		}
	}
	return -1
}
