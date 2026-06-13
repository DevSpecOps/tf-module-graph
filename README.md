
# tf-module-graph

[![CI](https://github.com/DevSpecOps/tf-module-graph/actions/workflows/ci.yaml/badge.svg)](https://github.com/DevSpecOps/tf-module-graph/actions/workflows/ci.yaml)
[![License](https://img.shields.io/github/license/DevSpecOps/tf-module-graph)](LICENSE)

**Static analysis for Terraform module dependencies** – builds a dependency graph, detects cycles, dead variables, and cost regressions.

🚧 **MVP in progress** – currently parses module blocks.

## Quick start

```bash
git clone https://github.com/DevSpecOps/tf-module-graph.git
cd tf-module-graph
go build -o tf-module-graph ./cmd/tf-module-graph
./tf-module-graph --path ./test/fixtures
```

## Next features

- Dependency graph construction
- Cycle detection
- Dead variable analysis
- Cost regression warnings (Infracost integration)

## License

Apache 2.0

