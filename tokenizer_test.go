package gocabalparser

import (
	"fmt"
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
				testMakeToken(tokenTypeKey, "Name"),
				testMakeToken(tokenTypeValue, "3d-graphics-examples"),
				testMakeToken(tokenTypeKey, "Version"),
				testMakeToken(tokenTypeValue, "0.0.0.2"),
				testMakeToken(tokenTypeKey, "Cabal-Version"),
				testMakeToken(tokenTypeValue, ">= 1.8"),
				testMakeToken(tokenTypeKey, "Build-Type"),
				testMakeToken(tokenTypeValue, "Simple"),
				testMakeToken(tokenTypeKey, "License"),
				testMakeToken(tokenTypeValue, "BSD3"),
				testMakeToken(tokenTypeKey, "License-File"),
				testMakeToken(tokenTypeValue, "LICENSE"),
			},
		},
		{
			name:     "table with arrays",
			filename: "2.cabal",
			expected: tokens{
				testMakeToken(tokenTypeKey, "Name"),
				testMakeToken(tokenTypeValue, "3d-graphics-examples"),
				testMakeToken(tokenTypeKey, "Version"),
				testMakeToken(tokenTypeValue, "0.0.0.2"),
				testMakeToken(tokenTypeKey, "Cabal-Version"),
				testMakeToken(tokenTypeValue, ">= 1.8"),
				testMakeToken(tokenTypeKey, "Build-Type"),
				testMakeToken(tokenTypeValue, "Simple"),
				testMakeToken(tokenTypeKey, "License"),
				testMakeToken(tokenTypeValue, "BSD3"),
				testMakeToken(tokenTypeKey, "License-File"),
				testMakeToken(tokenTypeValue, "LICENSE"),
				testMakeToken(tokenTypeKey, "Copyright"),
				testMakeToken(tokenTypeValue, "© 2006      Matthias Reisner;"),
				testMakeToken(tokenTypeValue, "© 2012–2015 Wolfgang Jeltsch"),
				testMakeToken(tokenTypeKey, "Author"),
				testMakeToken(tokenTypeValue, "Matthias Reisner"),
				testMakeToken(tokenTypeKey, "Maintainer"),
				testMakeToken(tokenTypeValue, "wolfgang@cs.ioc.ee"),
				testMakeToken(tokenTypeKey, "Stability"),
				testMakeToken(tokenTypeValue, "provisional"),
			},
		},
		{
			name:     "scopes",
			filename: "3.cabal",
			expected: tokens{
				testMakeToken(tokenTypeKey, "Source-Repository"),
				testMakeToken(tokenTypeScopeName, "head"),
				testMakeToken(tokenTypeKey, "Type"),
				testMakeToken(tokenTypeValue, "darcs"),
				testMakeToken(tokenTypeKey, "Location"),
				testMakeToken(tokenTypeValue, "http://darcs.wolfgang.jeltsch.info/haskell/3d-graphics-examples/main"),
				testMakeToken(tokenTypeKey, "Source-Repository"),
				testMakeToken(tokenTypeScopeName, "this"),
				testMakeToken(tokenTypeKey, "Type"),
				testMakeToken(tokenTypeValue, "darcs"),
				testMakeToken(tokenTypeKey, "Location"),
				testMakeToken(tokenTypeValue, "http://darcs.wolfgang.jeltsch.info/haskell/3d-graphics-examples/main"),
				testMakeToken(tokenTypeKey, "Tag"),
				testMakeToken(tokenTypeValue, "3d-graphics-examples-0.0.0.2"),
				testMakeToken(tokenTypeKey, "Executable"),
				testMakeToken(tokenTypeScopeName, "mountains"),
				testMakeToken(tokenTypeKey, "Build-Depends"),
				testMakeToken(tokenTypeValue, "base   >= 3.0 && < 5"),
				testMakeToken(tokenTypeValue, "GLUT   >= 2.4 && < 2.8"),
				testMakeToken(tokenTypeValue, "OpenGL >= 2.8 && < 3.1"),
				testMakeToken(tokenTypeValue, "random >= 1.0 && < 1.2"),
				testMakeToken(tokenTypeKey, "Extensions"),
				testMakeToken(tokenTypeValue, "FlexibleContexts"),
				testMakeToken(tokenTypeKey, "Main-Is"),
				testMakeToken(tokenTypeValue, "Mountains.hs"),
				testMakeToken(tokenTypeKey, "Other-Modules"),
				testMakeToken(tokenTypeValue, "Utilities"),
				testMakeToken(tokenTypeKey, "HS-Source-Dirs"),
				testMakeToken(tokenTypeValue, "src src/mountains"),
				testMakeToken(tokenTypeKey, "Executable"),
				testMakeToken(tokenTypeScopeName, "l-systems"),
				testMakeToken(tokenTypeKey, "Build-Depends"),
				testMakeToken(tokenTypeValue, "base   >= 3.0 && < 5"),
				testMakeToken(tokenTypeValue, "GLUT   >= 2.4 && < 2.8"),
				testMakeToken(tokenTypeValue, "OpenGL >= 2.8 && < 3.1"),
				testMakeToken(tokenTypeKey, "Extensions"),
				testMakeToken(tokenTypeValue, "FlexibleContexts"),
				testMakeToken(tokenTypeKey, "Main-Is"),
				testMakeToken(tokenTypeValue, "LSystems.hs"),
				testMakeToken(tokenTypeKey, "Other-Modules"),
				testMakeToken(tokenTypeValue, "Utilities"),
				testMakeToken(tokenTypeValue, "ConiferLSystem"),
				testMakeToken(tokenTypeValue, "IslandLSystem"),
				testMakeToken(tokenTypeValue, "KochLSystem"),
				testMakeToken(tokenTypeValue, "LSystem"),
				testMakeToken(tokenTypeValue, "TreeLSystem"),
				testMakeToken(tokenTypeValue, "Turtle"),
				testMakeToken(tokenTypeKey, "HS-Source-Dirs"),
				testMakeToken(tokenTypeValue, "src src/l-systems"),
			},
		},
		{
			name:     "full",
			filename: "4.cabal",
			expected: tokens{
				testMakeToken(tokenTypeKey, "Name"),
				testMakeToken(tokenTypeValue, "3d-graphics-examples"),
				testMakeToken(tokenTypeKey, "Version"),
				testMakeToken(tokenTypeValue, "0.0.0.2"),
				testMakeToken(tokenTypeKey, "Cabal-Version"),
				testMakeToken(tokenTypeValue, ">= 1.8"),
				testMakeToken(tokenTypeKey, "Build-Type"),
				testMakeToken(tokenTypeValue, "Simple"),
				testMakeToken(tokenTypeKey, "License"),
				testMakeToken(tokenTypeValue, "BSD3"),
				testMakeToken(tokenTypeKey, "License-File"),
				testMakeToken(tokenTypeValue, "LICENSE"),
				testMakeToken(tokenTypeKey, "Copyright"),
				testMakeToken(tokenTypeValue, "© 2006      Matthias Reisner;"),
				testMakeToken(tokenTypeValue, "© 2012–2015 Wolfgang Jeltsch"),
				testMakeToken(tokenTypeKey, "Author"),
				testMakeToken(tokenTypeValue, "Matthias Reisner"),
				testMakeToken(tokenTypeKey, "Maintainer"),
				testMakeToken(tokenTypeValue, "wolfgang@cs.ioc.ee"),
				testMakeToken(tokenTypeKey, "Stability"),
				testMakeToken(tokenTypeValue, "provisional"),
				testMakeToken(tokenTypeKey, "Homepage"),
				testMakeToken(tokenTypeValue, "http://darcs.wolfgang.jeltsch.info/haskell/3d-graphics-examples"),
				testMakeToken(tokenTypeKey, "Package-URL"),
				testMakeToken(tokenTypeValue, "http://hackage.haskell.org/packages/archive/3d-graphics-examples/0.0.0.2/3d-graphics-examples-0.0.0.2.tar.gz"),
				testMakeToken(tokenTypeKey, "Synopsis"),
				testMakeToken(tokenTypeValue, "Examples of 3D graphics programming with OpenGL"),
				testMakeToken(tokenTypeKey, "Description"),
				testMakeToken(tokenTypeValue, "This package demonstrates how to program simple interactive 3D"),
				testMakeToken(tokenTypeValue, "graphics with OpenGL. It contains two programs, which are both"),
				testMakeToken(tokenTypeValue, "about fractals:"),
				testMakeToken(tokenTypeValue, "."),
				testMakeToken(tokenTypeValue, "[L-systems] generates graphics from Lindenmayer systems"),
				testMakeToken(tokenTypeValue, "(L-systems). It defines a language for L-systems as an embedded"),
				testMakeToken(tokenTypeValue, "DSL."),
				testMakeToken(tokenTypeValue, "."),
				testMakeToken(tokenTypeValue, "[Mountains] uses the generalized Brownian motion to generate"),
				testMakeToken(tokenTypeValue, "graphics that resemble mountain landscapes."),
				testMakeToken(tokenTypeValue, "."),
				testMakeToken(tokenTypeValue, "The original versions of these programs were written by Matthias"),
				testMakeToken(tokenTypeValue, "Reisner as part of a student project at the Brandenburg"),
				testMakeToken(tokenTypeValue, "University of Technology at Cottbus, Germany. Wolfgang Jeltsch,"),
				testMakeToken(tokenTypeValue, "who supervised this student project, is now maintaining these"),
				testMakeToken(tokenTypeValue, "programs."),
				testMakeToken(tokenTypeKey, "Category"),
				testMakeToken(tokenTypeValue, "Graphics, Fractals"),
				testMakeToken(tokenTypeKey, "Tested-With"),
				testMakeToken(tokenTypeValue, "GHC == 8.0.1"),
				testMakeToken(tokenTypeKey, "Source-Repository"),
				testMakeToken(tokenTypeScopeName, "head"),
				testMakeToken(tokenTypeKey, "Type"),
				testMakeToken(tokenTypeValue, "darcs"),
				testMakeToken(tokenTypeKey, "Location"),
				testMakeToken(tokenTypeValue, "http://darcs.wolfgang.jeltsch.info/haskell/3d-graphics-examples/main"),
				testMakeToken(tokenTypeKey, "Source-Repository"),
				testMakeToken(tokenTypeScopeName, "this"),
				testMakeToken(tokenTypeKey, "Type"),
				testMakeToken(tokenTypeValue, "darcs"),
				testMakeToken(tokenTypeKey, "Location"),
				testMakeToken(tokenTypeValue, "http://darcs.wolfgang.jeltsch.info/haskell/3d-graphics-examples/main"),
				testMakeToken(tokenTypeKey, "Tag"),
				testMakeToken(tokenTypeValue, "3d-graphics-examples-0.0.0.2"),
				testMakeToken(tokenTypeKey, "Executable"),
				testMakeToken(tokenTypeScopeName, "mountains"),
				testMakeToken(tokenTypeKey, "Build-Depends"),
				testMakeToken(tokenTypeValue, "base   >= 3.0 && < 5"),
				testMakeToken(tokenTypeValue, "GLUT   >= 2.4 && < 2.8"),
				testMakeToken(tokenTypeValue, "OpenGL >= 2.8 && < 3.1"),
				testMakeToken(tokenTypeValue, "random >= 1.0 && < 1.2"),
				testMakeToken(tokenTypeKey, "Extensions"),
				testMakeToken(tokenTypeValue, "FlexibleContexts"),
				testMakeToken(tokenTypeKey, "Main-Is"),
				testMakeToken(tokenTypeValue, "Mountains.hs"),
				testMakeToken(tokenTypeKey, "Other-Modules"),
				testMakeToken(tokenTypeValue, "Utilities"),
				testMakeToken(tokenTypeKey, "HS-Source-Dirs"),
				testMakeToken(tokenTypeValue, "src src/mountains"),
				testMakeToken(tokenTypeKey, "Executable"),
				testMakeToken(tokenTypeScopeName, "l-systems"),
				testMakeToken(tokenTypeKey, "Build-Depends"),
				testMakeToken(tokenTypeValue, "base   >= 3.0 && < 5"),
				testMakeToken(tokenTypeValue, "GLUT   >= 2.4 && < 2.8"),
				testMakeToken(tokenTypeValue, "OpenGL >= 2.8 && < 3.1"),
				testMakeToken(tokenTypeKey, "Extensions"),
				testMakeToken(tokenTypeValue, "FlexibleContexts"),
				testMakeToken(tokenTypeKey, "Main-Is"),
				testMakeToken(tokenTypeValue, "LSystems.hs"),
				testMakeToken(tokenTypeKey, "Other-Modules"),
				testMakeToken(tokenTypeValue, "Utilities"),
				testMakeToken(tokenTypeValue, "ConiferLSystem"),
				testMakeToken(tokenTypeValue, "IslandLSystem"),
				testMakeToken(tokenTypeValue, "KochLSystem"),
				testMakeToken(tokenTypeValue, "LSystem"),
				testMakeToken(tokenTypeValue, "TreeLSystem"),
				testMakeToken(tokenTypeValue, "Turtle"),
				testMakeToken(tokenTypeKey, "HS-Source-Dirs"),
				testMakeToken(tokenTypeValue, "src src/l-systems"),
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
