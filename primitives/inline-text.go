package primitives


type InlineText struct {
    text string
}

func NewInlineText(text string) DocumentTextInlineNode {
    return &InlineText{
        text: text,
    }
}

func (InlineText) ImplInlineNode() {}
func (InlineText) ImplDocumentNode() {}
func (t *InlineText) ExtractText() string {
    return t.text
}

func (InlineText) Children() []DocumentTextInlineNode {
    return []DocumentTextInlineNode{}
}
