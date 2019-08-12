# Zero Knowledge

*In cryptography, a zero-knowledge proof or zero-knowledge protocol is a method by which one party (the prover) can prove to another party (the verifier) that they know a value x, without conveying any information apart from the fact that they know the value x.* [Source](https://en.wikipedia.org/wiki/Zero-knowledge_proof) 

In other words, zero-knowledge proofs allow us to prove something specific without revealing any extra information.

Why do we care? Simply put, when we're talking about claims, sometimes we want to prove things in a private way.

## Examples

### Nightclub entry

Say you want to enter a nightclub, and you need to prove to the bouncer that you are over 18. But you don't want to reveal to him your name, address, or anything else that's not relevant.

With a zero-knowledge proof you can prove that you hold the key that belongs to an identity that the state says is over 18, without revealing anything else about that identity.

### ICO participation

Say an ICO is only available to KYC or authorized users. With ZK proofs you can prove that you are an authorized person to participate in the ICO without revealing who you are or how much you spent.

### Anonymous Voting

Similar to the above, ZK proofs allow you to prove that you are an eligible identity, without revealing your identity.

## Non-reusable proofs

A non-reusable proof is a received proof that is not valid to send to a third identity.

For example, imagine that you belong to a political party, P. And P has made a private claim that you belong to it.

Say that you want to prove to another identity that you belong to P, but you don't want that other identity to be able to pass on that proof to others. In other words, you want to make sure the proof stays between the two of you.

We can do this using zero-knowledge proofs.

How?

To prove something -- let's call it A -- we can create a new proof B that is valid either if A is valid or we know the private key of the recipient, R.

*[image]*

Clearly we don't know R's private key, so when we share a valid proof B with R, R knows that A must be valid.

*[image]*

To see why B is non-reusable. Suppose R wants to share B with another recipient R'.

*[image]*

Now, from the perspective of R', B is valid either if A is valid or R knows her own private key.

*[image]*

But since R clearly knows her own private key, R' can't tell whether A is valid.

*[image]*
