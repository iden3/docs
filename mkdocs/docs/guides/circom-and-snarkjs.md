# How to Use Circom and SnarkJS?

Hello and welcome!!

In this guide, we'll explain how to create your first [zero-knowledge Snark](../../basics/key-concepts#zk-snarks) circuit using [Circom](https://docs.circom.io) and [SnarkJS](https://github.com/iden3/snarkjs).

[Circom](https://circom.iden3.io) is a library that allows you to build circuits to be used in zero-knowledge proofs. 

While [SnarkJS](https://github.com/iden3/snarkjs) is an independent implementation of the zk-SNARK protocol (fully written in JavaScript), Circom is designed to work with SnarkJS. In other words, any circuit you build in Circom can be used in SnarkJS.

We'll start by covering various techniques to write circuits; then we shall move on to creating and verifying a proof off-chain, and finally, finish it off by repeating this process on-chain (on Ethereum).

If you have zero knowledge about zero-knowledge 😋 or are unsure about what a zk-SNARK is, we recommend you read [this page](../../basics/key-concepts) first.

To get started, click [here](https://docs.circom.io/getting-started/installation/).

