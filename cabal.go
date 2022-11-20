package gocabalparser

import (
	"bufio"
	"errors"
	"fmt"
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

type tokenType int

const (
	tokenName tokenType = iota
	tokenValue
	tokenScopeType
	tokenScopeName
	tokenScopeValueName
	tokenScopeValueValue
)

func (t tokenType) String() string {
	switch t {
	case tokenName:
		return "Name"
	case tokenValue:
		return "Value"
	case tokenScopeType:
		return "ScopeType"
	case tokenScopeName:
		return "ScopeName"
	case tokenScopeValueName:
		return "ScopeValueName"
	case tokenScopeValueValue:
		return "ScopeValueValue"
	default:
		return fmt.Sprintf("unknown token: %d", t)
	}
}

type tokenizerState int

const (
	tokenizerStateInit tokenizerState = iota
	tokenizerStateName
	tokenizerStateValueStart
	tokenizerStateValue
	tokenizerStateScopeStart
	tokenizerStateScopeEntryInit
	tokenizerStateScopeEntryNameStart
	tokenizerStateScopeEntryName
	tokenizerStateScopeEntryValueStart
	tokenizerStateScopeEntryValue
)

func (s tokenizerState) String() string {
	switch s {
	case tokenizerStateInit:
		return "TokenizerStateInit"
	case tokenizerStateName:
		return "TokenizerStateName"
	case tokenizerStateValueStart:
		return "TokenizerStateValueStart"
	case tokenizerStateValue:
		return "TokenizerStateValue"
	case tokenizerStateScopeStart:
		return "TokenizerStateScopeStart"
	case tokenizerStateScopeEntryInit:
		return "TokenizerStateScopeEntryInit"
	case tokenizerStateScopeEntryNameStart:
		return "TokenizerStateScopeEntryNameStart"
	case tokenizerStateScopeEntryName:
		return "TokenizerStateScopeEntryName"
	case tokenizerStateScopeEntryValueStart:
		return "TokenizerStateScopeEntryValueStart"
	case tokenizerStateScopeEntryValue:
		return "TokenizerStateScopeEntryValue"
	default:
		return fmt.Sprintf("unknown state: %d", s)
	}
}

type token struct {
	Type  tokenType
	Value string
}

type tokens []*token

func (t tokens) String() string {
	r := ""

	for _, v := range t {
		r += fmt.Sprintf("[%s: %s]\n", v.Type, v.Value)
	}

	return r
}

type tokenizer struct {
	states []tokenizerState
}

func newTokenizer() *tokenizer {
	return &tokenizer{
		states: []tokenizerState{
			tokenizerStateInit,
		},
	}
}

func (t *tokenizer) TokenizeReader(r io.Reader) (tokens, error) {
	res := make(tokens, 0)
	state := tokenizerStateInit

	buf := make([]byte, 4096)
	val := make([]byte, 0)
	index := 0

	for {
		readn, err := r.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, err
		}

		for index < readn {
			v := buf[index]
			if state != t.states[len(t.states)-1] {
				t.states = append(t.states, state)
			}

			switch state {
			case tokenizerStateInit:
				{
					switch v {
					case '\t', ' ':
						state = tokenizerStateValueStart
					default:
						if v != '\n' {
							state = tokenizerStateName
							val = append(val, v)
						}
					}
				}

			case tokenizerStateName:
				{
					switch v {
					case ':':
						{
							t := &token{
								Type:  tokenName,
								Value: string(val),
							}

							res = append(res, t)
							state = tokenizerStateValueStart
							val = make([]byte, 0)
						}
					case ' ':
						{
							t := &token{
								Type:  tokenScopeType,
								Value: string(val),
							}

							res = append(res, t)
							state = tokenizerStateScopeStart
							val = make([]byte, 0)
						}
					default:
						{
							val = append(val, v)
						}
					}
				}

			case tokenizerStateValueStart:
				{
					if v != ' ' && v != '\t' && v != '\n' {
						val = append(val, v)
						state = tokenizerStateValue
					}
				}

			case tokenizerStateValue:
				{
					if v == '\n' {
						t := &token{
							Type:  tokenValue,
							Value: string(val),
						}

						res = append(res, t)
						state = tokenizerStateInit
						val = make([]byte, 0)
					} else {
						val = append(val, v)
					}
				}

			case tokenizerStateScopeStart:
				{
					if v == '\n' {
						t := &token{
							Type:  tokenScopeName,
							Value: string(val),
						}

						res = append(res, t)
						state = tokenizerStateScopeEntryInit
						val = make([]byte, 0)
					} else {
						val = append(val, v)
					}
				}

			case tokenizerStateScopeEntryInit:
				{
					if v == ' ' || v == '\t' {
						state = tokenizerStateScopeEntryNameStart
					} else {
						if v != '\n' {
							val = append(val, v)
						}

						state = tokenizerStateInit
					}
				}

			case tokenizerStateScopeEntryNameStart:
				{
					if v != ' ' && v != '\t' {
						state = tokenizerStateScopeEntryName
						val = append(val, v)
					}
				}

			case tokenizerStateScopeEntryName:
				{
					switch v {
					case ':':
						{
							t := &token{
								Type:  tokenScopeValueName,
								Value: string(val),
							}

							res = append(res, t)
							state = tokenizerStateScopeEntryValueStart
							val = make([]byte, 0)
						}
					case '\n':
						{
							t := &token{
								Type:  tokenScopeValueValue,
								Value: string(val),
							}

							res = append(res, t)
							state = tokenizerStateScopeEntryInit
							val = make([]byte, 0)
						}
					default:
						val = append(val, v)
					}
				}

			case tokenizerStateScopeEntryValueStart:
				{
					if v != ' ' && v != '\t' && v != '\n' {
						state = tokenizerStateScopeEntryValue
						val = append(val, v)
					}
				}

			case tokenizerStateScopeEntryValue:
				{
					if v == '\n' || v == ',' {
						t := &token{
							Type:  tokenScopeValueValue,
							Value: string(val),
						}

						res = append(res, t)
						val = make([]byte, 0)

						if v == '\n' {
							state = tokenizerStateScopeEntryInit
						} else {
							state = tokenizerStateScopeEntryValueStart
						}
					} else {
						val = append(val, v)
					}
				}
			default:
				val = append(val, v)
			}

			index++
		}

		index = 0
	}

	if state != t.states[len(t.states)-1] {
		t.states = append(t.states, state)
	}

	return res, nil
}
