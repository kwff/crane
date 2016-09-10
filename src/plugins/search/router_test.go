package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDocument Document = Document{
	ID:      "1",
	Name:    "test",
	Type:    "test",
	GroupId: uint64(123),
	Param:   map[string]string{"id": "test"},
}

func TestNewDocumentStorage(t *testing.T) {
	ds := NewDocumentStorage()
	assert.Empty(t, ds.Store, "NewDocumentStorage should be empty")
}

func TestEmpty(t *testing.T) {
	ds := NewDocumentStorage()
	ds.Store = map[string]Document{
		"document": testDocument,
	}
	ds.Empty()
	assert.Empty(t, ds.Store, "should be empty")
}

func TestCopy(t *testing.T) {
	ds := NewDocumentStorage()
	ds.Store = map[string]Document{
		"document": testDocument,
	}
	copidDS := ds.Copy()
	assert.Equal(t, copidDS.Store["document"], testDocument, "should be equal")
}

func TestIndices(t *testing.T) {
	ds := NewDocumentStorage()
	ds.Store = map[string]Document{
		"document": testDocument,
	}
	indicesDS := ds.Indices()
	assert.Equal(t, indicesDS, []string{"document"})
}

func TestSet(t *testing.T) {
	ds := NewDocumentStorage()
	ds.Set("test", testDocument)
	assert.Equal(t, ds.Store["test"], testDocument)
}

func TestGet(t *testing.T) {
	ds := NewDocumentStorage()
	ds.Store = map[string]Document{
		"document": testDocument,
	}
	dc := ds.Get("document")
	assert.Equal(t, dc, &testDocument)
}