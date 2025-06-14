package helper

import (
	"math"
	"net/http"

	"github.com/gofiber/fiber/v2"

	pkgError "github.com/adityaeka26/deptech-test-backend/pkg/error"
)

func init() {}

// Result common output
type Result struct {
	Data     any
	MetaData any
	Error    error
	Count    int64
}

type response struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func getErrorStatusCode(err error) int {
	errString, ok := err.(*pkgError.ErrorString)
	if ok {
		return errString.Code()
	}

	return http.StatusInternalServerError
}

func RespSuccess(c *fiber.Ctx, data any, message string) error {
	return c.Status(http.StatusOK).JSON(response{
		Message: message,
		Data:    data,
		Code:    http.StatusOK,
		Success: true,
	})
}

func RespError(c *fiber.Ctx, err error) error {
	return c.Status(getErrorStatusCode(err)).JSON(response{
		Message: err.Error(),
		Data:    nil,
		Code:    getErrorStatusCode(err),
		Success: false,
	})
}

type MetaData struct {
	Page      int64 `json:"page"`
	Count     int64 `json:"count"`
	TotalPage int64 `json:"total_page"`
	TotalData int64 `json:"total_data"`
}

func GenerateMetaData(totalData, count int64, page, limit int64) MetaData {
	metaData := MetaData{
		Page:      page,
		Count:     count,
		TotalPage: int64(math.Ceil(float64(totalData) / float64(limit))),
		TotalData: totalData,
	}

	return metaData
}

type paginationResponse struct {
	Success bool     `json:"success"`
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Meta    MetaData `json:"meta"`
	Data    any      `json:"data"`
}

func RespPagination(c *fiber.Ctx, data any, metadata MetaData, message string) error {
	return c.Status(http.StatusOK).JSON(paginationResponse{
		Message: message,
		Meta:    metadata,
		Data:    data,
		Code:    http.StatusOK,
		Success: true,
	})
}
