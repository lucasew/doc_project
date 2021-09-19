package primitives

import (
	"bytes"

	"github.com/lucasew/doc_project/components"
)

type InlineItalic struct {
    children []components.DocumentTextInlineNode
}

func NewItalicText(nodes ...components.DocumentTextInlineNode) components.DocumentTextInlineNode {
    return &InlineItalic{
        children: nodes,
    }
}

func (InlineItalic) NodeKind() string {
    return "inline-italic"
}

func (b *InlineItalic) Children() []components.DocumentTextInlineNode {
    return b.children
}

func (b *InlineItalic) ExtractText() string {
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

func (InlineItalic) ImplInlineNode() {}
func (InlineItalic) ImplDocumentNode() {}
