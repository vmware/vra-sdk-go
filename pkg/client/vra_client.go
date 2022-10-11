// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/client/about"
	"github.com/vmware/vra-sdk-go/pkg/client/about_the_service"
	"github.com/vmware/vra-sdk-go/pkg/client/blueprint"
	"github.com/vmware/vra-sdk-go/pkg/client/blueprint_requests"
	"github.com/vmware/vra-sdk-go/pkg/client/blueprint_terraform_integrations"
	"github.com/vmware/vra-sdk-go/pkg/client/blueprint_validation"
	"github.com/vmware/vra-sdk-go/pkg/client/catalog_admin_items"
	"github.com/vmware/vra-sdk-go/pkg/client/catalog_entitlements"
	"github.com/vmware/vra-sdk-go/pkg/client/catalog_item_types"
	"github.com/vmware/vra-sdk-go/pkg/client/catalog_items"
	"github.com/vmware/vra-sdk-go/pkg/client/catalog_sources"
	"github.com/vmware/vra-sdk-go/pkg/client/certificates"
	"github.com/vmware/vra-sdk-go/pkg/client/cloud_account"
	"github.com/vmware/vra-sdk-go/pkg/client/cluster_plans"
	"github.com/vmware/vra-sdk-go/pkg/client/compute"
	"github.com/vmware/vra-sdk-go/pkg/client/compute_gateway"
	"github.com/vmware/vra-sdk-go/pkg/client/compute_nat"
	"github.com/vmware/vra-sdk-go/pkg/client/content_source"
	"github.com/vmware/vra-sdk-go/pkg/client/custom_integrations"
	"github.com/vmware/vra-sdk-go/pkg/client/custom_naming"
	"github.com/vmware/vra-sdk-go/pkg/client/data_collector"
	"github.com/vmware/vra-sdk-go/pkg/client/deployment"
	"github.com/vmware/vra-sdk-go/pkg/client/deployment_actions"
	"github.com/vmware/vra-sdk-go/pkg/client/deployments"
	"github.com/vmware/vra-sdk-go/pkg/client/disk"
	"github.com/vmware/vra-sdk-go/pkg/client/endpoints"
	"github.com/vmware/vra-sdk-go/pkg/client/executions"
	"github.com/vmware/vra-sdk-go/pkg/client/fabric_aws_volume_types"
	"github.com/vmware/vra-sdk-go/pkg/client/fabric_azure_disk_encryption_sets"
	"github.com/vmware/vra-sdk-go/pkg/client/fabric_azure_storage_account"
	"github.com/vmware/vra-sdk-go/pkg/client/fabric_compute"
	"github.com/vmware/vra-sdk-go/pkg/client/fabric_flavors"
	"github.com/vmware/vra-sdk-go/pkg/client/fabric_images"
	"github.com/vmware/vra-sdk-go/pkg/client/fabric_network"
	"github.com/vmware/vra-sdk-go/pkg/client/fabric_vsphere_datastore"
	"github.com/vmware/vra-sdk-go/pkg/client/fabric_vsphere_storage_policies"
	"github.com/vmware/vra-sdk-go/pkg/client/flavor_profile"
	"github.com/vmware/vra-sdk-go/pkg/client/flavors"
	"github.com/vmware/vra-sdk-go/pkg/client/folders"
	"github.com/vmware/vra-sdk-go/pkg/client/icons"
	"github.com/vmware/vra-sdk-go/pkg/client/image_profile"
	"github.com/vmware/vra-sdk-go/pkg/client/images"
	"github.com/vmware/vra-sdk-go/pkg/client/installers"
	"github.com/vmware/vra-sdk-go/pkg/client/integration"
	"github.com/vmware/vra-sdk-go/pkg/client/kubernetes_clusters"
	"github.com/vmware/vra-sdk-go/pkg/client/kubernetes_zones"
	"github.com/vmware/vra-sdk-go/pkg/client/limit_ranges"
	"github.com/vmware/vra-sdk-go/pkg/client/load_balancer"
	"github.com/vmware/vra-sdk-go/pkg/client/location"
	"github.com/vmware/vra-sdk-go/pkg/client/login"
	"github.com/vmware/vra-sdk-go/pkg/client/namespaces"
	"github.com/vmware/vra-sdk-go/pkg/client/network"
	"github.com/vmware/vra-sdk-go/pkg/client/network_ip_range"
	"github.com/vmware/vra-sdk-go/pkg/client/network_profile"
	"github.com/vmware/vra-sdk-go/pkg/client/notification_scenario_configuration"
	"github.com/vmware/vra-sdk-go/pkg/client/onboarding_blueprints"
	"github.com/vmware/vra-sdk-go/pkg/client/onboarding_deployments"
	"github.com/vmware/vra-sdk-go/pkg/client/onboarding_plan_execution"
	"github.com/vmware/vra-sdk-go/pkg/client/onboarding_plans"
	"github.com/vmware/vra-sdk-go/pkg/client/onboarding_resources"
	"github.com/vmware/vra-sdk-go/pkg/client/pks_endpoints"
	"github.com/vmware/vra-sdk-go/pkg/client/perspective_sync"
	"github.com/vmware/vra-sdk-go/pkg/client/pipelines"
	"github.com/vmware/vra-sdk-go/pkg/client/policies"
	"github.com/vmware/vra-sdk-go/pkg/client/policy_decisions"
	"github.com/vmware/vra-sdk-go/pkg/client/policy_types"
	"github.com/vmware/vra-sdk-go/pkg/client/pricing_card_assignments"
	"github.com/vmware/vra-sdk-go/pkg/client/pricing_cards"
	"github.com/vmware/vra-sdk-go/pkg/client/project"
	"github.com/vmware/vra-sdk-go/pkg/client/projects"
	"github.com/vmware/vra-sdk-go/pkg/client/property"
	"github.com/vmware/vra-sdk-go/pkg/client/property_groups"
	"github.com/vmware/vra-sdk-go/pkg/client/provider_requests"
	"github.com/vmware/vra-sdk-go/pkg/client/query_for_discovered_machines"
	"github.com/vmware/vra-sdk-go/pkg/client/request"
	"github.com/vmware/vra-sdk-go/pkg/client/requests"
	"github.com/vmware/vra-sdk-go/pkg/client/resource_actions"
	"github.com/vmware/vra-sdk-go/pkg/client/resource_quotas"
	"github.com/vmware/vra-sdk-go/pkg/client/resource_types"
	"github.com/vmware/vra-sdk-go/pkg/client/resources"
	"github.com/vmware/vra-sdk-go/pkg/client/security_group"
	"github.com/vmware/vra-sdk-go/pkg/client/source_control_sync"
	"github.com/vmware/vra-sdk-go/pkg/client/storage_profile"
	"github.com/vmware/vra-sdk-go/pkg/client/supervisor_clusters"
	"github.com/vmware/vra-sdk-go/pkg/client/supervisor_namespaces"
	"github.com/vmware/vra-sdk-go/pkg/client/t_m_c_endpoints"
	"github.com/vmware/vra-sdk-go/pkg/client/tags"
	"github.com/vmware/vra-sdk-go/pkg/client/triggers"
	"github.com/vmware/vra-sdk-go/pkg/client/unregister_machines"
	"github.com/vmware/vra-sdk-go/pkg/client/user_events"
	"github.com/vmware/vra-sdk-go/pkg/client/user_operations"
	"github.com/vmware/vra-sdk-go/pkg/client/vcf"
	"github.com/vmware/vra-sdk-go/pkg/client/vsphere_endpoints"
	"github.com/vmware/vra-sdk-go/pkg/client/variables"
)

