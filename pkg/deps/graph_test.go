package deps

import (
	"testing"
)

func TestGraphNoCycle(t *testing.T) {
	g := NewGraph()
	g.AddNode("A", "path/A")
	g.AddNode("B", "path/B")
	g.AddNode("C", "path/C")
	g.AddDependency("A", "B")
	g.AddDependency("B", "C")
	if g.Cyclic() {
		t.Error("Expected no cycle, but got cycle")
	}
}

func TestGraphCycle(t *testing.T) {
	g := NewGraph()
	g.AddNode("A", "path/A")
	g.AddNode("B", "path/B")
	g.AddDependency("A", "B")
	g.AddDependency("B", "A")
	if !g.Cyclic() {
		t.Error("Expected cycle, but got none")
	}
	cycles := g.DetectCycles()
	if len(cycles) == 0 {
		t.Fatal("Expected at least one cycle")
	}
}

func TestGraphMultipleCycles(t *testing.T) {
	g := NewGraph()
	g.AddNode("A", "")
	g.AddNode("B", "")
	g.AddNode("C", "")
	g.AddDependency("A", "B")
	g.AddDependency("B", "C")
	g.AddDependency("C", "A") // cycle A-B-C
	g.AddNode("D", "")
	g.AddDependency("D", "D") // self-cycle
	cycles := g.DetectCycles()
	if len(cycles) != 2 {
		t.Errorf("Expected 2 cycles, got %d", len(cycles))
	}
}