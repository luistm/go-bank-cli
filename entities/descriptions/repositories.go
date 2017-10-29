package descriptions

import (
	"fmt"

	"github.com/luistm/go-bank-cli/entities"
)

var saveStatement = "INSERT INTO descriptions(slug, friendlyName ) VALUES (?, ?)"

type repository struct {
	SQLStorage entities.SQLDatabaseHandler
}

func (r *repository) Save(ent entities.Entity) error {

	if r.SQLStorage == nil {
		return entities.ErrInfrastructureUndefined
	}

	d := ent.(*Description)
	err := r.SQLStorage.Execute(saveStatement, d.rawName, d.friendlyName)
	if err != nil {
		return &entities.ErrInfrastructure{Msg: err.Error()}
	}

	return nil
}

func (r *repository) Get(d string) (entities.Entity, error) {
	return &Description{}, nil
}

// GetAll fetches all descriptions
func (r *repository) GetAll() ([]entities.Entity, error) {
	statement := "SELECT * FROM descriptions"
	rows, err := r.SQLStorage.Query(statement)
	if err != nil {
		return []entities.Entity{}, fmt.Errorf("Database failure: %s", err)
	}

	descriptions := []entities.Entity{}

	for rows.Next() {
		var rawName int
		var friendlyName string
		err := rows.Scan(&rawName, &friendlyName)
		if err != nil {
			return nil, err
		}
		descriptions = append(descriptions, &Description{})
	}

	return descriptions, nil
}