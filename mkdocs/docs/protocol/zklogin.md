# Server-side login (ZKP Login)

## **Introduction**

Iden3 is SSI solution that allows users to leverage their pre-existing validated identities to prove they are who they claim to be based on zero-knowledge proofs. One of the direct applications of iden3’s technology is to allow web applications to reuse these identities for login into their portals.

![Login workflow](../../imgs/login.png)

Login workflow

In a simple example application request a user identifier, which is done through zero-knowledge proof (zk proof) generation.

The server generates an authentication request

_Auth request_

```json
    {
      "type": "https://iden3-communication.io/authorization-request/v1",
      "data": {
        "callbackUrl": "https://test.com/callbackurl",
        "audience": "1125GJqgw6YEsKFwj63GY87MMxPL9kwDKxPUiwMLNZ",
        "scope": [
            {
              "circuit_id": "auth",
              "type": "zeroknowledge",
              "rules": {
                  "challenge": 12345
              }
          }
       ]
      }
    }
```

This is an example of an authorization request. Scope field is a set of objects describes an array of proofs that must be generated on a user device and presented later.
Each scope member has a unique definition of a circuit that must be used and rules (public inputs) that must be applied.

This message can be delivered to user through different communication channels: QR code, email, deep-linking, .etc .
On scan mobile has to implement:

1. Parsing the auth request and understand which proof handler it should use
2. Resolve verifier identifier if it’s needed.
3. Generate proofs using a specific handler. It can be signature proof or zero-knowledge
4. Prepare an authentication response message.


