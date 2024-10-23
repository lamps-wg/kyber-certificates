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

The Module-Lattice-Based Key-Encapsulation Mechanism (ML-KEM) standardized
in {{!FIPS203=DOI.10.6028/NIST.FIPS.203}} is a quantum-resistant key-encapsulation mechanism (KEM)
standardized by the US National Institute of Standards and Technology (NIST)
PQC Project {{NIST-PQC}}. Prior to standardization, the mechanism was known
as Kyber. ML-KEM and Kyber are not compatible. This document specifies the
use of ML-KEM in Public Key Infrastructure X.509 (PKIX) certificates {{!RFC5280}}
at three security levels: ML-KEM-512, ML-KEM-768, and ML-KEM-1024, using object
identifiers assigned by NIST. The private key format is also specified.

## Applicability Statement

ML-KEM certificates are used in protocols where the public key is used to
generate and encapsulate a shared secret used to derive a symmetric key used to
encrypt a payload; see {{?I-D.ietf-lamps-cms-kyber}}. To be used in
TLS, ML-KEM certificates could only be used as end-entity identity
certificates and would require significant updates to the protocol; see
{{?I-D.celi-wiggers-tls-authkem}}.

# Conventions and Definitions

{::boilerplate bcp14-tagged}

# Algorithm Identifiers

The AlgorithmIdentifier type, which is included herein for convenience,
is defined as follows:

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

The fields in AlgorithmIdentifier have the following meanings:

* algorithm identifies the cryptographic algorithm with an object
  identifier.

* parameters, which are optional, are the associated parameters for
  the algorithm identifier in the algorithm field.

The AlgorithmIdentifier for a ML-KEM public key MUST use one of the
id-alg-ml-kem object identifiers listed below, based on the security
level. The parameters field of the AlgorithmIdentifier for the ML-KEM
public key MUST be absent.

When any of the ML-KEM AlgorithmIdentifier appears in the
SubjectPublicKeyInfo field of an X.509 certificate, the key usage
certificate extension MUST only contain keyEncipherment
{{Section 4.2.1.3 of RFC5280}}.

~~~
  nistAlgorithms OBJECT IDENTIFIER ::= { joint-iso-ccitt(2)
    country(16) us(840) organization(1) gov(101) csor(3)
    nistAlgorithm(4) }

  kems OBJECT IDENTIFIER ::= { nistAlgorithms 4 }

  id-alg-ml-kem-512 OBJECT IDENTIFIER ::= { kems 1 }

  id-alg-ml-kem-768 OBJECT IDENTIFIER ::= { kems 2 }

  id-alg-ml-kem-1024 OBJECT IDENTIFIER ::= { kems 3 }

  pk-ml-kem-512 PUBLIC-KEY ::= {
    IDENTIFIER id-alg-ml-kem-512
    -- KEY no ASN.1 wrapping --
    PARAMS ARE absent
    CERT-KEY-USAGE { keyEncipherment }
    --- PRIVATE-KEY no ASN.1 wrapping --
    }

  pk-ml-kem-768 PUBLIC-KEY ::= {
    IDENTIFIER id-alg-ml-kem-768
    -- KEY no ASN.1 wrapping --
    PARAMS ARE absent
    CERT-KEY-USAGE { keyEncipherment }
    --- PRIVATE-KEY no ASN.1 wrapping --
    }

  pk-ml-kem-1024 PUBLIC-KEY ::= {
    IDENTIFIER id-alg-ml-kem-1024
    -- KEY no ASN.1 wrapping --
    PARAMS ARE absent
    CERT-KEY-USAGE { keyEncipherment }
    --- PRIVATE-KEY no ASN.1 wrapping --
    }

    ML-KEM-PublicKey ::= OCTET STRING

    ML-KEM-PrivateKey ::= OCTET STRING
~~~

No additional encoding of the ML-KEM public key value is applied in
the SubjectPublicKeyInfo field of an X.509 certificate {{RFC5280}}.
However, whenever the ML-KEM public key value appears outside of a
certificate, it MAY be encoded as an OCTET STRING.

No additional encoding of the ML-KEM private key value is applied in
the PrivateKeyInfo field of an Asymmetric Key Package {{RFC5958}}.
However, whenever the ML-KEM private key value appears outside of a
Asymmetric Key Package, it MAY be encoded as an OCTET STRING.

# Subject Public Key Fields

In the X.509 certificate, the subjectPublicKeyInfo field has the
SubjectPublicKeyInfo type, which has the following ASN.1 syntax:

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

