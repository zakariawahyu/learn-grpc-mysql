package services

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"learn-grpc-mysql/cmd/helpers"
	paginationPb "learn-grpc-mysql/pb/pagination"
	productPb "learn-grpc-mysql/pb/product"
	"log"
)

type ProductServices struct {
	productPb.UnimplementedProductServiceServer
	DB *gorm.DB
}

func (p *ProductServices) GetProducts(ctx context.Context, param *productPb.Page) (*productPb.Products, error) {
	var page int64 = 1

	if param.GetPage() != 0 {
		page = param.GetPage()
	}

	pagination := paginationPb.Pagination{}
	products := []*productPb.Product{}

	sql := p.DB.Table("products As p").
		Joins("LEFT JOIN categories as c ON p.category_id = c.id").
		Select("p.id, p.name, p.price, p.stock, c.id as category_id, c.name as category_name")

	offset, limit := helpers.Pagination(sql, page, &pagination)
	rows, err := sql.Offset(int(offset)).Limit(int(limit)).Rows()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		product := productPb.Product{}
		category := productPb.Category{}

		if err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &category.Id, &category.Name); err != nil {
			log.Fatal("Error rows data", err.Error())
		}
		product.Category = &category

		products = append(products, &product)
	}

	respose := &productPb.Products{
		Pagination: &pagination,
		Data:       products,
	}

	return respose, nil
}

func (p *ProductServices) GetProduct(ctx context.Context, id *productPb.Id) (*productPb.Product, error) {
	row := p.DB.Table("products As p").
		Joins("LEFT JOIN categories as c ON p.category_id = c.id").
		Select("p.id, p.name, p.price, p.stock, c.id as category_id, c.name as category_name").
		Where("p.id = ?", id.GetId()).
		Row()

	product := productPb.Product{}
	category := productPb.Category{}

	if err := row.Scan(&product.Id, &product.Name, &product.Price, &product.Stock, &category.Id, &category.Name); err != nil {
		log.Fatal("Error rows data", err.Error())
	}
	product.Category = &category

	return &product, nil
}
func (p *ProductServices) CreateProduct(ctx context.Context, productData *productPb.Product) (*productPb.Id, error) {
	response := productPb.Id{}

	err := p.DB.Transaction(func(tx *gorm.DB) error {
		category := productPb.Category{
			Id:   0,
			Name: productData.GetCategory().GetName(),
		}

		if err := tx.Table("categories").
			Where("LCASE(name) = ?", category.GetName()).
			FirstOrCreate(&category).Error; err != nil {
			return err
		}

		product := struct {
			Id          uint64
			Name        string
			Price       float64
			Stock       uint32
			Category_id uint32
		}{
			Id:          productData.GetId(),
			Name:        productData.GetName(),
			Price:       productData.GetPrice(),
			Stock:       productData.GetStock(),
			Category_id: category.GetId(),
		}

		if err := tx.Table("products").Create(&product).Error; err != nil {
			return err
		}

		response.Id = product.Id

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &response, nil
}
