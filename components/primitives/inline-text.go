package primitives

import "github.com/lucasew/doc_project/components"


type InlineText struct {
    text string
}

func NewInlineText(text string) components.DocumentTextInlineNode {
    return &InlineText{
        text: text,
    }
}

func (InlineText) NodeKind() string {
    return "inline-text"
}

func (InlineText) ImplInlineNode() {}
func (InlineText) ImplDocumentNode() {}
func (t *InlineText) ExtractText() string {
    return t.text
}

func (InlineText) Children() []components.DocumentTextInlineNode {
    return []components.DocumentTextInlineNode{}
}
