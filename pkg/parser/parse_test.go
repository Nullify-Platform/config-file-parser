package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseConfiguration(t *testing.T) {
	configStr := "ignore_dirs: [\"data\"]"
	config, err := ParseConfiguration([]byte(configStr))
	require.NoError(t, err)
	assert.Equal(t, 1, len(config.IgnoreDirs))
	assert.Equal(t, "data", config.IgnoreDirs[0])
}
