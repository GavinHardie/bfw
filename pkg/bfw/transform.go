package bfw

// Transform is a data transformation that is between source and target.
type Transform interface {
	getInputSchema() Schema
	getOutputSchema() Schema

	transform(Row) Row
}
