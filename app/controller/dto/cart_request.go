package dto

type AccessCartRequest struct {
	ID string `json:"id"` // cart id
}

type UpdProductRequest struct {
	ID        string `json:"id"`      // cart id
	Product   string `json:"product"` // product id
	Size      string `json:"size"`
	Delta     int    `json:"delta"`
	Remaining int    `json:"remaining"`
}

type RemoveItemRequest struct {
	ID      string `json:"id"`
	Product string `json:"product"`
}
