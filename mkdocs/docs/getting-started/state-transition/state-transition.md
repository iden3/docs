# Intro to State Transtion

When an identity adds a new claim to her Claims Tree, the root of the tree and, consequently, the identity state change. The process of moving from one state to another is defined [**State Transition**](https://docs.iden3.io/protocol/spec/#identity-state-transition-function).

The State Transtion is executed inside a circuit. The `stateTransition` [circuit](../../circuits/main-circuits.md#statetransition) encodes a set of rules that must be respected to complete the state transition such as:

- The prover is the owner of the identity (checked using a digital signature by the private key corresponding the `authClaim`)
- The `authClaim` of the prover hasn't been revoked.
 
The identity state gets updated by calling the `transitState` [smart contract function](https://github.com/iden3/contracts/blob/master/contracts/State.sol#L87.). To call this function, it is necessary to pass in the proof generated previously.

On verification, the `identities` mapping gets updated associating the `ID` with a new `IdS`.

This tutorial is split in 3 parts:

1. [Add Claim to the Claims Tree](./new-identity-state.md)
2. [Generate Proof for State Transition](./state-transition-proof.md)
3. [Verify the Proof On-Chain](./on-chain-state-transition.md)

> Note: The Identity State Transition happens not only when an identity adds a new claim to the Claims Tree, but also when a claim gets updated or revoked (by adding it to the revocation tree).