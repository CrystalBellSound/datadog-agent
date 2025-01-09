module github.com/DataDog/datadog-agent/pkg/trace/stats/oteltest

go 1.22.0

require (
	github.com/DataDog/datadog-agent/comp/otelcol/otlp/components/statsprocessor v0.56.0-rc.3
	github.com/DataDog/datadog-agent/pkg/proto v0.56.0-rc.3
	github.com/DataDog/datadog-agent/pkg/trace v0.56.0-rc.3
	github.com/DataDog/datadog-go/v5 v5.6.0
	github.com/DataDog/opentelemetry-mapping-go/pkg/otlp/attributes v0.22.0
	github.com/google/go-cmp v0.6.0
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/component/componenttest v0.115.0
	go.opentelemetry.io/collector/pdata v1.21.0
	go.opentelemetry.io/collector/semconv v0.115.0
	go.opentelemetry.io/otel/metric v1.33.0
	google.golang.org/protobuf v1.36.1
)

require go.opentelemetry.io/collector/component v0.115.0 // indirect

require (
	github.com/DataDog/datadog-agent/comp/core/tagger/origindetection v0.0.0-20241217122454-175edb6c74f2 // indirect
	github.com/DataDog/datadog-agent/comp/trace/compression/def v0.56.0-rc.3 // indirect
	github.com/DataDog/datadog-agent/comp/trace/compression/impl-gzip v0.56.0-rc.3 // indirect
	github.com/DataDog/datadog-agent/pkg/obfuscate v0.56.0-rc.3 // indirect
	github.com/DataDog/datadog-agent/pkg/remoteconfig/state v0.56.0-rc.3 // indirect
	github.com/DataDog/datadog-agent/pkg/util/cgroups v0.56.0-rc.3 // indirect
	github.com/DataDog/datadog-agent/pkg/util/log v0.59.0 // indirect
	github.com/DataDog/datadog-agent/pkg/util/pointer v0.59.0 // indirect
	github.com/DataDog/datadog-agent/pkg/util/scrubber v0.59.0 // indirect
	github.com/DataDog/datadog-agent/pkg/version v0.59.1 // indirect
	github.com/DataDog/go-sqllexer v0.0.19 // indirect
	github.com/DataDog/go-tuf v1.1.0-0.5.2 // indirect
	github.com/DataDog/sketches-go v1.4.6 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575 // indirect
	github.com/containerd/cgroups/v3 v3.0.4 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/ebitengine/purego v0.8.1 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/mock v1.7.0-rc.1 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/karrick/godirwalk v1.17.0 // indirect
	github.com/lufia/plan9stats v0.0.0-20220913051719-115f729f3c8c // indirect
	github.com/moby/sys/userns v0.1.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/opencontainers/runtime-spec v1.2.0 // indirect
	github.com/outcaste-io/ristretto v0.2.3 // indirect
	github.com/philhofer/fwd v1.1.3-0.20240916144458-20a13a1f6b7c // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/power-devops/perfstat v0.0.0-20220216144756-c35f1ee13d7c // indirect
	github.com/secure-systems-lab/go-securesystemslib v0.8.0 // indirect
	github.com/shirou/gopsutil/v4 v4.24.11 // indirect
	github.com/tinylib/msgp v1.2.4 // indirect
	github.com/tklauser/go-sysconf v0.3.14 // indirect
	github.com/tklauser/numcpus v0.8.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.115.0 // indirect
	go.opentelemetry.io/otel v1.33.0 // indirect
	go.opentelemetry.io/otel/sdk v1.33.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/trace v1.33.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/time v0.8.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241230172942-26aa7a208def // indirect
	google.golang.org/grpc v1.69.2 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/DataDog/datadog-agent/comp/core/tagger/origindetection => ../../../../comp/core/tagger/origindetection
	github.com/DataDog/datadog-agent/comp/otelcol/otlp/components/metricsclient => ../../../../comp/otelcol/otlp/components/metricsclient
	github.com/DataDog/datadog-agent/comp/otelcol/otlp/components/statsprocessor => ../../../../comp/otelcol/otlp/components/statsprocessor
	github.com/DataDog/datadog-agent/comp/trace/compression/def => ../../../../comp/trace/compression/def
	github.com/DataDog/datadog-agent/comp/trace/compression/impl-gzip => ../../../../comp/trace/compression/impl-gzip
	github.com/DataDog/datadog-agent/comp/trace/compression/impl-zstd => ../../../../comp/trace/compression/impl-zstd
	github.com/DataDog/datadog-agent/pkg/networkdevice/profile => ../../../networkdevice/profile
	github.com/DataDog/datadog-agent/pkg/obfuscate => ../../../obfuscate
	github.com/DataDog/datadog-agent/pkg/proto => ../../../proto
	github.com/DataDog/datadog-agent/pkg/remoteconfig/state => ../../../remoteconfig/state
	github.com/DataDog/datadog-agent/pkg/security/secl => ../../../security/secl
	github.com/DataDog/datadog-agent/pkg/serializer => ../../../serializer/
	github.com/DataDog/datadog-agent/pkg/status/health => ../../../status/health
	github.com/DataDog/datadog-agent/pkg/tagset => ../../../tagset/
	github.com/DataDog/datadog-agent/pkg/telemetry => ../../../telemetry/
	github.com/DataDog/datadog-agent/pkg/trace => ../../
	github.com/DataDog/datadog-agent/pkg/util/cgroups => ../../../util/cgroups
	github.com/DataDog/datadog-agent/pkg/util/log => ../../../util/log
	github.com/DataDog/datadog-agent/pkg/util/option => ../../../util/option
	github.com/DataDog/datadog-agent/pkg/util/pointer => ../../../util/pointer
	github.com/DataDog/datadog-agent/pkg/util/scrubber => ../../../util/scrubber
)

replace github.com/DataDog/datadog-agent/pkg/version => ../../../version
