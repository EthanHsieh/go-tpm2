package tpm2

import (
	"encoding/binary"
	"fmt"
)

type ResourceContext interface {
	Handle() Handle
	Name() Name
}

type resourceContextPrivate interface {
	Tpm() *tpmContext
	SetTpm(t *tpmContext)
}

type permanentContext struct {
	tpm    *tpmContext
	handle Handle
}

func (r *permanentContext) Handle() Handle {
	return r.handle
}

func (r *permanentContext) Name() Name {
	name := make([]byte, 4)
	binary.BigEndian.PutUint32(name, uint32(r.handle))
	return Name(name)
}

func (r *permanentContext) Tpm() *tpmContext {
	return r.tpm
}

func (r *permanentContext) SetTpm(t *tpmContext) {
	r.tpm = t
}

type objectContext struct {
	tpm    *tpmContext
	handle Handle
	public Public
	name   Name
}

func (r *objectContext) Handle() Handle {
	return r.handle
}

func (r *objectContext) Name() Name {
	return r.name
}

func (r *objectContext) Tpm() *tpmContext {
	return r.tpm
}

func (r *objectContext) SetTpm(t *tpmContext) {
	r.tpm = t
}

type nvIndexContext struct {
	tpm    *tpmContext
	handle Handle
	name   Name
}

func (r *nvIndexContext) Handle() Handle {
	return r.handle
}

func (r *nvIndexContext) Name() Name {
	return r.name
}

func (r *nvIndexContext) Tpm() *tpmContext {
	return r.tpm
}

func (r *nvIndexContext) SetTpm(t *tpmContext) {
	r.tpm = t
}

type sessionContext struct {
	tpm           *tpmContext
	handle        Handle
	hashAlg       AlgorithmId
	boundResource ResourceContext
	sessionKey    []byte
	nonceCaller   Nonce
	nonceTPM      Nonce
}

func (r *sessionContext) Handle() Handle {
	return r.handle
}

func (r *sessionContext) Name() Name {
	return nil
}

func (r *sessionContext) Tpm() *tpmContext {
	return r.tpm
}

func (r *sessionContext) SetTpm(t *tpmContext) {
	r.tpm = t
}

func makeNVIndexContext(t *tpmContext, handle Handle) (ResourceContext, error) {
	_, name, err := t.nvReadPublic(handle)
	if err != nil {
		return nil, err
	}
	return &nvIndexContext{handle: handle, name: name}, nil
}

func makeObjectContext(t *tpmContext, handle Handle) (ResourceContext, error) {
	pub, name, _, err := t.readPublic(handle)
	if err != nil {
		return nil, err
	}
	return &objectContext{handle: handle, public: *pub, name: name}, nil
}

func (t *tpmContext) evictResourceContext(rc ResourceContext) {
	rcp := rc.(resourceContextPrivate)
	if rcp.Tpm() == nil {
		return
	}
	if rcp.Tpm() != t {
		panic("Attempting to evict a resource for another TPM instance")
	}
	rcp.SetTpm(nil)
	delete(t.resources, rc.Handle())
}

func (t *tpmContext) addResourceContext(rc ResourceContext) {
	rcp := rc.(resourceContextPrivate)
	if rcp.Tpm() != nil {
		panic("Attempting to add a resource to more than one TPM instance")
	}
	if _, exists := t.resources[rc.Handle()]; exists {
		panic("Resource object for handle already exists")
	}
	rcp.SetTpm(t)
	t.resources[rc.Handle()] = rc

}

func (t *tpmContext) checkResourceContextParam(rc ResourceContext, name string) error {
	if rc == nil {
		return fmt.Errorf("invalid resource context for %s: nil", name)
	}
	rcp := rc.(resourceContextPrivate)
	if rcp.Tpm() == nil {
		return fmt.Errorf("invalid resource context for %s: resource has been closed", name)
	}
	if rcp.Tpm() != t {
		return fmt.Errorf("invalid resource context for %s: resource belongs to another tpm2.TPM "+
			"instance", name)
	}
	return nil
}

func (t *tpmContext) WrapHandle(handle Handle) (ResourceContext, error) {
	if rc, exists := t.resources[handle]; exists {
		return rc, nil
	}

	var rc ResourceContext
	var err error

	switch (handle & 0xff000000) >> 24 {
	case 0x00:
		// PCR
	case 0x01:
		rc, err = makeNVIndexContext(t, handle)
	case 0x02:
		// HMAC session
	case 0x03:
		// Policy session
	case 0x40:
		rc = &permanentContext{handle: handle}
	case 0x80, 0x81:
		rc, err = makeObjectContext(t, handle)
	}

	if err != nil {
		return nil, err
	}

	t.addResourceContext(rc)

	return rc, nil
}
