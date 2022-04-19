package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/svanellewee/todoclient/client/operations"
	"github.com/svanellewee/todoclient/models"
	//"log"
)

func resourceTask() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider scaffolding.",

		CreateContext: resourceTaskCreate,
		ReadContext:   resourceTaskRead,
		UpdateContext: resourceTaskUpdate,
		DeleteContext: resourceTaskDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				// This description is used by the documentation generator and the language server.
				Description: "Sample attribute.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func resourceTaskCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	client := meta.(*apiClient)
	params := operations.NewPostTaskParams()
	tflog.Debug(ctx, fmt.Sprintf("..........................%#v\n", d))
	tflog.Debug(ctx, fmt.Sprintf("..........................%#v\n", d.Get("description")))
	var (
		description string
		isString    bool
	)

	if description, isString = d.Get("description").(string); !isString {
		return diag.Errorf("could not get description") //, a ...interface{})
	}
	//description, ok := d.State().Attributes["description"]
	//if !ok {
	//	log.Fatal("could nto cast description")
	//}
	//fmt.Println(">>>>>>>>>>>>>>>>>", description)
	//params.Body.Description = swag.String(description)
	params.Body = &models.Task{}
	params.Body.Description = swag.String(description)
	result, err := client.todoClient.Operations.PostTask(params)
	if err != nil {
		return diag.FromErr(err)
	}
	//client.client.Operations.PostTask
	d.SetId(fmt.Sprintf("%d", result.Payload.ID))

	// write logs using the tflog package
	// see https://pkg.go.dev/github.com/hashicorp/terraform-plugin-log/tflog
	// for more information
	tflog.Trace(ctx, "created a resource")
	resourceTaskRead(ctx, d, meta)
	return diag.Diagnostics{} //diag.Errorf("not implemented")
}

func resourceTaskRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*apiClient)

	params := operations.NewGetTaskIDParams()
	var diags diag.Diagnostics
	var err error
	params.ID, err = strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	result, err := client.todoClient.Operations.GetTaskID(params)
	if err != nil {
		return diag.FromErr(err)
	}

	//d.Set("completed", swag.Bool(result.Payload.Completed))
	d.Set("description", result.Payload.Description)
	//d.Set("id", swag.Int64(result.Payload.ID))
	tflog.Debug(ctx, "BLAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	return diags
}

func resourceTaskUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceTaskDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}
