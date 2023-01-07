package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-restapi/helper"
	"golang-restapi/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

//Fungsi untuk simpan data
func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "INSERT INTO categories(name) values (?)"
	result, err := tx.ExecContext(ctx, query, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	category.Id = int(id)
	return category
}

//Fungsi untuk Update Data
func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "INSERT categories set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, query, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	query := "DELETE from categories where id = ?"
	_, err := tx.ExecContext(ctx, query, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	query := "SELECT id, name from categories where id = ?"
	rows, err := tx.QueryContext(ctx, query, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()
	category := domain.Category{}

	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("Category is Not Found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	query := "SELECT id, name from categories"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
