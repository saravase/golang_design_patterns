package concurrency

import "testing"

func TestGeneratorPattern(t *testing.T) {
	GeneratePattern()
	SinkPattern()
	ProcessorPattern()
	FanoutPattern()
	FaninPattern()
}
