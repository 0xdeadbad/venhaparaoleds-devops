package main

type cType int

const (
	Str cType = iota
	ListStr
	Void
	Eof
)

func (t cType) String() string {
	switch t {
	case Str:
		return "Str"
	case ListStr:
		return "ListStr"
	case Void:
		return "Void"
	case Eof:
		return "Eof"
	default:
		return "Unknown"
	}
}

type cParser struct {
	src      string
	current  int
	previous int
}

func newCParser(src string) *cParser {

	return &cParser{
		current:  0,
		previous: 0,
		src:      src,
	}
}

func (c *cParser) Next() (cType, any) {
	for c.current < len(c.src) {
		switch c.src[c.current] {
		case ']', ',', ' ', '\n', '\t':
			c.current += 1
			if c.src[c.current] == ' ' || c.src[c.current] == '\t' || c.src[c.current] == '\n' {
				c.SkipWhitespace()
			}
			return c.Next()
		case '[':
			c.current += 1
			return ListStr, c.ParseListStr()
		default:
			return Str, c.ParseStr()
		}
	}

	return Eof, nil
}

func (c *cParser) ParseStr() string {
	var str string
	for c.current < len(c.src) && c.current <= c.NextBreak() && c.src[c.current] != ',' && c.src[c.current] != '[' && c.src[c.current] != ']' && c.src[c.current] != '\n' {
		str += string(c.src[c.current])
		c.current += 1
	}

	return str
}

func (c *cParser) ParseListStr() []string {
	var lstr []string
	for c.current < len(c.src) && c.current < c.NextBreak() && c.src[c.current] != ']' && c.src[c.current] != '\n' {
		c.SkipWhitespace()
		s := c.ParseStr()
		lstr = append(lstr, s)
		c.current += 1
	}

	return lstr
}

func (c *cParser) SkipWhitespace() {
	for c.src[c.current] == ' ' || c.src[c.current] == '\t' {
		c.current += 1
	}
}

func (c *cParser) NextBreak() int {
	i := 0
	for i = c.current; i < len(c.src) && i != '\n'; {
		i += 1
	}

	return i
}
