package primitives

import (
	"bytes"

	"github.com/lucasew/doc_project/components"
)

type InlineUnderline struct {
    children []components.DocumentTextInlineNode
}

func NewUnderlineText(nodes ...components.DocumentTextInlineNode) components.DocumentTextInlineNode {
    return &InlineUnderline{
        children: nodes,
    }
}

func (InlineUnderline) NodeKind() string {
    return "inline-underline"
}

func (b *InlineUnderline) Children() []components.DocumentTextInlineNode {
    return b.children
}

func (b *InlineUnderline) ExtractText() string {
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

func (InlineUnderline) ImplInlineNode() {}
func (InlineUnderline) ImplDocumentNode() {}
