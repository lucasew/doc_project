package primitives

type BlockPictureWithCaption struct {
    image Image
    caption string
}

func NewBlockPictureWithCaption(image Image, caption string) DocumentBlockNode {
    return &BlockPictureWithCaption{
        image: image,
        caption: caption,
    }
}

func (n *BlockPictureWithCaption) Children() []DocumentBlockNode {
    return []DocumentBlockNode{
        NewBlockPicture(n.image),
        NewBlockInline(NewInlineTextNode(NewInlineText(n.caption))),
    }
}

func (BlockPictureWithCaption) ImplBlockNode() {}
func (BlockPictureWithCaption) ImplDocumentNode() {}
