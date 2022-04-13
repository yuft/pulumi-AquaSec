// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aquasec
{
    [AquasecResourceType("aquasec:index/containerRuntimePolicy:ContainerRuntimePolicy")]
    public partial class ContainerRuntimePolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// List of executables that are allowed for the user.
        /// </summary>
        [Output("allowedExecutables")]
        public Output<ImmutableArray<string>> AllowedExecutables { get; private set; } = null!;

        /// <summary>
        /// Indicates the application scope of the service.
        /// </summary>
        [Output("applicationScopes")]
        public Output<ImmutableArray<string>> ApplicationScopes { get; private set; } = null!;

        /// <summary>
        /// If true, all network activity will be audited.
        /// </summary>
        [Output("auditAllNetworkActivity")]
        public Output<bool?> AuditAllNetworkActivity { get; private set; } = null!;

        /// <summary>
        /// If true, all process activity will be audited.
        /// </summary>
        [Output("auditAllProcessesActivity")]
        public Output<bool?> AuditAllProcessesActivity { get; private set; } = null!;

        /// <summary>
        /// If true, full command arguments will be audited.
        /// </summary>
        [Output("auditFullCommandArguments")]
        public Output<bool?> AuditFullCommandArguments { get; private set; } = null!;

        /// <summary>
        /// Username of the account that created the service.
        /// </summary>
        [Output("author")]
        public Output<string> Author { get; private set; } = null!;

        /// <summary>
        /// If true, prevent containers from running with access to host network.
        /// </summary>
        [Output("blockAccessHostNetwork")]
        public Output<bool?> BlockAccessHostNetwork { get; private set; } = null!;

        /// <summary>
        /// If true, prevent containers from running with adding capabilities with `--cap-add` privilege.
        /// </summary>
        [Output("blockAddingCapabilities")]
        public Output<bool?> BlockAddingCapabilities { get; private set; } = null!;

        /// <summary>
        /// If true, exec into a container is prevented.
        /// </summary>
        [Output("blockContainerExec")]
        public Output<bool?> BlockContainerExec { get; private set; } = null!;

        /// <summary>
        /// If true, prevent containers from running with the capability to bind in port lower than 1024.
        /// </summary>
        [Output("blockLowPortBinding")]
        public Output<bool?> BlockLowPortBinding { get; private set; } = null!;

        /// <summary>
        /// If true, running non-compliant image in the container is prevented.
        /// </summary>
        [Output("blockNonCompliantImages")]
        public Output<bool?> BlockNonCompliantImages { get; private set; } = null!;

        /// <summary>
        /// If true, running containers in non-compliant pods is prevented.
        /// </summary>
        [Output("blockNonCompliantWorkloads")]
        public Output<bool?> BlockNonCompliantWorkloads { get; private set; } = null!;

        /// <summary>
        /// If true, prevent containers from running with privileged container capability.
        /// </summary>
        [Output("blockPrivilegedContainers")]
        public Output<bool?> BlockPrivilegedContainers { get; private set; } = null!;

        /// <summary>
        /// If true, prevent containers from running with root user.
        /// </summary>
        [Output("blockRootUser")]
        public Output<bool?> BlockRootUser { get; private set; } = null!;

        /// <summary>
        /// If true, running images in the container that are not registered in Aqua is prevented.
        /// </summary>
        [Output("blockUnregisteredImages")]
        public Output<bool?> BlockUnregisteredImages { get; private set; } = null!;

        /// <summary>
        /// If true, prevent containers from running with the privilege to use the IPC namespace.
        /// </summary>
        [Output("blockUseIpcNamespace")]
        public Output<bool?> BlockUseIpcNamespace { get; private set; } = null!;

        /// <summary>
        /// If true, prevent containers from running with the privilege to use the PID namespace.
        /// </summary>
        [Output("blockUsePidNamespace")]
        public Output<bool?> BlockUsePidNamespace { get; private set; } = null!;

        /// <summary>
        /// If true, prevent containers from running with the privilege to use the user namespace.
        /// </summary>
        [Output("blockUseUserNamespace")]
        public Output<bool?> BlockUseUserNamespace { get; private set; } = null!;

        /// <summary>
        /// If true, prevent containers from running with the privilege to use the UTS namespace.
        /// </summary>
        [Output("blockUseUtsNamespace")]
        public Output<bool?> BlockUseUtsNamespace { get; private set; } = null!;

        /// <summary>
        /// If true, prevents containers from using specific Unix capabilities.
        /// </summary>
        [Output("blockedCapabilities")]
        public Output<ImmutableArray<string>> BlockedCapabilities { get; private set; } = null!;

        /// <summary>
        /// List of executables that are prevented from running in containers.
        /// </summary>
        [Output("blockedExecutables")]
        public Output<ImmutableArray<string>> BlockedExecutables { get; private set; } = null!;

        /// <summary>
        /// List of files that are prevented from being read, modified and executed in the containers.
        /// </summary>
        [Output("blockedFiles")]
        public Output<ImmutableArray<string>> BlockedFiles { get; private set; } = null!;

        /// <summary>
        /// List of blocked inbound ports.
        /// </summary>
        [Output("blockedInboundPorts")]
        public Output<ImmutableArray<string>> BlockedInboundPorts { get; private set; } = null!;

        /// <summary>
        /// List of blocked outbound ports.
        /// </summary>
        [Output("blockedOutboundPorts")]
        public Output<ImmutableArray<string>> BlockedOutboundPorts { get; private set; } = null!;

        /// <summary>
        /// Prevent containers from reading, writing, or executing all files in the list of packages.
        /// </summary>
        [Output("blockedPackages")]
        public Output<ImmutableArray<string>> BlockedPackages { get; private set; } = null!;

        /// <summary>
        /// List of volumes that are prevented from being mounted in the containers.
        /// </summary>
        [Output("blockedVolumes")]
        public Output<ImmutableArray<string>> BlockedVolumes { get; private set; } = null!;

        /// <summary>
        /// The description of the container runtime policy
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// If true, executables that are not in the original image is prevented from running.
        /// </summary>
        [Output("enableDriftPrevention")]
        public Output<bool?> EnableDriftPrevention { get; private set; } = null!;

        /// <summary>
        /// If true, fork bombs are prevented in the containers.
        /// </summary>
        [Output("enableForkGuard")]
        public Output<bool?> EnableForkGuard { get; private set; } = null!;

        /// <summary>
        /// If true, detect and prevent communication from containers to IP addresses known to have a bad reputation.
        /// </summary>
        [Output("enableIpReputationSecurity")]
        public Output<bool?> EnableIpReputationSecurity { get; private set; } = null!;

        /// <summary>
        /// If true, detects port scanning behavior in the container.
        /// </summary>
        [Output("enablePortScanDetection")]
        public Output<bool?> EnablePortScanDetection { get; private set; } = null!;

        /// <summary>
        /// Indicates if the runtime policy is enabled or not.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Indicates that policy should effect container execution (not just for audit).
        /// </summary>
        [Output("enforce")]
        public Output<bool?> Enforce { get; private set; } = null!;

        /// <summary>
        /// Indicates the number of days after which the runtime policy will be changed to enforce mode.
        /// </summary>
        [Output("enforceAfterDays")]
        public Output<int?> EnforceAfterDays { get; private set; } = null!;

        /// <summary>
        /// List of files and directories to be excluded from the read-only list.
        /// </summary>
        [Output("exceptionalReadonlyFilesAndDirectories")]
        public Output<ImmutableArray<string>> ExceptionalReadonlyFilesAndDirectories { get; private set; } = null!;

        /// <summary>
        /// Process limit for the fork guard.
        /// </summary>
        [Output("forkGuardProcessLimit")]
        public Output<int?> ForkGuardProcessLimit { get; private set; } = null!;

        /// <summary>
        /// If true, prevents the container from obtaining new privileges at runtime. (only enabled in enforce mode)
        /// </summary>
        [Output("limitNewPrivileges")]
        public Output<bool?> LimitNewPrivileges { get; private set; } = null!;

        /// <summary>
        /// Name of the container runtime policy
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// List of files and directories to be restricted as read-only
        /// </summary>
        [Output("readonlyFilesAndDirectories")]
        public Output<ImmutableArray<string>> ReadonlyFilesAndDirectories { get; private set; } = null!;

        /// <summary>
        /// Logical expression of how to compute the dependency of the scope variables.
        /// </summary>
        [Output("scopeExpression")]
        public Output<string> ScopeExpression { get; private set; } = null!;

        /// <summary>
        /// List of scope attributes.
        /// </summary>
        [Output("scopeVariables")]
        public Output<ImmutableArray<Outputs.ContainerRuntimePolicyScopeVariable>> ScopeVariables { get; private set; } = null!;


        /// <summary>
        /// Create a ContainerRuntimePolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ContainerRuntimePolicy(string name, ContainerRuntimePolicyArgs? args = null, CustomResourceOptions? options = null)
            : base("aquasec:index/containerRuntimePolicy:ContainerRuntimePolicy", name, args ?? new ContainerRuntimePolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ContainerRuntimePolicy(string name, Input<string> id, ContainerRuntimePolicyState? state = null, CustomResourceOptions? options = null)
            : base("aquasec:index/containerRuntimePolicy:ContainerRuntimePolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ContainerRuntimePolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ContainerRuntimePolicy Get(string name, Input<string> id, ContainerRuntimePolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new ContainerRuntimePolicy(name, id, state, options);
        }
    }

    public sealed class ContainerRuntimePolicyArgs : Pulumi.ResourceArgs
    {
        [Input("allowedExecutables")]
        private InputList<string>? _allowedExecutables;

        /// <summary>
        /// List of executables that are allowed for the user.
        /// </summary>
        public InputList<string> AllowedExecutables
        {
            get => _allowedExecutables ?? (_allowedExecutables = new InputList<string>());
            set => _allowedExecutables = value;
        }

        [Input("applicationScopes")]
        private InputList<string>? _applicationScopes;

        /// <summary>
        /// Indicates the application scope of the service.
        /// </summary>
        public InputList<string> ApplicationScopes
        {
            get => _applicationScopes ?? (_applicationScopes = new InputList<string>());
            set => _applicationScopes = value;
        }

        /// <summary>
        /// If true, all network activity will be audited.
        /// </summary>
        [Input("auditAllNetworkActivity")]
        public Input<bool>? AuditAllNetworkActivity { get; set; }

        /// <summary>
        /// If true, all process activity will be audited.
        /// </summary>
        [Input("auditAllProcessesActivity")]
        public Input<bool>? AuditAllProcessesActivity { get; set; }

        /// <summary>
        /// If true, full command arguments will be audited.
        /// </summary>
        [Input("auditFullCommandArguments")]
        public Input<bool>? AuditFullCommandArguments { get; set; }

        /// <summary>
        /// If true, prevent containers from running with access to host network.
        /// </summary>
        [Input("blockAccessHostNetwork")]
        public Input<bool>? BlockAccessHostNetwork { get; set; }

        /// <summary>
        /// If true, prevent containers from running with adding capabilities with `--cap-add` privilege.
        /// </summary>
        [Input("blockAddingCapabilities")]
        public Input<bool>? BlockAddingCapabilities { get; set; }

        /// <summary>
        /// If true, exec into a container is prevented.
        /// </summary>
        [Input("blockContainerExec")]
        public Input<bool>? BlockContainerExec { get; set; }

        /// <summary>
        /// If true, prevent containers from running with the capability to bind in port lower than 1024.
        /// </summary>
        [Input("blockLowPortBinding")]
        public Input<bool>? BlockLowPortBinding { get; set; }

        /// <summary>
        /// If true, running non-compliant image in the container is prevented.
        /// </summary>
        [Input("blockNonCompliantImages")]
        public Input<bool>? BlockNonCompliantImages { get; set; }

        /// <summary>
        /// If true, running containers in non-compliant pods is prevented.
        /// </summary>
        [Input("blockNonCompliantWorkloads")]
        public Input<bool>? BlockNonCompliantWorkloads { get; set; }

        /// <summary>
        /// If true, prevent containers from running with privileged container capability.
        /// </summary>
        [Input("blockPrivilegedContainers")]
        public Input<bool>? BlockPrivilegedContainers { get; set; }

        /// <summary>
        /// If true, prevent containers from running with root user.
        /// </summary>
        [Input("blockRootUser")]
        public Input<bool>? BlockRootUser { get; set; }

        /// <summary>
        /// If true, running images in the container that are not registered in Aqua is prevented.
        /// </summary>
        [Input("blockUnregisteredImages")]
        public Input<bool>? BlockUnregisteredImages { get; set; }

        /// <summary>
        /// If true, prevent containers from running with the privilege to use the IPC namespace.
        /// </summary>
        [Input("blockUseIpcNamespace")]
        public Input<bool>? BlockUseIpcNamespace { get; set; }

        /// <summary>
        /// If true, prevent containers from running with the privilege to use the PID namespace.
        /// </summary>
        [Input("blockUsePidNamespace")]
        public Input<bool>? BlockUsePidNamespace { get; set; }

        /// <summary>
        /// If true, prevent containers from running with the privilege to use the user namespace.
        /// </summary>
        [Input("blockUseUserNamespace")]
        public Input<bool>? BlockUseUserNamespace { get; set; }

        /// <summary>
        /// If true, prevent containers from running with the privilege to use the UTS namespace.
        /// </summary>
        [Input("blockUseUtsNamespace")]
        public Input<bool>? BlockUseUtsNamespace { get; set; }

        [Input("blockedCapabilities")]
        private InputList<string>? _blockedCapabilities;

        /// <summary>
        /// If true, prevents containers from using specific Unix capabilities.
        /// </summary>
        public InputList<string> BlockedCapabilities
        {
            get => _blockedCapabilities ?? (_blockedCapabilities = new InputList<string>());
            set => _blockedCapabilities = value;
        }

        [Input("blockedExecutables")]
        private InputList<string>? _blockedExecutables;

        /// <summary>
        /// List of executables that are prevented from running in containers.
        /// </summary>
        public InputList<string> BlockedExecutables
        {
            get => _blockedExecutables ?? (_blockedExecutables = new InputList<string>());
            set => _blockedExecutables = value;
        }

        [Input("blockedFiles")]
        private InputList<string>? _blockedFiles;

        /// <summary>
        /// List of files that are prevented from being read, modified and executed in the containers.
        /// </summary>
        public InputList<string> BlockedFiles
        {
            get => _blockedFiles ?? (_blockedFiles = new InputList<string>());
            set => _blockedFiles = value;
        }

        [Input("blockedInboundPorts")]
        private InputList<string>? _blockedInboundPorts;

        /// <summary>
        /// List of blocked inbound ports.
        /// </summary>
        public InputList<string> BlockedInboundPorts
        {
            get => _blockedInboundPorts ?? (_blockedInboundPorts = new InputList<string>());
            set => _blockedInboundPorts = value;
        }

        [Input("blockedOutboundPorts")]
        private InputList<string>? _blockedOutboundPorts;

        /// <summary>
        /// List of blocked outbound ports.
        /// </summary>
        public InputList<string> BlockedOutboundPorts
        {
            get => _blockedOutboundPorts ?? (_blockedOutboundPorts = new InputList<string>());
            set => _blockedOutboundPorts = value;
        }

        [Input("blockedPackages")]
        private InputList<string>? _blockedPackages;

        /// <summary>
        /// Prevent containers from reading, writing, or executing all files in the list of packages.
        /// </summary>
        public InputList<string> BlockedPackages
        {
            get => _blockedPackages ?? (_blockedPackages = new InputList<string>());
            set => _blockedPackages = value;
        }

        [Input("blockedVolumes")]
        private InputList<string>? _blockedVolumes;

        /// <summary>
        /// List of volumes that are prevented from being mounted in the containers.
        /// </summary>
        public InputList<string> BlockedVolumes
        {
            get => _blockedVolumes ?? (_blockedVolumes = new InputList<string>());
            set => _blockedVolumes = value;
        }

        /// <summary>
        /// The description of the container runtime policy
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If true, executables that are not in the original image is prevented from running.
        /// </summary>
        [Input("enableDriftPrevention")]
        public Input<bool>? EnableDriftPrevention { get; set; }

        /// <summary>
        /// If true, fork bombs are prevented in the containers.
        /// </summary>
        [Input("enableForkGuard")]
        public Input<bool>? EnableForkGuard { get; set; }

        /// <summary>
        /// If true, detect and prevent communication from containers to IP addresses known to have a bad reputation.
        /// </summary>
        [Input("enableIpReputationSecurity")]
        public Input<bool>? EnableIpReputationSecurity { get; set; }

        /// <summary>
        /// If true, detects port scanning behavior in the container.
        /// </summary>
        [Input("enablePortScanDetection")]
        public Input<bool>? EnablePortScanDetection { get; set; }

        /// <summary>
        /// Indicates if the runtime policy is enabled or not.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Indicates that policy should effect container execution (not just for audit).
        /// </summary>
        [Input("enforce")]
        public Input<bool>? Enforce { get; set; }

        /// <summary>
        /// Indicates the number of days after which the runtime policy will be changed to enforce mode.
        /// </summary>
        [Input("enforceAfterDays")]
        public Input<int>? EnforceAfterDays { get; set; }

        [Input("exceptionalReadonlyFilesAndDirectories")]
        private InputList<string>? _exceptionalReadonlyFilesAndDirectories;

        /// <summary>
        /// List of files and directories to be excluded from the read-only list.
        /// </summary>
        public InputList<string> ExceptionalReadonlyFilesAndDirectories
        {
            get => _exceptionalReadonlyFilesAndDirectories ?? (_exceptionalReadonlyFilesAndDirectories = new InputList<string>());
            set => _exceptionalReadonlyFilesAndDirectories = value;
        }

        /// <summary>
        /// Process limit for the fork guard.
        /// </summary>
        [Input("forkGuardProcessLimit")]
        public Input<int>? ForkGuardProcessLimit { get; set; }

        /// <summary>
        /// If true, prevents the container from obtaining new privileges at runtime. (only enabled in enforce mode)
        /// </summary>
        [Input("limitNewPrivileges")]
        public Input<bool>? LimitNewPrivileges { get; set; }

        /// <summary>
        /// Name of the container runtime policy
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("readonlyFilesAndDirectories")]
        private InputList<string>? _readonlyFilesAndDirectories;

        /// <summary>
        /// List of files and directories to be restricted as read-only
        /// </summary>
        public InputList<string> ReadonlyFilesAndDirectories
        {
            get => _readonlyFilesAndDirectories ?? (_readonlyFilesAndDirectories = new InputList<string>());
            set => _readonlyFilesAndDirectories = value;
        }

        /// <summary>
        /// Logical expression of how to compute the dependency of the scope variables.
        /// </summary>
        [Input("scopeExpression")]
        public Input<string>? ScopeExpression { get; set; }

        [Input("scopeVariables")]
        private InputList<Inputs.ContainerRuntimePolicyScopeVariableArgs>? _scopeVariables;

        /// <summary>
        /// List of scope attributes.
        /// </summary>
        public InputList<Inputs.ContainerRuntimePolicyScopeVariableArgs> ScopeVariables
        {
            get => _scopeVariables ?? (_scopeVariables = new InputList<Inputs.ContainerRuntimePolicyScopeVariableArgs>());
            set => _scopeVariables = value;
        }

        public ContainerRuntimePolicyArgs()
        {
        }
    }

    public sealed class ContainerRuntimePolicyState : Pulumi.ResourceArgs
    {
        [Input("allowedExecutables")]
        private InputList<string>? _allowedExecutables;

        /// <summary>
        /// List of executables that are allowed for the user.
        /// </summary>
        public InputList<string> AllowedExecutables
        {
            get => _allowedExecutables ?? (_allowedExecutables = new InputList<string>());
            set => _allowedExecutables = value;
        }

        [Input("applicationScopes")]
        private InputList<string>? _applicationScopes;

        /// <summary>
        /// Indicates the application scope of the service.
        /// </summary>
        public InputList<string> ApplicationScopes
        {
            get => _applicationScopes ?? (_applicationScopes = new InputList<string>());
            set => _applicationScopes = value;
        }

        /// <summary>
        /// If true, all network activity will be audited.
        /// </summary>
        [Input("auditAllNetworkActivity")]
        public Input<bool>? AuditAllNetworkActivity { get; set; }

        /// <summary>
        /// If true, all process activity will be audited.
        /// </summary>
        [Input("auditAllProcessesActivity")]
        public Input<bool>? AuditAllProcessesActivity { get; set; }

        /// <summary>
        /// If true, full command arguments will be audited.
        /// </summary>
        [Input("auditFullCommandArguments")]
        public Input<bool>? AuditFullCommandArguments { get; set; }

        /// <summary>
        /// Username of the account that created the service.
        /// </summary>
        [Input("author")]
        public Input<string>? Author { get; set; }

        /// <summary>
        /// If true, prevent containers from running with access to host network.
        /// </summary>
        [Input("blockAccessHostNetwork")]
        public Input<bool>? BlockAccessHostNetwork { get; set; }

        /// <summary>
        /// If true, prevent containers from running with adding capabilities with `--cap-add` privilege.
        /// </summary>
        [Input("blockAddingCapabilities")]
        public Input<bool>? BlockAddingCapabilities { get; set; }

        /// <summary>
        /// If true, exec into a container is prevented.
        /// </summary>
        [Input("blockContainerExec")]
        public Input<bool>? BlockContainerExec { get; set; }

        /// <summary>
        /// If true, prevent containers from running with the capability to bind in port lower than 1024.
        /// </summary>
        [Input("blockLowPortBinding")]
        public Input<bool>? BlockLowPortBinding { get; set; }

        /// <summary>
        /// If true, running non-compliant image in the container is prevented.
        /// </summary>
        [Input("blockNonCompliantImages")]
        public Input<bool>? BlockNonCompliantImages { get; set; }

        /// <summary>
        /// If true, running containers in non-compliant pods is prevented.
        /// </summary>
        [Input("blockNonCompliantWorkloads")]
        public Input<bool>? BlockNonCompliantWorkloads { get; set; }

        /// <summary>
        /// If true, prevent containers from running with privileged container capability.
        /// </summary>
        [Input("blockPrivilegedContainers")]
        public Input<bool>? BlockPrivilegedContainers { get; set; }

        /// <summary>
        /// If true, prevent containers from running with root user.
        /// </summary>
        [Input("blockRootUser")]
        public Input<bool>? BlockRootUser { get; set; }

        /// <summary>
        /// If true, running images in the container that are not registered in Aqua is prevented.
        /// </summary>
        [Input("blockUnregisteredImages")]
        public Input<bool>? BlockUnregisteredImages { get; set; }

        /// <summary>
        /// If true, prevent containers from running with the privilege to use the IPC namespace.
        /// </summary>
        [Input("blockUseIpcNamespace")]
        public Input<bool>? BlockUseIpcNamespace { get; set; }

        /// <summary>
        /// If true, prevent containers from running with the privilege to use the PID namespace.
        /// </summary>
        [Input("blockUsePidNamespace")]
        public Input<bool>? BlockUsePidNamespace { get; set; }

        /// <summary>
        /// If true, prevent containers from running with the privilege to use the user namespace.
        /// </summary>
        [Input("blockUseUserNamespace")]
        public Input<bool>? BlockUseUserNamespace { get; set; }

        /// <summary>
        /// If true, prevent containers from running with the privilege to use the UTS namespace.
        /// </summary>
        [Input("blockUseUtsNamespace")]
        public Input<bool>? BlockUseUtsNamespace { get; set; }

        [Input("blockedCapabilities")]
        private InputList<string>? _blockedCapabilities;

        /// <summary>
        /// If true, prevents containers from using specific Unix capabilities.
        /// </summary>
        public InputList<string> BlockedCapabilities
        {
            get => _blockedCapabilities ?? (_blockedCapabilities = new InputList<string>());
            set => _blockedCapabilities = value;
        }

        [Input("blockedExecutables")]
        private InputList<string>? _blockedExecutables;

        /// <summary>
        /// List of executables that are prevented from running in containers.
        /// </summary>
        public InputList<string> BlockedExecutables
        {
            get => _blockedExecutables ?? (_blockedExecutables = new InputList<string>());
            set => _blockedExecutables = value;
        }

        [Input("blockedFiles")]
        private InputList<string>? _blockedFiles;

        /// <summary>
        /// List of files that are prevented from being read, modified and executed in the containers.
        /// </summary>
        public InputList<string> BlockedFiles
        {
            get => _blockedFiles ?? (_blockedFiles = new InputList<string>());
            set => _blockedFiles = value;
        }

        [Input("blockedInboundPorts")]
        private InputList<string>? _blockedInboundPorts;

        /// <summary>
        /// List of blocked inbound ports.
        /// </summary>
        public InputList<string> BlockedInboundPorts
        {
            get => _blockedInboundPorts ?? (_blockedInboundPorts = new InputList<string>());
            set => _blockedInboundPorts = value;
        }

        [Input("blockedOutboundPorts")]
        private InputList<string>? _blockedOutboundPorts;

        /// <summary>
        /// List of blocked outbound ports.
        /// </summary>
        public InputList<string> BlockedOutboundPorts
        {
            get => _blockedOutboundPorts ?? (_blockedOutboundPorts = new InputList<string>());
            set => _blockedOutboundPorts = value;
        }

        [Input("blockedPackages")]
        private InputList<string>? _blockedPackages;

        /// <summary>
        /// Prevent containers from reading, writing, or executing all files in the list of packages.
        /// </summary>
        public InputList<string> BlockedPackages
        {
            get => _blockedPackages ?? (_blockedPackages = new InputList<string>());
            set => _blockedPackages = value;
        }

        [Input("blockedVolumes")]
        private InputList<string>? _blockedVolumes;

        /// <summary>
        /// List of volumes that are prevented from being mounted in the containers.
        /// </summary>
        public InputList<string> BlockedVolumes
        {
            get => _blockedVolumes ?? (_blockedVolumes = new InputList<string>());
            set => _blockedVolumes = value;
        }

        /// <summary>
        /// The description of the container runtime policy
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If true, executables that are not in the original image is prevented from running.
        /// </summary>
        [Input("enableDriftPrevention")]
        public Input<bool>? EnableDriftPrevention { get; set; }

        /// <summary>
        /// If true, fork bombs are prevented in the containers.
        /// </summary>
        [Input("enableForkGuard")]
        public Input<bool>? EnableForkGuard { get; set; }

        /// <summary>
        /// If true, detect and prevent communication from containers to IP addresses known to have a bad reputation.
        /// </summary>
        [Input("enableIpReputationSecurity")]
        public Input<bool>? EnableIpReputationSecurity { get; set; }

        /// <summary>
        /// If true, detects port scanning behavior in the container.
        /// </summary>
        [Input("enablePortScanDetection")]
        public Input<bool>? EnablePortScanDetection { get; set; }

        /// <summary>
        /// Indicates if the runtime policy is enabled or not.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Indicates that policy should effect container execution (not just for audit).
        /// </summary>
        [Input("enforce")]
        public Input<bool>? Enforce { get; set; }

        /// <summary>
        /// Indicates the number of days after which the runtime policy will be changed to enforce mode.
        /// </summary>
        [Input("enforceAfterDays")]
        public Input<int>? EnforceAfterDays { get; set; }

        [Input("exceptionalReadonlyFilesAndDirectories")]
        private InputList<string>? _exceptionalReadonlyFilesAndDirectories;

        /// <summary>
        /// List of files and directories to be excluded from the read-only list.
        /// </summary>
        public InputList<string> ExceptionalReadonlyFilesAndDirectories
        {
            get => _exceptionalReadonlyFilesAndDirectories ?? (_exceptionalReadonlyFilesAndDirectories = new InputList<string>());
            set => _exceptionalReadonlyFilesAndDirectories = value;
        }

        /// <summary>
        /// Process limit for the fork guard.
        /// </summary>
        [Input("forkGuardProcessLimit")]
        public Input<int>? ForkGuardProcessLimit { get; set; }

        /// <summary>
        /// If true, prevents the container from obtaining new privileges at runtime. (only enabled in enforce mode)
        /// </summary>
        [Input("limitNewPrivileges")]
        public Input<bool>? LimitNewPrivileges { get; set; }

        /// <summary>
        /// Name of the container runtime policy
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("readonlyFilesAndDirectories")]
        private InputList<string>? _readonlyFilesAndDirectories;

        /// <summary>
        /// List of files and directories to be restricted as read-only
        /// </summary>
        public InputList<string> ReadonlyFilesAndDirectories
        {
            get => _readonlyFilesAndDirectories ?? (_readonlyFilesAndDirectories = new InputList<string>());
            set => _readonlyFilesAndDirectories = value;
        }

        /// <summary>
        /// Logical expression of how to compute the dependency of the scope variables.
        /// </summary>
        [Input("scopeExpression")]
        public Input<string>? ScopeExpression { get; set; }

        [Input("scopeVariables")]
        private InputList<Inputs.ContainerRuntimePolicyScopeVariableGetArgs>? _scopeVariables;

        /// <summary>
        /// List of scope attributes.
        /// </summary>
        public InputList<Inputs.ContainerRuntimePolicyScopeVariableGetArgs> ScopeVariables
        {
            get => _scopeVariables ?? (_scopeVariables = new InputList<Inputs.ContainerRuntimePolicyScopeVariableGetArgs>());
            set => _scopeVariables = value;
        }

        public ContainerRuntimePolicyState()
        {
        }
    }
}