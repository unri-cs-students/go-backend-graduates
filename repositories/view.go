package repositories

import (
	"github.com/SemmiDev/go-backend/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type viewRepository struct {
	db *gorm.DB
}

func NewViewRepository(db *gorm.DB) entities.ViewRepository {
	return &viewRepository{db: db}
}

func (repo *viewRepository) AddOne(IdWisudawan string, IP string, Time time.Time) error {
	v := entities.View{
		GraduatesID: IdWisudawan,
		IP:          IP,
		AccessTime:  Time,
	}
	if err := repo.db.Create(&v).Error; err != nil {
		return err
	}
	return nil
}

func (repo *viewRepository) GetLast(IdWisudawan string, IP string) (entities.View, error) {
	var view entities.View
	if err := repo.db.
		Order("access_time DESC").
		Where("graduates_id = ?", IdWisudawan).
		Where("ip = ?", IP).
		Last(&view).Error; err != nil {
		return view, err
	}
	return view, nil
}

func (repo *viewRepository) GetTop5() ([]entities.GetViewGraduates, error) {
	var result []entities.GetViewGraduates
	size := repo.db.Find(&[]entities.View{}).RowsAffected
	if size > 5 {
		if err := repo.db.Raw("SELECT graduates.id,count from (SELECT graduates_id as id, count(id) as count FROM \"view\" GROUP BY \"graduates_id\" LIMIT 5) T INNER JOIN graduates ON T.id = graduates.id INNER JOIN major ON graduates.major_id = major.id ORDER BY count desc").Scan(&result).Error; err != nil {
			return nil, err
		}
		for i := range result {
			var graduates entities.Graduates
			if err := repo.db.Preload(clause.Associations).Find(&graduates, "id = ?", result[i].Graduates.ID).Error; err != nil {
				return nil, err
			}
			result[i].Graduates = graduates
		}
	}
	return result, nil
}