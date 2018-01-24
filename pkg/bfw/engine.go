package bfw

func run(pipe Pipe) {

	source := pipe.getSource()
	target := pipe.getTarget()

	// Source -> Transform -> Transform -> Target
	// Check that each inputSchema can read from the prevous outputSchema
	incomingSchema := source.getOutputSchema()
	for _, transform := range pipe.getTransforms() {
		if !transform.getInputSchema().canReadFrom(incomingSchema) {
			// Throw error?
		}
		incomingSchema = transform.getOutputSchema()
	}
	if !target.getInputSchema().canReadFrom(incomingSchema) {

	}

	for row := source.read(); row != nil; row = source.read() {
		for _, transform := range pipe.getTransforms() {
			row = transform.transform(row)
		}
		target.write(row)
	}

}
