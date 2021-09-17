package primitives

import "bytes"

type InlineUnderline struct {
    children []DocumentTextInlineNode
}

func NewUnderlineText(nodes ...DocumentTextInlineNode) DocumentTextInlineNode {
    return &InlineUnderline{
        children: nodes,
    }
}

func (b *InlineUnderline) Children() []DocumentTextInlineNode {
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
