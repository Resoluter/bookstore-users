package users

import (
	"fmt"
	"net/http"

	"github.com/Resoluter/bookstore-users.git/domain/users"
	"github.com/Resoluter/bookstore-users.git/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	// fmt.Println(user)
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err.Error())
	// }

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
