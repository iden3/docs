# OnChain Identity

An OnChain Identity is an identity that is controlled by a smart contract. It is a special case of Ethereum-controlled Identity.

An OnChain Identity manages its own identity trees on-chain (claims, revocation, and roots trees). It can **issue credentials** (by adding them to its claims tree), **revoke credentials**, and **perform state transitions** to update its overall state.

## OnChain Issuer

An OnChain Issuer is a specialized OnChain Identity used to issue credentials to other identities.

## Libraries and Contracts

* [IdentityLib.sol](https://github.com/iden3/contracts/blob/master/contracts/lib/IdentityLib.sol) - A library that can create identities, manage trees, issue or revoke credentials, and perform state transitions.
* [GenesisUtils.sol](https://github.com/iden3/contracts/blob/master/contracts/lib/GenesisUtils.sol) - A library that can generate an identity ID from an Ethereum address or genesis identity state, and verify the ID's structure.
* [IdentityBase.sol](https://github.com/iden3/contracts/blob/master/contracts/lib/IdentityBase.sol) - A base contract for building OnChain Identity or Issuer contracts, with the required public interfaces implemented.

## Benefits & Possible Use Cases

### Programmable Identity
OnChain Identities can have **custom logic** implemented in their smart contract code **to manage credentials and keys**. For example, an OnChain Identity can require **multi-signature approval** for issuing or revoking credentials, or implement **time-based restrictions** on certain actions.

### Composable
OnChain Identities are smart contracts, which means they can include other smart contract logic to extend their functionality. 
For example, an OnChain Identity can be a **Smart Account Wallet** that allows users to manage their **funds and identity in a single contract** and benefit from features like **social recovery**, **gasless transactions**, and more.
They can also interact with other smart contracts and DeFi protocols, enabling **identity-aware** interactions in use cases such as reputation-based lending or insurance.

### Transparent and Auditable
Smart contract code **defines rules**, such as who and how is issuing credentials or rotating keys. This code is public and can be **audited by anyone**. This increases trust in the system, as users can verify that the identity behaves as expected.

### Self-Issuance / Trustless Decentralized Issuance
Users can interact with the smart contract to **issue themselves credentials** that are valid and verifiable, exactly like those issued by regular off-chain issuers. This allows users to **make their existing reputation portable** and usable **in a privacy-preserving manner**, such as proving balance, NFT ownership, or trade volumes without revealing their addresses and exact amounts.
And there is **no need to trust a third party** for proper user verification and credential issuance. The smart contract can **enforce the correct behavior**.

### ZK-Enabled Self-Issuance with Private Data
It is also possible to **leverage private data sources**, such as **digitally signed documents** (e.g., **e-Passports**, **emails** with DKIM signatures, signed **PDFs**), **to issue credentials in a privacy-preserving manner using zero-knowledge proofs**. A user can create credentials on their own device and prove that they were created correctly, following the rules of a smart contract and a specific ZK circuit.
For example, **a user uses a mobile phone to read their biometric passport and generates a verifiable credential based on its data**. A ZK circuit then proves that the resulting credential corresponds to the passport data, and that this data is properly signed by valid government keys. **Only the hash of the credential, along with a zero-knowledge proof, is sent on-chain** to be verified and added to the OnChain Identity's claims tree. This ensures that **private data never leaves the user's device**, while the user still receives a valid credential.

### Identity Recovery
If a user loses access to their keys, it is possible to implement a **recovery mechanism** in the smart contract. For example, a set of **trusted parties (friends, family members, DAO members, etc.) can vote to give access back to the user** by adding a new authentication key to user's OnChain Identity. **Alternatively, the user can present a zero-knowledge proof that they control another identity with a credential stating it is the same real-world identity** (e.g., based on a self-issued zk passport credential or a KYC credential from a trusted issuer), and request to add a new authentication key and revoke the old one.

### Smart Governance
OnChain Identities can implement governance mechanisms allowing **multiple parties to manage the identity collaboratively**.
For example, a DAO may vote to issue or revoke a credential (such as an executive or validator role) of a specific user.

### Organizational Identity Management
Organizations can use OnChain Identities to manage their members' and clients' credentials. For example, a company can have an OnChain Identity that issues credentials to its employees, such as right to represent the company in certain transactions or access specific resources such as instant messaging channels, internal tools and physical locations. The company can also revoke credentials when an employee leaves, so that access is automatically removed.

## OnChain Identity State Transition
An OnChain Identity state transition is performed by calling the `transitState` function of the `IdentityLib` library. This function calculates new state from the claims, revocations and roots tree roots, which collectively define the current identity's status. It also verifies whether the roots have changed since the last state transition, due to claims or revocations added to the relevant trees.

Finally, it calls `transitStateGeneric` function of the `State` contract, which is designed to be generic and may be used in the future to perform state transitions for other types of identities or transition logic.

#### State data consistency warning
Please note that neither the `State` contract nor `IdentityLib` impose restrictions on users who perform a state transition using `IdentityLib.transitState()`, and then subsequently perform another state transition using the `State.transitState()` contract with BJJ key authentication. This sequence of actions may result in inconsistent state data between the `State` and OnChain Identity smart contracts. It is the **sole responsibility of the OnChain Identity owner to prevent such situations**.
