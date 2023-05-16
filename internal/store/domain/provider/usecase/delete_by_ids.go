package usecase

import "context"

func (u providerUseCase) DeleteProviderByIDs(ctx context.Context, providerID []uint) error {
	return u.providerSto.DeleteProviderByIDs(ctx, providerID)
}
