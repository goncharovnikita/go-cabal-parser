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
