# Revocation

Revocation is the process of invalidating a claim. For example, if a claim is used to prove that a person is resident in a country, and the person moves to another country, the claim can be revoked by the issuer.

This is done by adding the claim revocation nonce to the [revocation tree](https://docs.iden3.io/protocol/spec/#revocation-tree). The revocation tree contains revocation nonces of all the claims that were revoked. The root of the revocation tree is stored in the [identity state](./identity/identity-state.md).

To revoke a claim, the revocation nonce of the claim must be added to the revocation tree. The revocation nonce is a number that is added to the claim [data structure](./claim/generic-claim.md) when it is created. 

```go
	// revocation nonce of the claim to be revoked
	revocationNonce := uint64(1909830690)

	// add the revocation nonce to the revocation tree
	ret.Add(ctx, new(big.Int).SetUint64(revNonce), big.NewInt(0))
```

The action of adding the revocation nonce to the revocation tree modifies the root of the revocation tree and, consequently, the identity state. 

To finalize the revocation process, the identity state must be updated on-chain by executing a [state transition](./state-transition/state-transition.md).
