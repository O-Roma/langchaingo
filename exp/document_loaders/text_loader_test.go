package document_loaders

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tmc/langchaingo/exp/text_splitters"
)

func TestTextLoader_Load(t *testing.T) {
	t.Parallel()

	loader := NewTextLoaderFromFile("./testdata/test.txt")

	docs, err := loader.Load()
	require.NoError(t, err)
	require.Len(t, docs, 1)

	expectedPageContent := "Foo Bar Baz"
	assert.Equal(t, docs[0].PageContent, expectedPageContent)

	expectedMetadata := map[string]any{
		"source": "./testdata/test.txt",
	}

	assert.Equal(t, docs[0].Metadata, expectedMetadata)
}

func TestTextLoader_LoadAndSplit(t *testing.T) {
	t.Parallel()

	loader := NewTextLoaderFromFile("./testdata/test.txt")
	splitter := text_splitters.NewRecursiveCharactersSplitter()
	splitter.ChunkOverlap = 1
	splitter.ChunkSize = 3

	docs, err := loader.LoadAndSplit(splitter)
	require.NoError(t, err)
	require.Len(t, docs, 3)

	expectedMetadata := map[string]interface{}{
		"source": "./testdata/test.txt",
	}

	assert.Equal(t, docs[0].PageContent, "Foo")
	assert.Equal(t, docs[0].Metadata, expectedMetadata)

	assert.Equal(t, docs[1].PageContent, "Bar")
	assert.Equal(t, docs[1].Metadata, expectedMetadata)

	assert.Equal(t, docs[2].PageContent, "Baz")
	assert.Equal(t, docs[2].Metadata, expectedMetadata)
}
