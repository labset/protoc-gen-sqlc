package helpers

import "strings"

func ExtractDomain(packageName string) string {
	parts := strings.Split(packageName, ".")
	if len(parts) < 2 {
		return ""
	}

	// Support patterns:
	// domain.version (e.g., todo.v1)
	// domain.subdomain.version (e.g., user.auth.v1)

	return parts[0]
}
