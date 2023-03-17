package store

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInfo_GetInfo(t *testing.T) {
	ctx := context.Background()

	info := NewInfo(ctx, time.Second)

	id1, _, date1 := info.GetInfo()
	id2, _, date2 := info.GetInfo()
	assert.Equal(t, id1, id2)
	assert.Equal(t, date1, date2)

	<-time.After(time.Second)

	id2, _, date2 = info.GetInfo()

	assert.NotEqual(t, id1, id2)
	assert.NotEqual(t, date1, date2)
}
