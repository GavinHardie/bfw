package bfw

// Source is where the data originates from.
type Source interface {
	getOutputSchema() Schema
	read() Row
}

// TableSource reads all the data from a SQL table.
type TableSource struct {
}

// SelectSource reads all the data from a select statement.
type SelectSource struct {
}

func (ts *TableSource) getOutputSchema() {

}

func (ts *TableSource) read() Row {
	return nil
}

func (ss *SelectSource) getOutputSchema() {

}

func (ss *SelectSource) read() Row {
	return nil
}
