package category

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"saglyk-backend/internal/appresult"
	"saglyk-backend/internal/config"
	"saglyk-backend/internal/handlers"
	"saglyk-backend/pkg/logging"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	getAllURL = "/get-all"
	getOneURL = "/get"
	addURL    = "/add"
	deleteURL = "/delete"

	// sub category
	getOneProductURL = "/get-product"
	addSubURL        = "/add-product"
	deleteSubURL     = "/delete-product"

	fileUploadURL = "/file"
)

type handler struct {
	repository Repository
	logger     *logging.Logger
}

func NewHandler(repository Repository, logger *logging.Logger) handlers.Handler {
	return &handler{
		repository: repository,
		logger:     logger,
	}
}

func (h *handler) Register(router *gin.RouterGroup) {
	router.GET(getAllURL, h.GetAll)
	router.POST(getOneURL, h.GetOne)
	router.POST(addURL, h.Create)
	router.POST(deleteURL, h.Delete)
	//sub categories
	router.POST(getOneProductURL, h.GetProduct)
	router.POST(addSubURL, h.CreateSub)
	router.POST(deleteSubURL, h.DeleteSub)

	router.POST(fileUploadURL, h.FileUpload)

}

// GetAll godoc
// @Description all data categories
// @Tags        categories
// @Accept      json
// @Produce     json
// @Security BearerAuth
// @Success     200  {array}  CategoryDTO
// @Router      /categories/get-all [get]
func (h *handler) GetAll(c *gin.Context) {
	result, err := h.repository.GetAllData(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, appresult.ErrInternalServer)
		return
	}

	successResult := appresult.Success
	successResult.Data = result
	c.JSON(http.StatusOK, successResult)

	return
}

// GetOne godoc
// @Description one data category
// @Tags        categories
// @Accept      json
// @Produce     json
// @Security BearerAuth
// @Param       categories  body ReqIdDTO true  "Get One JSON"
// @Success     200  {object}  DataDTO
// @Router      /categories/get [post]
func (h *handler) GetOne(c *gin.Context) {
	body, errBody := io.ReadAll(c.Request.Body)
	if errBody != nil {
		c.JSON(http.StatusBadRequest, appresult.ErrMissingParam)
		return
	}

	reqDTO := ReqIdDTO{}
	err := json.Unmarshal(body, &reqDTO)
	result, err := h.repository.GetOneData(context.Background(), reqDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, appresult.ErrInternalServer)
		return
	}

	successResult := appresult.Success
	successResult.Data = result
	c.JSON(http.StatusOK, successResult)

	return
}

// Create godoc
// @Description create and update data  categories
// @Tags        categories
// @Accept      json
// @Produce     json
// @Security BearerAuth
// @Param       categories  body CategoryIdDTO  true  "Create JSON"
// @Success     200  {string}  string
// @Router      /categories/add [post]
func (h *handler) Create(c *gin.Context) {
	var (
		err error
	)

	body, errBody := io.ReadAll(c.Request.Body)
	if errBody != nil {
		c.JSON(http.StatusBadRequest, appresult.ErrMissingParam)
		return
	}

	reqDTO := CategoryIdDTO{}
	err = json.Unmarshal(body, &reqDTO)

	fmt.Println("uuid :::", reqDTO.Id)
	if reqDTO.Id != 0 {
		err = h.repository.UpdateData(context.Background(), reqDTO)
	} else {
		err = h.repository.AddData(context.Background(), reqDTO)
	}

	if err != nil {
		return
	}

	successResult := appresult.Success
	successResult.Data = ""
	c.JSON(http.StatusOK, successResult)
	return
}

// Delete godoc
// @Description delete data  banner
// @Tags        categories
// @Accept      json
// @Produce     json
// @Security BearerAuth
// @Param       categories  body ReqIdDTO  true  "Delete JSON"
// @Success     200  {string}  string
// @Router      /categories/delete [post]
func (h *handler) Delete(c *gin.Context) {

	body, errBody := io.ReadAll(c.Request.Body)
	if errBody != nil {
		c.JSON(http.StatusBadRequest, appresult.ErrMissingParam)
		return
	}

	reqDTO := ReqIdDTO{}
	err := json.Unmarshal(body, &reqDTO)

	err = h.repository.DeleteData(context.Background(), reqDTO.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, appresult.ErrMissingParam)
		return
	}

	successResult := appresult.Success
	successResult.Data = ""
	c.JSON(http.StatusOK, successResult)
	return
}

