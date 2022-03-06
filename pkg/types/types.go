package types

type (
	Type          string
	IntrinsicType string
)

const (
	Integer IntrinsicType = "integer"
	Double  IntrinsicType = "double"
	String  IntrinsicType = "string"
	Date    IntrinsicType = "date"
)

var TypesMap = map[string]IntrinsicType{
	"int32":  Integer,
	"int64":  Integer,
	"double": Double,
	"string": String,
	"date":   Date,
}
