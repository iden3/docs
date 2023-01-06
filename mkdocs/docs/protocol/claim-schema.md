# Claim Schema 

In order to reuse claims across different services is necessary to keep consistent data formatting. A Claim Schema encodes the structure of a particular claim by defining the usage of data slots.

[Iden3 claims](claims-structure.md) store data inside four data slots: two index slots(i_2,i_3) and two value slots (v_2, v_3). To properly design and fill a claim with information, it is necessary to define which data should be stored inside which data slots. These rules are encoded inside the Claim Schema.

Schemas are described via JSON-LD documents. 

Here an example of Claim Schema of type [**KYCCountryOfResidenceCredential**](https://github.com/iden3/claim-schema-vocab/blob/main/schemas/json-ld/kyc-v2.json-ld#L27).
The `countryCode` should be stored in IndexDataSlotA while the `documentType` in ValueDataSlotB:

```json
...
"countryCode": { 
    "@id": "kyc-vocab:countryCode", 
    "@type": "serialization:IndexDataSlotA" 
}, 
"documentType": { 
    "@id": "kyc-vocab:documentType", 
    "@type": "serialization:ValueDataSlotB" 
}
...
```

A claim issuer could reuse [existing claim schemas](https://github.com/iden3/claim-schema-vocab/tree/main/schemas/json-ld) or create new ones from scratch.

## Example: Auth Claim Schema 

- [Github Document](https://github.com/iden3/claim-schema-vocab/blob/main/schemas/json-ld/auth.json-ld)

```json
{
  "@context": [{
    "@version": 1.1,
    "@protected": true,
    "id": "@id",
    "type": "@type",
    "AuthBJJCredential": {
      "@id": "https://raw.githubusercontent.com/iden3/claim-schema-vocab/main/schemas/json-ld/auth.json-ld#AuthBJJCredential",
      "@context": {
        "@version": 1.1,
        "@protected": true,
        "id": "@id",
        "type": "@type",
        "auth-vocab": "https://github.com/iden3/claim-schema-vocab/blob/main/credentials/auth.md#",
        "serialization": "https://github.com/iden3/claim-schema-vocab/blob/main/credentials/serialization.md#",
        "x": {
          "@id": "auth-vocab:x",
          "@type": "serialization:IndexDataSlotA"
        },
        "y": {
          "@id": "auth-vocab:y",
          "@type": "serialization:IndexDataSlotB"
        }
      }
    }
  }]
}
```

This document describes the schema for a claim of type `AuthBJJCredential`.
The `@id` contains the unique url that contains the JSON-LD Document.
The `auth-vocab` contains the url that describes the value types stored inside this claim, in this case `x` and `y`.
The `serialization` contains the instructions need to parse the raw claim into a JSON-LD document (and viceversa).
The last part of the document contains a reference to the value types `x` and `y`. Their `@id` is corresponding description in the auth-vocab while the `@type` indicates where the values should be stored inside the claim. In this case x and y should, respectively, be stored in `IndexDataSlotA` and `IndexDataSlotB`.

## Schema Hash

The first index slot (i_0) of a [Claim](./claims-structure.md) should store the claim schema itself. Storing the whole JSON-LD document inside the claim would be highly inefficient, so only an hash is stored inside the claim. 

The Schema Hash stored inside a claim is the last 16 bytes of the result of hashing together: 

- `schemaBytes`, the Claim Schema JSON-LD document in bytes format
- `credentialType`, the credential type in bytes format.

```golang
var sHash core.SchemaHash
hash := Keccak256(schemaBytes, []byte(credentialType))
copy(sHash[:], hash[len(hash)-16:])
```

For example, in the case of the Auth Claim the schema hash would be 

```golang
var sHash core.SchemaHash
h := Keccak256(schemaBytes, []byte("AuthBJJCredential"))
copy(sHash[:], h[len(h)-16:])
// sHash = ca938857241db9451ea329256b9c06e5
```
