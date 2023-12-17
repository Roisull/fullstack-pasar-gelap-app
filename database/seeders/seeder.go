package seeders

import (
	"pasar-gelap/database/fakers"

	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{}
}

func RegisterSeeders(db *gorm.DB) []Seeder {
	return []Seeder{
		{Seeder: fakers.UserFaker(db)},
		{Seeder: fakers.ProductFaker(db)},
	}
}

func DBSeed(db *gorm.DB) error {
	for _, s := range RegisterSeeders(db) {
		err := db.Debug().Create(s.Seeder).Error
		if err != nil {
			return err
		}
	}

	return nil
}
