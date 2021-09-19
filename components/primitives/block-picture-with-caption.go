package primitives

import (
	"github.com/lucasew/doc_project/components"
	"github.com/lucasew/doc_project/providers"
)

type BlockPictureWithCaption struct {
    image providers.Image
    caption string
}

func NewBlockPictureWithCaption(image providers.Image, caption string) components.DocumentBlockNode {
    return &BlockPictureWithCaption{
        image: image,
        caption: caption,
    }
}

func (BlockPictureWithCaption) NodeKind() string {
    return "block-picture-with-caption"
}

func (n *BlockPictureWithCaption) Children() []components.DocumentBlockNode {
    return []components.DocumentBlockNode{
        NewBlockPicture(n.image),
        NewBlockInline(NewInlineTextNode(NewInlineText(n.caption))),
    }
}

func (BlockPictureWithCaption) ImplBlockNode() {}
func (BlockPictureWithCaption) ImplDocumentNode() {}
