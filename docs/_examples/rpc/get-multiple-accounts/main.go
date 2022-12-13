package main

import (
	"context"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/doesnotcommit/solana-go-sdk/client"
	"github.com/doesnotcommit/solana-go-sdk/rpc"
)

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)
	accountInfos, err := c.GetMultipleAccounts(
		context.Background(),
		[]string{
			"AngPj3DQh1zNW68E1sa2yWfy71D6mcd7iwzhbLiLMxcR",
			"4T2BWArMHpFiwAMkMsJUCxa5ux9vYZBnVigJznxzbeVx",
		},
	)
	if err != nil {
		log.Fatalf("failed to get multiple account infos, err: %v", err)
	}

	spew.Dump(accountInfos)
}
