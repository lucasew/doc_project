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
    // ImplInlineNode: just a signal when typechecking
    Children() []DocumentInlineNode
    ImplInlineNode()
}

type DocumentTextInlineNode interface {
    DocumentNode
    ImplInlineNode()
    ExtractText() string
    Children() []DocumentTextInlineNode
}

type DocumentBlockNode interface {
    DocumentNode
    // ImplBlockNode: just a signal when typechecking
    ImplBlockNode()
    Children() []DocumentNode
}