The fields in SubjectPublicKeyInfo have the following meaning:

* algorithm is the algorithm identifier and parameters for the
  public key (see above).

* subjectPublicKey contains the byte stream of the public key.

{{example-public}} contains an example of an id-alg-ml-kem-768 public key
encoded using the textual encoding defined in {{?RFC7468}}.

# Private Key Format

An ML-KEM private key is encoded by storing its 64-octet seed in
the privateKey field as follows.

{{FIPS203}} specifies two formats for an ML-KEM private key: a 64-octet
seed and an (expanded) private key, which is referred to as the
decapsulation key. The expanded private key (and public key)
is computed from the seed using `ML-KEM.KeyGen_internal(d,z)` (algorithm 16)
using the first 32 octets as *d* and the remaining 32 octets as *z*.

"Asymmetric Key Packages" {{!RFC5958}} describes how to encode a private
key in a structure that both identifies what algorithm the private key
is for and allows for the public key and additional attributes about the
key to be included as well. For illustration, the ASN.1 structure
OneAsymmetricKey is replicated below. The algorithm-specific details of
how a private key is encoded are left for the document describing the
algorithm itself.

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

When used in a OneAsymmetricKey type, the privateKey OCTET STRING contains
the raw octet string encoding of the 64-octet seed. The publicKey field
SHOULD be omitted because the public key can be computed as noted earlier
in this section.

<aside markdown="block">
  NOTE: There exist some private key import functions that have not
  implemented the new ASN.1 structure OneAsymmetricKey that is defined in
  {{RFC5958}}. This means that they will not accept a private key
  structure that contains the public key field.  This means a balancing
  act needs to be done between being able to do a consistency check on the
  key pair and widest ability to import the key.
</aside>

{{example-private}} contains an example of an id-alg-ml-kem-768 private key
encoded using the textual encoding defined in {{?RFC7468}}.

# Security Considerations

The Security Considerations section of {{RFC5280}} applies to this specification as well.

<aside markdown="block">
  To Do: Discuss side-channels for Kyber TBD1.
</aside>

# IANA Considerations

For the ASN.1 Module in {{asn1}}, IANA is requested to assign an
object identifier (OID) for the module identifier (TBD2) with a
Description of "id-mod-x509-ml-kem-2024".  The OID for the module
should be allocated in the "SMI Security for PKIX Module Identifier"
registry (1.3.6.1.5.5.7.0).


--- back


# ASN.1 Module {#asn1}

This appendix includes the ASN.1 module {{X680}} for the ML-KEM.  Note that
as per {{RFC5280}}, certificates use the Distinguished Encoding Rules; see
{{X690}}. This module imports objects from {{RFC5912}} and {{!RFC9629}}.

~~~
<CODE BEGINS>
{::include X509-ML-KEM-2024.asn}
<CODE ENDS>
~~~

# Examples {#examples}

This appendix contains examples of ML-KEM public keys, private keys and certificates.


## Example Private Key {#example-private}

The following is an example of a ML-KEM-512 private key with hex seed `0001…3f`:

~~~
{::include ./example/ML-KEM-512.priv}
~~~

The following is an example of a ML-KEM-768 private key from the same seed.

~~~
{::include ./example/ML-KEM-768.priv}
~~~

The following is an example of a ML-KEM-1024 private key from the same seed.

~~~
{::include ./example/ML-KEM-1024.priv}
~~~

## Example Public Key {#example-public}

The following is the ML-KEM-512 public key corresponding to the private
key in the previous section.

~~~
{::include ./example/ML-KEM-512.pub}
~~~

The following is the ML-KEM-768 public key corresponding to the private
key in the previous section.

~~~
{::include ./example/ML-KEM-768.pub}
~~~

The following is the ML-KEM-1024 public key corresponding to the private
key in the previous section.

~~~
{::include ./example/ML-KEM-1024.pub}
~~~

The following example, in addition to encoding the ML-KEM-768 private key,
has an attribute included as well as the public key:

~~~
  -----BEGIN PRIVATE KEY-----
  TODO insert example private key with attribute
  -----END PRIVATE KEY-------
~~~

## Example Certificate {#example-certificate}

~~~
  TODO insert ASN.1 Pretty Print
~~~

~~~
  -----BEGIN CERTIFICATE-----
  TODO Certificate
  -----END CERTIFICATE-------
~~~

# Acknowledgments
{:numbered="false"}

TODO acknowledge.
