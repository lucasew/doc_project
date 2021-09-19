package primitives

import (
    "github.com/lucasew/doc_project/components"
)

type BlockInline struct {
    children []components.DocumentInlineNode
}

// NewBlockInline Create a text block
func NewBlockInline(children ...components.DocumentInlineNode) components.DocumentBlockNode {
    return &BlockInline{
        children: children,
    }
}

func (BlockInline) NodeKind() string {
    return "block-inline"
}

// Children returns empty always because the children is of inline nodes
func (n *BlockInline) Children() []components.DocumentBlockNode {
    return components.NewEmptyDocumentBlockNodeList()
}

func (BlockInline) ImplBlockNode() {}
func (BlockInline) ImplDocumentNode() {}
