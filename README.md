# smsutils-volc

Volcengine (VolcEngine) SMS driver for [smsutils](https://go.gh.ink/smsutils).

This module implements the `model.Driver` / `model.Client` interfaces of the core
`smsutils` library on top of the Volcengine `volc-sdk-golang` SMS service.

Driver name: `volc`

## Installation

```bash
go get go.gh.ink/smsutils/volc/v3
```

## Usage

Blank-import this package so it registers itself, then configure it via the core client.

```go
import (
	"go.gh.ink/smsutils/v3/client"
	"go.gh.ink/smsutils/v3/model"

	_ "go.gh.ink/smsutils/volc/v3"
)

clients, err := client.NewClient(model.Config{
	Credentials: model.C{
		"volc": {
			"accessKey":  "<your-access-key>",
			"secretKey":  "<your-secret-key>",
			"smsAccount": "<your-sms-account>",
		},
	},
})
if err != nil {
	panic(err)
}

// "sender" is the Sign; "ST_..." is the TemplateID.
err = clients["volc"].SendMessage("+8617601205205", "<sign>", "ST_12345", model.Vars{
	{Key: "code", Value: "1234"},
})
```

## Credentials

| Key | Constant | Required | Description |
|-----|----------|----------|-------------|
| `accessKey` | `AccessKey` | Yes | Volcengine Access Key |
| `secretKey` | `SecretKey` | Yes | Volcengine Secret Key |
| `smsAccount` | `SmsAccount` | Yes (per send) | Volcengine SMS account ID |

Missing `accessKey` or `secretKey` yields `errors.ErrDriverCredentialInvalid`.

## Behavior

- The destination is normalized via `utils.ProcessNumberForChinese` before sending.
- `template` maps to the Volcengine `TemplateID`; `sender` maps to `Sign`.
- `vars` are passed **by key**: they are flattened to a `map[string]string`, JSON-marshalled
  (using the configured `Marshal`), and sent as `TemplateParam`. Each `Var.Key` must match a
  placeholder in your template.
- A non-`200` HTTP status code or a non-nil `ResponseMetadata.Error` produces
  `errors.ErrDriverSendFailed`, decorated with the status code, provider message, request ID
  and raw result.

> **Note:** the underlying Volcengine SDK uses a process-global `DefaultInstance`; the access
> key and secret key are set on that shared instance when a client is created. This makes the
> driver effectively single-credential per process — constructing multiple `volc` clients with
> different keys is not supported.

## Requirements

- Go 1.25.0+
- `go.gh.ink/smsutils/v3`
- Volcengine SDK: `github.com/volcengine/volc-sdk-golang`

## License

[Apache License 2.0](LICENSE)
