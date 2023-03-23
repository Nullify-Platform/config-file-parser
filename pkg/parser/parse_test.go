package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const configStr string = `
minimum_comment_severity: medium
ignore_dirs: ["data"]
`

func TestParseConfiguration(t *testing.T) {
	config, err := ParseConfiguration([]byte(configStr))
	require.NoError(t, err)
	assert.Equal(t, "medium", config.MinimumCommentSeverity)
	assert.Equal(t, 1, len(config.IgnoreDirs))
	assert.Equal(t, "data", config.IgnoreDirs[0])
}
