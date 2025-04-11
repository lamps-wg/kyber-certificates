---
title: >
  Internet X.509 Public Key Infrastructure - Algorithm Identifiers
  for the Module-Lattice-Based Key-Encapsulation Mechanism (ML-KEM)
abbrev: ML-KEM in Certificates
category: std

docname: draft-ietf-lamps-kyber-certificates-latest
submissiontype: IETF
number:
date:
consensus: true
v: 3
area: SEC
workgroup: LAMPS
keyword:
  ML-KEM
  Kyber
  KEM
  Certificate
  X.509
  PKIX
venue:
  group: "Limited Additional Mechanisms for PKIX and SMIME (lamps)"
  type: "Working Group"
  mail: "spasm@ietf.org"
  arch: "https://mailarchive.ietf.org/arch/browse/spasm/"
  github: "lamps-wg/kyber-certificates"
  latest: "https://lamps-wg.github.io/kyber-certificates/#go.draft-ietf-lamps-kyber-certificates.html"

author:
 -
    name: Sean Turner
    organization: sn3rd
    email: sean@sn3rd.com
 -
    ins: P. Kampanakis
    name: Panos Kampanakis
    org: AWS
    email: kpanos@amazon.com
 -
    ins: J. Massimo
    name: Jake Massimo
    organization: AWS
    email: jakemas@amazon.com
 -
    ins: B. Westerbaan
    name: Bas Westerbaan
    organization: Cloudflare
    email: bas@westerbaan.name

normative:
  X680:
    target: https://www.itu.int/rec/T-REC-X.680
    title: >
      Information technology - Abstract Syntax Notation One (ASN.1):
      Specification of basic notation
    date: 2021-02
    author:
    -  org: ITU-T
    seriesinfo:
      ITU-T Recommendation: X.680
      ISO/IEC: 8824-1:2021
  X690:
    target: https://www.itu.int/rec/T-REC-X.690
    title: >
      Information technology - Abstract Syntax Notation One (ASN.1):
      ASN.1 encoding rules: Specification of Basic Encoding Rules (BER),
      Canonical Encoding Rules (CER) and Distinguished Encoding Rules (DER)
    date: 2021-02
    author:
    -  org: ITU-T
    seriesinfo:
      ITU-T Recommendation: X.690
      ISO/IEC: 8825-1:2021

informative:
  CDM23:
    title: "Keeping Up with the KEMs: Stronger Security Notions for KEMs and automated analysis of KEM-based protocols"
    target: https://eprint.iacr.org/2023/1933.pdf
    date: 2023
    author:
      -
        ins: C. Cremers
        name: Cas Cremers
        org: CISPA Helmholtz Center for Information Security
      -
        ins: A. Dax
        name: Alexander Dax
        org: CISPA Helmholtz Center for Information Security
      -
        ins: N. Medinger
        name: Niklas Medinger
        org: CISPA Helmholtz Center for Information Security
  KEMMY24:
    title: "Unbindable Kemmy Schmidt: ML-KEM is neither MAL-BIND-K-CT nor MAL-BIND-K-PK"
    target: https://eprint.iacr.org/2024/523.pdf
    date: 2024
    author:
      -
        ins: S. Schmieg
        name: Sophie Schmieg
  NIST-PQC:
    target: https://csrc.nist.gov/projects/post-quantum-cryptography
    title: >
      Post-Quantum Cryptography Project
    author:
    - org: National Institute of Standards and Technology (NIST)
    date: 2016-12-20

--- abstract

The Module-Lattice-Based Key-Encapsulation Mechanism (ML-KEM) is a
quantum-resistant key-encapsulation mechanism (KEM). This document
describes the conventions for using the ML-KEM in X.509 Public Key
Infrastructure. The conventions for the subject public keys and
private keys are also described.

--- middle

# Introduction

The Module-Lattice-Based Key-Encapsulation Mechanism (ML-KEM) standardized in
{{!FIPS203=DOI.10.6028/NIST.FIPS.203}} is a quantum-resistant
key-encapsulation mechanism (KEM) standardized by the US National Institute
of Standards and Technology (NIST) PQC Project {{NIST-PQC}}. Prior to
standardization, the earlier versions of the mechanism were known as
Kyber. ML-KEM and Kyber are not compatible. This document specifies the use
of ML-KEM in Public Key Infrastructure X.509 (PKIX) certificates {{!RFC5280}}
at three security levels: ML-KEM-512, ML-KEM-768, and ML-KEM-1024, using
object identifiers assigned by NIST. The private key format is also
specified.

