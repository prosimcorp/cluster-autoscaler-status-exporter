package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

// GetNodeGroupsNames return an array with the names of the node-groups in the same order they are in the status
func GetNodeGroupsNames(status string) []string {
	var nodeGroupNames []string
	nameRe := regexp.MustCompile(`(Name:\s*)([a-zA-Z0-9_-]+)`)
	nodeGroupMatches := nameRe.FindAllStringSubmatch(status, -1)

	for _, match := range nodeGroupMatches {
		nodeGroupNames = append(nodeGroupNames, match[2])
	}

	return nodeGroupNames
}

// GetNodeGroupsHealthArguments return an array where each element is a string with all the arguments of one nodegroup
func GetNodeGroupsHealthArguments(status string) []string {
	// Look for the node group health arguments
	nameRe := regexp.MustCompile(`(Health:\s*)([a-zA-Z0-9]+)\s*\((?P<args>(.*)+)\)`)
	nodeGroupMatches := nameRe.FindAllStringSubmatch(status, -1)

	// Filter arguments string
	var healthArgs []string
	symbolsRe := regexp.MustCompile(`[^\w=]`)

	for _, match := range nodeGroupMatches {
		// Delete all the symbols from the match
		match[3] = symbolsRe.ReplaceAllString(match[3], " ")

		// Separate string into an array of arguments
		if !strings.Contains(match[3], "minSize") || !strings.Contains(match[3], "maxSize") {
			continue
		}

		healthArgs = append(healthArgs, match[3])
	}

	return healthArgs
}

// GetNodeGroupsObject return a NodeGroups type with all the data from status ConfigMap already parsed
func GetNodeGroupsObject(nodeGroupsNames []string, nodeGroupsHealthStatus []string) (*NodeGroups, error) {
	var nodeGroup NodeGroup
	var nodeGroups NodeGroups

	stringRe := regexp.MustCompile(`(\w+)=([0-9]+)`)
	replacePattern := `"$1":"$2",`

	for i, name := range nodeGroupsNames {
		// Include NG name
		nodeGroup.Name = name

		// Parse NG args
		arguments := stringRe.ReplaceAllString(nodeGroupsHealthStatus[i], replacePattern)
		arguments = strings.TrimSpace(arguments)       // Fix extra spaces
		arguments = strings.TrimSuffix(arguments, ",") // Fix trailing colon
		arguments = fmt.Sprintf("{%s}", arguments)     // Complete json syntax

		// Parse string args into NG Object
		err := json.Unmarshal([]byte(arguments), &nodeGroup.Health)
		if err != nil {
			return &nodeGroups, err
		}

		// Append NG to the NG pool
		nodeGroups = append(nodeGroups, nodeGroup)
	}

	return &nodeGroups, nil
}
