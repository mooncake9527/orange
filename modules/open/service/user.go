package service

import (
	"github.com/mooncake9527/orange-core/core/base"
	"github.com/mooncake9527/orange/modules/open/models"
	"github.com/mooncake9527/x/xerrors/xerror"
)

type UserService struct {
	*base.BaseService
}

var SerUser = UserService{
	base.NewService("default"),
}

func (x *UserService) Create(data *models.User) error {
	err := x.DB().Create(data).Error
	if err != nil {
		return xerror.New(err.Error())
	}
	return nil
}

func (x *UserService) GetAll(list *[]*models.User) (err error) {
	err = x.DB().Model(&models.User{}).Find(list).Error
	if err != nil {
		return xerror.Wrap(err, "query User all error")
	}
	return nil
}

func (x *UserService) Get(req, u *models.User) (err error) {
	q := x.DB()
	if req.Id != 0 {
		q = q.Where("id=?", req.Id)
	}
	err = q.First(&u).Error
	if err != nil {
		return xerror.Wrap(err, "query User error")
	}
	return nil
}
