# Glossary

## Rollup
I don't think we can do much better than Ed Felton's definition:

> Rollup is a general approach to scaling open contracts, that is, contracts that everyone can see and interact with. In rollup, calls to the contract and their arguments are written on-chain as calldata, but the actual computation and storage of the contract are done off-chain. Somebody posts an on-chain assertion about what the contract will do. You can think of the assertion as "rolling up" all of the calls and their results into a single on-chain transaction. Where rollup systems differ is in how they ensure that the assertions are correct.

## zkSnark
A zkSnark is a short (and efficiently checkable) cryptographic proof that allow us to prove something specific without revealing any extra information.

## zkRollup
ZkRollup relies on zkSnarks to prove that on-chain assertions are correct. In other words, every assertion is accompanied by an easily verifiable proof which proves that the computations and data described in the assertion are correct.

## BabyJubJub
BabyJubJub is an elliptic curve defined over a large prime sub-group. It's useful in zkSnark proofs.

## Batch
A batch is a rollup block.

## Operator
An operator is a rollup block producer.

## Forging
Forging refers to the creation of a batch (off-chain) and the subsequent (on-chain) verification of the attached zkSnark.
