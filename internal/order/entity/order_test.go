package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyId_WhenCreateNewOrder_ThenSHouldReceiveError(t *testing.T) {
	order := Order{}
	assert.Error(t, order.isValid(), "isvalid id")
}

func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenShoudlReveiveAnError(t *testing.T) {
	order := Order{Price: 10.0}
	assert.Error(t, order.isValid(), "invalid price")
}

func TestGivenAnEmptyTax_WhenCreateANewOrder_ThenShoudlReveiveAnError(t *testing.T) {
	order := Order{Tax: 1}
	assert.Error(t, order.isValid(), "invalid tax")
}

func TestGivenAValidParams_WhenICallNewOrder_ThenIShouldReceiveCreatedOrderWithAllParams(t *testing.T) {
	order := Order{ID: "123", Price: 10.0, Tax: 2.0}
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	assert.Nil(t, order.isValid())
}

func TestGivenAValidParams_WhenICallNewOrderFunc_ThenIShouldReceiveCreatedOrderWithAllParams(t *testing.T) {
	order, error := NewOrder("123", 10.0, 2.0)
	assert.Nil(t, error)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
}

func TestGivenAPriceAndTax_WhenICallCalculateTax_ThenIShouldSetFinalPrice(t *testing.T) {
	order, error := NewOrder("123", 10.0, 2.0)
	assert.Nil(t, error)
	assert.Nil(t, order.CalculateFinalPrice())
	assert.Equal(t, 12.0, order.FinalPrice)
}
