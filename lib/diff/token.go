package diff

import (
	"context"
	"fmt"
	"strings"
)

// Type of token
type Type int

const (
	// Newline type
	Newline Type = iota
	// Space type
	Space

	// ChunkStart type
	ChunkStart
	// ChunkEnd type
	ChunkEnd

	// SameSymbol type
	SameSymbol
	// SameLine type
	SameLine

	// AddSymbol type
	AddSymbol
	// AddLine type
	AddLine

	// DelSymbol typ
	DelSymbol
	// DelLine type
	DelLine

	// SameWords type
	SameWords
	// AddWords type
	AddWords
	// DelWords type
	DelWords

	// EmptyLine type
	EmptyLine
)

// Token presents a symbol in diff layout
type Token struct {
	Type    Type
	Literal string
}

// TokenizeText text block a and b into diff tokens.
func TokenizeText(ctx context.Context, x, y string) []*Token {
	xls := NewText(x) // x lines
	yls := NewText(y) // y lines

	s := xls.LCS(ctx, yls)

	ts := []*Token{}

	xNum, yNum, sNum := numFormat(xls, yls)

	for i, j, k := 0, 0, 0; i < len(xls) || j < len(yls); {
		if i < len(xls) && (k == len(s) || neq(xls[i], s[k])) {
			ts = append(ts,
				&Token{DelSymbol, fmt.Sprintf(xNum, i+1) + "-"},
				&Token{Space, " "},
				&Token{DelLine, xls[i].String()},
				&Token{Newline, "\n"})
			i++
		} else if j < len(yls) && (k == len(s) || neq(yls[j], s[k])) {
			ts = append(ts,
				&Token{AddSymbol, fmt.Sprintf(yNum, j+1) + "+"},
				&Token{Space, " "},
				&Token{AddLine, yls[j].String()},
				&Token{Newline, "\n"})
			j++
		} else {
			ts = append(ts,
				&Token{SameSymbol, fmt.Sprintf(sNum, i+1, j+1) + " "},
				&Token{Space, " "},
				&Token{SameLine, s[k].String() + "\n"})
			i, j, k = i+1, j+1, k+1
		}
	}

	return ts
}

// TokenizeLine two different lines
func TokenizeLine(ctx context.Context, x, y string) ([]*Token, []*Token) {
	xs := NewString(x)
	ys := NewString(y)

	s := xs.LCS(ctx, ys)

	xTokens := []*Token{}
	yTokens := []*Token{}

	for i, j, k := 0, 0, 0; i < len(xs) || j < len(ys); {
		if i < len(xs) && (k == len(s) || neq(xs[i], s[k])) {
			xTokens = append(xTokens, &Token{DelWords, xs[i].String()})
			i++
		} else if j < len(ys) && (k == len(s) || neq(ys[j], s[k])) {
			yTokens = append(yTokens, &Token{AddWords, ys[j].String()})
			j++
		} else {
			xTokens = append(xTokens, &Token{SameWords, s[k].String()})
			yTokens = append(yTokens, &Token{SameWords, s[k].String()})
			i, j, k = i+1, j+1, k+1
		}
	}

	return xTokens, yTokens
}

func numFormat(x, y []Comparable) (string, string, string) {
	xl := len(fmt.Sprintf("%d", len(x)))
	yl := len(fmt.Sprintf("%d", len(y)))

	return fmt.Sprintf("%%0%dd "+strings.Repeat(" ", yl+1), xl),
		fmt.Sprintf(strings.Repeat(" ", xl)+" %%0%dd ", yl),
		fmt.Sprintf("%%0%dd %%0%dd ", xl, yl)
}
