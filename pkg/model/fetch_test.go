package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinks(t *testing.T) {
	assert.Equal(t, []string{}, ExtractLinks("https://fcp7.com/#/schema/"))
}
