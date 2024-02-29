package services

import (
	"github.com/couchbase/gocb/v2"
	"github.com/saitunurlu/golang/models"
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
	query := gocb.NewN1qlQuery("SELECT * FROM `" + s.collectionName + "` WHERE META().id = $1")
	params := []interface{}{docKey}
	rows, err := s.scope.ExecuteN1qlQuery(query, params)
	if err != nil {
		return nil, err
	}

	var categorysData models.Categorys
	if err := rows.One(&categorysData); err != nil {
		return nil, err
	}
	return &categorysData, nil
}