## Applicability Statement

ML-KEM certificates are used in protocols where the public key is used to
generate and encapsulate a shared secret used to derive a symmetric key used
to encrypt a payload; see {{?I-D.ietf-lamps-cms-kyber}}. To be used in TLS,
ML-KEM certificates could only be used as end-entity identity certificates
and would require significant updates to the protocol; see
{{?I-D.celi-wiggers-tls-authkem}}.

# Conventions and Definitions

{::boilerplate bcp14-tagged}

# Algorithm Identifiers {#oids}

The `AlgorithmIdentifier` type is defined in {{!RFC5912}} as follows:

~~~
  AlgorithmIdentifier{ALGORITHM-TYPE, ALGORITHM-TYPE:AlgorithmSet} ::=
    SEQUENCE {
      algorithm   ALGORITHM-TYPE.&id({AlgorithmSet}),
      parameters  ALGORITHM-TYPE.
                    &Params({AlgorithmSet}{@algorithm}) OPTIONAL
    }
~~~

<aside markdown="block">
  NOTE: The above syntax is from {{!RFC5912}} and is compatible with the
  2021 ASN.1 syntax {{X680}}. See {{RFC5280}} for the 1988 ASN.1 syntax.
</aside>

The fields in `AlgorithmIdentifier` have the following meanings:

* `algorithm` identifies the cryptographic algorithm with an object
  identifier.

* `parameters`, which are optional, are the associated parameters for
  the algorithm identifier in the `algorithm` field.

The `AlgorithmIdentifier` for an ML-KEM public key MUST use one of the
`id-alg-ml-kem` object identifiers listed below, based on the security
level. The `parameters` field of the `AlgorithmIdentifier` for the ML-KEM
public key MUST be absent.

~~~
  nistAlgorithms OBJECT IDENTIFIER ::= { joint-iso-ccitt(2)
    country(16) us(840) organization(1) gov(101) csor(3)
    nistAlgorithm(4) }

  kems OBJECT IDENTIFIER ::= { nistAlgorithms 4 }

  id-alg-ml-kem-512 OBJECT IDENTIFIER ::= { kems 1 }

  id-alg-ml-kem-768 OBJECT IDENTIFIER ::= { kems 2 }

  id-alg-ml-kem-1024 OBJECT IDENTIFIER ::= { kems 3 }

~~~

# Subject Public Key Fields  {#pub-key}

In the X.509 certificate, the `subjectPublicKeyInfo` field has the
`SubjectPublicKeyInfo` type, which has the following ASN.1 syntax:

~~~
  SubjectPublicKeyInfo {PUBLIC-KEY: IOSet} ::= SEQUENCE {
      algorithm        AlgorithmIdentifier {PUBLIC-KEY, {IOSet}},
      subjectPublicKey BIT STRING
  }
~~~

<aside markdown="block">
  NOTE: The above syntax is from {{RFC5912}} and is compatible with the
  2021 ASN.1 syntax {{X680}}. See {{RFC5280}} for the 1988 ASN.1 syntax.
</aside>

The fields in `SubjectPublicKeyInfo` have the following meaning:

* `algorithm` is the algorithm identifier and parameters for the
  public key (see above).

* `subjectPublicKey` contains the byte stream of the public key.

The `PUBLIC-KEY` ASN.1 type for ML-KEM are defined here:

~~~
  pk-ml-kem-512 PUBLIC-KEY ::= {
    IDENTIFIER id-alg-ml-kem-512
    -- KEY no ASN.1 wrapping; 800 octets --
    PARAMS ARE absent
    CERT-KEY-USAGE { keyEncipherment }
    PRIVATE-KEY ML-KEM-512-PrivateKey -- defined in Section 6
    }

  pk-ml-kem-768 PUBLIC-KEY ::= {
    IDENTIFIER id-alg-ml-kem-768
    -- KEY no ASN.1 wrapping; 1184 octets --
    PARAMS ARE absent
    CERT-KEY-USAGE { keyEncipherment }
    PRIVATE-KEY ML-KEM-768-PrivateKey -- defined in Section 6
    }

  pk-ml-kem-1024 PUBLIC-KEY ::= {
    IDENTIFIER id-alg-ml-kem-1024
    -- KEY no ASN.1 wrapping; 1568 octets --
    PARAMS ARE absent
    CERT-KEY-USAGE { keyEncipherment }
    PRIVATE-KEY ML-KEM-1024-PrivateKey -- defined in Section 6
  }

  ML-KEM-512-PublicKey ::= OCTET STRING (SIZE (800))

  ML-KEM-768-PublicKey ::= OCTET STRING (SIZE (1184))

  ML-KEM-1024-PublicKey ::= OCTET STRING (SIZE (1568))
