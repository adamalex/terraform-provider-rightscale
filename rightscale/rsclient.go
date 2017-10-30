package rightscale

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/rightscale/rsc.v6/rsapi"
	"gopkg.in/rightscale/rsc.v6/ss"
	"net/http"
	"strconv"
	"strings"
)

type RsClient struct {
	ApiHostname   string
	RefreshToken  string
	AccountNumber int
	API           *ss.API
}

func (rsc *RsClient) deploymentRead(href string) (*Deployment, error) {
	source := fmt.Sprintf(`define main() return $resource do
		@resource = rs_cm.get(href: "%s")
		$resource = {
			name: @resource.name,
			description: @resource.description
		}
	end`, href)

	process, err := rsc.processExecuteUntilComplete(source)
	if err != nil {
		return nil, err
	}

	output := process.singleOutput().(map[string]interface{})

	return &Deployment{
		Name:        output["name"].(map[string]interface{})["value"].(string),
		Description: output["description"].(map[string]interface{})["value"].(string),
	}, nil
}

func (rsc *RsClient) resourceExists(href string, account int, hostName string) (bool, error) {
	source := fmt.Sprintf(`define main() return $code do
		$rs_endpoint = "https://%s"

		$response = http_get(
			url: $rs_endpoint+"%s",
			headers: {
				"X-Api-Version": "1.6",
				"X-Account": "%s"
			}
		)

		$code = $response["code"]
	end`, hostName, href, strconv.Itoa(account))

	process, err := rsc.processExecuteUntilComplete(source)
	if err != nil {
		return false, err
	}

	code := process.singleOutput().(float64)

	switch code {
	case http.StatusOK:
		return true, nil
	case http.StatusNotFound:
		return false, nil
	default:
		return false, fmt.Errorf("unexpected response code: %v", code)
	}
}

func (rsc *RsClient) resourceUpdate(href string, resource *Resource) error {
	resourceFields, err := json.Marshal((*resource).Fields)
	if err != nil {
		return err
	}

	source := fmt.Sprintf(`define main() return $href do
		@resource = rs_cm.get(href: "%s")
		@resource.update("%s": %s)
		$href = @resource.href
	end`, href, strings.TrimSuffix((*resource).Type, "s"), resourceFields)

	_, err = rsc.processExecuteUntilComplete(source)
	if err != nil {
		return err
	}

	return nil
}

func (rsc *RsClient) resourceCreate(resource *Resource) (string, error) {
	resourceJson, err := json.Marshal(resource)
	if err != nil {
		return "", err
	}

	source := fmt.Sprintf(`define main() return $href do
		@new_resource = %s
		provision(@new_resource)
		$href = @new_resource.href
	end`, resourceJson)

	process, err := rsc.processExecuteUntilComplete(source)
	if err != nil {
		return "", err
	}

	return process.singleOutput().(string), nil
}

func (rsc *RsClient) resourceDelete(href string) error {
	source := fmt.Sprintf(`define main() return $href do
		@resource = rs_cm.get(href: "%s")
		delete(@resource)
		$href = ""
	end`, href)

	_, err := rsc.processExecuteUntilComplete(source)
	if err != nil {
		return err
	}

	return nil
}

func (rsc *RsClient) processRead(id, view string) (*ProcessMedia, error) {
	res, err := rsc.requestCreate(
		"get",
		"/cwf/v1/accounts/"+rsc.accountId()+"/processes/"+id,
		rsapi.APIParams{
			"view": view,
		},
		rsapi.APIParams{},
	)
	if err != nil {
		return nil, err
	}

	var process ProcessMedia
	err = mapstructure.Decode(res, &process)
	if err != nil {
		return nil, err
	}

	return &process, nil
}

func (rsc *RsClient) processExecuteUntilComplete(source string) (*ProcessMedia, error) {
	var process *ProcessMedia

	payload := rsapi.APIParams{
		"source":      source,
		"main":        "main",
		"rcl_version": "2",
		"parameters":  nil,
		"application": "cwfconsole",
		"created_by": map[string]interface{}{
			"id":    0,
			"name":  "Terraform",
			"email": "support@rightscale.com",
		},
		"refresh_token": rsc.RefreshToken,
	}
	res, err := rsc.requestCreate("post", "/cwf/v1/accounts/"+rsc.accountId()+"/processes", nil, payload)
	if err != nil {
		return nil, err
	}

	processHref := res.(map[string]interface{})["Location"]
	processHrefParts := strings.Split(processHref.(string), "/")
	processId := processHrefParts[len(processHrefParts)-1]

	for {
		process, err = rsc.processRead(processId, "expanded")
		if err != nil {
			return nil, err
		}
		status := process.Status
		if status != "running" {
			if status == "completed" {
				return process, nil
			} else {
				return nil, errors.New(fmt.Sprintf("cwf process status: %s\n%s", status, process.Tasks[0].Error.Message))
			}
		}
	}
}

func (rsc *RsClient) requestCreate(method, url string, params, payload rsapi.APIParams) (interface{}, error) {
	req, err := rsc.API.BuildHTTPRequest(
		strings.ToUpper(method),
		url,
		"1.0",
		params,
		payload,
	)
	if err != nil {
		return nil, err
	}

	req.Host = rsc.cwfHostname()

	res, err := rsc.API.PerformRequest(req)
	if err != nil {
		return nil, err
	}

	resp, err := rsc.API.LoadResponse(res)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (rsc *RsClient) authenticate() error {
	auth := rsapi.NewOAuthAuthenticator(rsc.RefreshToken, rsc.AccountNumber)
	rsc.API = ss.New(rsc.ApiHostname, auth)
	return rsc.API.CanAuthenticate()
}

func (rsc *RsClient) cwfHostname() string {
	return strings.Replace(rsc.ApiHostname, "us-", "cloud-workflow", 1)
}

func (rsc *RsClient) accountId() string {
	return strconv.Itoa(rsc.AccountNumber)
}

func (p *ProcessMedia) singleOutput() interface{} {
	return p.Outputs[0].Value.Value
}
