package legalentities

import (
	"errors"

	"github.com/google/uuid"
	"github.com/krisch/crm-backend/dto"
	"github.com/krisch/crm-backend/internal/helpers"
	"github.com/krisch/crm-backend/pkg/postgres"
	"github.com/prometheus/client_golang/prometheus"
	"gorm.io/gorm"
)

type Repository struct {
	gorm    *postgres.GDB
	counter *prometheus.CounterVec
}

func (r *Repository) WithCounter() *Repository {
	if r.counter != nil {
		c := *r.counter
		c.With(prometheus.Labels{"repo": "total"}).Inc()
	}
	return r
}

func NewRepository(db *postgres.GDB, metrics *helpers.MetricsCounters) *Repository {
	return &Repository{
		gorm:    db,
		counter: metrics.RepoCounter,
	}
}

func (r *Repository) getAllLegalEntities() (items []LegalEntity, err error) {
	if r.counter != nil {
		r.counter.With(prometheus.Labels{"repo": "getAllLegalEntities"}).Inc()
	}
	err = r.gorm.DB.
		Model(&LegalEntity{}).
		Where("deleted_at is null").
		Order("created_at DESC").
		Find(&items).Error
	return items, err
}

func (r *Repository) CreateLegalEntity(name string) (item LegalEntity, err error) {
	if r.counter != nil {
		r.counter.With(prometheus.Labels{"repo": "CreateLegalEntity"}).Inc()
	}
	item = LegalEntity{Name: name}
	err = r.gorm.DB.Create(&item).Error
	return item, err
}

func (r *Repository) UpdateLegalEntity(uid uuid.UUID, name string) error {
	if r.counter != nil {
		r.counter.With(prometheus.Labels{"repo": "UpdateLegalEntity"}).Inc()
	}
	res := r.gorm.DB.
		Model(&LegalEntity{}).
		Where("uuid = ? AND deleted_at is null", uid).
		Updates(map[string]interface{}{
			"name":       name,
			"updated_at": gorm.Expr("now()"),
		})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return dto.NotFoundErr("юридическое лицо не найдено")
	}
	return nil
}

func (r *Repository) DeleteLegalEntity(uid uuid.UUID) error {
	if r.counter != nil {
		r.counter.With(prometheus.Labels{"repo": "DeleteLegalEntity"}).Inc()
	}
	res := r.gorm.DB.
		Model(&LegalEntity{}).
		Where("uuid = ? AND deleted_at is null", uid).
		Updates(map[string]interface{}{
			"deleted_at": gorm.Expr("now()"),
			"updated_at": gorm.Expr("now()"),
		})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return dto.NotFoundErr("юридическое лицо уже удалено или не найдено")
	}
	return nil
}

func (r *Repository) getByUUID(uid uuid.UUID, fields ...string) (item LegalEntity, err error) {
	if len(fields) == 0 {
		fields = []string{"uuid", "name", "created_at", "updated_at", "deleted_at"}
	}
	err = r.gorm.DB.
		Model(&LegalEntity{}).
		Where("uuid = ?", uid).
		Select(fields).
		First(&item).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return item, dto.NotFoundErr("юридическое лицо не найдено")
	}
	return item, err
}
