package banner

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"saglyk-backend/internal/appresult"
	"saglyk-backend/internal/handlers"
	"saglyk-backend/pkg/logging"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	getAllURL = ""
	getOneURL = "/:id"
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
	router.GET(getOneURL, h.GetOne)

	router.POST(getAllURL, h.Create)
	router.PUT(getOneURL, h.Update)
	router.DELETE(getOneURL, h.Delete)

}

// GetAll godoc
// @Description all data banner
// @Tags        banner
// @Accept      json
// @Produce     json
// @Security BearerAuth
// @Param       limit    query     int  false  "limit"
// @Param       page     query     int  false  "page"
// @Success     200  {object}  DataBannerDTO
// @Router      /banner [get]
func (h *handler) GetAll(c *gin.Context) {
	limitS, _ := c.GetQuery("limit")
	pageS, _ := c.GetQuery("page")

	limit, err := strconv.Atoi(limitS)
	page, err := strconv.Atoi(pageS)

	result, err := h.repository.GetAllData(context.Background(), limit, page)
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
// @Description one data banner
// @Tags        banner
// @Accept      json
// @Produce     json
// @Security BearerAuth
// @Param id path int true "banner ID"
// @Success     200  {object}  GetBannerDTO
// @Router      /banner/{id} [get]
func (h *handler) GetOne(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("--------------", id)
	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, appresult.ErrNotAcceptable)
		return
	}
	result, err := h.repository.GetOneData(context.Background(), i)
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
// @Description create banner
// @Tags        banner
// @Accept      json
// @Produce     json
// @Security BearerAuth
// @Param       banner  body AddBannerDTO  true  "Create JSON"
// @Success     200  {object}  GetBannerDTO
// @Router      /banner [post]
func (h *handler) Create(c *gin.Context) {
	var (
		err error
	)

	body, errBody := io.ReadAll(c.Request.Body)
	if errBody != nil {
		c.JSON(http.StatusBadRequest, appresult.ErrMissingParam)
		return
	}

	reqDTO := AddBannerDTO{}
	err = json.Unmarshal(body, &reqDTO)
	res, err := h.repository.AddData(context.Background(), reqDTO)

	if err != nil {
		c.JSON(http.StatusInternalServerError, appresult.ErrInternalServer)
		return
	}

	successResult := appresult.Success
	successResult.Data = res
	c.JSON(http.StatusOK, successResult)
	return
}

// Update godoc
// @Description update banner
// @Tags        banner
// @Accept      json
// @Produce     json
// @Security BearerAuth
// @Param id path int true "banner ID"
// @Param       banner  body AddBannerDTO  true  "Create JSON"
// @Success     200  {string}  string
// @Router      /banner/{id} [put]
func (h *handler) Update(c *gin.Context) {
	var (
		res GetBannerDTO
	)
	bannerId := c.Param("id")
	fmt.Println("--------------", bannerId)
	id, err := strconv.Atoi(bannerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, appresult.ErrNotAcceptable)
		return
	}

	body, errBody := io.ReadAll(c.Request.Body)
	if errBody != nil {
		c.JSON(http.StatusBadRequest, appresult.ErrMissingParam)
		return
	}

	reqDTO := AddBannerDTO{}
	err = json.Unmarshal(body, &reqDTO)

	fmt.Println("id :::", id)
	if id != 0 {
		res, err = h.repository.UpdateData(context.Background(), reqDTO, id)
	}

	if err != nil {
		return
	}

	successResult := appresult.Success
	successResult.Data = res
	c.JSON(http.StatusOK, successResult)
	return
}

// Delete godoc
// @Description delete banner
// @Tags        banner
// @Accept      json
// @Produce     json
// @Security BearerAuth
// @Param id path int true "banner ID"
// @Success     200  {string}  string
// @Router      /banner/{id} [delete]
func (h *handler) Delete(c *gin.Context) {

	bannerId := c.Param("id")
	fmt.Println("--------------", bannerId)
	id, err := strconv.Atoi(bannerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, appresult.ErrNotAcceptable)
		return
	}

	err = h.repository.DeleteData(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, appresult.ErrMissingParam)
		return
	}

	successResult := appresult.Success
	successResult.Data = ""
	c.JSON(http.StatusOK, successResult)
	return
}
