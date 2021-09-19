package primitives

import "github.com/lucasew/doc_project/components"

type InlineTextNode struct {
    children []components.DocumentTextInlineNode
}

func NewInlineTextNode(children ...components.DocumentTextInlineNode) components.DocumentInlineNode {
    return &InlineTextNode{children: children}
}

func (InlineTextNode) NodeKind() string {
    return "inline-text-node"
}

func (InlineTextNode) Children() []components.DocumentInlineNode {
    return components.NewEmptyDocumentInlineNodeList()
}

func (InlineTextNode) ImplInlineNode() {}
func (InlineTextNode) ImplDocumentNode() {}
