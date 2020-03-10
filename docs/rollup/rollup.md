# zkRollup

## Introduction
[zkRollup](https://github.com/barryWhiteHat/roll_up) is a project to scale ethereum with zkSnarks.

In the words of [Vitalik](https://github.com/barryWhiteHat/roll_up):

>We can actually scale asset transfer transactions on ethereum by a huge amount, without using layer 2’s that introduce liveness assumptions (eg. channels, plasma), by using ZK-SNARKs to mass-validate transactions. 

By some estimates, rollup could help ethereum scale to [3000 transactions-per-second](https://blog.iden3.io/istanbul-zkrollup-ethereum-throughput-limits-analysis.html). To put this in context, the ethereum blockchain currently supports roughly 15 tps. And 2000 tps is what the Visa network currently averages.

For a more complete introduction, see this [video](https://www.youtube.com/watch?v=TtsDNneTDDY).

## Our implementation

We've just released our first public testnet. Checkout [zkrollup.io](http://zkrollup.io) for an overvire of how it works.

It's an implementation that uses both [circom](https://github.com/iden3/circom) and [circomlib](https://github.com/iden3/circomlib).

>Note: this is very much a work in progress. If you have any comments at all, please don't hesitate to join our [Telegram group](https://t.me/joinchat/G89XThj_TdahM0HASZEHwg). All feedback is welcome :)
