package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/notblessy/go-listing/model"
	"github.com/notblessy/go-listing/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

var productColumns = []string{"id", "name", "price", "description", "quantity", "created_at", "updated_at"}

func initPostgreMock() (db *gorm.DB, mock sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	db, err = gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return
}

// TestProductRepo_Create :nodoc:
func TestProductRepo_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbMock, sqlMock := initPostgreMock()

	pr := productRepo{
		db: dbMock,
	}

	product := &model.Product{
		ID:          utils.GenerateID(),
		Name:        "Oud Vibrant Leather",
		Price:       399999,
		Description: "A Perfume from Zara",
		Quantity:    10,
	}

	t.Run("Success", func(t *testing.T) {
		sqlMock.ExpectBegin()
		queryResult := sqlmock.NewRows([]string{"id"}).
			AddRow(product.ID)
		sqlMock.ExpectQuery("INSERT INTO \"products\"").
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(queryResult)
		sqlMock.ExpectCommit()

		err := pr.db.Create(product).Error
		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		sqlMock.ExpectBegin()
		queryResult := sqlmock.NewRows([]string{"id"}).
			AddRow(product.ID)
		sqlMock.ExpectQuery("INSERT INTO \"products\"").
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(queryResult)
		sqlMock.ExpectCommit()

		err := pr.db.Create(product).Error
		assert.Error(t, err)
	})

}

// TestProductRepo_FindAll :nodoc:
func TestProductRepo_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbMock, sqlMock := initPostgreMock()

	pr := productRepo{
		db: dbMock,
	}

	product := &model.Product{
		ID:          utils.GenerateID(),
		Name:        "Oud Vibrant Leather",
		Price:       399999,
		Description: "A Perfume from Zara",
		Quantity:    10,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.Run("Success", func(t *testing.T) {
		sqlMock.ExpectBegin()

		dbMock.Begin()
		rows := sqlmock.NewRows(productColumns).
			AddRow(product.ID, product.Name, product.Price, product.Description, product.Quantity, product.CreatedAt, product.UpdatedAt)

		sqlMock.ExpectQuery("^SELECT .+ FROM \"products\"").
			WillReturnRows(rows)

		err := pr.db.Find(&[]model.Product{}).Error
		assert.NoError(t, err)
	})

	t.Run("Success", func(t *testing.T) {
		sqlMock.ExpectBegin()

		dbMock.Begin()
		sqlMock.ExpectQuery("^SELECT .+ FROM \"products\"").
			WillReturnError(gorm.ErrRecordNotFound)

		err := pr.db.Find(&[]model.Product{}).Error
		assert.Error(t, err)
	})
}
