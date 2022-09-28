# Identifier

Each identity has a unique identifier. `ID` is:

- Permanent: it remains the same for the entire existence of an identity.
- Unique: No two identities can have the same ID.

The `ID` is deterministically [calculated](https://docs.iden3.io/protocol/spec/#genesis-id) from the Genesis State. 

**Retrieve the Identifier `ID`**

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
ID:
11AbuG9EKnWVXK1tooT2NyStQod2EnLhfccSajkwJA
```

Hereafter, this identity is represented as a mapping: `ID => IdS`. This gets published, together with all other identities, inside the `identities` mapping, which is part of the `State.sol` [contract](../../contracts/overview.md). While the ID remains constant, the Identity State will get updated as soon as the identity adds or revokes claims in its trees. 

> No Personal Identifiable Information (PPI) is stored on-chain. From the IdS is impossible to retrieve any information (represented as claim) stored inside the Identity Claims Tree

The Identity State hasn't been published on-chain yet as claims haven't been issued yet. This is the subject of the next section.

> The executable code can be found [here](https://github.com/0xPolygonID/tutorial-examples/blob/main/issuer-protocol/main.go#L133)
