package service


import (
	"github.com/lucasferreirajs/grpc/internal/database"
	"github.com/lucasferreirajs/grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB
	}


}


func (UnimplementedCategoryServiceServer) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil{
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id: category.Id,
		Name: category.Name,
		Description: category.Description,
	}

	return &pb.categoryResponse{
		Category: categoryResponse,
	}, nil
}