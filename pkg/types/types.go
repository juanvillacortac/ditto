package types

type (
	Type          string
	IntrinsicType string
)

const (
	Int32   IntrinsicType = "int32"
	Int64   IntrinsicType = "int64"
	Double  IntrinsicType = "double"
	String  IntrinsicType = "string"
	Boolean IntrinsicType = "boolean"
	Date    IntrinsicType = "date"
)

var TypesMap = map[string]IntrinsicType{
	"int32":  Int32,
	"int64":  Int64,
	"double": Double,
	"string": String,
	"bool":   Boolean,
	"date":   Date,
}
