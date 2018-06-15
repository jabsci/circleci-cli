package filetree

import (
	"fmt"
	"os"
	"path/filepath"
)

// Node represents a leaf in the filetree
type Node struct {
	FullPath string
	Info     os.FileInfo
	Children []*Node
	Parent   *Node
}

// Helper function that returns true if a path exists in excludes array
func excluded(exclude []string, path string) bool {
	for _, n := range exclude {
		if path == n {
			return true
		}
	}
	return false
}

// NewTree creates a new filetree starting at the root
func NewTree(root string) (*Node, error) {
	parents := make(map[string]*Node)
	var result *Node

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// check if file is in exclude slice and skip it
		// need to pass this in as an array
		exclude := []string{"path/to/skip"}
		if excluded(exclude, path) {
			fmt.Printf("skipping: %+v \n", info.Name())
			return filepath.SkipDir
		}

		parents[path] = &Node{
			FullPath: path,
			Info:     info,
			Children: make([]*Node, 0),
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	for path, node := range parents {
		parentPath := filepath.Dir(path)
		parent, exists := parents[parentPath]
		if exists {
			node.Parent = parent
			parent.Children = append(parent.Children, node)
		} else {
			result = node
		}

	}
	return result, err
}