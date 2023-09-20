package models

import "context"

type Admin struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func GetAdmin(ctx context.Context, username, password string) (admin *Admin) {
	admin = &Admin{}
	db.Where(Admin{Username: username, Password: password}).First(admin)
	if admin.ID > 0 {
		return admin
	}

	return nil
}

func ListAdmins(ctx context.Context, req *ListRequest) (resp *ListResponse) {
	resp = &ListResponse{
		Pagination: &Pagination{
			Page:     req.Pagination.Page,
			PageSize: req.Pagination.PageSize,
		},
	}

	pagination := req.Pagination
	page := pagination.Page
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pagination.PageSize
	var admins []*Admin
	db.Select([]string{"id", "username", "role"}).Offset(offset).Limit(pagination.PageSize).Find(&admins)
	for _, item := range admins {
		resp.List = append(resp.List, item)
	}

	if req.QueryCount {
		var cnt int
		db.Model(&Admin{}).Count(&cnt)
		resp.Pagination.Total = cnt
	}

	return
}
