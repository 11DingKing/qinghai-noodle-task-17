package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask17(t *testing.T) {
	s := NewService(NewRegistry(), time.Now)
	require.NoError(t, s.CheckListingPrices(context.Background(), ProductListing{OnlinePriceCents: 1000, StorePriceCents: 1000}))
}
