package service

import (
	"animation/entity"
	"animation/entity/link"
	"animation/luff"
	"animation/luff/middleware/jwt"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

//用户注册，返回数据
func Register(ctx *gin.Context, user entity.User) (p *entity.User,
	err error) {
	var u entity.User
	orm := link.Orm

	t := orm.Where("id>0")
	t.Where("tel=?", user.Tel)
	t.Get(&u)
	if u.Id > 0 {
		err = errors.New("该账户已存在")
		return
	}
	user.Password = luff.Md5encode(user.Password)
	user.Id, err = orm.InsertOne(user)

	var j jwt.JWT
	var UserClaims jwt.CustomClaims
	UserClaims.Id = u.Id
	UserClaims.Tel = u.Tel
	UserClaims.NickName = u.NickName

	tokenString, err := j.CreateToken(UserClaims)
	if err != nil {
		fmt.Printf("create dell err: %s", err)
	}
	fmt.Println(tokenString)

}
