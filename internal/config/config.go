/*
Copyright 2025.

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

package config

import (
	//
	"os"
	"reflect"

	//
	"gopkg.in/yaml.v3"

	//
	"cluster-autoscaler-status-exporter/api/v1alpha1"
)

// Default values for the status config map
const (
	defaultStatusConfigMapNamespace = "kube-system"
	defaultStatusConfigMapName      = "cluster-autoscaler-status"
)

// ReadFile TODO
func ReadFile(filepath string) (config v1alpha1.ConfigSpec, err error) {
	var fileBytes []byte

	fileBytes, err = os.ReadFile(filepath)
	if err != nil {
		return config, err
	}

	// Expand environment variables present in the config
	// This will cause expansion in the following way: field: "$FIELD" -> field: "value_of_field"
	fileExpandedEnv := os.ExpandEnv(string(fileBytes))

	err = yaml.Unmarshal([]byte(fileExpandedEnv), &config)
	if err != nil {
		return config, err
	}

	// Load defaults if not set
	if reflect.ValueOf(config.StatusConfigMap).IsZero() {
		config.StatusConfigMap = v1alpha1.StatusConfigMapSpec{}
	}
	if config.StatusConfigMap.Name == "" {
		config.StatusConfigMap.Name = defaultStatusConfigMapName
	}
	if config.StatusConfigMap.Namespace == "" {
		config.StatusConfigMap.Namespace = defaultStatusConfigMapNamespace
	}

	return config, err
}
