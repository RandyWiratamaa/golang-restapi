package service

import (
	"context"
	"golang-restapi/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ct context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.WebResponse
	findAll(ctx context.Context) []web.CategoryResponse
}
