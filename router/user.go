package router

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"project1/domain"
	"project1/util/api_handle"
)

type loginJson struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type IdUri struct {
	ID int32 `uri:"id"`
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

	// Bind data
	var user_info domain.User
	if err := c.ShouldBind(&user_info); err != nil {
		fmt.Println(err)
		api_handle.BadRequesResponse(c, "Some information is wrong")
		return
	}

	// Create new user
	new_user, err := route.u.Create(c.Request.Context(), user_info)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, domain.ErrUserIsExists) {
			api_handle.BadRequesResponse(c, "Username already exists")
		} else {
			api_handle.ServerErrorResponse(c)
		}
		return
	}

	api_handle.SuccessResponse(c, new_user)
}

func (route *userRoute) Update(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic in Update User\n", r)
			api_handle.ServerErrorResponse(c)
		}
	}()

	var id IdUri
	if err := c.ShouldBindUri(&id); err != nil {
		fmt.Println(err)
		api_handle.BadRequesResponse(c, "ID uri must be integer")
		return
	}

	new_user_info := make(map[string]interface{})

	if _, err := c.MultipartForm(); err != nil {
		api_handle.BadRequesResponse(c, "Some information is wrong")
		return
	}

	for key, value := range c.Request.Form {
		new_user_info[key] = value[0]
	}

	if err := route.u.Update(c.Request.Context(), id.ID, new_user_info); err != nil {
		if errors.Is(err, domain.ErrUserNotExists) {
			api_handle.NotFoundResponse(c, "Not found user with id "+strconv.Itoa(int(id.ID)))
		} else {
			api_handle.ServerErrorResponse(c)
		}
		return
	}

	api_handle.SuccessResponse(c, "Update user success, id "+strconv.Itoa(int(id.ID)))
}

func (route *userRoute) Delete(c *gin.Context) {

}
