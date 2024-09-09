package tree_sitter_circom_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_circom "github.com/tree-sitter/tree-sitter-circom/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_circom.Language())
	if language == nil {
		t.Errorf("Error loading Circom grammar")
	}
}
