package orders

import (
	"fmt"
	. "sgstory/orders"
	"sgstory/requests"
)

type OrderStream struct {
	HasNext        bool
	LastOrderID    string
	AuthToken      string
	TotalPages     int
	CurrentPage    int
	TotalReadCount int
	TotalOrders    int
}

func CreateOrderStream(authToken string) OrderStream {
	return OrderStream{
		HasNext:        true,
		LastOrderID:    "",
		CurrentPage:    -1,
		TotalReadCount: 0,
		AuthToken:      authToken,
	}
}

func (os *OrderStream) Next() ([]Order, error) {
	if os.CurrentPage != 0 {
		//Check if we have consumed all the pages
		if os.CurrentPage >= os.TotalPages {
			os.HasNext = false
			return []Order{}, nil
		}
	}

	//Fetch the orders
	orderResponse, err := requests.FetchOrders(os.AuthToken, os.LastOrderID)

	if err != nil {
		return []Order{}, err
	}

	//We have fetched a page
	os.CurrentPage = os.CurrentPage + 1

	if os.CurrentPage == 0 {
		//Set the total pages only sent on the first page
		os.TotalOrders = orderResponse.Data.TotalOrders
		fmt.Printf("Total orders: %d\n", os.TotalOrders)
		totalPages := os.TotalOrders / 10
		os.TotalPages = totalPages
	}

	items := len(orderResponse.Data.Orders)

	if os.CurrentPage >= os.TotalPages || items == 0 {
		os.HasNext = false
	}

	if items > 0 {
		os.TotalReadCount = os.TotalReadCount + items
		os.LastOrderID = fmt.Sprintf("%.0f", orderResponse.Data.Orders[items-1].OrderID)
	}

	return orderResponse.Data.Orders, nil

}
