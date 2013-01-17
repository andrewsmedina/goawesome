package goawesome

import (
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})

func (s *S) TestNumber(c *C) {
	c.Assert(tokenizer([]byte("1")), DeepEquals, []token{{"type": "NUMBER", "value": "1"}})
}

func (s *S) TestOperator(c *C) {
	c.Assert(tokenizer([]byte("+")), DeepEquals, []token{{"type": "+", "value": "+"}})
	c.Assert(tokenizer([]byte("||")), DeepEquals, []token{{"type": "||", "value": "||"}})
}
