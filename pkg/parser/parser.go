package parser

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

type Module struct {
	Name   string
	Source string
	Path   string
}

func ParseModules(filename string, src []byte) ([]Module, error) {
	file, diags := hclsyntax.ParseConfig(src, filename, hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		return nil, diags
	}
	body := file.Body
	// Получаем все блоки модулей
	content, _, diags := body.PartialContent(&hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{
				Type:       "module",
				LabelNames: []string{"name"},
			},
		},
	})
	if diags.HasErrors() {
		return nil, diags
	}
	var modules []Module
	for _, block := range content.Blocks {
		if block.Type == "module" {
			name := block.Labels[0]
			// Для получения атрибута source нужно обратиться к телу блока
			// Используем JustAttributes, чтобы извлечь все атрибуты блока
			attrs, diags := block.Body.JustAttributes()
			if diags.HasErrors() {
				continue
			}
			var source string
			if attr, exists := attrs["source"]; exists {
				val, diags := attr.Expr.Value(nil)
				if !diags.HasErrors() && val.Type() == cty.String {
					source = val.AsString()
				}
			}
			modules = append(modules, Module{
				Name:   name,
				Source: source,
				Path:   filename,
			})
		}
	}
	return modules, nil
}