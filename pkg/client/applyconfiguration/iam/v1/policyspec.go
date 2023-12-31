/*
Copyright 2019-2021 OpenFaaS Authors

Licensed under the MIT license. See LICENSE file in the project root for full license information.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// PolicySpecApplyConfiguration represents an declarative configuration of the PolicySpec type for use
// with apply.
type PolicySpecApplyConfiguration struct {
	Statement []PolicyStatementApplyConfiguration `json:"statement,omitempty"`
}

// PolicySpecApplyConfiguration constructs an declarative configuration of the PolicySpec type for use with
// apply.
func PolicySpec() *PolicySpecApplyConfiguration {
	return &PolicySpecApplyConfiguration{}
}

// WithStatement adds the given value to the Statement field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Statement field.
func (b *PolicySpecApplyConfiguration) WithStatement(values ...*PolicyStatementApplyConfiguration) *PolicySpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithStatement")
		}
		b.Statement = append(b.Statement, *values[i])
	}
	return b
}
