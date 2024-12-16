package server

import (
	"context"
	"github/order/clients"
	"github/order/db"
	order "github/order/gen"
	"github/order/gen/cars"
	"github/order/model"
	"github/order/utils"
	"google.golang.org/protobuf/types/known/emptypb"
)

type OrderServiceServer struct {
	order.UnimplementedOrderServiceServer
}

func (o *OrderServiceServer) GetAll(ctx context.Context, body *emptypb.Empty) (*order.OrderList, error) {
	var orders []*model.Order

	db.DB.Find(&orders)

	var orderList []*order.Order
	for _, orderEntity := range orders {
		car, err := clients.CarsClient.GetOne(context.Background(), &cars.OneCarRequest{
			Id: int32(orderEntity.CarsID),
		})

		if err != nil {
			orderList = append(orderList, &order.Order{
				Id:  int32(orderEntity.ID),
				Car: nil,
			})

			continue
		}

		orderList = append(orderList, &order.Order{
			Id:  int32(orderEntity.ID),
			Car: car.Car,
		})
	}

	return &order.OrderList{Orders: orderList}, nil
}

func (o *OrderServiceServer) GetOne(ctx context.Context, body *order.OneOrderRequest) (*order.OneOrder, error) {
	var orderEntity = &model.Order{}

	db.DB.First(&orderEntity, body.Id)

	var car *cars.Car
	oneCars, err := clients.CarsClient.GetOne(context.Background(), &cars.OneCarRequest{
		Id: int32(orderEntity.CarsID),
	})

	if err != nil {
		car = nil
	} else {
		car = oneCars.Car
	}

	if orderEntity.ID == 0 {
		return &order.OneOrder{Order: nil}, nil
	}

	return &order.OneOrder{
		Order: &order.Order{
			Id:  int32(orderEntity.ID),
			Car: car,
		},
	}, nil
}

func (o *OrderServiceServer) Create(ctx context.Context, body *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	err := utils.VerifyToken(ctx)
	if err != nil {
		return nil, err
	}

	oneCars, err := clients.CarsClient.GetOne(context.Background(), &cars.OneCarRequest{
		Id: body.CarId,
	})

	if err != nil {
		return nil, err
	}

	createOrder := &model.Order{
		CarsID: uint64(oneCars.Car.Id),
	}
	db.DB.Create(&createOrder)

	return &order.CreateOrderResponse{
		OrderId: int32(createOrder.ID),
	}, nil
}
