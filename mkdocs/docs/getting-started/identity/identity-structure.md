# Identity Structure

Each identity consists of [Three Sparse Merkle Trees](https://docs.iden3.io/protocol/spec/#identity-state-update):

- `ClT`: A Claims tree that contains the claims issued by that particular identity
- `ReT`: A Revocation tree that contains the revocation nonces of the claims that have been revoked by that particular identity
- `RoT`: A Roots tree that contains the history of the tree roots from the Claims tree

Claims issued by an identity are added to the Claims tree (we'll see in a while why that's not always the case). The position of a claim inside the Sparse Merkle Tree is determined by the hash of the claim's `Index` while the value stored inside the leaf will be the hash of the claim's `Value`.

An identity must issue at least one `Auth Claim` to operate properly. This is the first claim that is issued by an identity and that **must** be added to the `ClT`.

**Create identity trees and add authClaim**

```go
package main

import (
    "github.com/iden3/go-merkletree-sql"
)

// Generate the three identity trees
func main() {

    // Create empty Claims tree
    clt, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 32) 

    // Create empty Revocation tree
    ret, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 32) 

    // Create empty Roots tree
    rot, _ := merkletree.NewMerkleTree(ctx, memory.NewMemoryStorage(), 32) 

    // Get the Index and the Value of the authClaim
    hIndex, hValue, _ := authClaim.HiHv()

    // add auth claim to claims tree with value hValue at index hIndex
    clt.Add(ctx, hIndex, hValue)

    // print the roots
    fmt.Println(clt.Root().BigInt(), ret.Root().BigInt(), rot.Root().BigInt())
}
```

We just generated the three identity trees! For now, we only added a leaf correponding to the `authClaim` to the Claims tree `ClT`. The Revocation tree `ReT` and the `RoT` remain empty. In particular:

- The revocation tree gets updated whenever an identity decides to revoke a claim. For instance, if a user decides to rotate her keys, then she generates a key pair, creates a new authClaim with the public key from the key pair and adds the claim to the Claims Tree. Now the user can revoke the old public key, so she adds an entry to the Revocation Tree with the claim revocation nonce as an Index and zero as a Value. 
- The Roots Tree gets updated whenever the Identity Claims Tree root gets updated.

> The executable code can be found [here](https://github.com/0xPolygonID/tutorial-examples/blob/main/issuer-protocol/main.go#L104)
