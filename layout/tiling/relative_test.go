package tiling

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelativeUp(t *testing.T) {
	tiling, clients, _ := prepareTestClients(3)
	client, err := tiling.RelativeUp(clients[2])
	assert.NoError(t, err)
	assert.Equal(t, clients[1], client)
}

func TestRelativeDown(t *testing.T) {
	tiling, clients, _ := prepareTestClients(3)
	client, err := tiling.RelativeDown(clients[1])
	assert.NoError(t, err)
	assert.Equal(t, clients[2], client)
}
