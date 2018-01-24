package bfw

// Pipe is an executable data processing pipeline.
// A pipe consists of a Source and a Target and optional Transforms.
type Pipe interface {
	getSource() Source
	getTarget() Target
	getTransforms() []Transform
}
