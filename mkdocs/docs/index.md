<div align="center">
<img src="logo-dark.svg" align="center" width="128px"/>
<br /><br />
</div>

[![Chat on Twitter][ico-twitter]][link-twitter]
[![Chat on Telegram][ico-telegram]][link-telegram]
[![Website][ico-website]][link-website]
<!-- [![GitHub repo][ico-github]][link-github] -->
<!-- ![Issues](https://img.shields.io/github/issues-raw/iden3/docs?color=blue) -->
<!-- ![GitHub top language](https://img.shields.io/github/languages/top/iden3/docs) -->
<!-- ![Contributors](https://img.shields.io/github/contributors-anon/iden3/docs) -->

[ico-twitter]: https://img.shields.io/twitter/url?color=black&label=Iden3&logoColor=black&style=social&url=https%3A%2F%2Ftwitter.com%2Fidenthree
[ico-telegram]: https://img.shields.io/badge/telegram-telegram-black
[ico-website]: https://img.shields.io/website?up_color=black&up_message=iden3.io&url=https%3A%2F%2Fiden3.io
<!-- [ico-github]: https://img.shields.io/github/last-commit/iden3/docs?color=black -->

[link-twitter]: https://twitter.com/identhree
[link-telegram]: https://t.me/iden3io
[link-website]: https://iden3.io
<!-- [link-github]: https://github.com/iden3/docs -->

---

# Iden3 Docs

Welcome to the documentation site of the Iden3 project, future-proof tech stack for self-sovereign identity.

---

# <div align="center"><b>[Iden3 on GitHub](https://github.com/iden3)</b></div>

---

## Versatility of applications

The main idea of the iden3 protocol is that each identity is self-soverign and can issue claims on another identity (which can be for an individual, an organisation or a system/machine).

This simple and unique characteristics can lead to creation complex adaptive systems and the following use cases:

<ul>
    <li>Decentralised trust models / web-of-trust</li>
    <li>Decentralised ID verification / proof-of-personhood</li>
    <li>Decentralised voting systems</li>
    <li>Interaction with DeFi / dApps / Web3</li>
    <li>Decentralised payment identifiers</li>
    <li>Private access control</li>
    <li>Authentication and authorisation</li>
    <li>Signing documents and private messaging</li>
    <li>Supply chain and IoT</li>
    <li>NFT ownership</li>
</ul>

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
