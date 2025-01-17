// Copyright 2021 Canonical Ltd.
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package util

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"

	"github.com/canonical/go-tpm2"
	"github.com/canonical/go-tpm2/templates"
)

// NewExternalRSAPublicKey creates a public area from the supplied RSA
// public key with the specified name algorithm, key usage and scheme,
// for use with the TPM2_LoadExternal command. If nameAlg is
// HashAlgorithmNull, then HashAlgorithmSHA256 is used. If no usage is
// specified, the public area will include both sign and decrypt attributes.
func NewExternalRSAPublicKey(nameAlg tpm2.HashAlgorithmId, usage templates.KeyUsage, scheme *tpm2.RSAScheme, key *rsa.PublicKey) *tpm2.Public {
	pub := templates.NewRSAKey(nameAlg, usage, scheme, uint16(len(key.N.Bytes())*8))
	pub.Attrs &^= (tpm2.AttrFixedTPM | tpm2.AttrFixedParent | tpm2.AttrSensitiveDataOrigin | tpm2.AttrUserWithAuth)
	pub.Params.RSADetail.Exponent = uint32(key.E)
	pub.Unique = &tpm2.PublicIDU{RSA: key.N.Bytes()}

	return pub
}

// NewExternalRSAPublicKeyWithDefaults creates a public area from the
// supplied RSA with the specified key usage, SHA256 as the name algorithm
// and the scheme unset, for use with the TPM2_LoadExternal command. If no
// usage is specified, the public area will include both sign and decrypt
// attributes.
func NewExternalRSAPublicKeyWithDefaults(usage templates.KeyUsage, key *rsa.PublicKey) *tpm2.Public {
	return NewExternalRSAPublicKey(tpm2.HashAlgorithmNull, usage, nil, key)
}

// NewExternalECCPublicKey creates a public area from the supplied
// elliptic public key with the specified name algorithm, key usage
// and scheme, for use with the TPM2_LoadExternal command. If nameAlg
// is HashAlgorithmNull, then HashAlgorithmSHA256 is used. If no usage is
// specified, the public area will include both sign and decrypt attributes.
func NewExternalECCPublicKey(nameAlg tpm2.HashAlgorithmId, usage templates.KeyUsage, scheme *tpm2.ECCScheme, key *ecdsa.PublicKey) *tpm2.Public {
	var curve tpm2.ECCCurve
	switch key.Curve {
	case elliptic.P224():
		curve = tpm2.ECCCurveNIST_P224
	case elliptic.P256():
		curve = tpm2.ECCCurveNIST_P256
	case elliptic.P384():
		curve = tpm2.ECCCurveNIST_P384
	case elliptic.P521():
		curve = tpm2.ECCCurveNIST_P521
	default:
		panic("unsupported curve")
	}

	pub := templates.NewECCKey(nameAlg, usage, scheme, curve)
	pub.Attrs &^= (tpm2.AttrFixedTPM | tpm2.AttrFixedParent | tpm2.AttrSensitiveDataOrigin | tpm2.AttrUserWithAuth)
	pub.Unique = &tpm2.PublicIDU{
		ECC: &tpm2.ECCPoint{
			X: zeroExtendBytes(key.X, key.Params().BitSize/8),
			Y: zeroExtendBytes(key.Y, key.Params().BitSize/8)}}

	return pub
}

// NewExternalECCPublicKeyWithDefaults creates a public area from the
// supplied elliptic public key with the specified key usage, SHA256
// as the name algorithm and the scheme unset, for use with the
// TPM2_LoadExternal command. If no usage is specified, the public area
// will include both sign and decrypt attributes.
func NewExternalECCPublicKeyWithDefaults(usage templates.KeyUsage, key *ecdsa.PublicKey) *tpm2.Public {
	return NewExternalECCPublicKey(tpm2.HashAlgorithmNull, usage, nil, key)
}

// NewSealedObject creates both the public and sensitive areas for a
// sealed object containing the supplied data, with the specified name
// algorithm and authorization value. If nameAlgorithm is HashAlgorithmNull,
// then HashAlgorithmSHA256 is used.
//
// It will panic if authValue is larger than the size of the name algorithm.
//
// The returned public and sensitive areas can be made into a duplication
// object with CreateDuplicationObjectFromSensitive for importing into a TPM.
//
// The public area has the AttrUserWithAuth set in order to permit authentication
// for the user auth role using the sensitive area's authorization value. In order
// to require authentication for the user auth role using an authorization policy,
// remove the AttrUserWithAuth attribute.
func NewSealedObject(nameAlg tpm2.HashAlgorithmId, authValue tpm2.Auth, data []byte) (*tpm2.Public, *tpm2.Sensitive) {
	pub := templates.NewSealedObject(nameAlg)
	pub.Attrs &^= (tpm2.AttrFixedTPM | tpm2.AttrFixedParent)

	if len(authValue) > pub.NameAlg.Size() {
		panic("authValue too large")
	}

	sensitive := &tpm2.Sensitive{
		Type:      tpm2.ObjectTypeKeyedHash,
		AuthValue: make(tpm2.Auth, pub.NameAlg.Size()),
		SeedValue: make(tpm2.Digest, pub.NameAlg.Size()),
		Sensitive: &tpm2.SensitiveCompositeU{Bits: data}}
	copy(sensitive.AuthValue, authValue)
	rand.Read(sensitive.SeedValue)

	h := nameAlg.NewHash()
	h.Write(sensitive.SeedValue)
	h.Write(sensitive.Sensitive.Bits)
	pub.Unique = &tpm2.PublicIDU{KeyedHash: h.Sum(nil)}

	return pub, sensitive
}
