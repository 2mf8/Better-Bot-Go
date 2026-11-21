package v1

import (
	"context"

	"github.com/2mf8/Better-Bot-Go/dto"
)

// PutInteraction 更新 interaction
func (o *openAPI) PutInteraction(ctx context.Context,
	interactionID string, req *dto.InteractionRespReq) error {
	_, err := o.request(ctx).
		SetPathParam("interaction_id", interactionID).
		SetBody(req).
		Put(o.getURL(interactionsURI))
	return err
}
