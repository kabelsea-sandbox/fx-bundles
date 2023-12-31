module github.com/kabelsea-sandbox/fx-bundles

go 1.21.5

replace (
	github.com/aws/aws-sdk-go-v2 => github.com/aws/aws-sdk-go-v2 v1.22.1
	github.com/aws/aws-sdk-go-v2/config => github.com/aws/aws-sdk-go-v2/config v1.20.0
	github.com/aws/aws-sdk-go-v2/credentials => github.com/aws/aws-sdk-go-v2/credentials v1.14.0
	github.com/aws/aws-sdk-go-v2/service/sqs => github.com/aws/aws-sdk-go-v2/service/sqs v1.26.0
)

require (
	// We cannot use the latest version of the aws-sdk-go-v2 because it BREAKS the compatibility with the YMQ.
	// https://github.com/aws/aws-sdk-go-v2/issues/2370
	github.com/aws/aws-sdk-go-v2 v1.22.1
	github.com/aws/aws-sdk-go-v2/config v1.20.0
	github.com/aws/aws-sdk-go-v2/credentials v1.14.0
	// We cannot use the latest version of the sqs because it uses new JSON protocol v1 which is not supported
	// by YMQ. Latest version of the sqs which uses query protocol is v1.26.0. And as it is version from 2023-11-01
	// it is not compatible with the aws-sdk-go-v2 greater than v1.22.2.
	github.com/aws/aws-sdk-go-v2/service/sqs v1.26.0
)

require (
	github.com/TheZeroSlave/zapsentry v1.20.2
	github.com/getsentry/sentry-go v0.25.0
	github.com/go-playground/validator/v10 v10.16.0
	github.com/go-resty/resty/v2 v2.11.0
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.2-0.20221020003552-4126fa611266
	github.com/json-iterator/go v1.1.12
	github.com/labstack/echo-contrib v0.15.0
	github.com/labstack/echo/v4 v4.11.4
	github.com/mcuadros/go-defaults v1.2.0
	github.com/spf13/viper v1.18.2
	github.com/ydb-platform/ydb-go-sdk/v3 v3.54.3
	go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho v0.46.1
	go.opentelemetry.io/otel v1.21.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.21.0
	go.opentelemetry.io/otel/exporters/prometheus v0.44.0
	go.opentelemetry.io/otel/sdk v1.21.0
	go.opentelemetry.io/otel/sdk/metric v1.21.0
	go.uber.org/fx v1.20.1
	go.uber.org/zap v1.26.0
	golang.org/x/exp v0.0.0-20240103183307-be819d1f06fc
)

require (
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.14.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.2.1 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.5.1 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.4.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.10.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.16.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.18.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.24.0 // indirect
	github.com/aws/smithy-go v1.16.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jonboulle/clockwork v0.4.0 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.1.1 // indirect
	github.com/prometheus/client_golang v1.18.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.45.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/ydb-platform/ydb-go-genproto v0.0.0-20231215113745-46f6d30f974a // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.21.0 // indirect
	go.opentelemetry.io/otel/metric v1.21.0 // indirect
	go.opentelemetry.io/otel/trace v1.21.0 // indirect
	go.opentelemetry.io/proto/otlp v1.0.0 // indirect
	go.uber.org/dig v1.17.1 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240102182953-50ed04b92917 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240102182953-50ed04b92917 // indirect
	google.golang.org/grpc v1.60.1 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
