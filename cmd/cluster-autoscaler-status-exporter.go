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

package main

import (
	//
	"context"
	"flag"
	"log"

	//
	"go.uber.org/zap"

	//
	"cluster-autoscaler-status-exporter/api/v1alpha1"
	"cluster-autoscaler-status-exporter/internal/config"
	"cluster-autoscaler-status-exporter/internal/kubernetes"
	"cluster-autoscaler-status-exporter/internal/logging"
	"cluster-autoscaler-status-exporter/internal/server"
)

func main() {

	// Configure application's context
	app := v1alpha1.ApplicationSpec{
		Config:  &v1alpha1.ConfigSpec{},
		Context: context.Background(),
	}

	// Get configPath from command line arguments
	configPath := flag.String("config", "config.yaml", "config file path")

	// Parse arguments
	flag.Parse()

	// Get and parse the config
	configContent, err := config.ReadFile(*configPath)
	if err != nil {
		log.Fatalf("Error parsing configuration file: %v", err)
	}
	app.Config = &configContent

	// Configure logging
	err = logging.ConfigureLogger(&app)
	if err != nil {
		log.Fatalf("Error configuring logger: %v", err)
	}

	// Get Kubernetes client
	app.KubeRawClient, err = kubernetes.NewKubernetesClient()
	if err != nil {
		app.Logger.Fatal("Error creating Kubernetes client", zap.Error(err))
	}

	// Run the metrics server to expose the metrics
	err = server.RunMetricsServer(&app)
	if err != nil {
		app.Logger.Error("Error starting server", zap.Error(err))
	}

}
