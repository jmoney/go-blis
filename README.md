# Go-Lapack

Linear algebra package in golang using native bindings to various architecture specific linear algebra libraries.

The difference between this repository [gonum](https://github.com/gonum/gonum) is gonum is a golang native implementation while this repo uses native bindings to the blas or blis.

## go-blis

Go-Blis is a Go wrapper for the [BLIS](https://github.com/flame/blis) library. BLIS is a portable software framework for instantiating high-performance BLAS-like dense linear algebra libraries. The framework was designed to isolate essential kernels of computation that, when optimized, immediately enable optimized implementations of most of its commonly used and computationally intensive operations.

## Local Testing

go-blis uses cgo to interface with the BLIS library. The BLIS library must be installed on the system and the library must be in the system's library path. Local development has been done on a Mac M1 processor. To install on a mac homebrew can be used.

```bash
brew bundle --file=./Brewfile
```

This installs blis and any other dependencies.

Then to run the tests:

```bash
make test
```
