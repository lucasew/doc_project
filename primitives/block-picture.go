package primitives

import (
	"image"
	"net/url"
)

type BlockPicture struct {
    image Image
}

type Image interface {
    ToDataURL() url.URL
    ToRaw() RawImage
    ToImage() image.Image
}

type RawImage interface {
    ImageExtension() string
    ImageBytes() []byte
}

func NewBlockPicture(image Image) DocumentBlockNode {
    return &BlockPicture{image: image}
}

func (BlockPicture) Children() []DocumentBlockNode {
    return NewEmptyDocumentBlockNodeList()
}

func (BlockPicture) ImplBlockNode() {}
func (BlockPicture) ImplDocumentNode() {}

