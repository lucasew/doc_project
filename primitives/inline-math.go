package primitives

import (
	"fmt"
)

type InlineMath struct {
    stmt string
}

func NewMathText(stmt string) DocumentTextInlineNode {
    return &InlineMath{
        stmt: stmt,
    }
}

func (b *InlineMath) Children() []DocumentTextInlineNode {
    return []DocumentTextInlineNode{}
}

func (b *InlineMath) ExtractText() string {
    return fmt.Sprintf("$ %s $", b.stmt)
}

func (InlineMath) ImplInlineNode() {}
func (InlineMath) ImplDocumentNode() {}
