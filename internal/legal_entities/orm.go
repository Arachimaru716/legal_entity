package legalentities

import (
	"time"

	"github.com/google/uuid"
)

type LegalEntity struct {
	UUID      uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();not null:false;primary_key:true"`
	Name      string     `gorm:"type:varchar(255);default:'';not null"`
	CreatedAt time.Time  `gorm:"type:timestamptz;default:now();not null"`
	UpdatedAt time.Time  `gorm:"type:timestamptz;default:now();not null"`
	DeletedAt *time.Time `gorm:"type:timestamptz;default:NULL;"`
}

func (LegalEntity) TableName() string { return "legal_entities" }
