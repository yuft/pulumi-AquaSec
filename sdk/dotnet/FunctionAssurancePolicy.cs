// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aquasec
{
    [AquasecResourceType("aquasec:index/functionAssurancePolicy:FunctionAssurancePolicy")]
    public partial class FunctionAssurancePolicy : Pulumi.CustomResource
    {
        [Output("allowedImages")]
        public Output<ImmutableArray<string>> AllowedImages { get; private set; } = null!;

        [Output("applicationScopes")]
        public Output<ImmutableArray<string>> ApplicationScopes { get; private set; } = null!;

        [Output("assuranceType")]
        public Output<string?> AssuranceType { get; private set; } = null!;

        [Output("auditOnFailure")]
        public Output<bool?> AuditOnFailure { get; private set; } = null!;

        [Output("author")]
        public Output<string> Author { get; private set; } = null!;

        [Output("autoScanConfigured")]
        public Output<bool?> AutoScanConfigured { get; private set; } = null!;

        [Output("autoScanEnabled")]
        public Output<bool?> AutoScanEnabled { get; private set; } = null!;

        [Output("autoScanTimes")]
        public Output<ImmutableArray<Outputs.FunctionAssurancePolicyAutoScanTime>> AutoScanTimes { get; private set; } = null!;

        [Output("blacklistPermissions")]
        public Output<ImmutableArray<string>> BlacklistPermissions { get; private set; } = null!;

        [Output("blacklistPermissionsEnabled")]
        public Output<bool?> BlacklistPermissionsEnabled { get; private set; } = null!;

        [Output("blacklistedLicenses")]
        public Output<ImmutableArray<string>> BlacklistedLicenses { get; private set; } = null!;

        [Output("blacklistedLicensesEnabled")]
        public Output<bool?> BlacklistedLicensesEnabled { get; private set; } = null!;

        [Output("blockFailed")]
        public Output<bool?> BlockFailed { get; private set; } = null!;

        [Output("controlExcludeNoFix")]
        public Output<bool?> ControlExcludeNoFix { get; private set; } = null!;

        [Output("customChecks")]
        public Output<ImmutableArray<Outputs.FunctionAssurancePolicyCustomCheck>> CustomChecks { get; private set; } = null!;

        [Output("customChecksEnabled")]
        public Output<bool?> CustomChecksEnabled { get; private set; } = null!;

        [Output("customSeverityEnabled")]
        public Output<bool?> CustomSeverityEnabled { get; private set; } = null!;

        [Output("cvesBlackListEnabled")]
        public Output<bool?> CvesBlackListEnabled { get; private set; } = null!;

        [Output("cvesBlackLists")]
        public Output<ImmutableArray<string>> CvesBlackLists { get; private set; } = null!;

        [Output("cvesWhiteListEnabled")]
        public Output<bool?> CvesWhiteListEnabled { get; private set; } = null!;

        [Output("cvesWhiteLists")]
        public Output<ImmutableArray<string>> CvesWhiteLists { get; private set; } = null!;

        [Output("cvssSeverity")]
        public Output<string?> CvssSeverity { get; private set; } = null!;

        [Output("cvssSeverityEnabled")]
        public Output<bool?> CvssSeverityEnabled { get; private set; } = null!;

        [Output("cvssSeverityExcludeNoFix")]
        public Output<bool?> CvssSeverityExcludeNoFix { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("disallowMalware")]
        public Output<bool?> DisallowMalware { get; private set; } = null!;

        [Output("dockerCisEnabled")]
        public Output<bool?> DockerCisEnabled { get; private set; } = null!;

        [Output("domain")]
        public Output<string?> Domain { get; private set; } = null!;

        [Output("domainName")]
        public Output<string?> DomainName { get; private set; } = null!;

        [Output("dtaEnabled")]
        public Output<bool?> DtaEnabled { get; private set; } = null!;

        [Output("dtaSeverity")]
        public Output<string?> DtaSeverity { get; private set; } = null!;

        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        [Output("enforce")]
        public Output<bool?> Enforce { get; private set; } = null!;

        [Output("enforceAfterDays")]
        public Output<int?> EnforceAfterDays { get; private set; } = null!;

        [Output("enforceExcessivePermissions")]
        public Output<bool?> EnforceExcessivePermissions { get; private set; } = null!;

        [Output("exceptionalMonitoredMalwarePaths")]
        public Output<ImmutableArray<string>> ExceptionalMonitoredMalwarePaths { get; private set; } = null!;

        [Output("failCicd")]
        public Output<bool?> FailCicd { get; private set; } = null!;

        [Output("forbiddenLabels")]
        public Output<ImmutableArray<Outputs.FunctionAssurancePolicyForbiddenLabel>> ForbiddenLabels { get; private set; } = null!;

        [Output("forbiddenLabelsEnabled")]
        public Output<bool?> ForbiddenLabelsEnabled { get; private set; } = null!;

        [Output("forceMicroenforcer")]
        public Output<bool?> ForceMicroenforcer { get; private set; } = null!;

        [Output("functionIntegrityEnabled")]
        public Output<bool?> FunctionIntegrityEnabled { get; private set; } = null!;

        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Output("id")]
        public Output<string> Id { get; private set; } = null!;

        [Output("ignoreRecentlyPublishedVln")]
        public Output<bool?> IgnoreRecentlyPublishedVln { get; private set; } = null!;

        [Output("ignoreRecentlyPublishedVlnPeriod")]
        public Output<int> IgnoreRecentlyPublishedVlnPeriod { get; private set; } = null!;

        [Output("ignoreRiskResourcesEnabled")]
        public Output<bool?> IgnoreRiskResourcesEnabled { get; private set; } = null!;

        [Output("ignoredRiskResources")]
        public Output<ImmutableArray<string>> IgnoredRiskResources { get; private set; } = null!;

        [Output("images")]
        public Output<ImmutableArray<string>> Images { get; private set; } = null!;

        [Output("kubeCisEnabled")]
        public Output<bool?> KubeCisEnabled { get; private set; } = null!;

        [Output("labels")]
        public Output<ImmutableArray<string>> Labels { get; private set; } = null!;

        [Output("malwareAction")]
        public Output<string?> MalwareAction { get; private set; } = null!;

        [Output("maximumScore")]
        public Output<double?> MaximumScore { get; private set; } = null!;

        [Output("maximumScoreEnabled")]
        public Output<bool?> MaximumScoreEnabled { get; private set; } = null!;

        [Output("maximumScoreExcludeNoFix")]
        public Output<bool?> MaximumScoreExcludeNoFix { get; private set; } = null!;

        [Output("monitoredMalwarePaths")]
        public Output<ImmutableArray<string>> MonitoredMalwarePaths { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("onlyNoneRootUsers")]
        public Output<bool?> OnlyNoneRootUsers { get; private set; } = null!;

        [Output("packagesBlackListEnabled")]
        public Output<bool?> PackagesBlackListEnabled { get; private set; } = null!;

        [Output("packagesBlackLists")]
        public Output<ImmutableArray<Outputs.FunctionAssurancePolicyPackagesBlackList>> PackagesBlackLists { get; private set; } = null!;

        [Output("packagesWhiteListEnabled")]
        public Output<bool?> PackagesWhiteListEnabled { get; private set; } = null!;

        [Output("packagesWhiteLists")]
        public Output<ImmutableArray<Outputs.FunctionAssurancePolicyPackagesWhiteList>> PackagesWhiteLists { get; private set; } = null!;

        [Output("partialResultsImageFail")]
        public Output<bool?> PartialResultsImageFail { get; private set; } = null!;

        [Output("readOnly")]
        public Output<bool?> ReadOnly { get; private set; } = null!;

        [Output("registries")]
        public Output<ImmutableArray<string>> Registries { get; private set; } = null!;

        [Output("registry")]
        public Output<string?> Registry { get; private set; } = null!;

        [Output("requiredLabels")]
        public Output<ImmutableArray<Outputs.FunctionAssurancePolicyRequiredLabel>> RequiredLabels { get; private set; } = null!;

        [Output("requiredLabelsEnabled")]
        public Output<bool?> RequiredLabelsEnabled { get; private set; } = null!;

        [Output("scanNfsMounts")]
        public Output<bool?> ScanNfsMounts { get; private set; } = null!;

        [Output("scanSensitiveData")]
        public Output<bool?> ScanSensitiveData { get; private set; } = null!;

        [Output("scapEnabled")]
        public Output<bool?> ScapEnabled { get; private set; } = null!;

        [Output("scapFiles")]
        public Output<ImmutableArray<string>> ScapFiles { get; private set; } = null!;

        [Output("scopes")]
        public Output<ImmutableArray<Outputs.FunctionAssurancePolicyScope>> Scopes { get; private set; } = null!;

        [Output("trustedBaseImages")]
        public Output<ImmutableArray<Outputs.FunctionAssurancePolicyTrustedBaseImage>> TrustedBaseImages { get; private set; } = null!;

        [Output("trustedBaseImagesEnabled")]
        public Output<bool?> TrustedBaseImagesEnabled { get; private set; } = null!;

        [Output("whitelistedLicenses")]
        public Output<ImmutableArray<string>> WhitelistedLicenses { get; private set; } = null!;

        [Output("whitelistedLicensesEnabled")]
        public Output<bool?> WhitelistedLicensesEnabled { get; private set; } = null!;


        /// <summary>
        /// Create a FunctionAssurancePolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FunctionAssurancePolicy(string name, FunctionAssurancePolicyArgs args, CustomResourceOptions? options = null)
            : base("aquasec:index/functionAssurancePolicy:FunctionAssurancePolicy", name, args ?? new FunctionAssurancePolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FunctionAssurancePolicy(string name, Input<string> id, FunctionAssurancePolicyState? state = null, CustomResourceOptions? options = null)
            : base("aquasec:index/functionAssurancePolicy:FunctionAssurancePolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FunctionAssurancePolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FunctionAssurancePolicy Get(string name, Input<string> id, FunctionAssurancePolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new FunctionAssurancePolicy(name, id, state, options);
        }
    }

    public sealed class FunctionAssurancePolicyArgs : Pulumi.ResourceArgs
    {
        [Input("allowedImages")]
        private InputList<string>? _allowedImages;
        public InputList<string> AllowedImages
        {
            get => _allowedImages ?? (_allowedImages = new InputList<string>());
            set => _allowedImages = value;
        }

        [Input("applicationScopes", required: true)]
        private InputList<string>? _applicationScopes;
        public InputList<string> ApplicationScopes
        {
            get => _applicationScopes ?? (_applicationScopes = new InputList<string>());
            set => _applicationScopes = value;
        }

        [Input("assuranceType")]
        public Input<string>? AssuranceType { get; set; }

        [Input("auditOnFailure")]
        public Input<bool>? AuditOnFailure { get; set; }

        [Input("autoScanConfigured")]
        public Input<bool>? AutoScanConfigured { get; set; }

        [Input("autoScanEnabled")]
        public Input<bool>? AutoScanEnabled { get; set; }

        [Input("autoScanTimes")]
        private InputList<Inputs.FunctionAssurancePolicyAutoScanTimeArgs>? _autoScanTimes;
        public InputList<Inputs.FunctionAssurancePolicyAutoScanTimeArgs> AutoScanTimes
        {
            get => _autoScanTimes ?? (_autoScanTimes = new InputList<Inputs.FunctionAssurancePolicyAutoScanTimeArgs>());
            set => _autoScanTimes = value;
        }

        [Input("blacklistPermissions")]
        private InputList<string>? _blacklistPermissions;
        public InputList<string> BlacklistPermissions
        {
            get => _blacklistPermissions ?? (_blacklistPermissions = new InputList<string>());
            set => _blacklistPermissions = value;
        }

        [Input("blacklistPermissionsEnabled")]
        public Input<bool>? BlacklistPermissionsEnabled { get; set; }

        [Input("blacklistedLicenses")]
        private InputList<string>? _blacklistedLicenses;
        public InputList<string> BlacklistedLicenses
        {
            get => _blacklistedLicenses ?? (_blacklistedLicenses = new InputList<string>());
            set => _blacklistedLicenses = value;
        }

        [Input("blacklistedLicensesEnabled")]
        public Input<bool>? BlacklistedLicensesEnabled { get; set; }

        [Input("blockFailed")]
        public Input<bool>? BlockFailed { get; set; }

        [Input("controlExcludeNoFix")]
        public Input<bool>? ControlExcludeNoFix { get; set; }

        [Input("customChecks")]
        private InputList<Inputs.FunctionAssurancePolicyCustomCheckArgs>? _customChecks;
        public InputList<Inputs.FunctionAssurancePolicyCustomCheckArgs> CustomChecks
        {
            get => _customChecks ?? (_customChecks = new InputList<Inputs.FunctionAssurancePolicyCustomCheckArgs>());
            set => _customChecks = value;
        }

        [Input("customChecksEnabled")]
        public Input<bool>? CustomChecksEnabled { get; set; }

        [Input("customSeverityEnabled")]
        public Input<bool>? CustomSeverityEnabled { get; set; }

        [Input("cvesBlackListEnabled")]
        public Input<bool>? CvesBlackListEnabled { get; set; }

        [Input("cvesBlackLists")]
        private InputList<string>? _cvesBlackLists;
        public InputList<string> CvesBlackLists
        {
            get => _cvesBlackLists ?? (_cvesBlackLists = new InputList<string>());
            set => _cvesBlackLists = value;
        }

        [Input("cvesWhiteListEnabled")]
        public Input<bool>? CvesWhiteListEnabled { get; set; }

        [Input("cvesWhiteLists")]
        private InputList<string>? _cvesWhiteLists;
        public InputList<string> CvesWhiteLists
        {
            get => _cvesWhiteLists ?? (_cvesWhiteLists = new InputList<string>());
            set => _cvesWhiteLists = value;
        }

        [Input("cvssSeverity")]
        public Input<string>? CvssSeverity { get; set; }

        [Input("cvssSeverityEnabled")]
        public Input<bool>? CvssSeverityEnabled { get; set; }

        [Input("cvssSeverityExcludeNoFix")]
        public Input<bool>? CvssSeverityExcludeNoFix { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("disallowMalware")]
        public Input<bool>? DisallowMalware { get; set; }

        [Input("dockerCisEnabled")]
        public Input<bool>? DockerCisEnabled { get; set; }

        [Input("domain")]
        public Input<string>? Domain { get; set; }

        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        [Input("dtaEnabled")]
        public Input<bool>? DtaEnabled { get; set; }

        [Input("dtaSeverity")]
        public Input<string>? DtaSeverity { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("enforce")]
        public Input<bool>? Enforce { get; set; }

        [Input("enforceAfterDays")]
        public Input<int>? EnforceAfterDays { get; set; }

        [Input("enforceExcessivePermissions")]
        public Input<bool>? EnforceExcessivePermissions { get; set; }

        [Input("exceptionalMonitoredMalwarePaths")]
        private InputList<string>? _exceptionalMonitoredMalwarePaths;
        public InputList<string> ExceptionalMonitoredMalwarePaths
        {
            get => _exceptionalMonitoredMalwarePaths ?? (_exceptionalMonitoredMalwarePaths = new InputList<string>());
            set => _exceptionalMonitoredMalwarePaths = value;
        }

        [Input("failCicd")]
        public Input<bool>? FailCicd { get; set; }

        [Input("forbiddenLabels")]
        private InputList<Inputs.FunctionAssurancePolicyForbiddenLabelArgs>? _forbiddenLabels;
        public InputList<Inputs.FunctionAssurancePolicyForbiddenLabelArgs> ForbiddenLabels
        {
            get => _forbiddenLabels ?? (_forbiddenLabels = new InputList<Inputs.FunctionAssurancePolicyForbiddenLabelArgs>());
            set => _forbiddenLabels = value;
        }

        [Input("forbiddenLabelsEnabled")]
        public Input<bool>? ForbiddenLabelsEnabled { get; set; }

        [Input("forceMicroenforcer")]
        public Input<bool>? ForceMicroenforcer { get; set; }

        [Input("functionIntegrityEnabled")]
        public Input<bool>? FunctionIntegrityEnabled { get; set; }

        [Input("ignoreRecentlyPublishedVln")]
        public Input<bool>? IgnoreRecentlyPublishedVln { get; set; }

        [Input("ignoreRiskResourcesEnabled")]
        public Input<bool>? IgnoreRiskResourcesEnabled { get; set; }

        [Input("ignoredRiskResources")]
        private InputList<string>? _ignoredRiskResources;
        public InputList<string> IgnoredRiskResources
        {
            get => _ignoredRiskResources ?? (_ignoredRiskResources = new InputList<string>());
            set => _ignoredRiskResources = value;
        }

        [Input("images")]
        private InputList<string>? _images;
        public InputList<string> Images
        {
            get => _images ?? (_images = new InputList<string>());
            set => _images = value;
        }

        [Input("kubeCisEnabled")]
        public Input<bool>? KubeCisEnabled { get; set; }

        [Input("labels")]
        private InputList<string>? _labels;
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        [Input("malwareAction")]
        public Input<string>? MalwareAction { get; set; }

        [Input("maximumScore")]
        public Input<double>? MaximumScore { get; set; }

        [Input("maximumScoreEnabled")]
        public Input<bool>? MaximumScoreEnabled { get; set; }

        [Input("maximumScoreExcludeNoFix")]
        public Input<bool>? MaximumScoreExcludeNoFix { get; set; }

        [Input("monitoredMalwarePaths")]
        private InputList<string>? _monitoredMalwarePaths;
        public InputList<string> MonitoredMalwarePaths
        {
            get => _monitoredMalwarePaths ?? (_monitoredMalwarePaths = new InputList<string>());
            set => _monitoredMalwarePaths = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("onlyNoneRootUsers")]
        public Input<bool>? OnlyNoneRootUsers { get; set; }

        [Input("packagesBlackListEnabled")]
        public Input<bool>? PackagesBlackListEnabled { get; set; }

        [Input("packagesBlackLists")]
        private InputList<Inputs.FunctionAssurancePolicyPackagesBlackListArgs>? _packagesBlackLists;
        public InputList<Inputs.FunctionAssurancePolicyPackagesBlackListArgs> PackagesBlackLists
        {
            get => _packagesBlackLists ?? (_packagesBlackLists = new InputList<Inputs.FunctionAssurancePolicyPackagesBlackListArgs>());
            set => _packagesBlackLists = value;
        }

        [Input("packagesWhiteListEnabled")]
        public Input<bool>? PackagesWhiteListEnabled { get; set; }

        [Input("packagesWhiteLists")]
        private InputList<Inputs.FunctionAssurancePolicyPackagesWhiteListArgs>? _packagesWhiteLists;
        public InputList<Inputs.FunctionAssurancePolicyPackagesWhiteListArgs> PackagesWhiteLists
        {
            get => _packagesWhiteLists ?? (_packagesWhiteLists = new InputList<Inputs.FunctionAssurancePolicyPackagesWhiteListArgs>());
            set => _packagesWhiteLists = value;
        }

        [Input("partialResultsImageFail")]
        public Input<bool>? PartialResultsImageFail { get; set; }

        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        [Input("registries")]
        private InputList<string>? _registries;
        public InputList<string> Registries
        {
            get => _registries ?? (_registries = new InputList<string>());
            set => _registries = value;
        }

        [Input("registry")]
        public Input<string>? Registry { get; set; }

        [Input("requiredLabels")]
        private InputList<Inputs.FunctionAssurancePolicyRequiredLabelArgs>? _requiredLabels;
        public InputList<Inputs.FunctionAssurancePolicyRequiredLabelArgs> RequiredLabels
        {
            get => _requiredLabels ?? (_requiredLabels = new InputList<Inputs.FunctionAssurancePolicyRequiredLabelArgs>());
            set => _requiredLabels = value;
        }

        [Input("requiredLabelsEnabled")]
        public Input<bool>? RequiredLabelsEnabled { get; set; }

        [Input("scanNfsMounts")]
        public Input<bool>? ScanNfsMounts { get; set; }

        [Input("scanSensitiveData")]
        public Input<bool>? ScanSensitiveData { get; set; }

        [Input("scapEnabled")]
        public Input<bool>? ScapEnabled { get; set; }

        [Input("scapFiles")]
        private InputList<string>? _scapFiles;
        public InputList<string> ScapFiles
        {
            get => _scapFiles ?? (_scapFiles = new InputList<string>());
            set => _scapFiles = value;
        }

        [Input("scopes")]
        private InputList<Inputs.FunctionAssurancePolicyScopeArgs>? _scopes;
        public InputList<Inputs.FunctionAssurancePolicyScopeArgs> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<Inputs.FunctionAssurancePolicyScopeArgs>());
            set => _scopes = value;
        }

        [Input("trustedBaseImages")]
        private InputList<Inputs.FunctionAssurancePolicyTrustedBaseImageArgs>? _trustedBaseImages;
        public InputList<Inputs.FunctionAssurancePolicyTrustedBaseImageArgs> TrustedBaseImages
        {
            get => _trustedBaseImages ?? (_trustedBaseImages = new InputList<Inputs.FunctionAssurancePolicyTrustedBaseImageArgs>());
            set => _trustedBaseImages = value;
        }

        [Input("trustedBaseImagesEnabled")]
        public Input<bool>? TrustedBaseImagesEnabled { get; set; }

        [Input("whitelistedLicenses")]
        private InputList<string>? _whitelistedLicenses;
        public InputList<string> WhitelistedLicenses
        {
            get => _whitelistedLicenses ?? (_whitelistedLicenses = new InputList<string>());
            set => _whitelistedLicenses = value;
        }

        [Input("whitelistedLicensesEnabled")]
        public Input<bool>? WhitelistedLicensesEnabled { get; set; }

        public FunctionAssurancePolicyArgs()
        {
        }
    }

    public sealed class FunctionAssurancePolicyState : Pulumi.ResourceArgs
    {
        [Input("allowedImages")]
        private InputList<string>? _allowedImages;
        public InputList<string> AllowedImages
        {
            get => _allowedImages ?? (_allowedImages = new InputList<string>());
            set => _allowedImages = value;
        }

        [Input("applicationScopes")]
        private InputList<string>? _applicationScopes;
        public InputList<string> ApplicationScopes
        {
            get => _applicationScopes ?? (_applicationScopes = new InputList<string>());
            set => _applicationScopes = value;
        }

        [Input("assuranceType")]
        public Input<string>? AssuranceType { get; set; }

        [Input("auditOnFailure")]
        public Input<bool>? AuditOnFailure { get; set; }

        [Input("author")]
        public Input<string>? Author { get; set; }

        [Input("autoScanConfigured")]
        public Input<bool>? AutoScanConfigured { get; set; }

        [Input("autoScanEnabled")]
        public Input<bool>? AutoScanEnabled { get; set; }

        [Input("autoScanTimes")]
        private InputList<Inputs.FunctionAssurancePolicyAutoScanTimeGetArgs>? _autoScanTimes;
        public InputList<Inputs.FunctionAssurancePolicyAutoScanTimeGetArgs> AutoScanTimes
        {
            get => _autoScanTimes ?? (_autoScanTimes = new InputList<Inputs.FunctionAssurancePolicyAutoScanTimeGetArgs>());
            set => _autoScanTimes = value;
        }

        [Input("blacklistPermissions")]
        private InputList<string>? _blacklistPermissions;
        public InputList<string> BlacklistPermissions
        {
            get => _blacklistPermissions ?? (_blacklistPermissions = new InputList<string>());
            set => _blacklistPermissions = value;
        }

        [Input("blacklistPermissionsEnabled")]
        public Input<bool>? BlacklistPermissionsEnabled { get; set; }

        [Input("blacklistedLicenses")]
        private InputList<string>? _blacklistedLicenses;
        public InputList<string> BlacklistedLicenses
        {
            get => _blacklistedLicenses ?? (_blacklistedLicenses = new InputList<string>());
            set => _blacklistedLicenses = value;
        }

        [Input("blacklistedLicensesEnabled")]
        public Input<bool>? BlacklistedLicensesEnabled { get; set; }

        [Input("blockFailed")]
        public Input<bool>? BlockFailed { get; set; }

        [Input("controlExcludeNoFix")]
        public Input<bool>? ControlExcludeNoFix { get; set; }

        [Input("customChecks")]
        private InputList<Inputs.FunctionAssurancePolicyCustomCheckGetArgs>? _customChecks;
        public InputList<Inputs.FunctionAssurancePolicyCustomCheckGetArgs> CustomChecks
        {
            get => _customChecks ?? (_customChecks = new InputList<Inputs.FunctionAssurancePolicyCustomCheckGetArgs>());
            set => _customChecks = value;
        }

        [Input("customChecksEnabled")]
        public Input<bool>? CustomChecksEnabled { get; set; }

        [Input("customSeverityEnabled")]
        public Input<bool>? CustomSeverityEnabled { get; set; }

        [Input("cvesBlackListEnabled")]
        public Input<bool>? CvesBlackListEnabled { get; set; }

        [Input("cvesBlackLists")]
        private InputList<string>? _cvesBlackLists;
        public InputList<string> CvesBlackLists
        {
            get => _cvesBlackLists ?? (_cvesBlackLists = new InputList<string>());
            set => _cvesBlackLists = value;
        }

        [Input("cvesWhiteListEnabled")]
        public Input<bool>? CvesWhiteListEnabled { get; set; }

        [Input("cvesWhiteLists")]
        private InputList<string>? _cvesWhiteLists;
        public InputList<string> CvesWhiteLists
        {
            get => _cvesWhiteLists ?? (_cvesWhiteLists = new InputList<string>());
            set => _cvesWhiteLists = value;
        }

        [Input("cvssSeverity")]
        public Input<string>? CvssSeverity { get; set; }

        [Input("cvssSeverityEnabled")]
        public Input<bool>? CvssSeverityEnabled { get; set; }

        [Input("cvssSeverityExcludeNoFix")]
        public Input<bool>? CvssSeverityExcludeNoFix { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("disallowMalware")]
        public Input<bool>? DisallowMalware { get; set; }

        [Input("dockerCisEnabled")]
        public Input<bool>? DockerCisEnabled { get; set; }

        [Input("domain")]
        public Input<string>? Domain { get; set; }

        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        [Input("dtaEnabled")]
        public Input<bool>? DtaEnabled { get; set; }

        [Input("dtaSeverity")]
        public Input<string>? DtaSeverity { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("enforce")]
        public Input<bool>? Enforce { get; set; }

        [Input("enforceAfterDays")]
        public Input<int>? EnforceAfterDays { get; set; }

        [Input("enforceExcessivePermissions")]
        public Input<bool>? EnforceExcessivePermissions { get; set; }

        [Input("exceptionalMonitoredMalwarePaths")]
        private InputList<string>? _exceptionalMonitoredMalwarePaths;
        public InputList<string> ExceptionalMonitoredMalwarePaths
        {
            get => _exceptionalMonitoredMalwarePaths ?? (_exceptionalMonitoredMalwarePaths = new InputList<string>());
            set => _exceptionalMonitoredMalwarePaths = value;
        }

        [Input("failCicd")]
        public Input<bool>? FailCicd { get; set; }

        [Input("forbiddenLabels")]
        private InputList<Inputs.FunctionAssurancePolicyForbiddenLabelGetArgs>? _forbiddenLabels;
        public InputList<Inputs.FunctionAssurancePolicyForbiddenLabelGetArgs> ForbiddenLabels
        {
            get => _forbiddenLabels ?? (_forbiddenLabels = new InputList<Inputs.FunctionAssurancePolicyForbiddenLabelGetArgs>());
            set => _forbiddenLabels = value;
        }

        [Input("forbiddenLabelsEnabled")]
        public Input<bool>? ForbiddenLabelsEnabled { get; set; }

        [Input("forceMicroenforcer")]
        public Input<bool>? ForceMicroenforcer { get; set; }

        [Input("functionIntegrityEnabled")]
        public Input<bool>? FunctionIntegrityEnabled { get; set; }

        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("ignoreRecentlyPublishedVln")]
        public Input<bool>? IgnoreRecentlyPublishedVln { get; set; }

        [Input("ignoreRecentlyPublishedVlnPeriod")]
        public Input<int>? IgnoreRecentlyPublishedVlnPeriod { get; set; }

        [Input("ignoreRiskResourcesEnabled")]
        public Input<bool>? IgnoreRiskResourcesEnabled { get; set; }

        [Input("ignoredRiskResources")]
        private InputList<string>? _ignoredRiskResources;
        public InputList<string> IgnoredRiskResources
        {
            get => _ignoredRiskResources ?? (_ignoredRiskResources = new InputList<string>());
            set => _ignoredRiskResources = value;
        }

        [Input("images")]
        private InputList<string>? _images;
        public InputList<string> Images
        {
            get => _images ?? (_images = new InputList<string>());
            set => _images = value;
        }

        [Input("kubeCisEnabled")]
        public Input<bool>? KubeCisEnabled { get; set; }

        [Input("labels")]
        private InputList<string>? _labels;
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        [Input("malwareAction")]
        public Input<string>? MalwareAction { get; set; }

        [Input("maximumScore")]
        public Input<double>? MaximumScore { get; set; }

        [Input("maximumScoreEnabled")]
        public Input<bool>? MaximumScoreEnabled { get; set; }

        [Input("maximumScoreExcludeNoFix")]
        public Input<bool>? MaximumScoreExcludeNoFix { get; set; }

        [Input("monitoredMalwarePaths")]
        private InputList<string>? _monitoredMalwarePaths;
        public InputList<string> MonitoredMalwarePaths
        {
            get => _monitoredMalwarePaths ?? (_monitoredMalwarePaths = new InputList<string>());
            set => _monitoredMalwarePaths = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("onlyNoneRootUsers")]
        public Input<bool>? OnlyNoneRootUsers { get; set; }

        [Input("packagesBlackListEnabled")]
        public Input<bool>? PackagesBlackListEnabled { get; set; }

        [Input("packagesBlackLists")]
        private InputList<Inputs.FunctionAssurancePolicyPackagesBlackListGetArgs>? _packagesBlackLists;
        public InputList<Inputs.FunctionAssurancePolicyPackagesBlackListGetArgs> PackagesBlackLists
        {
            get => _packagesBlackLists ?? (_packagesBlackLists = new InputList<Inputs.FunctionAssurancePolicyPackagesBlackListGetArgs>());
            set => _packagesBlackLists = value;
        }

        [Input("packagesWhiteListEnabled")]
        public Input<bool>? PackagesWhiteListEnabled { get; set; }

        [Input("packagesWhiteLists")]
        private InputList<Inputs.FunctionAssurancePolicyPackagesWhiteListGetArgs>? _packagesWhiteLists;
        public InputList<Inputs.FunctionAssurancePolicyPackagesWhiteListGetArgs> PackagesWhiteLists
        {
            get => _packagesWhiteLists ?? (_packagesWhiteLists = new InputList<Inputs.FunctionAssurancePolicyPackagesWhiteListGetArgs>());
            set => _packagesWhiteLists = value;
        }

        [Input("partialResultsImageFail")]
        public Input<bool>? PartialResultsImageFail { get; set; }

        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        [Input("registries")]
        private InputList<string>? _registries;
        public InputList<string> Registries
        {
            get => _registries ?? (_registries = new InputList<string>());
            set => _registries = value;
        }

        [Input("registry")]
        public Input<string>? Registry { get; set; }

        [Input("requiredLabels")]
        private InputList<Inputs.FunctionAssurancePolicyRequiredLabelGetArgs>? _requiredLabels;
        public InputList<Inputs.FunctionAssurancePolicyRequiredLabelGetArgs> RequiredLabels
        {
            get => _requiredLabels ?? (_requiredLabels = new InputList<Inputs.FunctionAssurancePolicyRequiredLabelGetArgs>());
            set => _requiredLabels = value;
        }

        [Input("requiredLabelsEnabled")]
        public Input<bool>? RequiredLabelsEnabled { get; set; }

        [Input("scanNfsMounts")]
        public Input<bool>? ScanNfsMounts { get; set; }

        [Input("scanSensitiveData")]
        public Input<bool>? ScanSensitiveData { get; set; }

        [Input("scapEnabled")]
        public Input<bool>? ScapEnabled { get; set; }

        [Input("scapFiles")]
        private InputList<string>? _scapFiles;
        public InputList<string> ScapFiles
        {
            get => _scapFiles ?? (_scapFiles = new InputList<string>());
            set => _scapFiles = value;
        }

        [Input("scopes")]
        private InputList<Inputs.FunctionAssurancePolicyScopeGetArgs>? _scopes;
        public InputList<Inputs.FunctionAssurancePolicyScopeGetArgs> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<Inputs.FunctionAssurancePolicyScopeGetArgs>());
            set => _scopes = value;
        }

        [Input("trustedBaseImages")]
        private InputList<Inputs.FunctionAssurancePolicyTrustedBaseImageGetArgs>? _trustedBaseImages;
        public InputList<Inputs.FunctionAssurancePolicyTrustedBaseImageGetArgs> TrustedBaseImages
        {
            get => _trustedBaseImages ?? (_trustedBaseImages = new InputList<Inputs.FunctionAssurancePolicyTrustedBaseImageGetArgs>());
            set => _trustedBaseImages = value;
        }

        [Input("trustedBaseImagesEnabled")]
        public Input<bool>? TrustedBaseImagesEnabled { get; set; }

        [Input("whitelistedLicenses")]
        private InputList<string>? _whitelistedLicenses;
        public InputList<string> WhitelistedLicenses
        {
            get => _whitelistedLicenses ?? (_whitelistedLicenses = new InputList<string>());
            set => _whitelistedLicenses = value;
        }

        [Input("whitelistedLicensesEnabled")]
        public Input<bool>? WhitelistedLicensesEnabled { get; set; }

        public FunctionAssurancePolicyState()
        {
        }
    }
}
