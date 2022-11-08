package orders

type Order struct {
	OrderID                float64      `json:"order_id"`
	OrderDeliveryBoy       DeliveryBoy  `json:"delivery_boy"`
	RestaurantId           string       `json:"restaurant_id"`
	RestaurantName         string       `json:"restaurant_name"`
	RestaurantLocality     string       `json:"restaurant_locality"`
	RestaurantCusine       []string     `json:"restaurant_cuisine"`
	OrderTotalWithTip      float64      `json:"order_total_with_tip"`
	NetTotal               float64      `json:"net_total"`
	ItemTotal              float64      `json:"item_total"`
	OrderDiscount          float64      `json:"order_discount"`
	CouponCode             string       `json:"coupon_code"`
	OnTime                 bool         `json:"on_time"`
	DeliveredTimeInSeconds string       `json:"delivered_time_in_seconds"`
	DeliveryTimeInSeconds  string       `json:"delivery_time_in_seconds"`
	Rating                 RatingMeta   `json:"rating_meta"`
	RestaurantCoverImage   string       `json:"restaurant_cover_image"`
	OrderItems             []OrderItems `json:"order_items"`
}

type OrderItems struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
}

type Restaurant struct {
	RestaurantId         string   `json:"restaurant_id"`
	RestaurantCoverImage string   `json:"restaurant_cover_image"`
	RestaurantName       string   `json:"restaurant_name"`
	RestaurantLocality   string   `json:"restaurant_locality"`
	RestaurantCusine     []string `json:"restaurant_cuisine"`
}

type RatingMeta struct {
	RestaurantRating struct {
		Rating int `json:"rating"`
	} `json:"restaurant_rating"`
	DeliveryRating struct {
		Rating int `json:"rating"`
	} `json:"delivery_rating"`
}
type DeliveryBoy struct {
	Name string `json:"name"`
}

func RestaurantFromOrder(order Order) Restaurant {
	return Restaurant{
		RestaurantId:         order.RestaurantId,
		RestaurantCoverImage: order.RestaurantCoverImage,
		RestaurantName:       order.RestaurantName,
		RestaurantLocality:   order.RestaurantLocality,
		RestaurantCusine:     order.RestaurantCusine,
	}
}
