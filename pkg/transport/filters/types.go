// SPDX-FileCopyrightText: 2021 SAP SE or an SAP affiliate company and Gardener contributors.
//
// SPDX-License-Identifier: Apache-2.0
package filters

import (
	cdv2 "github.com/gardener/component-spec/bindings-go/apis/v2"
)

// Filter
type Filter interface {
	// Matches matches a component descriptor and a resource against the filter
	Matches(cdv2.ComponentDescriptor, cdv2.Resource) bool
}
