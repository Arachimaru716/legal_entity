package web

import (
	"context"

	"github.com/krisch/crm-backend/dto"
	oapi "github.com/krisch/crm-backend/internal/web/ofederation"
)

// GET /legal_entity.
func (a *Web) GetLegalEntity(ctx context.Context, _ oapi.GetLegalEntityRequestObject) (oapi.GetLegalEntityResponseObject, error) {
	items, err := a.app.LegalEntitiesService.GetLegalEntities(ctx)
	if err != nil {
		return nil, err
	}

	itemsDTO := make([]dto.LegalEntityDTO, 0, len(items))
	for _, it := range items {
		itemsDTO = append(itemsDTO, dto.LegalEntityDTO{
			UUID:      it.UUID,
			Name:      it.Name,
			CreatedAt: it.CreatedAt,
			UpdatedAt: it.UpdatedAt,
		})
	}

	return oapi.GetLegalEntity200JSONResponse{
		Count: len(itemsDTO),
		Items: itemsDTO,
	}, nil
}

// POST /legal_entity.
func (a *Web) PostLegalEntity(ctx context.Context, req oapi.PostLegalEntityRequestObject) (oapi.PostLegalEntityResponseObject, error) {
	entity, err := a.app.LegalEntitiesService.CreateLegalEntity(req.Body.Name)
	if err != nil {
		return nil, err
	}
	return oapi.PostLegalEntity200JSONResponse{Uuid: entity.UUID}, nil
}

// PATCH /legal_entity/{UUID}/name.
func (a *Web) PatchLegalEntityUUIDName(ctx context.Context, req oapi.PatchLegalEntityUUIDNameRequestObject) (oapi.PatchLegalEntityUUIDNameResponseObject, error) {
	if err := a.app.LegalEntitiesService.UpdateLegalEntityName(ctx, req.UUID, req.Body.Name); err != nil {
		return nil, err
	}
	return oapi.PatchLegalEntityUUIDName200Response{}, nil
}

// DELETE /legal_entity/{UUID}.
func (a *Web) DeleteLegalEntityUUID(ctx context.Context, req oapi.DeleteLegalEntityUUIDRequestObject) (oapi.DeleteLegalEntityUUIDResponseObject, error) {
	if err := a.app.LegalEntitiesService.DeleteLegalEntity(req.UUID); err != nil {
		return nil, err
	}
	return oapi.DeleteLegalEntityUUID200Response{}, nil
}
