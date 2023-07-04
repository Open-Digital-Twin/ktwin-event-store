package repository

import (
	"agwermann/dt-service/internal/app/context/twininterface/domain"
	"agwermann/dt-service/internal/app/infra/db"
	"time"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

var TWIN_INTERFACE_METADATA = table.Metadata{
	Name:    "twin_interface",
	Columns: []string{"id", "name", "active", "parent", "data_definition", "created_at"},
	PartKey: []string{"id"},
	SortKey: []string{},
}

var TWIN_INTERFACE_TABLE = table.New(TWIN_INTERFACE_METADATA)

type TwinInterface struct {
	Id             string    `db:"id"`
	Name           string    `db:"name"`
	Active         bool      `db:"active"`
	Parent         string    `db:"parent"`
	DataDefinition string    `db:"data_definition"`
	CreatedAt      time.Time `db:"created_at"`
}

type TwinInterfaceRepository interface {
	GetAllTwinInterfaces() ([]domain.TwinInterface, error)
	GetOneTwinInterface(id string) (domain.TwinInterface, error)
	DeleteTwinInterface(id string) error
	CreateTwinInterface(twinInterface domain.TwinInterface) error
}

func NewTwinInterfaceRepository(
	dbConnection db.DBConnection,
	mapper TwinInterfaceMapper,
) TwinInterfaceRepository {
	return &twinInterfaceRepository{
		dbConnection: dbConnection,
		mapper:       mapper,
	}
}

type twinInterfaceRepository struct {
	dbConnection db.DBConnection
	mapper       TwinInterfaceMapper
}

func (t *twinInterfaceRepository) GetAllTwinInterfaces() ([]domain.TwinInterface, error) {
	var twinInstances []TwinInterface

	err := t.dbConnection.GetManyWithParameters(TWIN_INTERFACE_TABLE, qb.M{}, &twinInstances)

	if err != nil {
		return []domain.TwinInterface{}, err
	}

	return t.mapper.ToDomainList(twinInstances), err
}

func (t *twinInterfaceRepository) GetOneTwinInterface(id string) (domain.TwinInterface, error) {
	var twinInstance TwinInterface

	err := t.dbConnection.GetOneWithParameters(TWIN_INTERFACE_TABLE, t.getConditions(id), &twinInstance)

	if err != nil {
		return domain.TwinInterface{}, err
	}

	return t.mapper.ToDomain(twinInstance), err
}

func (t *twinInterfaceRepository) CreateTwinInterface(twinInterface domain.TwinInterface) error {
	twinInstanceDB := TwinInterface(twinInterface)
	twinInstanceDB.CreatedAt = time.Now()

	return t.dbConnection.InsertQueryDB(TWIN_INTERFACE_TABLE, twinInstanceDB)
}

func (t *twinInterfaceRepository) DeleteTwinInterface(id string) error {
	return t.dbConnection.DeleteQuery(TWIN_INTERFACE_TABLE, t.getConditions(id))
}

func (t *twinInterfaceRepository) getConditions(id string) qb.M {
	return qb.M{"id": id}
}
