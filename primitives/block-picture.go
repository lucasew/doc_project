package primitives

import (
    "github.com/lucasew/doc_project/providers"
)

type BlockPicture struct {
    image providers.Image
}

func NewBlockPicture(image providers.Image) DocumentBlockNode {
    return &BlockPicture{image: image}
}

func (BlockPicture) Children() []DocumentBlockNode {
    return NewEmptyDocumentBlockNodeList()
}

func (BlockPicture) ImplBlockNode() {}
func (BlockPicture) ImplDocumentNode() {}

