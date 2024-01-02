package services

import (
	"fmt"
	"strings"
)

type nodeType int

const (
	nodeTypeUser nodeType = iota
	nodeTypeRepository
	nodeTypeIssue
	nodeTypePullRequest
	nodeTypeProject
	nodeTypeProjectItem
)

func getTypeOfNode(id string) (nodeType, error) {
	split := strings.Split(id, "_")
	if len(split) < 2 {
		return nodeType(-1), fmt.Errorf("invalid id %s", id)
	}

	switch split[0] {
	case "ISSUE":
		return nodeTypeIssue, nil
	case "PJ":
		return nodeTypeProject, nil
	case "PVTI":
		return nodeTypeProjectItem, nil
	case "PR":
		return nodeTypePullRequest, nil
	case "REPO":
		return nodeTypeRepository, nil
	case "U":
		return nodeTypeUser, nil
	}

	return nodeType(-1), fmt.Errorf("undefined id type %s", split[0])
}
