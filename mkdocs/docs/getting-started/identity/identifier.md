# Identifier

Each identity has one main identifier - **Genesis ID**, and can have many additional identifiers - [Identity Profiles](./identity-profile.md).

Genesis ID is:

- Permanent: it remains the same for the entire existence of an identity.
- Unique: No two identities can have the same ID.

The Genesis ID is deterministically [calculated](https://docs.iden3.io/protocol/spec/#genesis-id) from the Genesis State.

**Calculate Genesis ID from the Genesis State**

```go
package main

import (
"fmt"

core "github.com/iden3/go-iden3-core"
)

// Retrieve ID
func main() {

    id, _ := core.IdGenesisFromIdenState(core.TypeDefault, state.BigInt())

    fmt.Println("ID:", id)

}
```

Here is what the output would look like: 

```bash
ID: 11AbuG9EKnWVXK1tooT2NyStQod2EnLhfccSajkwJA
```

The identity gets published, together with all other identities, inside the struct `StateLib.Data internal _stateData` state variable, which is part of the `State` [contract](../../contracts/state.md). While the ID remains constant, the Identity State will get updated as soon as the identity adds or revokes claims in its trees. 

> No Personal Identifiable Information (PPI) is stored on-chain. From the IdState is impossible to retrieve any information (represented as claim) stored inside the Identity Claims Tree

The Identity State hasn't been published on-chain yet as claims haven't been issued yet. This is the subject of the next section.

> The executable code can be found [here](https://github.com/0xPolygonID/tutorial-examples/blob/main/issuer-protocol/main.go#L133)

## Genesis State

There are two [identity types](./identity-types.md), which differ in many aspects and because of this have two different ways to generate the Genesis State:

1. **Regular Identity**: The Genesis State is the initial Identity State (hash of Identity SMT Roots). This identity is primarily controlled by Baby JubJub keys. At least one BJJ public key must be added into Claims Tree during the identity creation.
2. **Ethereum-controlled Identity**: The Genesis State is derived from the Ethereum address. This identity is primarily controlled by Ethereum account from which its Genesis State and Identifier are derived.

## W3C DID representation

Decentralized Identifier can be generated from the ID by prepending the DID method and network parameters in the following way:

```
did:<method>:<network>:<subnet>:<id>
```

Example of valid DIDs:

1.
    ```
    did:iden3:eth:mainnet:11AbuG9EKnWVXK1tooT2NyStQod2EnLhfccSajkwJA
    ```
    where:
    * `did:iden3` is DID method  
    * `eth:mainnet` is the network identifier for the Ethereum Mainnet
    * `id`, base58-encoded id.  

2.
    ```
    did:polygonid:polygon:mumbai:2qCU58EJgrEMAMwdTehMoxtopwP1gKXCEt9GGeVDaG
    ```
    where:
    * `did:polygonid` is DID method  
    * `polygon:mumbai` is the network identifier for the Mumbai testnet  
    * `id`, base58-encoded id.  
