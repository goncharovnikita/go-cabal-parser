package gocabalparser

import (
	"reflect"
	"testing"
)

func TestTokensParser(t *testing.T) {
	cases := []struct {
		name     string
		tokens   tokens
		expected *CabalPackage
	}{
		{
			name: "basic fields",
			tokens: tokens{
				testMakeToken(tokenTypeKey, "Name"),
				testMakeToken(tokenTypeValue, "Some name"),
				testMakeToken(tokenTypeKey, "Version"),
				testMakeToken(tokenTypeValue, "v1.0.0.0"),
				testMakeToken(tokenTypeKey, "Cabal-Version"),
				testMakeToken(tokenTypeValue, "v1.0.1.1"),
				testMakeToken(tokenTypeKey, "Build-Type"),
				testMakeToken(tokenTypeValue, "Simple"),
				testMakeToken(tokenTypeKey, "License"),
				testMakeToken(tokenTypeValue, "BSD3"),
			},
			expected: &CabalPackage{
				Name:         "Some name",
				Version:      "v1.0.0.0",
				CabalVersion: "v1.0.1.1",
				BuildType:    "Simple",
				License:      "BSD3",
			},
		},
		{
			name: "array fields",
			tokens: tokens{
				testMakeToken(tokenTypeKey, "Copyright"),
				testMakeToken(tokenTypeValue, "© 2006      Matthias Reisner;"),
				testMakeToken(tokenTypeValue, "© 2012–2015 Wolfgang Jeltsch"),
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
			},
			expected: &CabalPackage{
				Copyright: []string{
					"© 2006      Matthias Reisner;",
					"© 2012–2015 Wolfgang Jeltsch",
				},
				Description: []string{
					"This package demonstrates how to program simple interactive 3D",
					"graphics with OpenGL. It contains two programs, which are both",
					"about fractals:",
					".",
					"[L-systems] generates graphics from Lindenmayer systems",
					"(L-systems). It defines a language for L-systems as an embedded",
					"DSL.",
					".",
					"[Mountains] uses the generalized Brownian motion to generate",
					"graphics that resemble mountain landscapes.",
					".",
					"The original versions of these programs were written by Matthias",
					"Reisner as part of a student project at the Brandenburg",
					"University of Technology at Cottbus, Germany. Wolfgang Jeltsch,",
					"who supervised this student project, is now maintaining these",
					"programs.",
				},
			},
		},
		{
			name: "repo fields",
			tokens: tokens{
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
			},
			expected: &CabalPackage{
				Repositories: map[string]*SourceRepository{
					"head": {
						Type:     "darcs",
						Location: "http://darcs.wolfgang.jeltsch.info/haskell/3d-graphics-examples/main",
					},
					"this": {
						Type:     "darcs",
						Location: "http://darcs.wolfgang.jeltsch.info/haskell/3d-graphics-examples/main",
						Tag:      "3d-graphics-examples-0.0.0.2",
					},
				},
			},
		},
		{
			name: "executable fields",
			tokens: tokens{
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
			expected: &CabalPackage{
				Executables: map[string]*Executable{
					"mountains": {
						BuildDepends: []string{
							"base   >= 3.0 && < 5",
							"GLUT   >= 2.4 && < 2.8",
							"OpenGL >= 2.8 && < 3.1",
							"random >= 1.0 && < 1.2",
						},
						Extensions: []string{
							"FlexibleContexts",
						},
						MainIs: "Mountains.hs",
						OtherModules: []string{
							"Utilities",
						},
						HSSourceDirs: []string{
							"src src/mountains",
						},
					},
					"l-systems": {
						BuildDepends: []string{
							"base   >= 3.0 && < 5",
							"GLUT   >= 2.4 && < 2.8",
							"OpenGL >= 2.8 && < 3.1",
						},
						Extensions: []string{
							"FlexibleContexts",
						},
						MainIs: "LSystems.hs",
						OtherModules: []string{
							"Utilities",
							"ConiferLSystem",
							"IslandLSystem",
							"KochLSystem",
							"LSystem",
							"TreeLSystem",
							"Turtle",
						},
						HSSourceDirs: []string{
							"src src/l-systems",
						},
					},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := newTokensParser().Parse(tc.tokens)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Logf("expected: %+v", tc.expected)
				t.Logf("actual: %+v", actual)

				t.FailNow()
			}
		})
	}
}
