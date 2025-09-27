# Adverax/configs

It is a lightweight library for managing application configurations using YAML files. It provides a simple and efficient way to organize, load, and validate configuration data for your projects.

## Features
- Supports JSON and YAML formats.
- Supports loading from files.
- Supports loading from environment variables.
- Supports loading from memory (or from command-line-arguments).
- Supports migrations from one version to another.
- Merge multiple configuration sources.
- Support for nested and hierarchical configuration structures.
- Support dynamic reloading of sources.
- Extensible and customizable.

## Installation

```bash
go get github.com/adverax/configs
```

## Usage

### Basic examples
- see adverax/configs/formats/yaml/example_test.go for more information.
- see adverax/configs/formats/json/example_test.go for more information.
- see adverax/configs/formats/mix/example_test.go for more information.
- see migrations_test.go for more information.
- see adverax/configs/dynamic/example_test.go for more information.

```go

## License

This project is licensed under the MIT License. See the LICENSE file for details.
