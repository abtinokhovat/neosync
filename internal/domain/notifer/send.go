package notifer

import (
	"context"
	"fmt"
)

// SendNotification is a mock notification to show how the flow of the sending notification can happen
// this can be replaced with a real implementation by implementing an adapter for external services like kavenegar and calling it here
func (s Service) SendNotification(ctx context.Context, message string, userID uint) error {
	phoneNumber, err := s.recipient.GetPhoneNumber(ctx, userID)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("sending notfication to %s", phoneNumber))
	fmt.Println(message)

	return nil
}
