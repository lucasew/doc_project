package providers

import (
	"image"
	"net/url"
)

type Image interface {
    ToDataURL() url.URL
    ToRaw() RawImage
    ToImage() image.Image
}

type RawImage interface {
    ImageExtension() string
    ImageBytes() []byte
}
