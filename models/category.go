package models

type Category struct {
	Cid  int
	Name string
}
type CategoryResponse struct {
	*HomeResponse
	CategoryName string
}
