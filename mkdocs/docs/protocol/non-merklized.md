# Non-merklized credentials

## Motivation

While Iden3 protocol allows to have a root of  merklized JSON-LD document for proving the inclusion of the credential entries in the tree, that increases level of privacy and flexibility for the credential data there is also a need to issue a core claim with a data itself. This is useful as for onchain issuers that can’t work with ld schemas in smart contracts and also for use case when the actual data must be saved, e.g. Auth BJJ credential public key representation.

**Core Concept**

Non merklized credentials now can be created relying on JSON-LD schemas.

Example of such context: `https://schema.iden3.io/core/jsonld/auth.jsonld`

```json
{
  "@context": [{
    "@version": 1.1,
    "@protected": true,
    "id": "@id",
    "type": "@type",
    "AuthBJJCredential": {
      "@id": "https://schema.iden3.io/core/jsonld/auth.jsonld#AuthBJJCredential",
      "@context": {
        "@version": 1.1,
        "@protected": true,
        "id": "@id",
        "type": "@type",
        "iden3_serialization": "iden3:v1:slotIndexA=x&slotIndexB=y",
        "xsd": "http://www.w3.org/2001/XMLSchema#",
        "auth-vocab": "https://schema.iden3.io/core/vocab/auth.md#",
        "x": {
          "@id": "auth-vocab:x",
          "@type": "xsd:positiveInteger"
        },
        "y": {
          "@id": "auth-vocab:y",
          "@type": "xsd:positiveInteger"
        }
      }
    }
  }]
}
```

Schema is defined as a schema for non-merklized credentials by utilizing `iden3_serialization` attribute.

It contains a map string represented mapping, where:

`iden3:v1` - version of protocol serialization, constant.

`slotIndexA=x`  slotIndexA is index data slot  A with index 2 for path to field  `x` in credential

`slotIndexB=y`  slotIndexB is index data slot B  with index 3 for path to field `y` in credential

other possible values:

`slotValueA`  slotIndexB is value data slot  with index 6

`slotValueB`  slotValueB is value data slot  with index 7

See more information regarding data slot indexes [here](https://www.notion.so/Identity-Core-77fe2f04c8ad4e0296e2b90400d7ae3a?pvs=21). When user creates schema he should choose the slot to put the field.

Nested structures are supported and path are created using concatenation with `.` e.g for birthday field in the Index Data Slot A mapping entry looks like that: `"...slotIndexA=passportInfo.birthday"`

```json
{
  ...
			...
          "passportInfo": {
            "@id": "vocab:passportInfo",
            "@context": {
              "@version": 1.1,
              "@protected": true,
              "kyc-vocab": "https://github.com/iden3/claim-schema-vocab/blob/main/credentials/kyc.md#",
              "xsd": "http://www.w3.org/2001/XMLSchema#",
              "id": "@id",
              "type": "@type",
              "birthday": {
                "@type": "xsd:integer",
                "@id": "vocab:birthday"
              }
            }
          }
      ...
  ...
}
```

**Important**:

1. Fields in index slots make influence on the uniqueness of the claim in the clams tree of issuer, data in the value slots - don’t.
2. Data Slots number is 4, so there is a restriction to have only 4 fields for non-merklized credentials.

Meanwhile `@type` filed for each field must contain one of the supported primitive types, so value can be written according to the data type.

List of supported data types:

```jsx
XSD namespace {
  Boolean = 'http://www.w3.org/2001/XMLSchema#boolean',
  Integer = 'http://www.w3.org/2001/XMLSchema#integer',
  NonNegativeInteger = 'http://www.w3.org/2001/XMLSchema#nonNegativeInteger',
  NonPositiveInteger = 'http://www.w3.org/2001/XMLSchema#nonPositiveInteger',
  NegativeInteger = 'http://www.w3.org/2001/XMLSchema#negativeInteger',
  PositiveInteger = 'http://www.w3.org/2001/XMLSchema#positiveInteger',
  DateTime = 'http://www.w3.org/2001/XMLSchema#dateTime',
  Double = 'http://www.w3.org/2001/XMLSchema#double'
}
```

Libraries that support non-merklized credentials:
https://github.com/iden3/go-schema-processor/releases/tag/v2.0.1
https://github.com/0xPolygonID/js-sdk/releases/tag/v1.1.0

To create non-merklized credential in go-schema-processor / JS-sdk  merklized Root Position must be set to None (default value) and ld context must contain `iden3_serialization` attribute.

```go
processor.CoreClaimOptions{
		  ..,
			MerklizedRootPosition: verifiable.CredentialMerklizedRootPositionNone,
      ...
		}
```

If context contains serialization attribute but MerklizedRootPosition is set to Index / Value error will be thrown.

In case context doesn’t contain serialization attribute and MerklizedRootPosition is set to Index / Value. Merkle root will be written to corresponding position. If the MerklizedRootPosition is set to None Merkle root will be written to Index.
