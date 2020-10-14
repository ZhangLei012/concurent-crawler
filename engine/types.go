package engine

// Request is the parser request
type Request struct {
	URL        string
	ParserFunc func([]byte) ParserResult
}

// ParserResult is the parser result
type ParserResult struct {
	Requests []Request
	Items    []interface{}
}

// NilParser does nothing
func NilParser(content []byte) ParserResult {
	return ParserResult{}
}
