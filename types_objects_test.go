// Copyright 2019 Canonical Ltd.
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package tpm2_test

import (
	"bytes"
	"reflect"
	"testing"

	. "github.com/canonical/go-tpm2"
	"github.com/canonical/go-tpm2/mu"
	"github.com/canonical/go-tpm2/testutil"

	. "gopkg.in/check.v1"
)

type typesObjectsSuite struct{}

var _ = Suite(&typesObjectsSuite{})

func (s *typesObjectsSuite) TestPublicIsStorageRSAValid(c *C) {
	pub := Public{
		Type:    ObjectTypeRSA,
		NameAlg: HashAlgorithmSHA256,
		Attrs:   AttrRestricted | AttrDecrypt,
		Params: &PublicParamsU{
			RSADetail: &RSAParams{
				Symmetric: SymDefObject{
					Algorithm: SymObjectAlgorithmAES,
					KeyBits:   &SymKeyBitsU{Sym: 128},
					Mode:      &SymModeU{Sym: SymModeCFB},
				},
				Scheme:   RSAScheme{Scheme: RSASchemeNull},
				KeyBits:  2048,
				Exponent: 0}}}
	c.Check(pub.IsStorage(), testutil.IsTrue)
}

func (s *typesObjectsSuite) TestPublicIsStorageECCValid(c *C) {
	pub := Public{
		Type:    ObjectTypeECC,
		NameAlg: HashAlgorithmSHA256,
		Attrs:   AttrRestricted | AttrDecrypt,
		Params: &PublicParamsU{
			ECCDetail: &ECCParams{
				Symmetric: SymDefObject{
					Algorithm: SymObjectAlgorithmAES,
					KeyBits:   &SymKeyBitsU{Sym: 128},
					Mode:      &SymModeU{Sym: SymModeCFB},
				},
				Scheme:  ECCScheme{Scheme: ECCSchemeNull},
				CurveID: ECCCurveNIST_P256,
				KDF:     KDFScheme{Scheme: KDFAlgorithmNull}}}}
	c.Check(pub.IsStorage(), testutil.IsTrue)
}

func (s *typesObjectsSuite) TestPublicIsStorageRSASign(c *C) {
	pub := Public{
		Type:    ObjectTypeRSA,
		NameAlg: HashAlgorithmSHA256,
		Attrs:   AttrRestricted | AttrSign,
		Params: &PublicParamsU{
			RSADetail: &RSAParams{
				Symmetric: SymDefObject{Algorithm: SymObjectAlgorithmNull},
				Scheme: RSAScheme{
					Scheme: RSASchemeRSAPSS,
					Details: &AsymSchemeU{
						RSAPSS: &SigSchemeRSAPSS{HashAlg: HashAlgorithmSHA256},
					},
				},
				KeyBits:  2048,
				Exponent: 0}}}
	c.Check(pub.IsStorage(), testutil.IsFalse)
}

func (s *typesObjectsSuite) TestPublicIsStorageSymmetric(c *C) {
	pub := Public{
		Type:    ObjectTypeSymCipher,
		NameAlg: HashAlgorithmSHA256,
		Attrs:   AttrRestricted | AttrDecrypt,
		Params: &PublicParamsU{
			SymDetail: &SymCipherParams{
				Sym: SymDefObject{
					Algorithm: SymObjectAlgorithmAES,
					KeyBits:   &SymKeyBitsU{Sym: 128},
					Mode:      &SymModeU{Sym: SymModeCFB},
				}}}}
	c.Check(pub.IsStorage(), testutil.IsFalse)
}

func (s *typesObjectsSuite) TestPublicIsStorageKeyedHash(c *C) {
	pub := Public{
		Type:    ObjectTypeKeyedHash,
		NameAlg: HashAlgorithmSHA256,
		Attrs:   AttrRestricted | AttrDecrypt,
		Params: &PublicParamsU{
			KeyedHashDetail: &KeyedHashParams{
				Scheme: KeyedHashScheme{
					Scheme: KeyedHashSchemeXOR,
					Details: &SchemeKeyedHashU{
						XOR: &SchemeXOR{
							HashAlg: HashAlgorithmSHA256,
							KDF:     KDFAlgorithmKDF1_SP800_108}}}}}}
	c.Check(pub.IsStorage(), testutil.IsFalse)
}

