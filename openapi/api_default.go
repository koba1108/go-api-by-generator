package openapi

import (
	"database/sql"
	"github.com/go-api-by-generator/external"
	"github.com/go-api-by-generator/models"
	"gopkg.in/guregu/null.v3"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// UsersGet -
func UsersGet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	u, _ := models.UserByID(external.GetDB(), uint64(id))
	if u == nil {
		c.JSON(http.StatusOK, []UserResponse{})
		return
	}
	var response []UserResponse
	response = append(response, UserResponse{
		Id:        int32(u.ID),
		Name:      null.NewString(u.Name.String, u.Name.Valid),
		Age:       int32(u.Age),
		Weight:    null.NewInt(u.Weight.Int64, u.Weight.Valid),
		CreatedAt: u.CreatedAt.Format("2006/1/2 15:04:05"),
		UpdatedAt: u.UpdatedAt.Format("2006/1/2 15:04:05"),
	})
	c.JSON(http.StatusOK, response)
}

// UsersUserIdPost -
func UsersUserIdPost(c *gin.Context) {
	user := models.User{
		Name:      sql.NullString{String: "koba", Valid: true},
		Age:       32,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := user.Insert(external.GetDB())
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "OK",
	})
}
