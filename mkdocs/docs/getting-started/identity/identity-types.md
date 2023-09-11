# Identity Types

In iden3 protocol there are two types of identities, which differ in many aspects:

1. **Regular Identity**, which is generated from three identity trees and controlled by Baby JubJub keys.
2. **Ethereum-controlled Identity**, which is primarily controlled by Ethereum account from which it's Genesis State and
   Identifier are derived.

## Regular Identity

Regular identity is created from three merkle trees (Genesis State is a hash of Identity SMT Roots). This identity is
primarily controlled by Baby JubJub keys. At least one BJJ public key must be added into Claims Tree during the identity
creation.

```
genesisState = Hash(ClaimsTreeRoot || RevocationsTreeRoot || RootsTreeRoot)

genesisId = idType + genesisStateCut + checksum
```

where:

* idType - identifier of DID method and blockchain network & subnetwork, 2 bytes
* genesisStateCut - first 27 bytes of genesisState, 27 bytes
* checksum - control checksum, 2 bytes

### Limitations of Regular Identity

Currently, adding an Ethereum key to the Claims Tree and using it for authentication and proving is not practical, because
it's very computationally expensive to verify ECDSA signatures in zk-circuits. Also, having the user's Ethereum address
there's no way to get a user identifier from it, so dApps would need to authenticate the user additionally to get the
identifier.

## Ethereum-controlled Identity

This type of identity was introduced to overcome some of the limitations of regular identity - allow using Ethereum
accounts to authenticate, prove statements and control identity (perform state transitions). That eliminates strict
requirement to have Baby JubJub keys.

Genesis state is always zero for Ethereum-controlled Identity.

Genesis Identifier is directly derived from the Ethereum address in the following way:

```
genesisId = idType + zeroPadding + ethAddress + checksum
```

where:

* idType - identifier of DID method and blockchain network & subnetwork, 2 bytes
* zeroPadding - 7 zero bytes
* ethAddress - Ethereum address of controlling Ethereum account, 20 bytes
* checksum - control checksum, 2 bytes

Example:

```
idType: 0212 // DID method: PolygonID; Network: Polygon Mumbai
+
ethAddress: 0x0dcd1bf9a1b36ce34237eeafef220932846bcd82
+
// uint16 sum of bytes of the byte string: idType + zeroPadding + ethAddress.
// Note that the bytes of the uint16 are in reversed order, e.g. if sum is 0x0a45 then checksum is 0x450a
checksum: 450a
===
id: 0A4582CD6B84320922EFAFEE3742E36CB3A1F91BCD0D000000000000001202 (bytes, reversed order)
id: A5tDcNxacVgBQ4yHRvqv1FMR7cqNG74xGDhBWMidaq (base58)
```

Note, that smart contracts use little-endian byte order, so the resulting identifier is reversed.


### DID representation

Canonical form is the same as for Regular Identity:

```
did:polygonid:polygon:mumbai:2qCU58EJgrELSJT6EzT27Rw9DhvwamAdbMLpePztYq
```

### Authentication Method

Ethereum-controlled Identity can be authenticated by verifying if the Identifier matches the Ethereum account that sent a
transaction (`msg.sender`).

In case Ethereum-controlled Identity performs State Transition and adds BJJ keys to it's Claims Tree, it can also
perform authentication by ZKP using BJJ keys.

### Ethereum Account Types

Ethereum has two types of accounts:

* Smart Contract (SC)
* Externally Owned Account (EOA)

### Smart Contract (SC)

Smart contracts can control identity and perform state transitions. In case such a smart contract manages its identity
trees on chain, it becomes an OnChain Identity.

### Externally Owned Account (EOA)

Note: The EOA-controlled identity is not yet fully supported in iden3 protocol (WIP).

### Limitations of Ethereum-controlled Identity

* Only one Ethereum account can control identity
* No support for Ethereum key rotation and revocation. It's embedded into Identifier directly
* No support for profiles. It's not possible to hide genesis identifier and use profile instead when using Ethereum key
  to authenticate & prove statements.
* No support for credential issuance with BJJ signature, only SMT proofs available.

The last two limitations can be overcome by adding BJJ keys to Claims Tree and performing state transition. Afterwards, it's possible to use BJJ keys for authentication, proving & credential issuance. See [Identity Type Comparison](#identity-type-comparison) for more details.

## Identity Type Comparison

| Aspect                             | Regular Identity            | Ethereum-controlled Identity | Ethereum-controlled Identity with added BJJ keys           |
|------------------------------------|-----------------------------|------------------------------|------------------------------------------------------------|
| Genesis State                      | Identity State              | Ethereum Address             | Ethereum Address                                           |
| Keys                               | BJJ keys                    | Ethereum Account (SC or EOA) | Ethereum Account (SC or EOA) + BJJ keys                    |
| Authentication Method (off-chain)  | JWZ with ZKP using BJJ keys | JWS with Ethereum Signature  | JWS with Ethereum Signature or JWZ with ZKP using BJJ keys |
| Authentication Method (on-chain)   | ZKP using BJJ keys          | Ethereum Account             | Ethereum Account or ZKP using BJJ keys                     |
| State Transition Method            | ZKP using BJJ keys          | Ethereum Account             | Ethereum Account or ZKP using BJJ keys                     |
| Key Rotation Support               | Only BJJ keys               | Can add BJJ keys             | Only BJJ keys                                              |
| Profiles Support                   | Yes                         | No                           | Yes                                                        |
| Credential Issuance with MTP proof | Yes                         | Yes                          | Yes                                                        |
| Credential Issuance with Sig proof | Yes                         | No                           | Yes                                                        |
| Credential Revocation Support      | Yes                         | Yes                          | Yes                                                        |

