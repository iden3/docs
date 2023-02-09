# Main Circuits

This iden3 circuits are the heart of the protocol. The main ones are: 

- [`stateTransition.circom`](main-circuits.md#statetransition), checks the execution of the [identity state transition](../getting-started/state-transition/state-transition.md) by taking the old identity state and the new identity state as inputs.
- [`authV2.circom`](main-circuits.md#authentication-v1), checks that the prover is owner of an identity.
- [`credentialAtomicQueryMTPV2.circom`](main-circuits.md#credentialatomicquerymtp), checks that a claim issued to the prover (and added to issuer's Claims Tree) satisfies a query set by the verifier.
- [`credentialAtomicQueryMTPV2OnChain.circom`](main-circuits.md#credentialatomicquerymtp), checks that a claim issued to the prover (and added to issuer's Claims Tree) satisfies a query set by the verifier and the verifier is a smart contract.
- [`credentialAtomicQuerySig.circom`](main-circuits.md#credentialatomicquerysig) checks that a claim issued to the prover (and signed by the Issuer) satisfies a query set by the verifier.
- [`credentialAtomicQuerySigV2OnChain.circom`](main-circuits.md#credentialatomicquerysig) checks that a claim issued to the prover (and signed by the Issuer) satisfies a query set by the verifier and the verifier is a smart contract.

> You can find all the source code on [Github - Iden3 Circuits](https://github.com/iden3/circuits). All the proving and verification keys necessary to use the circuits were generated after a Trusted Setup Ceremony. Details here:  [Iden3 Protocol Phase2 Trusted Setup Ceremony](https://github.com/iden3/phase2ceremony#auth-circuit)

## stateTransition

- [**Github**](https://github.com/iden3/circuits/blob/master/circuits/lib/stateTransition.circom)

- [**Example of instantiation**](https://github.com/iden3/circuits/blob/master/circuits/stateTransition.circom)

#### Instantiation Parameters

- `idOwnershipLevels` Merkle tree depth level for Identity Trees (claims Tree, revocation Tree and roots Tree)

#### Inputs

| Input                          | Description              | Public or Private
| -----------                    | -----------          |  ----------
| userID                      | Prover's (Genesis) Identifier                | Public
| oldUserState             | Prover's Identity State (before transition)                 | Public
| newUserState     | Prover's Identity State (after transition)           | Public
| isOldStateGenesis               | "1" indicates that the old state is genesis: it means that this is the first State Transition, otherwise "0"              | Public
| claimsTreeRoot                | Prover's Claims Tree Root                | Private
| authClaimMtp[idOwnershipLevels] | Merkle Tree Proof of Auth Claim inside Prover's Claims tree                 | Private
| authClaim[8]    | Prover's Auth Claim                | Private
| revTreeRoot    | Prover's Revocation Tree Root                 | Private
| authClaimNonRevMtp[idOwnershipLevels]    | Merkle Tree Proof of non membership of Auth Claim inside Prover's Revocation Tree                | Private
| authClaimNonRevMtpNoAux              | Flag that indicates whether to check the auxiliary Node                  | Private
| authClaimNonRevMtpAuxHv               | Auxiliary Node Value              | Private
| authClaimNonRevMtpAuxHi          | Auxiliary Node Index           | Private
| rootsTreeRoot          | Prover's Roots Tree Root            | Private
| signatureR8x            | Signature of the challenge (Rx point)           | Private
| signatureR8y            | Signature of the challenge (Ry point)           | Private
| signatureS            | Signature of the challenge (S point)             | Private
| newClaimsTreeRoot            | Claim Tree Root of the Prover after State Transtion is executed             | Private
| newAuthClaimMtp[IdOwnershipLevels];            | Merkle Tree Proof of existance of the Prover's Auth Claim inside the Claims Tree after State Transtion is executed        | Private
| newRevTreeRoot            | Revocation Tree Root of the Prover after State Transtion is executed             | Private
| newRootsTreeRoot            | Roots Tree Root of the Prover after State Transtion is executed             | Private


#### Scope

- If oldState is genesis, verifies that userID is derived from the oldUserState (= genesis state). Performed using [`cutId()`](https://github.com/iden3/circuits/blob/master/circuits/lib/utils/treeUtils.circom#L184), [`cutState()`](https://github.com/iden3/circuits/blob/master/circuits/lib/utils/treeUtils.circom#L198)and [`isEqual()`](https://github.com/iden3/circomlib/blob/master/circuits/comparators.circom#L37) templates
- newUserState is different than zero using [`isZero()`](https://github.com/iden3/circomlib/blob/master/circuits/comparators.circom#L24) comparator
- oldUserState and newUserState are different using `isEqual()`
- Verifies user's identity ownership using [`idOwnershipBySignature(IdOwnershipLevels)`](./main-circuits.md#idownershipbysignature) template. The challenge signed by the user is `H(oldstate, newstate)` where `H` is a Poseidon hash function executed inside the [`Poseidon(nInputs)`](https://github.com/iden3/circomlib/blob/master/circuits/poseidon.circom#L198) template
- Verifies that the auth claim exists in the `newClaimsTreeRoot` using [`checkClaimExists(IdOwnershipLevels)` template](https://github.com/iden3/circuits/blob/master/circuits/lib/stateTransition.circom#L91)
- Verifies that the new state (`newUserState`) matches the hash of the new claims tree root (`newClaimsTreeRoot`), revocation tree root (`newRevTreeRoot`) and roots tree root (`newRootsTreeRoot`) using [`checkIdenStateMatchesRoots()`](https://github.com/iden3/circuits/blob/master/circuits/lib/stateTransition.circom#L96)

#### Circuit Specific Files (From Trusted Setup)

- [Final zkey `circuit_final.zkey`](https://iden3-circuits-bucket.s3.eu-west-1.amazonaws.com/circuits/v0.1.0/auth/circuit_final.zkey)
- [Verification Key `verification_key.json`](https://iden3-circuits-bucket.s3.eu-west-1.amazonaws.com/circuits/v0.1.0/auth/verification_key.json)
- [WASM Witness Generator `circuit.wasm`](https://iden3-circuits-bucket.s3.eu-west-1.amazonaws.com/circuits/v0.1.0/auth/circuit.wasm)

<!-- ## authV1 (Deprecated)

- [**Github**](https://github.com/iden3/circuits/blob/master/circuits/lib/authentication.circom)

- [**Example of instantiation**](https://github.com/iden3/circuits/blob/master/circuits/auth.circom)

#### Instantiation Parameters

- `IdOwnershipLevels` Merkle tree depth level for Claims tree

#### Inputs

| Input                          | Description              | Public or Private
| -----------                    | -----------          |  ----------
| userClaimsTreeRoot                      | Prover's Claims Tree Root                | Private
| userAuthClaimMtp[IdOwnershipLevels]             | Merkle Tree Proof of Auth Claim inside Prover's Claims tree                 | Private
| userAuthClaim[8]     | Prover's Auth Claim           | Private
| userRevTreeRoot    | Prover's Revocation Tree Root                 | Private
| userAuthClaimNonRevMtp[IdOwnershipLevels]               | Merkle Tree Proof of non membership of Auth Claim inside Prover's Revocation Tree              | Private
| userAuthClaimNonRevMtpNoAux                | Flag that indicates whether to check the auxiliary Node                | Private
| userAuthClaimNonRevMtpAuxHv | Auxiliary Node Value                 | Private
| userAuthClaimNonRevMtpAuxHi    | Auxiliary Node Index                | Private
| userRootsTreeRoot    | Prover's Roots Tree Root                 | Private
| challenge    | Message to be signed by the Prover to prove control of an Identity                | Public
| challengeSignatureR8x              | Signature of the challenge (Rx point)                    | Private
| challengeSignatureR8y               | Signature of the challenge (Ry point)                | Private
| challengeSignatureS          |  Signature of the challenge (S point)           | Private
| userState          | Prover's Identity State            | Public
| userID            | Prover's (Genesis) Identifier           | Public

#### Scope

- Prover is owner of an identity by signing a message using [`idOwnershipBySignature` template](./template-circuits.md#idownershipbysignature)
- Contains `userID` as unconstrained input. This is needed as public input as it should be used by the verifier to authenticate the prover. -->

## authV2 

- [**Github**](https://github.com/iden3/circuits/blob/master/circuits/auth/authV2.circom)

- [**Example of instantiation**](https://github.com/iden3/circuits/blob/master/circuits/authV2.circom)

#### Instantiation Parameters

- `IdOwnershipLevels` Merkle tree depth levels for Identity Trees (claims Tree, revocation Tree and roots Tree)
- `onChainLevels` Merkle tree depth of [GIST](./spec.md#GIST-(NEW)) stored on chain

#### Inputs

| Input                          | Description              | Public or Private
| -----------                    | -----------          |  ----------
| genesisID                      | genesis ID of the prover                | Private
| profileNonce                      | Random number, stored by the user              | Private
| state                      | Prover's Identity State                | Private
| claimsTreeRoot                      | Prover's Claims Tree Root                | Private
| revTreeRoot    | Prover's Revocation Tree Root                 | Private
| rootsTreeRoot    | Prover's Roots Tree Root                 | Private
| authClaim[8]     | Prover's Auth Claim           | Private
| authClaimIncMtp[IdOwnershipLevels]             | Merkle Tree Proof of Auth Claim inclusion inside Prover's Claims tree                 | Private
| authClaimNonRevMtp[IdOwnershipLevels]               | Merkle Tree Proof of non inclusion of Auth Claim Nonce inside Prover's Revocation Tree              | Private
| authClaimNonRevMtpNoAux                | Flag that indicates whether to check the auxiliary Node                | Private
| authClaimNonRevMtpAuxHi    | Auxiliary Node Index                | Private
| authClaimNonRevMtpAuxHv | Auxiliary Node Value                 | Private
| challenge    | Message to be signed by the Prover to prove control of an Identity                | Public
| challengeSignatureR8x              | Signature of the challenge (Rx point)                    | Private
| challengeSignatureR8y               | Signature of the challenge (Ry point)                | Private
| challengeSignatureS          |  Signature of the challenge (S point)           | Private
| gistRoot          |  Root of the GIST stored on chain          | Private
| gistMtp[onChainLevels]          |  Merkle Tree Proof of Inclusion of the user state inside the global state          | Private
| gistMtpAuxHi          |  Auxiliary Node Index           | Private
| gistMtpAuxHv          |  Auxiliary Node Value      | Private
| gistMtpNoAux          |  Flag that indicates whether to check the auxiliary Node           | Private

#### Output

| userID          |  Identifier of the user, assigned to H(genesisID, nonce) if nonce != 0, assigned to genesisID if nonce = 0  | Private

#### Scope

- Prover is owner of an identity by signing a message using [`idOwnershipBySignature` template](./template-circuits.md#idownershipbysignature)
- Checks that the user state is included in the [GIST](../protocol/spec.md#gist-new) by using the [SMTVerifier(onChainLevels)](https://github.com/iden3/circuits/blob/master/circuits/auth/authV2.circom#L90)
- Calculate the `userID` as H(genesisID, nonce) if nonce != 0, assigned to genesisID if nonce = 0 as output it. This is the public [Identity Profile](../protocol/spec.md#identity-profiles-new) of the user

## credentialAtomicQueryMTPV2

- [**Github**](https://github.com/iden3/circuits/blob/master/circuits/lib/query/credentialAtomicQueryMTP.circom)

- [**Example of instantiation**](https://github.com/iden3/circuits/blob/master/circuits/offchain/credentialAtomicQueryMTPOffChain.circom)

The circuit takes a query by a verifier and a claim owned by the prover and generate a proof that the claim satisfies the query. In particular, it checks that: 

1. Checks that the prover is owner of an identity by idOwnershipBySignature template
2. Verifies that the identity is the subject of the claim
3. Verifier that the claim is included in the issuer's claim tree
4. Verifies that the claim schema matches the one in the query
5. Verifies that the claim is not revoked by the issuer and is not expired
6. Verifies that the query posed by the verifier is satisfied by the claim

## credentialAtomicQueryMTPV2OnChain 

- [**Github**](https://github.com/iden3/circuits/blob/master/circuits/lib/query/credentialAtomicQueryMTP.circom)

- [**Example of instantiation**](https://github.com/iden3/circuits/blob/master/circuits/credentialAtomicQueryMTPV2OnChain.circom)

> WIP

## credentialAtomicQuerySigV2

- [**Github**](https://github.com/iden3/circuits/blob/master/circuits/onchain/credentialAtomicQueryMTPOnChain.circom)

- [**Example of instantiation**](https://github.com/iden3/circuits/blob/master/circuits/credentialAtomicQuerySigOffChain.circom)

This circuit checks that an issuer has issued a claim for identity and validates ownership of that identity in the following manner:

1. Checks that the prover is owner of an identity by idOwnershipBySignature template
2. Verifies that the identity is the subject of the claim
3. Verifier that the claim was signed by the issuer
4. Verifies that the claim schema matches the one in the query
5. Verifies that the claim is not revoked by the issuer and is not expired
6. Verifies that the query posed by the verifier is satisfied by the claim

## credentialAtomicQueryMTPV2OnChain 

- [**Github**](https://github.com/iden3/circuits/blob/master/circuits/lib/query/credentialAtomicQuerySIG.circom)

- [**Example of instantiation**](https://github.com/iden3/circuits/blob/master/circuits/credentialAtomicQuerySIGV2OnChain.circom)

> WIP





