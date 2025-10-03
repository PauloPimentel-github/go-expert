package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTaxAndSave(t *testing.T) {
	//cenário
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil).Twice()
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax!"))

	//ação
	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)
	err = CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	//ação
	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "error saving tax")

	//validação
	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 3)
}
