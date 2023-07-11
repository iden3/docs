package main

import (
	"context"
	"fmt"
	"math/big"

	core "github.com/iden3/go-iden3-core"
	"github.com/iden3/go-merkletree-sql/v2"
	"github.com/iden3/go-merkletree-sql/v2/db/memory"
)

// Generate the three identity trees
func main() {

	ctx := context.Background()

	// Create empty Claims tree
	clt, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 40)

	// Create empty Revocation tree
	ret, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 40)

	// Create empty Roots tree
	rot, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 40)

	// Generate a proper BJJ key pair here and provide x & y coordinates
	x := big.NewInt(1)
	y := big.NewInt(2)

	authClaim, err := core.NewClaim(core.AuthSchemaHash,
		core.WithIndexDataInts(x, y),
		core.WithRevocationNonce(0))

	if err != nil {
		panic(err)
	}

	// Get the Index and the Value of the authClaim
	hIndex, hValue, _ := authClaim.HiHv()

	// add auth claim to claims tree with value hValue at index hIndex
	clt.Add(ctx, hIndex, hValue)

	// print the roots
	fmt.Println(clt.Root().BigInt(), ret.Root().BigInt(), rot.Root().BigInt())
}
