OTEL := https://github.com/open-telemetry/opentelemetry-collector-releases/releases/download/
OTELV := v0.115.1/otelcol-contrib_0.115.1_linux_amd64.tar.gz
otel/install:
	wget -O otel.tar.gz $(OTEL)$(OTELV)
	tar zxvf otel.tar.gz

export NGINX_ACCESS_LOG_FILE := /var/log/nginx/access.json
export SIGNOZ_OTEL_ENDPOINT := http://ytr.temama.soy:34317
otel/run:
	./otelcol-contrib --config otel-config.yml
