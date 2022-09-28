# Identity State

An Identity State `IdS` is represented by the hash of the roots of these three merkle trees. 

`IdS = H(ClR || ReR || RoR)` where:

- `H`: Hashing Function
- `ClR`: Claims Tree Root
- `ReR`: Revocation Tree Root
- `RoR`: Roots Tree Root

The identity state gets stored on-chain and represents the status of an identity at a certain point in time.

<div align="center">
<img src= "../../../imgs/identity-state-diagram.png" align="center" width="600"/>
<div align="center"><span style="font-size: 17px;"><b> Identity State Diagram </b></div>
</div>

**Retrieve the Identity State `IdS`**

```go
package main

import (
    "fmt"
    "github.com/iden3/go-merkletree-sql"
)

// Retrieve Identity State
func main() {

    // calculate Identity State as a hash of the three roots
    state, _ := merkletree.HashElems(
        clt.Root().BigInt(),
        ret.Root().BigInt(),
        rot.Root().BigInt())

    fmt.Println("Identity State:", state)

}
```

Here is what the output would look like: 

```bash
Identity State: 
20698226269617404048572275736120991936409000313072409404791246779211976957795
```

> The very first identity state of an identity is defined as **Genesis State**

Every verification inside Iden3 protocol is executed against the Identity State. For instance, to prove the validity of a specific claim issued by A to B (in case if the claims gets added to the claims tree):

- user B needs to produce a merkle proof of the existence of that claim inside user's A `ClR`
- user B needs to produce a merkle proof of non existence of the corresponding revocation nonce inside user's A `ReT`

> The executable code can be found [here](https://github.com/0xPolygonID/tutorial-examples/blob/main/issuer-protocol/main.go#L124)
