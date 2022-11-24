# Go cabal parser

`.cabal` parser go lib

## Usage

```
f, _ := os.Open("lib.cabal")
cabalPackage, _ := gocabalparser.NewParser().ParseReader(f)

// use cabal package
```

![tests status](https://github.com/goncharovnikita/go-cabal-parser/actions/workflows/build-go.yaml/badge.svg)

