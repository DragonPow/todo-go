package router

import (
	"github.com/gin-gonic/gin"

	"project1/domain"
)

type userRoute struct {
	u domain.UserUsecase
}

func BuildUserRoute(router *gin.RouterGroup, userUsecase domain.UserUsecase) {
	user := &userRoute{u: userUsecase}
	router.POST("/login", user.Login)
	router.POST("/signup", user.Signup)
}

func (t *userRoute) Login(c *gin.Context) {

}

func (t *userRoute) Signup(c *gin.Context) {

}
