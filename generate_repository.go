package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func generateRepository(directory string, name string) {
	camelName := snakeCaseToCamelCase(name)
	titleName := strings.Title(camelName)
	path := directory + "repository/"

	f, err := os.Create(path + name + "_repository.go")
	check(err)

	defer f.Close()

	d1 := []byte(`
package repository

import (
	log "github.com/sirupsen/logrus"
)

type ` + titleName + `RepositoryInterface interface {
	ArangoBaseRepositoryInterface
}

type ` + camelName + `Repository struct {
	arangoBaseRepository
}

func New` + titleName + `Repository(logger *log.Logger, arangoDB arangodb.ArangoDB, rmq rabbitmq.RabbitMQ) ` + titleName + `RepositoryInterface {
	repo := ` + camelName + `Repository{
		arangoBaseRepository: arangoBaseRepository{
			collection: "` + name + `s",
			logger:     logger,
			arangoDB:   arangoDB,
			rmq:        rmq,
		},
	}

	return &repo
}

func (r *` + camelName + `Repository) First(c *gin.Context, request arango.DigitalPaymentDisbursementTransaction) (arango.DigitalPaymentDisbursementTransaction, error) {
	err := r.arangoBaseRepository.first(c, &request)

	return request, err
}

func (r *` + camelName + `Repository) All(c *gin.Context, request arango.DigitalPaymentDisbursementTransaction, baseFilter dto.LazyFilters) ([]arango.DigitalPaymentDisbursementTransaction, int64, error) {
	var response []arango.DigitalPaymentDisbursementTransaction

	result, totalRecords, err := r.arangoBaseRepository.all(c, &request, baseFilter)
	if err != nil {
		return response, 0, err
	}

	bytedata, _ := json.Marshal(result)
	json.Unmarshal(bytedata, &response)

	return response, totalRecords, err
}

func (r *` + camelName + `Repository) Create(c *gin.Context, request *arango.DigitalPaymentDisbursementTransaction) error {
	return r.arangoBaseRepository.create(c, request)
}

func (r *` + camelName + `Repository) Update(c *gin.Context, request *arango.DigitalPaymentDisbursementTransaction) error {
	return r.arangoBaseRepository.update(c, request)
}

func (r *` + camelName + `Repository) Delete(c *gin.Context, request *arango.DigitalPaymentDisbursementTransaction) error {
	return r.arangoBaseRepository.delete(c, request)
}
			`)
	err = ioutil.WriteFile(path+name+"_repository.go", d1, 0777)
	check(err)
}
