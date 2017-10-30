package rightscale

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDeployment() *schema.Resource {
	return &schema.Resource{
		Create: resourceDeploymentCreate,
		Read:   resourceDeploymentRead,
		Update: resourceDeploymentUpdate,
		Delete: resourceDeploymentDelete,
		Exists: resourceDeploymentExists,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceDeploymentCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(ProviderConfiguration).client
	resource := Resource{
		Namespace: "rs_cm",
		Type:      "deployments",
		Fields: Deployment{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
		},
	}

	resourceHref, err := client.resourceCreate(&resource)
	if err != nil {
		return err
	}

	d.SetId(resourceHref)
	return nil
}

func resourceDeploymentRead(d *schema.ResourceData, m interface{}) error {
	client := m.(ProviderConfiguration).client
	deployment, err := client.deploymentRead(d.Id())
	if err != nil {
		return err
	}

	d.Set("name", deployment.Name)
	d.Set("description", deployment.Description)
	return nil
}

func resourceDeploymentUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(ProviderConfiguration).client

	resource := &Resource{
		Namespace: "rs_cm",
		Type:      "deployments",
		Fields: Deployment{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
		},
	}

	err := client.resourceUpdate(d.Id(), resource)
	if err != nil {
		return err
	}
	return nil
}

func resourceDeploymentDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(ProviderConfiguration).client
	err := client.resourceDelete(d.Id())
	if err != nil {
		return err
	}
	return nil
}

func resourceDeploymentExists(d *schema.ResourceData, m interface{}) (bool, error) {
	config := m.(ProviderConfiguration)
	client := config.client
	exists, err := client.resourceExists(d.Id(), config.accountNumber, config.apiHostname)
	if err != nil {
		return false, fmt.Errorf("could not check existence of %s: %+v", d.Id(), err)
	}
	return exists, nil
}
