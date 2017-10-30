package rightscale

import (
	"time"
)

type Resource struct {
	Namespace string      `json:"namespace"`
	Type      string      `json:"type"`
	Fields    interface{} `json:"fields"`
}

type Deployment struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProviderConfiguration struct {
	client        *RsClient
	accountNumber int
	apiHostname   string
}

type ProcessMedia struct {
	Id           string                       `json:"id"`
	Href         string                       `json:"href"`
	Name         string                       `json:"name"`
	Source       string                       `json:"source"`
	Main         string                       `json:"main"`
	Parameters   []ParameterMedia             `json:"parameters"`
	Application  string                       `json:"application"`
	Status       string                       `json:"status"`
	References   []ReferenceMedia             `json:"references"`
	Variables    []VariableMedia              `json:"variables"`
	Outputs      []OutputMedia                `json:"outputs"`
	Tasks        []TaskMedia                  `json:"tasks"`
	CreatedBy    UserMedia                    `json:"created_by"`
	CreatedAt    time.Time                    `json:"created_at"`
	UpdatedAt    time.Time                    `json:"updated_at"`
	FinishedAt   time.Time                    `json:"finished_at"`
	FinishReason string                       `json:"finish_reason"`
	Links        map[string]map[string]string `json:"links"`
}

type UserMedia struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type OutputMedia struct {
	Name  string         `json:"name"`
	Value ParameterMedia `json:"value"`
}

type TaskMedia struct {
	Id               string                       `json:"id"`
	Href             string                       `json:"href"`
	Name             string                       `json:"name"`
	Label            string                       `json:"label"`
	Process          *ProcessMedia                `json:"process"`
	ParentExpression ExpressionMedia              `json:"parent_expression"`
	Expressions      []ExpressionMedia            `json:"expressions"`
	Progress         ProgressMedia                `json:"progress"`
	Sequence         int                          `json:"sequence"`
	Callstack        []string                     `json:"callstack"`
	Status           string                       `json:"status"`
	Activity         ActivityMedia                `json:"activity"`
	Error            *ErrorMedia                  `json:"error"`
	CreatedAt        time.Time                    `json:"created_at"`
	UpdatedAt        time.Time                    `json:"updated_at"`
	FinishedAt       time.Time                    `json:"finished_at"`
	Links            map[string]map[string]string `json:"links"`
}

type ExpressionMedia struct {
	Id               string           `json:"id"`
	Href             string           `json:"href"`
	Task             *TaskMedia       `json:"task"`
	Source           string           `json:"source"`
	JsonResult       JsonMedia        `json:"json_result"`
	CollectionResult CollectionMedia  `json:"collection_result"`
	Variables        []VariableMedia  `json:"variables"`
	References       []ReferenceMedia `json:"references"`
	Error            *ErrorMedia      `json:"error"`
}

type ProgressMedia struct {
	Percent    int             `json:"percent"`
	Summary    string          `json:"summary"`
	Expression ExpressionMedia `json:"expression"`
}

type ParameterMedia struct {
	Kind  string      `json:"kind"`
	Value interface{} `json:"value"`
}

type ReferenceMedia struct {
	Id         string          `json:"id"`
	Href       string          `json:"href"`
	Name       string          `json:"name"`
	Value      CollectionMedia `json:"value"`
	Expression ExpressionMedia `json:"expression"`
	Process    ProcessMedia    `json:"process"`
}

type VariableMedia struct {
	Id         string          `json:"id"`
	Href       string          `json:"href"`
	Name       string          `json:"name"`
	Value      JsonMedia       `json:"value"`
	Expression ExpressionMedia `json:"expression"`
	Process    ProcessMedia    `json:"process"`
}

type ActivityMedia struct {
	Id          string                       `json:"id"`
	Href        string                       `json:"href"`
	User        string                       `json:"user"`
	Application string                       `json:"application"`
	Trace       bool                         `json:"trace"`
	Resources   CollectionMedia              `json:"resources"`
	Action      ActionMedia                  `json:"action"`
	State       map[string]string            `json:"state"`
	Links       map[string]map[string]string `json:"links"`
}

type ErrorMedia struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type JsonMedia struct {
	Kind  string      `json:"kind"`
	Value interface{} `json:"value"`
}

type CollectionMedia struct {
	Kind  string `json:"kind"`
	Value struct {
		Namespace string                   `json:"namespace"`
		Type      string                   `json:"type"`
		Hrefs     []string                 `json:"hrefs"`
		Details   []map[string]interface{} `json:"details"`
	} `json:"value"`
}

type ActionMedia struct {
	Name      string            `json:"name"`
	Arguments map[string]string `json:"arguments"`
}
