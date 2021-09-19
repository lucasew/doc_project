package primitives

import (
	"bytes"

	"github.com/lucasew/doc_project/components"
)

type InlineBold struct {
    children []components.DocumentTextInlineNode
}

func NewBoldText(nodes ...components.DocumentTextInlineNode) components.DocumentTextInlineNode {
    return &InlineBold{
        children: nodes,
    }
}

func (InlineBold) NodeKind() string {
    return "inline-bold"
}

func (b *InlineBold) Children() []components.DocumentTextInlineNode {
    return b.children
}

func (b *InlineBold) ExtractText() string {
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

func (InlineBold) ImplInlineNode() {}
func (InlineBold) ImplDocumentNode() {}
