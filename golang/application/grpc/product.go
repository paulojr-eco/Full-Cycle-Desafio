package grpc

import (
	"context"

	"github.com/paulojr-eco/full-cycle-desafio/application/grpc/pb"
	"github.com/paulojr-eco/full-cycle-desafio/application/usecase"
)

type ProductGrpcService struct {
	ProductUseCase usecase.ProductUseCase
	pb.UnimplementedProductServiceServer
}

func (p *ProductGrpcService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product, err := p.ProductUseCase.AddProduct(in.Name, in.Description, in.Price)
	if err != nil {
		return &pb.CreateProductResponse{
			Product: &pb.Product{},
		}, err
	}

	return &pb.CreateProductResponse{
		Product: &pb.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		},
	}, nil
}

func (p *ProductGrpcService) FindProducts(context.Context, *pb.FindProductsRequest) (*pb.FindProductsResponse, error) {
	products := p.ProductUseCase.FindProducts()

	productsPB := make([]*pb.Product, 0)

	for _, product := range *products {
		productPB := &pb.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		}
		productsPB = append(productsPB, productPB)
	}

	return &pb.FindProductsResponse{
		Products: productsPB,
	}, nil
}

func NewProductGrpcService(usecase usecase.ProductUseCase) *ProductGrpcService {
	return &ProductGrpcService{
		ProductUseCase: usecase,
	}
}
