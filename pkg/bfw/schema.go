package bfw

type Schema interface {
	canReadFrom(Schema) bool
	canWriteTo(Schema) bool
}
