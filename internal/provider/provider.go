package provider

import (
	"context"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	apiclient "github.com/svanellewee/todoclient/client"
	//"github.com/svanellewee/todoclient/client/operations"
	//"github.com/svanellewee/todoclient/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// postParams := operations.NewPostTaskParams()
	// description := "hello"
	// postParams.Body = &models.Task{
	// 	Description: &description,
	// }

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			DataSourcesMap: map[string]*schema.Resource{
				"task": dataSourceTask(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"task": resourceTask(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

type apiClient struct {
	// Add whatever fields, client or connection info, etc. here
	// you would need to setup to communicate with the upstream
	// API.
	todoClient *apiclient.TodoServer
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
		var diags diag.Diagnostics
		// Setup a User-Agent for your API client (replace the provider name for yours):
		// userAgent := p.UserAgent("terraform-provider-scaffolding", version)
		// TODO: myClient.UserAgent = userAgent
		transport := httptransport.New("127.0.0.1:8080", "", []string{"http"})
		client := apiclient.New(transport, strfmt.Default)

		return &apiClient{
			todoClient: client,
		}, diags
	}
}
