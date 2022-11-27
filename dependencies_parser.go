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
		gt   float64
		gte  float64
		lt   float64
		lte  float64
	)

	for _, v := range chunks[1:] {
		if v == "&&" || v == " " || v == "\t" || v == "" {
			continue
		}

		if sign == "" {
			switch v {
			case ">", ">=", "<", "<=":
				sign = v
			default:
				return nil, fmt.Errorf("unexpected token: %s", v)
			}

			continue
		}

		switch sign {
		case ">":
			v, err := parseFloat64(v)
			if err != nil {
				return nil, fmt.Errorf("expected valid float: %v", err)
			}

			gt = v
		case ">=":
			v, err := parseFloat64(v)
			if err != nil {
				return nil, fmt.Errorf("expected valid float: %v", err)
			}

			gte = v
		case "<":
			v, err := parseFloat64(v)
			if err != nil {
				return nil, fmt.Errorf("expected valid float: %v", err)
			}

			lt = v
		case "<=":
			v, err := parseFloat64(v)
			if err != nil {
				return nil, fmt.Errorf("expected valid float: %v", err)
			}

			lte = v
		default:
			return nil, fmt.Errorf("unexpected sign: %s", sign)
		}

		sign = ""
	}

	if gt > 0 && gte > 0 {
		return nil, fmt.Errorf("gt and gte simultaneous declaration")
	}

	if lt > 0 && lte > 0 {
		return nil, fmt.Errorf("gt and gte simultaneous declaration")
	}

	return &Dependency{
		Name:               chunks[0],
		IsLatest:           gt == 0 && gte == 0 && lt == 0 && lte == 0,
		GreaterThan:        gt,
		GreaterOrEqualThan: gte,
		LessThan:           lt,
		LessOrEqualThan:    lte,
	}, nil
}

func parseFloat64(f string) (float64, error) {
	return strconv.ParseFloat(f, 64)
}
