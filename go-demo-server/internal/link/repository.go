package link

import (
	"goadvancedserver/pkg/db"

	"gorm.io/gorm/clause"
)

type Repository struct {
	Database *db.DB
}

func NewRepository(db *db.DB) *Repository {
	return &Repository{
		Database: db,
	}
}

func (repo *Repository) Create(link *Link) (*Link, error) {
	result := repo.Database.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *Repository) Find(id uint) (*Link, error) {
	var link Link
	result := repo.Database.DB.First(&link, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *Repository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repo.Database.DB.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *Repository) Update(link *Link) (*Link, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(link)

	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *Repository) Delete(id uint) error {
	result := repo.Database.DB.Delete(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *Repository) GetList(limit, offset int) []Link {
	var links []Link

	repo.Database.
		Table("links").
		Where("deleted_at IS NULL").
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Scan(&links)

	return links
}

func (repo *Repository) Count() int64 {
	var count int64

	repo.Database.
		Table("links").
		Where("deleted_at IS NULL").
		Count(&count)

	return count
}