func (s *typesObjectsSuite) TestPublicIsParentRSAValid(c *C) {
	pub := Public{
		Type:    ObjectTypeRSA,
		NameAlg: HashAlgorithmSHA256,
		Attrs:   AttrRestricted | AttrDecrypt,
		Params: &PublicParamsU{
			RSADetail: &RSAParams{
				Symmetric: SymDefObject{
					Algorithm: SymObjectAlgorithmAES,
					KeyBits:   &SymKeyBitsU{Sym: 128},
					Mode:      &SymModeU{Sym: SymModeCFB},
				},
				Scheme:   RSAScheme{Scheme: RSASchemeNull},
				KeyBits:  2048,
				Exponent: 0}}}
	c.Check(pub.IsParent(), testutil.IsTrue)
}

func (s *typesObjectsSuite) TestPublicIsParentECCValid(c *C) {
	pub := Public{
		Type:    ObjectTypeECC,
		NameAlg: HashAlgorithmSHA256,
		Attrs:   AttrRestricted | AttrDecrypt,
		Params: &PublicParamsU{
			ECCDetail: &ECCParams{
				Symmetric: SymDefObject{
					Algorithm: SymObjectAlgorithmAES,
					KeyBits:   &SymKeyBitsU{Sym: 128},
					Mode:      &SymModeU{Sym: SymModeCFB},
				},
				Scheme:  ECCScheme{Scheme: ECCSchemeNull},
				CurveID: ECCCurveNIST_P256,
				KDF:     KDFScheme{Scheme: KDFAlgorithmNull}}}}
	c.Check(pub.IsParent(), testutil.IsTrue)
}

func (s *typesObjectsSuite) TestPublicIsParentSymmetric(c *C) {
	pub := Public{
		Type:    ObjectTypeSymCipher,
		NameAlg: HashAlgorithmSHA256,
		Attrs:   AttrRestricted | AttrDecrypt,
		Params: &PublicParamsU{
			SymDetail: &SymCipherParams{
				Sym: SymDefObject{
					Algorithm: SymObjectAlgorithmAES,
					KeyBits:   &SymKeyBitsU{Sym: 128},
					Mode:      &SymModeU{Sym: SymModeCFB},
				}}}}
	c.Check(pub.IsParent(), testutil.IsTrue)
}

func (s *typesObjectsSuite) TestPublicIsParentKeyedHash(c *C) {
	pub := Public{
		Type:    ObjectTypeKeyedHash,
		NameAlg: HashAlgorithmSHA256,
		Attrs:   AttrRestricted | AttrDecrypt,
		Params: &PublicParamsU{
			KeyedHashDetail: &KeyedHashParams{
				Scheme: KeyedHashScheme{
					Scheme: KeyedHashSchemeXOR,
					Details: &SchemeKeyedHashU{
						XOR: &SchemeXOR{
							HashAlg: HashAlgorithmSHA256,
							KDF:     KDFAlgorithmKDF1_SP800_108}}}}}}
	c.Check(pub.IsParent(), testutil.IsFalse)
}

func (s *typesObjectsSuite) TestPublicIsParentRSASign(c *C) {
	pub := Public{
		Type:    ObjectTypeRSA,
		NameAlg: HashAlgorithmSHA256,
		Attrs:   AttrRestricted | AttrSign,
		Params: &PublicParamsU{
			RSADetail: &RSAParams{
				Symmetric: SymDefObject{Algorithm: SymObjectAlgorithmNull},
				Scheme: RSAScheme{
					Scheme: RSASchemeRSAPSS,
					Details: &AsymSchemeU{
						RSAPSS: &SigSchemeRSAPSS{HashAlg: HashAlgorithmSHA256},
					},
				},
				KeyBits:  2048,
				Exponent: 0}}}
	c.Check(pub.IsParent(), testutil.IsFalse)
}

