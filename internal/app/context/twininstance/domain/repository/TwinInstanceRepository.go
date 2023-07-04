package repository

import (
	"agwermann/dt-service/internal/app/context/twininstance/domain"
	"agwermann/dt-service/internal/app/infra/db"
	"time"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

var TWIN_INSTANCE_METADATA = table.Metadata{
	Name:    "twin_instance",
	Columns: []string{"id", "name", "interface_id", "active", "parent", "last_event_data", "created_at"},
	PartKey: []string{"interface_id", "id"},
	SortKey: []string{},
}

var TWIN_INSTANCE_TABLE = table.New(TWIN_INSTANCE_METADATA)

type TwinInstance struct {
	Id            string    `db:"id"`
	Name          string    `db:"name"`
	InterfaceId   string    `db:"interface_id"`
	Active        bool      `db:"active"`
	Parent        string    `db:"parent"`
	LastEventData string    `db:"last_event_data"`
	CreatedAt     time.Time `db:"created_at"`
}

func NewTwinInstanceRepository(
	dbConnection db.DBConnection,
	mapper TwinInstanceMapper,
) TwinInstanceRepository {
	return &twinInstanceRepository{
		dbConnection: dbConnection,
		mapper:       mapper,
	}
}

type TwinInstanceRepository interface {
	GetAllTwinInstances() ([]domain.TwinInstance, error)
	GetOneTwinInstance(interfaceId string, id string) (domain.TwinInstance, error)
	InsertTwinInstance(twinInstance domain.TwinInstance) error
	DeleteTwinInstance(interfaceId string, id string) error
}

type twinInstanceRepository struct {
	dbConnection db.DBConnection
	mapper       TwinInstanceMapper
}

func (t *twinInstanceRepository) GetAllTwinInstances() ([]domain.TwinInstance, error) {
	var twinInstances []TwinInstance

	err := t.dbConnection.GetManyWithParameters(TWIN_INSTANCE_TABLE, qb.M{}, &twinInstances)

	if err != nil {
		return []domain.TwinInstance{}, err
	}

	return t.mapper.ToDomainList(twinInstances), err
}

func (t *twinInstanceRepository) GetOneTwinInstance(interfaceId string, id string) (domain.TwinInstance, error) {
	var twinInstance TwinInstance

	err := t.dbConnection.GetOneWithParameters(TWIN_INSTANCE_TABLE, t.getConditions(interfaceId, id), &twinInstance)

	if err != nil {
		return domain.TwinInstance{}, err
	}

	return t.mapper.ToDomain(twinInstance), err
}

func (t *twinInstanceRepository) InsertTwinInstance(twinInstance domain.TwinInstance) error {
	twinInstanceDB := TwinInstance(twinInstance)
	twinInstanceDB.CreatedAt = time.Now()

	return t.dbConnection.InsertQueryDB(TWIN_INSTANCE_TABLE, twinInstanceDB)
}

func (t *twinInstanceRepository) DeleteTwinInstance(interfaceId string, id string) error {
	return t.dbConnection.DeleteQuery(TWIN_INSTANCE_TABLE, t.getConditions(interfaceId, id))
}

func (t *twinInstanceRepository) getConditions(interfaceId string, id string) qb.M {
	return qb.M{"interface_id": interfaceId, "id": id}
}
