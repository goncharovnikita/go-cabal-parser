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
				testMakeToken(tokenName, "Name"),
				testMakeToken(tokenValue, "Some name"),
				testMakeToken(tokenName, "Version"),
				testMakeToken(tokenValue, "v1.0.0.0"),
				testMakeToken(tokenName, "Cabal-Version"),
				testMakeToken(tokenValue, "v1.0.1.1"),
				testMakeToken(tokenName, "Build-Type"),
				testMakeToken(tokenValue, "Simple"),
				testMakeToken(tokenName, "License"),
				testMakeToken(tokenValue, "BSD3"),
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
				testMakeToken(tokenName, "Copyright"),
				testMakeToken(tokenValue, "© 2006      Matthias Reisner;"),
				testMakeToken(tokenValue, "© 2012–2015 Wolfgang Jeltsch"),
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
				testMakeToken(tokenName, "Source-Repository"),
				testMakeToken(tokenScopeName, "head"),
				testMakeToken(tokenScopeValueName, "Type"),
				testMakeToken(tokenScopeValueValue, "darcs"),
				testMakeToken(tokenScopeValueName, "Location"),
				testMakeToken(tokenScopeValueValue, "http://darcs.wolfgang.jeltsch.info/haskell/3d-graphics-examples/main"),
				testMakeToken(tokenName, "Source-Repository"),
				testMakeToken(tokenScopeName, "this"),
				testMakeToken(tokenScopeValueName, "Type"),
				testMakeToken(tokenScopeValueValue, "darcs"),
				testMakeToken(tokenScopeValueName, "Location"),
				testMakeToken(tokenScopeValueValue, "http://darcs.wolfgang.jeltsch.info/haskell/3d-graphics-examples/main"),
				testMakeToken(tokenScopeValueName, "Tag"),
				testMakeToken(tokenScopeValueValue, "3d-graphics-examples-0.0.0.2"),
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
				testMakeToken(tokenName, "Executable"),
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
				testMakeToken(tokenName, "Executable"),
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
