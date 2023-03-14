package funcs

type ModelFunc interface {
	String() string
	FullName() string
	Req() string
	Resp() string
	Name() string
	ModelName() string
}
