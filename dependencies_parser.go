package gocabalparser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type dependenciesParser struct{}

func newDependenciesParser() *dependenciesParser {
	return &dependenciesParser{}
}

func (p *dependenciesParser) ParseString(s string) (*Dependency, error) {
	chunks := strings.Split(s, " ")

	if len(chunks) == 0 {
		return nil, errors.New("empty dependency")
	}

	var (
		sign string
		eq   string
		gt   string
		gte  string
		lt   string
		lte  string
	)

	for _, v := range chunks[1:] {
		if v == "&&" || v == " " || v == "\t" || v == "" {
			continue
		}

		if sign == "" {
			switch v {
			case "==", ">", ">=", "<", "<=":
				sign = v
			default:
				return nil, fmt.Errorf("unexpected token: %s", v)
			}

			continue
		}

		switch sign {
		case ">":
			gt = v
		case ">=":
			gte = v
		case "<":
			lt = v
		case "<=":
			lte = v
		case "==":
			eq = v
		default:
			return nil, fmt.Errorf("unexpected sign: %s", sign)
		}

		sign = ""
	}

	if gt != "" && gte != "" {
		return nil, fmt.Errorf("gt and gte simultaneous declaration")
	}

	if lt != "" && lte != "" {
		return nil, fmt.Errorf("gt and gte simultaneous declaration")
	}

	return &Dependency{
		Name:     chunks[0],
		IsLatest: eq != "" && gt != "" && gte != "" && lt != "" && lte != "",
		Eq:       eq,
		Gt:       gt,
		Gte:      gte,
		Lt:       lt,
		Lte:      lte,
	}, nil
}

func parseFloat64(f string) (float64, error) {
	return strconv.ParseFloat(f, 64)
}