// Default vmware cloud assembly blueprint API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.mgmt.cloud.vmware.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new vmware cloud assembly blueprint API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *API {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new vmware cloud assembly blueprint API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *API {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new vmware cloud assembly blueprint API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *API {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(API)
	cli.Transport = transport
	cli.About = about.New(transport, formats)
	cli.AboutTheService = about_the_service.New(transport, formats)
	cli.Blueprint = blueprint.New(transport, formats)
	cli.BlueprintRequests = blueprint_requests.New(transport, formats)
	cli.BlueprintTerraformIntegrations = blueprint_terraform_integrations.New(transport, formats)
	cli.BlueprintValidation = blueprint_validation.New(transport, formats)
	cli.CatalogAdminItems = catalog_admin_items.New(transport, formats)
	cli.CatalogEntitlements = catalog_entitlements.New(transport, formats)
	cli.CatalogItemTypes = catalog_item_types.New(transport, formats)
	cli.CatalogItems = catalog_items.New(transport, formats)
	cli.CatalogSources = catalog_sources.New(transport, formats)
	cli.Certificates = certificates.New(transport, formats)
	cli.CloudAccount = cloud_account.New(transport, formats)
	cli.ClusterPlans = cluster_plans.New(transport, formats)
	cli.Compute = compute.New(transport, formats)
	cli.ComputeGateway = compute_gateway.New(transport, formats)
	cli.ComputeNat = compute_nat.New(transport, formats)
	cli.ContentSource = content_source.New(transport, formats)
	cli.CustomIntegrations = custom_integrations.New(transport, formats)
	cli.CustomNaming = custom_naming.New(transport, formats)
	cli.DataCollector = data_collector.New(transport, formats)
	cli.Deployment = deployment.New(transport, formats)
	cli.DeploymentActions = deployment_actions.New(transport, formats)
	cli.Deployments = deployments.New(transport, formats)
	cli.Disk = disk.New(transport, formats)
	cli.Endpoints = endpoints.New(transport, formats)
	cli.Executions = executions.New(transport, formats)
	cli.FabricawsVolumeTypes = fabric_aws_volume_types.New(transport, formats)
	cli.FabricAzureDiskEncryptionSets = fabric_azure_disk_encryption_sets.New(transport, formats)
	cli.FabricAzureStorageAccount = fabric_azure_storage_account.New(transport, formats)
	cli.FabricCompute = fabric_compute.New(transport, formats)
	cli.FabricFlavors = fabric_flavors.New(transport, formats)
	cli.FabricImages = fabric_images.New(transport, formats)
	cli.FabricNetwork = fabric_network.New(transport, formats)
	cli.FabricvSphereDatastore = fabric_vsphere_datastore.New(transport, formats)
	cli.FabricvSphereStoragePolicies = fabric_vsphere_storage_policies.New(transport, formats)
	cli.FlavorProfile = flavor_profile.New(transport, formats)
	cli.Flavors = flavors.New(transport, formats)
	cli.Folders = folders.New(transport, formats)
	cli.Icons = icons.New(transport, formats)
	cli.ImageProfile = image_profile.New(transport, formats)
	cli.Images = images.New(transport, formats)
	cli.Installers = installers.New(transport, formats)
	cli.Integration = integration.New(transport, formats)
	cli.KubernetesClusters = kubernetes_clusters.New(transport, formats)
	cli.KubernetesZones = kubernetes_zones.New(transport, formats)
	cli.LimitRanges = limit_ranges.New(transport, formats)
	cli.LoadBalancer = load_balancer.New(transport, formats)
	cli.Location = location.New(transport, formats)
	cli.Login = login.New(transport, formats)
	cli.Namespaces = namespaces.New(transport, formats)
	cli.Network = network.New(transport, formats)
	cli.NetworkIPRange = network_ip_range.New(transport, formats)
	cli.NetworkProfile = network_profile.New(transport, formats)
	cli.NotificationScenarioConfiguration = notification_scenario_configuration.New(transport, formats)
	cli.OnboardingBlueprints = onboarding_blueprints.New(transport, formats)
	cli.OnboardingDeployments = onboarding_deployments.New(transport, formats)
	cli.OnboardingPlanExecution = onboarding_plan_execution.New(transport, formats)
	cli.OnboardingPlans = onboarding_plans.New(transport, formats)
	cli.OnboardingResources = onboarding_resources.New(transport, formats)
	cli.PksEndpoints = pks_endpoints.New(transport, formats)
	cli.PerspectiveSync = perspective_sync.New(transport, formats)
	cli.Pipelines = pipelines.New(transport, formats)
	cli.Policies = policies.New(transport, formats)
	cli.PolicyDecisions = policy_decisions.New(transport, formats)
	cli.PolicyTypes = policy_types.New(transport, formats)
	cli.PricingCardAssignments = pricing_card_assignments.New(transport, formats)
	cli.PricingCards = pricing_cards.New(transport, formats)
	cli.Project = project.New(transport, formats)
	cli.Projects = projects.New(transport, formats)
	cli.Property = property.New(transport, formats)
	cli.PropertyGroups = property_groups.New(transport, formats)
	cli.ProviderRequests = provider_requests.New(transport, formats)
	cli.QueryForDiscoveredMachines = query_for_discovered_machines.New(transport, formats)
	cli.Request = request.New(transport, formats)
	cli.Requests = requests.New(transport, formats)
	cli.ResourceActions = resource_actions.New(transport, formats)
	cli.ResourceQuotas = resource_quotas.New(transport, formats)
	cli.ResourceTypes = resource_types.New(transport, formats)
	cli.Resources = resources.New(transport, formats)
	cli.SecurityGroup = security_group.New(transport, formats)
	cli.SourceControlSync = source_control_sync.New(transport, formats)
	cli.StorageProfile = storage_profile.New(transport, formats)
	cli.SupervisorClusters = supervisor_clusters.New(transport, formats)
	cli.SupervisorNamespaces = supervisor_namespaces.New(transport, formats)
	cli.TmcEndpoints = t_m_c_endpoints.New(transport, formats)
	cli.Tags = tags.New(transport, formats)
	cli.Triggers = triggers.New(transport, formats)
	cli.UnregisterMachines = unregister_machines.New(transport, formats)
	cli.UserEvents = user_events.New(transport, formats)
	cli.UserOperations = user_operations.New(transport, formats)
	cli.Vcf = vcf.New(transport, formats)
	cli.VSphereEndpoints = vsphere_endpoints.New(transport, formats)
	cli.Variables = variables.New(transport, formats)
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

