# GoSmart-Lint: Early Detection of PolkaVM Incompatibilities in Go Code

## Hackathon Submission

A Semgrep-based static analysis tool for detecting issues in Go code that would be incompatible with PolkaVM bytecode. This tool is part of the initiative to make the developer experience more streamlined when working with JAM.

## Problem Statement

When compiling Go code to PolkaVM bytecode, developers often encounter issues (e.g. JAM doesn’t support multi-threading) only at the final stage of the compilation pipeline. These late-stage errors require time-consuming rewrites and can significantly delay development. We need a way to catch these issues earlier in the development process.

## Our Solution

GoSmart-Lint provides static analysis for Go code to catch PolkaVM compatibility issues during development rather than during final compilation. By identifying problematic patterns early, developers can avoid rework and accelerate their development cycles.

## Our Journey & Research Process

During this hackathon, we analyzed the PVM toolkit and considered various potential projects such as:

- **Full Stack Analysis**: We started by analyzing the entire pipeline for generating PVM bytecode to identify key pain points
- **Fuzz Testing**: We explored potential fuzz testing methodologies to detect issues in generated PVM code
- **LLM Integration**: We investigated LLM tooling for improving PVM code generation and reverse engineering of compiled projects, exploring if we could feed LLMs with high-level source code and compiled bytecode to generate more readable PVM
- **Early Detection**: We ultimately settled on using Semgrep at the high-level language (Go in this case) to detect practices which are incompatible with PVM

After evaluating these options, we determined that static analysis with custom rules would provide the most immediate value to developers by catching issues at the earliest possible stage in the development process.

## Key Features

The analyzer focuses on detecting patterns that are incompatible with PolkaVM's execution model, with a strong emphasis on multi-threading issues:

- **Multi-threading and concurrency**: Detection of goroutines, channels, mutexes, and other concurrency primitives that are incompatible with PolkaVM's execution model
- File and network operations
- Floating point operations

## Quick Start

```bash
# Install Semgrep
pip install semgrep

# Clone this repository
git clone https://github.com/dutragustavo/gosmart-lint.git
cd gosmart-lint

# Run against example code
semgrep --config go-threading.yaml main.go

```

## Example Detection

```go
// This code would be flagged as incompatible with PVM:
package main

import "fmt"

func main() {
    ch := make(chan string)
    go func() { ch <- "Hello from another thread!" }()
    fmt.Println(<-ch)
}
```

## Value Proposition

1. **Development Speed**: Catch incompatibility issues early, avoiding costly late-stage rewrites
2. **Learning Tool**: Help developers understand the constraints of PolkaVM execution
3. **CI Integration**: Automatically enforce PolkaVM compatibility in continuous integration
4. **Documentation**: Rules serve as living documentation of PolkaVM limitations

## Future Work

- Expand rule coverage to other PolkaVM incompatibility patterns
- Create IDE extensions for real-time feedback
- Add automatic code transformation suggestions
- Support additional languages targeting PolkaVM

## Try It Yourself

Set up the tool and run it against your own Go code with:

```yaml
rules:
  - id: no-go-concurrency
    pattern-either:
      - pattern: go $F(...)
      - pattern: |
          go func() {
              $BODY
          }()
      - pattern: make(chan $TYPE, ...)
      - pattern: var $X sync.Mutex
      - pattern: var $X sync.WaitGroup
      - pattern: var $X sync.RWMutex
      - pattern: $CHAN <- $VAL
      - pattern: <-$CHAN
      - pattern: atomic.$F(...)
    message: "PVM does not support multi-threading!!!"
    languages: [go]
    severity: ERROR
```

## License

MIT