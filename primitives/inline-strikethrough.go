package primitives

import "bytes"

type InlineStrikethrough struct {
    children []DocumentTextInlineNode
}

func NewStrikethroughText(nodes ...DocumentTextInlineNode) DocumentTextInlineNode {
    return &InlineStrikethrough{
        children: nodes,
    }
}

func (b *InlineStrikethrough) Children() []DocumentTextInlineNode {
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
