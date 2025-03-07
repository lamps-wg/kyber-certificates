package main

import (
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/cloudflare/circl/kem/schemes"
)

type subjectPublicKeyInfo struct {
	Algorithm pkix.AlgorithmIdentifier
	PublicKey asn1.BitString
}

type mlkemPrivateKey struct {
	Version    int
	Algorithm  pkix.AlgorithmIdentifier
	PrivateKey []byte
}

func generatePrivateKeyBytes(format string, seed []byte, expandedKey []byte) ([]byte, error) {
	switch format {
	case "seed":
		// Create [0] OCTET STRING structure
		seedValue := asn1.RawValue{
			Class:      asn1.ClassContextSpecific,
			Tag:        0,
			Bytes:      seed,
			IsCompound: false,
		}
		return asn1.Marshal(seedValue)
	case "expanded":
		return asn1.Marshal(expandedKey)
	case "both":
		sequence := struct {
			Seed        []byte
			ExpandedKey []byte
		}{
			Seed:        seed,
			ExpandedKey: expandedKey,
		}
		return asn1.Marshal(sequence)
	}
	return nil, fmt.Errorf("unknown format")
}

func generatePrivateKey(format string, alg pkix.AlgorithmIdentifier, seed []byte, expandedKey []byte) (mlkemPrivateKey, error) {
	ask := mlkemPrivateKey{
		Version:   0,
		Algorithm: alg,
	}

	// Generate the inner CHOICE structure
	innerBytes, err := generatePrivateKeyBytes(format, seed, expandedKey)
	if err != nil {
		return ask, err
	}

	// Set as the private key bytes
	ask.PrivateKey = innerBytes
	return ask, nil
}

func example(name string) {
	scheme := schemes.ByName(name)
	seed := make([]byte, scheme.SeedSize())
	if len(seed) != 64 {
		panic("Unexpected seed size")
	}
	for i := 0; i < len(seed); i++ {
		seed[i] = byte(i)
	}
	pk, sk := scheme.DeriveKeyPair(seed)

	expandedKey, _ := sk.MarshalBinary()

	ppk, _ := pk.MarshalBinary()

	var oid int
	switch name {
	case "ML-KEM-512":
		oid = 1
	case "ML-KEM-768":
		oid = 2
	case "ML-KEM-1024":
		oid = 3
	default:
		panic("unknown")
	}

	alg := pkix.AlgorithmIdentifier{
		Algorithm: asn1.ObjectIdentifier{2, 16, 840, 1, 101, 3, 4, 4, oid},
	}

	apk := subjectPublicKeyInfo{
		Algorithm: alg,
		PublicKey: asn1.BitString{
			BitLength: len(ppk) * 8,
			Bytes:     ppk,
		},
	}

	formats := []string{"seed", "expanded", "both"}
	for _, format := range formats {
		ask, err := generatePrivateKey(format, alg, seed, expandedKey)
		if err != nil {
			panic(err)
		}

		pask, err := asn1.Marshal(ask)
		if err != nil {
			panic(err)
		}

		fsk, err := os.Create(fmt.Sprintf("%s-%s.priv", name, format))
		if err != nil {
			panic(err)
		}
		defer fsk.Close()

		if err = pem.Encode(fsk, &pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: pask,
		}); err != nil {
			panic(err)
		}
	}

	papk, err := asn1.Marshal(apk)
	if err != nil {
		panic(err)
	}

	fpk, err := os.Create(fmt.Sprintf("%s.pub", name))
	if err != nil {
		panic(err)
	}
	defer fpk.Close()

	if err = pem.Encode(fpk, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: papk,
	}); err != nil {
		panic(err)
	}
}

func main() {
	example("ML-KEM-512")
	example("ML-KEM-768")
	example("ML-KEM-1024")
}
