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

{{example-public}} contains examples for ML-KEM public keys
encoded using the textual encoding defined in {{?RFC7468}}.

# Private Key Format

In short, an ML-KEM private key is encoded by storing its 64-octet seed in
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
OneAsymmetricKey is replicated below.

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

{{example-private}} contains examples for ML-KEM private keys
encoded using the textual encoding defined in {{?RFC7468}}.

# Security Considerations

The Security Considerations section of {{RFC5280}} applies to this
specification as well.

<aside markdown="block">
  To Do: Discuss side-channels for Kyber TBD1.
</aside>

# IANA Considerations

For the ASN.1 Module in {{asn1}}, IANA is requested to assign an
object identifier (OID) for the module identifier (TBD) with a
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

This appendix contains examples of ML-KEM public keys, private keys and
certificates.


## Example Private Key {#example-private}

The following is an example of a ML-KEM-512 private key with hex seed `0001…3f`:

~~~
{::include ./example/ML-KEM-512.priv}
~~~

~~~
0  82: SEQUENCE
2   2:  INTEGER 0
5  11:  SEQUENCE {
7   9:   OBJECT IDENTIFIER '2.16.840.1.101.3.4.4.1'
     :   }
18 64:  OCTET STRING
     :    00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f
     :    10 11 12 13 14 15 16 17 18 19 1a 1b 1c 1d 1e 1f
     :    20 21 22 23 24 25 26 27 28 29 2a 2b 2c 2d 2e 2f
     :    30 31 32 33 34 35 36 37 38 39 3a 3b 3c 3d 3e 3f
     :  }
~~~

The following is an example of a ML-KEM-768 private key from the same seed.

~~~
{::include ./example/ML-KEM-768.priv}
~~~

~~~
0  82: SEQUENCE
2   2:  INTEGER 0
5  11:  SEQUENCE {
7   9:   OBJECT IDENTIFIER '2.16.840.1.101.3.4.4.2'
     :   }
18 64:  OCTET STRING
     :    00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f
     :    10 11 12 13 14 15 16 17 18 19 1a 1b 1c 1d 1e 1f
     :    20 21 22 23 24 25 26 27 28 29 2a 2b 2c 2d 2e 2f
     :    30 31 32 33 34 35 36 37 38 39 3a 3b 3c 3d 3e 3f
     :  }
~~~

The following is an example of a ML-KEM-1024 private key from the same seed.

~~~
{::include ./example/ML-KEM-1024.priv}
~~~

~~~
0  82: SEQUENCE
2   2:  INTEGER 0
5  11:  SEQUENCE {
7   9:   OBJECT IDENTIFIER '2.16.840.1.101.3.4.4.2'
     :   }
18 64:  OCTET STRING
     :    00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f
     :    10 11 12 13 14 15 16 17 18 19 1a 1b 1c 1d 1e 1f
     :    20 21 22 23 24 25 26 27 28 29 2a 2b 2c 2d 2e 2f
     :    30 31 32 33 34 35 36 37 38 39 3a 3b 3c 3d 3e 3f
     :  }
~~~

<aside markdown="block">
  NOTE: The private key is the seed and all three examples keys
  use the same seed; therefore, the private above are the same
  except for the OID used to represent the ML-KEM algorithm's
  security strength.
</aside>

## Example Public Key {#example-public}

The following is the ML-KEM-512 public key corresponding to the private
key in the previous section.

~~~
{::include ./example/ML-KEM-512.pub}
~~~

