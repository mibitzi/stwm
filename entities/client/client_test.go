package client

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mibitzi/stwm/test"
)

func TestNew(t *testing.T) {
	win := test.NewWindow()
	client, err := New(win)
	assert.NoError(t, err, "new: New client")
	assert.NotNil(t, client, "new: New client")
	assert.Equal(t, client.Id(), win.Id(), "new: Client id")
}
