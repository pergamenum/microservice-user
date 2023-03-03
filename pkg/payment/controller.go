package payment

import (
	"context"
	"time"

	commonpb "github.com/pergamenum/protobuf/golang/common"
	paymentpb "github.com/pergamenum/protobuf/golang/payment"
)

type Controller struct {
	paymentpb.UnimplementedCustomerServer
	paymentpb.UnimplementedOrderServer
}

func NewController() *Controller {
	return &Controller{}
}

func (c Controller) GetCustomer(_ context.Context, request *paymentpb.GetCustomerRequest) (*paymentpb.GetCustomerResponse, error) {

	println(time.Now().String())

	RID := request.GetId()

	if RID == "666" {
		e := &commonpb.Error{
			Status:  "666",
			Message: "Evil Error",
		}
		result := &paymentpb.GetCustomerResponse_Error{Error: e}
		response := &paymentpb.GetCustomerResponse{Result: result}
		return response, nil
	}

	is := []Item{
		{
			ID:    "300",
			Name:  "Minor",
			Price: 300,
		},
		{
			ID:    "700",
			Name:  "Major",
			Price: 700,
		},
	}

	os := []Order{{
		ID:         "1",
		CustomerID: RID,
		Items:      is,
		Total:      0,
	}}

	for i := range os {
		os[i].calcTotal()
	}

	cus := customerToPB(Customer{
		ID:           RID,
		OrderHistory: os,
	})
	result := &paymentpb.GetCustomerResponse_Customer{Customer: cus}
	response := &paymentpb.GetCustomerResponse{Result: result}
	return response, nil
}

func customerToPB(input Customer) *paymentpb.Customer {

	var pbos []*paymentpb.Order
	for _, o := range input.OrderHistory {

		var pbis []*paymentpb.Item
		for _, i := range o.Items {
			pbi := &paymentpb.Item{
				Id:    i.ID,
				Name:  i.Name,
				Price: i.Price,
			}
			pbis = append(pbis, pbi)
		}

		pbo := &paymentpb.Order{
			Id:         o.ID,
			CustomerId: o.CustomerID,
			Items:      pbis,
			Total:      o.Total,
		}
		pbos = append(pbos, pbo)
	}

	r := &paymentpb.Customer{
		Id:           input.ID,
		OrderHistory: pbos,
	}
	return r
}

func (c Controller) GetOrder(_ context.Context, _ *paymentpb.GetOrderRequest) (*paymentpb.GetOrderResponse, error) {
	panic("implement me")
}

type Customer struct {
	ID           string
	OrderHistory []Order
}

type Order struct {
	ID         string
	CustomerID string
	Items      []Item
	Total      float32
}

func (o *Order) calcTotal() {

	if len(o.Items) == 0 {
		return
	}

	var sum float32
	for _, i := range o.Items {
		sum = sum + i.Price
	}
	o.Total = sum
}

type Item struct {
	ID    string
	Name  string
	Price float32
}
