# Merkle tree

A Merkle tree is a [binary tree](https://en.wikipedia.org/wiki/Binary_tree) built using hash pointers (if you're unfamiliar with what a hash pointer or function is, see the definitions section at the bottom of the page).

We care about Merkle trees because we want to build a data structure that:

1. Can store lots of data (**scalability**)
2. Makes it easy to prove that some data exists (**proof of membership**)
3. Allows us to check that data hasn't been altered (**tamper resistance**)

Merkle trees satisfy these three properties.

## Specification

Before we take a closer look at the above properties, let's go through how to build a Merkle tree given some data.

Suppose we have a number of blocks containing data. And that these blocks make up the leaves of our tree.

*[image]*

The first step is to group these data blocks into pairs.

*[image]*

Then for each pair of blocks, we build a data structure that has two hash
pointers, one to each block.

*[image]*

In other words, the hash of each block is stored in a parent node. And these parent nodes make up the next level of the tree.

*[image]*

Next, the parent nodes are in turn grouped in pairs, and their hashes stored one level up the tree.

*[image]*

We continue doing this until we reach a single block, the root of the tree.

*[image]*

## Tamper resistance

It turns out that any attempt to tamper with any piece of data can be detected by simply remembering
the hash pointer at the root of the tree.

To understand why this is the case, let’s look at what happens if an adversary wants to tamper with a data block.

If an adversary tampers with a block at the leaf of our tree.

*[image]*

That will cause the hash pointer that's one level up to not match.

*[image]*

So she'll have to tamper with that too.

*[image]*

Which means, she'll have to tamper with the hash pointer one level up from there.

*[image]*

And so on... Eventually she'll get to the root. If she tries to tamper with the hash pointer here, we'll know because this is the pointer we've remembered.

*[image]*

## Proof of membership

Merkle trees allow us to quickly check membership. What do we mean by that?

Say that, as usual, we remember just the root. And we want to prove that a certain data block is a member of the Merkle tree.

*[image]*

All we need to find is this data block, and the blocks on the path from the data block to the root.

*[image]*

We can ignore the rest of the tree, as the blocks on this path are enough to allow us to verify the hashes all the way up to the root of the tree.

*[image]*

In technical terms:

*This means that if there are n nodes in the tree, only about log(n) items need to be shown. And since each step just requires computing the hash of the child block, it takes about log(n) time for us to verify it. And so even if the Merkle tree contains a very large number of blocks, we can still prove membership in a relatively short time. Verification thus runs in time and space that’s logarithmic in the number of nodes in the tree.* [Source](https://d28rh4a8wq0iu5.cloudfront.net/bitcointech/readings/princeton_bitcoin_book.pdf) (pg 35)

## Scalability

Storing data on a blockchain is expensive. Merkle trees help us minimize the amount of data stored on chain.

How so? As we saw in the previous sections, to ensure tamper resistance and proof of membership we only need to store the root of the tree, not the whole tree. This means that, no matter how big the tree is, the only piece of data we actually need to store on chain is the root.

## Why we use Merkle trees at iden3

At iden3, one of our major goals is scalability. Specifically, we believe anybody should be able to create as many identities as they want. And that **any identity should be able to generate as many claims as they want.**

Imagine if you had to make a new transaction to the blockchain every time you wanted to make a new claim? Even worse, imagine you're a government and you're responsible for making millions of claims every day...

To achieve this goal requires minimizing the amount of data stored on chain. This is where Merkle trees come in.

Even if you're a government that's making millions of claims a day, you can just contruct a tree (off chain) with each claim as a separate data block, and simply calculate and store the root on chain.

In other words, Merkle trees allow prolific claim generators to add/modify **millions of claims** in a single transaction.

This makes it easy to scale claims.

## Definitions


### Hash functions

A hash function basically maps an input string of any size to an output string of a fixed size.

It must be efficiently computable (by that we mean that for any given input string, we can figure out the output of the hash function in a reasonable amount of time. More technically, computing the hash of an n‐bit string should have a running time that is O(n).

For a hash function to be cryptographically secure, it must have three additional properties: 

1. Collision resistance
2. Hiding
3. Puzzle-friendliness

While we won't get into the details here, let's briefly discuss what each of these properties mean.

**Collision resistance** means that nobody can find two inputs that map to the same output.

**Hiding** means that given an output there's no feasible way to figure out the input that generated it.

**Puzzle-friendliness** is a little more complicated. Intuitively it means it's very hard to target the hash function to come out to some particular output value y. Don't worry if you don't see why this property is useful, for our purposes, it isn't very important.


### Hash pointers

A hash pointer is simply a pointer to where some information is stored together with a cryptographic hash of the
information. **A pointer gives you a way to retrieve the information, whereas a hash pointer also gives you a way to verify that the information hasn’t changed.**

In other words, a hash pointer is a pointer to where data is stored together with a cryptographic hash of the value of that data at some fixed point in time.

If at some point in the future, we want to check the data hasn't changed, we simply hash the data again and check that the new output (cryptographic hash) matches the previous output. This works because we know by the **collision resistance** property of the hash function, that nobody can find two inputs that map to the same output. So if the output is the same, the input must also have been the same.
