# Baby Jubjub Key Pair

In Iden3 Protocol the public and private key pair is used to manage an identity and to authenticate in the name of an identity. In particular, Baby Jubjub is the elliptic curve used in Iden3. [This curve](https://github.com/iden3/iden3-docs/blob/master/source/docs/Baby-Jubjub.pdf) is designed to work efficiently with zkSNARKs.

1. **Initiate a Go Module**

    ```bash
    go mod init example/iden3-tutorial
    ```

2. **Update the required dependencies.**

    ```bash
    go get github.com/iden3/go-iden3-crypto/babyjub
    ```
    
3. **Generate a baby jubjub public key.**
    ``` go
    package main

    import (
        "fmt"
        "github.com/iden3/go-iden3-crypto/babyjub"
    )

    // BabyJubJub key
    func main() {

        // generate babyJubjub private key randomly
        babyJubjubPrivKey := babyjub.NewRandPrivKey()

        // generate public key from private key
        babyJubjubPubKey := babyJubjubPrivKey.Public()

        // print public key
	    fmt.Println(babyJubjubPubKey)
    }
    ```

Here is an example of a public key generated using Baby Jubjub:

```bash
500d43e1c3daa864995a9615b6f9e3a4fd0af018548c583773b6e422b14201a3
```

> The executable code can be found [here](https://github.com/0xPolygonID/tutorial-examples/blob/main/issuer-protocol/main.go#L21)
