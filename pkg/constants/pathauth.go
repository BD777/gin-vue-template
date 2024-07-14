package constants

type RoleSet map[Role]struct{}

func newRoleSet(roles ...Role) RoleSet {
	resp := make(RoleSet)
	for _, role := range roles {
		resp[role] = struct{}{}
	}
	return resp
}

type PageAuth struct {
	Page  string            `json:"page"`
	Roles map[Role]struct{} `json:"roles"`
}

var PageAuths = []*PageAuth{
	{Page: "home", Roles: newRoleSet(RoleAdmin, RoleUser)},
	{Page: "setting", Roles: newRoleSet(RoleAdmin, RoleUser)},
}
