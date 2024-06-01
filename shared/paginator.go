package shared

import (
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"main.go/models"
)

type PagedUser struct {
	Data     []models.UserModel             `json:"data"`
	PageInfo mongopagination.PaginationData `json:"PageInfo"`
}
