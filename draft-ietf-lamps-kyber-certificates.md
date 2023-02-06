---
title: Internet X.509 Public Key Infrastructure - Algorithm Identifiers for Kyber
abbrev: PQC Kyber in Certificates
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
    title: "Information technology - Abstract Syntax Notation One (ASN.1): Specification of basic notation"
    date: Feburary 2021
    author:
      org: ITU-T
    seriesinfo:
        ISO/IEC: 8824-1:2021
  X690:
    target: https://www.itu.int/rec/T-REC-X.690
    title: "Information technology - Abstract Syntax Notation One (ASN.1): ASN.1 encoding rules: Specification of Basic Encoding Rules (BER), Canonical Encoding Rules (CER) and Distinguished Encoding Rules (DER)"
    date: Feburary 2021
    author:
      org: ITU-T
    seriesinfo:
        ISO/IEC: 8825-1:2021

informative:
  PQCProj:
    target: https://csrc.nist.gov/projects/post-quantum-cryptography
    title: Post-Quantum Cryptography Project
    author:
      - org: National Institute of Standards and Technology
    date: 2016-12-20

--- abstract

Kyber is a key-encapsulation mechanism (KEM). This document specifies
algorithm identifiers and ASN.1 encoding format for Kyber in public
key certificates. The encoding for public and private keys are also
provided.

\[EDNOTE:
This document is not expected to be finalized before the NIST PQC
Project has standardized PQ algorithms. This specification will use
object identifiers for the new algorithms that are assigned by NIST,
and will use placeholders until these are released.]

--- middle

# Introduction

Kyber is a key-encapsulation mechanism (KEM) standardized by the US NIST
PQC Project {{PQCProj}}. This document specifies the use of the Kyber
algorithm at three security levels: Kyber512, Kyber768, and Kyber1024,
in X.509 public key certificates; see {{!RFC5280}}. Public and private
key encodings are also specified.

## ASN.1 and Kyber Identifiers

An ASN.1 module {{X680}} is included for reference purposes. Note that
as per {{RFC5280}}, certificates use the Distinguished Encoding Rules;
see {{X690}}. Also note that NIST defined the object identifiers for
the Kyber algorithms in an ASN.1 modulle; see (TODO insert reference).

## Applicability Statement

Kyber certificates are used in protocols where the public key is used to
generate and encapsulate a shared secret used to derive a symmetric key used to
encrypt a payload; see {{?I-D.ietf-lamps-kyber}}. To be used in
TLS, Kyber certificates could only be used as end-entity identity
certificates and would require significant updates to the protocol; see
{{?I-D.celi-wiggers-tls-authkem}}.

# Conventions and Definitions

{::boilerplate bcp14-tagged}


# Algorithm Identifiers {#Alg-IDs}

Certificates conforming to {{RFC5280}} can convey a public key for any
public key algorithm. The certificate indicates the algorithm through
an algorithm identifier. An algorithm identifier consists of an object
identifier and optional parameters.

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
2021 ASN.1 syntax {{X680}}.
</aside>

The fields in AlgorithmIdentifier have the following meanings:

* algorithm identifies the cryptographic algorithm with an object
  identifier.

* parameters, which are optional, are the associated parameters for
  the algorithm identifier in the algorithm field.

{{Kyber-TBD1}} includes object identifiers for Kyber-512, Kyber-768, and
Kyber-1024. For all of these OIDs, the parameters MUST be absent.

<aside markdown="block">
NOTE: It is possible to find systems that require the parameters to be
present. This can be due to either a defect in the original 1997
syntax or a programming error where developers never got input where
this was not true. The optimal solution is to fix these systems;
where this is not possible, the problem needs to be restricted to
that subsystem and not propagated to the Internet.
</aside>


# Candidate TBD1 {#Kyber-TBD1}

TODO insert object-identifiers for Kyber-768, Kyber-1024 etc.


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
2021 ASN.1 syntax {{X680}}.
</aside>

The fields in SubjectPublicKeyInfo have the following meaning:

* algorithm is the algorithm identifier and parameters for the
  public key (see above).

* subjectPublicKey contains the byte stream of the public key.  The
  algorithms defined in this document always encode the public key
  as TODO pick format e.g., exact multiple of 8 bits?.

The following is an example of a Kyber-512 public key encoded using the
textual encoding defined in {{?RFC7468}}:

~~~
  -----BEGIN PUBLIC KEY-----
  TODO insert example public key
  -----END PUBLIC KEY-------
~~~

# Key Usage Bits

The intended application for the key is indicated in the keyUsage
certificate extension; see {{Section 4.2.1.3 of RFC5280}}.

If the keyUsage extension is present in a certificate that indicates
Kyber TBD1 in SubjectPublicKeyInfo, then the following
MUST be present:

~~~
  keyEncipherment;
~~~

# Private Key Format

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
      privateKeyAlgorithm      PrivateKeyAlgorithmIdentifier,
      privateKey               PrivateKey,
      attributes           [0] IMPLICIT Attributes OPTIONAL,
      ...,
      [[2: publicKey       [1] IMPLICIT PublicKey OPTIONAL ]],
      ...
  }

  PrivateKey ::= OCTET STRING

  PublicKey ::= BIT STRING
~~~

<aside markdown="block">
NOTE: The above syntax is from {{RFC5958}} and matches the version used
therein, i.e., the 2002 ASN.1 syntax. The syntax used therein is
compatible with the 2015 ASN.1 syntax.
</aside>

For the keys defined in this document, the private key is always an
opaque byte sequence. The ASN.1 type PqckemPrivateKey is defined in
this document to hold the byte sequence. Thus, when encoding a
OneAsymmetricKey object, the private key is wrapped in a
PqckemPrivateKey object and wrapped by the OCTET STRING of the
"privateKey" field.

~~~
  PqckemPrivateKey ::= OCTET STRING
~~~

The following is an example of a Kyber TBD private key encoded using the
textual encoding defined in {{RFC7468}}.

~~~
  -----BEGIN PRIVATE KEY-----
  TODO iser example private key
  -----END PRIVATE KEY-------
~~~

The following example, in addition to encoding the Kyber TBD private key,
has an attribute included as well as the public key. As with the
prior example, the textual encoding defined in {{RFC7468}} is used.

~~~
  -----BEGIN PRIVATE KEY-----
  TODO insert example private key with attribute
  -----END PRIVATE KEY-------
~~~

<aside markdown="block">
NOTE: There exist some private key import functions that have not
implemented the new ASN.1 structure OneAsymmetricKey that is defined in
{{RFC5958}}. This means that they will not accept a private key
structure that contains the public key field.  This means a balancing
act needs to be done between being able to do a consistency check on the
key pair and widest ability to import the key.
</aside>

# ASN.1 Module

TODO ASN.1 Module


# Security Considerations

The Security Considerations section of {{RFC5280}} applies to this specification as well.

\[EDNOTE: Discuss side-channels for Kyber TBD1.\]


# IANA Considerations

This document will have some IANA actions.


--- back

# Acknowledgments
{:numbered="false"}

TODO acknowledge.
