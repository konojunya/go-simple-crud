package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/konojunya/go-simple-crud/service"
)

// GET /users
func FindUser(c *gin.Context) {
	users, err := service.FindUserAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, users)
}

// GET /users/konojunya
func FindUserByName(c *gin.Context) {
	name := c.Param("name")
	ok, err := service.ExistsUserByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	if !ok {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	user, err := service.FindUserByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, user)
}

// POST /users
func CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	ok, err := service.ExistsUserByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"message": "そのユーザー名おるよー",
		})
		return
	}

	err = service.CreateUser(name, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

// PUT /users/konojunya
func EditUser(c *gin.Context) {
	name := c.Param("name")
	password := c.PostForm("password")
	ok, err := service.ExistsUserByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	if !ok {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	user, err := service.FindUserByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	user, err = service.UpdatePasswordById(user.ID, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}

// DELETE /users/konojunya
func DeleteUser(c *gin.Context) {
	name := c.Param("name")
	ok, err := service.ExistsUserByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	if !ok {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	user, err := service.FindUserByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	err = service.DeleteUserById(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
