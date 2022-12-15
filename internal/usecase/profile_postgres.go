package usecase

import (
	"fmt"
	"log"
	"ptok/internal/entity"
	"ptok/pkg/postgres"
)

type ProfilePostgres struct {
	*postgres.Postgres
}

func NewPG(pg *postgres.Postgres) *ProfilePostgres {
	return &ProfilePostgres{pg}
}

func (p *ProfilePostgres) GetProfiles() ([]entity.Profile, error) {

	rows, err := p.Postgres.DB.Query("SELECT ID, NAME, AGE FROM TEST.PROFILES")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	profiles := make([]entity.Profile, 0)

	for rows.Next() {
		prfl := entity.Profile{}

		err := rows.Scan(&prfl.Id, &prfl.Name, &prfl.Age)
		if err != nil {
			log.Println("GetProfiles from Postgres: ", err)
			continue
		}

		profiles = append(profiles, prfl)

	}

	if len(profiles) == 0 {
		return nil, fmt.Errorf("profiles from postgres is empty")
	}
	return profiles, nil
}

func (p *ProfilePostgres) GetProfilesV2(profileChan chan entity.Profile) error {

	rows, err := p.Postgres.DB.Query("SELECT ID, NAME, AGE FROM TEST.PROFILES")
	if err != nil {
		return err
	}

	defer rows.Close()
	//go func() {
	for rows.Next() {
		prfl := entity.Profile{}

		err := rows.Scan(&prfl.Id, &prfl.Name, &prfl.Age)
		if err != nil {
			log.Println("GetProfiles from Postgres: ", err)
			continue
		}

		profileChan <- prfl

	}
	//}()
	defer close(profileChan) // приходится искать где закрывается канал, но так работает

	return nil
}
