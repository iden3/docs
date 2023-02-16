# Generate Proof for State Transition

The output of the previous tutorial was the result of a locally executed computation by the Identity Owner, namely moving from the genesis state to state 1. 

>What if the person that executed the state transition wasn't actually the Identity Owner? What if the key used to sign the message was revoked? 

To ensure the state transition happens in a trustless way, it must be verified inside a circuit. 

The State Transition inputs generated earlier will be passed into the [State Transition Circuit](../../protocol/main-circuits.md#statetransition) to generate a proof of the executed state transition. 

1.**Install [Circom and SnarkJS.](https://docs.circom.io/getting-started/installation/#installing-circom)**

2.**Clone the repository that contains the compiled circuit**

```bash
git clone https://github.com/iden3/tutorial-examples.git
```

This repository contains the `stateTransition` compiled circuit after a trusted setup.

3.**Create a .json file with the state transition inputs from the previous tutorial**

For this, create a file named `input.json` inside the `.stateTransition/stateTransition_js` and then paste the inputs you generated in the previous tutorial. These inputs will be passed to the circuit and will be used to generate the zk proof.

4.**Generate the proof**

From the compiled-circuits folder run:

```bash 
./generate.sh stateTransition
```

If everything worked fine, your terminal should display: 

```bash
[INFO]  snarkJS: OK!
```
 
5.**Display the proof**

You should now have 2 new files inside the /stateTransition/stateTransition_js directory, namely proof.json and public.json:

- `proof.json` contains the actual proof represented by the three arrays `a, b, and c`. It contains all the raw data of the proof that the SnarkJS library uses for verification of the proof.

- `public.json` is an array of the four elements representing the public inputs of the circuit. These are `userID,oldUserState,newUserState,isOldStateGenesis`

6.**Export the proof in the Solidity calldata.**

The two files from the above step can also be exported as Solidity calldata in order to execute the verification on-chain. From the `stateTransition_js` directory run `snarkjs generatecall`.

```bash
snarkjs generatecall
```
Here is what the output would look like: 

```bash
["0x0c98dbb5bcdc4810a976b9804972c6086e855532740ab2c611fbcf4a5d939f91", "0x1f3b6aa1cfe69a2a3f5e8e7db5ccae0d269fc66be6d0c364469486d5718431ee"],[["0x21f67821a25f3b0eb008e8aa840706c6dd9c1cff16ec6f138d7745aff350dbbb", "0x255b9f12a90b1f1089af5edcda19fb6d592096f6ba7ce2438ce4ecc48399687d"],["0x1568f9a5a84d72a31b90d26b5035030b0b02544dcba18f0a3740f80b9632942d", "0x28dcba6dd58878a3383fd556d27118a3e905f424d23afa30b71de3ac000822de"]],["0x15adbb5f1abe4418a7ea7f876164b57bf70f88183fa7d85406be4cb5f8fee261", "0x04466d6e7131a89fdcf5136b52ed2b52e00755ad77c97bb87e8afa690eeef5e4"],["0x000a501c057d28c0c50f91062730531a247474274ff6204a4f7da6d4bcb70000","0x1c057d28c0c50f91062730531a247474274ff6204a4f7da6d4bcb7d23be4d605","0x203034fdafe4563e84962f2b16fefe8ebedb1be5c05b7d5e5e30898d799192fd","0x0000000000000000000000000000000000000000000000000000000000000001"]
```

The Solidity calldata output represents: 

    - `a[2]`, `b[2][2]`, `c[2]`, namely the proof
    - `public[4]`, namely the public inputs of the circuit 

In the next tutorial, we shall pass this proof to the State.sol smart contract in order to complete the State Transition function.
