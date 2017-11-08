package rightscale

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"reflect"
	"strconv"
	"strings"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"refresh_token": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("RS_REFRESH_TOKEN", nil),
				Description: "RightScale refresh token for OAuth API authentication",
			},
			"api_hostname": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("RS_HOSTNAME", nil),
				Description: "RightScale hostname for API, e.g. us-3.rightscale.com",
			},
			"account_number": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("RS_ACCOUNT_NUMBER", nil),
				Description: "RightScale account number for this infrastructure",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"rightscale_deployment": resourceDeployment(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(r *schema.ResourceData) (interface{}, error) {
	apiHostname := r.Get("api_hostname").(string)
	refreshToken := r.Get("refresh_token").(string)
	accountNumberString := r.Get("account_number").(string)
	accountNumber, err := strconv.Atoi(accountNumberString)
	if err != nil {
		return nil, err
	}

	rsc := &RsClient{
		ApiHostname:   apiHostname,
		RefreshToken:  refreshToken,
		AccountNumber: accountNumber,
	}

	err = rsc.authenticate()
	if err != nil {
		return nil, err
	}

	return ProviderConfiguration{
		client:        rsc,
		accountNumber: accountNumber,
		apiHostname:   apiHostname,
	}, nil
}

func resourceSchema(resourceType reflect.Type) map[string]*schema.Schema {
	result := make(map[string]*schema.Schema)

	for i := 0; i < resourceType.NumField(); i++ {
		var fieldSchema schema.Schema
		field := resourceType.Field(i)
		attributeName := strings.Split(field.Tag.Get("json"), ",")[0]
		attributeRequired := field.Tag.Get("required")

		switch attributeType := field.Type.Name(); attributeType {
		case "string":
			fieldSchema.Type = schema.TypeString
		default:
			panic(fmt.Sprintf("unknown type: %v", attributeType))
		}

		if attributeRequired == "true" {
			fieldSchema.Required = true
		} else {
			fieldSchema.Optional = true
		}

		result[attributeName] = &fieldSchema
	}

	return result
}

func resourceDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(ProviderConfiguration).client
	err := client.resourceDelete(d.Id())
	if err != nil {
		return err
	}
	return nil
}

func resourceExists(d *schema.ResourceData, m interface{}) (bool, error) {
	config := m.(ProviderConfiguration)
	client := config.client
	exists, err := client.resourceExists(d.Id(), config.accountNumber, config.apiHostname)
	if err != nil {
		return false, fmt.Errorf("could not check existence of %s: %+v", d.Id(), err)
	}
	return exists, nil
}
