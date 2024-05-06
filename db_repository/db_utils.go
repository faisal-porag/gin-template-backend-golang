package db_repository

import (
	"errors"
	"github.com/faisal-porag/gin-template-backend-golang/data_responses/common_data_responses"
)

var NoDataFound = errors.New("no data found")
var NoDataFoundById = errors.New("no data found by this id")
var RowsPerPageZeroErrors = errors.New("rows per page value must be grater than zero")

func GetPaginationInfoData(totalCount int64, CurrentPage int, limitPerPage int) (*common_data_responses.PageInfo, error) {
	// Calculate the total number of pages
	if limitPerPage == 0 {
		return nil, RowsPerPageZeroErrors
	}
	totalPages := totalCount / int64(limitPerPage)
	if totalCount%int64(limitPerPage) > 0 {
		totalPages++
	}

	hasNextPage := int64(CurrentPage) < totalPages
	pagingData := &common_data_responses.PageInfo{
		TotalCount:     totalCount,
		RowsPerPage:    limitPerPage,
		CurrentPage:    CurrentPage,
		TotalPageCount: totalPages,
		HasNextPage:    hasNextPage,
	}

	return pagingData, nil
}
