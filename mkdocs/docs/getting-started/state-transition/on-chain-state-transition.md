# Verify the Proof On-Chain

In order to complete the State Transition process it is necessary to verify the proof inside the `State.sol` contract.

The `transitState` public function of the contract takes the proof generated in the previous section and verifies it on-chain. On verification, the `identities` mapping associated with the `ID` that is executing the transition gets updated.


### Hardhat

   1. **Add Mumbai Network inside your `hardhat.config.js`**

      ```js
      networks: {
         mumbai: {
            url: `${process.env.ALCHEMY_MUMBAI_URL}`,
            accounts: [`0x${process.env.MUMBAI_PRIVATE_KEY}`],
         } 
      ...
      }
      ```

   2. **Add [State.sol](https://github.com/iden3/contracts/blob/master/contracts/State.sol) contract inside the contracts folder**

   3. **Import the state contract from the existing Mumbai testnet address**

      ```js
      const contract = await hre.ethers.getContractAt("State", "0x46Fd04eEa588a3EA7e9F055dd691C688c4148ab3");
      ```

   4. **Add inputs from the proof generated in the previous section**

      ```js
      const id = "0x00e00c0ee273921f3aa97ba2de5480f140b17e2d35943a8f17a7f45aa04f0000"
      const oldState = "0x0ee273921f3aa97ba2de5480f140b17e2d35943a8f17a7f45aa04fb715a18685"
      const newState = "0x2ba2ba06e0fec5e71fb55019925946590743750a181744fe8eeb8da62e0709db"
      const isOldStateGenesis = "0x0000000000000000000000000000000000000000000000000000000000000001"

      const a = ["0x2b256e25496ac8584bf5714d347821cf9ac8f2472310306033d1ebd4613d12e9", "0x2cca3d40ba395135a38b4ac8c6f8daf81e968ab7082d26d778a82aad9c39d8e3"]
      const b = [["0x2b92b4fc713b659225bfc2b2560b4a1af7901b2a5ee4a3ed07465a88f70e71b3", "0x241ce1ba397c4e1d65059779cacf30fd8d977ed89e6964fa4aa84daec7965254"],["0x27099d3f5cac46fa58c031913c5cd68e24634e9d80281a3d0c0c091bdf574786", "0x08df6f588353293a926660cb1b65a13ad8c5094a42e76dc46d2963ca1cacc096"]]
      const c = ["0x0873f0c6ad05f760775b74a8a6e391beb5b5d3a040a3259f6f5c2429b9d37f8d", "0x15ff3cb9c37c9a07b0fdb2f24cad7bf56adc632c625d9d236841676d731f661b"]
      ```

      > Note: Do not use these same inputs for the next section of the tutorial. I already executed the State Transition using these inputs, so the transaction will fail. Instead, use the inputs that you locally generated.

   5. **Fetch identity state before state transition**

      ```js
      let identityState0 = await contract.getState(id);
      // 0
      ```

   6. **Execute state transition function**

      ```js
      await contract.transitState(id, oldState, newState, isOldStateGenesis, a, b, c);
      ```

   7. **Fetch identity state after state transition**

      ```js
      let identityState1 = await contract.getState(id);
      // 19736965623849496899943145128310994086117058864343685620577405145725675178459
      ```



Congratulations! You have successfully completed the identity state transition. 

Starting from the identifier, people will be able to track the status of an identity in a timestamped and tamper-proof way. The identifier remains fixed for the entire existence of an identity, while the identity state changes every time an identity gets updated, for example, when issuing or revoking a claim. As we'll see in the next section, every ZK proof generated from an identity will be checked against the identity state published on-chain.

It is important to underline that:

- The mapping that associates an identifier with its current identity state is the only piece of information stored on-chain. 
- Starting from the identifier and the identity state, it is impossible to retrieve any information stored in the identity trees, for example, reading the content of a claim (which is stored off-chain).
- There is no association between the ECDSA (Elliptical Curve Digital Signature Algorithm) key pair associated with the Ethereum address that executes the State Transition and the Baby Jubjub key pair which is used to control an identity.

> The executable code can be found [here](https://github.com/0xPolygonID/tutorial-examples/tree/main/hardhat-transit-state)