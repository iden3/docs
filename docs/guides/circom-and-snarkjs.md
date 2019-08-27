# How to use circom and snarkjs

Hello and welcome!

In this guide we'll guide you through the creation of your first zero-knowledge zk-snark circuit using [circom](https://github.com/iden3/circom) and [snarkjs](https://github.com/iden3/snarkjs).

[Circom](https://github.com/iden3/circom) is a library that allows you to build circuits to be used in zero knowledge proofs. 

While [snarkjs](https://github.com/iden3/snarkjs) is an independent implementation of the zk-snarks protocol -- fully written in JavaScript.

Circom is designed to work with snarkjs. In other words, any circuit you build in circom can be used in snarkjs.

We'll start by covering the various techniques to write circuits, then move on to creating and verifying a proof off-chain, and finish off by doing the same thing on-chain on Ethereum.

If you have zero knowledge about zero-knowledge 😋 or are unsure about what a zk-snark is, we recommend you read [this page](basics/glossary/zeroknowledge.md) first.


## 1. Installing the tools

### 1.1 Prerequisites

First off, we need to be sure we have a recent version of `Node.js` installed.

While any version after `8.12.0` should work fine, we recommend you install version `10.12.0` or later.

Why? These later versions of Node include big integer libraries natively. `snarkjs` makes use of this feature (if available) to improve performance by up to **10x**.

To see which version of Node you have installed, from the command line run:

```node -v```

To download the latest version of Node, [click here](https://nodejs.org/en/download/).

### 1.2 Installing **circom** and **snarkjs**

As stated in the introduction, circom and snarkjs are the libraries we use to create zero-knowledge proofs.

If you haven't done so already, you can install them from NPM by running the following commands:
```
   npm install -global circom
   npm install -global snarkjs
```

Hopefully both libraries installed successfully.

If you're on a Unix machine and you're seeing some errors (e.g. `node-gyp rebuild`) it's probably because you need to update your version of Node to the latest long term support (LTS) version, at the time of writing this is `v10.15.3`.

If you're seeing one or more errors that look like:

`EACCES: permission denied`

It's probably because you originally installed Node with root permissions. Because of this, writing to your npm directory (`npm install -global`) requires root permissions too.

While it's not a good idea to have Node  installed this way, one way to quickly give yourself root permissions is to run the slightly modified commands:

```
sudo npm install -global --unsafe-perm circom
sudo npm install -global --unsafe-perm snarkjs
```

An arguably better way to fix this is to follow the steps outlined in this [stackoverflow answer.](https://stackoverflow.com/a/24404451)

## 2. Building a circuit with circom

### 2.1 Definition
First off, let's define what we mean by a circuit.

For our purposes, a circuit is equivalent to a **statement** or **deterministic program** which has an output and one or more inputs. 

**[insert image]**

There are two types of possible inputs to a circuit: `private` and `public`. The difference being that a `private` input is hidden from the verifier.

### 2.2 Motivation

The idea here is that given a `circom` circuit and its inputs, the prover can run the circuit and generate a proof -- using `snarkjs` -- that she ran it correctly.

With the proof, the output, and the public input(s), the prover can then prove to the verifier that she knows one or more private inputs that satisfy the constraints of the circuit, **without revealing anything about the private input(s)**.

In other words, even though the verifier has **zero knowledge about the private inputs** to the circuit, the proof, the output, and the public inputs(s) are enough to convince her that the prover's statement is valid (hence the term zero-knowledge proof).

### 2.3 Toy example

Don't worry if some (or all) of the above sounded a little abstract. In this section we'll go through an example that should help clarify things.

Let’s create a circuit that tries to prove to someone (the verifier) that we are able to factor an integer `c`.

It turns out that factoring an integer can be quite difficult -- in particular, the prime factorization of very large numbers can be [very difficult](https://www.reddit.com/r/math/comments/2jo786/why_is_the_prime_factorization_of_very_large/cldj3a9/).

For very large numbers, no efficient, non-quantum integer factorization algorithm is known. However it has not been proven that no efficient algorithm exists.

The presumed difficulty of this problem is at the heart of widely used algorithms in cryptography such as [RSA](https://en.wikipedia.org/wiki/RSA_(cryptosystem)).

If this problem were easy to solve, cryptography as we know it would break down. Which means there's a big chance that cryptocurrencies would cease to exist from one day to the next!

In this toy example we'll neither work with very large numbers, nor restrict their factors to primes. Nevertheless the general principle remains the same.

We want to prove that we know two numbers (call them `a` and `b`) that multiply together to give `c`. Without revealing `a` and `b`.

1. The first step is to create (and move into) a new directory called ``factor`` where we'll put all the files that we want to use in this guide.
```
mkdir factor
cd factor
```

   >Note: if we were designing a circuit for actual use, we'd probably be better off creating a ``git`` repository with a ``circuits`` directory containing the necessary scripts to build all our circuits, and a ``test`` directory with all our tests.

2. Next, we want to create a new file (in `factor`) named `circuit.circom` which looks like this:
```
   template Multiplier() {
       signal private input a;
       signal private input b;
       signal output c;
       
       c <== a*b;
   }

   component main = Multiplier();
   ```
   As you can see, this circuit has **two private input** signals named ``a`` and `b` and **one output** signal named `c`.
   
   The only thing this circuit does is force the signal `c` to be
   the value of `a*b`.

   **[Explain <==]**
   
   >Note: after declaring the ``Multiplier`` template, we instantiate it with a component named ``main``. When compiling a circuit a component named ``main`` must always exist.
   
3. We are now ready to compile the circuit. To compile the circuit to a file named `circuit.json`, run the following command:
```
circom circuit.circom -o circuit.json
```
**[Explain the need to compile and explain circuit.json]**

## 3. Taking the compiled circuit to *snarkjs*

Now that the circuit is compiled, we will continue with ``snarkjs``.
Please note that you can always access the help of ``snarkjs`` by
typing:

`snarkjs --help`

### 3.1 View information and stats regarding a circuit

To show general statistics of this circuit, you can run:

`snarkjs info -c circuit.json`

You can also print the constraints of the circuit by running:

`snarkjs printconstraints -c circuit.json`

### 3.2 Setting up using *snarkjs*

Ok, let’s run a setup for our circuit:

`snarkjs setup`

   By default ``snarkjs`` will look for and use ``circuit.json``. You
   can always specify a different circuit file by adding
   ``-c <circuit JSON file name>``

The output of the setup will in the form of 2 files:
``proving_key.json`` and ``verification_key.json``

### 3.3. Calculating a witness

Before creating any proof, we need to calculate all the signals of the
circuit that match (all) the constrains of the circuit.

``snarkjs`` calculates these for you. You need to provide a file with
the inputs and it will execute the circuit and calculate all the
intermediate signals and the output. This set of signals is the
*witness*.

The zero knowledge proofs prove that you know a set of signals (witness)
that match all the constraints but without revealing any of the signals
except the public inputs plus the outputs.

For example, Imagine that you want to prove that you are able to factor
33 that means that you know two numbers ``a`` and ``b`` that when you
multiply them, it results in 33.

   Of course you can always use one and the same number as ``a`` and
   ``b``. We will deal with this problem later.

So you want to prove that you know 3 and 11.

Let’s create a file named ``input.json``, with the following content:

`{"a": 3, "b": 11}`

And now let’s calculate the witness:

`snarkjs calculatewitness`

You may want to take a look at ``witness.json`` file with all the
signals.

Create the proof
~~~~~~~~~~~~~~~~

Now that we have the witness generated, we can create the proof.

`snarkjs proof`

This command will use the ``prooving_key.json`` and the ``witness.json``
files by default to generate ``proof.json`` and ``public.json``

The ``proof.json`` file will contain the actual proof. And the
``public.json`` file will contain just the values of the public inputs
and the outputs.

### 3.4 Verifying the proof

To verify the proof run:

`snarkjs verify`

This command will use ``verification_key.json``, ``proof.json`` and
``public.json`` to verify that is valid.

Here we are veifying that we know a witness that the public inputs and
the outputs matches the ones in the ``public.json`` file.

If the proof is ok, you will see an ``OK`` in the screen or ``INVALID``
otherwise.

### 3.5 Generate the solidity verifier

`snarkjs generateverifier`

This command will take the ``verification_key.json`` and generate a
solidity code in ``verifier.sol`` file.

You can take the code in ``verifier.sol`` and cut and paste in remix.

This code contains two contracts: Pairings and Verifier. You just need
to deploy the ``Verifier`` contract.

   You may want to use a test net like Rinkeby, Kovan or Ropsten. You
   can also use the Javascript VM, but in some browsers, the
   verification takes long and it may hang the page.

### 3.6 Verifying the proof on-chain


The verifier contract deployed in the last step has a ``view`` function
called ``verifyProof``.

This function will return true if the proof and the inputs are valid.

To facilitiate the call, you can use snarkjs to generate the parameters
of the call by typing:

`snarkjs generatecall`

Just cut and paste the output to the parameters field of the
``verifyProof`` method in Remix.

If every thing works ok, this method should return true.

If you just change any bit in the parameters, you can check that the
result will be false.

### 3.7 Bonus track

We can fix the circuit to not accept one as any of the values by adding
some extra constraints.

Here the trick is that we use the property that 0 has no inverse. so
``(a-1)`` should not have an inverse.

that means that ``(a-1)*inv = 1`` will be inpossible to match if ``a``
is one.

We just calculate inv by ``1/(a-1)``

So let’s modify the circuit:

```
   template Multiplier() {
       signal private input a;
       signal private input b;
       signal output c;
       signal inva;
       signal invb;
       
       inva <-- 1/(a-1);
       (a-1)*inva === 1;
       
       invb <-- 1/(b-1);
       (b-1)*invb === 1;    
       
       c <== a*b;
   }

   component main = Multiplier();
```

A nice thing of circom language is that you can split a <== into two
independent acions: <– and ===

The <– and –> operators Just assign a value to a signal without creating
any constraints.

The === operator just adds a constraint without assigning any value to
any signal.

The circuit has also another problem and it’s that the operation works
in Zr, so we need to guarantee too that the multiplication does not
overflow. This can be done by binarizing the inputs and checking the
ranges, but we will reserve it for future tutorials.

## 4. Where to go from here.

You may want to read the [README](https://github.com/iden3/circom) to
learn more features about circom.

You can also check a a library with many basic circuits lib
binaritzations, comparators, eddsa, hashes, merkle trees etc
[here](https://github.com/iden3/circomlib) (Work in progress).

Or a exponentiation in the Baby Jub curve
[here](https://github.com/iden3/circomlib) (Work in progress).

## 5. Final note

There is nothing worst for a dev than working with a buggy compiler.
This is a very early stage of the compiler, so there are many bugs and
lots of works needs to be done. Please have it present if you are doing
anything serious with it.

And please contact us for any isue you have. In general, a github issue
with a small piece of code with the bug is very worthy!.

Enjoy zero knowledge proving!
