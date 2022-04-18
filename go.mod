module github.com/svanellewee/terraform-provider-task-list

go 1.16

replace github.com/svanellewee/todoclient => /root/source/todoclient

require (
	github.com/Masterminds/goutils v1.1.0
	github.com/Masterminds/semver v1.5.0
	github.com/apparentlymart/go-textseg/v13 v13.0.0
	github.com/davecgh/go-spew v1.1.1
	github.com/go-openapi/runtime v0.23.3
	github.com/go-openapi/strfmt v0.21.2
	github.com/go-openapi/swag v0.21.1
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.7
	github.com/google/uuid v1.1.2
	github.com/hashicorp/go-cty v1.4.1-0.20200414143053-d3edf31b6320
	github.com/hashicorp/go-hclog v1.2.0
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hashicorp/go-plugin v1.4.3
	github.com/hashicorp/go-uuid v1.0.2
	github.com/hashicorp/go-version v1.4.0
	github.com/hashicorp/hc-install v0.3.1
	github.com/hashicorp/hcl/v2 v2.11.1
	github.com/hashicorp/logutils v1.0.0
	github.com/hashicorp/terraform-exec v0.16.0
	github.com/hashicorp/terraform-json v0.13.0
	github.com/hashicorp/terraform-plugin-docs v0.7.0
	github.com/hashicorp/terraform-plugin-go v0.8.0
	github.com/hashicorp/terraform-plugin-log v0.3.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.13.0
	github.com/hashicorp/terraform-registry-address v0.0.0-20210412075316-9b2996cce896
	github.com/huandu/xstrings v1.3.2
	github.com/imdario/mergo v0.3.12
	github.com/mattn/go-colorable v0.1.12
	github.com/mattn/go-isatty v0.0.14
	github.com/mitchellh/cli v1.1.2
	github.com/mitchellh/copystructure v1.2.0
	github.com/mitchellh/go-testing-interface v1.14.1
	github.com/mitchellh/mapstructure v1.4.3
	github.com/mitchellh/reflectwalk v1.0.2
	github.com/posener/complete v1.1.1
	github.com/russross/blackfriday v1.6.0
	github.com/svanellewee/todoclient v0.0.0-00010101000000-000000000000
	github.com/vmihailenco/msgpack v4.0.4+incompatible
	github.com/vmihailenco/msgpack/v4 v4.3.12
	github.com/zclconf/go-cty v1.10.0
	golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e
	golang.org/x/net v0.0.0-20220412020605-290c469a71a5
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e
	golang.org/x/text v0.3.7
	google.golang.org/appengine v1.6.6
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.27.1
)
