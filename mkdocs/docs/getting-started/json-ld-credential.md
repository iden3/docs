# JSON-LD Credentials

### JSON-LD Credential and Claim

A [Claim](./generic-claim.md) is the core data structure used by Iden3 to represent information. A claim by itself doesn't contain enough meaningful information to be read, understood and consumed (e.g. by the wallet). For example it doesn't tell anything about the meaning of values stored inside the data slots. 

The JSON-LD credential is able to pack the information contained in a claim in a more human-readable way. Furthermore, a JSON-LD credential does not only contain the claim itself but other proofs needed for the subject of the claim needs to consume the claim with other Verifiers.

Let's anaylise what is a JSON-LD Credential with a practical example. The first tab contains a claim attesting to someone's date of birth (this is the same claim generated in the [Generic Claim Section](./claim/generic-claim.md)). The second tab contains the corresponding JSON-LD Credential of type `KYCAgeCredential`.

=== "Core Claim Format"

    ```go
    Index:
    {
    "3613283249068442770038516118105710406958", // Claim Schema hash
    "86645363564555144061174553487309804257148595648980197130928167920533372928", // ID Subject of the claim
    "19960424", // First index data slot stores the date of birth
    "1"  //  Second index data slot stores the document type
    }

    Value:
    { 
    "227737944108667786680629310498", // Revocation nonce 
    "0",
    "0", // first value data slot
    "0"  // second value data slot
    }	
    ```

=== "JSON-LD Credential"

    ``` json
    {
        "id": "eca80230-6ed1-4251-8fe9-3c0204ba10ba",
        "@context": [
            "https://raw.githubusercontent.com/iden3/claim-schema-vocab/main/schemas/json-ld/iden3credential.json-ld",
            "https://raw.githubusercontent.com/iden3/claim-schema-vocab/main/schemas/json-ld/kyc-v2.json-ld"
        ],
        "@type": [
            "Iden3Credential"
        ],
        "expiration": "2361-03-22T00:44:48+05:30",
        "updatable": false,
        "version": 0,
        "rev_nonce": 1909830690,
        "credentialSubject": {
            "birthday": 19960424,
            "documentType": 1,
            "id": "113TCVw5KMeMp99Qdvub9Mssfz7krL9jWNvbdB7Fd2",
            "type": "KYCAgeCredential"
        },
        "credentialStatus": {
            "id": "https://fe03-49-248-235-75.in.ngrok.io/api/v1/claims/revocation/status/1909830690",
            "type": "SparseMerkleTreeProof"
        },
        "subject_position": "index",
        "credentialSchema": {
            "@id": "https://raw.githubusercontent.com/iden3/claim-schema-vocab/main/schemas/json-ld/kyc-v2.json-ld",
            "type": "KYCAgeCredential"
        },
        "proof": [
            {
                "@type": "BJJSignature2021",
                "issuer_data": {
                    "id": "113TCVw5KMeMp99Qdvub9Mssfz7krL9jWNvbdB7Fd2",
                    "state": {
                        "claims_tree_root": "ea5774fac8d72478d4db8a57a46193597bb61475fc9e72bdc74a0ce35aa85518",
                        "value": "5ccc30d5d0360170a29188d5a907381098801a1ab09003493d9833fa4d95271f"
                    },
                    "auth_claim": [
                        "304427537360709784173770334266246861770",
                        "0",
                        "6610201945487752676983171932113332232608162355365546060896064454271869708127",
                        "11380100573862532536254569563965305187461805636461289256869908267462627351172",
                        "0",
                        "0",
                        "0",
                        "0"
                    ],
                    "mtp": {
                        "existence": true,
                        "siblings": []
                    },
                    "revocation_status": "https://fe03-49-248-235-75.in.ngrok.io/api/v1/claims/revocation/status/1909830690"
                },
                "signature": "5e1356754a061c9f691496a4b4bd4cab5d1d74eb835ef7575fc6b2c1e8b4311dab9e2b544f9c3f4701324b1e0b3a8c09de22425de9038c2a08f98f6963f17102"
            }
        ]
    }
    ```


The core claim (1st tab) contains a limited set of information such as the schema hash, the identity subject of the claim, the data slots stored inside the claim (in this case the date of birth) and the revocation nonce of the claim. It's worth noting that the claim by itself doesn't say anything about the meaning of this content. How can someone infer that that value refers to a birthday? Furthermore the claim doesn't reveal information about the issuer, nor whether it has been revoked or not. All these set of extended information about a claim are included in the JSON-LD format in order to allow other parties to consume and understand the content of a claim.

In particular the first part of the JSON-LD Credential contains the details of the claim: 

- `id` namely the identifier of the credential itself
- `context`, as per JSON-LD spec is used to establish a description of the common terms that we will be using such as "expiration", "updatable" .... In particular the first one is standard for iden3 credential vocabulary while the second is specific to this type of claim. 
- `type` defines the type of the credential itself, when dealing with iden3 claim the credential should always be named Iden3Credential
- `expriation` which is a field contained in the claim in v_0(specifically, the hash of this value is contained in the claim)
- `updatable` which is a field contained in the claim in i_0
- `version` which is a field contained in the claim in i_0
- `rev_nonce` which is a field contained in the claim in v_0 
- `credentialSubject` which contains all the details about the subject of the claim and the information contained in the claim regarding the subject
- `credentialStatus` which contains a url to fetch the revocation status of the claim 
- `subject_position` indicates whether the identifier of the subject of the claim should be stored inside index data slot (i_1) or value data slot (v_1)
-  `credentialSchema` defines the [Schema](./claim/claim-schema.md) of the claim itself

The second part ofthe JSON-LD Credential contains a cryptographic proof that the credentials was issued by a specific issuer:

- `@type` indicates the way the proof was generated. It can be either "BJJSignature2021" or "Iden3SparseMerkleProof"
- `issuer_data` contains all the data related to the issuer of the claim. Including its [identifier](./identity/identifier.md) (`id`), its [identity state](./identity/identity-state.md) value at the time of the issuance (`id`), its [Auth Claim](./claim/auth-claim.md) (`auth_claim`), the merkle tree proof that the Auth Claim belongs to the Claims Tree at the time of the issuance (`mtp`) and, lastly, a url to fetch the revocation status of the issuer's auth claim (`revocation_status`)
- `signature` contains the signed claim [generated using the issuer's private key](./signature-claim/signature.md)

> In this case the claim was issued by signature, in the case of claim of Merkle Tree Type the proof array would also contain a second value, namely the `Iden3SparseMerkleProof` of inclusion of the issued claim inside the issuer's Claims Tree 

The subject of the claim will store the JSON-LD format credential inside their wallet. Starting from the details contained inside the Credential he/she will be able to generate zk proofs and present it to Verifiers in the form of [JWZ](../verifier/verification-library/jwz.md).

- [] Why we need that JSON-LD Credential?
- [] How am I able to get from the credential to core-claim? Is it gonna match? Is it also able to parse revocation nonce etcetera? 
- [] Does the parsing also work in reverse? From claim to VC?
- [] What is `signature`?
- [] Add part to explain how to parse the VC into a core claim

