# tf-module-graph

[![CI](https://github.com/DevSpecOps/tf-module-graph/actions/workflows/ci.yaml/badge.svg)](https://github.com/DevSpecOps/tf-module-graph/actions/workflows/ci.yaml)
[![GitHub release](https://img.shields.io/github/v/release/DevSpecOps/tf-module-graph)](https://github.com/DevSpecOps/tf-module-graph/releases)
[![License](https://img.shields.io/github/license/DevSpecOps/tf-module-graph)](LICENSE)
[![Docker](https://img.shields.io/badge/docker-ghcr.io-blue)](https://github.com/DevSpecOps/tf-module-graph/pkgs/container/tf-module-graph)

**Static analysis for Terraform module dependencies** – builds a dependency graph, detects cycles, dead variables, and cost regressions.

## 🚀 Quick start

### Local binary
```bash
git clone https://github.com/DevSpecOps/tf-module-graph.git
cd tf-module-graph
go build -o tf-module-graph ./cmd/tf-module-graph
./tf-module-graph --path ./test/fixtures/deps --deps
```

### Docker
```bash
docker run --rm -v $(pwd):/workspace ghcr.io/devspecops/tf-module-graph --path /workspace --deps
```

### GitHub Action
```yaml
- uses: DevSpecOps/tf-module-graph@v0.1.0
  with:
    path: './terraform'
    deps: 'true'
```

## 📋 Features

- 🔍 Parse module blocks from `.tf` files
- 🕸️ Build dependency graph
- 🔄 Detect cycles (e.g., `a -> b -> a`)
- 📊 JSON output
- 🐳 Docker image available
- 💖 Donation page (BTC/ETH)

## 🛠 Development

```bash
go test ./... -v
```

## 💖 Support

If you find this tool useful, consider [donating](docs/DONATE.md).

## 📄 License

Apache 2.0