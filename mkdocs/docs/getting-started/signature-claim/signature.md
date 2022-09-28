# Signature Claim Issuance

To issue a claim by signing it, the only thing needed is access to your [Baby Jubjub private key](../babyjubjub.md).

1. **Retrieve hash of Claim's Index and hash of Claim's Value**

    Starting from the [Generic Claim](../claim/generic-claim.md) previously created the first step we first need to extract the hash of its index and the hash of its value

    ```go
    claimIndex, claimValue := claim.RawSlots()
	indexHash, _ := poseidon.Hash(core.ElemBytesToInts(claimIndex[:]))
	valueHash, _ := poseidon.Hash(core.ElemBytesToInts(claimValue[:]))
    ```

2. **Hash the `indexHash` and the `valueHash` together and sign it**

    ```go
	// Poseidon Hash the indexHash and the valueHash together to get the claimHash
	claimHash, _ := merkletree.HashElems(indexHash, valueHash)

	// Sign the claimHash with the private key of the issuer
	claimSignature := babyJubjubPrivKey.SignPoseidon(claimHash.BigInt())
    ```

> The executable code can be found [here](https://github.com/0xPolygonID/tutorial-examples/blob/main/issuer-protocol/main.go#L139)