~~~
0  818: SEQUENCE {
4   11:   SEQUENCE {
6    9:     OBJECT IDENTIFIER '2.16.840.1.101.3.4.4.1'
      :     }
17 801:   BIT STRING
      :     00 39 95 81 5e 59 7d 10 43 55 cf 29 aa 53 33 c9
      :     32 51 86 9d 5b cd be 48 71 24 f6 02 b8 b6 a6 6c
      :     7f 0a c0 76 b0 c6 2e fa 32 81 53 e7 ca 57 01 69
      :     9f 13 05 f1 e6 bc 6f 90 b0 e4 9b 69 35 12 b6 ce
      :     99 2a 8b 80 16 dd fc 1a 66 2c 7e 3f 96 19 cb d8
      :     69 dd 77 1a f3 08 96 cc d5 91 8a c6 cb 77 46 6c
      :     5e 77 99 96 d6 7f f9 aa bc 97 50 3f 2c 7b 7e 2d
      :     00 0d 86 45 0f b1 80 7c a4 ca bd a4 65 82 5a 31
      :     c7 89 a1 b7 a4 91 ab 38 72 76 5d 32 0d 0b 71 92
      :     0f a2 13 c9 40 93 41 6b 83 b8 12 4e 69 f6 5e 62
      :     cb 50 00 dc c3 7a a9 a0 ff f7 39 70 c4 77 2f 35
      :     7d 24 18 9c a6 f5 30 55 68 c0 e2 37 6a 37 62 a6
      :     8c 60 5e 56 3c 5d 20 95 72 e0 fc 75 32 ca 29 47
      :     29 53 55 67 b5 fc 41 3c 5e 87 92 d2 46 45 36 cc
      :     80 8f 98 ad d7 46 64 f1 41 56 6f 90 16 a9 0a 54
      :     18 29 a9 8a 04 64 ce 41 a8 bb 44 c2 d4 fa 3c 2c
      :     20 94 60 72 8e f1 4a 1a 7c 4c 9b 98 d1 22 03 b4
      :     cc 35 29 16 0a 9a b2 d7 83 8f 7f f6 b5 3a e0 5a
      :     a3 1a 7d 64 6b 7a fa 6c 45 93 25 26 a3 c3 75 56
      :     19 be 99 4c 21 1c 2a 31 c0 5b 34 47 83 6c b2 15
      :     0b e1 82 9d ae 6b 04 c5 53 5c ff 54 6e 39 2b a7
      :     97 41 17 20 f9 24 f4 90 a5 ac 54 95 f2 13 56 d5
      :     50 b7 82 a6 4c 16 88 b6 b6 55 bc c7 84 21 97 a4
      :     34 c2 f6 56 3b 5b 7f 09 a7 8b cc 48 82 32 78 35
      :     61 d1 6f 4c ba b6 75 54 00 05 07 81 57 0c 66 60
      :     4b 81 7a d1 25 22 94 73 6e 8b 01 86 1a 4b 5a 74
      :     51 9b 8b 6f e5 14 89 a5 07 23 92 e5 87 62 6c 71
      :     37 76 57 5d 33 80 6a 1c 8e 27 32 af 97 c2 68 0f
      :     51 66 63 31 c4 eb 8b bc 04 31 c4 f9 68 32 da f1
      :     b3 c4 55 28 fb a1 53 f6 c7 8b 1c 19 87 02 94 7c
      :     cd 33 77 27 a4 6f b5 3b a1 1d e5 cb 41 91 34 68
      :     59 51 6c b6 ad 72 40 0f 3c f2 09 b2 36 ae f3 5a
      :     58 0a c8 7e b3 e3 0f af d6 69 73 ca 8a 7d d2 67
      :     5a f4 1f 7a 17 b6 14 33 cd 1a f8 0f 77 08 86 9f
      :     66 54 88 49 79 80 b1 ac 10 a0 cd cb 63 6a 00 ed
      :     86 81 b3 5e 42 91 24 ca 80 35 07 25 b8 5f 83 a5
      :     ea c3 a4 a3 cc 16 00 90 3e 65 29 35 60 b9 b3 36
      :     e5 af 0d 52 9d ac 1a 04 81 19 30 2c b7 a9 bc c1
      :     10 b9 48 51 bf 02 11 7f 19 9d c4 85 a8 52 b7 47
      :     3f 09 b8 31 a6 83 1d 5b 54 c0 b7 90 d2 25 cf 6b
      :     b9 2d 94 62 a2 6c db 33 dd a5 12 3c 7a af 0e 26
      :     a0 b8 36 55 ee a2 8b f3 a8 07 47 25 01 8f d6 ba
      :     e4 b6 01 cf 61 ba ab 71 a7 a3 d3 51 97 a3 43 e7
      :     4b 4a 27 2c 12 5d 54 08 96 42 6d 85 b7 95 8d 3b
      :     38 a6 ba 98 7e c3 72 25 c7 b4 4c db 12 dd e4 53
      :     9b 4a b0 82 36 36 83 f0 4b f7 a0 9c c5 c4 1d fe
      :     83 0a 1b 16 2e 0b 32 43 34 36 2f 08 4a 14 46 77
      :     23 34 4b ad d0 00 f8 d8 c5 37 c4 8f 99 8f 05 30
      :     7c eb d1 ed e0 b8 1c 3b c5 9a 06 5a 1b 6d 63 b2
      :     6c
      :   }
