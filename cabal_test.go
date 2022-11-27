package gocabalparser

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	cases := []struct {
		name     string
		filename string
		expected *CabalPackage
	}{
		{
			name:     "text values",
			filename: "1.cabal",
			expected: &CabalPackage{
				Name:         "3d-graphics-examples",
				Version:      "0.0.0.2",
				CabalVersion: ">= 1.8",
				BuildType:    "Simple",
				License:      "BSD3",
				LicenseFile:  "LICENSE",
			},
		},
		{
			name:     "text values with array",
			filename: "2.cabal",
			expected: &CabalPackage{
				Name:         "3d-graphics-examples",
				Version:      "0.0.0.2",
				CabalVersion: ">= 1.8",
				BuildType:    "Simple",
				License:      "BSD3",
				LicenseFile:  "LICENSE",
				Copyright: []string{"© 2006      Matthias Reisner;",
					"© 2012–2015 Wolfgang Jeltsch",
				},
				Author:     "Matthias Reisner",
				Maintainer: "wolfgang@cs.ioc.ee",
				Stability:  "provisional",
			},
		},
		{
			name:     "scopes",
			filename: "3.cabal",
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
				Executables: map[string]*Executable{
					"mountains": {
						BuildDepends: []*Dependency{
							{
								Name: "base",
								Gte:  "3.0",
								Lt:   "5",
							},
							{
								Name: "GLUT",
								Gte:  "2.4",
								Lt:   "2.8",
							},
							{
								Name: "OpenGL",
								Gte:  "2.8",
								Lt:   "3.1",
							},
							{
								Name: "random",
								Gte:  "1.0",
								Lt:   "1.2",
							},
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
						BuildDepends: []*Dependency{
							{
								Name: "base",
								Gte:  "3.0",
								Lt:   "5",
							},
							{
								Name: "GLUT",
								Gte:  "2.4",
								Lt:   "2.8",
							},
							{
								Name: "OpenGL",
								Gte:  "2.8",
								Lt:   "3.1",
							},
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
		{
			name:     "full",
			filename: "4.cabal",
			expected: &CabalPackage{
				Name:         "3d-graphics-examples",
				Version:      "0.0.0.2",
				CabalVersion: ">= 1.8",
				BuildType:    "Simple",
				License:      "BSD3",
				LicenseFile:  "LICENSE",
				Copyright: []string{"© 2006      Matthias Reisner;",
					"© 2012–2015 Wolfgang Jeltsch",
				},
				Author:     "Matthias Reisner",
				Maintainer: "wolfgang@cs.ioc.ee",
				Stability:  "provisional",
				Homepage:   "http://darcs.wolfgang.jeltsch.info/haskell/3d-graphics-examples",
				PackageURL: "http://hackage.haskell.org/packages/archive/3d-graphics-examples/0.0.0.2/3d-graphics-examples-0.0.0.2.tar.gz",
				Synopsis: []string{
					"Examples of 3D graphics programming with OpenGL",
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
				Category:   "Graphics, Fractals",
				TestedWith: "GHC == 8.0.1",
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
				Executables: map[string]*Executable{
					"mountains": {
						BuildDepends: []*Dependency{
							{
								Name: "base",
								Gte:  "3.0",
								Lt:   "5",
							},
							{
								Name: "GLUT",
								Gte:  "2.4",
								Lt:   "2.8",
							},
							{
								Name: "OpenGL",
								Gte:  "2.8",
								Lt:   "3.1",
							},
							{
								Name: "random",
								Gte:  "1.0",
								Lt:   "1.2",
							},
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
						BuildDepends: []*Dependency{
							{
								Name: "base",
								Gte:  "3.0",
								Lt:   "5",
							},
							{
								Name: "GLUT",
								Gte:  "2.4",
								Lt:   "2.8",
							},
							{
								Name: "OpenGL",
								Gte:  "2.8",
								Lt:   "3.1",
							},
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
		{
			name:     "with library",
			filename: "5.cabal",
			expected: &CabalPackage{
				Name:         "base",
				Version:      "4.17.0.0",
				CabalVersion: "3.0",
				BuildType:    "Configure",
				License:      "BSD-3-Clause",
				LicenseFile:  "LICENSE",
				Maintainer:   "libraries@haskell.org",
				Synopsis: []string{
					"Basic libraries",
				},
				Description: []string{
					"This package contains the Standard Haskell \"Prelude\" and its support libraries,",
					"and a large collection of useful libraries ranging from data",
					"structures to parsing combinators and debugging utilities.",
				},
				Category: "Prelude",
				Repositories: map[string]*SourceRepository{
					"head": {
						Type:     "git",
						Location: "https://gitlab.haskell.org/ghc/ghc.git",
					},
				},
				Library: &Library{
					BuildDepends: []*Dependency{
						{
							Name: "rts",
							Eq:   "1.0.*",
						},
						{
							Name: "ghc-prim",
							Gte:  "0.5.1.0",
							Lt:   "0.10",
						},
						{
							Name: "ghc-bignum",
							Gte:  "1.0",
							Lt:   "2.0",
						},
					},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f, err := os.Open(fmt.Sprintf("./test/testdata/%s", tc.filename))
			if err != nil {
				t.Fatal(err)
			}

			defer f.Close()

			actual, err := NewParser().ParseReader(f)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(tc.expected, actual) {
				t.Logf("expected: %+v", tc.expected)
				t.Log("p-----p")
				t.Logf("actual: %+v", actual)
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
			f, err := os.Open(fmt.Sprintf("./test/testdata/%s", fd.Name()))
			if err != nil {
				return nil, err
			}

			tf = append(tf, f)
		}
	}

	return tf, nil
}
