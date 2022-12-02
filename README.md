<div align="center">
<img src="mkdocs/docs/logo-dark.svg" width="150"/>
</div>
<br />
<div align="center">

[![Chat on Twitter][ico-twitter]][link-twitter]
[![Chat on Telegram][ico-telegram]][link-telegram]
[![Website][ico-website]][link-website]
<!-- [![GitHub repo][ico-github]][link-github] -->

</div>

[ico-twitter]: https://img.shields.io/twitter/url?color=black&label=Iden3&logoColor=black&style=social&url=https%3A%2F%2Ftwitter.com%2Fidenthree
[ico-telegram]: https://img.shields.io/badge/telegram-telegram-black
[ico-website]: https://img.shields.io/website?up_color=black&up_message=iden3.io&url=https%3A%2F%2Fiden3.io
<!-- [ico-github]: https://img.shields.io/github/last-commit/iden3/docs?color=black -->

[link-twitter]: https://twitter.com/identhree
[link-telegram]: https://t.me/iden3io
[link-website]: https://iden3.io
<!-- [link-github]: https://github.com/iden3/docs -->

# Identity protocol 

## Prove your access rights, not your identity

iden3 is a next-generation private access control based on self-sovereign identity, designed for decentralised and trust-minimised environments.

## Privacy for all

Everyone has the right to liberty and equality, the right freely to participate in their community, and the right to privacy.

The aim of the iden3 protocol is to empower people and create a more inclusive and egalitarian foundation for better human relationships through open-source cryptography and decentralised technologies.

<div align="center">
<br />

Privacy by design            | Decentralised                     |  Open source 
:---------------------------:|:---------------------------------:|:-------------------------------:
![](mkdocs/docs/imgs/icons/privacy.svg)  | ![](mkdocs/docs/imgs/icons/decentralised.svg) | ![](mkdocs/docs/imgs/icons/open-source.svg)

<br />
</div>

## Iden3 protocol libraries

- **Crypto library ([go-iden3-crypto](https://github.com/iden3/go-iden3-crypto))**
    <br />Implementation of Poseidon hash and Baby JubJub Eliptic curve

- **Merkle tree sql library ([go-merkletree-sql](https://github.com/iden3/go-merkletree-sql))**
    <br />Implementation of Sparse Merkle tree

- **Core library ([go-iden3-core](https://github.com/iden3/go-iden3-core))**
    <br />Identity core primitives

- **Circuits ([circuits](https://github.com/iden3/circuits))**
    <br />Identity circuits

- **Go-circuits ([go-circuits](https://github.com/iden3/go-circuits))**
    <br />Library for transformation go-core primitives to json inputs for identity circuits

- **Prover server ([prover-server](https://github.com/iden3/prover-server))**
    <br />Wrapper on snarkjs for ZK proof generation

- **Authorization library ([go-iden3-auth](https://github.com/iden3/go-iden3-auth))**
    <br />Library for authentication with zkp verification (edited)

---

# How to run this documentation (locally)
## Install mkdocs
```
pip3 install mkdocs
```

In case you have a rendering problem with the pieces of code, please execute:
```
pip install --upgrade mkdocs
```

## Install mkdocs-material theme
```
pip install mkdocs-material
```

## Install mkdocs-markdown-graphviz (1.3)
```
pip3 install mkdocs-markdown-graphviz==1.3
```

## Run the webserver
At the mkdocs directory execute:

```
mkdocs serve
```
auth 

    ** signal input genesisID;
    // random number, which should be stored by user
    // if there is a need to generate the same userID (ProfileID) output for different proofs
    signal input profileNonce;

    // user state
    ** signal input state;
    ** signal input claimsTreeRoot;
    ** signal input revTreeRoot;
    ** signal input rootsTreeRoot;

    // Auth claim
    ** signal input authClaim[8];

    // auth claim. merkle tree proof of inclusion to claim tree
    ** signal input authClaimIncMtp[IdOwnershipLevels];

    // auth claim - rev nonce. merkle tree proof of non-inclusion to rev tree
    ** signal input authClaimNonRevMtp[IdOwnershipLevels];
    ** signal input authClaimNonRevMtpNoAux;
    ** signal input authClaimNonRevMtpAuxHi;
    ** signal input authClaimNonRevMtpAuxHv;

    // challenge signature
    ** signal input challenge;
    ** signal input challengeSignatureR8x;
    ** signal input challengeSignatureR8y;
    ** signal input challengeSignatureS;

    // global identity state tree on chain
    signal input gistRoot;
    // proof of inclusion or exclusion of the user in the global state
    signal input gistMtp[onChainLevels];
    signal input gistMtpAuxHi;
    signal input gistMtpAuxHv;
    signal input gistMtpNoAux;