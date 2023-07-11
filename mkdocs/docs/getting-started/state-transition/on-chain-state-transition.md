# Verify the Proof On-Chain

In order to complete the State Transition process it is necessary to verify the proof inside the `StateV2.sol` contract.

The `transitState` public function of the contract takes the proof generated in the previous section and verifies it on-chain. On verification, the `identities` mapping associated with the `ID` that is executing the transition gets updated.

### Hardhat

1.**Add Mumbai Network inside your `hardhat.config.js`**

```js
networks: {
   mumbai: {
      url: `${process.env.MUMBAI_RPC_URL}`,
      accounts: [`${process.env.MUMBAI_PRIVATE_KEY}`]
   }
}
```

2.**Add [StateV2.sol](https://github.com/iden3/contracts/blob/master/contracts/state/StateV2.sol#L148) contract and its dependencies ([Poseidon.sol](https://github.com/iden3/contracts/blob/master/contracts/lib/Poseidon.sol) and [Smt.sol](https://github.com/iden3/contracts/blob/master/contracts/lib/Smt.sol)) inside the contracts folder**

3.**Import the state contract from the existing Mumbai testnet address**

```js
const contract = await hre.ethers.getContractAt("StateV2", "0xEA9aF2088B4a9770fC32A12fD42E61BDD317E655");
```

4.**Add inputs from the proof generated in the previous section**

```js
   const id = "0x000a501c057d28c0c50f91062730531a247474274ff6204a4f7da6d4bcb70000"
  const oldState = "0x1c057d28c0c50f91062730531a247474274ff6204a4f7da6d4bcb7d23be4d605"
  const newState = "0x203034fdafe4563e84962f2b16fefe8ebedb1be5c05b7d5e5e30898d799192fd"
  const isOldStateGenesis = "0x0000000000000000000000000000000000000000000000000000000000000001"

  const a = ["0x0c98dbb5bcdc4810a976b9804972c6086e855532740ab2c611fbcf4a5d939f91", "0x1f3b6aa1cfe69a2a3f5e8e7db5ccae0d269fc66be6d0c364469486d5718431ee"]
  const b = [["0x21f67821a25f3b0eb008e8aa840706c6dd9c1cff16ec6f138d7745aff350dbbb", "0x255b9f12a90b1f1089af5edcda19fb6d592096f6ba7ce2438ce4ecc48399687d"],["0x1568f9a5a84d72a31b90d26b5035030b0b02544dcba18f0a3740f80b9632942d", "0x28dcba6dd58878a3383fd556d27118a3e905f424d23afa30b71de3ac000822de"]]
  const c = ["0x15adbb5f1abe4418a7ea7f876164b57bf70f88183fa7d85406be4cb5f8fee261", "0x04466d6e7131a89fdcf5136b52ed2b52e00755ad77c97bb87e8afa690eeef5e4"]
```

> Note: Do not use these same inputs for the next section of the tutorial. I already executed the State Transition using these inputs, so the transaction will fail. Instead, use the inputs that you locally generated.

5.**Execute state transition function**

```js
await contract.transitState(id, oldState, newState, isOldStateGenesis, a, b, c);
```

6.**Fetch identity state after state transition**

```js
// Get state of identity without BigNumber
let identityState = await contract.getStateInfoById(id);

console.log("Identity State after state transition", identityState.id);
// 18221365812082731933036101625854358571814024255404073202903829924181114880
```

Congratulations! You have successfully completed the identity state transition. 

Starting from the identifier, people will be able to track the status of an identity in a timestamped and tamper-proof way. The identifier remains fixed for the entire existence of an identity, while the identity state changes every time an identity gets updated, for example, when issuing or revoking a claim. As we'll see in the next section, every ZK proof generated from an identity will be checked against the identity state published on-chain.

It is important to underline that:

- The mapping that associates an identifier with its current identity state is the only piece of information stored on-chain. 
- Starting from the identifier and the identity state, it is impossible to retrieve any information stored in the identity trees, for example, reading the content of a claim (which is stored off-chain).
- There is no association between the ECDSA (Elliptical Curve Digital Signature Algorithm) key pair associated with the Ethereum address that executes the State Transition and the Baby Jubjub key pair which is used to control an identity.

> The executable code can be found [here](https://github.com/0xPolygonID/tutorial-examples/tree/main/hardhat-transit-state)
