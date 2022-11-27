package gocabalparser

import (
	"errors"
	"fmt"
	"strings"
)

var (
	repoProperties = map[string]struct{}{
		"type":     {},
		"location": {},
		"tag":      {},
	}

	executableProperties = map[string]struct{}{
		"build-depends":  {},
		"extensions":     {},
		"main-is":        {},
		"other-modules":  {},
		"hs-source-dirs": {},
	}
)

type tokensParser struct{}

func newTokensParser() *tokensParser {
	return &tokensParser{}
}

func (p *tokensParser) Parse(tokens []*token) (*CabalPackage, error) {
	iterator := newTokensIterator(tokens)
	res := &CabalPackage{}

	for iterator.Next() {
		token := iterator.Val()

		if token.Type != tokenTypeKey {
			return nil, fmt.Errorf("name declaration expected, but got: %s", token.Value)
		}

		var err error

		switch strings.ToLower(token.Value) {
		case "name":
			err = parseString(&res.Name, iterator)
		case "version":
			err = parseString(&res.Version, iterator)
		case "cabal-version":
			err = parseString(&res.CabalVersion, iterator)
		case "build-type":
			err = parseString(&res.BuildType, iterator)
		case "license":
			err = parseString(&res.License, iterator)
		case "license-file":
			err = parseString(&res.LicenseFile, iterator)
		case "author":
			err = parseString(&res.Author, iterator)
		case "maintainer":
			err = parseString(&res.Maintainer, iterator)
		case "stability":
			err = parseString(&res.Stability, iterator)
		case "homepage":
			err = parseString(&res.Homepage, iterator)
		case "package-url":
			err = parseString(&res.PackageURL, iterator)
		case "category":
			err = parseString(&res.Category, iterator)
		case "tested-with":
			err = parseString(&res.TestedWith, iterator)
		case "copyright":
			err = parseStringArr(&res.Copyright, iterator)
		case "description":
			err = parseStringArr(&res.Description, iterator)
		case "synopsis":
			err = parseStringArr(&res.Synopsis, iterator)
		case "source-repository":
			if res.Repositories == nil {
				res.Repositories = make(map[string]*SourceRepository)
			}

			err = parseRepository(res.Repositories, iterator)
		case "executable":
			if res.Executables == nil {
				res.Executables = make(map[string]*Executable)
			}

			err = parseExecutable(res.Executables, iterator)
		default:
			return nil, fmt.Errorf("unsupported property: %s", token.Value)
		}

		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func parseString(to *string, iterator *tokensIterator) error {
	if !iterator.Next() {
		return errors.New("property value expected")
	}

	token := iterator.Val()
	if token.Type != tokenTypeValue {
		return errors.New("property value expected")
	}

	*to = token.Value

	return nil
}

func parseStringArr(to *[]string, iterator *tokensIterator) error {
	nextToken, ok := iterator.Seek()
	if !ok || nextToken.Type != tokenTypeValue {
		return errors.New("array value expected")
	}

	for {
		token, ok := iterator.Seek()
		if !ok || token.Type != tokenTypeValue {
			break
		}

		*to = append(*to, token.Value)
		iterator.Next()
	}

	return nil
}

func parseDependencies(to *[]*Dependency, iterator *tokensIterator) error {
	stringDeps := make([]string, 0)

	if err := parseStringArr(&stringDeps, iterator); err != nil {
		return err
	}

	p := newDependenciesParser()

	for _, d := range stringDeps {
		dep, err := p.ParseString(d)
		if err != nil {
			return err
		}

		*to = append(*to, dep)
	}

	return nil
}

func parseRepository(to map[string]*SourceRepository, iterator *tokensIterator) error {
	if !iterator.Next() {
		return errors.New("repository name expected")
	}

	token := iterator.Val()
	if token.Type != tokenTypeScopeName {
		return errors.New("repository name expected")
	}

	repo := &SourceRepository{}
	repoName := token.Value

	for {
		token, ok := iterator.Seek()
		if !ok || !isRepoProperty(token) {
			break
		}

		iterator.Next()

		var err error

		switch strings.ToLower(token.Value) {
		case "type":
			err = parseString(&repo.Type, iterator)
		case "location":
			err = parseString(&repo.Location, iterator)
		case "tag":
			err = parseString(&repo.Tag, iterator)
		default:
			return fmt.Errorf("unsupported repo property: '%s'", token.Value)
		}

		if err != nil {
			return err
		}
	}

	to[repoName] = repo

	return nil
}

func parseExecutable(to map[string]*Executable, iterator *tokensIterator) error {
	if !iterator.Next() {
		return errors.New("executable name expected")
	}

	token := iterator.Val()
	if token.Type != tokenTypeScopeName {
		return errors.New("executable name expected")
	}

	ex := &Executable{}
	exName := token.Value

	for {
		token, ok := iterator.Seek()
		if !ok || !isExecutableProperty(token) {
			break
		}

		iterator.Next()

		var err error

		switch strings.ToLower(token.Value) {
		case "build-depends":
			err = parseDependencies(&ex.BuildDepends, iterator)
		case "extensions":
			err = parseStringArr(&ex.Extensions, iterator)
		case "main-is":
			err = parseString(&ex.MainIs, iterator)
		case "other-modules":
			err = parseStringArr(&ex.OtherModules, iterator)
		case "hs-source-dirs":
			err = parseStringArr(&ex.HSSourceDirs, iterator)
		default:
			return fmt.Errorf("unsupported executable property: '%s'", token.Value)
		}

		if err != nil {
			return err
		}
	}

	to[exName] = ex

	return nil
}

func isRepoProperty(t *token) bool {
	if t.Type != tokenTypeKey {
		return false
	}

	_, ok := repoProperties[strings.ToLower(t.Value)]

	return ok
}

func isExecutableProperty(t *token) bool {
	if t.Type != tokenTypeKey {
		return false
	}

	_, ok := executableProperties[strings.ToLower(t.Value)]

	return ok
}
