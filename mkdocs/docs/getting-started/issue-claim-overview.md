# Overview

In the last section, you have initiated an identity and created different types of claims. As of now, the claims haven’t been published or issued in any way. The goal of this section is to *issue* claims so that the receiver can start using them within other applications. 

Starting from the same [core claim](./claim/generic-claim.md) there are two ways of issuing it: via Signature or via Merkle tree.

## Via Signature

The claim gets signed by the issuer using her private key. The proof of issuance is the signature itself. This action doesn’t modify the identity state of the issuer, so there’s no on-chain transaction involved.

## Via Merkle tree

The claim gets added to the issuer’s [Claims Tree](./identity/identity-structure.md). This action modifies the structure of the Merkle Tree and, therefore, the state has to be updated with the new Merkle root. The **state transition** involves an on-chain transaction. In this case, the proof of issuance is the membership of the claim iself inside the issuer’s Claims Tree.

## Similarities and Differences

Both approaches guarantee that the claim is tamper-resistant. The private zk-based verification of the claim is equally guaranteed in both cases.

The process of updating the on-chain state (in the case of Merkle Tree (MT)) may take around 10/20 seconds, so the claim could not be immediately available to use for the user. Instead, with Signature (S), the claim is immediately ready for use. The biggest advantage of MT approach is that it offers timestamp proof of an identity state: a user could always prove the existence of a claim at a specific point in time according to the block number when it was added to the issuer tree. Naturally, this comes at a cost: the gas fees needed to update the state (458,228 gas used inside the transaction). No on-chain transactions take place in the case of S. 
A further element of difference regards the uniqueness of the claim: in the MT case, there couldn’t be two claims with the same index [Claim Data Structure](https://docs.iden3.io/protocol/claims-structure). This is guaranteed by the characteristic of [Sparse Merkle Tree](./mt.md). With S, an issuer could sign as many claims they want with the same index. Let’s consider an example: a passport issued as a claim. This claim contains the identifier of the passport owner inside its index. MT approach provides a cryptographic guarantee that the issuer cannot duplicate the passport by issuing a claim with the same identifier. S doesn’t.

> Note: This section describes the claim issuance on a protocol level. The way in which issuers and users’ wallets communicate and transfer claims is defined on a platform level. This will be the subject of Polygon ID light issuer tutorial coming out this Autumn. 
 