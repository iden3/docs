[**State.sol - Github**](https://github.com/iden3/contracts/blob/master/contracts/state/State.sol)

The State Contract stores the [Global Identity State Tree](../protocol/spec.md#gist-new). The GIST State represents a snapshot of the states of all the identities operating in the system. The design of the State Contract allows identities to authenticate themselves using [Identity Profiles](../protocol/spec.md#identity-profiles-new)

Every time that an identity is updated, for example when a credential is issued using SMT Proof or revoked, it needs to perform a [State Transition](../getting-started/state-transition/on-chain-state-transition.md). This process consists of generating a zk-proof or a digitally signed message that proves that the identity is authorized to perform the state transition.
Then State contract verifies the proof on-chain via its [transitState](https://github.com/iden3/contracts/blob/master/contracts/state/State.sol) (for zk-proofs) or [transitStateGeneric](https://github.com/iden3/contracts/blob/master/contracts/state/State.sol) (generic as name suggests) function.
Note that the actual zk-proof verification is performed by calling the `verifyProof` function inside the [verifier.sol](https://github.com/iden3/contracts/blob/master/contracts/lib/verifier.sol) from the [`transitState`](https://github.com/iden3/contracts/blob/master/contracts/state/State.sol) function inside the State Contract.

Whenever an identity is updated, the State contract updates the corresponding leaf of the GIST Tree. This process is managed by the [SMTLib](https://github.com/iden3/contracts/blob/master/contracts/lib/SmtLib.sol) which is a Sparse Merkle Tree implementation that manages the GIST Tree and keeps track of its history.

The `verifier.sol` contract is automatically generated using circom and can be used as a standalone contract to verify state transition zk-proof. `State` implements further logic once the proof is verified (such as updating the GIST State).

### State contract addresses

- Ethereum: [0x3C9acB2205Aa72A05F6D77d708b5Cf85FCa3a896](https://etherscan.io/address/0x3C9acB2205Aa72A05F6D77d708b5Cf85FCa3a896)
- Ethereum Sepolia: [0x3C9acB2205Aa72A05F6D77d708b5Cf85FCa3a896](https://sepolia.etherscan.io/address/0x3C9acB2205Aa72A05F6D77d708b5Cf85FCa3a896)
- Polygon Mainnet: [0x624ce98D2d27b20b8f8d521723Df8fC4db71D79D](https://polygonscan.com/address/0x624ce98D2d27b20b8f8d521723Df8fC4db71D79D)
- Polygon Amoy Testnet: [0x1a4cC30f2aA0377b0c3bc9848766D90cb4404124](https://www.oklink.com/amoy/address/0x1a4cc30f2aa0377b0c3bc9848766d90cb4404124)
- Polygon zkEVM: [0x3C9acB2205Aa72A05F6D77d708b5Cf85FCa3a896](https://zkevm.polygonscan.com/address/0x3C9acB2205Aa72A05F6D77d708b5Cf85FCa3a896)
- Polygon zkEVM Cardona: [0x3C9acB2205Aa72A05F6D77d708b5Cf85FCa3a896](https://cardona-zkevm.polygonscan.com/address/0x3C9acB2205Aa72A05F6D77d708b5Cf85FCa3a896)
- Linea: [0x3C9acB2205Aa72A05F6D77d708b5Cf85FCa3a896](https://lineascan.build/address/0x3C9acB2205Aa72A05F6D77d708b5Cf85FCa3a896)
- Linea-Sepolia: [0x3C9acB2205Aa72A05F6D77d708b5Cf85FCa3a89](https://sepolia.lineascan.build/address/0x3C9acB2205Aa72A05F6D77d708b5Cf85FCa3a896)

<br/>
[**IdentityTreeStore.sol - Github**](https://github.com/iden3/contracts/blob/master/contracts/identitytreestore/IdentityTreeStore.sol)

The identity tree store contract is responsible for storing revocation and roots tree nodes of Identity. In case
when identity is using onchain [RHS](https://docs.iden3.io/services/rhs/) and [Iden3OnchainSparseMerkleTreeProof2023](https://iden3-communication.io/w3c/status/overview/) credential status.

### IdentityTreeStore contract addresses (On-chain RHS)

- Ethereum: [0x7dF78ED37d0B39Ffb6d4D527Bb1865Bf85B60f81](https://etherscan.io/address/0x7dF78ED37d0B39Ffb6d4D527Bb1865Bf85B60f81)
- Ethereum Sepolia: [0x7dF78ED37d0B39Ffb6d4D527Bb1865Bf85B60f81](https://sepolia.etherscan.io/address/0x7dF78ED37d0B39Ffb6d4D527Bb1865Bf85B60f81)
- Polygon Mainnet: [0x7dF78ED37d0B39Ffb6d4D527Bb1865Bf85B60f81](https://polygonscan.com/address/0x7dF78ED37d0B39Ffb6d4D527Bb1865Bf85B60f81)
- Polygon Amoy Testnet: [0x7dF78ED37d0B39Ffb6d4D527Bb1865Bf85B60f81](https://www.oklink.com/amoy/address/0x7dF78ED37d0B39Ffb6d4D527Bb1865Bf85B60f81)
- Polygon zkEVM: [0x7dF78ED37d0B39Ffb6d4D527Bb1865Bf85B60f81](https://zkevm.polygonscan.com/address/0x7dF78ED37d0B39Ffb6d4D527Bb1865Bf85B60f81)
- Polygon zkEVM Cardona: [0x7dF78ED37d0B39Ffb6d4D527Bb1865Bf85B60f81](https://cardona-zkevm.polygonscan.com/address/0x7dF78ED37d0B39Ffb6d4D527Bb1865Bf85B60f81)
- Linea: [0x7dF78ED37d0B39Ffb6d4D527Bb1865Bf85B60f81](https://lineascan.build/address/0x7dF78ED37d0B39Ffb6d4D527Bb1865Bf85B60f81)
- Linea-Sepolia: [0x7dF78ED37d0B39Ffb6d4D527Bb1865Bf85B60f81](https://sepolia.lineascan.build/address/0x7dF78ED37d0B39Ffb6d4D527Bb1865Bf85B60f81)

### VALIDATOR_MTP_V2 contract addresses

- Ethereum: [0x27bDFFCeC5478a648f89764E22fE415486A42Ede](https://etherscan.io/address/0x27bDFFCeC5478a648f89764E22fE415486A42Ede)
- Ethereum Sepolia: [0x27bDFFCeC5478a648f89764E22fE415486A42Ede](https://sepolia.etherscan.io/address/0x27bDFFCeC5478a648f89764E22fE415486A42Ede)
- Polygon Mainnet: [0x27bDFFCeC5478a648f89764E22fE415486A42Ede](https://polygonscan.com/address/0x27bDFFCeC5478a648f89764E22fE415486A42Ede)
- Polygon Amoy Testnet: [0x27bDFFCeC5478a648f89764E22fE415486A42Ede](https://www.oklink.com/amoy/address/0x27bDFFCeC5478a648f89764E22fE415486A42Ede)
- Polygon zkEVM: [0x27bDFFCeC5478a648f89764E22fE415486A42Ede](https://zkevm.polygonscan.com/address/0x27bDFFCeC5478a648f89764E22fE415486A42Ede)
- Polygon zkEVM Cardona: [0x27bDFFCeC5478a648f89764E22fE415486A42Ede](https://cardona-zkevm.polygonscan.com/address/0x27bDFFCeC5478a648f89764E22fE415486A42Ede)
- Linea: [0x27bDFFCeC5478a648f89764E22fE415486A42Ede](https://lineascan.build/address/0x27bDFFCeC5478a648f89764E22fE415486A42Ede)
- Linea-Sepolia: [0x27bDFFCeC5478a648f89764E22fE415486A42Ede](https://sepolia.lineascan.build/address/0x27bDFFCeC5478a648f89764E22fE415486A42Ede)

### VALIDATOR_SIG_V2 contract addresses

- Ethereum: [0x59B347f0D3dd4B98cc2E056Ee6C53ABF14F8581b](https://etherscan.io/address/0x59B347f0D3dd4B98cc2E056Ee6C53ABF14F8581b)
- Ethereum Sepolia: [0x59B347f0D3dd4B98cc2E056Ee6C53ABF14F8581b](https://sepolia.etherscan.io/address/0x59B347f0D3dd4B98cc2E056Ee6C53ABF14F8581b)
- Polygon Mainnet: [0x59B347f0D3dd4B98cc2E056Ee6C53ABF14F8581b](https://polygonscan.com/address/0x59B347f0D3dd4B98cc2E056Ee6C53ABF14F8581b)
- Polygon Amoy Testnet: [0x59B347f0D3dd4B98cc2E056Ee6C53ABF14F8581b](https://www.oklink.com/amoy/address/0x59B347f0D3dd4B98cc2E056Ee6C53ABF14F8581b)
- Polygon zkEVM: [0x59B347f0D3dd4B98cc2E056Ee6C53ABF14F8581b](https://zkevm.polygonscan.com/address/0x59B347f0D3dd4B98cc2E056Ee6C53ABF14F8581b)
- Polygon zkEVM Cardona: [0x59B347f0D3dd4B98cc2E056Ee6C53ABF14F8581b](https://cardona-zkevm.polygonscan.com/address/0x59B347f0D3dd4B98cc2E056Ee6C53ABF14F8581b)
- Linea: [0x59B347f0D3dd4B98cc2E056Ee6C53ABF14F8581b](https://lineascan.build/address/0x59B347f0D3dd4B98cc2E056Ee6C53ABF14F8581b)
- Linea-Sepolia: [0x59B347f0D3dd4B98cc2E056Ee6C53ABF14F8581b](https://sepolia.lineascan.build/address/0x59B347f0D3dd4B98cc2E056Ee6C53ABF14F8581b)

### VALIDATOR_V3 contract addresses

- Ethereum: [0xd179f29d00Cd0E8978eb6eB847CaCF9E2A956336](https://etherscan.io/address/0xd179f29d00Cd0E8978eb6eB847CaCF9E2A956336)
- Ethereum Sepolia: [0xd179f29d00Cd0E8978eb6eB847CaCF9E2A956336](https://sepolia.etherscan.io/address/0xd179f29d00Cd0E8978eb6eB847CaCF9E2A956336)
- Polygon Mainnet: [0xd179f29d00Cd0E8978eb6eB847CaCF9E2A956336](https://polygonscan.com/address/0xd179f29d00Cd0E8978eb6eB847CaCF9E2A956336)
- Polygon Amoy Testnet: [0xd179f29d00Cd0E8978eb6eB847CaCF9E2A956336](https://www.oklink.com/amoy/address/0xd179f29d00Cd0E8978eb6eB847CaCF9E2A956336)
- Polygon zkEVM: [0xd179f29d00Cd0E8978eb6eB847CaCF9E2A956336](https://zkevm.polygonscan.com/address/0xd179f29d00Cd0E8978eb6eB847CaCF9E2A956336)
- Polygon zkEVM Cardona: [0xd179f29d00Cd0E8978eb6eB847CaCF9E2A956336](https://cardona-zkevm.polygonscan.com/address/0xd179f29d00Cd0E8978eb6eB847CaCF9E2A956336)
- Linea: [0xd179f29d00Cd0E8978eb6eB847CaCF9E2A956336](https://lineascan.build/address/0xd179f29d00Cd0E8978eb6eB847CaCF9E2A956336)
- Linea-Sepolia: [0xd179f29d00Cd0E8978eb6eB847CaCF9E2A956336](https://sepolia.lineascan.build/address/0xd179f29d00Cd0E8978eb6eB847CaCF9E2A956336)