// API is a client for vmware cloud assembly blueprint API
type API struct {
	About about.ClientService

	AboutTheService about_the_service.ClientService

	Blueprint blueprint.ClientService

	BlueprintRequests blueprint_requests.ClientService

	BlueprintTerraformIntegrations blueprint_terraform_integrations.ClientService

	BlueprintValidation blueprint_validation.ClientService

	CatalogAdminItems catalog_admin_items.ClientService

	CatalogEntitlements catalog_entitlements.ClientService

	CatalogItemTypes catalog_item_types.ClientService

	CatalogItems catalog_items.ClientService

	CatalogSources catalog_sources.ClientService

	Certificates certificates.ClientService

	CloudAccount cloud_account.ClientService

	ClusterPlans cluster_plans.ClientService

	Compute compute.ClientService

	ComputeGateway compute_gateway.ClientService

	ComputeNat compute_nat.ClientService

	ContentSource content_source.ClientService

	CustomIntegrations custom_integrations.ClientService

	CustomNaming custom_naming.ClientService

	DataCollector data_collector.ClientService

	Deployment deployment.ClientService

	DeploymentActions deployment_actions.ClientService

	Deployments deployments.ClientService

	Disk disk.ClientService

	Endpoints endpoints.ClientService

	Executions executions.ClientService

	FabricawsVolumeTypes fabric_aws_volume_types.ClientService

	FabricAzureDiskEncryptionSets fabric_azure_disk_encryption_sets.ClientService

	FabricAzureStorageAccount fabric_azure_storage_account.ClientService

	FabricCompute fabric_compute.ClientService

	FabricFlavors fabric_flavors.ClientService

	FabricImages fabric_images.ClientService

	FabricNetwork fabric_network.ClientService

	FabricvSphereDatastore fabric_vsphere_datastore.ClientService

	FabricvSphereStoragePolicies fabric_vsphere_storage_policies.ClientService

	FlavorProfile flavor_profile.ClientService

	Flavors flavors.ClientService

	Folders folders.ClientService

	Icons icons.ClientService

	ImageProfile image_profile.ClientService

	Images images.ClientService

	Installers installers.ClientService

	Integration integration.ClientService

	KubernetesClusters kubernetes_clusters.ClientService

	KubernetesZones kubernetes_zones.ClientService

	LimitRanges limit_ranges.ClientService

	LoadBalancer load_balancer.ClientService

	Location location.ClientService

	Login login.ClientService

	Namespaces namespaces.ClientService

	Network network.ClientService

	NetworkIPRange network_ip_range.ClientService

	NetworkProfile network_profile.ClientService

	NotificationScenarioConfiguration notification_scenario_configuration.ClientService

	OnboardingBlueprints onboarding_blueprints.ClientService

	OnboardingDeployments onboarding_deployments.ClientService

	OnboardingPlanExecution onboarding_plan_execution.ClientService

	OnboardingPlans onboarding_plans.ClientService

	OnboardingResources onboarding_resources.ClientService

	PksEndpoints pks_endpoints.ClientService

	PerspectiveSync perspective_sync.ClientService

	Pipelines pipelines.ClientService

	Policies policies.ClientService

	PolicyDecisions policy_decisions.ClientService

	PolicyTypes policy_types.ClientService

	PricingCardAssignments pricing_card_assignments.ClientService

	PricingCards pricing_cards.ClientService

	Project project.ClientService

	Projects projects.ClientService

	Property property.ClientService

	PropertyGroups property_groups.ClientService

	ProviderRequests provider_requests.ClientService

	QueryForDiscoveredMachines query_for_discovered_machines.ClientService

	Request request.ClientService

	Requests requests.ClientService

	ResourceActions resource_actions.ClientService

	ResourceQuotas resource_quotas.ClientService

	ResourceTypes resource_types.ClientService

	Resources resources.ClientService

	SecurityGroup security_group.ClientService

	SourceControlSync source_control_sync.ClientService

	StorageProfile storage_profile.ClientService

	SupervisorClusters supervisor_clusters.ClientService

	SupervisorNamespaces supervisor_namespaces.ClientService

	TmcEndpoints t_m_c_endpoints.ClientService

	Tags tags.ClientService

	Triggers triggers.ClientService

	UnregisterMachines unregister_machines.ClientService

	UserEvents user_events.ClientService

	UserOperations user_operations.ClientService

	Vcf vcf.ClientService

	VSphereEndpoints vsphere_endpoints.ClientService

	Variables variables.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *API) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.About.SetTransport(transport)
	c.AboutTheService.SetTransport(transport)
	c.Blueprint.SetTransport(transport)
	c.BlueprintRequests.SetTransport(transport)
	c.BlueprintTerraformIntegrations.SetTransport(transport)
	c.BlueprintValidation.SetTransport(transport)
	c.CatalogAdminItems.SetTransport(transport)
	c.CatalogEntitlements.SetTransport(transport)
	c.CatalogItemTypes.SetTransport(transport)
	c.CatalogItems.SetTransport(transport)
	c.CatalogSources.SetTransport(transport)
	c.Certificates.SetTransport(transport)
	c.CloudAccount.SetTransport(transport)
	c.ClusterPlans.SetTransport(transport)
	c.Compute.SetTransport(transport)
	c.ComputeGateway.SetTransport(transport)
	c.ComputeNat.SetTransport(transport)
	c.ContentSource.SetTransport(transport)
	c.CustomIntegrations.SetTransport(transport)
	c.CustomNaming.SetTransport(transport)
	c.DataCollector.SetTransport(transport)
	c.Deployment.SetTransport(transport)
	c.DeploymentActions.SetTransport(transport)
	c.Deployments.SetTransport(transport)
	c.Disk.SetTransport(transport)
	c.Endpoints.SetTransport(transport)
	c.Executions.SetTransport(transport)
	c.FabricawsVolumeTypes.SetTransport(transport)
	c.FabricAzureDiskEncryptionSets.SetTransport(transport)
	c.FabricAzureStorageAccount.SetTransport(transport)
	c.FabricCompute.SetTransport(transport)
	c.FabricFlavors.SetTransport(transport)
	c.FabricImages.SetTransport(transport)
	c.FabricNetwork.SetTransport(transport)
	c.FabricvSphereDatastore.SetTransport(transport)
	c.FabricvSphereStoragePolicies.SetTransport(transport)
	c.FlavorProfile.SetTransport(transport)
	c.Flavors.SetTransport(transport)
	c.Folders.SetTransport(transport)
	c.Icons.SetTransport(transport)
	c.ImageProfile.SetTransport(transport)
	c.Images.SetTransport(transport)
	c.Installers.SetTransport(transport)
	c.Integration.SetTransport(transport)
	c.KubernetesClusters.SetTransport(transport)
	c.KubernetesZones.SetTransport(transport)
	c.LimitRanges.SetTransport(transport)
	c.LoadBalancer.SetTransport(transport)
	c.Location.SetTransport(transport)
	c.Login.SetTransport(transport)
	c.Namespaces.SetTransport(transport)
	c.Network.SetTransport(transport)
	c.NetworkIPRange.SetTransport(transport)
	c.NetworkProfile.SetTransport(transport)
	c.NotificationScenarioConfiguration.SetTransport(transport)
	c.OnboardingBlueprints.SetTransport(transport)
	c.OnboardingDeployments.SetTransport(transport)
	c.OnboardingPlanExecution.SetTransport(transport)
	c.OnboardingPlans.SetTransport(transport)
	c.OnboardingResources.SetTransport(transport)
	c.PksEndpoints.SetTransport(transport)
	c.PerspectiveSync.SetTransport(transport)
	c.Pipelines.SetTransport(transport)
	c.Policies.SetTransport(transport)
	c.PolicyDecisions.SetTransport(transport)
	c.PolicyTypes.SetTransport(transport)
	c.PricingCardAssignments.SetTransport(transport)
	c.PricingCards.SetTransport(transport)
	c.Project.SetTransport(transport)
	c.Projects.SetTransport(transport)
	c.Property.SetTransport(transport)
	c.PropertyGroups.SetTransport(transport)
	c.ProviderRequests.SetTransport(transport)
	c.QueryForDiscoveredMachines.SetTransport(transport)
	c.Request.SetTransport(transport)
	c.Requests.SetTransport(transport)
	c.ResourceActions.SetTransport(transport)
	c.ResourceQuotas.SetTransport(transport)
	c.ResourceTypes.SetTransport(transport)
	c.Resources.SetTransport(transport)
	c.SecurityGroup.SetTransport(transport)
	c.SourceControlSync.SetTransport(transport)
	c.StorageProfile.SetTransport(transport)
	c.SupervisorClusters.SetTransport(transport)
	c.SupervisorNamespaces.SetTransport(transport)
	c.TmcEndpoints.SetTransport(transport)
	c.Tags.SetTransport(transport)
	c.Triggers.SetTransport(transport)
	c.UnregisterMachines.SetTransport(transport)
	c.UserEvents.SetTransport(transport)
	c.UserOperations.SetTransport(transport)
	c.Vcf.SetTransport(transport)
	c.VSphereEndpoints.SetTransport(transport)
	c.Variables.SetTransport(transport)
}
