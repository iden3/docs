# Generic Claim

A [Claim](https://docs.iden3.io/protocol/spec/#claims) is a statement made by one identity about another identity or about itself. In general, claim is a flexible and modular data primitive that can be used to represent any identity-related information.

Claims can be viewed as Soul Bound Tokens (SBTs) on steroids. Similar to SBTs, the ownership is cryptographically guaranteed allowing control and reusability across platforms. Differently to SBTs, claims live off-chain ensuring users privacy over their Personal Identifiable Information.

1. **Update the required dependencies.**

	```bash
	go get github.com/iden3/go-iden3-core
	```

2. **Define the claim schema**.

	A [claim schema](./claim-schema.md) defines how a set of data must be stored inside a claim. In this example, we will use a schema called [`KYCAgeCredential`](https://github.com/iden3/claim-schema-vocab/blob/main/schemas/json-ld/kyc-v2.json-ld). According to this schema the birthday is stored in the first index slot of the [claim data structure](https://docs.iden3.io/protocol/claims-structure), while the documentType is stored in the second data slot.

    The hash of the schema is generated from the content of the schema document following the [Claim Schema Generation Rules](./claim-schema.md). For our example, the hash of the schema is: *`2e2d1c11ad3e500de68d7ce16a0a559e`*

3. **Create a generic claim.**  

	```go
	package main

	import (
		"encoding/json"
		"fmt"
		"math/big"
		"time"

		core "github.com/iden3/go-iden3-core"
	)

	// create basic claim
	func main() {

		// set claim expriation date to 2361-03-22T00:44:48+05:30
		t := time.Date(2361, 3, 22, 0, 44, 48, 0, time.UTC)
		
		// set schema
		ageSchema, _ := core.NewSchemaHashFromHex ("2e2d1c11ad3e500de68d7ce16a0a559e")  

		// define data slots
		birthday := big.NewInt(19960424)
		documentType := big.NewInt(1)	
		
		// set revocation nonce 
		revocationNonce := uint64(1909830690)

		// set ID of the claim subject
		id, _ := core.IDFromString("113TCVw5KMeMp99Qdvub9Mssfz7krL9jWNvbdB7Fd2")

		// create claim 
		claim, _ := core.NewClaim(ageSchema, core.WithExpirationDate(t), core.WithRevocationNonce(revocationNonce), core.WithIndexID(id), core.WithIndexDataInts(birthday, documentType))

		// transform claim from bytes array to json 
		claimToMarshal, _ := json.Marshal(claim)

		fmt.Println(string(claimToMarshal))
	}
	```

Here is what the claim would look like:
```
["3613283249068442770038516118105710406958","86645363564555144061174553487309804257148595648980197130928167920533372928","19960424","1","227737944108667786680629310498","0","0","0"]
```

In particular, the first 4 values of the claim represent the `Index` part of the claim while the last 4 represent the `Value`.
```
Index:
{
"3613283249068442770038516118105710406958", // Claim Schema hash
"86645363564555144061174553487309804257148595648980197130928167920533372928", // ID Subject of the claim
"19960424", // First index data slot stores the date of birth
"1"  //  Second index data slot stores the document type
}

Value:
{ 
"227737944108667786680629310498", // Revocation nonce 
"0",
"0", // first value data slot
"0"  // second value data slot
}	
```

The data stored in the first position of the Index contains a reference to the schemahash of the claim. As defined in the [`KYCAgeCredential` schema](https://github.com/iden3/claim-schema-vocab/blob/main/schemas/json-ld/kyc-v2.json-ld), the value birthday must be stored in the first index data slot while the second index stores the documentType. Other schemas may provide different rules on where to store the data.

> The executable code can be found [here](https://github.com/0xPolygonID/tutorial-examples/blob/main/issuer-protocol/main.go#L63)
