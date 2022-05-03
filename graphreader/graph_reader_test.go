package graphreader

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGraphReader(t *testing.T) {
	graph, err := Read("../graph.yml")
	require.NoError(t, err)
	require.NotNil(t, graph)
}
