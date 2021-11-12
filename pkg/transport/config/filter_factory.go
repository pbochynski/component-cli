// SPDX-FileCopyrightText: 2021 SAP SE or an SAP affiliate company and Gardener contributors.
//
// SPDX-License-Identifier: Apache-2.0
package config

import (
	"encoding/json"
	"fmt"

	"sigs.k8s.io/yaml"

	"github.com/gardener/component-cli/pkg/transport/filters"
)

const (
	// ComponentNameFilterType defines the type of a component name filter
	ComponentNameFilterType = "ComponentNameFilter"

	// ResourceTypeFilterType defines the type of a resource type filter
	ResourceTypeFilterType = "ResourceTypeFilter"

	// ResourceAccessTypeFilterType defines the type of a resource access filter
	ResourceAccessTypeFilterType = "ResourceAccessTypeFilter"
)

// NewFilterFactory creates a new filter factory
func NewFilterFactory() *FilterFactory {
	return &FilterFactory{}
}

// FilterFactory defines a helper struct for creating filters
type FilterFactory struct{}

// Create creates a new filter defined by a type and a spec
func (f *FilterFactory) Create(filterType string, spec *json.RawMessage) (filters.Filter, error) {
	switch filterType {
	case ComponentNameFilterType:
		return f.createComponentNameFilter(spec)
	case ResourceTypeFilterType:
		return f.createResourceTypeFilter(spec)
	case ResourceAccessTypeFilterType:
		return f.createAccessTypeFilter(spec)
	default:
		return nil, fmt.Errorf("unknown filter type %s", filterType)
	}
}

func (f *FilterFactory) createComponentNameFilter(rawSpec *json.RawMessage) (filters.Filter, error) {
	type filterSpec struct {
		IncludeComponentNames []string `json:"includeComponentNames"`
	}

	var spec filterSpec
	err := yaml.Unmarshal(*rawSpec, &spec)
	if err != nil {
		return nil, fmt.Errorf("unable to parse spec: %w", err)
	}

	return filters.NewComponentNameFilter(spec.IncludeComponentNames...)
}

func (f *FilterFactory) createResourceTypeFilter(rawSpec *json.RawMessage) (filters.Filter, error) {
	type filterSpec struct {
		IncludeResourceTypes []string `json:"includeResourceTypes"`
	}

	var spec filterSpec
	err := yaml.Unmarshal(*rawSpec, &spec)
	if err != nil {
		return nil, fmt.Errorf("unable to parse spec: %w", err)
	}

	return filters.NewResourceTypeFilter(spec.IncludeResourceTypes...)
}

func (f *FilterFactory) createAccessTypeFilter(rawSpec *json.RawMessage) (filters.Filter, error) {
	type filterSpec struct {
		IncludeAccessTypes []string `json:"includeAccessTypes"`
	}

	var spec filterSpec
	err := yaml.Unmarshal(*rawSpec, &spec)
	if err != nil {
		return nil, fmt.Errorf("unable to parse spec: %w", err)
	}

	return filters.NewResourceAccessTypeFilter(spec.IncludeAccessTypes...)
}
