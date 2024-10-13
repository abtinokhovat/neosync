package provider

import "context"

func (s Service) GetAll(ctx context.Context) (map[uint]Provider, error) {
	providers, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return providers, nil
}
