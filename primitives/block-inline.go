package primitives

type BlockInline struct {
    children []DocumentInlineNode
}

// NewBlockInline Create a text block
func NewBlockInline(children ...DocumentInlineNode) DocumentBlockNode {
    return &BlockInline{
        children: children,
    }
}

// Children returns empty always because the children is of inline nodes
func (n *BlockInline) Children() []DocumentBlockNode {
    return NewEmptyDocumentBlockNodeList()
}

func (BlockInline) ImplBlockNode() {}
func (BlockInline) ImplDocumentNode() {}
