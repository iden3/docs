# Add Claim to the Claims Tree

At t=0, the situation is the same left from the *Identity* section of tutorial. Our identity is still at the `Genesis State`. The Claims Tree contains only the `authClaim`. The revocation and roots trees are empty. The state hasn't been published on-chain yet.

Let's see what happens when if we decide to add a new claim to the Claims Tree.

1. **Update the required dependencies.**

    ```bash
    go get github.com/iden3/go-circuits
    ```

2. **Add a new claim and fetch the new state**

    ```go
    package main

    import (
        "encoding/hex"
        "fmt"
        "math/big"

        "github.com/iden3/go-circuits"
        core "github.com/iden3/go-iden3-core"
        "github.com/iden3/go-iden3-crypto/poseidon"
        "github.com/iden3/go-merkletree-sql"
    )

    // Change Identity State
    func main() {
        // GENESIS STATE:

        // 1. Generate Merkle Tree Proof for authClaim at Genesis State
        authMTPProof, _, _ := clt.GenerateProof(ctx, hIndex, clt.Root())

        // 2. Generate the Non-Revocation Merkle tree proof for the authClaim at Genesis State
        authNonRevMTPProof, _, _ := ret.GenerateProof(ctx, new(big.Int).SetUint64(revNonce), ret.Root())

        // Snapshot of the Genesis State
        genesisTreeState := circuits.TreeState{
            State:          state,
            ClaimsRoot:     clt.Root(),
            RevocationRoot: ret.Root(),
            RootOfRoots:    rot.Root(),
        }
        // STATE 1:

        // Before updating the claims tree, add the claims tree root at Genesis state to the Roots tree.
	    rot.Add(ctx, clt.Root().BigInt(), big.NewInt(0))

        // Create a new random claim
        schemaHex := hex.EncodeToString([]byte("myAge_test_claim"))
        schema, _ := core.NewSchemaHashFromHex(schemaHex)

        code := big.NewInt(51)

        newClaim, _ := core.NewClaim(schema, core.WithIndexDataInts(code, nil))

        // Get hash Index and hash Value of the new claim
        hi, hv, _ := newClaim.HiHv()

        // Add claim to the Claims tree
        clt.Add(ctx, hi, hv)

        // Fetch the new Identity State
        newState, _ := merkletree.HashElems(
            clt.Root().BigInt(),
            ret.Root().BigInt(),
            rot.Root().BigInt())

        // Sign a message (hash of the genesis state + the new state) using your private key
        hashOldAndNewStates, _ := poseidon.Hash([]*big.Int{state.BigInt(), newState.BigInt()})

        signature := babyJubjubPrivKey.SignPoseidon(hashOldAndNewStates)

        // Generate state transition inputs
        stateTransitionInputs := circuits.StateTransitionInputs{
            ID:                id,
            OldTreeState:      genesisTreeState,
            NewState:          newState,
            IsOldStateGenesis: true,
            AuthClaim: circuits.Claim{
                Claim: authClaim,
                Proof: authMTPProof,
                NonRevProof: &circuits.ClaimNonRevStatus{
                    Proof: authNonRevMTPProof,
                },
            },
            Signature: signature,
        }

        // Perform marshalling of the state transition inputs
        inputBytes, _ := stateTransitionInputs.InputsMarshal()

        fmt.Println(string(inputBytes))

    }
    ```

After issuing a new claim, the claims tree gets modified and, therefore, the Identity State changes. To complete the state transition it is necessary to verify it inside a circuit. The type `StateTransitionInputs` lets us pack the inputs needed to generate a proof while the `InputsMarshal()` function turns it into a json file that can be used directly as State Transition Circuit inputs. These inputs will be used in the next section.

> The executable code can be found [here](https://github.com/0xPolygonID/tutorial-examples/blob/main/issuer-protocol/main.go#L156)
