package atm

type RowsMap = map[string]interface{}

type ResultWithPage[T any] struct {
	Total int64 `json:"total"`
	Data  []T   `json:"data"`
}
