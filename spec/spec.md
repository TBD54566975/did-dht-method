The DID DHT Method Specification 1.0
==================

**Specification Status**: Working Draft

**Latest Draft:** [https://did-dht.com](https://did-dht.com)

**Registry:** [https://did-dht.com/registry](https://did-dht.com/registry)

**Draft Created:** October 20, 2023

**Latest Update:** March 22, 2024

**Editors:**
~ [Gabe Cohen](https://github.com/decentralgabe)
~ [Daniel Buchner](https://github.com/csuwildcat)

**Contributors:**
~ [Moe Jangda](https://github.com/mistermoe)
~ [Frank Hinek](https://github.com/frankhinek)

**Participate:**
~ [GitHub repo](https://github.com/TBD54566975/did-dht-method)
~ [File a bug](https://github.com/TBD54566975/did-dht-method/issues)
~ [Commit history](https://github.com/TBD54566975/did-dht-method/commits/main)

## Abstract

A DID Method [[spec:DID-CORE]] based on [[ref:Pkarr]] and [[ref:Mainline DHT]], identified by the prefix `did:dht`.

<style id="protocol-stack-styles">
  #protocol-stack-styles + table {
    display: table;
    width: 400px;
    border-radius: 4px;
    box-shadow: 0 1px 3px -1px rgb(0 0 0 / 80%);
    overflow: hidden;
  }
  #protocol-stack-styles + table tr, #protocol-stack-styles + table td {
    border: none;
  }
  #protocol-stack-styles + table tr {
    text-shadow: 0 1px 2px rgb(255 255 255 / 75%);
  }
  #protocol-stack-styles + table tr:nth-child(1) {
    background: hsl(149deg 100% 86%);
  }
  #protocol-stack-styles + table tr:nth-child(2) {
    background: hsl(0deg 100% 82%);
  }
  #protocol-stack-styles + table tr:nth-child(3) {
    background: hsl(198deg 100% 87%);
  }
</style>

:--------------------------------------------------------: |
DID DHT                                                    |
[Pkarr](https://github.com/nuhvi/pkarr)                    |
[Mainline DHT](https://en.wikipedia.org/wiki/Mainline_DHT) |

[[ref:Pkarr]] stands for **P**ublic **K**ey **A**ddressable **R**esource **R**ecords. [[ref:Pkarr]] makes use of 
[[ref:Mainline DHT]], specifically [[ref:BEP44]] to store signed mutable records. This DID method uses [[ref:Pkarr]] to 
store _[[ref:DID Documents]]_.

As [[ref:Pkarr]] states, [[def:Mainline]] is used for the following reasons:

1. It has a proven track record of 15 years.
2. It is the biggest DHT in existence with an estimated 10 million servers.
3. It is fairly generous in retaining data.
4. It has been implemented in most languages and is stable.

The syntax of the identifier and accompanying data model used by the protocol is conformant with the [[spec:DID-Core]]
specification and shall be registered with the [[spec:DID-Spec-Registries]].

## Conformance

The keywords MAY, MUST, MUST NOT, RECOMMENDED, SHOULD, and SHOULD NOT in this document are to be interpreted as described in [BCP 14](https://www.rfc-editor.org/info/bcp14) [[spec:RFC2119]] [[spec:RFC8174]] when, and only when, they appear in all capitals, as shown here.

## Terminology

[[def:Decentralized Identifier, Decentralized Identifier, DID, DIDs, DID Document, DID Documents]]
~ A [W3C specification](https://www.w3.org/TR/did-core/) describing an _identifier that enables verifiable, 
decentralized digital identity_. A DID identifier is associated with a JSON document containing cryptograhpic keys,
services, and other properties outlined in the specification.

[[def:DID Suffix, Suffix]]
~ The unique identifier string within a DID URI. For DID DHT the suffix is the [[ref:z-base-32]] encoded [[ref:Identity Key]].

[[def:Identity Key]]
~ An [Identity Key](#identity-key) is a [[ref:z-base-32]] encoded [[ref:Ed25519]] public key required to authenticate all records in 
[[ref:Mainline DHT]]. The encoded string comprises the [[ref:Suffix]] of `did:dht` identifier. This key is guaranteed to be present
in each `did:dht` document.

[[def:DID DHT Service]]
~ A service that provides a [[ref:DHT]] interface to the [[ref:Pkarr]] network, extended to support this [[ref:DID]] method.

[[def:Pkarr]]
~ An [open-source project](https://github.com/nuhvi/pkarr), based on [[ref:Mainline]] which aims to provide the 
_simplest possible streamlined integration between the Domain Name System and peer-to-peer overlay networks, 
enabling self-issued public keys to function as sovereign, publicly addressable domains_.

[[def:Mainline DHT, DHT, Mainline, Mainline Server]]
~ [Mainline DHT](https://en.wikipedia.org/wiki/Mainline_DHT) is the name given to the 
[Distributed Hash Table](https://en.wikipedia.org/wiki/Distributed_hash_table) used by the 
[BitTorrent protocol](https://github.com/bittorrent/bittorrent.org). It is a distributed system for storing and
finding data on a peer-to-peer network. It is based on [Kademlia](https://en.wikipedia.org/wiki/Kademlia) and is primarily
used to store and retrieve peer data. It is estimated to consistently have tens of millions of concurrent active users.

[[def:Gateway, Gateways, DID DHT, Bitcoin-anchored Gateway]]
~ A server that that facilitates DID DHT operations. The gateway may offer a set of APIs to interact with the DID DHT
method features such as providing [guaranteed retention](#retained-did-set), [historical resolution](#historical-resolution), 
and other features. Gateways can make themselves discoverable via a [[ref:Gateway Registry]].

[[def:Gateway Registry, Gateway Registries]]
~ A system used to aid in the discovery of [[ref:Gateways]]. One such system is the 
[registry provided by this specification]((registry/index.html#gateways)).

[[def:Client, Clients]]
~ A client is a piece of software that is responsible for generating a `did:dht` and submitting it to a 
[[ref:Mainline Server]] or [[ref:Gateway]]. Notably, a client has the ability to sign messages with the [[ref:Identity Key]].

[[def:Retained DID Set, Retained Set, Retention Set]]
~ The set of DIDs that a [[ref:Gateway]] is retaining and thus is responsible for republishing.

[[def:Retention Proof, Retention Proofs]]
~ A proof of work that is performed by the [[ref:DID]] controller to prove that they are still in control of the DID.
[[ref:Gateways]] use this proof to determine how long they should retain a DID.

[[def:Sequence Number, Sequence Numbers, Sequence]]
~ A sequence number, or `seq`, is a property of a mutable item as defined in [[ref:BEP44]]. It is a 64-bit integer that increases
in a consistent, unidirectional manner, ensuring that items are ordered sequentially. This specification requires that sequence 
numbers are [[ref:Unix Timestamps]] represented in seconds.

[[def:Unix Timestamp, Unix Timestamps]]
~ A value that measures time by the number of non-leap seconds that have elapsed since 00:00:00 UTC on 1 January 1970, the Unix epoch.

## DID DHT Method Specification

### Format

The format for `did:dht` conforms to the [[spec:DID Core]] specification. It consists of the `did:dht` prefix followed by 
the [[ref:Pkarr]] identifier. The [[ref:Pkarr]] identifier is a [[ref:z-base-32]]-encoded [[ref:Ed25519]] public key, which
we refer to as an [[ref:Identity Key]].

```text
did-dht-format := did:dht:<z-base-32-value>
z-base-32-value := [a-km-uw-z13-9]+
```

Alternatively, one can interpret the encoding rules as a series of transformation functions on the raw public key bytes:

```text
did-dht-format := did:dht:Z-BASE-32(raw-public-key-bytes)
```

### Identity Key

A unique property of DID DHT is its dependence on an a single non-rotatable key which we refer to as an _Identity Key_. This requirement
stems from [[ref:BEP44]], particularly within the _Mutable Items_ section:

> Mutable items can be updated, without changing their DHT keys. To authenticate that only the original publisher can update an item,
it is signed by a private key generated by the original publisher. The target ID mutable items are stored under is the SHA-1 hash of
the public key (as it appears in the put message).

This mechanism, as detailed in [[ref:BEP44]], ensures that all entries in the DHT are authenticated through a private key unique
to the initial publisher. Consequently, DHT records, including DID DHT Documents, are _independently verifiable_. This independence
implies that trust in a specific [[ref:Mainline]] or [[ref:Gateway]] server for providing unaltered messages is unnecessary. Instead,
all clients can, and should, verify messages themselves. This approach significantly mitigates risks associated with other DID methods,
where a compromised server or [DID resolver](https://www.w3.org/TR/did-core/#choosing-did-resolvers) might tamper with a [[ref:DID Document]]
which would be undecetable by a client.

Currently, [[ref:Mainline]] exclusively supports the [[ref:Ed25519]] key type. In turn, [[ref:Ed25519]] is required by DID DHT and is 
used to uniquely identify DID DHT Documents. DID DHT identifiers are formed by concatenating the `did:dht:` prefix with a [[ref:z-base-32]]
encoded Identity Key, which acts as its [[ref:suffix]]. Identity Keys always have the identifier `0` as both its Verification Method `id` and
JWK `kid` [[spec:RFC7517]]. While the system requires at least one [[ref:Ed25519]], a DID DHT Document can include any number of additional keys.
However, these additional key's types ****MUST**** be registered in the [Key Type Index](registry/index.html##key-type-index).

As a unique consequence of the requirement of the Identity Key, DID DHT Documents are able to be partially-resolved without contacting
[[ref:Maineline]] or [[ref:Gateway]] servers, though it is ****RECOMMENDED**** that deterministic resolution is only used as a fallback mechanism.
Similarly, the requirement of an Identity Key enables [interoperability with other DID methods](#interoperability-with-other-did-methods).

### DIDs as DNS Records

In this scheme, we encode the [[ref:DID Document]] as multiple [DNS TXT records](https://en.wikipedia.org/wiki/TXT_record).
Comprising a DNS packet [[spec:RFC1034]] [[spec:RFC1035]], which is then stored in the [[ref:DHT]].

| Name         | Type | TTL    | Rdata                                                        |
| ------------ | ---- | ------ | ------------------------------------------------------------ |
| _did.`<ID>`. | TXT  |  7200  | v=0;vm=k0,k1,k2;auth=k0;asm=k1;inv=k2;del=k2;srv=s0,s1,s2    |
| _k0._did.    | TXT  |  7200  | id=0;t=0;k=`<unpadded-b64url>`                               |
| _k1._did.    | TXT  |  7200  | id=1;t=1;k=`<unpadded-b64url>`                               |
| _k2._did.    | TXT  |  7200  | id=2;t=1;k=`<unpadded-b64url>`                               |
| _s0._did.    | TXT  |  7200  | id=domain;t=LinkedDomains;se=https://foo.com                 |
| _s1._did.    | TXT  |  7200  | id=dwn;t=DecentralizedWebNode;se=https://dwn.tbddev.org/dwn5 |

::: note
The recommended TTL value is 7200 seconds (2 hours), the default TTL for [[ref:Mainline]] records.
:::

- The root record's name ****MUST**** be of the form, `_did.<ID>.`, where `ID` is the [[ref:Pkarr]] identifier
associated with the DID (i.e. `did:dht:<ID>` becomes `_did.<ID>.`).

- The root record's value identifies the [property mapping](#property-mapping) for the document, along with
versioning information.

- The root record ****MUST**** include a version number. The version of the DNS packet representation for
the `did:dht` method is represented by a property `v=V` where `V` is a monotonically increasing version number.
The version number for this specification version is **0** (e.g. `v=0`). The version number is not present
in the corresponding DID Document.

- Additional records like `srv` (services), `vm` (verification methods), and verification relationships
(e.g., authentication, assertion, etc.) are represented as additional records in the format `<ID>._did.`.
These records contain the zero-indexed value of each `key` or `service` as attributes.

- All resource record names, aside from the root record, ****MUST**** end in `_did.`.

- The DNS packet ****MUST**** set the _Authoritative Answer_ flag since this is always an _Authoritative_ packet.

**Example Root Record**

| Name                                                       | Type | TTL    | Rdata                                                     
| ---------------------------------------------------------- | ---- | ------ | --------------------------------------------------------- 
| _did.o4dksfbqk85ogzdb5osziw6befigbuxmuxkuxq8434q89uj56uyy. | TXT  |  7200  | v=0;vm=k0,k1,k2;auth=k0;asm=k1;inv=k2;del=k2;srv=s0,s1,s2 

### Property Mapping

The following section describes mapping a [[ref:DID Document]] to a DNS packet.

- Resource names are aliased with zero-indexed values (e.g. `k0`, `k1`, `s0`, `s1`).

- Verification Methods, Verification Relationships, and Services are separated by a semicolon (`;`), while
values within each property are separated by a comma (`,`).

- Across all properties, distinct elements are separated by semicolons (`;`) while array elements are separated by
commas (`,`).

- Additional properties not defined by this specification ****MAY**** be represented in a [[ref:DID Document]] and
its corresponding DNS packet if the properties [are registered in the additional properties registry](registry/index.html#additional-properties).

The subsequent instructions serve as a reference for mapping DID Document properties to [DNS TXT records](https://en.wikipedia.org/wiki/TXT_record):

#### Identifiers

##### Controller

A [DID controller](https://www.w3.org/TR/did-core/#did-controller) ****MAY**** be present in a `did:dht` document.

If present, a DID controller ****MUST**** be represented as a `_cnt._did.` record in the form of a comma-separated
list of controller DID identifiers.

To ensure that the DID controller is authorized to make changes to the DID Document, the controller for the [[ref:Identity Key]] Verification Method ****MUST**** be contained within the controller property.

**Example Controller Record**

| Name       | Type | TTL  | Rdata            |
| ---------- | ---- | ---- | ---------------- |
| _cnt._did. | TXT  | 7200 | did:example:abcd |

##### Also Known As

A `did:dht` document ****MAY**** have multiple identifiers using the [alsoKnownAs](https://www.w3.org/TR/did-core/#also-known-as) property.

If present, alternate DID identifiers ****MUST**** be represented as `_aka_.did.` record in the form of a
comma-separated list of DID identifiers.

**Example AKA Record**

| Name       | Type | TTL  | Rdata                              |
| ---------- | ---- | ---- | ---------------------------------- |
| _aka._did. | TXT  | 7200 | did:example:efgh,did:example:ijkl  |

#### Verification Methods

- Each [Verification Method's](https://www.w3.org/TR/did-core/#verification-methods) **name** is represented 
as a `_kN._did.` record where `N` is the zero-indexed positional index of a given Verification Method (e.g. `_k0`, `_k1`).

- Each [Verification Method's](https://www.w3.org/TR/did-core/#verification-methods) **rdata** is represented by the form
`id=M;t=N;k=O` where `M` is the Verification Method's `id`, `N` is the index of the key's type from the
[key type index](registry/index.html#key-type-index), and `N` is the unpadded base64URL [[spec:RFC4648]] representation of
the public key.

  - Verification Method `id`s ****MAY**** be omitted. If omitted, they can be computed according to the 
  rules specified in the section on [representing keys](#representing-keys) when reconstructing the DID Document.

  - The [[ref:Identity Key]] ****MUST**** always be at index `_k0` with `id` `0`.

- [Verification Methods](https://www.w3.org/TR/did-core/#verification-methods) ****MAY**** have an _optional_ **controller** property
represented by `c=C` where `C` is the identifier of the verification method's controller (e.g. `t=N;k=O;c=C`). If omitted,
it is assumed that the controller of the Verification Method is the [[ref:Identity Key]].

::: note
Controllers are not cryptographically verified by [[ref:Gateways]] or this DID method. This means any DID may choose to list
a controller, even if there is no relationship between the identifiers. As such, DID controllers should be interrogated to 
assert the veracity of their relations.
:::

#### Verification Relationships

- Each [Verification Relationship](https://www.w3.org/TR/did-core/#verification-relationships) is represented as a part
of the root `_did.<ID>.` record (see: [Property Mapping](#property-mapping)).

The following table acts as a map between Verification Relationship types and their record name:

##### Verification Relationship Index

| Verification Relationship  | Record Name  |
| -------------------------- | ------------ |
| Authentication             | `auth`       |
| Assertion                  | `asm`        |
| Key Agreement              | `agm`        |
| Capability Invocation      | `inv`        |
| Capability Delegation      | `del`        |

The record data is uniform across [Verification Relationships](https://www.w3.org/TR/did-core/#verification-relationships),
represented as a comma-separated list of key references.

**Example Verification Relationship Records**

| Verification Relationship                          | Rdata in the Root Record                     |
|----------------------------------------------------|----------------------------------------------|
| "authentication": ["did:dht:example#0", "did:dht:example#HTsY9aMkoDomPBhGcUxSOGP40F-W4Q9XCJV1ab8anTQ"] | auth=0,HTsY9aMkoDomPBhGcUxSOGP40F-W4Q9XCJV1ab8anTQ |
| "assertionMethod": ["did:dht:example#0", "did:dht:example#HTsY9aMkoDomPBhGcUxSOGP40F-W4Q9XCJV1ab8anTQ"]| asm=0,HTsY9aMkoDomPBhGcUxSOGP40F-W4Q9XCJV1ab8anTQ  |
| "keyAgreement": ["did:dht:example#1"]              | agm=1                                        |
| "capabilityInvocation": ["did:dht:example#0"]      | inv=0                                        |
| "capabilityDelegation": ["did:dht:example#0"]      | del=0                                        |

#### Services

- Each [Service's](https://www.w3.org/TR/did-core/#services) **name** is represented as a `_sN._did.` record where `N` is
the zero-indexed positional index of the Service (e.g. `_s0`, `_s1`).

- Each [Service's](https://www.w3.org/TR/did-core/#services) **data** is represented with the form `id=M;t=N;se=O`
where `M` is the Service's ID, `N` is the Service's Type and `O` is the Service's URI.

  - Multiple service endpoints can be represented as an array (e.g. `id=dwn;t=DecentralizedWebNode;se=https://dwn.org/dwn1,https://dwn.org/dwn2`)

  - Additional properties ****MAY**** be present (e.g. `id=dwn;t=DecentralizedWebNode;se=https://dwn.org/dwn1;sig=1;enc=2`)

**Example Service Record**

| Name      | Type | TTL  | Rdata                                                    |
| --------- | ---- | ---- | -------------------------------------------------------- |
| _s0._did. | TXT  | 7200 | id=dwn;t=DecentralizedWebNode;se=https://example.com/dwn |

Each Service is represented as part of the root record (`_did.<ID>.`) as a list under the key `srv=<ids>` where `ids`
is a comma-separated list of all IDs for each Service.

#### Example

A sample transformation of a fully-featured DID Document to a DNS packet is exemplified as follows:

**DID Document**

```json
{
  "id": "did:dht:i9xkp8ddcbcg8jwq54ox699wuzxyifsqx4jru45zodqu453ksz6y",
  "controller": "did:example:abcd",
  "alsoKnownAs": ["did:example:efgh", "did:example:ijkl"],
  "verificationMethod": [
    {
      "id": "did:dht:i9xkp8ddcbcg8jwq54ox699wuzxyifsqx4jru45zodqu453ksz6y#0",
      "type": "JsonWebKey",
      "controller": "did:dht:i9xkp8ddcbcg8jwq54ox699wuzxyifsqx4jru45zodqu453ksz6y",
      "publicKeyJwk": {
        "kid": "0",
        "alg": "EdDSA",
        "crv": "Ed25519",
        "kty": "OKP",
        "x": "r96mnGNgWGOmjt6g_3_0nd4Kls5-kknrd4DdPW8qtfw"
      }
    },
    {
      "id": "did:dht:i9xkp8ddcbcg8jwq54ox699wuzxyifsqx4jru45zodqu453ksz6y#HTsY9aMkoDomPBhGcUxSOGP40F-W4Q9XCJV1ab8anTQ",
      "type": "JsonWebKey",
      "controller": "did:dht:i9xkp8ddcbcg8jwq54ox699wuzxyifsqx4jru45zodqu453ksz6y",
      "publicKeyJwk": {
        "kid": "HTsY9aMkoDomPBhGcUxSOGP40F-W4Q9XCJV1ab8anTQ",
        "alg": "ES256K",
        "crv": "secp256k1",
        "kty": "EC",
        "x": "KI0DPvL5cGvznc8EDOAA5T9zQfLDQZvr0ev2NMLcxDw",
        "y": "0iSbXxZo0jIFLtW8vVnoWd8tEzUV2o22BVc_IjVTIt8"
      }
    }
  ],
  "authentication": [
    "did:dht:i9xkp8ddcbcg8jwq54ox699wuzxyifsqx4jru45zodqu453ksz6y#0",
    "did:dht:i9xkp8ddcbcg8jwq54ox699wuzxyifsqx4jru45zodqu453ksz6y#HTsY9aMkoDomPBhGcUxSOGP40F-W4Q9XCJV1ab8anTQ"
  ],
  "assertionMethod": [
    "did:dht:i9xkp8ddcbcg8jwq54ox699wuzxyifsqx4jru45zodqu453ksz6y#0",
    "did:dht:i9xkp8ddcbcg8jwq54ox699wuzxyifsqx4jru45zodqu453ksz6y#HTsY9aMkoDomPBhGcUxSOGP40F-W4Q9XCJV1ab8anTQ"
  ],
  "capabilityInvocation": [
    "did:dht:i9xkp8ddcbcg8jwq54ox699wuzxyifsqx4jru45zodqu453ksz6y#0"
  ],
  "capabilityDelegation": [
    "did:dht:i9xkp8ddcbcg8jwq54ox699wuzxyifsqx4jru45zodqu453ksz6y#0"
  ],
  "service": [
    {
      "id": "did:dht:i9xkp8ddcbcg8jwq54ox699wuzxyifsqx4jru45zodqu453ksz6y#dwn",
      "type": "DecentralizedWebNode",
      "serviceEndpoint": ["https://example.com/dwn1", "https://example/dwn2"]
    }
  ]
}
```

**DNS Resource Records**

| Name         | Type | TTL   | Rdata                                                                              |
| ------------ | ---- | ----- | ---------------------------------------------------------------------------------- |
| _did.`<ID>`. | TXT  | 7200  | v=0;vm=k0,k1;auth=k0,k1;asm=k0,k1;inv=k0;del=k0;srv=s1                             |
| _cnt.did.    | TXT  | 7200  | did:example:abcd                                                                   |
| _aka.did.    | TXT  | 7200  | did:example:efgh,did:example:ijkl                                                  |
| _k0._did.    | TXT  | 7200  | t=0;k=afdea69c63605863a68edea0ff7ff49dde0a96ce7e9249eb7780dd3d6f2ab5fc             |
| _k1._did.    | TXT  | 7200  | t=1;k=AyiNAz7y-XBr853PBAzgAOU_c0Hyw0Gb69Hr9jTC3MQ8                                 |
| _s0._did.    | TXT  | 7200  | id=dwn;t=DecentralizedWebNode;se=https://example.com/dwn1,https://example.com/dwn2 |

### Operations

Entries to the [[ref:DHT]] require a signed record as per [[ref:BEP44]]. As such, the keypair used for the [[ref:Pkarr]]
identifier is also used to sign the [[ref:DHT]] record.

#### Create

To create a `did:dht`, the process is as follows:

1. Generate an [[ref:Ed25519]] keypair and encode the public key using the format provided in the [format section](#format).

2. Construct a conformant JSON representation of a [[ref:DID Document]].

    a. The document ****MUST**** include a [Verification Method](https://www.w3.org/TR/did-core/#verification-methods) with
    the [[ref:Identity Key]] encoded as a `publicKeyJwk` as per [[spec:RFC7517]] with an `id` of `0` and `type` of
    `JsonWebKey` defined by [[ref:VC-JOSE-COSE]].

    b. The document can include any number of other [core properties](https://www.w3.org/TR/did-core/#core-properties);
    always representing key material as a `JWK` as per [[spec:RFC7517]].

3. Map the output [[ref:DID Document]] to a DNS packet as outlined in [property mapping](#property-mapping).

4. Construct a [[ref:BEP44]] conformant mutable put message, as specified by [Pkarr](https://github.com/Nuhvi/pkarr/blob/main/design/relays.md#put).
  
   a. `v` ****MUST**** be set to a [[ref:bencoded]] DNS packet from the prior step.

   b. `seq` ****MUST**** be set to the current [[ref:Unix Timestamp]] in seconds.

5. Submit the result of to the [[ref:DHT]] via a [[ref:Pkarr]] relay, or a [[ref:Gateway]], with the identifier created in step 1.

This specification **does not** make use of JSON-LD. As such DID DHT Documents ****MUST NOT**** include the `@context` property.

#### Read

To read a `did:dht`, the process is as follows:

1. Take the [[ref:suffix]] of the DID, that is, the _[[ref:z-base-32]] encoded identifier key_, and submit it to a [[ref:Pkarr]]
relay or a [[ref:Gateway]].

2. Decode the resulting [[ref:BEP44]] response's `v` value using [[ref:bencode]].

3. Reverse the DNS [property mapping](#property-mapping) process and re-construct a conformant [[ref:DID Document]].

    a. Expand all identifiers (i.e. Verification Methods, Services, etc.) to their fully-qualified 
    form (e.g. `did:dht:uodqi99wuzxsz6yx445zxkp8ddwj9q54ocbcg8yifsqru45x63kj#0` 
    as opposed to `0` or `#0`, `did:dht:uodqi99wuzxsz6yx445zxkp8ddwj9q54ocbcg8yifsqru45x63kj#service-1` as opposed to `#service-1`).

::: note
As a fallback, if a `did:dht` value cannot be resolved via the network, it can be expanded to a conformant [[ref:DID Document]]
containing just the [[ref:Identity Key]].
:::

#### Update

Any valid [[ref:BEP44]] record written to the DHT is an update. As long as control of the
[[ref:Identity Key]] is retained any update is made possibly by signing and writing records with a unique incremental
[[ref:sequence number]] with [mutable items](https://www.bittorrent.org/beps/bep_0044.html).

It is ****RECOMMENDED**** that updates are infrequent, as caching of the DHT is highly encouraged.

#### Deactivate

To deactivate a [[ref:DID Document]], there are a couple options:

1. Let the DHT record expire and cease to publish it.

2. Publish a new DHT record where the `rdata` of the root DNS record is the string "deactivated."

| Name         | Type | TTL  | Rdata       |
| ------------ | ---- | ---- | ----------- |
| _did.`<ID>`. | TXT  | 7200 | deactivated |

::: note
If you have published your DID through a [[ref:Gateway]], you may need to contact the operator to have them remove the
record from their [[ref:Retained DID Set]].
:::

::: todo
[](https://github.com/TBD54566975/did-dht-method/issues/100)
Add guidance for rotating to a new DID after deactivation.
:::

### Type Indexing

Type indexing is an **OPTIONAL** feature that enables DIDs to become **discoverable**, by flagging themselves as being of
a particular type. Types are not included as a part of the DID Document, but rather as part of the DNS packet. This allows
for DIDs to be indexed by type by [[ref:Gateways]], and for DIDs to be resolved by type.

DIDs can be indexed by type by adding a `_typ._did.` record to the DNS packet. A DID may have **AT MOST** one type index
record. This record is a TXT record with the following format:

- The record **name** is represented as a `_typ._did.` record.

- The record **data** is represented with the form `id=0,1,2` where the value is a comma-separated list of integer types from
the [type index](#type-indexing).

**Example Type Index Record**

| Name       | Type | TTL  | Rdata     |
| ---------- | ---- | ---- | --------- |
| _typ._did. | TXT  | 7200 | id=0,1,2  |

Types can be found and registered in the [indexed types registry](registry/index.html#indexed-types).

::: note
Identifying entities through type-based indexing is a relatively unreliable practice. It serves
as an initial step in recognizing the identity linked to a [[ref:DID]]. To validate identity assertions more robustly,
it is essential to delve deeper, employing tools like verifiable credentials and the examination of related data.
:::

## Interoperability With Other DID Methods

As an **OPTIONAL** extension, some existing DID methods can leverage `did:dht` to expand their feature set. This
enhancement is most useful for DID methods that operate based on a single key and are compatible with the [[ref:Ed25519]]
key format. By adopting this optional extension, users can maintain their current DIDs without any changes. Additionally,
they gain the ability to add extra information to their DIDs. This is achieved by either publishing or retrieving
data from [[ref:Mainline]].

Interoperable DID methods ****MUST**** be registered in [this specification's registry](registry/index.html#interoperable-did-methods).

## Gateways

[[ref:Gateways]] serve as specialized servers, providing a range of DID-centric functionalities that extend
beyond the capabilities of a standard [[ref:Mainline DHT]] servers. This section elaborates on these unique features,
outlines the operational prerequisites for managing a gateway, and discusses various other facets, including the
optional integration of these gateways into a registry system.

::: note
[[ref:Gateways]] may choose to support interoperable methods in addition to `did:dht` as outlined in the
[section on interoperability](#interoperability-with-other-did-methods).
:::

### Discovering Gateways

As an **OPTIONAL** feature of the DID DHT Method, operators of a [[ref:Gateway]] may choose to make their server
discoverable through a [[ref:Gateway Registry]]. This feature allows for easy location through various internet-based 
discovery mechanisms. [[ref:Gateway Registries]] can vary in nature, encompassing a spectrum from centrally managed 
directories to diverse decentralized systems including databases, ledgers, or other structures.

One such registry is [provided by this specification](registry/index.html#gateways).

### Retained DID Set

A [[ref:Retained DID Set]] refers to the set of DIDs a [[ref:Gateway]] retains and republishes to the DHT. A
[[ref:Gateway]] may choose to surface additional [APIs](#gateway-api) based on this set, such as providing a
[type index](#type-indexing).

To safeguard equitable access to the resources of [[ref:Gateways]], which are publicly accessible and potentially
subject to [a high volume of requests](#rate-limiting), we suggest an ****OPTIONAL**** mechanism aimed at upholding
fairness in the retention and republishing of record sets by [[ref:Gateways]]. This mechanism, referred to as a
[[ref:Retention Proof]], requires clients to generate a proof value for write requests. This process guarantees that
the amount of work done by a client is proportional to the duration of data retention and republishing a [[ref:Gateway]]
performs. This mechanism enhances the overall reliability and effectiveness of [[ref:Gateways]] in managing requests.

#### Generating a Retention Proof

A [[ref:Retention Proof]] is a form of [Proof of Work](https://en.bitcoin.it/wiki/Proof_of_work) performed over a DID
identifier concatenated with the `retention` value of a given DID operation. The `retention` value is composed of a
hash value specified [by the gateway](registry/index.html#gateways), and a random
[nonce](https://en.wikipedia.org/wiki/Cryptographic_nonce) using the [SHA-256 hashing algorithm](https://en.wikipedia.org/wiki/SHA-2).
The resulting _Retention Proof Hash_ is used to determine the retention duration based on the number of leading zeros
of the hash, referred to as the _difficulty_, which ****MUST**** be no less than 26 bits of the 256-bit hash value.
The algorithm, in detail, is as follows:

1. Obtain a DID identifier and set it to `DID`.

2. Get the difficulty and recent hash from the server set to `DIFFICULTY` and `HASH`, respectively.

3. Generate a random 32-bit integer nonce value set to `NONCE`.

4. Compute the [SHA-256](https://en.wikipedia.org/wiki/SHA-2) hash over `ATTEMPT` where `ATTEMPT` = (`DID` + `HASH` + `NONCE`).

5. Inspect the result of `ATTEMPT`, and ensure it has >= `DIFFICULTY` bits of leading zeroes.

  a. If so, `ATTEMPT` = `RETENTION_PROOF`.

  b. Otherwise, regenerate `NONCE` and go to step 3.

6. Submit the `RETENTION_PROOF` to the [Gateway API](#register=or-update-a-did).

#### Managing the Retained DID Set

[[ref:Gateways]] adhering to the [[ref:Retention Set]] rules ****SHOULD**** sort DIDs they are retaining by the number of
_leading 0s_ in their [[ref:Retention Proofs]] in descending order, followed by the block hash's index number in
descending order. When a [[ref:Gateway]] needs to reduce its [[ref:retained set]] of DID entries, it ****SHOULD****
remove entries from the bottom of the list following this sort.

#### Reporting on Retention Status

[[ref:Gateways]] ****MUST**** include the approximate time until retention fall-off in the
[DID Resolution Metadata](https://www.w3.org/TR/did-core/#did-resolution-metadata) of a resolved
[[ref:DID Document]], to aid [[ref:clients]] in being able to assess whether resubmission is required.

:::todo
[](https://github.com/TBD54566975/did-dht-method/issues/74)
Specify how gateways can report on retention guarantees and provide guidance for clients to work with these guarantees.
:::

### Gateway API

At a minimum, a gateway ****MUST**** support the [Relay API defined by Pkarr](https://github.com/Nuhvi/pkarr/blob/main/design/relays.md).

Expanding on this API, a conforming [[ref:Gateway]] ****MUST**** support the following API, which is also made 
available via an [OpenAPI document](#open-api-definition).

#### Get the Current Difficulty

Difficulty is exposed as an **OPTIONAL** endpoint based on support of [retention proofs](#retained-did-set).

- **Method:** `GET`
- **Path:** `/difficulty`
- **Returns:**
  - `200` - Success.
    - `hash` - **string** - **REQUIRED** - The current hash.
    - `difficulty` - **integer** - **REQUIRED** - The current difficulty.
  - `501` - Retention proofs not supported by this gateway.

```json
{
  "hash": "000000000000000000022be0c55caae4152d023dd57e8d63dc1a55c1f6de46e7",
  "difficulty": 26
}
```

#### Register or Update a DID

- **Method:** `PUT`
- **Path:** `/did/:id`
  - `id` - **string** - **REQUIRED** - ID of the DID to publish.
    - `did` - **string** - **REQUIRED** - The DID to register or update.
    - `sig` - **string** - **REQUIRED** - An unpadded base64URL-encoded signature of the [[ref:BEP44]] payload.
    - `seq` - **integer** - **REQUIRED** - A [[ref:sequence number]] for the request. This number ****MUST**** be unique for each DID operation,
    which ****MUST**** be a [[ref:Unix Timestamp]] in seconds.
    - `v` - **string** - **REQUIRED** - An unpadded base64URL-encoded [[ref:bencoded]] DNS packet containing the DID Document.
    - `retention_proof` - **string** - **OPTIONAL** - A retention proof calculated according to the [retention proof algorithm](#generating-a-retention-proof).
- **Returns:**
  - `202` - Accepted. The server has accepted the request as valid and will publish to the DHT.
  - `400` - Invalid request.
  - `401` - Invalid signature.
  - `409` - DID already exists with a higher [[ref:sequence number]]. DID may be accepted if the [[ref:Gateway]] supports [historical resolution](#historical-resolution).

```json
{
  "did": "did:dht:example",
  "sig": "<unpadded-base64URL-encoded-signature>",
  "seq": 1234,
  "v": "<unpadded-base64URL-encoded bencoded DNS packet>"
}
```

Upon receiving a request to register a DID, the Gateway ****MUST**** verify the signature of the request and if valid
publish the DID Document to the DHT. If the DNS Packets contain a `_typ._did.` record, the [[ref:Gateway]] ****MUST**** index the
DID by its type.

#### Resolving a DID

- **Method:** `GET`
- **Path:** `/did/:id`
  - `id` - **string** - **REQUIRED** - ID of the DID to resolve.
- **Returns:**
  - `200` - Success.
    - `did` - **object** - **REQUIRED** - A JSON object representing the DID Document.
    - `pkarr` - **string** - **REQUIRED** - An unpadded base64URL-encoded representation of the full [[ref:Pkarr]] payload, represented as 64 bytes sig,
    8 bytes u64 big-endian seq, and 0-1000 bytes of v concatenated, enabling independent verification.
    - `types` - **array** - **OPTIONAL** - An array of [type integers](#type-indexing) for the DID.
    - `sequence_numbers` - **array** - **OPTIONAL** - An sorted array of seen [[ref:sequence numbers]], used with [historical resolution](#historical-resolution).
    - `400` - Invalid request.
  - `404` - DID not found.

```json
{
  "did": {
    "id": "did:dht:i9xkp8ddcbcg8jwq54ox699wuzxyifsqx4jru45zodqu453ksz6y",
    "verificationMethod": [
      {
        "id": "did:dht:i9xkp8ddcbcg8jwq54ox699wuzxyifsqx4jru45zodqu453ksz6y#0",
        "type": "JsonWebKey",
        "controller": "did:dht:i9xkp8ddcbcg8jwq54ox699wuzxyifsqx4jru45zodqu453ksz6y",
        "publicKeyJwk": {
          "kid": "0",
          "alg": "EdDSA",
          "crv": "Ed25519",
          "kty": "OKP",
          "x": "r96mnGNgWGOmjt6g_3_0nd4Kls5-kknrd4DdPW8qtfw"
        }
      }
    ],
    "authentication": [
      "did:dht:i9xkp8ddcbcg8jwq54ox699wuzxyifsqx4jru45zodqu453ksz6y#0"
    ],
    "assertionMethod": [
      "did:dht:i9xkp8ddcbcg8jwq54ox699wuzxyifsqx4jru45zodqu453ksz6y#0"
    ]
  },
  "pkarr": "<unpadded-base64URL-encoded pkarr payload as [sig][seq][v]>",
  "types": [1, 4],
  "sequence_numbers": [1700356854, 1700461736]
}
```

Upon receiving a request to resolve a DID, the [[ref:Gateway]] ****MUST**** query the DHT for the [[ref:DID Document]], and if found,
return the document. If the records are not found in the DHT, the [[ref:Gateway]] ****MAY**** fall back to its local storage.
If the DNS Packets contain a `_typ._did.` record, the [[ref:Gateway]] ****MUST**** return the type index.

This API is returns a `pkarr` property which matches the payload of a [Pkarr Get Request](https://github.com/Nuhvi/pkarr/blob/main/design/relays.md#get),
when encoded as an unpadded base64URL string. Implementers are ****RECOMMENDED**** to verify the integrity of the response using
the `pkarr` data and reconstruct the DID Document themselves. The `did` property is provided as a utility which, without independent verification,
****SHOULD NOT**** be trusted.

##### Historical Resolution

[[ref:Gateways]] ****MAY**** choose to support historical resolution, which is to surface different versions of the 
same [[ref:DID Document]], sorted by [[ref:sequence number]], according to the rules set out in the section on 
[conflict resolution](#conflict-resolution).

Upon [resolving a DID](#resolving-a-did), the Gateway will return the parameter `sequence_numbers` if there exists
historical state for a given [[ref:DID]]. The following API can be used with specific [[ref:sequence numbers]] to fetch historical state:

- **Method:** `GET`
- **Path:** `/did/:id?seq=:sequence_number`
  - `id` - **string** - **REQUIRED** - ID of the DID to resolve
  - `seq` - **integer** - **OPTIONAL** - [[ref:Sequence number]] of the DID to resolve
- **Returns**:
  - `200` - Success.
    - `did` - **object** - **REQUIRED** - A JSON object representing the DID Document.
    - `pkarr` - **string** - **REQUIRED** - An unpadded base64URL-encoded representation of the full [[ref:Pkarr]] payload, represented as 64 bytes sig,
    8 bytes u64 big-endian seq, and 0-1000 bytes of v concatenated, enabling independent verification.
    - `types` - **array** - **OPTIONAL** - An array of [type integers](#type-indexing) for the DID.
  - `400` - Invalid request.
  - `404` - DID not found for the given [[ref:sequence number]].
  - `501` - Historical resolution not supported by this gateway.

#### Deactivating a DID

To intentionally deactivate a DID, as opposed to letting the record cease being published to the DHT, a DID controller
follows the same process as [updating a DID](#register-or-update-a-did), but with a record format outlined in the
[section on deactivation](#deactivate).

Upon receiving a request to deactivate a DID, the Gateway ****MUST**** verify the signature of the request, and if valid,
stop republishing the DHT. If the DNS Packets contain a `_typ._did.` record, the Gateway ****MUST**** remove the type index.

#### Type Indexing

**Get Info**

- **Method:** `GET`
- **Path:** `/did/types`
- **Returns:**
  - `200` - Success.
    - **array** - An array of objects describing the known types of the following form:
      - `type` - **integer** - **REQUIRED** - An integer representing the [type](#type-indexing).
      - `description` - **string** - **REQUIRED** - A string describing the [type](#type-indexing).
  - `404` - Type indexing not supported.

```json
[
  {
    "type": 1,
    "description": "Organization"
  },
  {
    "type": 7,
    "description": "Financial Institution"
  }
]
```

**Get a Specific Type**

- **Method:** `GET`
- **Path:** `/did/types/:id`
  - `id` - **integer** - **REQUIRED** - The type to query from the index.
  - `offset` - **integer** - **OPTIONAL** - Specifies the starting position from where the type records should be retrieved (Default: `0`).
  - `limit` - **integer** - **OPTIONAL** - Specifies the maximum number of type records to retrieve (Default: `100`).
- **Returns:**
  - `200` - Success.
    - **array** - **REQUIRED** - An array of DID Identifiers matching the associated type.
  - `400` - Invalid request.
  - `404` - Type not found.
  - `501` - Types not supported by this gateway.

```json
[
  "did:dht:i9xkp8ddcbcg8jwq54ox699wuzxyifsqx4jru45zodqu453ksz6y",
  "did:dht:uodqi99wuzxsz6yx445zxkp8ddwj9q54ocbcg8yifsqru45x63kj"
]
```

A query to the type index returns an array of DIDs matching the associated type. If the type is not found, a `404` is
returned. If no DIDs match the type, an empty array is returned.

## Implementation Considerations

### Conflict Resolution

Per [[ref:BEP44]], [[ref:Gateway]] servers can leverage the `seq` [[ref:sequence number]] to handle conflicts:

> [[ref:Gateways]] receiving a put request where `seq` is lower than or equal to what's already stored on the server,
****MUST**** reject the request. If the [[ref:sequence number]] is equal, and the value is also the same, the server
****SHOULD**** reset its timeout counter.

When the [[ref:sequence number]] is equal, but the value is different, servers need to decide which value to accept and which
to reject. To make this determination [[ref:Gateways]] ****MUST**** compare the payloads lexicographically to determine a
[lexicographical order](https://en.wikipedia.org/wiki/Lexicographic_order), and reject the payload with a **lower**
lexicographical order.

### Size Constraints

[[ref:BEP44]] payload sizes are limited to 1000 bytes. Accordingly, we specify [an efficient representation of a
DID Document](#dids-as-dns-records) and leveraged DNS packet encoding to optimize our payload sizes. With this
encoding format, we recommend additional considerations to minimize payload sizes:

#### Representing Keys

The following guidance on representations of keys and their identifiers using the `JsonWebKey` type defined by
[[ref:VC-JOSE-COSE]] are ****REQUIRED****:

- For the [[ref:Identity Key]], both the Verification Method `id` and JWK `id` ****MUST**** be set to `0`.

- For all other Verification Methods, JWK identifiers (`kid`s) ****MUST**** be set to the key's JWK Thumbprint
as per [[spec:RFC7638]].

- If no Verification Method `id` is provided, the Verification Method `id` is set to the JWK's `kid` value.

- [[ref:DID Document]] representations of elliptic curve (EC) keys ****MUST**** include the x- and y-coordinate pair.
To conserve space in the DNS packet representation, compressed point encoding ****MUST**** be used to transmit the
x-coordinate and a sign bit for the y-coordinate. This practice reduces each public key's size from 65 to 33 bytes.

- [[ref:DID Document]] representations ****MUST**** always use fully-qualified identifiers when referring 
to Verification Methods (e.g. `did:dht:uodqi99wuzxsz6yx445zxkp8ddwj9q54ocbcg8yifsqru45x63kj#0` as opposed to `0` or `#0`).

#### Historical Key State

Rotating keys is a widely recommended security practice. However, if you frequently rotate keys in a
[[ref:DID Document]], this can lead to an increase in the document's size due to the accumulation of old keys.
This, in turn, can enlarge the size of the corresponding DNS packet. To manage this issue, while still distinguishing
between currently active keys and those that are no longer in use (but were valid in the past), users ****MAY****
utilize the [service property](https://www.w3.org/TR/did-core/#services). This property allows for the specification 
of services that are dedicated to storing signed records of the historical key states. By doing this, it helps to keep 
the [[ref:DID Document]] more concise.

### Republishing Data

[[ref:Mainline]] offers a limited duration (approximately 2 hours) for retaining records in the DHT. To ensure the
verifiability of data signed by a [[ref:DID]], consistent republishing of [[ref:DID Document]] records is crucial. To
address this, it is ****RECOMMENDED**** to use [[ref:Gateways]] equipped with [[ref:Retention Proofs]] support.

### Rate Limiting

To reduce the risk of [Denial of Service Attacks](https://www.cisa.gov/news-events/news/understanding-denial-service-attacks),
spam, and other unwanted traffic, it is ****RECOMMENDED**** that [[ref:Gateways]] require [[ref:Retention Proofs]]. The
use of [[ref:Retention Proofs]] can act as an attack prevention measure, as it would be costly to scale retention proof
calculations. [[ref:Gateways]] ****MAY**** choose to explore other rate limiting techniques, such as IP-limiting, or an
access-token-based approach.

### DID Resolution

The process for resoloving a DID DHT Document via a [[ref:Gateway]] is outlined in the [read section above](#read).
However, we provide additional guidance for [DID Resolvers](https://www.w3.org/TR/did-core/#dfn-did-resolvers) supplying
[DID Document Metadata](https://www.w3.org/TR/did-core/#did-document-metadata) as follows:

* The metadata's [`versionId` property](https://www.w3.org/TR/did-core/#dfn-versionid) ****MUST**** be set to the
[[ref:DID Document]]'s current [[ref:sequence number]].

* The metadata's [`created` property](https://www.w3.org/TR/did-core/#dfn-created) ****MUST**** be set to 
[XML Datetime](https://www.w3.org/TR/xmlschema11-2/#dateTime) representation of the earliest known sequence number
for the DID.

* The metadata's [`updated` property](https://www.w3.org/TR/did-core/#dfn-updated) ****MUST**** be set to the
[XML Datetime](https://www.w3.org/TR/xmlschema11-2/#dateTime) representation of the last known sequence number
for the DID.

* If the [[ref:DID Document]] has [been deactivated](#deactivate) the 
[`deactivated` property](https://www.w3.org/TR/did-core/#dfn-deactivated) ****MUST**** be set to `true`.

## Security and Privacy Considerations

When implementing and using the `did:dht` method, there are several security and privacy considerations to be aware of
to ensure expected and legitimate behavior.

### Data Conflicts

Malicious actors may try to force [[ref:Gateways]] into uncertain states by manipulating the [[ref:sequence number]] associated
with a record set. There are three such cases to be aware of:

- **Low Sequence Number** - If a [[ref:Gateway]] has yet to see [[ref:sequence numbers]] for a given record it ****MUST**** query
its peers to see if they have encountered the record. If a peer is found who has encountered the record, the record
with the latest sequence number must be selected. If the server has encountered greater [[ref:sequence numbers]] before, the
server ****MAY**** reject the record set. If the server supports [historical resolution](#historical-resolution) it
****MAY**** choose to accept the request and insert the record into its historical ordered state.

- **Conflicting Sequence Number** - When a malicious actor publishes _valid but conflicting_ records to two different
[[ref:Mainline Servers]] or [[ref:Gateways]]. Implementers are encouraged to follow the guidance outlined in [conflict
resolution](#conflict-resolution).

- **High Sequence Number** - Since [[ref:sequence numbers]] ****MUST**** be second representations of a [[ref:Unix Timestamp]],
it is ****RECOMMENDED**** that [[ref:Gateways]] reject [[ref:sequence numbers]] that represent timestamps greater than 
**2 hours** into the future to mitigate [timing attack](#data-conflicts) risks.

### Data Availability

Given the nature of decentralized distributed systems, there are no firm guarantees that all [[ref:Gateways]] have access
to the same state. It is ****RECOMMENDED**** to publish and read from multiple [[ref:Gateways]] to reduce such risks.
As an **optional** enhancement [[ref:Gateways]] ****MAY**** choose to share state amongst themselves via mechanisms
such as a [gossip protocol](https://en.wikipedia.org/wiki/Gossip_protocol).

### Data Authenticity

To enter into the DHT using [[ref:BEP44]] records ****MUST*** be signed by an [[ref:Ed25519]] private key, known as the
[[ref:Identity Key]]. When retrieving records either through a [[ref:Mainline Server]] or a [[ref:Gateway]] is it 
****RECOMMENDED**** that one verifies the cryptographic integrity of the record themselves instead of trusting a server
to have done the validation. Servers that do not return a signature value ****MUST NOT**** be trusted.

### Key Compromise

Since the `did:dht` uses a single, un-rotatable root key, there is a risk of root key compromise. Such a compromise
may be tough to detect without external assurances of identity. Implementers are encouraged to be aware of this
possibility and devise strategies that support entities transitioning to new [[ref:DIDs]] over time.

### Public Data

[[ref:Mainline]] is a public network. As such, there is risk in storing private, sensitive, or personally identifying
information (PII) on such a network. Storing such sensitive information on the network or in the contents of a `did:dht`
document is strongly discouraged.

### Data Retention

It is ****RECOMMENDED**** that [[ref:Gateways]] implement measures supporting the "[Right to be
Forgotten](https://en.wikipedia.org/wiki/Right_to_be_forgotten)," enabling precise control over the data retention duration.

### Cryptographic Risk

The security of data within the [[ref:Mainline DHT]] which relies on mutable records using [[ref:Ed25519]] keys—is
intrinsically tied to the strength of these keys and their underlying algorithms, as outlined in [[spec:RFC8032]].
Should vulnerabilities be discovered in [[ref:Ed25519]] or if advancements in quantum computing compromise its
cryptographic foundations, the [[ref:Mainline]] method could become obsolete.

## Appendix

### Test Vectors

#### Vector 1

A minimal DID Document.

**Identity Public Key JWK:**

```json
{
  "kty": "OKP",
  "crv": "Ed25519",
  "x": "YCcHYL2sYNPDlKaALcEmll2HHyT968M4UWbr-9CFGWE",
  "alg": "EdDSA",
  "kid": "0"
}
```

**DID Document:**

```json
{
  "id": "did:dht:cyuoqaf7itop8ohww4yn5ojg13qaq83r9zihgqntc5i9zwrfdfoo",
  "verificationMethod": [
    {
      "id": "did:dht:cyuoqaf7itop8ohww4yn5ojg13qaq83r9zihgqntc5i9zwrfdfoo#0",
      "type": "JsonWebKey",
      "controller": "did:dht:cyuoqaf7itop8ohww4yn5ojg13qaq83r9zihgqntc5i9zwrfdfoo",
      "publicKeyJwk": {
        "kty": "OKP",
        "crv": "Ed25519",
        "x": "YCcHYL2sYNPDlKaALcEmll2HHyT968M4UWbr-9CFGWE",
        "alg": "EdDSA",
        "kid": "0"
      }
    }
  ],
  "authentication": [
    "did:dht:cyuoqaf7itop8ohww4yn5ojg13qaq83r9zihgqntc5i9zwrfdfoo#0"
  ],
  "assertionMethod": [
    "did:dht:cyuoqaf7itop8ohww4yn5ojg13qaq83r9zihgqntc5i9zwrfdfoo#0"
  ],
  "capabilityInvocation": [
    "did:dht:cyuoqaf7itop8ohww4yn5ojg13qaq83r9zihgqntc5i9zwrfdfoo#0"
  ],
  "capabilityDelegation": [
    "did:dht:cyuoqaf7itop8ohww4yn5ojg13qaq83r9zihgqntc5i9zwrfdfoo#0"
  ]
}
```

**DNS Resource Records:**

| Name      | Type | TTL  | Rdata                                                  |
| --------- | ---- | ---- | ------------------------------------------------------ |
| _did.cyuoqaf7itop8ohww4yn5ojg13qaq83r9zihgqntc5i9zwrfdfoo. | TXT  | 7200 | v=0;vm=k0;auth=k0;asm=k0;inv=k0;del=k0 |
| _k0._did. | TXT  | 7200 | id=0;t=0;k=YCcHYL2sYNPDlKaALcEmll2HHyT968M4UWbr-9CFGWE |

#### Vector 2

A DID Document with two keys ([[ref:Identity Key]] and an uncompressed secp256k1 key), a service with multiple endpoints, two types to index, an aka, and controller properties.

**Identity Public Key JWK:**

```json
{
  "kty": "OKP",
  "crv": "Ed25519",
  "x": "YCcHYL2sYNPDlKaALcEmll2HHyT968M4UWbr-9CFGWE",
  "alg": "EdDSA",
  "kid": "0"
}
```

**secp256k1 Public Key JWK:**

With controller: `did:dht:i9xkp8ddcbcg8jwq54ox699wuzxyifsqx4jru45zodqu453ksz6y`.

```json
{
  "kty": "EC",
  "crv": "secp256k1",
  "x": "1_o0IKHGNamet8-3VYNUTiKlhVK-LilcKrhJSPHSNP0",
  "y": "qzU8qqh0wKB6JC_9HCu8pHE-ZPkDpw4AdJ-MsV2InVY",
  "alg": "ES256K",
  "kid": "0GkvkdCGu3DL7Mkv0W1DhTMCBT9-z0CkFqZoJQtw7vw"
}
```

**Key Purposes:** Assertion Method, Capability Invocation.

**Service:**

```json
{
  "id": "service-1",
  "type": "TestService",
  "serviceEndpoint": ["https://test-service.com/1", "https://test-service.com/2"]
}
```

**Types:** 1, 2, 3.

**DID Document:**

```json
{
  "id": "did:dht:cyuoqaf7itop8ohww4yn5ojg13qaq83r9zihgqntc5i9zwrfdfoo",
  "controller": "did:example:abcd",
  "alsoKnownAs": ["did:example:efgh", "did:example:ijkl"],
  "verificationMethod": [
    {
      "id": "did:dht:cyuoqaf7itop8ohww4yn5ojg13qaq83r9zihgqntc5i9zwrfdfoo#0",
      "type": "JsonWebKey",
      "controller": "did:dht:cyuoqaf7itop8ohww4yn5ojg13qaq83r9zihgqntc5i9zwrfdfoo",
      "publicKeyJwk": {
        "kty": "OKP",
        "crv": "Ed25519",
        "x": "YCcHYL2sYNPDlKaALcEmll2HHyT968M4UWbr-9CFGWE",
        "alg": "EdDSA",
        "kid": "0"
      }
    },
    {
      "id": "did:dht:cyuoqaf7itop8ohww4yn5ojg13qaq83r9zihgqntc5i9zwrfdfoo#0GkvkdCGu3DL7Mkv0W1DhTMCBT9-z0CkFqZoJQtw7vw",
      "type": "JsonWebKey",
      "controller": "did:dht:i9xkp8ddcbcg8jwq54ox699wuzxyifsqx4jru45zodqu453ksz6y",
      "publicKeyJwk": {
        "kty": "EC",
        "crv": "secp256k1",
        "x": "1_o0IKHGNamet8-3VYNUTiKlhVK-LilcKrhJSPHSNP0",
        "y": "qzU8qqh0wKB6JC_9HCu8pHE-ZPkDpw4AdJ-MsV2InVY",
        "alg": "ES256K",
        "kid": "0GkvkdCGu3DL7Mkv0W1DhTMCBT9-z0CkFqZoJQtw7vw"
      }
    }
  ],
  "authentication": [
    "did:dht:cyuoqaf7itop8ohww4yn5ojg13qaq83r9zihgqntc5i9zwrfdfoo#0"
  ],
  "assertionMethod": [
    "did:dht:cyuoqaf7itop8ohww4yn5ojg13qaq83r9zihgqntc5i9zwrfdfoo#0",
    "did:dht:cyuoqaf7itop8ohww4yn5ojg13qaq83r9zihgqntc5i9zwrfdfoo#0GkvkdCGu3DL7Mkv0W1DhTMCBT9-z0CkFqZoJQtw7vw"
  ],
  "capabilityInvocation": [
    "did:dht:cyuoqaf7itop8ohww4yn5ojg13qaq83r9zihgqntc5i9zwrfdfoo#0",
    "did:dht:cyuoqaf7itop8ohww4yn5ojg13qaq83r9zihgqntc5i9zwrfdfoo#0GkvkdCGu3DL7Mkv0W1DhTMCBT9-z0CkFqZoJQtw7vw"
  ],
  "capabilityDelegation": [
    "did:dht:cyuoqaf7itop8ohww4yn5ojg13qaq83r9zihgqntc5i9zwrfdfoo#0"
  ],
  "service": [
    {
      "id": "did:dht:cyuoqaf7itop8ohww4yn5ojg13qaq83r9zihgqntc5i9zwrfdfoo#service-1",
      "type": "TestService",
      "serviceEndpoint": ["https://test-service.com/1", "https://test-service.com/2"]
    }
  ]
}
```

**DNS Resource Records:**

| Name      | Type | TTL  | Rdata       |
| --------- | ---- | ---- | ----------- |
| _did.cyuoqaf7itop8ohww4yn5ojg13qaq83r9zihgqntc5i9zwrfdfoo. | TXT | 7200 | v=0;vm=k0,k1;svc=s0;auth=k0;asm=k0,k1;inv=k0,k1;del=k0 |
| _cnt.did. | TXT  | 7200 | did:example:abcd                                                                                       |
| _aka.did. | TXT  | 7200 | did:example:efgh,did:example:ijkl                                                                      |
| _k0.did.  | TXT  | 7200 | id=0;t=0;k=YCcHYL2sYNPDlKaALcEmll2HHyT968M4UWbr-9CFGWE;c=did:example:abcd                              |
| _k1.did.  | TXT  | 7200 | t=1;k=Atf6NCChxjWpnrfPt1WDVE4ipYVSvi4pXCq4SUjx0jT9                                                |
| _s0.did.  | TXT  | 7200 | id=service-1;t=TestService;se=https://test-service.com/1,https://test-service.com/2                    |
| _typ.did. | TXT  | 7200 | id=1,2,3                                                                                               |

### Open API Definition

```yaml
[[insert: api.yaml]]
```

## References

[[def:Ed25519]]
~ [Ed25519](https://ed25519.cr.yp.to/). D. J. Bernstein, N. Duif, T. Lange, P. Schwabe, B.-Y. Yang; 26 September 2011.
[ed25519.cr.yp.to](https://ed25519.cr.yp.to/).

[[def:BEP44]]
~ [BEP44](https://www.bittorrent.org/beps/bep_0044.html). Storing arbitrary data in the DHT. A. Norberg, S. Siloti;
19 December 2014. [Bittorrent.org](https://www.bittorrent.org/).

[[def:Bencode, Bencoded]]
~ [Bencode](https://wiki.theory.org/BitTorrentSpecification#Bencoding). A way to specify and organize data in a terse
format. [Bittorrent.org](https://www.bittorrent.org/).

[[def:z-base-32]]
~ [z-base-32](https://philzimmermann.com/docs/human-oriented-base-32-encoding.txt). Human-oriented base-32 encoding.
Z. O'Whielacronx; November 2002.

[[def:VC-JOSE-COSE]]
~ [Securing Verifiable Credentials using JOSE and COSE](https://www.w3.org/TR/vc-jose-cose/). O. Steele, M. Jones,
M. Prorock, G. Cohen; 26 February 2024. [W3C](https://www.w3.org/).

[[spec]]
