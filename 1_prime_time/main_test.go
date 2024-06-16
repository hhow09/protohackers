package main

import (
	"bufio"
	"encoding/json"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	num     float64
	isPrime bool
}

func TestHandlePrimeTime(t *testing.T) {
	tests := []testCase{
		{
			num:     1,
			isPrime: false,
		},
		{
			num:     123,
			isPrime: false,
		},
		{
			num:     7,
			isPrime: true,
		},
	}
	server, client := net.Pipe()
	client.SetDeadline(time.Now().Add(2 * time.Second))
	defer client.Close()
	go func() {
		handlePrimeTime(server)
	}()
	for _, tc := range tests {
		req := request{
			Method: "isPrime",
			Number: &tc.num,
		}
		b, err := json.Marshal(req)
		require.NoError(t, err)
		n, err := client.Write(append(b, '\n'))
		require.Equal(t, len(b)+1, n)
		require.NoError(t, err)

		bfio := bufio.NewReader(client)
		res, err := bfio.ReadBytes('\n')
		require.NoError(t, err)
		resp := &response{}
		err = json.Unmarshal(res, resp)
		require.NoError(t, err)
		require.Equal(t, tc.isPrime, resp.Prime)
		require.Equal(t, "isPrime", resp.Method)
	}
}
