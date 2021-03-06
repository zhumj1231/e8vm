[![BuildStatus](https://travis-ci.org/e8vm/e8vm.png?branch=master)](https://travis-ci.org/e8vm/e8vm)

```
go get -u e8vm.io/e8vm/...
```

# E8VM

[![Join the chat at https://gitter.im/e8vm/e8vm](https://badges.gitter.im/e8vm/e8vm.svg)](https://gitter.im/e8vm/e8vm?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Emul8ed Virtual Machine (E8VM) is a self-contained system that has its
own instruction set -- `arch8`, its own assembly language and
assembler -- `asm8`, its own system language -- `g8`, and its own
project building system -- `build8`. Using `g8` and `build8`, we can
build a small operating system [`os8`](https://github.com/e8vm/os8).

The project is written entirely in Go language. Plus, each file in the
project has no more than 300 lines, with each line no more than 80
characters. Among these small files, there are no circular
dependencies, and as a result, the project architecture can be
automatically visualized from static code analysis.

[Check the visualization of the architecture.](https://e8vm.io/e8vm)

The main project in this repository depends on nothing other than the
Go standard library. Hence, it is not a compiler project that based on
LLVM.

For Go language documentation on the package APIs, I recommend
[GoWalker](https://gowalker.org/e8vm.io/e8vm). I find it slightly
better than [godoc.org](https://godoc.org/e8vm.io/e8vm).

## To Use `make`

The project comes with a `makefile`, which formats the code files,
check lints, check circular dependencies and build tags. Running the
`makefile` requires installing some tools.

```
go get -u e8vm.io/tools/...
go get -u github.com/golang/lint/golint
go get -u github.com/jstemmer/gotags
```

## Copyright and License

The project developers own the copyright; my employer (Google) does
not own the copyright. Apache is the License.
