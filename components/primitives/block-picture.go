package primitives

import (
	"github.com/lucasew/doc_project/components"
	"github.com/lucasew/doc_project/providers"
)

type BlockPicture struct {
    image providers.Image
}

func NewBlockPicture(image providers.Image) components.DocumentBlockNode {
    return &BlockPicture{image: image}
}

func (BlockPicture) NodeKind() string {
    return "block-picture"
}

func (BlockPicture) Children() []components.DocumentBlockNode {
    return components.NewEmptyDocumentBlockNodeList()
}

func (BlockPicture) ImplBlockNode() {}
func (BlockPicture) ImplDocumentNode() {}
