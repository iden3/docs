### State Contract

[**State.sol - Github**](https://github.com/iden3/contracts/blob/master/contracts/State.sol)

The state contract stores the state of each identity operating within Polygon ID. Each identity has an [identifier](../getting-started/identity/identifier.md) and an [identity state](../getting-started/identity/identity-state.md) associated to it. Each identifier and the corresponding identity state are stored inside the [`identities`](https://github.com/iden3/contracts/blob/master/contracts/State.sol#L54) mapping. 

An identity gets updated by executing the [state transition function](../getting-started/state-transition/on-chain-state-transition.md). The State contract verifier the proof on-chain via its [`transitState`](https://github.com/iden3/contracts/blob/master/contracts/State.sol#L87) function.

The State contract provides a timestamp of the changes that occur inside an identity state. No personal information (such as claims) is stored on-chain nor it is inferrable from the information stored on-chain.

Note that the actual proof verification is executed by calling the `verifyProof` function inside the [verifier.sol](https://github.com/iden3/contracts/blob/master/contracts/lib/verifier.sol). 

The `verifier.sol` contract is automatically generated using circom and can be used as a standalone contract to verify the proof. `State.sol` implements further logic once the proof is verified (such as updating the identity state).

State contract addresses:

- [Mumbai: 0x46Fd04eEa588a3EA7e9F055dd691C688c4148ab3](https://mumbai.polygonscan.com/address/0x46Fd04eEa588a3EA7e9F055dd691C688c4148ab3)
- [Polygon Mainnet: 0xb8a86e138C3fe64CbCba9731216B1a638EEc55c8](https://polygonscan.com/address/0xb8a86e138C3fe64CbCba9731216B1a638EEc55c8)
