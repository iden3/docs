# Generate Proof for State Transition

The output of the previous tutorial was the result of a locally executed computation by the Identity Owner, namely moving from the genesis state to state 1. 

>What if the person that executed the state transition wasn't actually the Identity Owner? What if the key used to sign the message was revoked? 

To ensure the state transition happens in a trustless way, it must be verified inside a circuit. 

The State Transition inputs generated earlier will be passed into the [State Transition Circuit](../../circuits/main-circuits.md#statetransition) to generate a proof of the executed state transition. 

1. **Install [Circom and SnarkJS.](https://docs.circom.io/getting-started/installation/#installing-circom)**

2. **Clone the repository that contains the compiled circuit**

    ```bash
    git clone https://github.com/iden3/tutorial-examples.git
    ```

    This repository contains the `stateTransition` compiled circuit after a trusted setup.

3. **Create a .json file with the state transition inputs from the previous tutorial**

    For this, create a file named `input.json` inside the `.stateTransition/stateTransition_js` and then paste the inputs you generated in the previous tutorial. These inputs will be passed to the circuit and will be used to generate the zk proof.

4. **Generate the proof**

    From the compiled-circuits folder run:

    ```bash 
    ./generate.sh stateTransition
    ```

    If everything worked fine, your terminal should display: 

    ```bash
    [INFO]  snarkJS: OK!
    ```
 
5. **Display the proof**

    You should now have 2 new files inside the /stateTransition/stateTransition_js directory, namely proof.json and public.json:

    - `proof.json` contains the actual proof represented by the three arrays `a, b, and c`. It contains all the raw data of the proof that the SnarkJS library uses for verification of the proof.

    - `public.json` is an array of the four elements representing the public inputs of the circuit. These are `userID,oldUserState,newUserState,isOldStateGenesis`

6. **Export the proof in the Solidity calldata.**

    The two files from the above step can also be exported as Solidity calldata in order to execute the verification on-chain. From the `stateTransition_js` directory run `snarkjs generatecall`.

    ```bash
    snarkjs generatecall
    ```
    Here is what the output would look like: 

    ```bash
    ["0x2b256e25496ac8584bf5714d347821cf9ac8f2472310306033d1ebd4613d12e9", "0x2cca3d40ba395135a38b4ac8c6f8daf81e968ab7082d26d778a82aad9c39d8e3"],
    [["0x2b92b4fc713b659225bfc2b2560b4a1af7901b2a5ee4a3ed07465a88f70e71b3", "0x241ce1ba397c4e1d65059779cacf30fd8d977ed89e6964fa4aa84daec7965254"],["0x27099d3f5cac46fa58c031913c5cd68e24634e9d80281a3d0c0c091bdf574786", "0x08df6f588353293a926660cb1b65a13ad8c5094a42e76dc46d2963ca1cacc096"]],
    ["0x0873f0c6ad05f760775b74a8a6e391beb5b5d3a040a3259f6f5c2429b9d37f8d", "0x15ff3cb9c37c9a07b0fdb2f24cad7bf56adc632c625d9d236841676d731f661b"],
    ["0x00e00c0ee273921f3aa97ba2de5480f140b17e2d35943a8f17a7f45aa04f0000","0x0ee273921f3aa97ba2de5480f140b17e2d35943a8f17a7f45aa04fb715a18685","0x2ba2ba06e0fec5e71fb55019925946590743750a181744fe8eeb8da62e0709db","0x0000000000000000000000000000000000000000000000000000000000000001"]
    ```
The Solidity calldata output represents: 

    - `a[2]`, `b[2][2]`, `c[2]`, namely the proof
    - `public[4]`, namely the public inputs of the circuit 

In the next tutorial, we shall pass this proof to the State.sol smart contract in order to complete the State Transition function.
