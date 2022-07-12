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
	defer handlePanic(c, "Login")

	// Get username and password from query
	var loginInfo loginJson
	if err := c.ShouldBindQuery(&loginInfo); err != nil {
		api_handle.BadRequesResponse(c, "Username or password is needed")
		return
	}

	// Login
	account, err := route.u.Login(c.Request.Context(), loginInfo.Username, loginInfo.Password)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotExists) {
			message := fmt.Sprintf("Username or password is wrong")
			api_handle.NotFoundResponse(c, message)
		} else {
			api_handle.ServerErrorResponse(c)
		}
		return
	}

	// Build token
	// token := make(map[string]interface{})
	// token["token"] = "123"

	api_handle.SuccessResponse(c, getUserJsonResponse(account))
}

func (route *userRoute) GetByID(c *gin.Context) {
	defer handlePanic(c, "Get user by ID")

	var id IdUri
	// Get Id from uri
	if err := c.ShouldBindUri(&id); err != nil {
		api_handle.BadRequesResponse(c, "Id must be integer")
		return
	}

	// Get
	user_info, err := route.u.GetByID(c.Request.Context(), id.ID)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotExists) {
			api_handle.NotFoundResponse(c, "User with ID "+strconv.Itoa(int(id.ID))+" does not exist")
		} else {
			api_handle.ServerErrorResponse(c)
		}
		return
	}

	api_handle.SuccessResponse(c, getUserJsonResponse(user_info))
}

func (route *userRoute) Create(c *gin.Context) {
	defer handlePanic(c, "Create user")

	// Get user information from body
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

	api_handle.SuccessResponse(c, getUserJsonResponse(new_user))
}

func (route *userRoute) Update(c *gin.Context) {
	defer handlePanic(c, "Update user")

	// Get Id from uri
	var id IdUri
	if err := c.ShouldBindUri(&id); err != nil {
		fmt.Println(err)
		api_handle.BadRequesResponse(c, "ID uri must be integer")
		return
	}

	// Get user information from multipart form
	new_user_info := make(map[string]interface{})
	if _, err := c.MultipartForm(); err != nil {
		api_handle.BadRequesResponse(c, "Some information is wrong")
		return
	}
	for key, value := range c.Request.Form {
		new_user_info[key] = value[0]
	}

	// Update
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
	defer handlePanic(c, "Delete user")

	var id IdUri
	// Get Id from uri
	if err := c.ShouldBindUri(&id); err != nil {
		api_handle.BadRequesResponse(c, "ID uri must be integer")
		return
	}

	// Delete
	if err := route.u.Delete(c.Request.Context(), id.ID); err != nil {
		if errors.Is(err, domain.ErrUserNotExists) {
			api_handle.NotFoundResponse(c, "User with id "+strconv.Itoa(int(id.ID))+" not exists")
		} else {
			api_handle.ServerErrorResponse(c)
		}
		return
	}

	api_handle.SuccessResponse(c, "Delete id "+strconv.Itoa(int(id.ID))+" success")
}
