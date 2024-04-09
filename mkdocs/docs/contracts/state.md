### Contracts

[**State.sol - Github**](https://github.com/iden3/contracts/blob/master/contracts/state/State.sol)

The State Contract stores the [Global Identity State Tree](../protocol/spec.md#gist-new). The GIST State represents a snapshot of the states of all the identities operating in the system. The design of the State Contract allows identities to authenticate themselves using [Identity Profiles](../protocol/spec.md#identity-profiles-new)

Every time that an identity is updated, for example when a credential is issued using SMT Proof or revoked, it needs to perform a [State Transition](../getting-started/state-transition/on-chain-state-transition.md). This process consists of generating a zk-proof or a digitally signed message that proves that the identity is authorized to perform the state transition.
Then State contract verifies the proof on-chain via its [transitState](https://github.com/iden3/contracts/blob/master/contracts/state/State.sol) (for zk-proofs) or [transitStateGeneric](https://github.com/iden3/contracts/blob/master/contracts/state/State.sol) (generic as name suggests) function.

Note that the actual zk-proof verification is performed by calling the `verifyProof` function inside the [verifier.sol](https://github.com/iden3/contracts/blob/master/contracts/lib/verifier.sol) from the [`transitState`](https://github.com/iden3/contracts/blob/master/contracts/state/State.sol) function inside the State Contract.

Whenever an identity is updated, the State contract updates the corresponding leaf of the GIST Tree. This process is managed by the [SMTLib](https://github.com/iden3/contracts/blob/master/contracts/lib/SmtLib.sol) which is a Sparse Merkle Tree implementation that manages the GIST Tree and keeps track of its history.

The `verifier.sol` contract is automatically generated using circom and can be used as a standalone contract to verify state transition zk-proof. `State` implements further logic once the proof is verified (such as updating the GIST State).

### State contract addresses

- Polygon Mainnet: [0x624ce98D2d27b20b8f8d521723Df8fC4db71D79D](https://polygonscan.com/address/0x624ce98D2d27b20b8f8d521723Df8fC4db71D79D)
- Polygon Mumbai Testnet: [0x134B1BE34911E39A8397ec6289782989729807a4](https://mumbai.polygonscan.com/address/0x134B1BE34911E39A8397ec6289782989729807a4)
- Polygon Amoy Testnet: [0x1a4cC30f2aA0377b0c3bc9848766D90cb4404124](https://www.oklink.com/amoy/address/0x1a4cc30f2aa0377b0c3bc9848766d90cb4404124)


<br/>

[**IdentityTreeStore.sol - Github**](https://github.com/iden3/contracts/blob/master/contracts/identitytreestore/IdentityTreeStore.sol)

The identity tree store contract is responsible for storing revocation tree nodes and tree roots of Identity. In case
when identity is using onchain [RHS](https://docs.iden3.io/services/rhs/) and [Iden3OnchainSparseMerkleTreeProof2023](https://iden3-communication.io/w3c/status/overview/) credential status.

### IdentityTreeStore contract addresses (On-chain RHS)

- Polygon Mainnet: [0xbEeB6bB53504E8C872023451fd0D23BeF01d320B](https://polygonscan.com/address/0xbEeB6bB53504E8C872023451fd0D23BeF01d320B)
- Polygon Mumbai Testnet: [0x16A1ae4c460C0a42f0a87e69c526c61599B28BC9](https://mumbai.polygonscan.com/address/0x16A1ae4c460C0a42f0a87e69c526c61599B28BC9)
- Polygon Amoy Testnet: [0x3d3763eC0a50CE1AdF83d0b5D99FBE0e3fEB43fb](https://www.oklink.com/amoy/address/0x3d3763ec0a50ce1adf83d0b5d99fbe0e3feb43fb)