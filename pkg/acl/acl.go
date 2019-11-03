package acl

import "fmt"

type UserPermission struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	Groups   []string `json:"groups"`
}

// UserRoleCategory represents the top-level of the permission role system
type UserRoleCategory struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Roles       []*UserRole `json:"roles"`
}

// UserRole represents a single permission role.
type UserRole struct {
	Name        string              `json:"name"`
	Description string              `json:"description"`
	APIEndpoint []*UserRoleEndpoint `json:"api_endpoints"`
}

// UserRoleEndpoint represents the path and method of the API endpoint to be secured.
type UserRoleEndpoint struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

type Acl struct {
	RoleCategories []*UserRoleCategory
}

func (a *Acl) CheckRole(userRoles interface{}, method, path string) error {
	perm := a.getRequestRole(method, path)
	if perm == "" {
		return nil
	}
	for _, role := range userRoles.([]interface{}) {
		if role.(string) == perm {
			return nil
		}
	}
	return fmt.Errorf("required permission role %s", perm)
}

func (a *Acl) fullUserRoleName(category *UserRoleCategory, role *UserRole) string {
	return category.Name + role.Name
}

func (a *Acl) getRequestRole(method, path string) string {
	for _, category := range a.RoleCategories {
		for _, role := range category.Roles {
			for _, endpoint := range role.APIEndpoint {
				// If the http method & path match then return the role required for this endpoint
				if method == endpoint.Method && path == endpoint.Path {
					return a.fullUserRoleName(category, role)
				}
			}
		}
	}
	return ""
}