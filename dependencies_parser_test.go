package gocabalparser

import (
	"reflect"
	"testing"
)

func TestDependenciesParser_ParseString(t *testing.T) {
	cases := []struct {
		name             string
		dependencyString string
		expected         *Dependency
	}{
		{
			name:             "latest version",
			dependencyString: "base",
			expected: &Dependency{
				Name:     "base",
				IsLatest: true,
			},
		},
		{
			name:             "greater than",
			dependencyString: "base > 1.0",
			expected: &Dependency{
				Name:        "base",
				IsLatest:    false,
				GreaterThan: 1.0,
			},
		},
		{
			name:             "less than",
			dependencyString: "base < 1.0",
			expected: &Dependency{
				Name:     "base",
				IsLatest: false,
				LessThan: 1.0,
			},
		},
		{
			name:             "greater or equal",
			dependencyString: "base >= 1.0",
			expected: &Dependency{
				Name:               "base",
				IsLatest:           false,
				GreaterOrEqualThan: 1.0,
			},
		},
		{
			name:             "less or equal",
			dependencyString: "base <= 1.0",
			expected: &Dependency{
				Name:            "base",
				IsLatest:        false,
				LessOrEqualThan: 1.0,
			},
		},
		{
			name:             "greater and less",
			dependencyString: "base > 1.0 && < 2.0",
			expected: &Dependency{
				Name:        "base",
				IsLatest:    false,
				GreaterThan: 1.0,
				LessThan:    2.0,
			},
		},
		{
			name:             "greater or equal and less",
			dependencyString: "base >= 1.0 && < 2.0",
			expected: &Dependency{
				Name:               "base",
				IsLatest:           false,
				GreaterOrEqualThan: 1.0,
				LessThan:           2.0,
			},
		},
		{
			name:             "greater and less or equal",
			dependencyString: "base > 1.0 && <= 2.0",
			expected: &Dependency{
				Name:            "base",
				IsLatest:        false,
				GreaterThan:     1.0,
				LessOrEqualThan: 2.0,
			},
		},
		{
			name:             "greater or equal and less or equal",
			dependencyString: "base >= 1.0 && <= 2.0",
			expected: &Dependency{
				Name:               "base",
				IsLatest:           false,
				GreaterOrEqualThan: 1.0,
				LessOrEqualThan:    2.0,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := newDependenciesParser().ParseString(tc.dependencyString)

			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fatal("actual result don't match expectations")
			}
		})
	}
}
