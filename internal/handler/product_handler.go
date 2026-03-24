package handler

import (
	"database/sql"
	"encoding/json"
	"ims/internal/model"
	"ims/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	repo repository.ProductRepository
}

func NewProductHandler(repo repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

const ErrorMessageInvalidProductId = "Invalid product ID format"
const ErrorMessageInvalidRequestBody = "Invalid request body"
const ErrorMessageCreateProductFailed = "Failed to create product"
const ErrorMessageUpdateProductFailed = "Failed to update product"
const ErrorMessageDeleteProductFailed = "Failed to delete product"
const ErrorMessageProductNotFound = "Product not found or has been archived"

type UpdateProduct struct {
	Name  string  `json:"name" binding:"required,min=3,max=100"`
	Price float32 `json:"price" binding:"required,gt=0"`
}

func (handler *ProductHandler) CreateProduct(context *gin.Context) {

	var createProduct UpdateProduct

	decoder := json.NewDecoder(context.Request.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&createProduct); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": ErrorMessageInvalidRequestBody})
		return
	}

	product := model.Product{
		Name:  createProduct.Name,
		Price: createProduct.Price,
	}

	if err := handler.repo.Create(&product); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": ErrorMessageCreateProductFailed})
		return
	}

	context.JSON(http.StatusCreated, product)
}

func (handler *ProductHandler) GetProduct(context *gin.Context) {

	id, err := getIDParam(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": ErrorMessageInvalidProductId})
		return
	}

	product, err := handler.repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, gin.H{"error": ErrorMessageProductNotFound})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	context.JSON(http.StatusOK, product)
}

func (handler *ProductHandler) UpdateProduct(context *gin.Context) {

	id, err := getIDParam(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": ErrorMessageInvalidProductId})
		return
	}

	var updateProduct UpdateProduct

	decoder := json.NewDecoder(context.Request.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&updateProduct); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": ErrorMessageInvalidRequestBody})
		return
	}

	product := model.Product{
		ID:    id,
		Name:  updateProduct.Name,
		Price: updateProduct.Price,
	}

	err = handler.repo.Update(&product)
	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, gin.H{"error": ErrorMessageProductNotFound})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": ErrorMessageUpdateProductFailed})
		return
	}

	context.JSON(http.StatusOK, product)
}

func (handler *ProductHandler) DeleteProduct(context *gin.Context) {

	id, err := getIDParam(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": ErrorMessageInvalidProductId})
		return
	}

	err = handler.repo.Archive(id)
	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, gin.H{"error": ErrorMessageProductNotFound})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": ErrorMessageDeleteProductFailed})
		return
	}

	context.Status(http.StatusNoContent)
}

func getIDParam(context *gin.Context) (int, error) {
	idStr := context.Param("id")
	return strconv.Atoi(idStr)
}
