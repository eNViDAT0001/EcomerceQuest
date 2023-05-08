package paypal

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_viper"
	"github.com/plutov/paypal/v4"
	"os"
)

var (
	vp          = wrap_viper.GetViper()
	client      *paypal.Client
	accessToken *paypal.TokenResponse
)

func GetClient() (cl *paypal.Client, err error) {
	if client == nil {
		client, err = newClient()
		if err != nil {
			return nil, err
		}
	}
	accessToken, err = client.GetAccessToken(context.Background())
	return client, err
}
func newClient() (*paypal.Client, error) {
	// Create a client instance
	clientKey := vp.GetString("PAYPAL.CLIENT_KEY")
	secretKey := vp.GetString("PAYPAL.SECRET_KEY")
	c, err := paypal.NewClient(clientKey, secretKey, paypal.APIBaseSandBox)
	c.SetLog(os.Stdout) // Set log to terminal stdout

	return c, err
}
