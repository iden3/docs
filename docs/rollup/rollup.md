# zkRollup

## Introduction
[zkRollup](https://github.com/barryWhiteHat/roll_up) is a project to scale ethereum with zkSnarks.

In the words of [Vitalik](https://github.com/barryWhiteHat/roll_up):

>We can actually scale asset transfer transactions on ethereum by a huge amount, without using layer 2’s that introduce liveness assumptions (eg. channels, plasma), by using ZK-SNARKs to mass-validate transactions. 

By some estimates, rollup could help ethereum scale to [3000 transactions-per-second](https://blog.iden3.io/istanbul-zkrollup-ethereum-throughput-limits-analysis.html). To put this in context, the ethereum blockchain currently supports roughly 15 tps. And 2000 tps is what the Visa network currently averages.

For a more complete introduction, see this [video](https://www.youtube.com/watch?v=TtsDNneTDDY).

## Our implementation

For our implementation using [circom](https://github.com/iden3/circom) and [circomlib](https://github.com/iden3/circomlib) see [here](https://github.com/iden3/rollup).

>Note: this is very much a work in progress.

## Definitions

### Rollup
I don't think we can do much better than [Ed Felton's](https://medium.com/offchainlabs/whats-up-with-rollup-db8cd93b314e) definition :

>Rollup is a general approach to scaling open contracts, that is, contracts that everyone can see and interact with. In rollup, calls to the contract and their arguments are written on-chain as calldata, but the actual computation and storage of the contract are done off-chain. Somebody posts an on-chain assertion about what the contract will do. You can think of the assertion as "rolling up" all of the calls and their results into a single on-chain transaction. Where rollup systems differ is in how they ensure that the assertions are correct.

### zkSnark
A zkSnark is a short (and efficiently checkable) cryptographic proof that allow us to prove something specific without revealing any extra information.

### zkRollup
ZkRollup is a specific type of rollup that relies on zkSnarks to prove that on-chain assertions are correct. In other words, every assertion is accompanied by an easily verifiable proof which proves that the computations and data described in the assertion are correct.

### BabyJubJub
BabyJubJub is an elliptic curve defined over a large prime sub-group. It's useful in zkSnark proofs.
A batch is a rollup block.

### Operator
An operator is a rollup block producer.

### Forging
Forging refers to the creation of a batch (off-chain) and the subsequent (on-chain) verification of the attached zkSnark.
