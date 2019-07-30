// Copyright 2019 Canonical Ltd.
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package tpm2

import (
	"errors"
)

func (t *tpmContext) PolicySecret(authContext, policySession ResourceContext, cpHashA Digest, policyRef Nonce,
	expiration int32, authContextAuth interface{}, sessions ...*Session) (Timeout, *TkAuth, error) {
	if err := t.checkResourceContextParam(authContext, "authContext"); err != nil {
		return nil, nil, err
	}
	if err := t.checkResourceContextParam(policySession, "policySession"); err != nil {
		return nil, nil, err
	}

	sessionContext, isSession := policySession.(SessionContext)
	if !isSession {
		return nil, nil, errors.New("invalid resource context for policySession: not a session context")
	}

	var timeout Timeout
	var policyTicket TkAuth

	if err := t.RunCommand(CommandPolicySecret, ResourceWithAuth{Context: authContext, Auth: authContextAuth},
		policySession, Separator, sessionContext.NonceTPM(), cpHashA, policyRef, expiration, Separator,
		Separator, &timeout, &policyTicket, Separator, sessions); err != nil {
		return nil, nil, err
	}

	return timeout, &policyTicket, nil
}

func (t *tpmContext) PolicyOR(policySession ResourceContext, pHashList DigestList) error {
	if err := t.checkResourceContextParam(policySession, "policySession"); err != nil {
		return err
	}

	return t.RunCommand(CommandPolicyOR, policySession, Separator, pHashList)
}

func (t *tpmContext) PolicyPCR(policySession ResourceContext, pcrDigest Digest, pcrs PCRSelectionList,
	sessions ...*Session) error {
	if err := t.checkResourceContextParam(policySession, "policySession"); err != nil {
		return err
	}

	return t.RunCommand(CommandPolicyPCR, policySession, Separator, pcrDigest, pcrs, Separator, Separator,
		Separator, sessions)
}

func (t *tpmContext) PolicyAuthValue(policySession ResourceContext) error {
	if err := t.checkResourceContextParam(policySession, "policySession"); err != nil {
		return err
	}

	sc, isSessionContext := policySession.(*sessionContext)
	if !isSessionContext {
		return makeInvalidParamError("policySession", "not a session context")
	}

	if err := t.RunCommand(CommandPolicyAuthValue, policySession); err != nil {
		return err
	}

	sc.policyHMACType = policyHMACTypeAuth
	return nil
}

func (t *tpmContext) PolicyPassword(policySession ResourceContext) error {
	if err := t.checkResourceContextParam(policySession, "policySession"); err != nil {
		return err
	}

	sc, isSessionContext := policySession.(*sessionContext)
	if !isSessionContext {
		return makeInvalidParamError("policySession", "not a session context")
	}

	if err := t.RunCommand(CommandPolicyPassword, policySession); err != nil {
		return err
	}

	sc.policyHMACType = policyHMACTypePassword
	return nil
}

func (t *tpmContext) PolicyGetDigest(policySession ResourceContext) (Digest, error) {
	if err := t.checkResourceContextParam(policySession, "policySession"); err != nil {
		return nil, err
	}

	var policyDigest Digest

	if err := t.RunCommand(CommandPolicyGetDigest, policySession, Separator, Separator, Separator,
		&policyDigest); err != nil {
		return nil, err
	}

	return policyDigest, nil
}
