package dto

import "github.com/google/uuid"

type PurchasesRequest struct {
	PackageID uuid.UUID `json:"package_id" binding:"required"`
}
