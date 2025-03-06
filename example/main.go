package main

import (
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"log"
	"os"

	"github.com/cloudflare/circl/kem/schemes"
)

type subjectPublicKeyInfo struct {
	Algorithm pkix.AlgorithmIdentifier
	PublicKey asn1.BitString
}

type MLKEMPrivateKeyBoth struct {
    Seed        []byte `asn1:"tag:0"`
    ExpandedKey []byte
}

type MLKEMPrivateKey struct {
    Seed        []byte              `asn1:"tag:0,optional"`
    ExpandedKey []byte              `asn1:"optional"`
    Both        MLKEMPrivateKeyBoth `asn1:"optional"`
}

type oneAsymmetricKey struct {
	Version    int
	Algorithm  pkix.AlgorithmIdentifier
	PrivateKey MLKEMPrivateKey
	Attributes []asn1.RawValue       `asn1:"tag:0,optional"`
	PublicKey  *subjectPublicKeyInfo `asn1:"tag:1,optional"`
}

func createPrivateKey(seed []byte, expandedKey []byte, format string) MLKEMPrivateKey {
	switch format {
    case "seed":
        return MLKEMPrivateKey{
            Seed: seed,
        }
    case "expanded":
        return MLKEMPrivateKey{
            ExpandedKey: expandedKey,
        }
    case "both":
        return MLKEMPrivateKey{
            Both: MLKEMPrivateKeyBoth{
                Seed:        seed,
                ExpandedKey: expandedKey,
            },
        }
    default:
        panic("Unknown format")
    }
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
		ask := oneAsymmetricKey{
			Version:    0,
			Algorithm:  alg,
			PrivateKey: createPrivateKey(seed, expandedKey, format),
		}

		pask, err := asn1.Marshal(ask)
		if err != nil {
			log.Fatal(err)
		}

		fsk, err := os.Create(fmt.Sprintf("%s-%s.priv", name, format))
		if err != nil {
			log.Fatal(err)
		}
		defer fsk.Close()

		if err = pem.Encode(fsk, &pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: pask,
		}); err != nil {
			log.Fatal(err)
		}
	}

	papk, err := asn1.Marshal(apk)
	if err != nil {
		log.Fatal(err)
	}

	fpk, err := os.Create(fmt.Sprintf("%s.pub", name))
	if err != nil {
		log.Fatal(err)
	}
	defer fpk.Close()

	if err = pem.Encode(fpk, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: papk,
	}); err != nil {
		log.Fatal(err)
	}
}

func main() {
	example("ML-KEM-512")
	example("ML-KEM-768")
	example("ML-KEM-1024")
}
