# Protocol Specs V2

The new version of the Protocol contains features such as Identity Profiles and Global Identities State Tree (GIST). These features are designed to add a further level of identity privacy: users are now able to hide their [Identifier](../getting-started/identity/identifier.md) when interacting with others.

## Identity Profiles

This new feature allows users to hide their Identifier during interactions. Technically, the global unique Identifier of the user (hereafter defined as `GenesisID`) is replaced by a new Identifier, namely the `IdentityProfile`. 

An Identity Profile is generated starting from the `GenesisID` and hashing it with a (random) nonce. 

`Identity Profile` has the same [structure as the `Genesis ID`](./spec.md#identifier-format). It is a byte array of 31 bytes, encoded in base58.

[ `IDtype` (2 bytes) | `profile_state` (27 bytes) | `checksum` (2 bytes) ]

- `IDtype` :  same type as the one encoded in `Genesis ID`
- `profile_state` : First 27 bytes from the poseidonHash(`Genesis ID`, `profile_nonce`), where `profile_nonce` is any random number
- `checksum` Addition (with overflow) of all the ID bytes Little Endian 16 bits ([ `typeID`| `profile_state`])

> Here's how the [checksum](https://github.com/iden3/go-iden3-core/blob/2f1886532b353d1eb550ccc790cb5a6dc5bc7b32/core/id.go#L118) is calculated

Identity Profiles are irreversible and indistinguishable:

- **Irreversible**, thanks to the properties of the underlying hash function, meaning that it is impossible to retrieve the `Genesis ID` from an `Identity Profile`, unless you know the nonce.  
- **indistinguishable**, the data format of Identity Profiles is the same as Genesis IDs. It follows that an external party cannot tell if an identity is being identified by its Genesis ID or by one of its many Identity Profiles.

An Identity can now receive claims to a specific Identity Profile. An Identity Profile keeps all the properties of normal [Iden3 Identities](./spec.md#identity) while adding:

- **Anti-track**

Since users are no longer consistently identified with a persistent identifier in their interactions across different platforms, it becomes impossible to track the action of a single user. Even if platforms collude.

- **Faculty to decide which profile to show**

Users can decide which profiles to show as it is only based on the nonce. The Identity Profile is not tied to a specific Issuer or Verifier. An Identity can create an Identity Profile and reuse it across interaction with different actors, for example in the case of a Profile with all their business information just by reusing the same nonce. For interactions that require the maximum level of privacy, an Identity can create a single-use Identity Profile by choosing a random nonce and never reusing it again. 

- **Reusability of claims across different profiles**

Users can get claims issued to an Identity Profile (or to their global Genesis ID) and generate proof, based on these claims, from a different Identity Profile. The Verifier will be only able to see a valid proof coming from the Identity Profile that the user decided to use. No connection between the two identities is leaked.

Identity Profiles do not represent any additional attack vector for the security of the protocol. While the nonce has to be kept secret, losing the nonce will only reveal the link between the `Genesis ID` and an `Identity Profile` without any risk of losing control of the identity. **The control of an Identity is still managed by the underlying [Private Key](./spec#keys)**

## GIST

GIST, namely Global Identities State Tree, is a [Sparse Merkle Tree](../getting-started/mt.md) that contains the state of all the identities of the users that use Iden3 protocol. In particular, each leaf is indexed by the hash of its `Genesis ID` (key of the leaf) and contains the most recent state of that Identity (value of the leaf).

The root of the GIST is stored inside the [State Contract](../contracts/state.md). Every time a user executes a [State Transition function](../getting-started/state-transition/state-transition.md), the new state of an identity is [added to the GIST stored on-chain](https://github.com/iden3/contracts/blob/master/contracts/state/StateV2.sol#L190)

```solidity
_gistData.add(PoseidonUnit1L.poseidon([id]), newState);
```

This design allows users to prove ownership of an Identity without revealing which one is yours their Genesis ID!

<div align="center">
<img src= "../imgs/GIST.png" align="center" width="400"/>
<div align="center"><span style="font-size: 17px;"></div>
</div>

> The Global Identities State Tree doesn't replace the V1 Identity State Tree! The State Transition still updates user's [Identity State Tree](../getting-started/identity/identity-structure.md). The update of the GIST is am extra step executed inside the contract after the State Transition is executed

## Authentication Circuit V2

Authentication Circuit V2 supports the Identity Profile feature. The scope of the circuit is the same as the [Authentication Circuit V1](../protocol/main-circuits.md#authentication): to allow a user to prove that he/she is in control of an identity by signing a challenge. By verifying an auth proof, a subject can authenticate a user by their Identifier.

In V1 the Identifier of the user was always its `GenesisID`. In V2 the Identifier of the user can hide their actual GenesisID and authenticate themselves with a different one, namely the `Identity Profile`. So, how is that possible? 

The Auth V2 circuit doesn't modify the core logic of the Circuit. It maintains the same logic while adding two features: 

**Check the inclusion of the genesis ID inside the GIST**

The circuit takes the root of the GIST (stored on-chain inside the State Contract) and the merkle proof of inclusion of the user inside the GIST as [inputs](https://github.com/iden3/circuits/blob/feature/circuits_v0.2/circuits/lib/authV2.circom#L40). 

The logic of the circuit verifies whether the leaf (that contains the hash of the user's genesisID as a key and the user's state as value) is [included inside the GIST](https://github.com/iden3/circuits/blob/feature/circuits_v0.2/circuits/lib/authV2.circom#L76).

**Calculate the Identity Profile and return it as output**

The circuit takes a `profileNonce` as [input](https://github.com/iden3/circuits/blob/feature/circuits_v0.2/circuits/lib/authV2.circom#L14). 

The logic of the circuit [calculates](https://github.com/iden3/circuits/blob/feature/circuits_v0.2/circuits/lib/authV2.circom#L101) the `Identity Profile` by hashing together the `GenesisID` and the `profileNonce` and returns it as the only output of the circuit. 

If a user wants to authenticate using their `GenesisID` it is still possible by passing 0 as Profile Nonce.