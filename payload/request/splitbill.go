package request

type SplitBillRequest struct {
	To          []RequestTo
	Item        []Items
	Description string
	Nominal     int64
}

type RequestTo struct {
	CustomerID int
}

type Items struct {
	ProductID int
	Quantity  int
}