// ///////// sub-category //////////////////////

// GetProduct godoc
// @Description product data
// @Tags        categories
// @Accept      json
// @Produce     json
// @Security BearerAuth
// @Param       categories  body ReqIdDTO true  "Get One JSON"
// @Success     200  {object}  SubCategoryDTO
// @Router      /categories/get [post]
func (h *handler) GetProduct(c *gin.Context) {
	body, errBody := io.ReadAll(c.Request.Body)
	if errBody != nil {
		c.JSON(http.StatusBadRequest, appresult.ErrMissingParam)
		return
	}

	reqDTO := ReqIdDTO{}
	err := json.Unmarshal(body, &reqDTO)
	result, err := h.repository.GetProductData(context.Background(), reqDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, appresult.ErrInternalServer)
		return
	}

	successResult := appresult.Success
	successResult.Data = result
	c.JSON(http.StatusOK, successResult)

	return
}

// CreateSub godoc
// @Description create and update data  sub-categories
// @Tags        categories
// @Accept      json
// @Produce     json
// @Security BearerAuth
// @Param       categories  body SubCategoryIdDTO  true  "Create JSON"
// @Success     200  {string}  string
// @Router      /categories/add-product [post]
func (h *handler) CreateSub(c *gin.Context) {
	var (
		err error
	)

	body, errBody := io.ReadAll(c.Request.Body)
	if errBody != nil {
		c.JSON(http.StatusBadRequest, appresult.ErrMissingParam)
		return
	}

	reqDTO := SubCategoryIdDTO{}
	err = json.Unmarshal(body, &reqDTO)

	fmt.Println("id :::", reqDTO.Id)
	if reqDTO.Id != 0 {
		err = h.repository.UpdateSubData(context.Background(), reqDTO)
	} else {
		err = h.repository.AddSubData(context.Background(), reqDTO)
	}

	if err != nil {
		return
	}

	successResult := appresult.Success
	successResult.Data = ""
	c.JSON(http.StatusOK, successResult)
	return
}

// DeleteSub godoc
// @Description delete data  sub-categories
// @Tags        categories
// @Accept      json
// @Produce     json
// @Security BearerAuth
// @Param       categories  body ReqIdDTO  true  "Delete JSON"
// @Success     200  {string}  string
// @Router      /categories/delete-product [post]
func (h *handler) DeleteSub(c *gin.Context) {

	body, errBody := io.ReadAll(c.Request.Body)
	if errBody != nil {
		c.JSON(http.StatusBadRequest, appresult.ErrMissingParam)
		return
	}

	reqDTO := ReqIdDTO{}
	err := json.Unmarshal(body, &reqDTO)

	err = h.repository.DeleteSubData(context.Background(), reqDTO.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, appresult.ErrMissingParam)
		return
	}

	successResult := appresult.Success
	successResult.Data = ""
	c.JSON(http.StatusOK, successResult)
	return
}

// FileUpload godoc
// @Description file upload
// @Tags        file
//
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			file	formData	file			true	"this is a test file"
//
// @Security BearerAuth
// @Success     200  {string}  string
// @Router      /categories/file [post]
func (h *handler) FileUpload(c *gin.Context) {
	var (
		err error
	)

	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	cfg := config.GetConfig()
	appDirPath := cfg.PublicFilePath + "/images/"

	err = os.MkdirAll(appDirPath, os.ModePerm)
	if err != nil {
		fmt.Println("wwwww", err)
	}

	appName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(file.Filename))

	pathApp := fmt.Sprintf("%s/%s", appDirPath, appName)
	dbPath := "/public/images/" + appName

	if err := c.SaveUploadedFile(file, pathApp); err != nil {
		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}

	if err != nil {
		fmt.Println("file upload ::: ", err)
		return
	}

	successResult := appresult.Success
	successResult.Data = dbPath
	c.JSON(http.StatusOK, successResult)

	return
}