~~~

The following is the ML-KEM-768 public key corresponding to the private
key in the previous section.

~~~
{::include ./example/ML-KEM-768.pub}
~~~

~~~
0 1202: SEQUENCE {
4   11:   SEQUENCE {
6    9:     OBJECT IDENTIFIER '2.16.840.1.101.3.4.4.2'
      :     }
17 801:   BIT STRING
      :     00 29 8a a1 0d 42 3c 8d da 06 9d 02 bc 59 e6 cd
      :     f0 3a 09 6b 8b 3d a4 ca b9 b8 0c a4 a1 49 07 67
      :     2c ce f1 ec 4f af 23 4a 0b c5 b7 e9 d4 73 f2 b3
      :     13 3b 3b 26 a1 d1 75 cb 67 a7 80 59 19 69 9c 02
      :     f7 65 31 b9 9c 5f 89 18 07 04 bb 4c a4 53 5c 5b
      :     89 72 67 9c 66 0a 07 c5 e5 14 b8 70 09 c8 62 eb
      :     8f 51 57 69 5e fb 3f c4 0a 9d ef 6b 81 c1 cc 02
      :     a2 49 ae 4f 09 4a d0 d9 bd 34 85 c1 c1 c6 80 80
      :     52 0a 7c 8c 63 20 32 ce e7 38 15 4e 5c 51 76 c0
      :     7d a5 60 24 77 6a 43 0f e7 6e ac f6 65 a3 f7 b8
      :     32 10 22 15 bc 82 f1 09 39 c8 35 57 04 33 6a 8f
      :     ac 1d 81 e4 bb 04 85 aa 5d 7c 74 d6 b5 9b be 5c
      :     5e 97 2a 0d 8b ac 41 1b 55 b5 d5 55 7c d6 80 a1
      :     a8 f7 1b 4e b8 6b c4 8c 9a 05 09 73 1a 54 bd 9d
      :     72 90 b2 79 63 e4 37 2d c9 b1 99 cf dc ac 0b 01
      :     ac d2 8a 62 39 51 12 e4 c4 36 48 d6 22 c4 8c 82
      :     34 d0 14 40 e8 cc 37 6c 92 7f 23 a5 af c9 ac 04
      :     74 c6 62 27 4e 42 45 25 c8 55 2e ce 3b 3f e2 65
      :     16 de 90 1b c7 d5 15 bd e8 95 58 e6 26 c9 5c 80
      :     b9 33 42 f8 01 00 04 f3 9e 6c 6c 94 87 1c 5e 34
      :     4c ab 39 66 c8 35 f9 a9 6a 59 af d3 1c 40 28 6b
      :     38 b1 c1 a7 84 70 ba b9 47 51 89 34 45 3c e8 67
      :     36 a9 19 f1 f5 a6 d5 10 a8 6f 54 54 fc 39 80 cb
      :     5c 76 5b d2 bd 5f 7b 36 b1 41 0d 66 35 c8 ce b4
      :     7c 4d da 0d 76 a2 8e ac 93 9c 71 c3 02 48 04 86
      :     6c 71 62 66 58 44 21 63 c2 c2 21 17 e5 0a ce fc
      :     e6 37 8a 98 56 52 30 2a 4e f0 c2 ce 0c c7 16 b7
      :     79 6e 2b 6b 2e 37 77 df a1 ac 3d a2 59 a3 1b 5a
      :     9b 53 0f 8c b6 38 a8 1a 62 ac 30 18 49 ab af 95
      :     a7 30 1b da 30 06 89 09 bf db 7e 67 db cc bb 38
      :     a5 55 1a 25 b1 a3 a0 f6 85 74 8a d5 75 3d 88 80
      :     f0 01 6c 62 74 86 16 63 84 c5 57 1f e2 36 59 00
      :     36 4d 03 83 11 e2 d8 75 db 36 66 86 93 2b 5e c6
      :     02 43 0a 36 9e 87 a6 ef 5c 33 87 86 65 78 25 bd
      :     4c 05 7a ce b9 23 eb 09 35 e6 90 5e 63 b4 ce d7
      :     f8 08 57 a7 73 dd 64 b1 50 d2 66 12 ea 9a c1 20
      :     52 db 20 17 bf 18 43 cc b4 b3 28 1b 69 0d c7 28
      :     ad fa 85 c0 02 81 b8 e3 c0 92 87 33 5f 85 6b 4f
      :     c2 89 2f 69 a2 f5 79 21 ad a0 19 14 c4 09 88 66
      :     2d 57 76 96 62 a7 86 35 1b 9b 66 49 3d ab 79 59
      :     4d 98 6d e2 10 0d 65 ba 0f f4 ea 58 b8 15 38 d2
      :     4a 44 35 a2 58 fa c2 54 04 aa 7f 41 f6 58 b1 38
      :     50 65 e1 58 dc b6 01 15 73 27 20 f4 04 59 aa ac
      :     15 e4 06 95 3a 90 ac 52 99 7d 1c cd 07 00 60 ef
      :     c6 5d b9 e6 53 35 44 67 fa d5 6e c7 13 c8 6e 75
      :     40 c4 23 ac f2 66 9f 52 fa 6f 4a c6 88 8d 87 1e
      :     f3 e8 47 c0 29 a8 aa fb b9 2e 17 b2 4a a0 79 b1
      :     f4 19 ba 61 75 b4 42 af b1 19 09 d4 a5 6b 70 a0
      :     33 5b 28 73 92 18 aa 7c 93 48 e2 c3 c2 f3 eb 3d
      :     15 a4 1e 64 17 c0 dd 94 bf eb 21 41 9b 31 1a 7b
      :     b1 3a 18 0b be 83 32 18 a9 a6 b1 74 47 cc 85 f2
      :     25 85 95 87 a7 30 77 04 9a cb cf d4 4d 0f 02 54
      :     38 e1 5d 15 38 27 0d 58 6e 1b f8 31 92 a9 45 9c
      :     f6 3c 0e 97 2f 85 29 76 79 83 1e cf 12 15 09 85
      :     1c b8 34 0f 6f 10 7b 0f a1 a0 ef d1 b3 6a 81 89
      :     bc 08 5c 4f 5c b7 84 e5 53 f4 1b 91 8f 80 39 7c
      :     e1 95 6f 78 5b ee 37 7c a9 aa 8b e6 99 8a da 30
      :     c2 6b 7c 3d 8c 6b 55 25 4c c9 62 03 b2 0c 42 ae
      :     e0 ac 4e 1e bb 40 8e 49 a9 e3 f8 79 d0 ab 07 85
      :     eb 70 25 42 5d 13 05 a2 29 9c 01 5e 12 0d 16 3b
      :     0e 19 49 4c e5 72 53 d0 24 6d 18 27 45 cb 81 97
      :     ab 74 38 b3 c1 bb 79 72 be c5 a3 06 eb a3 56 78
      :     55 c0 14 69 9f ef 65 ae 54 c7 70 a0 d8 5c 18 40
      :     0c f6 42 ae dc 66 07 77 ba 4b 13 85 02 bd 5a 78
      :     12 f6 21 f8 4a 48 29 6b 98 dd 43 22 b6 f1 58 28
      :     b8 a8 f0 e0 0a 8b a4 4a 53 c3 a8 b1 43 57 1b 07
      :     40 ab d5 67 da f1 cd e9 c7 9c 20 4b 6d 5e 25 9d
      :     17 66 a3 1b bb cb 4e 6a 05 cf 45 02 17 6b 30 1c
      :     1c 2f 41 24 77 50 15 7b ce c8 5e 80 9b 30 a4 d6
      :     0d 77 47 cd d0 f5 b9 9a a8 c8 26 98 75 17 79 3a
      :     aa 80 80 a0 b1 24 a8 55 8d f7 2b be 37 b7 5f 4e
      :     db b6 be 82 16 d6 c6 33 fb 2b 22 80 e2 51 13 d8
      :     69 5e 43 48 1c 3e eb 39 7e b1 92 50 52 29 b6 7a
      :     20 1e a8 93 c3 e2 cb 32 da 8b c3 42 fa 4d ea 05
      :     78
      :   }
