package gee

import (
	"strings"
)

type trieNode struct {
	pattern  string
	part     string
	nextNode []*trieNode
	Wild     bool
}

func (tn *trieNode) matchChild(part string) *trieNode {
	for _, node := range tn.nextNode {
		if node.part == part || node.Wild {
			return node
		}
	}
	return nil
}

func (tn *trieNode) matchChildren(part string) []*trieNode {
	nodes := make([]*trieNode, 0)
	for _, node := range tn.nextNode {
		if node.part == part || node.Wild {
			nodes = append(nodes, node)
		}
	}
	return nodes
}

func (tn *trieNode) insertNode(pattern string, parts []string, height int) {
	if len(parts) == height {
		tn.pattern = pattern
		return
	}
	part := parts[height]

	child := tn.matchChild(part)
	if child == nil {
		child = &trieNode{part: part, Wild: part[0] == ':' || part[0] == '*'}
		tn.nextNode = append(tn.nextNode, child)
	}
	child.insertNode(pattern, parts, height+1)
}

func (tn *trieNode) searchNode(parts []string, height int) *trieNode {
	if len(parts) == height || strings.HasPrefix(tn.part, "*") {
		if tn.pattern == "" {
			return nil
		}
		return tn
	}
	part := parts[height]
	children := tn.matchChildren(part)

	for _, child := range children {
		result := child.searchNode(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
