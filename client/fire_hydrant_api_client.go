// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/client/ai"
	"github.com/firehydrant/api-client-go/client/alerts"
	"github.com/firehydrant/api-client-go/client/bootstrap"
	"github.com/firehydrant/api-client-go/client/catalogs"
	"github.com/firehydrant/api-client-go/client/change_types"
	"github.com/firehydrant/api-client-go/client/changes"
	"github.com/firehydrant/api-client-go/client/checklist_templates"
	"github.com/firehydrant/api-client-go/client/conversations"
	"github.com/firehydrant/api-client-go/client/current_user"
	"github.com/firehydrant/api-client-go/client/custom_fields"
	"github.com/firehydrant/api-client-go/client/entitlements"
	"github.com/firehydrant/api-client-go/client/environments"
	"github.com/firehydrant/api-client-go/client/form_configurations"
	"github.com/firehydrant/api-client-go/client/functionalities"
	"github.com/firehydrant/api-client-go/client/incident_roles"
	"github.com/firehydrant/api-client-go/client/incident_tags"
	"github.com/firehydrant/api-client-go/client/incident_types"
	"github.com/firehydrant/api-client-go/client/incidents"
	"github.com/firehydrant/api-client-go/client/infrastructures"
	"github.com/firehydrant/api-client-go/client/integrations"
	"github.com/firehydrant/api-client-go/client/metrics"
	"github.com/firehydrant/api-client-go/client/noauth"
	"github.com/firehydrant/api-client-go/client/nunc"
	"github.com/firehydrant/api-client-go/client/nunc_connections"
	"github.com/firehydrant/api-client-go/client/ping"
	"github.com/firehydrant/api-client-go/client/post_mortems"
	"github.com/firehydrant/api-client-go/client/priorities"
	"github.com/firehydrant/api-client-go/client/processing_log_entries"
	"github.com/firehydrant/api-client-go/client/reports"
	"github.com/firehydrant/api-client-go/client/runbook_audits"
	"github.com/firehydrant/api-client-go/client/runbooks"
	"github.com/firehydrant/api-client-go/client/saved_searches"
	"github.com/firehydrant/api-client-go/client/scheduled_maintenances"
	"github.com/firehydrant/api-client-go/client/schedules"
	"github.com/firehydrant/api-client-go/client/scim"
	"github.com/firehydrant/api-client-go/client/service_dependencies"
	"github.com/firehydrant/api-client-go/client/services"
	"github.com/firehydrant/api-client-go/client/severities"
	"github.com/firehydrant/api-client-go/client/severity_matrix"
	"github.com/firehydrant/api-client-go/client/signals"
	"github.com/firehydrant/api-client-go/client/signals_on_call"
	"github.com/firehydrant/api-client-go/client/status_update_templates"
	"github.com/firehydrant/api-client-go/client/task_lists"
	"github.com/firehydrant/api-client-go/client/teams"
	"github.com/firehydrant/api-client-go/client/ticketing"
	"github.com/firehydrant/api-client-go/client/users"
	"github.com/firehydrant/api-client-go/client/webhooks"
)

// Default fire hydrant API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.firehydrant.io"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new fire hydrant API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *FireHydrantAPI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new fire hydrant API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *FireHydrantAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new fire hydrant API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *FireHydrantAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(FireHydrantAPI)
	cli.Transport = transport
	cli.Ai = ai.New(transport, formats)
	cli.Alerts = alerts.New(transport, formats)
	cli.Bootstrap = bootstrap.New(transport, formats)
	cli.Catalogs = catalogs.New(transport, formats)
	cli.ChangeTypes = change_types.New(transport, formats)
	cli.Changes = changes.New(transport, formats)
	cli.ChecklistTemplates = checklist_templates.New(transport, formats)
	cli.Conversations = conversations.New(transport, formats)
	cli.CurrentUser = current_user.New(transport, formats)
	cli.CustomFields = custom_fields.New(transport, formats)
	cli.Entitlements = entitlements.New(transport, formats)
	cli.Environments = environments.New(transport, formats)
	cli.FormConfigurations = form_configurations.New(transport, formats)
	cli.Functionalities = functionalities.New(transport, formats)
	cli.IncidentRoles = incident_roles.New(transport, formats)
	cli.IncidentTags = incident_tags.New(transport, formats)
	cli.IncidentTypes = incident_types.New(transport, formats)
	cli.Incidents = incidents.New(transport, formats)
	cli.Infrastructures = infrastructures.New(transport, formats)
	cli.Integrations = integrations.New(transport, formats)
	cli.Metrics = metrics.New(transport, formats)
	cli.Noauth = noauth.New(transport, formats)
	cli.Nunc = nunc.New(transport, formats)
	cli.NuncConnections = nunc_connections.New(transport, formats)
	cli.Ping = ping.New(transport, formats)
	cli.PostMortems = post_mortems.New(transport, formats)
	cli.Priorities = priorities.New(transport, formats)
	cli.ProcessingLogEntries = processing_log_entries.New(transport, formats)
	cli.Reports = reports.New(transport, formats)
	cli.RunbookAudits = runbook_audits.New(transport, formats)
	cli.Runbooks = runbooks.New(transport, formats)
	cli.SavedSearches = saved_searches.New(transport, formats)
	cli.ScheduledMaintenances = scheduled_maintenances.New(transport, formats)
	cli.Schedules = schedules.New(transport, formats)
	cli.Scim = scim.New(transport, formats)
	cli.ServiceDependencies = service_dependencies.New(transport, formats)
	cli.Services = services.New(transport, formats)
	cli.Severities = severities.New(transport, formats)
	cli.SeverityMatrix = severity_matrix.New(transport, formats)
	cli.Signals = signals.New(transport, formats)
	cli.SignalsOnCall = signals_on_call.New(transport, formats)
	cli.StatusUpdateTemplates = status_update_templates.New(transport, formats)
	cli.TaskLists = task_lists.New(transport, formats)
	cli.Teams = teams.New(transport, formats)
	cli.Ticketing = ticketing.New(transport, formats)
	cli.Users = users.New(transport, formats)
	cli.Webhooks = webhooks.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// FireHydrantAPI is a client for fire hydrant API