~~~

<aside markdown="block">
  NOTE: The above syntax is from {{!RFC5912}} and is compatible with the
  2021 ASN.1 syntax {{X680}}. See {{RFC5280}} for the 1988 ASN.1 syntax.
</aside>

When an ML-KEM public key appears outside of a `SubjectPublicKeyInfo`
type in an environment that uses ASN.1 encoding, it can be encoded
as an OCTET STRING by using the `ML-KEM-512-PublicKey`,
`ML-KEM-768-PublicKey`, `ML-KEM-1024-PublicKey` types corresponding to
the correct key size.

{{!RFC5958}} describes the Asymmetric Key Package's `OneAsymmetricKey`
type for encoding asymmetric keypairs. When an ML-KEM private key or
keypair is encoded as a `OneAsymmetricKey`, it follows the description
in {{priv-key}}.

When the ML-KEM private key appears outside of an Asymmetric Key Package
in an environment that uses ASN.1 encoding, it can be encoded using one
of the the `ML-KEM-PrivateKey CHOICE` formats defined in {{priv-key}}. The
`seed` format is RECOMMENDED as it efficiently stores both the private and
public key.

{{example-public}} contains examples for ML-KEM public keys
encoded using the textual encoding defined in {{?RFC7468}}.

# Key Usage Bits

The intended application for the key is indicated in the keyUsage certificate
extension; see {{Section 4.2.1.3 of RFC5280}}. If the `keyUsage` extension is
present in a certificate that indicates `id-alg-ml-kem-*` in the `SubjectPublicKeyInfo`,
then `keyEncipherment` MUST be the only key usage set.

# Private Key Format {#priv-key}

{{FIPS203}} specifies two formats for an ML-KEM private key: a 64-octet
seed and an (expanded) private key, which is referred to as the
decapsulation key. The expanded private key (and public key)
is computed from the seed using `ML-KEM.KeyGen_internal(d,z)` (algorithm 16)
using the first 32 octets as *d* and the remaining 32 octets as *z*.
If the expanded private key is generated without exporting the seed,
`ML-KEM.KeyGen()` (algorithm 19), which combines seed generation with
`ML-KEM.KeyGen_internal(d,z)`, is used.

A keypair is generated by sampling 64 octets uniformly at random
for the seed (private key) from a cryptographically secure
pseudorandom number generator (CSPRNGs). The public key can then
be computed using `ML-KEM.KeyGen_internal(d,z)` as described earlier.

"Asymmetric Key Packages" {{RFC5958}} describes how to encode a private
key in a structure that both identifies which algorithm the private key
is for and allows for the public key and additional attributes about the
key to be included as well. For illustration, the ASN.1 structure
`OneAsymmetricKey` is replicated below.

