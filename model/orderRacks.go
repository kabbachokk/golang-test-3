package model

type OrderRack struct {
	PrimaryRack    string
	OrderID        int
	ProductID      int
	ProductName    string
	Qty            int
	SecondaryRacks string
}
