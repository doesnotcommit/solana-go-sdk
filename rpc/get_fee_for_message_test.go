package rpc

import (
	"context"
	"testing"

	"github.com/doesnotcommit/solana-go-sdk/pkg/pointer"
)

func TestGetFeeForMessage(t *testing.T) {
	tests := []testRpcCallParam{
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getFeeForMessage", "params":["AQABAyRn8Htq2L5KAQiNyByMm5M/q8rDpBu7qahSf2bBSZq4Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMqYD+EHyvdHM3SIIuGet5Q6BxOI26dTbdOzaCY8V1mtAQICAAEMAgAAAAEAAAAAAAAA"]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":112884237},"value":null},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetFeeForMessage(
					context.TODO(),
					"AQABAyRn8Htq2L5KAQiNyByMm5M/q8rDpBu7qahSf2bBSZq4Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMqYD+EHyvdHM3SIIuGet5Q6BxOI26dTbdOzaCY8V1mtAQICAAEMAgAAAAEAAAAAAAAA",
				)
			},
			ExpectedResponse: JsonRpcResponse[GetFeeForMessage]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetFeeForMessage{
					Context: Context{
						Slot: 112884237,
					},
					Value: nil,
				},
			},
			ExpectedError: nil,
		},
		{
			RequestBody:  `{"jsonrpc":"2.0", "id":1, "method":"getFeeForMessage", "params":["AQABAyRn8Htq2L5KAQiNyByMm5M/q8rDpBu7qahSf2bBSZq4Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMqYD+EHyvdHM3SIIuGet5Q6BxOI26dTbdOzaCY8V1mtAQICAAEMAgAAAAEAAAAAAAAA", {"commitment":"processed"}]}`,
			ResponseBody: `{"jsonrpc":"2.0","result":{"context":{"slot":112884286},"value":5000},"id":1}`,
			RpcCall: func(rc RpcClient) (any, error) {
				return rc.GetFeeForMessageWithConfig(
					context.TODO(),
					"AQABAyRn8Htq2L5KAQiNyByMm5M/q8rDpBu7qahSf2bBSZq4Bj5w2ZFXmNyj7tuRN89kxw/6+2LN04KBBSUL12sdbN4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMqYD+EHyvdHM3SIIuGet5Q6BxOI26dTbdOzaCY8V1mtAQICAAEMAgAAAAEAAAAAAAAA",
					GetFeeForMessageConfig{
						Commitment: CommitmentProcessed,
					},
				)
			},
			ExpectedResponse: JsonRpcResponse[GetFeeForMessage]{
				JsonRpc: "2.0",
				Id:      1,
				Error:   nil,
				Result: GetFeeForMessage{
					Context: Context{
						Slot: 112884286,
					},
					Value: pointer.Get[uint64](5000),
				},
			},
			ExpectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			testRpcCall(t, tt)
		})
	}
}
