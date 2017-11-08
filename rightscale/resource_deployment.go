package rightscale

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mitchellh/mapstructure"
	"reflect"
)

func resourceDeployment() *schema.Resource {
	return &schema.Resource{
		Schema: resourceSchema(reflect.TypeOf(Deployment{})),
		Create: resourceDeploymentCreate,
		Read:   resourceDeploymentRead,
		Update: resourceDeploymentUpdate,
		Delete: resourceDelete,
		Exists: resourceExists,
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
	resource, err := client.resourceRead(d.Id())
	if err != nil {
		return err
	}

	var deployment Deployment
	err = mapstructure.Decode(resource, &deployment)
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

	return err
}
