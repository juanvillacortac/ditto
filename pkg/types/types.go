package types

type (
	Type          interface{}
	PrimitiveType string
)

const (
	Integer PrimitiveType = "integer"
	Double  PrimitiveType = "double"
	String  PrimitiveType = "string"
)

var TypesMap = map[string]PrimitiveType{
	"int32":  Integer,
	"int64":  Integer,
	"double": Double,
	"string": String,
}
