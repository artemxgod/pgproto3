package pgproto3_test

import (
	"testing"

	"github.com/jackc/pgproto3/v2"
	"github.com/stretchr/testify/require"
)

func TestDataRowDecodeNegativeValueLength(t *testing.T) {
	t.Parallel()

	body := []byte{0, 1, 0xFF, 0xFF, 0xFF, 0xFE}

	var row pgproto3.DataRow
	err := row.Decode(body)
	require.Error(t, err)
}
