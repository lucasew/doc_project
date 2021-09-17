package primitives

import "time"

type Document interface {
    CreatedAt() time.Time
    DocumentTitle() string
    RawMetadata() map[string]interface{}
    Nodes() []DocumentNode
}

type DocumentNode interface {
    ImplDocumentNode()
}

type DocumentInlineNode interface {
    DocumentNode
    Children() []DocumentInlineNode
    // ImplInlineNode: just a signal when typechecking
    ImplInlineNode()
}

type DocumentTextInlineNode interface {
    DocumentNode
    ExtractText() string
    Children() []DocumentTextInlineNode
    ImplInlineNode()
}

type DocumentBlockNode interface {
    DocumentNode
    Children() []DocumentBlockNode
    // ImplBlockNode: just a signal when typechecking
    ImplBlockNode()
}

