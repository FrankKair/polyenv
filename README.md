# Polyenv

Run any programming language from the command line via [tio.run](https://tio.run).

Supports 680+ languages -- anything available on Try It Online.

## Install

```sh
go install github.com/FrankKair/polyenv@latest
```

## Usage

```sh
polyenv run python3 hello.py
polyenv run ruby script.rb
polyenv run rust main.rs

polyenv search haskell
polyenv languages
polyenv version
```

## Build from source

```sh
make build
./polyenv run python3 hello.py
```

## Lincese

MIT
