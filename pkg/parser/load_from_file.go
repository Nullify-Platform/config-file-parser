package parser

import (
	"os"

	"github.com/nullify-platform/config-file-parser/pkg/models"
)

func LoadFromFile(path string) (*models.Configuration, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ParseConfiguration(data)
}
