# Main Circuits

This iden3 circuits are the heart of the protocol. The main ones are: 

- [`stateTransition.circom`](main-circuits.md#statetransition), checks the execution of the [identity state transition](../getting-started/state-transition/state-transition.md) by taking the old identity state and the new identity state as inputs.
- [`authV2.circom`](main-circuits.md#authv2), checks that the prover is owner of an identity.
- [`credentialAtomicQueryMTPV2.circom`](./main-circuits.md#credentialatomicquerymtpv2), checks that a claim issued to the prover (added to issuer's Claims Tree) satisfies a query set by the verifier.
- [`credentialAtomicQueryMTPV2OnChain.circom`](./main-circuits.md#credentialatomicquerymtpv2onchain), checks that a claim issued to the prover (added to issuer's Claims Tree) satisfies a query set by the verifier and the verifier is a smart contract.
- [`credentialAtomicQuerySigV2.circom`](./main-circuits.md#credentialatomicquerysigv2) checks that a claim issued to the prover (signed by the Issuer) satisfies a query set by the verifier.
- [`credentialAtomicQuerySigV2OnChain.circom`](./main-circuits.md#credentialatomicquerymtpv2onchain) checks that a claim issued to the prover (signed by the Issuer) satisfies a query set by the verifier and the verifier is a smart contract.

> You can find all the source code on [Github - Iden3 Circuits](https://github.com/iden3/circuits). All the proving and verification keys necessary to use the circuits were generated after a Trusted Setup Ceremony. Details here:  [Iden3 Protocol Phase2 Trusted Setup Ceremony](https://github.com/0xPolygonID/phase2ceremony)

## Circuits that are in beta
- [`credentialAtomicQueryV3.circom`](./main-circuits.md#credentialatomicqueryv3) checks that a claim issued to the prover (signed by the Issuer or included to the Issuer's state) and satisfies a query set by the verifier.
- [`credentialAtomicQueryV3OnChain.circom`](./main-circuits.md#credentialatomicqueryv3onchain) checks that a claim issued to the prover (signed by the Issuer or included to the Issuer's state) and satisfies a query set by the verifier (smart contract). Authentication check inside circuit can be disabled in case Ethereum-based identity authenticates with Ethereum account.


## stateTransition

- [**Github**](https://github.com/iden3/circuits/blob/master/circuits/lib/stateTransition.circom)

- [**Example of instantiation**](https://github.com/iden3/circuits/blob/master/circuits/stateTransition.circom)

- [**Circuit Specific Files (From Trusted Setup)**](https://iden3-circuits-bucket.s3.eu-west-1.amazonaws.com/feature/trusted-setup-v1.0.0.zip)

#### Instantiation Parameters

- `idOwnershipLevels` Merkle tree depth level for Identity Trees (claims Tree, revocation Tree and roots Tree)

#### Inputs

| Input                                 | Description                                                                                                        | Public or Private |
|---------------------------------------|--------------------------------------------------------------------------------------------------------------------|-------------------|
| userID                                | Prover's (Genesis) Identifier                                                                                      | Public            |
| oldUserState                          | Prover's Identity State (before transition)                                                                        | Public            |
| newUserState                          | Prover's Identity State (after transition)                                                                         | Public            |
| isOldStateGenesis                     | "1" indicates that the old state is genesis: it means that this is the first State Transition, otherwise "0"       | Public            |
| claimsTreeRoot                        | Prover's Claims Tree Root                                                                                          | Private           |
| authClaimMtp[idOwnershipLevels]       | Merkle Tree Proof of Auth Claim inside Prover's Claims tree                                                        | Private           |
| authClaim[8]                          | Prover's Auth Claim                                                                                                | Private           |
| revTreeRoot                           | Prover's Revocation Tree Root                                                                                      | Private           |
| authClaimNonRevMtp[idOwnershipLevels] | Merkle Tree Proof of non membership of Auth Claim inside Prover's Revocation Tree                                  | Private           |
| authClaimNonRevMtpNoAux               | Flag that indicates whether to check the auxiliary Node                                                            | Private           |
| authClaimNonRevMtpAuxHv               | Auxiliary Node Value                                                                                               | Private           |
| authClaimNonRevMtpAuxHi               | Auxiliary Node Index                                                                                               | Private           |
| rootsTreeRoot                         | Prover's Roots Tree Root                                                                                           | Private           |
| signatureR8x                          | Signature of the challenge (Rx point)                                                                              | Private           |
| signatureR8y                          | Signature of the challenge (Ry point)                                                                              | Private           |
| signatureS                            | Signature of the challenge (S point)                                                                               | Private           |
| newClaimsTreeRoot                     | Claim Tree Root of the Prover after State Transtion is executed                                                    | Private           |
| newAuthClaimMtp[IdOwnershipLevels];   | Merkle Tree Proof of existance of the Prover's Auth Claim inside the Claims Tree after State Transtion is executed | Private           |
| newRevTreeRoot                        | Revocation Tree Root of the Prover after State Transtion is executed                                               | Private           |
| newRootsTreeRoot                      | Roots Tree Root of the Prover after State Transtion is executed                                                    | Private           |

#### Scope

- If oldState is genesis, verifies that userID is derived from the oldUserState (= genesis state). Performed using [`cutId()`](https://github.com/iden3/circuits/blob/master/circuits/lib/utils/treeUtils.circom#L184), [`cutState()`](https://github.com/iden3/circuits/blob/master/circuits/lib/utils/treeUtils.circom#L198)and [`isEqual()`](https://github.com/iden3/circomlib/blob/master/circuits/comparators.circom#L37) templates
- newUserState is different than zero using [`isZero()`](https://github.com/iden3/circomlib/blob/master/circuits/comparators.circom#L24) comparator
- oldUserState and newUserState are different using `isEqual()`
- Verifies user's identity ownership using [`idOwnershipBySignature(IdOwnershipLevels)`](./main-circuits.md#idownershipbysignature) template. The challenge signed by the user is `H(oldstate, newstate)` where `H` is a Poseidon hash function executed inside the [`Poseidon(nInputs)`](https://github.com/iden3/circomlib/blob/master/circuits/poseidon.circom#L198) template
- Verifies that the auth claim exists in the `newClaimsTreeRoot` using [`checkClaimExists(IdOwnershipLevels)` template](https://github.com/iden3/circuits/blob/master/circuits/lib/stateTransition.circom#L91)
- Verifies that the new state (`newUserState`) matches the hash of the new claims tree root (`newClaimsTreeRoot`), revocation tree root (`newRevTreeRoot`) and roots tree root (`newRootsTreeRoot`) using [`checkIdenStateMatchesRoots()`](https://github.com/iden3/circuits/blob/master/circuits/lib/stateTransition.circom#L96)


## authV2 

- [**Github**](https://github.com/iden3/circuits/blob/master/circuits/auth/authV2.circom)

- [**Example of instantiation**](https://github.com/iden3/circuits/blob/master/circuits/authV2.circom)

- [**Circuit Specific Files (From Trusted Setup)**](https://iden3-circuits-bucket.s3.eu-west-1.amazonaws.com/feature/trusted-setup-v1.0.0.zip)

#### Instantiation Parameters

- `IdOwnershipLevels` Merkle tree depth levels for Identity Trees (claims Tree, revocation Tree and roots Tree)
- `onChainLevels` Merkle tree depth of [GIST](./spec.md#GIST-(NEW)) stored on chain

#### Inputs

| Input                                 | Description                                                                            | Public or Private |
|---------------------------------------|----------------------------------------------------------------------------------------|-------------------|
| genesisID                             | genesis ID of the prover                                                               | Private           |
| profileNonce                          | Random number, stored by the user                                                      | Private           |
| state                                 | Prover's Identity State                                                                | Private           |
| claimsTreeRoot                        | Prover's Claims Tree Root                                                              | Private           |
| revTreeRoot                           | Prover's Revocation Tree Root                                                          | Private           |
| rootsTreeRoot                         | Prover's Roots Tree Root                                                               | Private           |
| authClaim[8]                          | Prover's Auth Claim                                                                    | Private           |
| authClaimIncMtp[IdOwnershipLevels]    | Merkle Tree Proof of Auth Claim inclusion inside Prover's Claims tree                  | Private           |
| authClaimNonRevMtp[IdOwnershipLevels] | Merkle Tree Proof of non inclusion of Auth Claim Nonce inside Prover's Revocation Tree | Private           |
| authClaimNonRevMtpNoAux               | Flag that indicates whether to check the auxiliary Node                                | Private           |
| authClaimNonRevMtpAuxHi               | Auxiliary Node Index                                                                   | Private           |
| authClaimNonRevMtpAuxHv               | Auxiliary Node Value                                                                   | Private           |
| challenge                             | Message to be signed by the Prover to prove control of an Identity                     | Public            |
| challengeSignatureR8x                 | Signature of the challenge (Rx point)                                                  | Private           |
| challengeSignatureR8y                 | Signature of the challenge (Ry point)                                                  | Private           |
| challengeSignatureS                   | Signature of the challenge (S point)                                                   | Private           |
| gistRoot                              | Root of the GIST stored on chain                                                       | Private           |
| gistMtp[onChainLevels]                | Merkle Tree Proof of Inclusion of the user state inside the global state               | Private           |
| gistMtpAuxHi                          | Auxiliary Node Index                                                                   | Private           |
| gistMtpAuxHv                          | Auxiliary Node Value                                                                   | Private           |
| gistMtpNoAux                          | Flag that indicates whether to check the auxiliary Node                                | Private           |

#### Output

| Input  | Description                                                                                               | Public or Private |
|--------|-----------------------------------------------------------------------------------------------------------|-------------------|
| userID | Identifier of the user, assigned to H(genesisID, nonce) if nonce != 0, assigned to genesisID if nonce = 0 | Public            |

#### Scope

- Prover is owner of an identity by signing a message using [`idOwnershipBySignature` template](./template-circuits.md#idownershipbysignature)
- Checks that the user state is included in the [GIST](../protocol/spec.md#gist-new) by using the [SMTVerifier(onChainLevels)](https://github.com/iden3/circuits/blob/master/circuits/auth/authV2.circom#L90)
- Calculate the `userID` as H(genesisID, nonce) if nonce != 0, assigned to genesisID if nonce = 0 as output it. This is the public [Identity Profile](../protocol/spec.md#identity-profiles-new) of the user

## credentialAtomicQueryMTPV2

- [**Github**](https://github.com/iden3/circuits/blob/master/circuits/offchain/credentialAtomicQueryMTPOffChain.circom)

- [**Example of instantiation**](https://github.com/iden3/circuits/blob/master/circuits/credentialAtomicQueryMTPV2.circom)

- [**Circuit Specific Files (From Trusted Setup)**](https://iden3-circuits-bucket.s3.eu-west-1.amazonaws.com/feature/trusted-setup-v1.0.0.zip)


The circuit takes a query by a verifier and a claim owned by the prover and generate a proof that the claim satisfies the query. In particular, it checks that: 

1. Checks that the prover is owner of an identity by idOwnershipBySignature template
2. Verifies that the identity is the subject of the claim
3. Verifier that the claim is included in the issuer's claim tree
4. Verifies that the claim schema matches the one in the query
5. Verifies that the claim is not revoked by the issuer and is not expired
6. Verifies that the query posed by the verifier is satisfied by the claim

## credentialAtomicQueryMTPV2OnChain 

- [**Github**](https://github.com/iden3/circuits/blob/master/circuits/onchain/credentialAtomicQueryMTPOnChain.circom)

- [**Example of instantiation**](https://github.com/iden3/circuits/blob/master/circuits/credentialAtomicQueryMTPV2OnChain.circom)

- [**Circuit Specific Files (From Trusted Setup)**](https://iden3-circuits-bucket.s3.eu-west-1.amazonaws.com/feature/trusted-setup-v1.0.0.zip)


This circuit should be used for smart contract verifiers. This circuits does all the checks that the [credentialAtomicQueryMTPV2](https://github.com/iden3/docs/blob/master/mkdocs/docs/protocol/main-circuits.md#credentialatomicquerymtpv2) circuit does, plus the following:

1. Check that prover controls the identity the same way as the AuthV2 circuit checks it
2. Calculates hash of the query inputs, like claimSchema, slotIndex, operator, claimPathKey, claimPathNotExists and values as an output for all the query related inputs.
This reduces the number of public inputs and much cheaper for Smart Contracts to verify the proof.

## credentialAtomicQuerySigV2

- [**Github**](https://github.com/iden3/circuits/blob/master/circuits/offchain/credentialAtomicQuerySigOffChain.circom)

- [**Example of instantiation**](https://github.com/iden3/circuits/blob/master/circuits/credentialAtomicQuerySigV2.circom)

- [**Circuit Specific Files (From Trusted Setup)**](https://iden3-circuits-bucket.s3.eu-west-1.amazonaws.com/feature/trusted-setup-v1.0.0.zip)


This circuit checks that an issuer has issued a claim for identity and validates ownership of that identity in the following manner:

1. Checks that the prover is owner of an identity by idOwnershipBySignature template
2. Verifies that the identity is the subject of the claim
3. Verifier that the claim was signed by the issuer
4. Verifies that the claim schema matches the one in the query
5. Verifies that the claim is not revoked by the issuer and is not expired
6. Verifies that the query posed by the verifier is satisfied by the claim

## credentialAtomicQuerySigV2OnChain 

- [**Github**](https://github.com/iden3/circuits/blob/master/circuits/onchain/credentialAtomicQuerySigOnChain.circom)

- [**Example of instantiation**](https://github.com/iden3/circuits/blob/master/circuits/credentialAtomicQuerySigV2OnChain.circom)

- [**Circuit Specific Files (From Trusted Setup)**](https://iden3-circuits-bucket.s3.eu-west-1.amazonaws.com/feature/trusted-setup-v1.0.0.zip)


This circuit should be used for smart contract verifiers. This circuits does all the checks that the the [credentialAtomicQuerySigV2](https://github.com/iden3/docs/blob/master/mkdocs/docs/protocol/main-circuits.md#credentialatomicquerysigv2) circuit does, plus the following:

1. Check that prover controls the identity the same way as the AuthV2 circuit checks it
2. Calculates hash of the query inputs, like claimSchema, slotIndex, operator, claimPathKey, claimPathNotExists and values as an output for all the query related inputs.
This reduces the number of public inputs and much cheaper for Smart Contracts to verify the proof.


## credentialAtomicQueryV3

- [**Github**](https://github.com/iden3/circuits/blob/develop/circuits/offchain/credentialAtomicQueryV3OffChain.circom)

- [**Example of instantiation**](https://github.com/iden3/circuits/blob/develop/circuits/credentialAtomicQueryV3.circom)

- [**Circuit Specific Files (version 1.0.0-beta.0, NO Trusted Setup!)**](https://iden3-circuits-bucket.s3.eu-west-1.amazonaws.com/latest.zip)


This circuit checks that an issuer has issued a claim for identity and validates ownership of that identity in the following manner:

1. Verifies that the identity or identity profile is the subject of the credential. 
2. Verifies that the schema in the core claim representation contains a hash of the credential type identifier.
3. Verifies that the credential is not expired.
4. Verifies that the credential is not revoked (in case the revocation check is not skipped).
5. Verifies that the provided issuer state for non-revocation check is built from the provided tree roots (in case the revocation check is not skipped).
6. Depending on the proof of the verifiable credential (Iden3SparseMerkleTreeProof or BJJSignature) determines the proof verification flow and the tree roots to verify.
    1. Verification of BJJSignature Proof:
        1. Verifies that AuthBJJ credential of the issuer (signing key) has a protocol-defined schema hash.
        2. Verifies that AuthBJJ credential of the issuer (signing key) is not revoked by the issuer.
        3. Verifies that the signature is valid and created with a private key corresponding to AuthBJJ credential of the issuer.
        4. Verifies that the core claim representation of AuthBJJ credential is included in the issuer state.
        5. Verifies that the provided issuer state for AuthBJJ issuance check is built from the provided tree roots.
    2. Verification of Iden3SparseMerkleTreeProof:
        1. Verifies that the core claim representation of the user credential is included in the issuer state.
        2. Verifies that the provided issuer state for issuance check is built from the provided tree roots.
7. Verifies query:
    1. Verifies that the credential field is a part of the merklized root from core claim representation (in case schema is for merklized credential).
    2. Verifies that the credential field is located at the expected data slot of core claim representation (in case schema is for non-merklized credential).
    3. Verifies that credential data satisfies the query condition. 
8. Calculates nullifier in case nullifier session id and verifierID are set and credential has been issued to the user profile.
9. Outputs the field value in case selective disclosure is requested.
10. Generates user profile in case profile nonce is set.
11. Calculates link id in case links session id is set.

## credentialAtomicQueryV3Onchain

- [**Github**](https://github.com/iden3/circuits/blob/develop/circuits/onchain/credentialAtomicQueryV3OnChain.circom)

- [**Example of instantiation**](https://github.com/iden3/circuits/blob/develop/circuits/credentialAtomicQueryV3OnChain.circom)

- [**Circuit Specific Files (version 1.0.0-beta.0, NO Trusted Setup!)**](https://iden3-circuits-bucket.s3.eu-west-1.amazonaws.com/latest.zip)


This circuit should be used for smart contract verifiers. This circuit does all the checks that the credentialAtomicQueryV3 circuit does, plus the following:

1. Checks that the prover controls the identity in the same way AuthV2 circuit checks it if auth is enabled.
2. Verifies credential query in the same way as credentialAtomicQueryV3 does.
3. Calculates hash of the query inputs, like claimSchema, slotIndex, operator, claimPathKey, claimPathNotExists, and values as an output for all the query-related inputs.
   This reduces the number of public inputs and makes it much cheaper for Smart Contracts to verify the proof.
