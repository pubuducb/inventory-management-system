package handler

import (
	"database/sql"
	"encoding/json"
	"ims/internal/model"
	"ims/internal/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

const ErrorMessageInvalidUserId = "Invalid user ID format"
const ErrorMessageInvalidRequestBody = "Invalid request body"
const ErrorMessageCreateUserFailed = "Failed to create user"
const ErrorMessageUpdateUserFailed = "Failed to update user"
const ErrorMessageDeleteUserFailed = "Failed to delete user"
const ErrorMessageUserNotFound = "User not found or has been archived"

type UpdateUser struct {
	Name  string `json:"name" binding:"required,min=3,max=100"`
	Email string `json:"email" binding:"required,email"`
}

func (handler *UserHandler) CreateUser(context *gin.Context) {

	var createUser UpdateUser

	decoder := json.NewDecoder(context.Request.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&createUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": ErrorMessageInvalidRequestBody})
		return
	}

	user := model.User{
		Name:  createUser.Name,
		Email: createUser.Email,
	}

	if err := handler.repo.Create(&user); err != nil {
		log.Printf("[UserHandler] %s: %s", ErrorMessageCreateUserFailed, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": ErrorMessageCreateUserFailed})
		return
	}

	context.JSON(http.StatusCreated, user)
}

func (handler *UserHandler) GetUser(context *gin.Context) {

	id, err := getIDParam(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": ErrorMessageInvalidUserId})
		return
	}

	user, err := handler.repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, gin.H{"error": ErrorMessageUserNotFound})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	context.JSON(http.StatusOK, user)
}

func (handler *UserHandler) UpdateUser(context *gin.Context) {

	id, err := getIDParam(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": ErrorMessageInvalidUserId})
		return
	}

	var updateUser UpdateUser

	decoder := json.NewDecoder(context.Request.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&updateUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": ErrorMessageInvalidRequestBody})
		return
	}

	user := model.User{
		ID:   id,
		Name: updateUser.Name,
	}

	err = handler.repo.Update(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, gin.H{"error": ErrorMessageUserNotFound})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": ErrorMessageUpdateUserFailed})
		return
	}

	context.JSON(http.StatusOK, user)
}

func (handler *UserHandler) DeleteUser(context *gin.Context) {

	id, err := getIDParam(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": ErrorMessageInvalidUserId})
		return
	}

	err = handler.repo.Archive(id)
	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, gin.H{"error": ErrorMessageUserNotFound})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": ErrorMessageDeleteUserFailed})
		return
	}

	context.Status(http.StatusNoContent)
}

func getIDParam(context *gin.Context) (int, error) {
	idStr := context.Param("id")
	return strconv.Atoi(idStr)
}
