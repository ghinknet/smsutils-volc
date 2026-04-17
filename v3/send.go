package volc

import (
	"github.com/ghinknet/smsutils/v3/errors"
	"github.com/ghinknet/smsutils/v3/model"
	"github.com/ghinknet/smsutils/v3/utils"
	"github.com/volcengine/volc-sdk-golang/service/sms"
)

func (c Client) SendMessage(dest string, sender string, template string, vars model.Vars) error {
	// Try to parse number
	to, _, _, _, err := utils.ProcessNumberForChinese(dest)

	// Preprocess vars
	params := make(map[string]string)
	for _, v := range vars {
		params[v.Key] = v.Value
	}

	// Marshal params
	marshalledParam, err := c.Marshal(params)
	if err != nil {
		return err
	}

	// Construct request
	req := &sms.SmsRequest{
		SmsAccount:    c.SmsAccount,
		Sign:          sender,
		TemplateID:    template,
		TemplateParam: string(marshalledParam),
		PhoneNumbers:  to,
	}

	// Send request
	result, statusCode, err := sms.DefaultInstance.Send(req)
	if err != nil {
		return err
	}
	if statusCode != 200 {
		if result.ResponseMetadata.Error != nil {
			return errors.ErrDriverSendFailed.
				WithDriverName(DriverName).
				WithDriverCode(statusCode).
				WithDriverMessage(result.ResponseMetadata.Error.Message).
				WithDriverRequestID(result.ResponseMetadata.RequestId).
				WithDriverResponse(result)
		}
		return errors.ErrDriverSendFailed.
			WithDriverName(DriverName).
			WithDriverCode(statusCode).
			WithDriverRequestID(result.ResponseMetadata.RequestId).
			WithDriverResponse(result)
	}

	return nil
}
