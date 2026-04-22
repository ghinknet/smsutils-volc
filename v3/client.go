package volc

import (
	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/sms"
	"go.gh.ink/smsutils/v3/errors"
	"go.gh.ink/smsutils/v3/model"
)

type Client struct {
	Client     *base.Client
	SmsAccount string
	// JSON
	Marshal   func(any) ([]byte, error)
	Unmarshal func([]byte, any) error
}
type Driver struct{}

func (d Driver) NewClient(params model.DriverClientParam) (model.Client, error) {
	// Check credential
	ak, sk := params.Credential[AccessKey], params.Credential[SecretKey]
	if ak == "" || sk == "" {
		return Client{}, errors.ErrDriverCredentialInvalid.WithDriverName(Name)
	}

	// Create volc client
	client := sms.NewInstance().Client

	// Set access key and secret key
	sms.DefaultInstance.Client.SetAccessKey(ak)
	sms.DefaultInstance.Client.SetSecretKey(sk)

	return Client{
		Client:     client,
		SmsAccount: params.Credential[SmsAccount],
		Marshal:    params.Marshal,
		Unmarshal:  params.Unmarshal,
	}, nil
}
