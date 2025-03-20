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

package kubernetes

import (

	//
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
)

// NewKubernetesClient return a new Kubernetes client from client-go SDK
func NewKubernetesClient() (client *kubernetes.Clientset, err error) {
	config, err := ctrl.GetConfig()
	if err != nil {
		return client, err
	}

	// Core client
	client, err = kubernetes.NewForConfig(config)
	if err != nil {
		return client, err
	}

	return client, err
}
