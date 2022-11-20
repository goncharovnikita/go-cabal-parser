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
