package usecase

import (
	"database/sql"
	"testing"

	"github.com/devfullcyle/gointensivo/internal/infra/database"
	"github.com/devfullcyle/gointensivo/internal/order/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type CalculateFinalPriceUseCaseTestSuite struct {
	suite.Suite
	OrderRepository database.OrderRepository
	Db              *sql.DB
}

func (suite *CalculateFinalPriceUseCaseTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
	suite.OrderRepository = *database.NewOrderRepository(db)
}

func (suite *CalculateFinalPriceUseCaseTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CalculateFinalPriceUseCaseTestSuite))
}

func (suite *CalculateFinalPriceUseCaseTestSuite) TestCalculateFinalPrice() {
	order, err := entity.NewOrder("123", 10.0, 2)
	suite.NoError(err)
	order.CalculateFinalPrice()

	calculateFinalPriceInput := OrderInputDTO{
		ID:    order.ID,
		Price: order.Price,
		Tax:   order.Tax,
	}

	calculateFinalPriceUseCase := NewCalculateFinalPriceUseCase(suite.OrderRepository)
	output, err := calculateFinalPriceUseCase.Execute(calculateFinalPriceInput)

	suite.NoError(err)
	suite.Equal(order.ID, output.ID)
	suite.Equal(order.Price, output.Price)
	suite.Equal(order.Tax, output.Tax)
}
