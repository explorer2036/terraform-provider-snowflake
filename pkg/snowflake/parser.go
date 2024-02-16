package snowflake

import (
	"fmt"
	"strings"
	"unicode"
)

// ViewSelectStatementExtractor is a simplistic parser that only exists to extract the select
// statement from a create view statement
//
// The implementation is optimized for undertandable and predictable behavior. So far we only seek
// to support queries of the sort that are generated by this project.
//
// Also there is little error handling and we assume queries are well-formed.
type ViewSelectStatementExtractor struct {
	input []rune
	pos   int
}

func NewViewSelectStatementExtractor(input string) *ViewSelectStatementExtractor {
	return &ViewSelectStatementExtractor{
		input: []rune(input),
	}
}

func (e *ViewSelectStatementExtractor) Extract() (string, error) {
	fmt.Printf("[DEBUG] extracting view query %s\n", string(e.input))
	e.consumeSpace()
	e.consumeToken("create")
	e.consumeSpace()
	e.consumeToken("or replace")
	e.consumeSpace()
	e.consumeToken("secure")
	e.consumeSpace()
	e.consumeToken("recursive")
	e.consumeSpace()
	e.consumeToken("view")
	e.consumeSpace()
	e.consumeToken("if not exists")
	e.consumeSpace()
	e.consumeID()
	// TODO column list
	e.consumeSpace()
	e.consumeToken("copy grants")
	e.consumeComment()
	e.consumeSpace()
	e.consumeComment()
	e.consumeSpace()
	e.consumeToken("as")
	e.consumeSpace()

	return string(e.input[e.pos:]), nil
}

func (e *ViewSelectStatementExtractor) ExtractMaterializedView() (string, error) {
	fmt.Printf("[DEBUG] extracting materialized view query: %s\n", string(e.input))
	e.consumeSpace()
	e.consumeToken("use warehouse")
	e.consumeSpace()
	e.consumeNonSpace() // warehouse name
	e.consumeSpace()
	e.consumeToken("create")
	e.consumeSpace()
	e.consumeToken("or replace")
	e.consumeSpace()
	e.consumeToken("secure")
	e.consumeSpace()
	e.consumeToken("materialized view")
	e.consumeSpace()
	e.consumeToken("if not exists")
	e.consumeSpace()
	e.consumeID()
	// TODO copy grants
	// TODO column list
	e.consumeComment()
	e.consumeSpace()
	e.consumeComment()
	e.consumeSpace()
	e.consumeToken("cluster by")
	e.consumeSpace()
	e.consumeClusterBy()
	e.consumeSpace()
	e.consumeToken("as")
	e.consumeSpace()

	return string(e.input[e.pos:]), nil
}

func (e *ViewSelectStatementExtractor) ExtractDynamicTable() (string, error) {
	fmt.Printf("[DEBUG] extracting dynamic table query %s\n", string(e.input))
	e.consumeSpace()
	e.consumeToken("create")
	e.consumeSpace()
	e.consumeToken("or replace")
	e.consumeSpace()
	e.consumeToken("dynamic table")
	e.consumeSpace()
	e.consumeID()
	// TODO column list
	e.consumeSpace()
	e.consumeComment()
	e.consumeSpace()
	e.consumeQuotedParameter("lag")
	e.consumeSpace()
	e.consumeTokenParameter("warehouse")
	e.consumeSpace()
	e.consumeTokenParameter("refresh_mode")
	e.consumeSpace()
	e.consumeTokenParameter("initialize")
	e.consumeSpace()
	e.consumeTokenParameter("warehouse")
	e.consumeSpace()
	e.consumeComment()
	e.consumeSpace()
	e.consumeToken("as")
	e.consumeSpace()

	return string(e.input[e.pos:]), nil
}

// consumeToken will move e.pos forward iff the token is the next part of the input. Comparison is
// case-insensitive. Will return true if consumed.
func (e *ViewSelectStatementExtractor) consumeToken(t string) bool {
	found := 0
	for i, r := range t {
		// it is annoying that we have to convert the runes back to strings to do a case-insensitive
		// comparison. Hopefully I am just missing something in the docs.
		if e.pos+i > len(e.input) || !strings.EqualFold(string(r), string(e.input[e.pos+i])) {
			break
		}
		found++
	}

	if found == len(t) {
		e.pos += len(t)
		return true
	}
	return false
}

func (e *ViewSelectStatementExtractor) consumeSpace() {
	found := 0
	for {
		if e.pos+found > len(e.input)-1 || !unicode.IsSpace(e.input[e.pos+found]) {
			break
		}
		found++
	}
	e.pos += found
}

func (e *ViewSelectStatementExtractor) consumeID() {
	e.consumeNonSpace()
}

func (e *ViewSelectStatementExtractor) consumeNonSpace() {
	found := 0
	for {
		if e.pos+found > len(e.input)-1 || unicode.IsSpace(e.input[e.pos+found]) {
			break
		}
		found++
	}
	e.pos += found
}

func (e *ViewSelectStatementExtractor) consumeComment() {
	e.consumeQuotedParameter("comment")
}

func (e *ViewSelectStatementExtractor) consumeQuotedParameter(param string) {
	if c := e.consumeToken(param); !c {
		return
	}

	e.consumeSpace()

	if c := e.consumeToken("="); !c {
		return
	}

	e.consumeSpace()

	if c := e.consumeToken("'"); !c {
		return
	}

	found := 0
	escaped := false
	for {
		if e.pos+found > len(e.input)-1 {
			break
		}

		if escaped { //nolint:gocritic // todo: please fix this to pass gocritic
			escaped = false
		} else if e.input[e.pos+found] == '\\' {
			escaped = true
		} else if e.input[e.pos+found] == '\'' {
			break
		}
		found++
	}
	e.pos += found

	if !e.consumeToken("'") {
		return
	}
}

func (e *ViewSelectStatementExtractor) consumeTokenParameter(param string) {
	if c := e.consumeToken(param); !c {
		return
	}

	e.consumeSpace()

	if c := e.consumeToken("="); !c {
		return
	}

	e.consumeSpace()
	e.consumeNonSpace()
}

func (e *ViewSelectStatementExtractor) consumeClusterBy() {
	if e.input[e.pos] != '(' {
		return
	}
	found := 1
	for {
		if e.pos+found > len(e.input)-1 || e.input[e.pos+found-1] == ')' {
			break
		}
		found++
	}
	e.pos += found
}
