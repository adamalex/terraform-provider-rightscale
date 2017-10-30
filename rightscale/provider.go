package rightscale

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"strconv"
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
