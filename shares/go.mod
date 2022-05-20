module shares

go 1.14

require (
	github.com/axgle/mahonia v0.0.0-20180208002826-3358181d7394
	github.com/chenjiandongx/ginprom v0.0.0-20201217063207-fe11b7f55a35
	github.com/gin-gonic/gin v1.7.7
	github.com/gmsec/goplugins v0.0.0-20220510122305-f8eb57238301
	github.com/gmsec/micro v0.0.0-20211022114905-169485e6fedf
	github.com/go-playground/validator/v10 v10.11.0 // indirect
	github.com/goccy/go-json v0.9.7 // indirect
	github.com/gookit/color v1.5.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/miekg/dns v1.1.49 // indirect
	github.com/mozillazg/go-pinyin v0.19.0
	github.com/prometheus/client_golang v1.11.0
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/asr v1.0.310
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.300
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/nlp v1.0.300
	github.com/ugorji/go v1.2.7 // indirect
	github.com/xxjwxc/ginrpc v0.0.0-20220416151621-d2a2864d29f4
	github.com/xxjwxc/public v0.0.0-20220427083745-2dca406fd512
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/crypto v0.0.0-20220518034528-6f7dac969898 // indirect
	golang.org/x/net v0.0.0-20220520000938-2e3eb7b945c2 // indirect
	golang.org/x/sys v0.0.0-20220519141025-dcacdad47464 // indirect
	golang.org/x/tools v0.1.10 // indirect
	golang.org/x/xerrors v0.0.0-20220517211312-f3a8303e98df // indirect
	google.golang.org/genproto v0.0.0-20220519153652-3a47de7e79bd // indirect
	google.golang.org/grpc v1.46.2
	google.golang.org/protobuf v1.28.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	gorm.io/datatypes v1.0.1
	gorm.io/gorm v1.21.15
	rpc v0.0.0-00010101000000-000000000000

)

replace rpc => ../apidoc/rpc/

replace google.golang.org/grpc v1.39.0 => google.golang.org/grpc v1.29.1

// replace github.com/xxjwxc/public => ../public

// replace github.com/xxjwxc/ginrpc => ../../ginrpc
// replace github.com/gmsec/goplugins => ../../goplugins

// replace github.com/gmsec/micro => ../../micro
