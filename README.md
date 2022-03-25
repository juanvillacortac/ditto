# Ditto ✨🔮 (previously Rosetta)

[![GoDoc](https://godoc.org/github.com/juanvillacortac/ditto?status.svg)](https://godoc.org/github.com/juanvillacortac/ditto)
[![Release CLI](https://github.com/juanvillacortac/ditto/actions/workflows/release.yml/badge.svg)](https://github.com/juanvillacortac/ditto/actions/workflows/release.yml)

An extensible schema-based CLI tool for generic code generation

![](https://c.tenor.com/IA266nP_INIAAAAC/ditto-pokemon.gif)

## Why?
Because in my actual work we have a tedious developing workflow. We don't use ORMs and we adopt a "microservices" project structure try.
I develop this tool for generate C# microservices, Angular models and services and SQLServer SPs based in a user given schema in a single shot.

## Features

* Multiformat (YAML, Protobuf and Json).
* Flexible.
* Based in templates.
* Fast.
* Multiple files generation.
* Multiple languages at same time.
* Configurable.

## Installation

You can compile from source, install from the [releases](https://github.com/juanvillacortac/ditto/releases) or simply install from NPM (the recommended way):

```bash
npm i -g ditto-gen
```

## Documentation

See an [example](https://github.com/juanvillacortac/ditto/tree/main/examples/angular) in action
