package http

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/t0239184/CleanArch/app"
	"github.com/t0239184/CleanArch/app/domain"
)

type UserHandler struct {
	UserUsecase domain.IUserUsecase
}

type CreateUserRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Id       int64  `json:"id,string"`
	Password string `json:"password"`
}

func NewUserHandler(e *gin.Engine, userUsecase domain.IUserUsecase) {
	handler := &UserHandler{
		UserUsecase: userUsecase,
	}
	e.GET("/api/v1/user/:id", handler.FindById)
	e.POST("/api/v1/user", handler.CreateUser)
	e.POST("/api/v1/user/:id", handler.UpdateUser)
	e.POST("/api/v1/user/:id/delete", handler.DeleteUser)
	e.POST("/api/v1/user/:id/unlock", handler.UnlockUser)
}

func (u *UserHandler) FindById(c *gin.Context) {
	request_id, _ := c.Get("request_id")
	fmt.Println(request_id)
	userId := c.Param("id")
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Invalid user id",
		})
		return
	}

	user, err := u.UserUsecase.FindById(&id)
	if err != nil {
		fmt.Println(err)
		if err == gorm.ErrRecordNotFound {
			e := app.ErrUserNotFound
			c.JSON(200, app.ErrorResponse(e))
			return
		}
		c.JSON(200, gin.H{
			"status":  500,
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  0,
		"message": "success",
		"data":    user,
	})
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	request := &CreateUserRequest{}
	if err := c.Bind(request); err != nil {
		fmt.Println(err)
		return
	}
	user := &domain.User{
		Account:  request.Account,
		Password: request.Password,
		Status:   "0",
	}

	id, err := u.UserUsecase.CreateUser(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"status":  500,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  0,
		"message": "success",
		"data": gin.H{
			"id": id,
		},
	})
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	request := &UpdateUserRequest{}
	if err := c.Bind(request); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  400,
			"message": "Invalid user id",
		})
		return
	}

	user := &domain.User{
		Id:       request.Id,
		Password: request.Password,
	}

	if err := u.UserUsecase.UpdateUser(user); err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"status":  500,
			"message": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"status":  0,
		"message": "success",
	})
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status":  400,
			"message": "Invalid user id",
		})
		return
	}

	if err := u.UserUsecase.DeleteUser(&id); err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"status":  500,
			"message": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"status":  0,
		"message": "success",
	})
}

func (u *UserHandler) UnlockUser(c *gin.Context) {
	userId := c.Param("id")
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status":  400,
			"message": "Invalid user id",
		})
		return
	}

	if err := u.UserUsecase.UnlockUser(&id); err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"status":  500,
			"message": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"status":  0,
		"message": "success",
	})
}
