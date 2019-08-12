# Claims

An [identity](basics/glossary/identity.md)  can provide a claim. You can think of a claim as a statement: something an identity is saying.

Most of the time, these statements refer to other identities. In other words **claims usually create relations between identities.**

For example, when a university (identity) says that a student (identity) has a degree, this is a statement (claim) about the student. This statement creates a relation between the student and the university.

**Claims can be public or private.** And it turns out that almost anything we say or do can be thought of as a claim. Company invoices, Facebook likes, email messages, can all be thought of as claims.

## Examples of claims

- A certificate (e.g. birth certificate)

- A debt recognition

- An invoice

- An Instagram "Like"

- An endorsement (reputation)

- An email

- A driving license

- A role in a company

- ... Almost anything!

## Direct claims

If an identity wants to create many claims, they can put them all in a database, construct a [Merkle tree](basics/glossary/merkletree.md) of that database, and just publish (with a transaction) the root of the Merkle tree on-chain.

If the identity wants to update the claims later, they repeat the same process and just publish the new root of the Merkle tree.

For example, one could imagine a government adding/modifying millions of claims in a single transaction.

## Indirect claims

While direct claims scale really well for identities that make a lot of claims (since million of claims can be batched in a single transaction), the average user will probably only need to make a few claims a day, and so won't benefit from this batching.

This is where indirect claims come in handy. Instead of having to pay gas everytime to update the Merkle root on-chain, indirect claims allow users to send claims off-chain to a **relayer**.

The idea is that with relayers, millions of users can create millions of claims on mainnet **without** spending any **gas** (since the relayer is responsible for batching the claims and publishing the transactions).

On top of this, using [zero knowledge proofs](basics/glossary/zeroknowledge.md) we can ensure that the relayer is trustless. In other words we can make sure the relayer can't lie about the claims we sent them. The worst they can do is not publish them (and if this happens we as the user always have the choice to change relayers).

