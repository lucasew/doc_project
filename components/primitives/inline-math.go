package primitives

import (
	"fmt"

	"github.com/lucasew/doc_project/components"
)

type InlineMath struct {
    stmt string
}

func NewMathText(stmt string) components.DocumentTextInlineNode {
    return &InlineMath{
        stmt: stmt,
    }
}

func (InlineMath) NodeKind() string {
    return "inline-math"
}

func (b *InlineMath) Children() []components.DocumentTextInlineNode {
    return []components.DocumentTextInlineNode{}
}

func (b *InlineMath) ExtractText() string {
    return fmt.Sprintf("$ %s $", b.stmt)
}

func (InlineMath) ImplInlineNode() {}
func (InlineMath) ImplDocumentNode() {}
