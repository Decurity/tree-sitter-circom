package tree_sitter_circom_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-circom"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_circom.Language())
	if language == nil {
		t.Errorf("Error loading Circom grammar")
	}
}