type FireHydrantAPI struct {
	Ai ai.ClientService

	Alerts alerts.ClientService

	Bootstrap bootstrap.ClientService

	Catalogs catalogs.ClientService

	ChangeTypes change_types.ClientService

	Changes changes.ClientService

	ChecklistTemplates checklist_templates.ClientService

	Conversations conversations.ClientService

	CurrentUser current_user.ClientService

	CustomFields custom_fields.ClientService

	Entitlements entitlements.ClientService

	Environments environments.ClientService

	FormConfigurations form_configurations.ClientService

	Functionalities functionalities.ClientService

	IncidentRoles incident_roles.ClientService

	IncidentTags incident_tags.ClientService

	IncidentTypes incident_types.ClientService

	Incidents incidents.ClientService

	Infrastructures infrastructures.ClientService

	Integrations integrations.ClientService

	Metrics metrics.ClientService

	Noauth noauth.ClientService

	Nunc nunc.ClientService

	NuncConnections nunc_connections.ClientService

	Ping ping.ClientService

	PostMortems post_mortems.ClientService

	Priorities priorities.ClientService

	ProcessingLogEntries processing_log_entries.ClientService

	Reports reports.ClientService

	RunbookAudits runbook_audits.ClientService

	Runbooks runbooks.ClientService

	SavedSearches saved_searches.ClientService

	ScheduledMaintenances scheduled_maintenances.ClientService

	Schedules schedules.ClientService

	Scim scim.ClientService

	ServiceDependencies service_dependencies.ClientService

	Services services.ClientService

	Severities severities.ClientService

	SeverityMatrix severity_matrix.ClientService

	Signals signals.ClientService

	SignalsOnCall signals_on_call.ClientService

	StatusUpdateTemplates status_update_templates.ClientService

	TaskLists task_lists.ClientService

	Teams teams.ClientService

	Ticketing ticketing.ClientService

	Users users.ClientService

	Webhooks webhooks.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *FireHydrantAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Ai.SetTransport(transport)
	c.Alerts.SetTransport(transport)
	c.Bootstrap.SetTransport(transport)
	c.Catalogs.SetTransport(transport)
	c.ChangeTypes.SetTransport(transport)
	c.Changes.SetTransport(transport)
	c.ChecklistTemplates.SetTransport(transport)
	c.Conversations.SetTransport(transport)
	c.CurrentUser.SetTransport(transport)
	c.CustomFields.SetTransport(transport)
	c.Entitlements.SetTransport(transport)
	c.Environments.SetTransport(transport)
	c.FormConfigurations.SetTransport(transport)
	c.Functionalities.SetTransport(transport)
	c.IncidentRoles.SetTransport(transport)
	c.IncidentTags.SetTransport(transport)
	c.IncidentTypes.SetTransport(transport)
	c.Incidents.SetTransport(transport)
	c.Infrastructures.SetTransport(transport)
	c.Integrations.SetTransport(transport)
	c.Metrics.SetTransport(transport)
	c.Noauth.SetTransport(transport)
	c.Nunc.SetTransport(transport)
	c.NuncConnections.SetTransport(transport)
	c.Ping.SetTransport(transport)
	c.PostMortems.SetTransport(transport)
	c.Priorities.SetTransport(transport)
	c.ProcessingLogEntries.SetTransport(transport)
	c.Reports.SetTransport(transport)
	c.RunbookAudits.SetTransport(transport)
	c.Runbooks.SetTransport(transport)
	c.SavedSearches.SetTransport(transport)
	c.ScheduledMaintenances.SetTransport(transport)
	c.Schedules.SetTransport(transport)
	c.Scim.SetTransport(transport)
	c.ServiceDependencies.SetTransport(transport)
	c.Services.SetTransport(transport)
	c.Severities.SetTransport(transport)
	c.SeverityMatrix.SetTransport(transport)
	c.Signals.SetTransport(transport)
	c.SignalsOnCall.SetTransport(transport)
	c.StatusUpdateTemplates.SetTransport(transport)
	c.TaskLists.SetTransport(transport)
	c.Teams.SetTransport(transport)
	c.Ticketing.SetTransport(transport)
	c.Users.SetTransport(transport)
	c.Webhooks.SetTransport(transport)
}
