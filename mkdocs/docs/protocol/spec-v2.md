# Protocol Specs V2

The new version of the Protocol contains new features such as Identity Profiles and Global Identities State Tree (GIST). These features are designed to add a further level of identity privacy: users are now able to hide their [Identifier](../getting-started/identity/identifier.md) when interacting with others. 

## Identity Profiles 

In the introduction, it was mentioned that these new features allow users to hide their Identifier during interactions. Technically, the global unique Identifier `ID` (deterministically [calculated](https://docs.iden3.io/protocol/spec/#genesis-id) from the Genesis State) is replaced by a new identifier, namely the `Profile ID`. 

A Profile ID is generated starting from the `ID` and hashing it with a (random) nonce. 

`ProfileID` has the same [structure as `ID`](./spec.md#identifier-format). It is a byte array of 31 bytes, in base58.

[ `IDtype` (2 bytes) | `profile_state` (27 bytes) | `checksum` (2 bytes) ]

- `IDtype` :  genesis ID type
- `profile_state` : First 27 bytes from the poseidonHash(`ID`, `profile_nonce`), where `profile_nonce` is any random number
- `checksum` Addition (with overflow) of all the ID bytes Little Endian 16 bits ([ `typeID`| `profile_state`])

> Here's how the [checksum](https://github.com/iden3/go-iden3-core/blob/2f1886532b353d1eb550ccc790cb5a6dc5bc7b32/core/id.go#L118) is calculated

Profile IDs are irreversible and indistinguishable:

- **Irreversible**, thanks to the properties of the underlying hash function, meaning the is impossible to retrieve the `ID` from a `Profile ID`, unless you know the nonce.  
- **indistinguishable**, the data format of Profile ID is the same as ID. It follows that an external party cannot tell if an identity is being identified by its Global Unique Identifier or by one of its many Profile IDs.

An Identity can now receive claims to a specific Identity Profile. An Identity Profile keeps all the properties of normal [Iden3 Identities](./spec.md#identity) while adding:

- **Anti-track**

Since users are no longer consistently identified with a persistent identifier in their interactions across different platforms, it becomes impossible to track the action of a single user. Even if platforms collude.

- **Faculty to decide which profile to show**

Users can decide which profiles to show as it is only based on the nonce. The Identity Profile is not tied to a specific Issuer or Verifier. An Identity can create an Identity Profile and reuse it across interaction with different actors, for example in the case of a Profile with all their business information. For interactions that require the maximum level of privacy, an Identity can create a single-use Identity Profile. 

- **Reusability of claims across different profiles**

Users can get claims issued to an Identity Profile (or to their global  `ID`) and generate proof, based on these claims, from a different Identity Profile. The Verifier will be only able to see a valid proof coming from the Identity Profile that the user decided to use. No connection between the two identities is leaked.

Identity Profiles do not represent any additional attack vector for the security of the protocol. While the nonce has to be kept secret, losing the nonce will only reveal the link between an `ID` and a `Profile ID` without any risk of losing control of the identity. **The control of an Identity is still managed by the underlying [Private Key](./spec#keys)**

## GIST

GIST, namely Global Identities State Tree, is a [Sparse Merkle Tree](../getting-started/mt.md) that contains the state of all the identities that use the protocol. In particular, each leaf is indexed by the hash of its `ID` (key of the leaf) and contains the most recent state of that Identity (value of the leaf). 

The root of the GIST is stored inside the [State Contract](../contracts/state.md). Everytime a user executes a state transition function the new state will be added to the GIST and the root of the GIST will be updated. 

This design allows users to prove ownership of an Identity without revealing which identity they own! 

<div align="center">
<img src= "../imgs/GIST.png" align="center" width="400"/>
<div align="center"><span style="font-size: 17px;"></div>
</div>

## Auth V2

> The Global Identities State Tree doesn't replace the V1 Identity State Tree. 
