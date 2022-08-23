package sqrt

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	value    float64
	expected float64
}

func TestMany(t *testing.T) {
	f, err := os.Open("testcases.csv")
	require.NoError(t, err)
	defer f.Close()
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	require.NoError(t, err)

	for _, v := range records {
		fIn, err := strconv.ParseFloat(strings.TrimSpace(v[0]), 64)
		require.NoError(t, err)
		fExp, err := strconv.ParseFloat(strings.TrimSpace(v[1]), 64)
		require.NoError(t, err)
		tst := testCase{value: fIn, expected: fExp}

		t.Run(fmt.Sprintf("%f", tst.value), func(t *testing.T) {
			out, err := Sqrt(tst.value)
			require.NoError(t, err)
			require.InDelta(t, tst.expected, out, 0.001)
		})
	}
}
