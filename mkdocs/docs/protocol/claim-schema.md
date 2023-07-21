# Claim Schema 

In order to reuse claims across different services is necessary to keep consistent data formatting. A Claim Schema encodes the structure of a particular claim by defining a type, the fields that must be included inside a claim, and a description for these fields.

Schemas are described via JSON-LD documents. A claim issuer could reuse existing claim schemas or create new ones from scratch.

## Example: KYCAgeCredential Schema 

- [Github Document](https://github.com/iden3/claim-schema-vocab/blob/main/schemas/json-ld/kyc-v3.json-ld)

```json
{
  "@context": [
    {
      "@version": 1.1,
      "@protected": true,
      "id": "@id",
      "type": "@type",
      "KYCAgeCredential": {
        "@id": "https://raw.githubusercontent.com/iden3/claim-schema-vocab/main/schemas/json-ld/kyc-v3.json-ld#KYCAgeCredential",
        "@context": {
          "@version": 1.1,
          "@protected": true,
          "id": "@id",
          "type": "@type",
          "kyc-vocab": "https://github.com/iden3/claim-schema-vocab/blob/main/credentials/kyc.md#",
          "xsd": "http://www.w3.org/2001/XMLSchema#",
          "birthday": {
            "@id": "kyc-vocab:birthday",
            "@type": "xsd:integer"
          },
          "documentType": {
            "@id": "kyc-vocab:documentType",
            "@type": "xsd:integer"
          }
        }
      },
    }
  ]
}
```

This document describes the schema for a claim of type `KYCAgeCredential`.
The `@id` contains the unique url that contains the JSON-LD Document.
The `kyc-vocab` contains a link to a url with a vocabulary description of the value types stored inside this claim, in this case `birthday` and `documentType`.
The last part of the document contains a reference to the value types `birthday` and `documentType`. Their `@id` corresponds to their description in the kyc-vocab while the `@type` indicates where the data type of each field. In this case, `birthday` is an integer and `documentType` is also an integer.

## Schema Hash

The Schema Hash is a unique identifier for a Claim Schema. It is derived by hashing the string that represents unique identifier `@id` of the Claim Schema type. In the previous example, the hash pre-image is the string `https://raw.githubusercontent.com/iden3/claim-schema-vocab/main/schemas/json-ld/kyc-v3.json-ld#KYCAgeCredential`.

For example, in the case of the Auth Claim the schema hash would be 

```golang
var sHash core.SchemaHash
h := Keccak256([]byte("https://raw.githubusercontent.com/iden3/claim-schema-vocab/main/schemas/json-ld/kyc-v3.json-ld#KYCAgeCredential"))
copy(sHash[:], h[len(h)-16:])
sHashHex, _ := sHash.MarshalText()

fmt.Println(string(sHashHex))
// f21e8faf5c95292b6bfbc53f8143a9d4
```