# Sparse Merkle Tree

A Merkle Tree or a hash tree is a cryptographically verifiable data structure where every "leaf" node of the tree contains the cryptographic hash of a data block, and every non-leaf node contains the cryptographic hash of its child nodes.

The Merkle Trees used in Iden3 protocol are Sparse. In [Sparse Merkle Trees](https://blog.iden3.io/sparse-merkle-trees-visual-introduction.html) each data block has an index associated to it that determines its position as leaf inside the tree. 

In addition to inheriting the tamper-resistance and proof-of-membership properties from standard merkle trees, a Sparse Merkle Tree has other features:

- The insert order of data blocks doesn't influence the final Merkle Tree Root. A data block `A` with index `1` and a data block `B` with index `4` will always occupy the same positions inside the tree despite the insert order
- Some leaves remain empty
- It's possible to prove that certain data is not included in the tree (**proof of non-membership**)

A Sparse Merkle Tree is the core data structure used in Iden3 protocol to represent an identity. In particular, the leaves of a Sparse Merkle Tree are the claims issued by an identity. 

1. **Update the required dependencies.**

    ```bash
    go get github.com/iden3/go-merkletree-sql 
    ```

2. **Design a Sparse Merkle Tree.**


    ```go
    package main

    import (
        "context"
        "fmt"
        "math/big"

        merkletree "github.com/iden3/go-merkletree-sql"
        "github.com/iden3/go-merkletree-sql/db/memory"
    )

    // Sparse MT
    func main() {

        ctx := context.Background()

        // Tree storage
        store := memory.NewMemoryStorage()

        // Generate a new MerkleTree with 32 levels
        mt, _ := merkletree.NewMerkleTree(ctx, store, 32)

        // Add a leaf to the tree with index 1 and value 10
        index1 := big.NewInt(1)
        value1 := big.NewInt(10)
        mt.Add(ctx, index1, value1)

        // Add another leaf to the tree
        index2 := big.NewInt(2)
        value2 := big.NewInt(15)
        mt.Add(ctx, index2, value2)

        // Proof of membership of a leaf with index 1
        proofExist, value, _ := mt.GenerateProof(ctx, index1, mt.Root())

        fmt.Println("Proof of membership:", proofExist.Existence)
        fmt.Println("Value corresponding to the queried index:", value)

        // Proof of non-membership of a leaf with index 4
        proofNotExist, _, _ := mt.GenerateProof(ctx, big.NewInt(4), mt.Root())

        fmt.Println("Proof of membership:", proofNotExist.Existence)
    }
    ```

A data block inside the tree is represented by a `index` and a `value`. The index represents the position in the tree and it must be unique. The value represents the associated value stored in the tree.

The `GenerateProof` method shown above allows verifying the membership of a leaf in the merkle tree starting from its root. 

> The executable code can be found [here](https://github.com/0xPolygonID/tutorial-examples/blob/main/issuer-protocol/main.go#L32)

