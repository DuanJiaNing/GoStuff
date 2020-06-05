Updating service [default]...failed.
ERROR: (gcloud.app.deploy) Error Response: [9] Cloud build 45ec40fc-a323-4aca-829b-4e7cbd42b7aa status: FAILURE.
Build error details: {"error":{"errorType":"BuildError","canonicalCode":"INVALID_ARGUMENT","errorId":"E89CAB1E","errorMessage":"2020/05/08 05:36:19 Building /tmp/staging/srv, with main package at .,
saving to /tmp/staging/usr/local/bin/start
2020/05/08 05:36:19 Running \u0026{/usr/local/go/bin/go [go build -o /tmp/staging/usr/local/bin/start .] [PATH=/go/bin:/usr/local/go/bin:/builder/google-cl
oud-sdk/bin/:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin HOSTNAME=7acb7ac85fc1 HOME=/builder/home BUILDER_OUTPUT=/builder/outputs DEBIAN_FRONTEND=noninteractive GOROOT=/usr/local/go/
 GOPATH=/go GO111MODULE=on GOCACHE=/tmp/cache GOPATH=/go] /tmp/staging/srv \u003cnil\u003e \u003cnil\u003e \u003cnil\u003e [] \u003cnil\u003e \u003cnil\u003e \u003cnil\u003e \u003cnil\u003e \u003cnil
\u003e false [] [] [] [] \u003cnil\u003e \u003cnil\u003e}
2020/05/08 05:38:21 Wrote build output to /builder/outputs/output
2020/05/08 05:38:21 Failed to build app: [go build -o /tmp/staging/usr/lo
cal/bin/start .] with env [PATH=/go/bin:/usr/local/go/bin:/builder/google-cloud-sdk/bin/:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin HOSTNAME=7acb7ac85fc1 HOME=/builder/home BUILDER_
OUTPUT=/builder/outputs DEBIAN_FRONTEND=noninteractive GOROOT=/usr/local/go/ GOPATH=/go GO111MODULE=on GOCACHE=/tmp/cache GOPATH=/go] failed: err=exit status 2, out=\"go: finding github.com/onsi/gome
ga v1.9.0\
go: finding github.com/onsi/ginkgo v1.12.0\
go: finding github.com/gorilla/mux v1.7.4\
go: finding github.com/golang/protobuf v1.3.5\
go: finding github.com/golang/mock v1.4.3\
go: fi
nding golang.org/x/net v0.0.0-20200301022130-244492dfa37a\
go: finding google.golang.org/genproto v0.0.0-20200317114155-1f3552e48f24\
go: finding golang.org/x/text v0.3.2\
go: finding cloud.google
.com/go v0.55.0\
go: finding cloud.google.com/go/storage v1.6.0\
go: finding github.com/go-redis/redis v6.15.7+incompatible\
go: finding cloud.google.com/go/datastore v1.1.0\
go: finding github.c
om/hpcloud/tail v1.0.0\
go: finding google.golang.org/api v0.20.0\
go: finding golang.org/x/sys v0.0.0-20191120155948-bd437916bb0e\
go: finding golang.org/x/tools v0.0.0-20190425150028-36563e24a26
2\
go: finding golang.org/x/tools v0.0.0-20190524140312-2c0ae7006135\
go: finding honnef.co/go/tools v0.0.0-20190523083050-ea95bdfd59fc\
go: finding github.com/google/martian v2.1.0+incompatible\\
ngo: finding google.golang.org/grpc v1.27.1\
go: finding golang.org/x/sys v0.0.0-20200212091648-12a6c2dcc1e4\
go: finding golang.org/x/lint v0.0.0-20200302205851-738671d3881b\
go: finding cloud.go
ogle.com/go/logging v1.0.0\
go: finding github.com/envoyproxy/go-control-plane v0.9.1-0.20191026205805-5f8ba28d4473\
go: finding github.com/golang/protobuf v1.3.3\
go: finding cloud.google.com/go
v0.53.0\
go: finding cloud.google.com/go/bigquery v1.5.0\
go: finding google.golang.org/api v0.7.0\
go: finding github.com/google/go-cmp v0.2.0\
go: finding golang.org/x/net v0.0.0-20190311183353
-d8887717615a\
go: finding golang.org/x/tools v0.0.0-20200224181240-023911ca70b2\
go: finding github.com/google/go-cmp v0.3.0\
go: finding github.com/onsi/ginkgo v1.6.0\
go: finding golang.org/x/
sys v0.0.0-20200317113312-5766fd39f98d\
go: finding golang.org/x/net v0.0.0-20200202094626-16171245cfb2\
go: finding google.golang.org/api v0.18.0\
go: finding honnef.co/go/tools v0.0.1-2019.2.3\\
ngo: finding google.golang.org/genproto v0.0.0-20200224152610-e50cd9704f63\
go: finding golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d\
go: finding golang.org/x/oauth2 v0.0.0-20190604053449
-0f29369cfe45\
go: finding cloud.google.com/go v0.38.0\
go: finding golang.org/x/lint v0.0.0-20190409202823-959b441ac422\
go: finding github.com/google/pprof v0.0.0-20200229191704-1ebb73c60ed3\
g
o: finding cloud.google.com/go v0.34.0\
go: finding honnef.co/go/tools v0.0.1-2020.1.3\
go: finding golang.org/x/tools v0.0.0-20200130002326-2f3ba24bd6e7\
go: finding github.com/rogpeppe/go-intern
al v1.3.0\
go: finding honnef.co/go/tools v0.0.0-20190418001031-e561f6794a2a\
go: finding github.com/prometheus/client_model v0.0.0-201.
Check the build log for errors: https://console.cloud.google.com/cloud-build/builds/45ec40fc-a323-4aca-829b-4e7cbd42b7aa?project=230743352788
