package middleware

import "testing"

func TestSplitPermission(t *testing.T) {
	cases := map[string]struct {
		resource, action string
	}{
		"companies.list":           {"companies", "list"},
		"special-equipment.create": {"special-equipment", "create"},
		"foo":                      {"*", "foo"},
		"a.b.c":                    {"a.b", "c"},
	}
	for input, want := range cases {
		r, a := splitPermission(input)
		if r != want.resource || a != want.action {
			t.Errorf("splitPermission(%q) = (%q, %q), want (%q, %q)",
				input, r, a, want.resource, want.action)
		}
	}
}
