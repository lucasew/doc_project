package primitives

import (
	"bytes"

	"github.com/lucasew/doc_project/components"
)

type InlineStrikethrough struct {
    children []components.DocumentTextInlineNode
}

func NewStrikethroughText(nodes ...components.DocumentTextInlineNode) components.DocumentTextInlineNode {
    return &InlineStrikethrough{
        children: nodes,
    }
}

func (InlineStrikethrough) NodeKind() string {
    return "inline-strikethrough"
}

func (b *InlineStrikethrough) Children() []components.DocumentTextInlineNode {
    return b.children
}

func (b *InlineStrikethrough) ExtractText() string {
    if len(b.children) == 0 {
        return ""
    }
    if len(b.children) == 1 {
        return b.children[0].ExtractText()
    }
    var err error
    buf := bytes.NewBufferString("")
    for _, child := range(b.children) {
        _, err = buf.WriteString(" ")
        if err != nil {
            panic(err)
        }
        _, err = buf.WriteString(child.ExtractText())
        if err != nil {
            panic(err)
        }
    }
    return buf.String()
}

func (InlineStrikethrough) ImplInlineNode() {}
func (InlineStrikethrough) ImplDocumentNode() {}
