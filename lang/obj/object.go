package obj

type Type string

const (
	TNumber   Type = "number"
	TString        = "string"
	TFunction      = "function"
	TList          = "list"
)

type Object interface {
	Type() Type
	AsString() string
	AsBool() bool
	AsNumber() float64
	SetDoc(string)
	GetDoc() string
}

type BaseObject struct {
	doc string
}

func (o *BaseObject) SetDoc(doc string) {
	o.doc = doc
}

func (o *BaseObject) GetDoc() string {
	return o.doc
}
