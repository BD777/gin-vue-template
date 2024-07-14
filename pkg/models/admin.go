package models

import (
	"context"
	"gin-vue-template/pkg/infra/logger"
)

type Admin struct {
	Model
	Username string `gorm:"type:varchar(63);not null" json:"username"`
	Password string `gorm:"type:varchar(63);not null" json:"password"`
	Role     string `gorm:"type:varchar(63);not null" json:"role"`
}

func (g *GormDAO) GetAdmin(ctx context.Context, username, password string) (admin *Admin) {
	admin = &Admin{}
	g.db.Where(Admin{Username: username, Password: password}).First(admin)
	if admin.ID > 0 {
		return admin
	}

	return nil
}

func (g *GormDAO) ListAdmins(ctx context.Context, req *ListRequest) (resp *ListResponse[*Admin]) {
	resp = &ListResponse[*Admin]{
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
	g.db.Select([]string{"id", "username", "role"}).Offset(int(offset)).Limit(int(pagination.PageSize)).Find(&admins)
	resp.List = admins

	if req.QueryCount {
		var cnt int64
		g.db.Model(&Admin{}).Count(&cnt)
		resp.Pagination.Total = cnt
	}

	logger.Info(ctx, "list admins.size = %d", len(admins))

	return
}
