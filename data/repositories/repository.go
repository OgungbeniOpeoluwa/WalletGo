package repositories

import (
	"WalletService/db"
	"WalletService/util"
	"fmt"
	"gorm.io/gorm"
)

type BaseRepositoryInterface[T any] interface {
	Save(value T) (T, error)
	FindById(id any) (T, error)
	GetAllBy(name string, id any) ([]T, error)
	FindByEmail(value string) (T, error)
	GetBy(name string, value any) (T, error)
}

type BaseRepository[T any] struct {
	baseRepository BaseRepositoryInterface[T]
	db             *gorm.DB
}

func NewBaseRepository[T any]() *BaseRepository[T] {
	//config.Load("../.env")
	database, _ := db.DbConnection()
	return &BaseRepository[T]{db: database}
}

func (repository *BaseRepository[T]) Save(value T) (T, error) {
	var err error
	result := repository.db.Save(&value)
	if result.Error != nil {
		err = result.Error
		return value, err
	}
	return value, err

}
func (repository *BaseRepository[T]) FindById(id any) (T, error) {
	var err error
	var results T
	result := repository.db.Where("id=?", id).First(&results)
	if result.Error != nil {
		err = util.ErrFetching
		return results, err
	}

	return results, err

}
func (repository *BaseRepository[T]) GetAllBy(name string, value any) ([]T, error) {
	var err error
	var results []T
	query := fmt.Sprintf("%s = ?", name)
	result := repository.db.Where(query, value).Find(&results)
	if result.Error != nil {
		err = result.Error
		return results, err
	}
	return results, err

}

func (repository *BaseRepository[T]) FindByEmail(email string) (T, error) {
	var err error
	var results T
	result := repository.db.Where("email=?", email).First(&results)
	if result.Error != nil {
		err = result.Error
		return results, err
	}

	return results, err

}

func (repository *BaseRepository[T]) GetBy(name string, value any) (T, error) {
	var err error
	var results T
	query := fmt.Sprintf("%s = ?", name)
	result := repository.db.Where(query, value).First(&results)
	if result.Error != nil {
		err = result.Error
		return results, err
	}
	return results, err

}