On mobile user generates ZK proof using [auth](https://github.com/iden3/circuits/blob/master/circuits/authentication.circom) circuit that will prove identity ownership, and send the response to the callback URL


_Auth response_

```json
{
  "type": "https://iden3-communication.io/authorization-response/v1",
  "data": {
    "scope": [
      {
        "type": "zeroknowledge",
        "circuit_id": "auth",
        "pub_signals": [
          "371135506535866236563870411357090963344408827476607986362864968105378316288",
          "12345",
          "16751774198505232045539489584666775489135471631443877047826295522719290880931"
        ],
        "proof_data": {
          "pi_a": [
            "8286889681087188684411199510889276918687181609540093440568310458198317956303",
            "20120810686068956496055592376395897424117861934161580256832624025185006492545",
            "1"
          ],
          "pi_b": [
            [
              "8781021494687726640921078755116610543888920881180197598360798979078295904948",
              "19202155147447713148677957576892776380573753514701598304555554559013661311518"
            ],
            [
              "15726655173394887666308034684678118482468533753607200826879522418086507576197",
              "16663572050292231627606042532825469225281493999513959929720171494729819874292"
            ],
            [
              "1",
              "0"
            ]
          ],
          "pi_c": [
            "9723779257940517259310236863517792034982122114581325631102251752415874164616",
            "3242951480985471018890459433562773969741463856458716743271162635077379852479",
            "1"
          ],
          "protocol": "groth16"
        }
      }
    ]
  }
}
```

Client after receiving authorization response performs  verification procedure:

  1. Zero-knowledge proof verification
  2. Extraction of metadata: (auth and circuit-specific)
  3. Verification of user identity states
  4. Verification of circuits public inputs (e.g. issuer state)


## Authentication based on zero-knowledge proof

ZK proof is based on [Circom 2.0](https://docs.circom.io/) language.

Auth circuit repository: [https://github.com/iden3/circuits/blob/master/circuits/authentication.circom](https://github.com/iden3/circuits/blob/master/circuits/authentication.circom)

The circuit verifies that the user is the owner of the identity and his auth key is not revoked in the provided user state.

## Prerequisite

Identity wallet installed

## Integration

### Back-end

Generate auth request

```go
request := auth.CreateAuthorizationRequest("<challenge>","<verifier identity|app-url>", "<callbackURI>") // create auth request
```

Validate auth request

```go
// unpack raw message
message, err := packer.Unpack(msgBytes) 
// call library to verify zkp proofs
err = auth.VerifyProofs(message)  
// extract metadata
token, err := auth.ExtractMetadata(message)
// verify state
stateInfo, err := token.VerifyState(ctx.Background(),"< rpc url >", "< state contract address >")

```

In future releases of auth library the verification procedure will be simplified and optimized for verifier.

### Front-end

On the front-end side, you need to embed a button to start the login process. After the button is pressed, the front-end makes a request to the back-end to generate an authentication request and displays it in QR code. When the user scans the QR code, the phone generates ZK proof and sends the proof to the call-back URL from QR-code.
Currently, we are working on js-iden3-auth library.

## Tutorial simple go app

We need a simple web server with two endpoint

- GET /sign-in should return auth request
- POST /call-back endpoint to receive callback request from the phone and validate the request

Let’s write a simple web-server

```go
func main() {

	http.HandleFunc("/sign-in", signIn)
	http.HandleFunc("/call-back", callBack)

	http.ListenAndServe(":8001", nil)
}

func signIn(w http.ResponseWriter, req *http.Request) {
	
}

func callBack(w http.ResponseWriter, req *http.Request) {

}
```

### Auth package

Add authorization package to the project.

```go
go get https://github.com/iden3/go-iden3-auth
```

### Sign in

To generate a ZK auth request we need a callback URL, to this URL we will receive a response from the mobile application with an authentication response. And verifier identity [do we really need this identity?]

go-iden3-auth library contains a method for generating the authentication request

[Descrition]

```go
func CreateAuthorizationRequest(challenge int64, aud, callbackURL string) *types.AuthorizationMessageRequest
```

Now we are ready to generate auth request

```go
const CallBackUrl      = "http:localhost:8001/call-back"
const VerifierIdentity = "1125GJqgw6YEsKFwj63GY87MMxPL9kwDKxPUiwMLNZ"

func signIn(w http.ResponseWriter, req *http.Request) {
	
	request := auth.CreateAuthorizationRequest(10, VerifierIdentity, callBackURI)

	msgBytes, _ := json.Marshal(request) // error handling ommited for simplification

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(msgBytes)
}
```

### Callback

When we receive a callback request with ZK response we have to do a couple of things to validate this response

- Validate ZK proof, and make sure that the proof is valid
- Validate identity state on-chain, we are doing it to verify that user identity state is valid and his auth keys are not revoked

First, let’s validate ZK proof, for this we have a function

```go
func VerifyProofs(message types.Message) (err error)
```

but before we can call it we need to unpack raw bytes to a message. Packer can be used to process encrypted message in future releases.

```go
p := &packer.PlainMessagePacker{}
// Unpack msg
message, _ := p.Unpack(msgBytes)	
// verify zkp
proofErr := auth.VerifyProofs(message)
```

Now ZK proof is verified and we can check identity status on chain

Fist we need access to RPC URL, and address of identity smart-contract

```go
const rpc = "https://polygon-mumbai.infura.io/v3/<your-token>"
const IdentityContract = "0x3e95a7B12e8905E01126E1beA3F1a52D1624A725"
```

Before we can verify a state, we need to extract metadata, and then verify it on chain

```go
token, _:= auth.ExtractMetadata(message)
	
// verify match identifier with the state on chain
stateInfo, err := token.VerifyState(ctx, rpc, IdentityContract)
```

## Verification procedure details

### Zero-knowledge proof verification

> Groth16 proof are supported now by auth library
>

Verification keys for circuits are known by the library itself. In the future, they can be resolved from circuits registries.

### Extraction of metadata

Each circuit has a schema of its public inputs that links the public signal name to its position in the resulted array.

This allows extracting user identifiers and challenges for authentication from proof.

Other signals are added to the user token ( scope field) as attributes of a specific circuit.

Circuit public signals schemas are known by this library or can be retrieved from some registry.

### Verification of user identity states

The blockchain verification algorithm is  used

1. Get state from the blockchain (address of id state contract and URL must be provided by the caller of the library):
    1. Empty state is returned - it means that identity state hasn’t been updated or updated state hasn’t been published. We need to compare id and state. If they are different it’s not a genesis state of identity then it’s not valid.
    2. The non-empty state is returned and equals to the state in provided proof which means that the user state is fresh enough and we work with the latest user state.
    3. The non-empty state is returned and it’s not equal to the state that the user has provided. Gets the time of the state transition. The verification party can make a decision if it can accept this state based on that time frame

2.  Verification party can make a decision to accept provided state or not.

### Verification of circuits public signals

It can be different kind of verification e.g.
1. Check of that issuer states of provided claim proofs are published on the blockchain (same as for identity state)
2. Check of query signals, so claim schema and specific values can be verified.

### Full example



[Link to github](https://github.com/iden3/go-iden3-auth) === TODO: add proper link
