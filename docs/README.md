# Iden3 Developer Portal
<img src="./imgs/iden3-icon2.png" style="float:right; max-width: 200px; margin-left: 30px;">

Welcome to the Iden3 developer portal. Here you'll find eveything you need to get started building with Iden3.

## What is Iden3 and what are its aims

Iden3 is an open source project offering a complete decentralized identity management solution over the public blockchain. It is designed with four fundamental goals in mind:

**Enable Self Sovereign Identities**, whereby identities are created, owned and controlled by the user.

**Accessibility**, by promoting the development of open standards and ease of use.

**Scalability**, by minimizing the number of accesses to the blockchain that identity management and claim processing requires.   

**Privacy by design** by using Zero Knowledge Proofs to prove claims without revealing any information other than the validity of the claim.

## Introduction

>Identity is a uniquely human concept. It is that ineffable “I” of self-consciousness, something that is understood worldwide by every person living in every culture. As René Descartes said, Cogito ergo sum — I think, therefore I am. [Source](http://www.lifewithalacrity.com/2016/04/the-path-to-self-soverereign-identity.html)

 What constitutes your identity? What makes you who you are? What is it about you that distinguishes you from others? Philosophers have argued over these questions since the beginning of civilization. Suffice to say there are no simple answers. Identity is a difficult concept to pin down.

 However, we don't need a precise definition to see that there are problems with how modern society thinks about identity.

 In the words of [Christopher Allen](http://www.lifewithalacrity.com/2016/04/the-path-to-self-soverereign-identity.html):

>Today, nations and corporations conflate driver’s licenses, social security cards, and other state-issued credentials with identity; this is problematic because it suggests a person can lose his very identity if a state revokes his credentials or even if he just crosses state borders. I think, but I am not.

How can we improve on this?

It's clear we're at an inflection point with respect to how the digital world interacts with the physical world.

The legacy systems of the physical world have not kept up with the digital world's rising importance to it. As both worlds continue merging, this will have to change.

This gives us an opportunity to create systems -- from the ground up -- that bridge the two. Systems that operate with a different conception of identity.

If we design them well, they will allow us to redefine how modern society thinks about identity. Perhaps getting us closer to that ineffable "I" of self-consciousness.

At Iden3 we're focused on building the tools and developing the protocols to make this happen.

## More on decentralized identity

### Why does identity matter?

In the [words of Vitalik](https://vitalik.ca/general/2019/04/03/collusion.html):

>Mechanisms that do not rely on identity cannot solve the problem of concentrated interests outcompeting dispersed communities; an identity-free mechanism that empowers distributed communities cannot avoid over-empowering centralized plutocrats pretending to be distributed communities.

In other words, without an identity mechanism, one can't ensure one human one address, or one human one vote. This means that however you structure the rules of the system, those with the most resources will be able to game it.

### How is the existing system failing us?

Since the emergence of the modern state, identities have typically been verified by credentials such as a passport or social network account issued by a central authority such as a state or corporation.

However, as noted in the paper [Verifying Identity as a Social Intersection](https://papers.ssrn.com/sol3/papers.cfm?abstract_id=3375436), such identity systems have several interrelated flaws:

1. They are insecure. Crucial data such as an ID number constantly has to be given out. Yet this is also sufficient to impersonate an individual. On top of this, since all data is stored in a single repository managed by the state or a corporation, it becomes particularly vulnerable to external hacking or internal corruption.

2. They narrow you down to one thing (in system or out, criminal or not, a credit score, etc.). The central database has little use for more information than this. This limits the functionality of the system and results in great injustices (for example convicted individuals find it hard to re-enter society as this is the only information about themselves they can reliably convey).

3. They are artificial, in the sense that the information stored about you usually bears little relation to what you or your friends think of as your identity.

To quote directly from [the paper](https://papers.ssrn.com/sol3/papers.cfm?abstract_id=3375436):

>Recently, new identity paradigms have tried to get around some of these elements. One approach, adopted by “big data” platforms like Facebook and Google, is to overcome thin- ness by storing enormous amounts of detailed information about each individual. we might call this “panoptic identity”. However, such solutions have greatly exacerbated the other two problems, as they require extremely artificial compromises to intimacy through the global sharing of data with platforms that would not otherwise store it, creating exceptional potential security risks.

### Why do we need this vision now?

Given the rising political polarization and the increasing amount of information about us collected, shared, and cross-correlated by governments and corporations, there's a real risk our information will be used against us in ways we cannot imagine.

Decentralized identity systems provide a natural technological check on the ability of governments and corporations to abuse their power. A check that goes beyond formal legal protections.

### How can the developing world benefit?

In the developing world, decentralized identity systems have the potential to help bring millions of people out of poverty.

To quote the words of [Timothy Ruff](https://medium.com/evernym/7-myths-of-self-sovereign-identity-67aea7416b1):

>Most of us take for granted that we can prove things about ourselves, unaware that over a billion people cannot. Identity is a prerequisite to financial inclusion, and financial inclusion is a big part of solving poverty.

### Use Cases

#### Liquid democracy

Imagine if you could vote every two weeks to express your political sentiment regarding interest rates.

Imagine if you could have a direct say in any decision, rather than relying on elected politicians to represent you.

Imagine if those in power were held accountable in real-time, rather than once every few years.

This is the promise of liquid democracy.

Liquid democracy exists somewhere in the sweetspot between direct and representative democracy. 

As with direct democracy, everyone has the opportunity to vote on every issue. However, unlike direct democracy, you also have the choice to delegate your vote to someone else.

You can even choose to delegate your votes on different issues to different people. For example, on environmental issues you might choose to delegate your vote to your favourite environmentalist. Whereas on issues concerning government debt and taxation you might choose your brother-in-law.

Note that this ability to delegate is recursive. Meaning that if your brother-in-law in turn chooses to delegate his vote on financial issues to his favourite economist, your vote will also be delegated to him.

If you're unhappy with how one of your delegates is voting, you can take that power away from him/her immediately and either vote yourself or redelegate to someone you deem more trustworthy.

Those with the most delegations essentially become our representatives. Except unlike representative democracy, they are held accountable in real time.

>A system like this addresses the uninformed voter issue that a direct democracy creates by allowing voters to allot their votes to experts in their fields. It also addresses the corruption issues of a representative democracy because citizens can rescind their vote from someone instantly, forcing delegates to vote in the best interest of their constituents. It is the best of both worlds that truly gives the power of influence to the voters. [Source](https://media.consensys.net/liquid-democracy-and-emerging-governance-models-df8f3ce712af)

On top of this all votes are transparent and easily verifiable by anyone (whilst preserving anonymity).

This sounds like a fair, transparent and corruption-free government... why haven't we implemented this before?

Since there's no central government under this form of democracy, before we can implement such a system we first need to figure out how to store and verify identities in a secure, private, and decentralized way.

We also need to ensure one person is not able to vote multiple times (what's known as a Sybil attack).

The key is a voting protocol with a built in (privacy-preserving) decentralized identity system  -- one that can resist Sybil attacks by requiring some basic verification and reputation for each user while still protecting their pseudonymous identity.

In other words, decentralized identity is the first big unlock that's needed to turn liquid democracy into a reality.

## Contributing

Our team can always use your feedback and help to improve the tutorials and materials included. If you don't understand something, or cannot find what you are looking for, or have any suggestion, help us make the documentation better by letting us know! You can do this by submitting an issue or pull request on the [GitHub repository](https://github.com/iden3/docs/issues).

P.S. Before you submit any changes, please make sure to read our [contribution guidelines]().


