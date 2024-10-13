package customer

import "context"

// GetPhoneNumber retrieves the phone number of a customer by their user ID.
func (s *Service) GetPhoneNumber(ctx context.Context, userID uint) (string, error) {
	user, err := s.repo.Get(ctx, userID)
	if err != nil {
		return "", err
	}

	return user.PhoneNumber, nil
}
