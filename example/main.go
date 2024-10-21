package main

import (
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"log"
	"os"

	"github.com/cloudflare/circl/kem/schemes"
)

type subjectPublicKeyInfo struct {
	Algorithm pkix.AlgorithmIdentifier
	PublicKey asn1.BitString
}

type oneAsymmetricKey struct {
	Version    int
	Algorithm  pkix.AlgorithmIdentifier
	PrivateKey []byte
	Attributes []asn1.RawValue       `asn1:"tag:0,optional"`
	PublicKey  *subjectPublicKeyInfo `asn1:"tag:1,optional"`
}

func main() {
	scheme := schemes.ByName("ML-KEM-768")
	seed := make([]byte, scheme.SeedSize())
	pk, sk := scheme.DeriveKeyPair(seed)

	ppk, _ := pk.MarshalBinary()
	psk, _ := sk.MarshalBinary()

	apk := subjectPublicKeyInfo{
		Algorithm: pkix.AlgorithmIdentifier{
			Algorithm: asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 4, 3},
		},
		PublicKey: asn1.BitString{
			BitLength: len(ppk) * 8,
			Bytes:     ppk,
		},
	}

	ask := oneAsymmetricKey{
		Version: 0,
		Algorithm: pkix.AlgorithmIdentifier{
			Algorithm: asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 4, 3},
		},
		PrivateKey: psk,
	}

	papk, err := asn1.Marshal(apk)
	if err != nil {
		log.Fatal(err)
	}

	pask, err := asn1.Marshal(ask)
	if err != nil {
		log.Fatal(err)
	}

	if err = pem.Encode(os.Stdout, &pem.Block{
		Type:  "ML-KEM-768 PRIVATE KEY",
		Bytes: pask,
	}); err != nil {
		log.Fatal(err)
	}

	if err = pem.Encode(os.Stdout, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: papk,
	}); err != nil {
		log.Fatal(err)
	}
}
