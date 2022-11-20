package gocabalparser

import (
	"bufio"
	"io"
	"strings"
)

type CabalPackage struct {
	Name string
}

type Parser interface {
	ParseReader(r io.Reader) (*CabalPackage, error)
}

type parser struct{}

func NewParser() Parser {
	return &parser{}
}

func (p *parser) ParseReader(r io.Reader) (*CabalPackage, error) {
	s := bufio.NewScanner(r)

	res := &CabalPackage{}

	for s.Scan() {

	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func (p *parser) parseLine(l string) {
	if strings.HasPrefix(l, "Name:") {

	}
}
