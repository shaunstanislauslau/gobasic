// token_test.go: Simple tests of our keywords

package token

import (
	"strings"
	"testing"
)

// Test looking up values succeeds, with both lower + upper-case
func TestKeywordLookup(t *testing.T) {

	for key, val := range keywords {

		// Obviously this will pass.
		if LookupIdentifier(string(key)) != val {
			t.Errorf("Lookup of %s failed", key)
		}

		// Once the keywords are uppercase they'll no longer
		// match - so we find them as identifiers.
		if LookupIdentifier(strings.ToUpper(string(key))) == IDENT {
			t.Errorf("Lookup of %s failed", key)
		}
	}
}

// Test looking up variables returns "idents", not keywords.
func TestIdentLookup(t *testing.T) {

	names := []string{"a", "a$", "name", "name$"}

	for _, ent := range names {

		// Obviously this will pass.
		if LookupIdentifier(string(ent)) != IDENT {
			t.Errorf("Failed to identify %s as an ident", ent)
		}
	}
}
