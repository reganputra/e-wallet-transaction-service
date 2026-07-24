package external

import (
	"context"
	"e-wallet-transaction-service/constant"
	"e-wallet-transaction-service/external/proto/notification"
	"e-wallet-transaction-service/helpers"
	"errors"
	"fmt"

	"google.golang.org/grpc"
)

func (*External) SendNotification(ctx context.Context, recipient, templateName string, placeholder map[string]string) error {
	conn, err := grpc.Dial(helpers.GetEnv("NOTIFICATION_GRPC_HOST", ""), grpc.WithInsecure())
	if err != nil {
		return errors.New("failed to dial notification service")
	}
	defer conn.Close()

	client := notification.NewNotificationServiceClient(conn)
	request := notification.SendNotificationRequest{
		Recipient:    recipient,
		TemplateName: templateName,
		Placeholders: placeholder,
	}

	resp, err := client.SendNotification(ctx, &request)
	if err != nil {
		return err
	}
	if resp.Message != constant.Success {
		return fmt.Errorf("get response error from notification: %s", resp.Message)
	}
	return nil
}
