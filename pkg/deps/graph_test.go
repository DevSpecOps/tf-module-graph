package deps

import (
	"testing"
)

func TestNoCycle(t *testing.T) {
	g := NewGraph()
	g.AddNode("A", "")
	g.AddNode("B", "")
	g.AddDependency("A", "B")
	if g.HasCycles() {
		t.Error("Expected no cycle")
	}
}

func TestCycle(t *testing.T) {
	g := NewGraph()
	g.AddNode("A", "")
	g.AddNode("B", "")
	g.AddDependency("A", "B")
	g.AddDependency("B", "A")
	if !g.HasCycles() {
		t.Error("Expected cycle")
	}
}