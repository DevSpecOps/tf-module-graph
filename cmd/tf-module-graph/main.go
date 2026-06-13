package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/DevSpecOps/tf-module-graph/pkg/deps"
	"github.com/DevSpecOps/tf-module-graph/pkg/parser"
)

var (
	path   string
	jsonOut bool
	depsFlag bool
)

func main() {
	flag.StringVar(&path, "path", ".", "Terraform file or directory to scan")
	flag.BoolVar(&jsonOut, "json", false, "Output findings as JSON")
	flag.BoolVar(&depsFlag, "deps", false, "Show dependency graph and detect cycles")
	flag.Parse()

	graph := deps.NewGraph()
	modulePaths := make(map[string]string) // name -> file path

	// First pass: collect all modules
	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(p, ".tf") {
			return nil
		}
		data, err := os.ReadFile(p)
		if err != nil {
			return nil
		}
		modules, err := parser.ParseModules(p, data)
		if err != nil {
			return nil
		}
		for _, m := range modules {
			graph.AddNode(m.Name, p)
			modulePaths[m.Name] = p
		}
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Walk error: %v\n", err)
		os.Exit(1)
	}

	// Second pass: add dependencies based on source paths
	err = filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(p, ".tf") {
			return nil
		}
		data, err := os.ReadFile(p)
		if err != nil {
			return nil
		}
		modules, err := parser.ParseModules(p, data)
		if err != nil {
			return nil
		}
		for _, m := range modules {
			if m.Source == "" {
				continue
			}
			// Extract referenced module name from source (e.g., "./modules/vpc" -> "vpc")
			parts := strings.Split(m.Source, "/")
			refName := parts[len(parts)-1]
			refName = strings.TrimSuffix(refName, ".tf")
			refName = strings.TrimSuffix(refName, ".json")
			if refName != "" {
				if _, exists := modulePaths[refName]; exists {
					graph.AddDependency(m.Name, refName)
				}
			}
		}
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Walk error: %v\n", err)
		os.Exit(1)
	}

	if depsFlag {
		cycles := graph.DetectCycles()
		if jsonOut {
			out := struct {
				Edges  map[string][]string `json:"edges"`
				Cycles [][]string          `json:"cycles"`
			}{
				Edges:  make(map[string][]string),
				Cycles: cycles,
			}
			for from, toMap := range graph.Edges {
				for to := range toMap {
					out.Edges[from] = append(out.Edges[from], to)
				}
			}
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			enc.Encode(out)
		} else {
			fmt.Println("Dependency graph:")
			for from, toMap := range graph.Edges {
				for to := range toMap {
					fmt.Printf("  %s -> %s\n", from, to)
				}
			}
			if graph.HasCycles() {
				fmt.Println("\n⚠️ Cycles detected:")
				for _, cycle := range cycles {
					fmt.Printf("  %s\n", strings.Join(cycle, " -> "))
				}
			} else {
				fmt.Println("\n✅ No cycles detected")
			}
		}
	} else {
		// Normal output (list modules) – optional
		var modules []string
		for name := range graph.Nodes {
			modules = append(modules, name)
		}
		if jsonOut {
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			enc.Encode(modules)
		} else {
			fmt.Println("Modules found:")
			for _, name := range modules {
				fmt.Printf("  - %s\n", name)
			}
		}
	}
}