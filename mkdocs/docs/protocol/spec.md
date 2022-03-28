#  Iden3 protocol specs (version 0)

> These specifications are still being built and updated regularly. Consider it work in progress.

## Basis

### Glossary

- Issuer: an actor who makes a claim.
- Holder: an actor who has received a claim.
- Verifier: an actor who verifies if the content of a claim is issued by a specific identity and held by another specific identity.
- Credential: data that is needed to prove that a claim is issued by a specific identity and held by another specific identity. This data is composed of a **claim and a proof**.

### MerkleTree

A Merkle tree (MT) or hash tree is a cryptographically verifiable data structure where every "leaf" node of the tree contains the cryptographic hash of a data block, and every non-leaf node contains the cryptographic hash of its child nodes.

The MTs used in the protocol has a few particularities:
- Binary: each node can only have two children.
- Sparse and Deterministic: the contained data is indexed, and each data block is placed at the leaf that corresponds to that data block's index, so insert order doesn't influence the final Merkle Tree Root. This also means that some nodes are empty.
- ZK friendly: the used hash function, [poseidon](https://www.poseidon-hash.info/), plays well with the Zero-Knowledge proofs (ZKP) used in different parts of the protocol.

In order to ensure that these particularities are respected and to have a history of all changes that occurred on different trees (without revealing the actual content stored in the leaves), **the root of each MT is indirectly stored on the blockchain**. EVM-based blockchains are chosen for this purpose.

The `Merkle Tree` specification is defined in [this document](https://github.com/iden3/iden3-docs/blob/master/source/docs/MerkleTree.pdf). In the future, the MT implementation could be changed.

### Zero-Knowledge proof (ZKP)

In cryptography, a zero-knowledge proof is a method by which one party (the prover) can prove to another party (the verifier) that prover knows a value x that fulfills some constraints, without revealing any information apart from the fact that he/she knows the value x.

The technologies that implement these methods are evolving rapidly. As of now, the protocol uses zkSNARKs Groth16, but in the future, the zk protocol could be changed.
zkSNARK stands for "Zero-Knowledge Succinct Non-Interactive Argument of Knowledge", and have the following properties:

- Non-interactive: with a single message (credential) from the prover, the verifier can verify the proof.  This is good because it allows sending proofs to a smart contract that can verify them immediately.
- Efficient verification: it's computationally efficient to verify proofs, both in terms of size and operations. This is good for the protocol because verification happens on the blockchain with its inherent costs.
- Heavy proof generation: generating a proof is computationally very expensive and can be time-consuming even with powerful hardware.
- Setup: a pre-existing setup between the prover and verifier is required for the construction of zkSNARKs. In order to ensure that the verifier can not cheat one has to be sure that the keys used for the setup were destroyed. There are protocols to ensure that, resulting in a "trusted setup".

More technical information about zkSNARKs on [this article](https://medium.com/@VitalikButerin/zk-snarks-under-the-hood-b33151a013f6) by Vitalik Buterin.

## Claims

### Definition

A claim is a statement made by one identity about another identity or itself.
Each claim is composed of two parts: index and value part.
Claims are stored in the leaves of an MT. The index is hashed and used to determine in which leaf position the value of the claim will be stored.

A special transition validation functions can be used to restrict how leaves are stored in the MT, e.g. make the MT append-only, so that leaves can't be updated or deleted, just added.

### Properties

- It's impossible to generate a proof of a statement on behalf of an identity without its consent.
- Claims can be revoked.
- Claims can be updated by creating new versions. When a claim is revoked, no further versions can be made.  Claims can be set to be updatable or not with a flag (See claim structure).

```mermaid
graph LR
    revoked(Revoked claim)
    no-claim-->v0
    v0-->v1
    v1-.->vN
    vN-->vN1
    vN1-->revoked
    
    no-claim(No claim)
    v0(Claim v0)
    v1(Claim v1)
    vN(Claim vN)
    vN1(Claim vN+1)
```

- Claims can be verified. This means that it's possible to demonstrate cryptographically that a given claim is:
    - Issued by a specific identity.
    - Not revoked.
    - Is of the last version of that claim if it's updatable.
- There are two types of claims regarding destination
    - Claims about self properties. Example: Operational Key, Ethereum Address, etc.
    - Claims about another identity property
        - (another) identity has a property: directional relation between an identity and a property (See claim structure: identity stored in hIndex, i_1)
        - property is owned by (another) identity: directional relation between a property and an identity (See claim structure: identity stored in hValue, v_1)

> NOTE: Some of these properties are only guaranteed by a transition validation function (explained above in this document).

### Structure

```
h_i = H(i_0, i_1, i_2, i_3)
h_v = H(v_0, v_1, v_2, v_3)
h_t = H(h_i, h_v)
```

```mermaid
graph TD
Hi-->i0
Hi-->i1
Hi-->i2
Hi-->i3

Hv-->v0
Hv-->v1
Hv-->v2
Hv-->v3

Ht-->Hi
Ht-->Hv
```

```
Index:
 i_0: [ 128 bits ] claim schema
      [ 32 bits ] header flags
          [3] Subject:
            000: A.1 Self
            001: invalid
            010: A.2.i OtherIden Index
            011: A.2.v OtherIden Value
            100: B.i Object Index
            101: B.v Object Value
          [1] Expiration: bool
          [1] Updatable: bool
          [27] 0
      [ 32 bits ] version (optional?)
      [ 61 bits ] 0 - reserved for future use
 i_1: [ 248 bits] identity (case b) (optional)
      [  5 bits ] 0
 i_2: [ 253 bits] 0
 i_3: [ 253 bits] 0
Value:
 v_0: [ 64 bits ]  revocation nonce
         [ 64 bits ]  expiration date (optional)
         [ 125 bits] 0 - reserved
 v_1: [ 248 bits] identity (case c) (optional)
        [  5 bits ] 0
 v_2: [ 253 bits] 0
 v_3: [ 253 bits] 0
```

### Reliability of a claim content

The correctness of what is said in a claim is not verifiable by the protocol, since every identity is free to claim whatever they want. Since it's possible to know which identity issued each claim, the trust / reputation that the issuer has can affect its credibility.

However, the protocol can guarantee exclusivity: there cannot be two claims with the same index. So it's impossible that an identity claims that a particular property (index part of the claim) is linked to two different identities (value part of the claim) at the same time.

## Keys

Keys are cryptographic elements that can be used to sign data. In the protocol keys are used to authenticate certain interactions.

These keys require the authorization of the identity who owns them to be used. This is done by adding a claim with a specific schema, linking the key(s) with the identity. 

This way each time that a key is used for signing, the identity can (and must) prove the ownership of that key and the fact that the key is not revoked.

<!-- ==**TODO**==
== - explain what role the keys play in the iden3 protocol. Authenticate! explain that the keys are used to sign to authenticate an identity ==
== - and that in the case of zero knowledge it is not done as it would be done 'traditionally', but that the zero knowledge proof itself already shows that the sender knows the private key == -->

### Types of keys

- Baby Jubjub: used for authentication. This type of key is designed to be efficient when working with zkSNARKs.
  The `Baby Jubjub Elliptic Curve` specification is defined in [this document](https://github.com/iden3/iden3-docs/blob/master/source/docs/Baby-Jubjub.pdf).


## Identity

### Definition

An `identity` is characterized by the claims, that the identity has issued, and the claims, that the identity has received from other identities, in other words: an identity is built by what the identity has said, and what others have said about the identity.

Each claim that identity issues can be cryptographically proved and verified, ensuring that the claim existed under identity at a certain timestamp.

```mermaid
graph TD
    Root-->A
    Root-->B
    
    A-->C
    A-->D
    B-->E
    B-->F
    
    C-->G
    C-->H
    D-->I
    D-->J
    E-->K
    E-->L
    F-->M
    F-->N
    
    G["claim"]
    H["claim"]
    I["claim"]
    J["claim"]
    K["claim"]
    L["claim"]
    M["claim"]
    N["claim"]
```

To accomplish this (and other properties covered in this document), identities are built by [MerkleTrees](#MerkleTree), where the claims are placed as the leaves, and the `Root` is published (indirectly through identity state) in the blockchain. With this construction, the identities can issue, update and revoke claims.

The protocol construction is designed to enable Zero-Knowledge features, which means for example that identities have the ability to prove with Zero-Knowledge the ownership of properties of claims in issued and received claims among other capabilities and verify that claim is not revoked.

### Genesis ID

#### Description

Each identity has a unique `identifier` that is determined by the initial identity state (hash of its MerkleTree roots), called `Genesis ID`, under which the initial claims are placed, that are the ones contained in the initial state of the identity.

<!-- TODO: For the initial implementation of the protocol, the Genesis Claims Tree will contain at least a [claim of authorization of the Operational key](#link-to-the-spec-of-the-claim-once-is-done), that allows operating in the name of identity. -->
For the initial implementation of the protocol, the Genesis Claims Tree will contain at least a claim of authorization of the Operational key, that allows operating in the name of identity.
 <!-- (identities operated by Smart Contracts are not specified in this version). -->

While an identity doesn't add, update or revoke claims after the Genesis State, its identity state does not need to be published on the blockchain, and the `Genesis Claims` can be verified directly against the `Genesis ID`, as it is built by the Merkle Root that holds that claims.

> NOTE: The Genesis ID is calculated with the Identity State as a hash of Genesis Claims Tree Root, an empty Revocation Tree Root and an empty Roots Tree Root.

#### Identifier format

The Identifier of an identity is determined by the identity type and the `Genesis Identity State`, what we call the `Genesis ID`. This is built by creating a MerkleTree that holds the initial state claims, calculating its Root, hashing it together with an empty Revocation Tree Root and an empty Roots Tree Root. Then taking the first 27 bytes of the result and adding 2 bytes at the beginning to specify the identity type, and 2 bytes at the end for checksum. In sum, the identifier is a byte array of 31 bytes, encoded in base58.

The **identity type** specifies the specs that the identity follows, such as the hash function used by the identity. In this way, when the hash function changes, the identifiers of the identities will change, allowing to identify of which type and protocol is one identity.

Identifier structure:
- `ID` (genesis): base58 [ `type` | `genesis_state` | `checksum` ]
	- `type`: 2 bytes specifying the type
	- `genesis_state`: first 27 bytes from the identity state (using the genesis claim merkle tree)
	- `checksum`: addition (with overflow) of all ID bytes Little Endian 16 bits ( [ `type` | `genesis_state` ] )


#### Identity state

The **identity states** are published on the blockchain under the identity identifier, anchoring the state of the identity with the timestamp when it is published. In this way, the claims of the identity can be proved against the anchored identity state at a certain timestamp. To transition from a state to another one, identities follow the transition functions.

The identity states can be published on the blockchain **directly** performing the transaction to publish the root, or **indirectly** using a Relay. 
<!-- The indirect method is described in the section [Indirect Identities](#Indirect-identities). -->

The `genesis state` is the initial state of any identity, and does not need to be published in the blockchain, as the claims under it can be verified against the identity identifier itself (that contains that identity state).

![](../../imgs/identity_state_transition.png)


#### Identity state transition function
The `ITF` (Identity state Transition Function) is verified each time that the State is updated in order to ensure that the identities follow the protocol when updating the state.

Identity MerkleTree is a sparse binary tree, that only allows the addition of leaves (no edition nor deletion), and to add new claims, update through versions and revoke, needs to be done according to the `ITF`. To ensure this we use Zero-Knowledge proofs, in a way that when an identity is publishing a new state to the Smart Contract, also sends a zero-knowledge proof (`π`) proving that the `ϕ` is satisfied following the `ITF`. In this way, all the identity states published on the blockchain are validated to be following the protocol.

> In the initial version of the implementation there will not be checks that trees are append-only in the Smart Contract due to the complex computation needed to generate zk-proofs of multiple claim additions, which is needed for scalability.

The full circuit can be found here: https://github.com/iden3/circuits/blob/master/circuits/idState.circom

<!--
##### Direct identity ITF_min

###### Addition of non-updatable claim
Prove that a leaf in the MT (claim) is only added but never deleted nor changed

- `R_i -> leaf_claim==0`
- `R_i+1 -> leaf_claim==claim`
> 2 MerkleProofs with same siblings

###### Addition of claim with versions
- `v0`
    - `R_i -> leaf_claim_v0==0`
    - `R_i+1 -> leaf_claim_v0==claim_v0`
> 2 MerkleProofs with same siblings

If an updatable claim is added to the MT with version `v!=0`, claim version `v-1` must already exist in the MT

- `v_n+1`
    - `R_i -> leaf_claim_vn==claim_vn`
    - `R_i -> leaf_claim_vn+1==0`
    - `R_i+1 -> leaf_claim_vn+1==claim_vn+1`
> 1 MerkleProof with siblings
> 2 MerkleProof with same siblings
-->

<!-- ##### Indirect identity ITF_min
==TODO== -->

### Identity Ownership

We prove the identity ownership inside a zkSNARK proof. This means that the user can generate a zk-proof that he/she knows a `private key` corresponding to `operational key for authorization` claim added to Claims Tree, without revealing the `claim` and its position. This is codified inside a circom circuit, which can be used in other circuits (such as the `id state update` circuit).

The full circuit can be found here: https://github.com/iden3/circuits/blob/master/circuits/idOwnershipBySignature.circom

### Identity key rotation

An identity can self-issue and revoke many `private keys` and corresponding `claims` of the type `operational key authorization` enabling key rotation in that way. To support verification of such claims identity state should be publicly available in the blockchain. An identity can publish the state to blockchain directly or via the Relay.

Any private key, for which a corresponding claim exists in the identity claims tree and does not exist in the identity revocation tree, can be used to create a zero-knowledge proof of valid credentials. Such proof should pass verification by a verifier as it is able to check the latest identity state in the blockchain.

In the same way, any valid and non-revoked identity private key can be used to create a valid zk_proof for identity state transition function.

Note: An identity may lose some privacy as far as it needs to disclose its state to a verifier, which can track all the proofs of the same identity in that way. However, this can be mitigated if the identity state is published to the blockchain via the Relay. In this case, only the Relay state needs to be disclosed to a verifier.

<!-- ==TBD: Identity recovery procedure in case of a private key and state database loss== -->

### Identity Revocation

When identity revokes all `claims` of the type `operational key authorization`, it is considered as `revoked`, because this identity can no longer create proofs.

## Interaction between Identity and Claims

### Identity State Update

The Identity State Update is the procedure used to update information about
what this Identity claims. This involves three different actions:
- Add a Claim
- Update Updatable Claim (by incrementing the version and changing the Claim value part)
- Revoke a Claim

<!--
The Identity State Update can be generalized as an `ITF_min` (minor Identity Transition
Function) that performs a minor update of the Identity (where a minor update
only concerns the Claims of the Identity but not the Identity itself).
-->

Definitions

- `IdS`: Identity State
- `ClT`: Claims Tree
    - `ClR`: Claims Tree Root
- `ReT`: Revocation Tree
    - `ReR`: Revocation Tree Root
- `RoT`: Roots Tree
    - `RoR`: Roots Tree Root

The `IdS` (Identity State) is calculated concatenating the roots of the three user trees:
- `IdS`: `H(ClR || ReR || RoR)`
    - Where `H` is the Hash function defined by the Identity Type (for example Poseidon)

All trees are SMT (Sparse Merkle Tree) and use the hash function defined by the Identity Type.
- Leaves in `ClT` (Claims Tree) are Claims ((4 + 4) * 253 bits = 253B)
```
See Claims Structure
```
- Leaves in `ReT` (Revocation Tree) are Revocation Nonce + Version (64 + 32 bits = 12B)
```
Revocation Tree Leaf:
leaf: [ 64 bits ] revocation nonce
      [ 32 bits ] version
      [157 bits ] 0
```
- Leaves in `RoT` (Roots Tree) are tree Roots (from the Claims Tree) (253 bits ≈ 32B)
```
Roots Tree Leaf:
leaf: [253 bits ] tree root
```

![](https://i.imgur.com/3ZS1ZvJ.png)
> Identity State Diagram for Direct Identity

As seen in the diagram, only the `IdS` is stored in the Blockchain.  In order
to save stored bytes in the blockchain, it is desirable that only one "hash"
representing the current state of the Identity is stored in the Smart Contract.
This one "hash" is the `IdS` (Identity State), which is linked to a timestamp
and a block in the blockchain.

All the public data must be made available for any Holder so that
they can build fresh merkle tree proofs of both the `ReT` and `RoT`.  This
allows the Holder to:

1. Prove recent non-revocation / "current" version without interaction with the issuer.
2. Hide a particular `ClR` from all the `ClR`s, to avoid allowing the issuer to
   discover a Claim hidden behind a ZK proof. For this purpose `ClR` added to `RoR`

The place and method to access the publicly available data is specified in the
Identities State Smart Contract. Two possible initial options are:

- IPFS, by adding a link to an [IPNS address](https://docs.ipfs.io/guides/concepts/ipns/) (example:
  `ipfs://ipns/QmSrPmbaUKA3ZodhzPWZnpFgcPMFWF4QsxXbkWfEptTBJd`) which contains
  a standardized structure of the data.
- HTTPS, by adding a link to an HTTPS endpoint (example:
  `https://kyc.iden3.io/api/v1/public-state/aabbccdd` which offers the data
  following a standardized API.

#### Publish Claims

Publishing a Claim involves first adding a new leaf to the `ClT`, which updates
the Identity `ClR`.  Claims can be optionally published in batches, adding more
than one leaf to the `ClT` in a single transaction.  After the `ClT` has been
updated, the Identity must follow an Identity State Update so that anyone is
able to verify the newly added Claims.  This involves adding the new `ClR` to
the `RoT` which in turn will update the `RoR`.  Afterwards, the new `IdS` is
calculated and through a transaction it is updated in the Identities State
Smart Contract (from now on, referred to as "the Smart Contract") in the
blockchain.  Once the updated `IdS` is in the Smart Contract, anyone can verify
the validity of the newly added Claims.

The updating procedure of the `IdS` in the Smart Contract can be achieved
through the following means with the following properties:
- *Bad scalability (no batch), good privacy, correctness*: The identity uploads
  the new `IdS` to the Smart Contract, with a proof of a correct transition from
  the old `IdS` to the new one. Only one claim is added to the `ClT` in the
  transition.
- *Good scalability (batch), good privacy, correctness*: Same as before, but
  many claims are added (batch) in the transition (with a single proof for all
  newly added claims)
- *Good scalability (batch), good privacy, no correctness*: The identity uploads
  the new `IdS` to the Smart Contract, without proving correctness on the
  transition.

The correctness properties mentioned here are the following:
- Revocation of a Claim can't be later undone
- Updatable Claims are only updated with increasing versions, and only one
  version is valid at a time.

The choice of having correctness guarantees or not is specified in the Identity
Type, so that any Verifier knows about the guarantees provided by the protocol
for the Issuer Claims.

> NOTE: Good scalability refers to the verification process and the costs related to the Smart Contract. Batching with zkSNARKs can have a high computation load to the prover.

#### Revocation tree

Sometimes it's desirable for an Identity to invalidate some statement made
through a Claim.  For regular Claims, this involves revoking, a process that's
ideally irreversible, and allows any verifier to be aware that an already
published Claim is made invalid by the Issuer Identity.  Similarly, for updatable
Claims, there must be a mechanism to invalidate old versions when a new one is
published.  Since confirming the current validity of a Claim is a parallel
process to confirming that a Claim was published at some point, the "current
validation" process can be separated.

Separating these two processes allows a design in which the `ClT` (Claim Tree)
remains private, but the revocation/version information is public, allowing a
holder to generate a fresh proof of "current validity" without requesting
access to the private `ClT`.

To achieve this, every Identity has a `ClT` (Claim Tree) and a separate `ReT`
(Revocation Tree).  While the claim tree would be private and only the root
public, the revocation tree would be entirely public.  The roots of both trees
(`ClT` and `ReT`) are linked via the `IdS` (Identity State) which is published
in the Smart Contract.  The revocation tree could be published in IPFS or any
other public storage system.

Proving that a claim is valid (and thus not revoked/updated) is separated into
two proofs:
1. Prove that the claim was issued at some time T (this proof is generated once
   by the issuer and uses a `IdS`-`ClR` at time T stored in the Smart Contract)
2. Prove that the claim has not been revoked/updated recently (this proof is
   generated by the holder with a recent `ReR` (Revocation Tree Root) by
   querying the public `ReT` (Revocation Tree), and verified against a recent
   `IdS`).

#### Revoke Claims

In order to not reveal anything about the content of the claim in the
`ReT`, the Claim contains a revocation nonce in the value part, which
is added as a leaf in the `ReT` to revocate the Claim.

In order to forbid undoing revocation of a claim, the `ReT` needs to follow some
transition rules like `ClT`, enforced by a ZK proof (for space and verification
efficiency).

Apart from the revoking procedure, there's a method to define the validity of a
Claim based on expiration, by explicitly setting an expiration date in the
Claim (See Claim Structure).  Revoking and Expiration are compatible methods to
invalidate Claims.

#### Update Claims

To update a Claim, first, a new Claim is added to the `ClT` with an increased version
value in the index position in the claim (notice that the previous version of the 
claim is not touched). Then, a leaf is added to the `ReT` containing the revocation
nonce and the maximum invalid version (that is, all Claims with that nonce and
version equal or lower to the one in the leaf are invalid).  This means that
when a Claim is updated, the same revocation nonce is used in the Claim.

In order to forbid downgrading the version of a Claim, and forcing to have only
one valid updatable Claim at a time, the `ReT` needs to follow some transition
rules like the `ClT` does, enforced by a ZK proof (for space and verification
efficiency).

Updating and Revoking are compatible methods to invalidate Claims: an updatable
Claim can be revoked, meaning that no future (or past) updates will be valid.

In case when a claim needs to be revoked completely, without possibility to update. 
The maximum version and the revocation nonce should be added to `ReT`

### Prove Claims (Credentials)

Nomenclature
- MTP: Merkle Tree Proof. The list of siblings in a path from a leaf to the
  root.

#### Prove that a claim was issued at time at least t

- Requires proving a link between the Claim and an `IdS_t` (Identity State at
  time t) published in the Smart Contract.  This proof requires:
    - Claim
    - t
    - MTP Claim -> `ClR_t`
    - `RoR_t` (Roots Tree at time t)
    - `ReR_t` (Revocation Tree Root at time t)
    - `IdS_t`

Where `t` is any time.

#### Prove that the claim is currently valid

##### Prove that a claim hasn't been recently revoked

Where `t` is a recent time.

- Requires proving the inexistence of a link between the Claim revocation nonce
  and a recent `IdS_t` (`t` must be recent according to the verifier
  requirements [1]) published in the Smart Contract.  This proof requires:
    - Claim (Nonce)
    - t
    - MTP !Nonce -> `ReR_t`
    - `ClR_t`
    - `RoR_t`
    - `IdS_t`

[1] The verifier needs to decide a time span to define how recent the
`IdS_t` used in the proof needs to be.  Always requiring the current `IdS`
could lead to data races, so it's better to require an `IdS` that is no more
than X hours old.

##### Proof of last version

This is very similar to proving that a claim hasn't been recently revoked,
except that not only the nonce in the claim is checked, but also the version.

- Requires proving the inexistence of a link between the Claim revocation nonce
  + version and a recent `IdS_t` (`t` must be recent according to the verifier
    requirements [1]) published in the Smart Contract.  This proof requires:
    - Claim (Nonce, Version)
    - t
    - MTP !(Nonce, Version) -> `ReR_t`
    - `ClR_t`
    - `RoR_t`
    - `IdS_t`

Where `t` is a recent time.

##### Proof of non-Expiration

A Claim can be expirable by setting the expiration flag in the options and
specifying an expiration date in unix timestamp format in the corresponding
claim value part (see Claim Format).

#### Zero-Knowledge proof of valid Credential

A Zero-Knowledge proof allows hiding some information about a Claim while
proving that it was issued by a particular Identity and that it's currently
valid. The same checks performed in the following sections are done:
- Prove that a claim was issued at time at least t
- Prove that the claim is currently valid

In the proof that shows "that a claim was issued at time at least t" there's an
additional part that is added to hide the particular `IdS_t1` that is used (in
order to hide the Claim from the Issuer, See Appendix Title 2).  The proof then
requires:
    - Claim
    - t
    - MTP Claim -> `ClR_t1`
    - `RoR_t1` (Roots Tree at time t1)
    - `ReR_t1` (Revocation Tree Root at time t1)
    - `IdS_t1`
    - MTP `ClR_t1` ->`RoR_t2`
    - `ClR_t2` (Claims Tree Root at time t2)
    - `ReR_t2` (Revocation Tree Root at time t2)
    - `IdS_t2`

Where `t1` is a any time and `t2` is a recent time.

The full circuit can be found at: https://github.com/iden3/circuits/blob/master/circuits/credential.circom


[//]: # (### Indirect Identities)

[//]: # ()
[//]: # (Sometimes it's not practical for an Identity to publish its Identity State)

[//]: # (directly to the Smart Contract.  Some reasons could be:)

[//]: # (- The Identity publishes very few Claims and wants a more lightweight)

[//]: # (  procedure.)

[//]: # (- The Identity doesn't own cryptocurrency to call the Smart Contract.)

[//]: # (- The Identity wants to hide identifier from verifier.)

[//]: # ()
[//]: # (In those cases, the Identity must specify <!--in it's creation--> that it will be)

[//]: # (using a Relay which will indirectly link its claims to the Smart Contract via)

[//]: # (the Relay Identity State.  In this case, we say that the Identity under the)

[//]: # (Relay is an indirect Identity, and the Claims that an indirect Identity issues)

[//]: # (are Indirect Claims.)

[//]: # ()
[//]: # (Indirect claims are claims that appear in the user merkle tree whose Identity)

[//]: # (State appears in a SetRoot Claim that appears in the Relay merkle tree whose)

[//]: # (Identity State is in the blockchain.)

[//]: # ()
[//]: # (As mentioned, the claims of the Relay contain Identity States of other)

[//]: # (Identities.  If correctness guarantees are desired, updating these claims requires two levels of proofs:)

[//]: # (- Correct transition between the Relay Identity States)

[//]: # (- Correct transition between the Identity States that appear in the Relay Claims &#40;in the Relay Claims Tree&#41;)

[//]: # ()
[//]: # (This can be achieved through the following means with the following properties)

[//]: # (- *Bad scalability &#40;no batch&#41;, no privacy, correctness*: All the proofs are)

[//]: # (  public &#40;and long&#41;)

[//]: # (- *Bad scalability &#40;no batch&#41;, some privacy, correctness*: Only the user root)

[//]: # (  correct transition proofs are private &#40;zero-knowledge&#41;)

[//]: # (- *Good scalability &#40;batch&#41;, medium privacy, correctness*: All the proofs are)

[//]: # (  private, but the relay builds the user root correct transaction proofs &#40;the)

[//]: # (  relay learns about user claims&#41;)

[//]: # (- *Good scalability &#40;batch&#41;, good privacy, correctness*: All the proofs are)

[//]: # (  private, the relay builds a relay root transition proof with zk proofs)

[//]: # (  from the users. Requires 1 level of recursion in zk proofs.)

[//]: # (- *Good scalability &#40;batch&#41;, good privacy, no correctness*: The relay uploads a)

[//]: # (  new root after adding set root claims with updated users claims.  User root)

[//]: # (  updates are not proved for correctness.  The relay can add arbitrary user)

[//]: # (  claims.)

[//]: # ()
[//]: # (```)

[//]: # (                     ETH)

[//]: # (                      /\)

[//]: # (                      \/)

[//]: # (                      ^)

[//]: # (                      |)

[//]: # (                      |)

[//]: # (                     root)

[//]: # (                      /\     Relay ID's)

[//]: # (                     /  \    Merkle Tree)

[//]: # (                    /____\)

[//]: # (                    |)

[//]: # (             [sub id state claim]   <-- Direct Claim)

[//]: # (                      ^)

[//]: # (                      |)

[//]: # (                      |)

[//]: # (                     root)

[//]: # (                      /\     Indirect ID's)

[//]: # (                     /  \    Merkle Tree)

[//]: # (                    /____\)

[//]: # (                      |)

[//]: # (                   [claim]    <--- Indirect Claim)

[//]: # ()
[//]: # (```)

[//]: # ()
[//]: # (The correctness properties mentioned here are the following:)

[//]: # (- Revocation of a Claim can't be later undone)

[//]: # (- Updatable Claims are only updated with increasing versions, and only one)

[//]: # (  version is valid at a time.)

[//]: # (- An indirect claim can only be published by whoever controls the kOp of the)

[//]: # (  Identity.)

[//]: # (- An indirect claim can only be revoked/updated by whoever controls the kOp of)

[//]: # (  the Identity.)

[//]: # ()
[//]: # (![]&#40;https://i.imgur.com/szBzG1x.png&#41;)

[//]: # ()
[//]: # (> Identity State Diagram for Indirect Identity, Option A)

[//]: # ()
[//]: # (![]&#40;https://i.imgur.com/mRIJa23.png&#41;)

[//]: # ()
[//]: # (> Identity State Diagram for Indirect Identity, Option B)

[//]: # ()
[//]: # (As seen in the two diagrams, the Roots Trees and Revocation Trees of the)

[//]: # (Identities under the Relay can be linked to the Relay State in two ways.)

[//]: # ()
[//]: # (==TODO: Decide which one==)

[//]: # ()
[//]: # (- Correctness of Identity State transition depending on the type of identity:)

[//]: # ()
[//]: # (```mermaid)

[//]: # (graph LR)

[//]: # (    claim-->direct)

[//]: # (    claim-->indirect)

[//]: # (    direct-->no_correctness)

[//]: # (    direct-->correctness_depth_1)

[//]: # (    indirect-->no_correctness)

[//]: # (    indirect-->correctness_depth_1)

[//]: # (    indirect-->correctness_depth_2+sig)

[//]: # (```)

[//]: # ()
[//]: # (Claims that must exist in the Genesis ID)

[//]: # (- Only direct claims &#40;A&#41;)

[//]: # (    - kOp claim)

[//]: # (- Indirect claims via Relay &#40;B&#41;)

[//]: # (    - kOp claim)

[//]: # (    - Relay Identity claim)

[//]: # ()
[//]: # (#### Prove Claims &#40;Credentials&#41;)

[//]: # ()
[//]: # (Proving Indirect Claims requires doubling the contents of a Direct Claim proof)

[//]: # (plus some extra information.)

[//]: # ()
[//]: # (nomenclature)

[//]: # (- `RIdS`: Relay Identity State)

[//]: # ()
[//]: # (##### Prove that a claim was issued at time at least T)

[//]: # ()
[//]: # (- Direct existence proof for Claim -> `IdS_t1`)

[//]: # (- Direct existence proof for SetStateClaim&#40;ID, v, `IdS_t1`&#41; -> `RIdS_t2`)

[//]: # ()
[//]: # (Where `t1` and `t2` are any times.)

[//]: # ()
[//]: # (Additionally, it must be proved that the Identity is Indirect, and that it has)

[//]: # (authorized the Relay used in the proof.)

[//]: # ()
[//]: # (Assuming that the authorization of the Relay is done via the genesis Id, and)

[//]: # (that the Relay can't be changed via a Claim update:)

[//]: # (- Direct proof for ClaimAuthRelay&#40;`RID`&#41; -> `IdS_0`)

[//]: # ()
[//]: # (##### Prove that the claim is currently valid)

[//]: # ()
[//]: # (- Direct validity proof for Claim -> `IdS_t2`)

[//]: # (- Direct validity proof for SetStateClaim&#40;ID, v, `IdS_t12`&#41; -> `RIdS_t3`)

[//]: # ()
[//]: # (Where `t2` and `t3` are recent times.)

[//]: # ()
[//]: # (The same proof of Relay authorization as in the previous section is needed.)

[//]: # (## Trust levels)

[//]: # (### Trust levels verifier-issuer)

[//]: # ()
[//]: # (The table shows what the protocol ensures to the verifier about assertions from the issuer &#40;and thus doesn't require the verifier to trust the issuer about it&#41;, in the cases:)

[//]: # (- **NC**: No transition Correctness)

[//]: # (- **C**: Transition Correctness)

[//]: # ()
[//]: # (|NC|C | Assertion |)

[//]: # (|--|--|-----------|)

[//]: # (| | | the contents of the Claims to be true |)

[//]: # (| | ✓ | the Claim was not previously revoked and now it's valid |)

[//]: # (| | ✓ | the Claim is revoked and will not be unrevoked |)

[//]: # (| | ✓ | there are no multiple valid versions of a Claim |)

[//]: # (| | ✓ | a current invalid version of a claim won't be valid in the future |)

[//]: # (| | ✓ | an indirect Claim is issued by the indirect Identity and not by the Relay |)

[//]: # (| ✓ | ✓ | the Claim was issued at least at time t |)

[//]: # (## CRL &#40;Claims Requirements Language&#41;)

[//]: # (**Claims Requirements Language** &#40;`CRL`&#41; is a standard that allows to define a set of rules &#40;`requirements`&#41; that must be accomplished in order to get access to operate with a certain action.)

[//]: # ()
[//]: # (It is defined by a specific syntax that allows entities to define a set of `requirements`, that later on an identity can take as input and from the identity's `ClaimsDB` generate a `proof` that satisfies those requirements.)

[//]: # (Then, a verifier can take the `requirements` and the `proof` and check that they are fulfilled by the prover identity.)

[//]: # ()
[//]: # (`CRL` allows to specify which Claims issued from which Identities and containing which parameters are needed to get validated. Even allows to add levels of recursion, such as defining that a certain Claim is needed that is issued from any identity that holds a Claim of a certain type from another identity holding another Claim of another type issued by another specific identity.)

[//]: # ()
[//]: # (This allows a vast amount of complex specifications of requirements with arbitrary logic.)

[//]: # ()
[//]: # (### Flow)

[//]: # (- Define a set of rules: `requirements`)

[//]: # (  - the `requirements` has a specified format, that a prover & verifier understand)

[//]: # (- a prover can take the `requirements` data packet and generate a `proof`)

[//]: # (- the verifier will take the `requirements` & `proof` check if it verifies.)

[//]: # ()
[//]: # ()
[//]: # (```mermaid)

[//]: # (graph TD)

[//]: # ()
[//]: # (subgraph Legend)

[//]: # (style e fill:#ecb3ff)

[//]: # (e&#40;&#40;entity&#41;&#41;)

[//]: # (style action fill:#b3ffd4)

[//]: # (style data fill:#c2efef)

[//]: # (data&#40;data&#41;)

[//]: # (action)

[//]: # (end)

[//]: # (```)

[//]: # ()
[//]: # (```mermaid)

[//]: # (graph LR)

[//]: # ()
[//]: # (style prover fill:#ecb3ff)

[//]: # (style verifier fill:#ecb3ff)

[//]: # ()
[//]: # (style D fill:#b3ffd4)

[//]: # (style G fill:#b3ffd4)

[//]: # (style V fill:#b3ffd4)

[//]: # ()
[//]: # (style db fill:#c2efef)

[//]: # (style requirements fill:#c2efef)

[//]: # (style proof fill:#c2efef)

[//]: # (style true fill:#c2efef)

[//]: # ()
[//]: # (verifier&#40;&#40;verifier&#41;&#41;-->D&#40;define&#41;)

[//]: # (D-->requirements)

[//]: # ()
[//]: # (prover&#40;&#40;prover&#41;&#41; -->G&#40;generate&#41;)

[//]: # (db -->G)

[//]: # (requirements-->G)

[//]: # (G-->proof)

[//]: # ()
[//]: # (requirements-->V&#40;verify&#41;)

[//]: # (proof-->V)

[//]: # ()
[//]: # (V-->true)

[//]: # (```)

[//]: # ()
[//]: # (Modules:)

[//]: # (- Automatic generation of proofs from a manifest by consulting the claim DB)

[//]: # (- Automatic generation of proofs with interactive user selection of claims)

[//]: # (- Automatic verification of a proof generated for a manifest)

[//]: # ()
[//]: # (### Levels of complexity when demonstrating claims)

[//]: # ()
[//]: # (1. demonstrate claim without Zero-Knowledge)

[//]: # (2. demonstrate claim with properties without Zero-Knowledge)

[//]: # (3. demonstrate claim with Zero-Knowledge)

[//]: # (4. demonstrate claim with properties with Zero-Knowledge)

[//]: # ()
[//]: # (Case 1 is a subset of 2, and 3 is a subset of 4. So we will work on cases 2 and 4.)

[//]: # ()
[//]: # (> Example of claim)

[//]: # (> ```)

[//]: # (> Claim)

[//]: # (>   From: ID)

[//]: # (>   Recip: self.ID)

[//]: # (>   Field1: 0xAABB)

[//]: # (>   Field2: 0xCCDD)

[//]: # (>   Field3: 0xFFEE)

[//]: # (> ```)

[//]: # ()
[//]: # (- case 2:)

[//]: # (    - claim:)

[//]: # (        - type: ClaimType)

[//]: # (        - ID: core.ID)

[//]: # (        - recip: self.ID)

[//]: # (        - fields: [)

[//]: # (            "field1": "0xAABB",)

[//]: # (            "field2": "0xCCDD")

[//]: # (        ])

[//]: # (- case 4:)

[//]: # (    - claim:)

[//]: # (        - type: ClaimType)

[//]: # (        - ID: core.ID)

[//]: # (        - recip: prv&#40;self.ID&#41;)

[//]: # (        - fields: [)

[//]: # (            "field1": prv&#40;"0xAABB"&#41; == F&#40;pub&#40;X&#41;&#41;,)

[//]: # (            "field2": "0xCCDD")

[//]: # (        ])

[//]: # ()
[//]: # (#### Things to demonstrate)

[//]: # (- combination of claims)

[//]: # (    - &#40;OR, AND, ...&#41;)

[//]: # (- links between claims)

[//]: # (    - claim1[field1]==claim2[field3])

[//]: # (    - claim1.from==claim2.from)

[//]: # (    - recursion/dependency)

[//]: # (        - claim2.from=="IDx")

[//]: # (        - claim1.from==claim2.recip)

[//]: # (        - claim1.recip==self.ID)

[//]: # (- link between claim properties and public variable)

[//]: # (    - claim[field1]==F&#40;X, Y&#41; &#40;where Y is public var&#41;)

[//]: # (- recursion)

## Identities Communications
    
### Issuer - Holder (Credential Request procedure)

The same procedure works for already issued claims, and new claims: 
- The Issuer has issued a claim linking a property to the Holder, and the Holder requests the credential of the issued claim.
- The Holder requests the issue of a new claim linking a property to the Holder.

> NOTE: In http, use polling to resolve the "Future".  In async messaging, request the resolution of the "Future" and wait for the reply.

#### Direct Claims

```mermaid
sequenceDiagram
    Holder->>IssuerServer: req. Credential + auth?
    IssuerServer->>IssuerServer: Auto/Manual check
    IssuerServer->>IssuerServer: Add Claim to MT
    IssuerServer->>Holder: Future(Credential)
    IssuerServer->>SmartContract: Publish Root
    SmartContract->>IssuerServer: Ok
    Holder->>IssuerServer: Poll(Future(Credential))
    IssuerServer->>Holder: Credential
```

##### Indirect Claims

```mermaid
sequenceDiagram
    Holder->>IssuerClient: req. Credential
    IssuerClient->>IssuerClient: Auto/Manual check
    IssuerClient->>IssuerClient: Add Claim to MT
    IssuerClient->>Relay: req. Credential (SetRoot)
    Relay->>Relay: Add Claim to MT
    Relay->>IssuerClient: Future(Credential1)
    IssuerClient->>Holder: Credential0, Future(Credential1)
    Relay->>SmartContract: Publish Root
    SmartContract->>Relay: Ok
    Holder->>Relay: Poll(Future(Credential1))
    Relay->>Holder: Credential
```

### Holder - Verifier

- Verifier requests a claim (or in general, a proof that involves some claims).
- Holder shows a proof of the claim (or in general, a proof that involves some claims) to the Verifier.

```mermaid
sequenceDiagram
    participant A
    Exchange_SC->>Exchange_SC: define CR with CRL
    A->>Exchange_SC: get CR
    Exchange_SC->>A: CR
    A->>A: build proof using Claim DB
    A->>Exchange_SC: proof
    Exchange_SC->>Exchange_SC: validate
    Exchange_SC->>Exchange_SC: action
    Exchange_SC->>A: result
```

[//]: # (## Naming system)

[//]: # (**Naming system** allows assigning unique human-readable names.)

[//]: # ()
[//]: # (==we had some ideas and implementations in the past, needs to be defined==)

[//]: # ()
[//]: # (## Identity Discovery)

[//]: # (**Identity discover** allows identities to retrieve the information about how to communicate with other identities through a decentralized p2p network.)

[//]: # ()
[//]: # (# Addendum)

[//]: # ()
[//]: # (## Appendix Title 2)

[//]: # ()
[//]: # ([2] A ZK proof of a valid claim needs to link the claim to the Identity State)

[//]: # (that is stored in the Identity State Smart Contract &#40;blockchain&#41;.  The link is)

[//]: # (done via zero-knowledge &#40;the merkle tree proof which links the tree leaf)

[//]: # (&#40;Claim&#41; to the Claims Tree Root &#40;linked to the Identity State&#41;&#41;.  When emitting)

[//]: # (a credential &#40;Claim+Proof&#41;, the issuer needs to generate the merkle tree proof)

[//]: # (for a particular Claims Tree Root &#40;and Identity State&#41;.  If the ZK proof)

[//]: # (reveals that particular Root, and this ZK is public to the issuer, the issuer)

[//]: # (can identify a small set of claims which contains the claim in the ZK proof)

[//]: # (&#40;that is, the set of claims for which the issuer calculated a proof for that)

[//]: # (particular Root&#41;.  This diminishes the privacy in great measure, and needs a)

[//]: # (solution.)

[//]: # ()
[//]: # (### Solution 1 - Periodic generation of all credentials)

[//]: # ()
[//]: # (The issuer can periodically generate proofs for every claim ever issued, and)

[//]: # (send these fresh credentials to the recipients.  This allows the recipients to)

[//]: # (use a proof with a recent root in a ZK proof so that the issuer is unable to)

[//]: # (distinguish the hidden claim between all the issued claims in that particular)

[//]: # (root.)

[//]: # ()
[//]: # (Cons:)

[//]: # (- The issuer has a periodic overhead &#40;computation + bandwidth&#41; that grows)

[//]: # (  linearly &#40;actually n*log&#40;n&#41;&#41; with the number of issued claims)

[//]: # ()
[//]: # (### Solution 2 - Root of roots)

[//]: # ()
[//]: # (A merkle tree of Roots is generated by the Issuer and uploaded to the Smart)

[//]: # (Contract.  This allows the recipients to hide the particular Root they are)

[//]: # (using to prove a claim, and to prove that the Root is valid by proving)

[//]: # (inclusion in the Roots Tree of the Issuer.)

[//]: # ()
[//]: # (Cons:)

[//]: # (- The Root Smart Contract requires a new field for each identity: the root of)

[//]: # (  the roots tree.)

[//]: # (- For correctness, the root update should prove that the root of the roots tree)

[//]: # (  is correctly updated.)

[//]: # (- To build the merkle tree proof of the roots tree, the leaves &#40;roots&#41; and)

[//]: # (  intermediate nodes are needed.  Although this data structure can be built from)

[//]: # (  public data &#40;the history of roots from the blockchain&#41;, it should be cached)

[//]: # (  somewhere &#40;the data structure size is O&#40;n&#41; where n is the number of root)

[//]: # (  updates made by the identity&#41;.)

[//]: # ()
[//]: # (## MerkleTree)

[//]: # ()
[//]: # (Not all trees, but:)

[//]: # ()
[//]: # (# Open Questions)

[//]: # ()
[//]: # (## Genesis breaking version rules)

[//]: # ()
[//]: # (What happens when the Genesis Merkle Tree has a version claim at v != 0?  This breaks the transaction rules, which the genesis doesn't follow!)

[//]: # ()
[//]: # (## Temporarily disabling an updatable claim)

[//]: # ()
[//]: # (If an updatable claim needs to be temporarily disabled/unused, it can be upgraded and its value set to all 0.  This could be application-dependent.)

[//]: # ()
[//]: # (Example: A subdomain points to a user identity, but the user doesn't pay for renewal.  The owner of the domain doesn't want the subdomain to point to the user identity, so they make it point to NULL.)

[//]: # ()
[//]: # (---)

[//]: # ()
[//]: # (# Annex)

[//]: # ()
[//]: # (## Types of identity)

[//]: # ()
[//]: # (==Discuss:== Would be defined by a flag in the type field of the ID==)

[//]: # (==The other option is that all the identities are updatable always &#40;does not exist immutable identity&#41;, but sometimes don't use that feature==)

[//]: # ()
[//]: # (- Updatable identity)

[//]: # (	- publishes the root &#40;direct or indirectly&#41; on the blockchain)

[//]: # (	- can add claims after the GenesisId)

[//]: # (	- Before adding/revoking claims post genesis, doesn't need to publish in the blockchain)

[//]: # (- Immutable identity)

[//]: # (	- does not publish the root on the blockchain)

[//]: # (	- can not add claims after the GenesisId)

[//]: # (	- useful for identities that don't issue claims)

[//]: # ()
[//]: # (## Protocol/rules)

[//]: # ()
[//]: # (The `Identity State Transition` must be done following the protocol. This means that all the identities when updating their MerkleTree must follow the protocol, which means that to publish their root is needed to provide a ZKproof that the `root update` is done following the set of rules of the protocol.)

[//]: # ()
[//]: # (The state transition rules are different for regular identities &#40;direct claims&#41; and relay identities &#40;indirect claims&#41;.)

[//]: # ()
[//]: # (In the cases that an identity needs to add thousands of claims &#40;eg:100,000&#41; that are part of a non-updatable set &#40;linked to a timestamp&#41;, as the current state of technology does not allow to compute as much zkproofs in a reasonable time, the identity can create a new sub MerkleTree, and then adding a Claim to the main MerkleTree with the Root of the sub MerkleTree. )

[//]: # (The claims in the sub MerkleTree need to be part of a non-updatable set because the rules of correctness of the sub MerkleTree can not be guaranteed for a possible sub tree update.)

[//]: # ()
[//]: # (![]&#40;https://i.imgur.com/k56Txjl.png&#41;)

[//]: # (## Diagram of trees and proofs)

[//]: # ()
[//]: # (The following diagram shows the merkle trees and calculations involved in minor)

[//]: # (identity state updates that allow the following actions:)

[//]: # (- issue claim)

[//]: # (- revoke claim)

[//]: # ()
[//]: # (The top right section shows the data stored in the identity states smart contract.)

[//]: # ()
[//]: # (On the right, the different kind of credential required data is shown:)

[//]: # (- **A**: Credential Of Existence. Allows proving that a claim existed at some)

[//]: # (  point.)

[//]: # (- **B**: Credential Of Validity. Allows proving that a claim existed and has)

[//]: # (  not been recently revoked.)

[//]: # (- **C**: Credential Input for Zero-Knowledge Proof. Allows proving the)

[//]: # (  validity of a claim in zero-knowledge.)

[//]: # ()
[//]: # (![]&#40;../../imgs/treesAndProofs.png&#41;)