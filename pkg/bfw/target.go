package bfw

// Target is the final destination where the data will be written
type Target interface {
	getInputSchema() Schema
	write(Row)
}
