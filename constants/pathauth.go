package constants

type PageAuth struct {
	Page string `json:"page"`
	Roles []Role `json:"roles"`
}

var PageAuths = []*PageAuth{
	{ Page: "", Roles: []Role{RoleAdmin, RoleUser} },
}
