package services

import (
	"github.com/couchbase-examples/golang-quickstart/models"
	"github.com/couchbase/gocb/v2"
)

type ICategorysService interface {
	CreateCategorys(string, *models.Categorys) error
	GetCategorys(string) (*models.Categorys, error)
}

type CategorysService struct {
	collectionName string
	scope          *gocb.Scope
}

func NewCategorysService(scope *gocb.Scope) *CategorysService {
	return &CategorysService{
		collectionName: "categorys",
		scope:          scope,
	}
}

func (s *CategorysService) CreateCategorys(docKey string, data *models.Categorys) error {
	_, err := s.scope.Collection(s.collectionName).Insert(docKey, data, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *CategorysService) GetCategorys(docKey string) (*models.Categorys, error) {
	getResult, err := s.scope.Collection(s.collectionName).Get(docKey, nil)
	if err != nil {
		return nil, err
	}
	var categorysData models.Categorys
	if err := getResult.Content(&categorysData); err != nil {
		return nil, err
	}
	return &categorysData, nil
}
