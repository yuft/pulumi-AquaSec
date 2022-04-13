// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aquasec
{
    [AquasecResourceType("aquasec:index/enforcerGroups:EnforcerGroups")]
    public partial class EnforcerGroups : Pulumi.CustomResource
    {
        [Output("admissionControl")]
        public Output<bool?> AdmissionControl { get; private set; } = null!;

        [Output("allowKubeEnforcerAudit")]
        public Output<bool?> AllowKubeEnforcerAudit { get; private set; } = null!;

        [Output("allowedApplications")]
        public Output<ImmutableArray<string>> AllowedApplications { get; private set; } = null!;

        [Output("allowedLabels")]
        public Output<ImmutableArray<string>> AllowedLabels { get; private set; } = null!;

        [Output("allowedRegistries")]
        public Output<ImmutableArray<string>> AllowedRegistries { get; private set; } = null!;

        [Output("antivirusProtection")]
        public Output<bool?> AntivirusProtection { get; private set; } = null!;

        [Output("aquaVersion")]
        public Output<string> AquaVersion { get; private set; } = null!;

        [Output("auditAll")]
        public Output<bool?> AuditAll { get; private set; } = null!;

        [Output("autoCopySecrets")]
        public Output<bool?> AutoCopySecrets { get; private set; } = null!;

        [Output("autoDiscoverConfigureRegistries")]
        public Output<bool?> AutoDiscoverConfigureRegistries { get; private set; } = null!;

        [Output("autoDiscoveryEnabled")]
        public Output<bool?> AutoDiscoveryEnabled { get; private set; } = null!;

        [Output("autoScanDiscoveredImagesRunningContainers")]
        public Output<bool?> AutoScanDiscoveredImagesRunningContainers { get; private set; } = null!;

        [Output("behavioralEngine")]
        public Output<bool?> BehavioralEngine { get; private set; } = null!;

        [Output("blockAdmissionControl")]
        public Output<bool?> BlockAdmissionControl { get; private set; } = null!;

        [Output("commands")]
        public Output<ImmutableArray<Outputs.EnforcerGroupsCommand>> Commands { get; private set; } = null!;

        [Output("connectedCount")]
        public Output<int> ConnectedCount { get; private set; } = null!;

        [Output("containerActivityProtection")]
        public Output<bool?> ContainerActivityProtection { get; private set; } = null!;

        [Output("containerAntivirusProtection")]
        public Output<bool?> ContainerAntivirusProtection { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("disconnectedCount")]
        public Output<int> DisconnectedCount { get; private set; } = null!;

        [Output("enforce")]
        public Output<bool?> Enforce { get; private set; } = null!;

        [Output("enforcerImageName")]
        public Output<string> EnforcerImageName { get; private set; } = null!;

        [Output("gatewayAddress")]
        public Output<string> GatewayAddress { get; private set; } = null!;

        [Output("gatewayName")]
        public Output<string> GatewayName { get; private set; } = null!;

        [Output("gateways")]
        public Output<ImmutableArray<string>> Gateways { get; private set; } = null!;

        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        [Output("highVulns")]
        public Output<int> HighVulns { get; private set; } = null!;

        [Output("hostAssurance")]
        public Output<bool?> HostAssurance { get; private set; } = null!;

        [Output("hostBehavioralEngine")]
        public Output<bool?> HostBehavioralEngine { get; private set; } = null!;

        [Output("hostNetworkProtection")]
        public Output<bool?> HostNetworkProtection { get; private set; } = null!;

        [Output("hostOs")]
        public Output<string> HostOs { get; private set; } = null!;

        [Output("hostProtection")]
        public Output<bool?> HostProtection { get; private set; } = null!;

        [Output("hostUserProtection")]
        public Output<bool?> HostUserProtection { get; private set; } = null!;

        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        [Output("hostsCount")]
        public Output<int> HostsCount { get; private set; } = null!;

        [Output("imageAssurance")]
        public Output<bool?> ImageAssurance { get; private set; } = null!;

        [Output("installCommand")]
        public Output<string> InstallCommand { get; private set; } = null!;

        [Output("kubeBenchImageName")]
        public Output<string?> KubeBenchImageName { get; private set; } = null!;

        [Output("lastUpdate")]
        public Output<int> LastUpdate { get; private set; } = null!;

        [Output("logicalName")]
        public Output<string?> LogicalName { get; private set; } = null!;

        [Output("lowVulns")]
        public Output<int> LowVulns { get; private set; } = null!;

        [Output("medVulns")]
        public Output<int> MedVulns { get; private set; } = null!;

        [Output("microEnforcerCertsSecretsName")]
        public Output<string?> MicroEnforcerCertsSecretsName { get; private set; } = null!;

        [Output("microEnforcerImageName")]
        public Output<string?> MicroEnforcerImageName { get; private set; } = null!;

        [Output("microEnforcerInjection")]
        public Output<bool?> MicroEnforcerInjection { get; private set; } = null!;

        [Output("microEnforcerSecretsName")]
        public Output<string?> MicroEnforcerSecretsName { get; private set; } = null!;

        [Output("negVulns")]
        public Output<int> NegVulns { get; private set; } = null!;

        [Output("networkActivityProtection")]
        public Output<bool?> NetworkActivityProtection { get; private set; } = null!;

        [Output("networkProtection")]
        public Output<bool?> NetworkProtection { get; private set; } = null!;

        [Output("orchestrators")]
        public Output<ImmutableArray<Outputs.EnforcerGroupsOrchestrator>> Orchestrators { get; private set; } = null!;

        [Output("pasDeploymentLink")]
        public Output<string> PasDeploymentLink { get; private set; } = null!;

        [Output("permission")]
        public Output<string?> Permission { get; private set; } = null!;

        [Output("riskExplorerAutoDiscovery")]
        public Output<bool?> RiskExplorerAutoDiscovery { get; private set; } = null!;

        [Output("runtimePolicyName")]
        public Output<string> RuntimePolicyName { get; private set; } = null!;

        [Output("runtimeType")]
        public Output<string> RuntimeType { get; private set; } = null!;

        [Output("syncHostImages")]
        public Output<bool?> SyncHostImages { get; private set; } = null!;

        [Output("syscallEnabled")]
        public Output<bool?> SyscallEnabled { get; private set; } = null!;

        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        [Output("userAccessControl")]
        public Output<bool?> UserAccessControl { get; private set; } = null!;


        /// <summary>
        /// Create a EnforcerGroups resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EnforcerGroups(string name, EnforcerGroupsArgs args, CustomResourceOptions? options = null)
            : base("aquasec:index/enforcerGroups:EnforcerGroups", name, args ?? new EnforcerGroupsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EnforcerGroups(string name, Input<string> id, EnforcerGroupsState? state = null, CustomResourceOptions? options = null)
            : base("aquasec:index/enforcerGroups:EnforcerGroups", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EnforcerGroups resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EnforcerGroups Get(string name, Input<string> id, EnforcerGroupsState? state = null, CustomResourceOptions? options = null)
        {
            return new EnforcerGroups(name, id, state, options);
        }
    }

    public sealed class EnforcerGroupsArgs : Pulumi.ResourceArgs
    {
        [Input("admissionControl")]
        public Input<bool>? AdmissionControl { get; set; }

        [Input("allowKubeEnforcerAudit")]
        public Input<bool>? AllowKubeEnforcerAudit { get; set; }

        [Input("allowedApplications")]
        private InputList<string>? _allowedApplications;
        public InputList<string> AllowedApplications
        {
            get => _allowedApplications ?? (_allowedApplications = new InputList<string>());
            set => _allowedApplications = value;
        }

        [Input("allowedLabels")]
        private InputList<string>? _allowedLabels;
        public InputList<string> AllowedLabels
        {
            get => _allowedLabels ?? (_allowedLabels = new InputList<string>());
            set => _allowedLabels = value;
        }

        [Input("allowedRegistries")]
        private InputList<string>? _allowedRegistries;
        public InputList<string> AllowedRegistries
        {
            get => _allowedRegistries ?? (_allowedRegistries = new InputList<string>());
            set => _allowedRegistries = value;
        }

        [Input("antivirusProtection")]
        public Input<bool>? AntivirusProtection { get; set; }

        [Input("auditAll")]
        public Input<bool>? AuditAll { get; set; }

        [Input("autoCopySecrets")]
        public Input<bool>? AutoCopySecrets { get; set; }

        [Input("autoDiscoverConfigureRegistries")]
        public Input<bool>? AutoDiscoverConfigureRegistries { get; set; }

        [Input("autoDiscoveryEnabled")]
        public Input<bool>? AutoDiscoveryEnabled { get; set; }

        [Input("autoScanDiscoveredImagesRunningContainers")]
        public Input<bool>? AutoScanDiscoveredImagesRunningContainers { get; set; }

        [Input("behavioralEngine")]
        public Input<bool>? BehavioralEngine { get; set; }

        [Input("blockAdmissionControl")]
        public Input<bool>? BlockAdmissionControl { get; set; }

        [Input("containerActivityProtection")]
        public Input<bool>? ContainerActivityProtection { get; set; }

        [Input("containerAntivirusProtection")]
        public Input<bool>? ContainerAntivirusProtection { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enforce")]
        public Input<bool>? Enforce { get; set; }

        [Input("gateways")]
        private InputList<string>? _gateways;
        public InputList<string> Gateways
        {
            get => _gateways ?? (_gateways = new InputList<string>());
            set => _gateways = value;
        }

        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        [Input("hostAssurance")]
        public Input<bool>? HostAssurance { get; set; }

        [Input("hostBehavioralEngine")]
        public Input<bool>? HostBehavioralEngine { get; set; }

        [Input("hostNetworkProtection")]
        public Input<bool>? HostNetworkProtection { get; set; }

        [Input("hostProtection")]
        public Input<bool>? HostProtection { get; set; }

        [Input("hostUserProtection")]
        public Input<bool>? HostUserProtection { get; set; }

        [Input("imageAssurance")]
        public Input<bool>? ImageAssurance { get; set; }

        [Input("kubeBenchImageName")]
        public Input<string>? KubeBenchImageName { get; set; }

        [Input("logicalName")]
        public Input<string>? LogicalName { get; set; }

        [Input("microEnforcerCertsSecretsName")]
        public Input<string>? MicroEnforcerCertsSecretsName { get; set; }

        [Input("microEnforcerImageName")]
        public Input<string>? MicroEnforcerImageName { get; set; }

        [Input("microEnforcerInjection")]
        public Input<bool>? MicroEnforcerInjection { get; set; }

        [Input("microEnforcerSecretsName")]
        public Input<string>? MicroEnforcerSecretsName { get; set; }

        [Input("networkActivityProtection")]
        public Input<bool>? NetworkActivityProtection { get; set; }

        [Input("networkProtection")]
        public Input<bool>? NetworkProtection { get; set; }

        [Input("orchestrators", required: true)]
        private InputList<Inputs.EnforcerGroupsOrchestratorArgs>? _orchestrators;
        public InputList<Inputs.EnforcerGroupsOrchestratorArgs> Orchestrators
        {
            get => _orchestrators ?? (_orchestrators = new InputList<Inputs.EnforcerGroupsOrchestratorArgs>());
            set => _orchestrators = value;
        }

        [Input("permission")]
        public Input<string>? Permission { get; set; }

        [Input("riskExplorerAutoDiscovery")]
        public Input<bool>? RiskExplorerAutoDiscovery { get; set; }

        [Input("syncHostImages")]
        public Input<bool>? SyncHostImages { get; set; }

        [Input("syscallEnabled")]
        public Input<bool>? SyscallEnabled { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("userAccessControl")]
        public Input<bool>? UserAccessControl { get; set; }

        public EnforcerGroupsArgs()
        {
        }
    }

    public sealed class EnforcerGroupsState : Pulumi.ResourceArgs
    {
        [Input("admissionControl")]
        public Input<bool>? AdmissionControl { get; set; }

        [Input("allowKubeEnforcerAudit")]
        public Input<bool>? AllowKubeEnforcerAudit { get; set; }

        [Input("allowedApplications")]
        private InputList<string>? _allowedApplications;
        public InputList<string> AllowedApplications
        {
            get => _allowedApplications ?? (_allowedApplications = new InputList<string>());
            set => _allowedApplications = value;
        }

        [Input("allowedLabels")]
        private InputList<string>? _allowedLabels;
        public InputList<string> AllowedLabels
        {
            get => _allowedLabels ?? (_allowedLabels = new InputList<string>());
            set => _allowedLabels = value;
        }

        [Input("allowedRegistries")]
        private InputList<string>? _allowedRegistries;
        public InputList<string> AllowedRegistries
        {
            get => _allowedRegistries ?? (_allowedRegistries = new InputList<string>());
            set => _allowedRegistries = value;
        }

        [Input("antivirusProtection")]
        public Input<bool>? AntivirusProtection { get; set; }

        [Input("aquaVersion")]
        public Input<string>? AquaVersion { get; set; }

        [Input("auditAll")]
        public Input<bool>? AuditAll { get; set; }

        [Input("autoCopySecrets")]
        public Input<bool>? AutoCopySecrets { get; set; }

        [Input("autoDiscoverConfigureRegistries")]
        public Input<bool>? AutoDiscoverConfigureRegistries { get; set; }

        [Input("autoDiscoveryEnabled")]
        public Input<bool>? AutoDiscoveryEnabled { get; set; }

        [Input("autoScanDiscoveredImagesRunningContainers")]
        public Input<bool>? AutoScanDiscoveredImagesRunningContainers { get; set; }

        [Input("behavioralEngine")]
        public Input<bool>? BehavioralEngine { get; set; }

        [Input("blockAdmissionControl")]
        public Input<bool>? BlockAdmissionControl { get; set; }

        [Input("commands")]
        private InputList<Inputs.EnforcerGroupsCommandGetArgs>? _commands;
        public InputList<Inputs.EnforcerGroupsCommandGetArgs> Commands
        {
            get => _commands ?? (_commands = new InputList<Inputs.EnforcerGroupsCommandGetArgs>());
            set => _commands = value;
        }

        [Input("connectedCount")]
        public Input<int>? ConnectedCount { get; set; }

        [Input("containerActivityProtection")]
        public Input<bool>? ContainerActivityProtection { get; set; }

        [Input("containerAntivirusProtection")]
        public Input<bool>? ContainerAntivirusProtection { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("disconnectedCount")]
        public Input<int>? DisconnectedCount { get; set; }

        [Input("enforce")]
        public Input<bool>? Enforce { get; set; }

        [Input("enforcerImageName")]
        public Input<string>? EnforcerImageName { get; set; }

        [Input("gatewayAddress")]
        public Input<string>? GatewayAddress { get; set; }

        [Input("gatewayName")]
        public Input<string>? GatewayName { get; set; }

        [Input("gateways")]
        private InputList<string>? _gateways;
        public InputList<string> Gateways
        {
            get => _gateways ?? (_gateways = new InputList<string>());
            set => _gateways = value;
        }

        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        [Input("highVulns")]
        public Input<int>? HighVulns { get; set; }

        [Input("hostAssurance")]
        public Input<bool>? HostAssurance { get; set; }

        [Input("hostBehavioralEngine")]
        public Input<bool>? HostBehavioralEngine { get; set; }

        [Input("hostNetworkProtection")]
        public Input<bool>? HostNetworkProtection { get; set; }

        [Input("hostOs")]
        public Input<string>? HostOs { get; set; }

        [Input("hostProtection")]
        public Input<bool>? HostProtection { get; set; }

        [Input("hostUserProtection")]
        public Input<bool>? HostUserProtection { get; set; }

        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        [Input("hostsCount")]
        public Input<int>? HostsCount { get; set; }

        [Input("imageAssurance")]
        public Input<bool>? ImageAssurance { get; set; }

        [Input("installCommand")]
        public Input<string>? InstallCommand { get; set; }

        [Input("kubeBenchImageName")]
        public Input<string>? KubeBenchImageName { get; set; }

        [Input("lastUpdate")]
        public Input<int>? LastUpdate { get; set; }

        [Input("logicalName")]
        public Input<string>? LogicalName { get; set; }

        [Input("lowVulns")]
        public Input<int>? LowVulns { get; set; }

        [Input("medVulns")]
        public Input<int>? MedVulns { get; set; }

        [Input("microEnforcerCertsSecretsName")]
        public Input<string>? MicroEnforcerCertsSecretsName { get; set; }

        [Input("microEnforcerImageName")]
        public Input<string>? MicroEnforcerImageName { get; set; }

        [Input("microEnforcerInjection")]
        public Input<bool>? MicroEnforcerInjection { get; set; }

        [Input("microEnforcerSecretsName")]
        public Input<string>? MicroEnforcerSecretsName { get; set; }

        [Input("negVulns")]
        public Input<int>? NegVulns { get; set; }

        [Input("networkActivityProtection")]
        public Input<bool>? NetworkActivityProtection { get; set; }

        [Input("networkProtection")]
        public Input<bool>? NetworkProtection { get; set; }

        [Input("orchestrators")]
        private InputList<Inputs.EnforcerGroupsOrchestratorGetArgs>? _orchestrators;
        public InputList<Inputs.EnforcerGroupsOrchestratorGetArgs> Orchestrators
        {
            get => _orchestrators ?? (_orchestrators = new InputList<Inputs.EnforcerGroupsOrchestratorGetArgs>());
            set => _orchestrators = value;
        }

        [Input("pasDeploymentLink")]
        public Input<string>? PasDeploymentLink { get; set; }

        [Input("permission")]
        public Input<string>? Permission { get; set; }

        [Input("riskExplorerAutoDiscovery")]
        public Input<bool>? RiskExplorerAutoDiscovery { get; set; }

        [Input("runtimePolicyName")]
        public Input<string>? RuntimePolicyName { get; set; }

        [Input("runtimeType")]
        public Input<string>? RuntimeType { get; set; }

        [Input("syncHostImages")]
        public Input<bool>? SyncHostImages { get; set; }

        [Input("syscallEnabled")]
        public Input<bool>? SyscallEnabled { get; set; }

        [Input("token")]
        public Input<string>? Token { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("userAccessControl")]
        public Input<bool>? UserAccessControl { get; set; }

        public EnforcerGroupsState()
        {
        }
    }
}