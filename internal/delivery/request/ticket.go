package request

import (
	"github.com/google/uuid"
	"go-gorm/entity"
	"go-gorm/pkg/enum"
	"time"
)

type TicketRequest struct {
	ID       string            `json:"id" validate:"omitempty,uuid"`
	Metadata string            `json:"metadata" validate:"omitempty"`
	RoomID   string            `json:"roomId" validate:"required,uuid"`
	IssuedBy string            `json:"issuedBy" validate:"omitempty,uuid"`
	Assigned string            `json:"assigned" validate:"omitempty,uuid"`
	Status   enum.TicketStatus `json:"status" validate:"omitempty,oneof=todo on-progress done"`
}

type TicketFilterRequest struct {
	RoomID     string            `json:"roomId" query:"roomid" validate:"required,uuid"`
	Status     enum.TicketStatus `json:"status" query:"status" validate:"omitempty,oneof=todo on-progress done default=todo"`
	IssuedBy   string            `json:"issuedBy" query:"issuedBy" validate:"omitempty,uuid"`
	Assigned   string            `json:"assigned" query:"assigned" validate:"omitempty,uuid"`
	TimeFilter time.Time         `json:"timeFilter" query:"timeFilter" validate:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
	Page       int               `json:"page" query:"page" validate:"omitempty,gte=1 default=1"`
	PageSize   int               `json:"pageSize" query:"pageSize" validate:"omitempty,gte=1,lte=100 default=10"`
}

func (t *TicketFilterRequest) ToTicketFilter(userId uuid.UUID) *entity.TicketFilter {
	var roomID uuid.UUID
	if t.RoomID != "" {
		roomID = uuid.Must(uuid.Parse(t.RoomID))
	}

	var issuedBy uuid.UUID
	if t.IssuedBy != "" {
		issuedBy = uuid.Must(uuid.Parse(t.IssuedBy))
	} else {
		issuedBy = userId
	}

	var assigned uuid.UUID
	if t.Assigned != "" {
		assigned = uuid.Must(uuid.Parse(t.Assigned))
	}

	if t.Page == 0 {
		t.Page = 1
	}
	if t.PageSize == 0 {
		t.PageSize = 10
	}

	return &entity.TicketFilter{
		RoomID:     roomID,
		Status:     t.Status,
		IssuedBy:   issuedBy,
		Assigned:   assigned,
		TimeFilter: t.TimeFilter,
		Page:       t.Page,
		PageSize:   t.PageSize,
	}
}

func (t *TicketRequest) ToDto() *entity.TicketDTO {
	var id uuid.UUID
	if t.ID != "" {
		id = uuid.Must(uuid.Parse(t.ID))
	}

	var issuedBy uuid.UUID
	if t.IssuedBy != "" {
		issuedBy = uuid.Must(uuid.Parse(t.IssuedBy))
	}

	var roomID uuid.UUID
	if t.RoomID != "" {
		roomID = uuid.Must(uuid.Parse(t.RoomID))
	}

	var assigned uuid.UUID
	if t.Assigned != "" {
		assigned = uuid.Must(uuid.Parse(t.Assigned))
	}

	return &entity.TicketDTO{
		ID:       id,
		Metadata: t.Metadata,
		RoomID:   roomID,
		IssuedBy: issuedBy,
		Assigned: assigned,
		Status:   t.Status,
	}
}