~~~

The following is the ML-KEM-1024 public key corresponding to the private
key in the previous section.

~~~
{::include ./example/ML-KEM-1024.pub}
~~~

~~~
0  1586: SEQUENCE {
4    11:   SEQUENCE {
6     9:     OBJECT IDENTIFIER '2.16.840.1.101.3.4.4.3'
       :     }
17 1569:   BIT STRING
       :     00 4b 94 c2 94 50 11 11 91 82 3b 35 14 c9 ac 1e
       :     a3 d9 82 5c cb 86 39 3a 2d fb 04 65 4f a2 19 2d
       :     37 bf ad 1c 49 7c 65 02 ee e5 ca 80 a7 3b fc e0
       :     ba f5 a5 4a 88 58 5a 40 13 97 a3 d2 32 f4 26 a7
       :     af b0 82 bc 21 a4 43 17 09 0e aa c7 59 2c 2e a8
       :     8a 65 3c 44 91 ea 19 39 31 33 5f 52 e9 89 a3 c4
       :     cc 56 d9 c5 53 73 2d 57 c4 70 fb 41 ab 75 9b 65
       :     d2 d0 44 45 38 2f cd 9c 4e 34 4a 11 28 fa 9e 11
       :     e0 43 58 e1 92 ed 01 4b 23 23 2a 7e e2 b2 2e 23
       :     71 7f 44 11 1e e3 35 75 39 9c 37 64 6d a9 81 3e
       :     c9 b2 12 af e9 4e 5d c5 c2 33 0a 72 94 cc 1f 42
       :     34 a6 d3 fb b4 f1 68 5a b8 89 2c 04 ac b1 7c d1
       :     c1 70 d7 b0 61 1b 6a 71 76 c7 94 cc 8c 67 f5 5f
       :     c9 23 c2 ad 20 31 00 f3 65 99 18 82 c3 02 43 d7
       :     78 13 84 3b 5e c7 c9 64 03 22 63 70 60 92 ec f0
       :     0c 75 16 be 64 e4 59 8c a4 22 6c 06 9b b5 e6 7e
       :     41 75 cf 22 86 c8 dd 5c 48 8a 6c 58 61 f3 1b aa
       :     0b d0 26 94 70 e8 b5 51 dd 3b cd 38 c8 6c 12 f9
       :     cd b1 76 c7 7d c8 b6 c0 2a 70 1f 47 89 02 c8 55
       :     3f 69 4c 0d 82 72 7b 4c 4a 5c 2c 10 41 21 2a a1
       :     27 48 08 b8 21 11 b3 77 ec 75 21 4e 9b 19 78 f7
       :     60 04 d4 13 9d 98 61 3f 4b 8e 98 d2 0a f7 b5 34
       :     07 3a 50 9a 95 9b 7a 75 64 f9 b4 0c a2 18 bf 61
       :     82 93 20 a8 50 20 17 95 4d 32 8d 7a c6 c7 69 ec
       :     29 70 07 56 e7 b0 68 5b 34 0d 5e 11 80 59 50 4a
       :     49 a9 a5 0a 10 19 8e b1 0a 57 84 67 8e b4 27 d7
       :     b4 ba bb 95 52 93 3b 06 28 97 97 3e 13 18 ea f0
       :     a0 ea c3 75 84 a6 54 01 b1 70 3e 04 2a cc d8 37
       :     53 14 83 f2 41 ca dc d1 c1 d3 78 11 9e 69 44 29
       :     db 19 9a c8 91 e4 c5 34 37 57 08 5b b3 ae 78 36
       :     67 35 0c 44 58 d9 76 72 e8 61 e8 0b 1d 26 79 51
       :     0e a3 a6 f2 36 0c 77 a4 69 42 c7 a0 6a 55 4d 22
       :     80 80 c8 4b 47 ae f1 4d b1 76 20 cb 16 c0 6a b3
       :     0a 1b e4 cd a7 08 2b e9 f8 7e 9c 21 1c 46 91 63
       :     49 a5 ba 8e aa 52 01 c7 29 4a 3c 08 85 b5 3b 65
       :     74 52 10 88 25 ec 64 6c 90 a0 46 12 32 4e e7 d0
       :     31 af e5 34 31 32 cb ef 67 b6 ef b1 a5 ec 28 09
       :     b7 73 53 8c e7 7b 3d 8b 04 eb 0b 3c 22 56 01 1e
       :     4c 71 6c 19 a8 ba 07 52 bf 71 49 21 17 64 9f 06
       :     15 c3 29 0f c2 9a 46 fd e4 bd 52 db 92 86 d6 03
       :     38 82 44 25 9c 15 a7 ac 2b 64 0a 60 cc 03 37 6a
       :     58 41 a3 fb 8a 47 35 68 fa 9b 1a 26 72 15 f3 4c
       :     01 69 7b 0f 0e 62 71 75 d7 21 05 b7 70 7c 29 b9
       :     e6 14 bd c3 3a 6f 6c 81 8a 95 37 0b 42 78 82 d7
       :     b4 76 79 6a 9e c6 eb 99 32 74 cd 9b 23 91 a8 2b
       :     a4 5e 33 93 d2 e9 ae 97 21 ca 9d 6c 1b 98 8b 58
       :     27 71 3f 90 a6 58 5d e9 43 35 28 c0 2b 03 ce 10
       :     bb 5f 72 01 38 d0 fb b4 c3 0c 12 66 b9 18 e5 29
       :     25 df e1 7b 37 f9 5d 22 bc a5 4f 47 59 19 ac 85
       :     90 98 c0 f0 d0 8a c5 87 5e f2 9b 56 fd 14 1e 6e
       :     f1 5f 70 0a 0b 66 f3 95 95 c5 88 17 73 73 c4 66
       :     9b 21 bc 07 1e 4c 3a a5 f0 b4 a3 1b 62 58 f3 5d
       :     a2 4a c3 cd 29 c7 f2 09 24 10 c5 07 83 55 b1 38
       :     fb 53 a6 b9 ae 6e 0b 9c 08 24 3e 7b aa 45 c4 73
       :     76 eb 8c 7f 13 d4 cf 51 aa 73 6f a3 15 40 c9 24
       :     1f 37 0d a5 44 bf 9f 9c 28 d9 a5 7e 2f 2a 7c a9
       :     5a 4e 4b 46 6e 64 1a b3 bc c7 6a df 11 39 d5 67
       :     a6 f1 2b 52 f3 a6 5e 7e c0 aa e2 6b ca a8 c5 58
       :     33 b0 4e 59 99 8e bc 9a 19 30 fb b6 d2 23 3c 53
       :     d2 c1 f8 b9 51 8e 3c 2d e7 3a 19 de e6 b3 80 a5
       :     b3 29 71 cf 64 e1 29 fd 6c 1f a6 e7 5d 4a 23 45
       :     01 e9 66 dd 3a 54 0a f5 c8 f4 f3 4a 6b 4a 25 3e
       :     e2 84 92 56 6d 5e 67 c6 f5 58 55 fc b0 50 6f b0
       :     6c 15 67 44 d9 a0 3a 31 a2 6f a9 4c ad 14 f1 57
       :     b7 f3 03 d0 7a 69 c7 73 76 8f cb 4d 07 9c 09 05
       :     97 03 a0 c3 a9 4d e4 b9 9e a3 a2 f1 65 83 d0 f9
       :     17 0a 39 50 db 07 b4 f0 bc 30 80 29 27 f9 f7 96
       :     1b 62 59 89 26 36 a9 50 2a 27 05 30 36 37 79 9d
       :     d3 44 da 45 1c 1c f7 bf 67 84 0c eb 30 79 ab 8c
       :     6b 8c 19 27 f6 40 53 c6 12 45 0c 45 c9 e6 03 bc
       :     16 66 6e 59 6b 34 71 e1 03 b6 f1 54 47 42 4d 17
       :     02 20 48 11 1f fb d3 7e 1c 67 0f 64 f1 4b 8a 7b
       :     32 b9 4c 1a 49 b4 5d d2 fc 38 cd 52 89 d9 10 ad
       :     63 60 2c f5 e1 30 42 c6 4a c6 79 7b 89 fb 55 1a
       :     d0 8e 05 a9 2d 20 0c cc b7 e7 12 ef 23 c9 31 2c
       :     b3 50 f0 29 ab 53 7e 28 73 47 fd 30 75 ac 10 90
       :     6a 78 3f 1c 6c 07 cc b8 8f 41 22 8c 4b e1 c6 40
       :     f7 90 b5 c3 a5 d5 d3 ca 79 24 95 d7 4b c4 61 56
       :     26 58 c0 7a c6 00 27 6b 92 4a b5 bc 9b e1 f0 49
       :     4c b7 6f 82 f4 60 a7 48 09 72 66 33 81 e1 69 99
       :     60 61 d7 99 85 9e c5 4d 4f 5c a5 c4 11 c0 1d b1
       :     59 7b 16 59 77 66 9d e1 3a 92 8a 34 af ba c2 58
       :     fe a8 c4 76 42 39 c9 42 1d c3 11 9b f5 b4 76 99
       :     20 69 78 32 7b 1c 53 45 ef 74 6a 79 83 84 1f 05
       :     6e 25 34 10 0a b2 4d 4e 9a bb d0 b1 7c 6a 95 bd
       :     4c 3c 0e 40 f6 9e 16 12 ac ee b2 8b 99 08 6c 95
       :     11 6e 72 04 27 38 93 39 0b f4 6b 89 9b 36 28 6b
       :     0e bf 19 47 bb 98 84 f7 32 ca 27 da 82 b1 9b 5d
       :     c0 cc 7f 88 85 71 49 10 88 8b 23 10 c4 f9 31 9d
       :     41 0b 34 e6 43 3b 90 03 e2 17 6b b9 95 25 74 56
       :     10 6e 89 52 16 3b 8b a5 92 53 0c c5 aa 0a eb 43
       :     ad 39 8f e9 e9 7b aa 52 3d 7a 44 31 67 7c 3d 3a
       :     f0 71 9e 47 5d b8 5c a9 5a f5 08 9b ea be b0 5b
       :     2f aa b4 89 6b a6 0f 81 c8 84 72 a5 7b 46 a8 28
       :     82 6a 0c df b4 46 f8 18 91 82 d2 bf 5e ac 4e c1
       :     cc 5d ea f5 99 c8 a1 3e 48 23 54 06 d1 7f fd dc
       :     83 44 b6 c6 69 84 a8 68 aa 92 fa 02 22 7a 08 69
       :     50 eb 0c 87 01 ed 58 dc 62 87 76 b9 83 88 2e 11
       :     75
       :   }
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
