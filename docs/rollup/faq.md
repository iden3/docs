# Frequently asked questions

## What is zkRollup?

In a nutshell, zkRollup is a layer 2 construction - similar to Plasma - which uses the Ethereum blockchain for data storage instead of computation. In other words, zkRollup does computation off-chain and handles data availability on-chain.

All funds are held by a smart contract on the main-chain. For every batch, a zkSnark is generated off-chain and verified by this contract.

This snark proves the validity of every transaction in the batch.

This means that instead of relying on the Ethereum main-chain to verify each signature transaction, we just need to verify the snark to prove the validity of the off-chain transactions.

The beauty here is that this verification can be done in constant time. In other words, verification time doesn't depend on the number of transactions! This ability to verify proofs both efficiently and in constant time is at the heart of zkRollup.

In addition to this, all transaction data is published cheaply on-chain, without signatures - under [calldata](https://ethereum.stackexchange.com/a/52992). Since the data is published on-chain, we get around the [data availability problems](https://github.com/ethereum/research/wiki/A-note-on-data-availability-and-erasure-coding) that have plagued other L2 solutions such as Plasma.

Importantly, anyone can reconstruct the current state and history from this on-chain data. This prevents censorship and avoids the centralization of operators (rollup block producers) - since anyone can build the state tree from scratch (and therefore become an operator).

A [post](https://medium.com/matter-labs/introducing-zk-sync-the-missing-link-to-mass-adoption-of-ethereum-14c9cea83f58) by MatterLabs, does a good job at expressing the guarantees that this architecture provides:

> 1. The Rollup validator(s) [(what we call an operator)] can never corrupt the state or steal funds (unlike Sidechains).
> 2. Users can always retrieve the funds from the Rollup even if validator(s) [operators] stop cooperating because the data is available (unlike Plasma).
> 3. Thanks to validity proofs [zkSnarks], neither users nor a single trusted third party needs to be online to monitor Rollup blocks in order to prevent fraud (unlike fraud-proof systems, such as payment channels or optimistic rollups). This excellent article dives deep into the overwhelming benefits of validity proofs over fraud proofs.

## Why do we need zkRollup?

Trust-minimised blockchain scaling mechanisms are sorely needed if blockchain applications are ever to achieve mass adoption.

For context, the Ethereum network can handle approximately 15 transactions per second (tps), while the Visa network averages around 2,000 tps.

As outlined in [a post we published last year](https://iden3.io/post/istanbul-zkrollup-ethereum-throughput-limits-analysis), zkRollup has the potential to increase the Ethereum network's maximum tps by two orders of magnitude, making it comparable to the Visa network's average.

![](https://cdn-images-1.medium.com/max/800/1*l7P0QmjAxUolC1r4g_5LSA.png)

## How is 2000 tps possible?

In sum, zkRollup uses zkSnarks to scale Ethereum by taking advantage of the succinctness provided by snarks.

We improve blockchain scalability by compressing each transaction to ~10 bytes:instead of including signatures on-chain, we send a zkSnark which proves that 1000's of signature verifications and other transaction validation checks have been correctly done off-chain.

Since signatures make up a large percentage of transaction costs (gas), in practice zkRollup has the effect of significantly reducing the average cost per transaction. This allows us to fit more transactions per block, which results in a greater overall throughput.

![](https://cdn-images-1.medium.com/max/800/1*Ry2am_lnNAjnZDNdDQTvDw.png)
