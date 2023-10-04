package tests

import (
	"bytes"
	"os"
	"testing"

	"github.com/nullify-platform/config-file-parser/pkg/parser"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestIntegration(t *testing.T) {
	config, err := parser.LoadFromFile("../data/nullify.yaml")
	require.NoError(t, err)

	var buf bytes.Buffer
	yamlEncoder := yaml.NewEncoder(&buf)
	yamlEncoder.SetIndent(2)
	err = yamlEncoder.Encode(config)
	require.NoError(t, err)

	originalData, err := os.ReadFile("../data/nullify.yaml")
	require.NoError(t, err)

	require.Equal(t, string(originalData), buf.String())
}
