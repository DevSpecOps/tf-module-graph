package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/DevSpecOps/tf-module-graph/pkg/parser"
)

type Finding struct {
	File    string `json:"file"`
	Module  string `json:"module,omitempty"`
	RuleID  string `json:"rule_id"`
	Message string `json:"message"`
}

var (
	path   string
	jsonOut bool
)

func main() {
	flag.StringVar(&path, "path", ".", "Terraform file or directory to scan")
	flag.BoolVar(&jsonOut, "json", false, "output findings as JSON")
	flag.Parse()

	var findings []Finding

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
			// skip files with parse errors for now
			return nil
		}
		for _, mod := range modules {
			findings = append(findings, Finding{
				File:    p,
				Module:  mod.Name,
				RuleID:  "MODULE_FOUND",
				Message: fmt.Sprintf("Module '%s' found", mod.Name),
			})
		}
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Walk error: %v\n", err)
		os.Exit(1)
	}

	if jsonOut {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		enc.Encode(findings)
	} else {
		for _, f := range findings {
			fmt.Printf("📦 %s: %s\n", f.File, f.Message)
		}
		if len(findings) == 0 {
			fmt.Println("✅ No module blocks found")
		}
	}

	if len(findings) > 0 {
		// В будущем здесь будет exit code 1, пока 0 для демонстрации
	}
}