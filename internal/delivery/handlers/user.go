package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"vlad_xp/internal/models"
	"vlad_xp/internal/service"
)

type UserHandler struct {
	service service.User
}

func InitUserHandler(userService service.User) UserHandler {
	return UserHandler{
		service: userService,
	}
}

// @Summary Create user
// @Tags user
// @Accept  json
// @Produce  json
// @Param data body models.CreateUser true "user create"
// @Success 200 {object} int "Successfully created user"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/create [post]
func (h UserHandler) Create(g *gin.Context) {
	var creation models.CreateUser

	if err := g.ShouldBindJSON(&creation); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	id, err := h.service.Create(ctx, creation)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary Get user
// @Tags user
// @Accept  json
// @Produce  json
// @Param id query int true "UserID"
// @Success 200 {object} int "Successfully get user"
// @Failure 400 {object} map[string]string "Invalid UserID"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/get/{id} [get]
func (h UserHandler) GetUser(g *gin.Context) {
	userID := g.Query("id")

	id, err := strconv.Atoi(userID)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	user, err := h.service.GetMe(ctx, id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"user": &user})
}

// @Summary Delete user
// @Tags user
// @Accept  json
// @Produce  json
// @Param id query int true "UserID"
// @Success 200 {object} int "Successfully deleted"
// @Failure 400 {object} map[string]string "Invalid id"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/delete/{id} [delete]
func (h UserHandler) Delete(g *gin.Context) {
	userID := g.Query("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = h.service.Delete(ctx, id)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"delete": id})
}

// @Summary Login user
// @Tags user
// @Accept  json
// @Produce  json
// @Param data body models.CreateUser true "user login"
// @Success 200 {object} int "Successfully login user"
// @Failure 400 {object} map[string]string "Invalid password"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/login [post]
func (h UserHandler) Login(g *gin.Context) {
	var User models.CreateUser

	if err := g.ShouldBindJSON(&User); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id, err := h.service.Login(ctx, User)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": id})
}
