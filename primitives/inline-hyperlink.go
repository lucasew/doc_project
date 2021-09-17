package primitives

import "bytes"

type InlineHyperlink struct {
    children []DocumentTextInlineNode
    url string
}

func NewHyperlinkText(url string, nodes ...DocumentTextInlineNode) DocumentTextInlineNode {
    return &InlineHyperlink{
        children: nodes,
        url: url,
    }
}

func (b *InlineHyperlink) Children() []DocumentTextInlineNode {
    return b.children
}

func (b *InlineHyperlink) ExtractText() string {
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

func (InlineHyperlink) ImplInlineNode() {}
func (InlineHyperlink) ImplDocumentNode() {}
