package primitives

type InlineTextNode struct {
    children []DocumentTextInlineNode
}

func NewInlineTextNode(children ...DocumentTextInlineNode) DocumentInlineNode {
    return &InlineTextNode{children: children}
}

func (InlineTextNode) Children() []DocumentInlineNode {
    return NewEmptyDocumentInlineNodeList()
}

func (InlineTextNode) ImplInlineNode() {}
func (InlineTextNode) ImplDocumentNode() {}
