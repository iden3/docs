# Identity Profiles

Identity Profiles allow users to hide their [`Genesis Identifier`](identifier.md) during interactions. Instead, users will be identified by their [`Identity Profile`](../../protocol/spec.md#identity-profiles-new).

```go
package main

import (
	"fmt"
	"math/big"

	core "github.com/iden3/go-iden3-core"
)

// Generate Identity Profile from Genesis Identifier
func main() {

    id, _ := core.IDFromString("11BBCPZ6Zq9HX1JhHrHT3QKUFD9kFDEyJFoAVMptVs")

    profile, _ := core.ProfileID(id, big.NewInt(50))

    fmt.Println(profile.String())

}
```