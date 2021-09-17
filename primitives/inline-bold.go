package primitives

import "bytes"

type InlineBold struct {
    children []DocumentTextInlineNode
}

func NewBoldText(nodes ...DocumentTextInlineNode) DocumentTextInlineNode {
    return &InlineBold{
        children: nodes,
    }
}

func (b *InlineBold) Children() []DocumentTextInlineNode {
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
