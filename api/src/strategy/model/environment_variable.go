package model

type EnvironmentVariable struct {
	Name     string
	Value    string
	DataType EnvDataType
}

type EnvDataType string

const (
	String  EnvDataType = "String"
	Number  EnvDataType = "Number"
	Boolean EnvDataType = "Boolean"
)
