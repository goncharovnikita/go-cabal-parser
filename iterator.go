package gocabalparser

type tokensIterator struct {
	tokens tokens
	curr   *token
	index  int
}

func newTokensIterator(tokens tokens) *tokensIterator {
	return &tokensIterator{
		tokens: tokens,
		curr:   nil,
		index:  0,
	}
}

func (it *tokensIterator) Next() bool {
	if it.index >= len(it.tokens) {
		return false
	}

	it.curr = it.tokens[it.index]
	it.index++

	return true
}

func (it *tokensIterator) Seek() (*token, bool) {
	if it.index >= len(it.tokens) {
		return nil, false
	}

	return it.tokens[it.index], true
}

func (it *tokensIterator) Val() *token {
	return it.curr
}
