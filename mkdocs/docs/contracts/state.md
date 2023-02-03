### State Contract

[**StateV2.sol - Github**](https://github.com/iden3/contracts/blob/master/contracts/state/StateV2.sol)

The State Contract stores the [GIST State](../protocol/spec.md#gist-new). The GIST State represents a snapshot of the state of all the identities operating in the system.

Every time that an identity is updated, for example when a claim is added or revoked, it needs to perform a [State Transition](../getting-started/state-transition/on-chain-state-transition.md). The output of the state transition function is a zero knowledge proof that can be verified by the State contract.
The State contract verifies the proof on-chain via its [`transitState`](https://github.com/iden3/contracts/blob/master/contracts/state/StateV2.sol#L148) function.

Note that the actual proof verification is executed by calling the `verifyProof` function inside the [verifier.sol](https://github.com/iden3/contracts/blob/master/contracts/lib/verifier.sol) from the [`transitState`](https://github.com/iden3/contracts/blob/master/contracts/state/StateV2.sol#L196) function inside the State Contract.

Whenever an identity is updated, the State contract [updates the corresponding leaf of the GIST Tree](https://github.com/iden3/contracts/blob/master/contracts/state/StateV2.sol#L214). This process is managed by the [Sparse Merkle Tree (SMT) Contract](https://github.com/iden3/contracts/blob/master/contracts/lib/Smt.sol) which is a Sparse Merkle Tree implementation that manages the GIST Tree and keeps track of its history.

The `verifier.sol` contract is automatically generated using circom and can be used as a standalone contract to verify the proof. `State.sol` implements further logic once the proof is verified (such as updating the GIST State).

State contract addresses:

- [Mumbai: 0x46Fd04eEa588a3EA7e9F055dd691C688c4148ab3](https://mumbai.polygonscan.com/address/0x46Fd04eEa588a3EA7e9F055dd691C688c4148ab3)
- [Polygon Mainnet: 0xb8a86e138C3fe64CbCba9731216B1a638EEc55c8](https://polygonscan.com/address/0xb8a86e138C3fe64CbCba9731216B1a638EEc55c8)
