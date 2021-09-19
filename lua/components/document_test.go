package lua_components

import (
	// "testing"

	// "github.com/lucasew/doc_project/components"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/lucasew/doc_project/components"
	"github.com/lucasew/doc_project/lua"
	lua "github.com/yuin/gopher-lua"
)

type DocumentMock struct {}

func (DocumentMock) CreatedAt() time.Time {
    return time.Now()
}

func (DocumentMock) DocumentTitle() string {
    return "test"
}

func (DocumentMock) Nodes() []components.DocumentBlockNode {
    return components.NewEmptyDocumentBlockNodeList()
}

func (DocumentMock) RawMetadata() map[string]interface{} {
    return map[string]interface{}{}
}

func NewDocumentMock() components.Document {
    return DocumentMock{}
}


func TestWrapDocument(t *testing.T) {
    L := app_lua.NewCommonState(lua.Options{})
    mock := NewDocumentMock()
    wrapped := WrapDocument(L, mock)
    spew.Dump(wrapped)
    L.SetGlobal("demo", wrapped)
    app_lua.RunREPL(L)
    err := L.DoString(`
    return demo:title()
    `)
    expected := mock.DocumentTitle()
    got := L.CheckString(1)
    if expected != got {
        t.Errorf("expected %s got %s", expected, got)
    }
    if err != nil {
        t.Error(err)
    }
}
