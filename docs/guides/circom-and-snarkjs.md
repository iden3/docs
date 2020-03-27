# How to use circom and snarkjs

Hello and welcome!

In this guide we'll guide you through the creation of your first [zero-knowledge snark](basics/key-concepts#zk-snarks) circuit using [circom](https://github.com/iden3/circom) and [snarkjs](https://github.com/iden3/snarkjs).

[Circom](https://github.com/iden3/circom) is a library that allows you to build circuits to be used in zero knowledge proofs. 

While [snarkjs](https://github.com/iden3/snarkjs) is an independent implementation of the zk-snarks protocol -- fully written in JavaScript.

Circom is designed to work with snarkjs. In other words, any circuit you build in circom can be used in snarkjs.

We'll start by covering the various techniques to write circuits, then move on to creating and verifying a proof off-chain, and finish off by doing the same thing on-chain on Ethereum.

If you have zero knowledge about zero-knowledge 😋 or are unsure about what a zk-snark is, we recommend you read [this page](basics/key-concepts#zero-knowledge-proofs) first.

To get started follow along from [here](https://github.com/iden3/circom/blob/master/TUTORIAL.md).

