### Remix

   1. Connect yout Metamask wallet to Polygon Mumbai Testnet.

   2. On the [Remix Homepage](https://remix.ethereum.org/), click "Load from GitHub" and import `State.sol` using the link: https://github.com/iden3/contracts/blob/master/contracts/State.sol

   3. Move to the "Solidity Compiler" section and compile `State.sol`.

   4. Move to the "Deploy and Run Transaction" section and modify the Environment to "Injected web3". If everything was set correctly, you should see `Custom (80001) network` below the environment drop-down menu.The system prompts you to connect to your MetaMask wallet. Make sure to select the "Mumbai" network on your Metamask before connecting the wallet. 

   5. Make sure that the State contract is selected in the contract drop-down menu and "Load contract from address" adding **0x46Fd04eEa588a3EA7e9F055dd691C688c4148ab3** as contract address.

   6. Check identity state at T=0. To check the identity state call the getState function on the State.sol passing in your identifier. The identifier is the first public input in the public array returned from the solidity calldata from the previous tutorial. The result is zero as there's no identity state associated with that identifier because the identity state has never been published on-chain (yet!) 

   7. Now update the identity state by calling the `transitState` function on State.sol. 

      The outputs generated from the previous tutorial are passed as inputs to the `transitState` function. See the one-to-one mapping between the outputs from state transition and the inputs to the `transitState` function in the diagram below:

      <div align="center">
      <img src= "../../../imgs/transitState-input-remix.png" align="center" width="400"/>
      <img src= "../../../imgs/inputs-to-transitState-function.png" align="center" width="400"/>
      <div align="center"><span style="font-size: 14px;"><b> transitState Function Inputs </b></div>
      </div>

   8. Check the new state. To check, call the `getState` function again by passing the value of the identifier you used above as an input to the `transitState` function. 

   You can see that the console displays a new state:

   `uint256:14531895531750268543323474544059484523319511522635242711319115705040584883009`
