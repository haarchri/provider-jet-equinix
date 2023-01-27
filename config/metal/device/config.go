/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package device

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure the device group with references to other resources
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("equinix_metal_device", func(r *config.Resource) {
		r.LateInitializer = config.LateInitializer{
			// NOTE(displague): These are ignored because they conflict with each other.
			// See the following for more details: https://github.com/upbound/upjet/issues/107
			IgnoredFields: []string{
				"metro",
				"facilities",
			},
		}
	})
}
