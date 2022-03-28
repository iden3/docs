# Claims data structure

```mermaid
graph TD
Hi-->i0
Hi-->i1
Hi-->i2
Hi-->i3

Hv-->v0
Hv-->v1
Hv-->v2
Hv-->v3

Ht-->Hi
Ht-->Hv
```

## Common structure

The claim always contains a subject:
- **Self**: the claim says something about themself.  The subject is implicit, and it's the claiming identity.
- **OtherIden**: the claim says something about another identity by its ID.
- **Object**: the claim says something about an object by its ID.

If the subject is _Self_ -  _identity_ sections  i_1, v_1 can be empty. 

if the subject is NOT _Self_, the id(OtherIden) of the Identity/Object can be in the Index(i_1)
or the Value(v_1) part of the Claim.  This is encoded in a header bit.

```go
h_i = H(i_0, i_1, i_2, i_3)
h_v = H(v_0, v_1, v_2, v_3)
h_t = H(h_i, h_v)

Index:
 i_0: [ 128 bits ] claim schema
      [ 32 bits ] header flags
          [3] Subject:
            000: A.1 Self
            001: invalid
            010: A.2.i OtherIden Index
            011: A.2.v OtherIden Value
            100: B.i Object Index
            101: B.v Object Value
          [1] Expiration: bool
          [1] Updatable: bool
          [27] 0
      [ 32 bits ] version (optional?)
      [ 61 bits ] 0 - reserved for future use
 i_1: [ 248 bits] identity (case b) (optional)
      [  5 bits ] 0
 i_2: [ 253 bits] 0
 i_3: [ 253 bits] 0
Value:
 v_0: [ 64 bits ]  revocation nonce
         [ 64 bits ]  expiration date (optional)
         [ 125 bits] 0 - reserved
 v_1: [ 248 bits] identity (case c) (optional)
        [  5 bits ] 0
 v_2: [ 253 bits] 0
 v_3: [ 253 bits] 0
```

**Claim shema** - schemas define the kind of data inside a claim [link](../spec#claims)

_Idex slots_ **i_2**, **i_3** and _value slots_ **v_2**, **v_3** are data slots for user data

####Index VS Value 
When to use index slots and when value?

Claims are stored in the Merkle tree and the hash of index slots ( hash(i_0,i_1,i_2,i_3) ) is unique for the whole tree. It means that you can not have two claims with the same index inside the tree. 

As opposite to index, values slots could be the same for different claims if their index is different. 