package operations_test

import (
	"testing"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/internal/testutil"
	. "github.com/omise/omise-go/operations"
	a "github.com/stretchr/testify/assert"
)

// TODO: No way to programmatically generates Dispute against the API yet.
//   so not sure how we can test this thoroughly.
func TestDispute(t *testing.T) {
	client, e := testutil.NewClient()
	if !a.NoError(t, e) {
		return
	}

	// only test JSON bindings for now.
	disputes, list := &omise.DisputeList{}, &ListDisputes{}
	if e := client.Do(disputes, list); !a.NoError(t, e) {
		return
	}

	a.Len(t, disputes.Data, 0)
}
