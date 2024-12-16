package server

import (
	"context"
	"github/cars/db"
	cars "github/cars/gen"
	"github/cars/model"
	"github/cars/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CarsServiceServer struct {
	cars.UnimplementedCarsServiceServer
}

func (c *CarsServiceServer) GetAll(ctx context.Context, empty *emptypb.Empty) (*cars.CarsList, error) {
	var carsEntity []model.Cars

	db.DB.Find(&carsEntity)

	var carsList []*cars.Car

	for _, car := range carsEntity {
		carsList = append(carsList, &cars.Car{
			Id:    int32(car.ID),
			Price: car.Price,
			Name:  car.Name,
		})
	}

	return &cars.CarsList{
		Cars: carsList,
	}, nil
}
func (c *CarsServiceServer) GetOne(ctx context.Context, body *cars.OneCarRequest) (*cars.OneCars, error) {
	var carEntity model.Cars

	result := db.DB.Where("id = ?", body.Id).First(&carEntity)

	if result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "Car not found")
	}

	return &cars.OneCars{
		Car: &cars.Car{
			Id:    int32(carEntity.ID),
			Name:  carEntity.Name,
			Price: carEntity.Price,
		},
	}, nil
}
func (c *CarsServiceServer) Create(ctx context.Context, body *cars.CreateCarRequest) (*cars.OneCars, error) {
	err := utils.VerifyToken(ctx)
	if err != nil {
		return nil, err
	}

	var newCar = &model.Cars{
		Name:  body.Name,
		Price: body.Price,
	}

	result := db.DB.Create(&newCar)
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.AlreadyExists, "Car already exists")
	}

	return &cars.OneCars{
		Car: &cars.Car{
			Id:    int32(newCar.ID),
			Name:  newCar.Name,
			Price: newCar.Price,
		},
	}, nil
}
func (c *CarsServiceServer) Update(ctx context.Context, body *cars.UpdateCarRequest) (*cars.OneCars, error) {
	err := utils.VerifyToken(ctx)
	if err != nil {
		return nil, err
	}

	var updateCar = &model.Cars{
		ID:    uint64(body.Id),
		Name:  body.Name,
		Price: body.Price,
	}

	result := db.DB.Updates(&updateCar)
	if result.RowsAffected == 0 {
		return nil, status.Error(codes.AlreadyExists, "Car already exists")
	}

	return &cars.OneCars{
		Car: &cars.Car{
			Id:    int32(updateCar.ID),
			Name:  updateCar.Name,
			Price: updateCar.Price,
		},
	}, nil
}
