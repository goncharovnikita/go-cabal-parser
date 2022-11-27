package gocabalparser

import (
	"io"
)

type Library struct {
	BuildDepends []*Dependency
}

type SourceRepository struct {
	Type     string
	Location string
	Tag      string
}

type Dependency struct {
	Name     string
	IsLatest bool
	Eq       string
	Gt       string
	Gte      string
	Lt       string
	Lte      string
}

type Executable struct {
	BuildDepends []*Dependency
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
	Library      *Library
}

type Parser interface {
	ParseReader(r io.Reader) (*CabalPackage, error)
}

type parser struct{}

func NewParser() Parser {
	return &parser{}
}

func (p *parser) ParseReader(r io.Reader) (*CabalPackage, error) {
	tokens, err := newTokenizer().TokenizeReader(r)
	if err != nil {
		return nil, err
	}

	return newTokensParser().Parse(tokens)
}
