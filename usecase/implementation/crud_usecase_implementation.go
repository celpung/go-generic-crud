package crud_usecase_implementation

import (
	crud_repository "github.com/celpung/go-generic-crud/repository"
	crud_usecase "github.com/celpung/go-generic-crud/usecase"
)

type UsecaseStruct[T any] struct {
	repository crud_repository.CrudRepositoryInterface[T]
}

func (u *UsecaseStruct[T]) Create(entity *T) (*T, error) {
	return u.repository.Create(entity)
}

func (u *UsecaseStruct[T]) Read(page, limit int, sortBy string, conditions map[string]any, preloadFields ...string) ([]*T, error) {
	return u.repository.Read(page, limit, sortBy, conditions, preloadFields...)
}

func (u *UsecaseStruct[T]) ReadByID(id uint, preloadFields ...string) (*T, error) {
	return u.repository.ReadByID(id, preloadFields...)
}

func (u *UsecaseStruct[T]) Update(entity *T) (*T, error) {
	return u.repository.Update(entity)
}

func (u *UsecaseStruct[T]) Delete(id uint) error {
	return u.repository.Delete(id)
}

func (u *UsecaseStruct[T]) Search(query string, conditions map[string]any, preloadFields ...string) ([]*T, error) {
	return u.repository.Search(query, conditions, preloadFields...)
}

func (u *UsecaseStruct[T]) Count() (int64, error) {
	return u.repository.Count()
}

func NewUsecase[T any](repository crud_repository.CrudRepositoryInterface[T]) crud_usecase.UsecaseInterface[T] {
	return &UsecaseStruct[T]{repository: repository}
}
