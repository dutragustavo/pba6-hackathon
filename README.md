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