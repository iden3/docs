# Rollup

[Rollup](https://github.com/barryWhiteHat/roll_up) is a project to scale ethereum with zk-snarks.

In the words of [Vitalik](https://github.com/barryWhiteHat/roll_up):

>We can actually scale asset transfer transactions on ethereum by a huge amount, without using layer 2’s that introduce liveness assumptions (eg. channels, plasma), by using ZK-SNARKs to mass-validate transactions. 

By some estimates, rollup could help ethereum scale to [17,000 transactions-per-second](https://ethresear.ch/t/roll-up-roll-back-snark-side-chain-17000-tps/3675/3). To put this in context, the ethereum blockchain currently supports roughly 15 tps compared to around 45,000 processed by Visa.

For an introduction, see this [video](https://www.youtube.com/watch?v=TtsDNneTDDY).

For an implementation using [circom](https://github.com/iden3/circom) and [circomlib](https://github.com/iden3/circomlib) see [here](https://github.com/iden3/rollup).

>Note: this is very much a work in progress.
