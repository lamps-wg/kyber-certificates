---
title: Algorithm Identifiers for NIST's PQC Algorithms for Use in the Internet X.509 Public Key Infrastructure
abbrev: PQC KEM for Certificates
docname: draft-turner-lamps-nist-pqc-kem-certificates
category: std

ipr: trust200902
area: SEC
workgroup: None
keyword:
  group: "Limited Additional Mechanisms for PKIX and SMIME (lamps)"
  type: "Working Group"
  mail: "spasm@ietf.org"
  arch: "https://mailarchive.ietf.org/arch/browse/spasm/"
  github: " seanturner/draft-turner-lamps-nist-pqc-kem-certificates"

stand_alone: yes
smart_quotes: no
pi: [toc, sortrefs, symrefs]

author:
 -
    name: Sean Turner
    organization: sn3rd
    email: sean@sn3rd.com

normative:

informative:
  PQCComp:
    target: https://csrc.nist.gov/projects/post-quantum-cryptography
    title: Post-Quantum Cryptography Project
    author:
      - org: National Insititue of Standards and Technology
    date: 2016-12-20


--- abstract

This document specifies algorithm identifiers and ASN.1 encoding format
for the US NIST's PQC KEM (United States National Institute of Standards
and Technology's Post Quantum Cryptography Key Encapsulation Mechanism)
algorithms. The algorithms covered are Candidate 1 and Candidate 2. The
encoding for public key and private key is also provided.

--- middle

# Introduction

The US NIST PQC competition has selected the Candidate 1 and Candidate 2
algorithms as winners of their PQC competition {{PQCComp}}. These
algorithms are KEM algorithms. NIST has also defined object identifiers
for these algorithms (TODO insert reference).

This document specifies the use of the Candidate 1 and Candidate 2
algorithms in X.509 public key certifiates, see {{!RFC5280}}. It also
specifies private key encoding. An ASN.1 module is included for
reference purposes.


# Conventions and Definitions

{::boilerplate bcp14-tagged}


# Algorithm Identifiers

Certificates conforming to {{RFC5280}} can convey a public key for any
public key algorithm. The certificate indicates the algorithm through
an algorithm identifier. An algorithm identifier consists of an object
identifier and optional parameters.

The AlgorithmIdentifier type, which is included herein for convenience,
is defined as follows:

   AlgorithmIdentifier  ::=  SEQUENCE  {
       algorithm   OBJECT IDENTIFIER,
       parameters  ANY DEFINED BY algorithm OPTIONAL
   }

<aside markdown="block">
NOTE: The above syntax is from {{RFC5280}} and matches the version used
therein, i.e., the 1988 ASN.1 syntax. See {{!RFC5912}} for ASN.1
copmatible with the 2015 ASN.1 syntax.
</aside>

The fields in AlgorithmIdentifier have the following meanings:

* algorithm identifies the cryptographic algorithm with an object
  identifier. XXX such OIDs are defined in Sections {{candidate-1}} and
  {{candidate-2}}.

* parameters, which are optional, are the associated parameters for
  the algorithm identifier in the algorithm field.

In this document, TODO (specify number) new OIDs for identifying the
different algorithm and parameter pairs. For all of the object
identifiers, the parameters MUST be absent.

It is possible to find systems that require the parameters to be
present. This can be due to either a defect in the original 1997
syntax or a programming error where developers never got input where
this was not true. The optimal solution is to fix these systems;
where this is not possible, the problem needs to be restricted to
that subsystem and not propagated to the Internet.


# Candidate 1 {#candidate-1}

TODO insert object-identifiers


# Candidate 2 {#candidate-2}

TODO insert object identifiers

# Subject Public Key Fields

In the X.509 certificate, the subjectPublicKeyInfo field has the
SubjectPublicKeyInfo type, which has the following ASN.1 syntax:

~~~
  SubjectPublicKeyInfo  ::=  SEQUENCE  {
      algorithm         AlgorithmIdentifier,
      subjectPublicKey  BIT STRING
  }
~~~

<aside markdown="block">
NOTE: The above syntax is from {{RFC5280}} and matches the version used
therein, i.e., the 1988 ASN.1 syntax. See {{!RFC5912}} for ASN.1
copmatible with the 2015 ASN.1 syntax.
</aside>

The fields in SubjectPublicKeyInfo have the following meanings:

* algorithm is the algorithm identifier and parameters for the
  public key (see above).

* subjectPublicKey contains the byte stream of the public key.  The
  algorithms defined in this document always encode the public key
  as TODO pick format e.g., exact multiple of 8 bits?.

The following is an example of a TBD public key encoded using the
textual encoding defined in {{?RFC7468}}.

~~~
  -----BEGIN PUBLIC KEY-----
  TODO insert example public key
  -----END PUBLIC KEY-------
~~~

# Key Usage Bits

The intended application for the key is indicated in the keyUsage
certificate extension; see {{Section 4.2.1.3 of RFC5280}}.

If the keyUsage extension is present in a certificate that indicates
Candidate 1 or Candidate 2 in SubjectPublicKeyInfo, then the following
MUST be present:

~~~
  keyEncipherment;
~~~

all of the following MUST be present:

~~~
  encipherOnly; or
  decipherOnly.
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

The following is an example of a TBD private key encoded using the
textual encoding defined in {{RFC7468}}.

~~~
  -----BEGIN PRIVATE KEY-----
  TODO iser example private key
  -----END PRIVATE KEY-------
~~~

The following example, in addition to encoding the TBD private key,
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

This document updates {{?RFC3279}}. The Security Considerations section of that document therefore applies to this specification as well.

\[EDNOTE: Discuss side-channels for Dilithium.\]


# IANA Considerations

This document will have some IANA actions.


--- back

# Acknowledgments
{:numbered="false"}

TODO acknowledge.
