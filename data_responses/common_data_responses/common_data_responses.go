package common_data_responses

type PageInfo struct {
	TotalCount     int64 `json:"total_count"`
	RowsPerPage    int   `json:"rows_per_page"`
	CurrentPage    int   `json:"current_page"`
	TotalPageCount int64 `json:"total_page_count"`
	HasNextPage    bool  `json:"has_next_page"`
}
