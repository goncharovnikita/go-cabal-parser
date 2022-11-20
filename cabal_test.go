package gocabalparser

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func testMakeToken(typ tokenType, val string) *token {
	return &token{
		Type:  typ,
		Value: val,
	}
}

func TestParse(t *testing.T) {
	cases := []struct {
		name     string
		filename string
		expected *CabalPackage
	}{
		{
			name:     "3d example",
			filename: "1.cabal",
			expected: &CabalPackage{
				Name: "3d-graphics-examples",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f, err := os.Open(fmt.Sprintf("./testdata/%s", tc.filename))
			if err != nil {
				t.Fatal(err)
			}

			defer f.Close()

			p, err := NewParser().ParseReader(f)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(tc.expected, p) {
				t.Fatalf("expected value not equal to actual")
			}
		})
	}
}

func TestTokenizer_tokenizeReader(t *testing.T) {
	cases := []struct {
		name     string
		filename string
		expected tokens
	}{
		{
			name:     "plain table",
			filename: "1.cabal",
			expected: tokens{
				testMakeToken(tokenName, "Name"),
				testMakeToken(tokenValue, "3d-graphics-examples"),
				testMakeToken(tokenName, "Version"),
				testMakeToken(tokenValue, "0.0.0.2"),
				testMakeToken(tokenName, "Cabal-Version"),
				testMakeToken(tokenValue, ">= 1.8"),
				testMakeToken(tokenName, "Build-Type"),
				testMakeToken(tokenValue, "Simple"),
				testMakeToken(tokenName, "License"),
				testMakeToken(tokenValue, "BSD3"),
				testMakeToken(tokenName, "License-File"),
				testMakeToken(tokenValue, "LICENSE"),
			},
		},
		{
			name:     "table with arrays",
			filename: "2.cabal",
			expected: tokens{
				testMakeToken(tokenName, "Name"),
				testMakeToken(tokenValue, "3d-graphics-examples"),
				testMakeToken(tokenName, "Version"),
				testMakeToken(tokenValue, "0.0.0.2"),
				testMakeToken(tokenName, "Cabal-Version"),
				testMakeToken(tokenValue, ">= 1.8"),
				testMakeToken(tokenName, "Build-Type"),
				testMakeToken(tokenValue, "Simple"),
				testMakeToken(tokenName, "License"),
				testMakeToken(tokenValue, "BSD3"),
				testMakeToken(tokenName, "License-File"),
				testMakeToken(tokenValue, "LICENSE"),
				testMakeToken(tokenName, "Copyright"),
				testMakeToken(tokenValue, "© 2006      Matthias Reisner;"),
				testMakeToken(tokenValue, "© 2012–2015 Wolfgang Jeltsch"),
				testMakeToken(tokenName, "Author"),
				testMakeToken(tokenValue, "Matthias Reisner"),
				testMakeToken(tokenName, "Maintainer"),
				testMakeToken(tokenValue, "wolfgang@cs.ioc.ee"),
				testMakeToken(tokenName, "Stability"),
				testMakeToken(tokenValue, "provisional"),
			},
		},
		{
			name:     "scopes",
			filename: "3.cabal",
			expected: tokens{
				testMakeToken(tokenScopeType, "Source-Repository"),
				testMakeToken(tokenScopeName, "head"),
				testMakeToken(tokenScopeValueName, "Type"),
				testMakeToken(tokenScopeValueValue, "darcs"),
				testMakeToken(tokenScopeValueName, "Location"),
				testMakeToken(tokenScopeValueValue, "http://darcs.wolfgang.jeltsch.info/haskell/3d-graphics-examples/main"),
				testMakeToken(tokenScopeType, "Source-Repository"),
				testMakeToken(tokenScopeName, "this"),
				testMakeToken(tokenScopeValueName, "Type"),
				testMakeToken(tokenScopeValueValue, "darcs"),
				testMakeToken(tokenScopeValueName, "Location"),
				testMakeToken(tokenScopeValueValue, "http://darcs.wolfgang.jeltsch.info/haskell/3d-graphics-examples/main"),
				testMakeToken(tokenScopeValueName, "Tag"),
				testMakeToken(tokenScopeValueValue, "3d-graphics-examples-0.0.0.2"),
				testMakeToken(tokenScopeType, "Executable"),
				testMakeToken(tokenScopeName, "mountains"),
				testMakeToken(tokenScopeValueName, "Build-Depends"),
				testMakeToken(tokenScopeValueValue, "base   >= 3.0 && < 5"),
				testMakeToken(tokenScopeValueValue, "GLUT   >= 2.4 && < 2.8"),
				testMakeToken(tokenScopeValueValue, "OpenGL >= 2.8 && < 3.1"),
				testMakeToken(tokenScopeValueValue, "random >= 1.0 && < 1.2"),
				testMakeToken(tokenScopeValueName, "Extensions"),
				testMakeToken(tokenScopeValueValue, "FlexibleContexts"),
				testMakeToken(tokenScopeValueName, "Main-Is"),
				testMakeToken(tokenScopeValueValue, "Mountains.hs"),
				testMakeToken(tokenScopeValueName, "Other-Modules"),
				testMakeToken(tokenScopeValueValue, "Utilities"),
				testMakeToken(tokenScopeValueName, "HS-Source-Dirs"),
				testMakeToken(tokenScopeValueValue, "src src/mountains"),
				testMakeToken(tokenScopeType, "Executable"),
				testMakeToken(tokenScopeName, "l-systems"),
				testMakeToken(tokenScopeValueName, "Build-Depends"),
				testMakeToken(tokenScopeValueValue, "base   >= 3.0 && < 5"),
				testMakeToken(tokenScopeValueValue, "GLUT   >= 2.4 && < 2.8"),
				testMakeToken(tokenScopeValueValue, "OpenGL >= 2.8 && < 3.1"),
				testMakeToken(tokenScopeValueName, "Extensions"),
				testMakeToken(tokenScopeValueValue, "FlexibleContexts"),
				testMakeToken(tokenScopeValueName, "Main-Is"),
				testMakeToken(tokenScopeValueValue, "LSystems.hs"),
				testMakeToken(tokenScopeValueName, "Other-Modules"),
				testMakeToken(tokenScopeValueValue, "Utilities"),
				testMakeToken(tokenScopeValueValue, "ConiferLSystem"),
				testMakeToken(tokenScopeValueValue, "IslandLSystem"),
				testMakeToken(tokenScopeValueValue, "KochLSystem"),
				testMakeToken(tokenScopeValueValue, "LSystem"),
				testMakeToken(tokenScopeValueValue, "TreeLSystem"),
				testMakeToken(tokenScopeValueValue, "Turtle"),
				testMakeToken(tokenScopeValueName, "HS-Source-Dirs"),
				testMakeToken(tokenScopeValueValue, "src src/l-systems"),
			},
		},
		{
			name:     "full",
			filename: "4.cabal",
			expected: tokens{
				testMakeToken(tokenName, "Name"),
				testMakeToken(tokenValue, "3d-graphics-examples"),
				testMakeToken(tokenName, "Version"),
				testMakeToken(tokenValue, "0.0.0.2"),
				testMakeToken(tokenName, "Cabal-Version"),
				testMakeToken(tokenValue, ">= 1.8"),
				testMakeToken(tokenName, "Build-Type"),
				testMakeToken(tokenValue, "Simple"),
				testMakeToken(tokenName, "License"),
				testMakeToken(tokenValue, "BSD3"),
				testMakeToken(tokenName, "License-File"),
				testMakeToken(tokenValue, "LICENSE"),
				testMakeToken(tokenName, "Copyright"),
				testMakeToken(tokenValue, "© 2006      Matthias Reisner;"),
				testMakeToken(tokenValue, "© 2012–2015 Wolfgang Jeltsch"),
				testMakeToken(tokenName, "Author"),
				testMakeToken(tokenValue, "Matthias Reisner"),
				testMakeToken(tokenName, "Maintainer"),
				testMakeToken(tokenValue, "wolfgang@cs.ioc.ee"),
				testMakeToken(tokenName, "Stability"),
				testMakeToken(tokenValue, "provisional"),
				testMakeToken(tokenName, "Homepage"),
				testMakeToken(tokenValue, "http://darcs.wolfgang.jeltsch.info/haskell/3d-graphics-examples"),
				testMakeToken(tokenName, "Package-URL"),
				testMakeToken(tokenValue, "http://hackage.haskell.org/packages/archive/3d-graphics-examples/0.0.0.2/3d-graphics-examples-0.0.0.2.tar.gz"),
				testMakeToken(tokenName, "Synopsis"),
				testMakeToken(tokenValue, "Examples of 3D graphics programming with OpenGL"),
				testMakeToken(tokenName, "Description"),
				testMakeToken(tokenValue, "This package demonstrates how to program simple interactive 3D"),
				testMakeToken(tokenValue, "graphics with OpenGL. It contains two programs, which are both"),
				testMakeToken(tokenValue, "about fractals:"),
				testMakeToken(tokenValue, "."),
				testMakeToken(tokenValue, "[L-systems] generates graphics from Lindenmayer systems"),
				testMakeToken(tokenValue, "(L-systems). It defines a language for L-systems as an embedded"),
				testMakeToken(tokenValue, "DSL."),
				testMakeToken(tokenValue, "."),
				testMakeToken(tokenValue, "[Mountains] uses the generalized Brownian motion to generate"),
				testMakeToken(tokenValue, "graphics that resemble mountain landscapes."),
				testMakeToken(tokenValue, "."),
				testMakeToken(tokenValue, "The original versions of these programs were written by Matthias"),
				testMakeToken(tokenValue, "Reisner as part of a student project at the Brandenburg"),
				testMakeToken(tokenValue, "University of Technology at Cottbus, Germany. Wolfgang Jeltsch,"),
				testMakeToken(tokenValue, "who supervised this student project, is now maintaining these"),
				testMakeToken(tokenValue, "programs."),
				testMakeToken(tokenName, "Category"),
				testMakeToken(tokenValue, "Graphics, Fractals"),
				testMakeToken(tokenName, "Tested-With"),
				testMakeToken(tokenValue, "GHC == 8.0.1"),
				testMakeToken(tokenScopeType, "Source-Repository"),
				testMakeToken(tokenScopeName, "head"),
				testMakeToken(tokenScopeValueName, "Type"),
				testMakeToken(tokenScopeValueValue, "darcs"),
				testMakeToken(tokenScopeValueName, "Location"),
				testMakeToken(tokenScopeValueValue, "http://darcs.wolfgang.jeltsch.info/haskell/3d-graphics-examples/main"),
				testMakeToken(tokenScopeType, "Source-Repository"),
				testMakeToken(tokenScopeName, "this"),
				testMakeToken(tokenScopeValueName, "Type"),
				testMakeToken(tokenScopeValueValue, "darcs"),
				testMakeToken(tokenScopeValueName, "Location"),
				testMakeToken(tokenScopeValueValue, "http://darcs.wolfgang.jeltsch.info/haskell/3d-graphics-examples/main"),
				testMakeToken(tokenScopeValueName, "Tag"),
				testMakeToken(tokenScopeValueValue, "3d-graphics-examples-0.0.0.2"),
				testMakeToken(tokenScopeType, "Executable"),
				testMakeToken(tokenScopeName, "mountains"),
				testMakeToken(tokenScopeValueName, "Build-Depends"),
				testMakeToken(tokenScopeValueValue, "base   >= 3.0 && < 5"),
				testMakeToken(tokenScopeValueValue, "GLUT   >= 2.4 && < 2.8"),
				testMakeToken(tokenScopeValueValue, "OpenGL >= 2.8 && < 3.1"),
				testMakeToken(tokenScopeValueValue, "random >= 1.0 && < 1.2"),
				testMakeToken(tokenScopeValueName, "Extensions"),
				testMakeToken(tokenScopeValueValue, "FlexibleContexts"),
				testMakeToken(tokenScopeValueName, "Main-Is"),
				testMakeToken(tokenScopeValueValue, "Mountains.hs"),
				testMakeToken(tokenScopeValueName, "Other-Modules"),
				testMakeToken(tokenScopeValueValue, "Utilities"),
				testMakeToken(tokenScopeValueName, "HS-Source-Dirs"),
				testMakeToken(tokenScopeValueValue, "src src/mountains"),
				testMakeToken(tokenScopeType, "Executable"),
				testMakeToken(tokenScopeName, "l-systems"),
				testMakeToken(tokenScopeValueName, "Build-Depends"),
				testMakeToken(tokenScopeValueValue, "base   >= 3.0 && < 5"),
				testMakeToken(tokenScopeValueValue, "GLUT   >= 2.4 && < 2.8"),
				testMakeToken(tokenScopeValueValue, "OpenGL >= 2.8 && < 3.1"),
				testMakeToken(tokenScopeValueName, "Extensions"),
				testMakeToken(tokenScopeValueValue, "FlexibleContexts"),
				testMakeToken(tokenScopeValueName, "Main-Is"),
				testMakeToken(tokenScopeValueValue, "LSystems.hs"),
				testMakeToken(tokenScopeValueName, "Other-Modules"),
				testMakeToken(tokenScopeValueValue, "Utilities"),
				testMakeToken(tokenScopeValueValue, "ConiferLSystem"),
				testMakeToken(tokenScopeValueValue, "IslandLSystem"),
				testMakeToken(tokenScopeValueValue, "KochLSystem"),
				testMakeToken(tokenScopeValueValue, "LSystem"),
				testMakeToken(tokenScopeValueValue, "TreeLSystem"),
				testMakeToken(tokenScopeValueValue, "Turtle"),
				testMakeToken(tokenScopeValueName, "HS-Source-Dirs"),
				testMakeToken(tokenScopeValueValue, "src src/l-systems"),
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f, err := os.Open(fmt.Sprintf("./testdata/%s", tc.filename))
			if err != nil {
				t.Fatal(err)
			}

			defer f.Close()

			tt := newTokenizer()
			p, err := tt.TokenizeReader(f)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(tc.expected, p) {
				t.Logf("expected: %v", tc.expected)
				t.Logf("actual: %v", p)
				t.Logf("states: %v", tt.states)
				t.Fatalf("expected value not equal to actual")
			}
		})
	}
}

func getTestFiles() ([]*os.File, error) {
	files, err := ioutil.ReadDir("./testdata")
	if err != nil {
		return nil, err
	}

	tf := make([]*os.File, 0)

	for _, fd := range files {
		if !fd.IsDir() {
			f, err := os.Open(fmt.Sprintf("./testdata/%s", fd.Name()))
			if err != nil {
				return nil, err
			}

			tf = append(tf, f)
		}
	}

	return tf, nil
}
