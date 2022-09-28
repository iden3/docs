# Key Authorization Claim

The most important building block of an identity is the Key Authorization Claim. This claim stores user's Baby Jubjub public key. 

An [Auth Claim](https://docs.iden3.io/protocol/bjjkey/) **must** be included as a leaf inside the  Identity Tree. All the actions performed by an Idenitity (such as claim issuance or revocation) require users to prove via a digital signature that they own the private key associated with the public key stored in the `AuthClaim`.

1. **Define the claim schema.**

    The [auth schema](https://github.com/iden3/claim-schema-vocab/blob/main/schemas/json-ld/auth.json-ld) is pre-defined and should always be the same when creating an `AuthClaim`. The schema hash is: *`ca938857241db9451ea329256b9c06e5`*. According to the this schema, the X and Y coordinate of the Baby Jubjub public key must be stored, respectively, in the first and second index data slot.

2. **Generate an AuthClaim.** 

    ```go
    package main

    import (
        "encoding/json"
        "fmt"

        "github.com/iden3/go-iden3-core"
        "github.com/iden3/go-iden3-crypto/babyjub"
    )

    // Create auth claim
    func main() {

        authSchemaHash, _ := core.NewSchemaHashFromHex("ca938857241db9451ea329256b9c06e5")

        // Add revocation nonce. Used to invalidate the claim. This may be a random number in the real implementation.
        revNonce := uint64(1)

        // Create auth Claim 
        authClaim, _ := core.NewClaim(authSchemaHash,
        core.WithIndexDataInts(babyJubjubPubKey.X, babyJubjubPubKey.Y),
        core.WithRevocationNonce(revNonce))

        authClaimToMarshal, _ := json.Marshal(authClaim)

        fmt.Println(string(authClaimToMarshal))
    }
    ```

Here is what the claim would look like: 

```
Claim:
["304427537360709784173770334266246861770","0","12360031355466667401641753955380306964012305931931806442343193949747916655340","7208907202894542671711125895887320665787554014901011121180092863817137691080","1","0","0","0"]
```

Let us destructure the output:

```
Index:
{
"304427537360709784173770334266246861770", // Schema hash
"0",
"12360031355466667401641753955380306964012305931931806442343193949747916655340",  // X coordinate of the pubkey 	
"7208907202894542671711125895887320665787554014901011121180092863817137691080"   // Y coordinate of the pubkey
}

Value:
{ 
"1", // revocation nonce
"0",
"0", // first value data slot
"0"  // second value data slot
}	
```

The data stored in position 1 of the Value contains the Revocation Nonce. This value will be used to revoke/invalidate an `AuthClaim`. More on that in the next section.

> The executable code can be found [here](https://github.com/0xPolygonID/tutorial-examples/blob/main/issuer-protocol/main.go#L89)
