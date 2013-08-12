package workspace

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mibitzi/stwm/entities/client"
	"github.com/mibitzi/stwm/test"
)

func TestAddClient(t *testing.T) {
	ws := New("ws")
	client, _ := client.New(test.NewWindow())
	assert.NoError(t, ws.AddClient(client))
	assert.True(t, ws.HasClient(client.Id()))
}
