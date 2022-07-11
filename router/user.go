package router

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"project1/domain"
	"project1/util/api_handle"
)

type loginJson struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type userRoute struct {
	u domain.UserUsecase
}

func BuildUserRoute(router *gin.RouterGroup, userUsecase domain.UserUsecase) {
	user := &userRoute{u: userUsecase}
	router.GET("/login", user.Login)
	router.GET("/:id", user.GetByID)
	router.POST("/", user.Create)
	router.PUT("/:id", user.Update)
	router.DELETE("/:id", user.Delete)
}

func (route *userRoute) Login(c *gin.Context) {
	var loginInfo loginJson
	if err := c.ShouldBindQuery(&loginInfo); err != nil {
		api_handle.BadRequesResponse(c, "Username or password is needed")
		return
	}

	account, err := route.u.Login(c.Request.Context(), loginInfo.Username, loginInfo.Password)
	if err != nil {
		message := fmt.Sprintf("User with username = %s and password = %s is wrong", loginInfo.Username, loginInfo.Password)
		api_handle.NotFoundResponse(c, message)
		return
	}

	api_handle.SuccessResponse(c,
		map[string]interface{}{
			"id":       account.ID,
			"username": account.Username,
			"name":     account.Name,
		})
}

func (route *userRoute) GetByID(c *gin.Context) {

}

func (route *userRoute) Create(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic in Create User\n", r)
			api_handle.ServerErrorResponse(c)
		}
	}()

	var user_info domain.User
	if err := c.ShouldBind(&user_info); err != nil {
		fmt.Println(err)
		api_handle.BadRequesResponse(c, "Some information is wrong")
		return
	}

	new_user, err := route.u.Create(c.Request.Context(), user_info)

	if err != nil {
		fmt.Println(err)
		domainError, ok := err.(*domain.DomainError)

		if ok {
			if domainError.Code == domain.UsernameIsExists {
				api_handle.BadRequesResponse(c, "Username already exists")
				return
			}
		}

		api_handle.ServerErrorResponse(c)
		return
	}

	api_handle.SuccessResponse(c, new_user)
}

func (route *userRoute) Update(c *gin.Context) {

}

func (route *userRoute) Delete(c *gin.Context) {

}