~~~
  OneAsymmetricKey ::= SEQUENCE {
    version                  Version,
    privateKeyAlgorithm      SEQUENCE {
    algorithm                PUBLIC-KEY.&id({PublicKeySet}),
    parameters               PUBLIC-KEY.&Params({PublicKeySet}
                               {@privateKeyAlgorithm.algorithm})
                                  OPTIONAL}
    privateKey               OCTET STRING (CONTAINING
                               PUBLIC-KEY.&PrivateKey({PublicKeySet}
                                 {@privateKeyAlgorithm.algorithm})),
    attributes           [0] Attributes OPTIONAL,
    ...,
    [[2: publicKey       [1] BIT STRING (CONTAINING
                               PUBLIC-KEY.&Params({PublicKeySet}
                                 {@privateKeyAlgorithm.algorithm})
                                 OPTIONAL,
    ...
  }
~~~

<aside markdown="block">
  NOTE: The above syntax is from {{RFC5958}} and is compatible with the
  2021 ASN.1 syntax {{X680}}.
</aside>

For ML-KEM private keys, the privateKey field in `OneAsymmetricKey` contains
one of the following DER-encoded `CHOICE` structures. The `seed`
format is a fixed 64 byte `OCTET STRING` (66 bytes total with the `0x8040`
tag and length) for all security levels, while the `expandedKey`
and `both` formats vary in size by security level:

~~~
  ML-KEM-512-PrivateKey ::= CHOICE {
    seed [0] OCTET STRING (SIZE (64)),
    expandedKey OCTET STRING (SIZE (1632)),
    both SEQUENCE {
      seed OCTET STRING (SIZE (64)),
      expandedKey OCTET STRING (SIZE (1632))
      }
    }

  ML-KEM-768-PrivateKey ::= CHOICE {
    seed [0] OCTET STRING (SIZE (64)),
    expandedKey OCTET STRING (SIZE (2400)),
    both SEQUENCE {
      seed OCTET STRING (SIZE (64)),
      expandedKey OCTET STRING (SIZE (2400))
      }
    }

  ML-KEM-1024-PrivateKey ::= CHOICE {
    seed [0] OCTET STRING (SIZE (64)),
    expandedKey OCTET STRING (SIZE (3168)),
    both SEQUENCE {
      seed OCTET STRING (SIZE (64)),
      expandedKey OCTET STRING (SIZE (3168))
      }
    }
~~~

<aside markdown="block">
  NOTE: The above syntax is from {{!RFC5912}} and is compatible with the
  2021 ASN.1 syntax {{X680}}. See {{RFC5280}} for the 1988 ASN.1 syntax.
</aside>

The `CHOICE` allows three representations of the private key:

* The `seed` format (tag [0]) contains just the 64-byte seed value
from which both the expanded private key and public key can be
derived using `ML-KEM.KeyGen_internal(d,z)` (algorithm 16) using
the first 32 octets as *d* and the remaining 32 octets as *z*.

* The `expandedKey` format contains the expanded private key
that was derived from the seed. If the seed is not exported, both the
expanded private key and public key can be derived using
`ML-KEM.KeyGen()` (algorithm 16).

* The `both` format contains both the seed and expanded private key, allowing
for interoperability; some may want to use and retain the seed and
others may only support expanded private keys.

When encoding an ML-KEM private key in a `OneAsymmetricKey` object, any
of these three formats may be used, though the `seed` format is RECOMMENDED
for storage efficiency.

The `privateKeyAlgorithm` field uses the `AlgorithmIdentifier` structure
with the appropriate OID as defined in {{oids}}. If present, the `publicKey`
field will hold the encoded public key as defined in {{pub-key}}.

NOTE: While the private key can be stored in multiple formats, the seed-only
format is RECOMMENDED as it is the most compact representation. Both the
expanded private key and the public key can be deterministically derived
from the seed using `ML-KEM.KeyGen_internal(d,z)` (algorithm 16) using the
first 32 octets as *d* and the remaining 32 octets as *z*.  Alternatively,
the public key can be extracted from the extended private key. While the `publicKey` field and
`expandedKey` format are technically redundant when using the seed-only format,
they MAY be included to enable keypair consistency checks during import operations.

 When parsing the private key, the ASN.1 tag explicitly indicates which
variant of `CHOICE` is present. Implementations should use tag `UNIVERSAL IMPLICIT [0]`
(raw value `0x80`) for `seed`, `OCTET STRING` (`0x04`) for `expandedKey`, and
`SEQUENCE` (`0x30`) for `both` to parse the private key, rather than any
other heuristic like length of the enclosing `OCTET STRING`.

{{example-private}} contains examples for ML-KEM private keys
encoded using the textual encoding defined in {{?RFC7468}}.

# Implementation Considerations

Though section 7.1 of {{FIPS203}} mentions the potential to save seed values for future expansion, Algorithm 19 does not make the seed values available to a caller for serialization.
Similarly, the algorithm that expands seed values is not listed as one of the "main algorithms" and features "internal" in the name even though it is clear that it is allowed to be exposed externally for the purposes of expanding a key from a seed.
Below are possible ways to extend the APIs defined in {{FIPS203}} to support serialization of seed values as private keys.

To support serialization of seed values as private keys, let Algorithm 19b denote the same procedure as Algorithm 19 in {{FIPS203}} except it returns (ek, dk, d, z) on line 7. Additionally, Algorithm 16 should be promoted to be a "main algorithm" for external use in expanding seed values.

Note also that unlike other private key compression methods in other algorithms, expanding a private key from a seed is a one-way function, meaning that once a full key is expanded from seed and the seed discarded, the seed cannot be re-created even if the full expanded private key is available. For this reason it is RECOMMENDED that implementations retain and export the seed, even when also exporting the expanded private key.

# Private Key Consistency Tesing

When receiving a private key that contains both the seed and the
expandedKey, the recipient SHOULD perform a seed consistency check to
ensure that the sender properly generated the private key.  Recipients
that do not perform this seed consistency check avoid keygen
and compare operations, but are unable to ensure that the `seed` and
`expandedKey` match.

If the check is done and the `seed` and the `expandedKey` are not consistent,
the recipient MUST reject the private key as malformed.

The seed consistency check consists of regenerating the expanded form from
the seed via `ML-KEM.KeyGen_internal(d,z)` (algorithm 16) using the first
32 octets as *d* and the remaining 32 octets as *z* and ensuring it is
bytewise equal to the value presented in the private key.

{{example-bad}} includes some examples of inconsistent seeds and
expanded private keys.

# Security Considerations

The Security Considerations section of {{RFC5280}} applies to this
specification as well.

Protection of the private-key information, i.e., the seed, is vital to
public-key cryptography.  Disclosure of the private-key material to another
entity can lead to masquerades.

For ML-KEM specific security considerations refer to
{{?I-D.sfluhrer-cfrg-ml-kem-security-considerations}}.

The generation of private keys relies on random numbers. The use of
inadequate pseudo-random number generators (PRNGs) to generate these
values can result in little or no security.  An attacker may find it
much easier to reproduce the PRNG environment that produced the keys,
searching the resulting small set of possibilities, rather than brute
force searching the whole key space.  The generation of quality
random numbers is difficult, and {{?RFC4086}} offers important guidance
in this area.

ML-KEM key generation as standardized in {{FIPS203}} has specific
requirements around randomness generation, described in section 3.3,
'Randomness generation'.

Many protocols only rely on the IND-CCA security of a KEM. Some
(implicitly) require further binding properties, formalized
in {{CDM23}}.
The private key format influences these binding properties.
Per {{KEMMY24}}, ML-KEM is LEAK-BIND-K-PK-secure and
LEAK-BIND-K-CT-secure when using the expanded private key format,
but not MAL-BIND-K-CT nor MAL-BIND-K-PK.
Using the 64-byte seed format provides a step up in binding security,
additionally providing MAL-BIND-K-CT security, but still not MAL-BIND-K-PK.
For more guidance, see {{?I-D.sfluhrer-cfrg-ml-kem-security-considerations}}.

# IANA Considerations

For the ASN.1 Module in {{asn1}}, IANA is requested to assign an
object identifier (OID) for the module identifier (TBD) with a
Description of "id-mod-x509-ml-kem-2025".  The OID for the module
should be allocated in the "SMI Security for PKIX Module Identifier"
registry (1.3.6.1.5.5.7.0).


--- back


# ASN.1 Module {#asn1}

This appendix includes the ASN.1 module {{X680}} for the ML-KEM.  Note that
as per {{RFC5280}}, certificates use the Distinguished Encoding Rules; see
{{X690}}. This module imports objects from {{RFC5912}} and {{!RFC9629}}.

~~~
<CODE BEGINS>
{::include X509-ML-KEM-2025.asn}
<CODE ENDS>
~~~

# Parameter Set Security and Sizes {#arnold}

Instead of defining the strength of a quantum algorithm in a traditional
manner using the imprecise notion of bits of security, NIST has
defined security levels by picking a reference scheme, which
NIST expects to offer notable levels of resistance to both quantum and
classical attack.  To wit, a KEM algorithm that achieves NIST PQC
security must require computational resources to break IND-CCA
security comparable or greater than that required for key search
on AES-128, AES-192, and AES-256 for Levels 1, 3, and 5, respectively.
Levels 2 and 4 use collision search for SHA-256 and SHA-384 as reference.

| Level | Parameter Set | Encap. Key | Decap. Key | Ciphertext | Secret |
|-      |-              |-           |-           |-           |-       |
| 1     | ML-KEM-512    | 800        | 1632       | 768        | 32     |
| 3     | ML-KEM-768    | 1184       | 2400       | 1952       | 32     |
| 5     | ML-KEM-1024   | 1568       | 3168       | 2592       | 32     |
{: #tab-strengths title="Mapping between NIST Security Level, ML-KEM parameter set, and sizes in bytes"}

# Examples {#examples}

This appendix contains examples of ML-KEM public keys, private keys,
certificates, and inconsistent seed and expanded private keys.


## Example Private Keys {#example-private}

The following examples show ML-KEM private keys in different formats,
all derived from the same seed `000102...1e1f`. For each security level,
we show the seed-only format (using a context-specific `[0]` primitive
tag with an implicit encoding of `OCTET STRING`), the expanded format,
and both formats together.

NOTE: All examples use the same seed value, showing how the same seed
produces different expanded private keys for each security level.

### ML-KEM-512 Private Key Examples

Each of the examples includes the textual encoding {{RFC7468}} followed by
the so-called "pretty print"; the private keys are the same.

#### Seed Format
~~~
{::include ./example/ML-KEM-512-seed.priv}
~~~

~~~
{::include ./example/ML-KEM-512-seed.priv.txt}
~~~

#### Expanded Format
~~~
{::include ./example/ML-KEM-512-expanded.priv}
~~~

~~~
{::include ./example/ML-KEM-512-expanded.priv.txt}
~~~

#### Both Format
~~~
{::include ./example/ML-KEM-512-both.priv}
~~~

~~~
{::include ./example/ML-KEM-512-both.priv.txt}
~~~

### ML-KEM-768 Private Key Examples

Each of the examples includes the textual encoding {{RFC7468}} followed by
the so-called "pretty print"; the private keys are the same.

#### Seed Format
~~~
{::include ./example/ML-KEM-768-seed.priv}
~~~

~~~
{::include ./example/ML-KEM-768-seed.priv.txt}
~~~

#### Expanded Format
~~~
{::include ./example/ML-KEM-768-expanded.priv}
~~~

~~~
{::include ./example/ML-KEM-768-expanded.priv.txt}
~~~

#### Both Format
~~~
{::include ./example/ML-KEM-768-both.priv}
~~~

~~~
{::include ./example/ML-KEM-768-both.priv.txt}
~~~

### ML-KEM-1024 Private Key Examples

Each of the examples includes the textual encoding {{RFC7468}} followed by
the so-called "pretty print"; the private keys are the same.

#### Seed Format
~~~
{::include ./example/ML-KEM-1024-seed.priv}
~~~

~~~
{::include ./example/ML-KEM-1024-seed.priv.txt}
~~~

#### Expanded Format
~~~
{::include ./example/ML-KEM-1024-expanded.priv}
~~~

~~~
{::include ./example/ML-KEM-1024-expanded.priv.txt}
~~~

#### Both Format
~~~
{::include ./example/ML-KEM-1024-both.priv}
~~~

~~~
{::include ./example/ML-KEM-1024-both.priv.txt}
~~~

## Example Public Keys {#example-public}

The following is the ML-KEM-512 public key corresponding to the private
key in the previous section. The textual encoding {{RFC7468}} is
followed by the so-called "pretty print"; the public keys are the same.

~~~
{::include ./example/ML-KEM-512.pub}
~~~

~~~
{::include ./example/ML-KEM-512.pub.txt}

~~~

The following is the ML-KEM-768 public key corresponding to the private
key in the previous section. The textual encoding {{RFC7468}} is
followed by the so-called "pretty print"; the public keys are the same.

~~~
{::include ./example/ML-KEM-768.pub}
~~~

~~~
{::include ./example/ML-KEM-768.pub.txt}
~~~

The following is the ML-KEM-1024 public key corresponding to the private
key in the previous section. The textual encoding {{RFC7468}} is
followed by the so-called "pretty print"; the public keys are the same.

~~~
{::include ./example/ML-KEM-1024.pub}
~~~

~~~
{::include ./example/ML-KEM-1024.pub.txt}
~~~

## Example Certificates {#example-certificate}

<aside markdown="block">
  RFC EDITOR: Please replace the following reference to
  {{?I-D.ietf-lamps-dilithium-certificates}} with a reference to the published RFC.
</aside>

The following is the ML-KEM-512 certificate that corresponding to the
public key in the previous section signed with the ML-DSA-44 private key
from {{I-D.ietf-lamps-dilithium-certificates}}. The textual encoding {{RFC7468}}
is followed by the so-called "pretty print"; the certificates are the same.

~~~
{::include ./example/ML-KEM-512.crt}
~~~

~~~
{::include ./example/ML-KEM-512.crt.txt}
~~~

<aside markdown="block">
  RFC EDITOR: Please replace the following reference to
  {{I-D.ietf-lamps-dilithium-certificates}} with a reference to the published RFC.
</aside>

The following is the ML-KEM-768 certificate that corresponding to the
public key in the previous section signed with the ML-DSA-65 private key
from {{I-D.ietf-lamps-dilithium-certificates}}. The textual encoding {{RFC7468}}
is followed by the so-called "pretty print"; the certificates are the same.

~~~
{::include ./example/ML-KEM-768.crt}
~~~

~~~
{::include ./example/ML-KEM-768.crt.txt}
~~~

<aside markdown="block">
  RFC EDITOR: Please replace the following reference to
  {{I-D.ietf-lamps-dilithium-certificates}} with a reference to the published RFC.
</aside>

The following is the ML-KEM-1024 certificate that corresponding to the
public key in the previous section signed with the ML-DSA-87 private key
from {{I-D.ietf-lamps-dilithium-certificates}}. The textual encoding {{RFC7468}}
is followed by the so-called "pretty print"; the certificates are the same.

~~~
{::include ./example/ML-KEM-1024.crt}
~~~

~~~
{::include ./example/ML-KEM-1024.crt.txt}
~~~

## Examples of Bad Private Keys {#example-bad}

<aside markdown="block">
   WARNING: These private keys are purposely bad do not use them in
   production systems.
</aside>

The following examples demonstrate inconsistent seed and
expanded private keys.

### ML-KEM Inconsistent Seed and Expanded Private Keys

Four `ML-KEM-512-PrivateKey` examples of inconsistent seed and
expanded private keys follow:

1. The first `ML-KEM-512-PrivateKey` example includes the
   `both CHOICE` , i.e., both `seed` and `expandedKey` are
   included. The `seed` and `expanded` values can be checked
   for inconsistencies.

2. The second `ML-KEM-512-PrivateKey` example includes only
   `expandedKey`. The expanded private key has a mutated `s_0`
   and a valid public key hash, but a pairwise consistency
   check would find that the public key fails to match private.

3. The third `ML-KEM-512-PrivateKey` example includes only
   `expandedKey`. The expanded private key has a mutated H(ek); both
   a public key digest check and a pairwise consistency check should fail.

4. The fourth `ML-KEM-512-PrivateKey` example includes the
   `both CHOICE` , i.e., both `seed` and `expandedKey` are
   included. There is mismatch of the seed and expanded private
   key in only the z implicit rejection secret; here the private
   and public vectors match and the pairwise consistency check passes,
   but z is different.

The following is the first example:

 ~~~
 {::include ./examples/bad-ML-KEM-512-1.priv}
 ~~~

 The following is the second example:

 ~~~
 {::include ./examples/bad-ML-KEM-512-2.priv}
 ~~~

 The following is the third example:

 ~~~
 {::include ./examples/bad-ML-KEM-512-3.priv}
 ~~~

The following is the fourth example:

 ~~~
 {::include ./examples/bad-ML-KEM-512-4.priv}
 ~~~

# Acknowledgments
{:numbered="false"}

The authors wish to thank the following people for their contributions
to this document: Deirdre Connolly, Viktor Dukhovni, Alicja Kario, Russ
Housley, Mike Ounsworth, Daniel Van Geest, Thom Wiggers, and Carl Wallace.

In addition, we would like to thank those who contributed to the private
key format discussion: Tony Arcieri, Bob Beck, Dmitry Belyavskiy, David
Benjamin, Daniel Bernstein, Uri Blumenthal, Theo Buehler, Stephen Farrell,
Jean-Pierre Fiset, Scott Fluhrer, Alex Gaynor, John Gray, Peter Gutmann,
David Hook, Tim Hudson, Paul Kehrer, John Kemp, Watson Ladd, Adam Langley,
John Mattsson, Damien Miller, Robert Relyea, Michael Richardson,
Markku-Juhani O. Saarinen, Rich Salz, Roland Shoemaker, Sophie Schmieg,
Simo Sorce, Michael St. Johns, Falko Strenzke, Filippo Valsorda, and
Wei-Jun Wang.
