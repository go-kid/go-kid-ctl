package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryInRegistry(t *testing.T) {
	inRegistry, err := QueryInRegistry("scan-component")
	assert.NoError(t, err)
	fmt.Println(inRegistry)
}
