package msfidosignatureassertion

import (
	"github.com/matthewmueller/joy/dom/msassertion"
	"github.com/matthewmueller/joy/dom/mscredentialtype"
	"github.com/matthewmueller/joy/dom/msfidosignature"
	"github.com/matthewmueller/joy/js"
)

var _ msassertion.MSAssertion = (*MSFIDOSignatureAssertion)(nil)

// MSFIDOSignatureAssertion struct
// js:"MSFIDOSignatureAssertion,omit"
type MSFIDOSignatureAssertion struct {
}

// Signature prop
// js:"signature"
func (*MSFIDOSignatureAssertion) Signature() (signature *msfidosignature.MSFIDOSignature) {
	js.Rewrite("$_.signature")
	return signature
}

// ID prop
// js:"id"
func (*MSFIDOSignatureAssertion) ID() (id string) {
	js.Rewrite("$_.id")
	return id
}

// Type prop
// js:"type"
func (*MSFIDOSignatureAssertion) Type() (kind *mscredentialtype.MSCredentialType) {
	js.Rewrite("$_.type")
	return kind
}