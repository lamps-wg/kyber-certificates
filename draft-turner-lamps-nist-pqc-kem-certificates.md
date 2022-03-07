---
title: Algorithm Identifiers for NIST's PQC Algorithms for Use in the Internet X.509 Public Key Infrastructure
abbrev: PQC KEM for Certificates
category: std

docname: draft-turner-lamps-nist-pqc-kem-certificates
ipr: trust200902
keyword: Internet-Draft
area: SEC
workgroup: None
venue:
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

informative:
  PQCProj:
    target: https://csrc.nist.gov/projects/post-quantum-cryptography
    title: Post-Quantum Cryptography Project
    author:
      - org: National Insititue of Standards and Technology
    date: 2016-12-20


--- abstract

This document specifies algorithm identifiers and ASN.1 encoding format
for the US NIST's PQC KEM (United States National Institute of Standards
and Technology's Post Quantum Cryptography Key Encapsulation Mechanism)
algorithms. The algorithms covered are Candidate TBD1. The
encoding for public key and private key is also provided.

/[EDNOTE:
This draft is not expected to be finalized before the NIST PQC Project
has standardized PQ algorithms. After NIST has standardized its first
algorithms, this document will replace TBD, with the appropriate
algorithms and parameters before proceeding to ratification. The
algorithm Dilithium has been added as an example in this draft, to
provide a more detailed illustration of the content - it by no means
indicates its inclusion in the final version. This specification will
use object identifiers for the new algorithms that are assigned by NIST,
and will use placeholders until these are released.]

--- middle

# Introduction

The US NIST PQC Project has selected the Candidate TBD1 
algorithms as winners of their PQC Project {{PQCProj}}. These
algorithms are KEM algorithms. NIST has also defined object identifiers
for these algorithms (TODO insert reference).

This document specifies the use of the Candidate TBD1 
algorithms in X.509 public key certifiates, see {{!RFC5280}}. 
It also specifies private key encoding. 
An ASN.1 module is included for reference purposes.

These certificates could be used as Issuers in CMS where the public key 
is used to encapsulate a shared secret used to derive a symmetric key 
used to encrypt content in CMS 
\[EDNOTE: Add reference draft-perret-prat-lamps-cms-pq-kem\]. 
To be used in TLS, these certificates could only be used as end-entity 
identity certificates and would require significant updates to the
protocol 
\[EDNOTE: Add reference draft-celi-wiggers-tls-authkem\]. 

# Conventions and Definitions

{::boilerplate bcp14-tagged}


# Algorithm Identifiers

Certificates conforming to {{RFC5280}} can convey a public key for any
public key algorithm. The certificate indicates the algorithm through
an algorithm identifier. An algorithm identifier consists of an object
identifier and optional parameters.

The AlgorithmIdentifier type, which is included herein for convenience,
is defined as follows:

~~~
   AlgorithmIdentifier  ::=  SEQUENCE  {
       algorithm   OBJECT IDENTIFIER,
       parameters  ANY DEFINED BY algorithm OPTIONAL
   }
~~~

<aside markdown="block">
NOTE: The above syntax is from {{RFC5280}} and matches the version used
therein, i.e., the 1988 ASN.1 syntax. See {{!RFC5912}} for ASN.1
copmatible with the 2015 ASN.1 syntax.
</aside>

The fields in AlgorithmIdentifier have the following meanings:

* algorithm identifies the cryptographic algorithm with an object
  identifier. XXX such OIDs are defined in Sections {{candidate-TBD1}}.

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


# Candidate TBD1 {#candidate-TBD1}

TODO insert object-identifiers
   

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
Candidate TBD1 in SubjectPublicKeyInfo, then the following
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
