package gocabalparser

import (
	"bufio"
	"io"
	"strings"
)

type SourceRepository struct {
	Type     string
	Location string
	Tag      string
}

type Executable struct {
	BuildDepends []string
	Extensions   []string
	MainIs       string
	OtherModules []string
	HSSourceDirs []string
}

type CabalPackage struct {
	Name         string
	Version      string
	CabalVersion string
	BuildType    string
	License      string
	LicenseFile  string
	Copyright    []string
	Author       string
	Maintainer   string
	Stability    string
	Homepage     string
	PackageURL   string
	Synopsis     []string
	Description  []string
	Category     string
	TestedWith   string
	Repositories map[string]*SourceRepository
	Executables  map[string]*Executable
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
