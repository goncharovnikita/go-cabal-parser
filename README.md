# Go cabal parser [![Go Report Card](https://goreportcard.com/badge/github.com/goncharovnikita/go-cabal-parser)](https://goreportcard.com/report/github.com/goncharovnikita/go-cabal-parser) ![tests status](https://github.com/goncharovnikita/go-cabal-parser/actions/workflows/build-go.yaml/badge.svg) ![GoDoc](https://godoc.org/github.com/goncharovnikita/go-cabal-parser)

`.cabal` parser go lib

## Usage

```
f, _ := os.Open("lib.cabal")
cabalPackage, _ := gocabalparser.NewParser().ParseReader(f)

// use cabal package
```

