package contextWithValue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ReadWriteContext(t *testing.T) {
	ctx := context.Background()
	userName := "joy"
	ctx = contextWithUser(ctx, userName)
	getUserName, ok := userFromContext(ctx)
	assert.True(t, ok)
	assert.Equal(t, userName, getUserName)
}
