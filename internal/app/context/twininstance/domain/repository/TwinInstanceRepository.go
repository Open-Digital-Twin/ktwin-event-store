package repository

import (
	"agwermann/dt-service/internal/app/context/twininstance/domain"
	"agwermann/dt-service/internal/app/infra/db"
	"time"

	"github.com/scylladb/gocqlx/v2/qb"
)

const (
	TWIN_INSTANCE_TABLE = "twin_instance"
)

func NewTwinInstanceRepository(
	dbConnection db.DBConnection,
	mapper TwinInstanceMapper,
) TwinInstanceRepository {
	return &twinInstanceRepository{
		dbConnection: dbConnection,
		mapper:       mapper,
	}
}

type TwinInstance struct {
	Id            string    `db:"id"`
	Name          string    `db:"name"`
	InterfaceId   string    `db:"interface_id"`
	Active        bool      `db:"active"`
	Parent        string    `db:"parent"`
	LastEventData string    `db:"last_event_data"`
	CreatedAt     time.Time `db:"created_at"`
}

type TwinInstanceRepository interface {
	GetAllTwinInstances() ([]domain.TwinInstance, error)
	GetTwinInstance(id string) (domain.TwinInstance, error)
	InsertTwinInstance(twinInstance domain.TwinInstance) error
	DeleteTwinInstance(id string) error
}

type twinInstanceRepository struct {
	dbConnection db.DBConnection
	mapper       TwinInstanceMapper
}

func (t *twinInstanceRepository) GetAllTwinInstances() ([]domain.TwinInstance, error) {
	var twinInstances []TwinInstance

	queryParameters := db.NewQueryParameters(TWIN_INSTANCE_TABLE, nil, nil)
	err := t.dbConnection.GetManyWithParameters(queryParameters, &twinInstances)

	if err != nil {
		return []domain.TwinInstance{}, err
	}

	return t.mapper.ToDomainList(twinInstances), err
}

func (t *twinInstanceRepository) GetTwinInstance(id string) (domain.TwinInstance, error) {
	var twinInstance TwinInstance

	whereStatement, parameterValues := t.getOneTwinInstanceWhereClause(id)
	queryParameters := db.NewQueryParameters(TWIN_INSTANCE_TABLE, whereStatement, parameterValues)
	err := t.dbConnection.GetOneWithParameters(queryParameters, &twinInstance)

	if err != nil {
		return domain.TwinInstance{}, err
	}

	return t.mapper.ToDomain(twinInstance), err
}

func (t *twinInstanceRepository) getOneTwinInstanceWhereClause(id string) ([]qb.Cmp, map[string]interface{}) {
	whereStatement := []qb.Cmp{}
	parameterValues := make(map[string]interface{}, 10)

	if id != "" {
		whereStatement = append(whereStatement, qb.Eq("id"))
		parameterValues["id"] = id
	}

	return whereStatement, parameterValues
}

func (*twinInstanceRepository) InsertTwinInstance(twinInstance domain.TwinInstance) error {
	return nil
}

func (*twinInstanceRepository) DeleteTwinInstance(id string) error {
	return nil
}