func (s *typesObjectsSuite) TestPublicIsParentRSANoNameAlg(c *C) {
	pub := Public{
		Type:    ObjectTypeRSA,
		NameAlg: HashAlgorithmNull,
		Attrs:   AttrRestricted | AttrDecrypt,
		Params: &PublicParamsU{
			RSADetail: &RSAParams{
				Symmetric: SymDefObject{
					Algorithm: SymObjectAlgorithmAES,
					KeyBits:   &SymKeyBitsU{Sym: 128},
					Mode:      &SymModeU{Sym: SymModeCFB},
				},
				Scheme:   RSAScheme{Scheme: RSASchemeNull},
				KeyBits:  2048,
				Exponent: 0}}}
	c.Check(pub.IsParent(), testutil.IsFalse)
}

type TestPublicIDUContainer struct {
	Alg    ObjectTypeId
	Unique *PublicIDU
}

func TestPublicIDUnion(t *testing.T) {
	for _, data := range []struct {
		desc string
		in   TestPublicIDUContainer
		out  []byte
		err  string
	}{
		{
			desc: "RSA",
			in: TestPublicIDUContainer{Alg: ObjectTypeRSA,
				Unique: &PublicIDU{RSA: PublicKeyRSA{0x01, 0x02, 0x03}}},
			out: []byte{0x00, 0x01, 0x00, 0x03, 0x01, 0x02, 0x03},
		},
		{
			desc: "KeyedHash",
			in: TestPublicIDUContainer{Alg: ObjectTypeKeyedHash,
				Unique: &PublicIDU{KeyedHash: Digest{0x04, 0x05, 0x06, 0x07}}},
			out: []byte{0x00, 0x08, 0x00, 0x04, 0x04, 0x05, 0x06, 0x07},
		},
		{
			desc: "InvalidSelector",
			in: TestPublicIDUContainer{Alg: ObjectTypeId(AlgorithmNull),
				Unique: &PublicIDU{Sym: Digest{0x04, 0x05, 0x06, 0x07}}},
			out: []byte{0x00, 0x10},
			err: "cannot unmarshal argument whilst processing element of type tpm2.PublicIDU: invalid selector value: TPM_ALG_NULL\n\n" +
				"=== BEGIN STACK ===\n" +
				"... tpm2_test.TestPublicIDUContainer field Unique\n" +
				"=== END STACK ===\n",
		},
	} {
		t.Run(data.desc, func(t *testing.T) {
			out, err := mu.MarshalToBytes(data.in)
			if err != nil {
				t.Fatalf("MarshalToBytes failed: %v", err)
			}

			if !bytes.Equal(out, data.out) {
				t.Fatalf("MarshalToBytes returned an unexpected byte sequence: %x", out)
			}

			var a TestPublicIDUContainer
			n, err := mu.UnmarshalFromBytes(out, &a)
			if data.err != "" {
				if err == nil {
					t.Fatalf("UnmarshalFromBytes was expected to fail")
				}
				if err.Error() != data.err {
					t.Errorf("UnmarshalFromBytes returned an unexpected error: %v", err)
				}
			} else {
				if err != nil {
					t.Fatalf("UnmarshalFromBytes failed: %v", err)
				}
				if n != len(out) {
					t.Errorf("UnmarshalFromBytes consumed the wrong number of bytes (%d)", n)
				}

				if !reflect.DeepEqual(data.in, a) {
					t.Errorf("UnmarshalFromBytes didn't return the original data")
				}
			}
		})
	}
}

func TestPublicName(t *testing.T) {
	tpm, _, closeTPM := testutil.NewTPMContextT(t, testutil.TPMFeatureOwnerHierarchy)
	defer closeTPM()

	primary := createRSASrkForTesting(t, tpm, nil)
	defer flushContext(t, tpm, primary)

	pub, _, _, err := tpm.ReadPublic(primary)
	if err != nil {
		t.Fatalf("ReadPublic failed: %v", err)
	}

	name, err := pub.Name()
	if err != nil {
		t.Fatalf("Public.Name() failed: %v", err)
	}

	// primary.Name() is what the TPM returned at object creation
	if !bytes.Equal(primary.Name(), name) {
		t.Errorf("Public.Name() returned an unexpected name")
	}
}