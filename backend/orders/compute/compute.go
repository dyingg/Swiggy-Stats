package orders

import (
	. "sgstory/orders"
	. "sgstory/orders/stream"
	. "sgstory/util"
	"strconv"
)

type ComputedStats struct {
	IncludedOrders        int                           `json:"included_orders"`
	OnTimeOrders          int                           `json:"on_time_orders"`
	OrderTotalWithTip     float64                       `json:"order_total_with_tip"`
	NetTotal              float64                       `json:"net_total"`
	ItemTotal             float64                       `json:"item_total"`
	OrderDiscount         float64                       `json:"order_discount"`
	TotalTimeWaited       int64                         `json:"total_time_waited"`
	DeliveryBoys          FrequencyCounter[DeliveryBoy] `json:"delivery_boys"`
	Restaurants           FrequencyCounter[Restaurant]  `json:"restaurants"`
	Cuisines              FrequencyCounter[string]      `json:"cuisines"`
	Items                 FrequencyCounter[string]      `json:"items"`
	RatingMapRestaurants  [5]int                        `json:"rating_map_restaurants"`
	RatingMapDeliveryBoys [5]int                        `json:"rating_map_delivery_boys"`
}

type UIShowStats struct {
	IncludedOrders        int                          `json:"included_orders"`
	OnTimeOrders          int                          `json:"on_time_orders"`
	OrderTotalWithTip     float64                      `json:"order_total_with_tip"`
	NetTotal              float64                      `json:"net_total"`
	ItemTotal             float64                      `json:"item_total"`
	OrderDiscount         float64                      `json:"order_discount"`
	TotalTimeWaited       int64                        `json:"total_time_waited"`
	DeliveryBoys          []FrequencyPair[DeliveryBoy] `json:"delivery_boys"`
	Restaurants           []FrequencyPair[Restaurant]  `json:"restaurants"`
	Cuisines              []FrequencyPair[string]      `json:"cuisines"`
	Items                 []FrequencyPair[string]      `json:"items"`
	RatingMapRestaurants  [5]int                       `json:"rating_map_restaurants"`
	RatingMapDeliveryBoys [5]int                       `json:"rating_map_delivery_boys"`
	Signature             string                       `json:"signature"`
}

func (stats ComputedStats) ToUIStats() UIShowStats {
	return UIShowStats{
		IncludedOrders:        stats.IncludedOrders,
		OnTimeOrders:          stats.OnTimeOrders,
		OrderTotalWithTip:     stats.OrderTotalWithTip,
		NetTotal:              stats.NetTotal,
		ItemTotal:             stats.ItemTotal,
		OrderDiscount:         stats.OrderDiscount,
		TotalTimeWaited:       stats.TotalTimeWaited,
		DeliveryBoys:          stats.DeliveryBoys.SortK(3),
		Restaurants:           stats.Restaurants.SortK(3),
		Cuisines:              stats.Cuisines.SortK(3),
		Items:                 stats.Items.SortK(5),
		RatingMapRestaurants:  stats.RatingMapRestaurants,
		RatingMapDeliveryBoys: stats.RatingMapDeliveryBoys,
		Signature:             "A product by Dying ;)",
	}
}

func ComputeStats(authToken string) (ComputedStats, error) {
	OrderStream := CreateOrderStream(authToken)

	var UserStats ComputedStats

	for OrderStream.HasNext {
		orders, err := OrderStream.Next()
		if err != nil {
			return UserStats, err
		}

		for _, order := range orders {
			//Do something with the order
			// fmt.Println(order)

			UserStats.IncludedOrders++
			UserStats.OrderTotalWithTip += order.OrderTotalWithTip
			UserStats.NetTotal += order.NetTotal
			UserStats.ItemTotal += order.ItemTotal
			UserStats.OrderDiscount += order.OrderDiscount

			timeTaken, err := strconv.ParseInt(order.DeliveryTimeInSeconds, 10, 64)

			if err == nil {
				UserStats.TotalTimeWaited += timeTaken
			}

			//1.) Frequency Counters

			for _, item := range order.OrderItems {

				itemQuantity, err := strconv.ParseInt(item.Quantity, 10, 64)

				if err == nil {
					UserStats.Items.Add(item.Name, item.Name, itemQuantity)
				}
			}

			//Index the restaurants
			UserStats.Restaurants.Add(order.RestaurantId, RestaurantFromOrder(order), 1)
			//Index the delivery boy
			if order.OrderDeliveryBoy.Name != "" {
				UserStats.DeliveryBoys.Add(order.OrderDeliveryBoy.Name, order.OrderDeliveryBoy, 1)
			}

			//Index the cuisine
			for _, cuisine := range order.RestaurantCusine {
				UserStats.Cuisines.Add(cuisine, cuisine, 1)
			}

			//Compute delivery ratings
			if order.Rating.DeliveryRating.Rating != 0 {
				UserStats.RatingMapDeliveryBoys[order.Rating.DeliveryRating.Rating-1]++
			}

			//2.) Ratings mappings

			//Compute Restaurant ratings
			if order.Rating.RestaurantRating.Rating != 0 {
				UserStats.RatingMapRestaurants[order.Rating.RestaurantRating.Rating-1]++
			}

			if order.OnTime {
				UserStats.OnTimeOrders++
			}

		}
	}

	return UserStats, nil

}
