// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aquasec

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HostAssurancePolicy struct {
	pulumi.CustomResourceState

	AllowedImages                    pulumi.StringArrayOutput                     `pulumi:"allowedImages"`
	ApplicationScopes                pulumi.StringArrayOutput                     `pulumi:"applicationScopes"`
	AssuranceType                    pulumi.StringPtrOutput                       `pulumi:"assuranceType"`
	AuditOnFailure                   pulumi.BoolPtrOutput                         `pulumi:"auditOnFailure"`
	Author                           pulumi.StringOutput                          `pulumi:"author"`
	AutoScanConfigured               pulumi.BoolPtrOutput                         `pulumi:"autoScanConfigured"`
	AutoScanEnabled                  pulumi.BoolPtrOutput                         `pulumi:"autoScanEnabled"`
	AutoScanTimes                    HostAssurancePolicyAutoScanTimeArrayOutput   `pulumi:"autoScanTimes"`
	BlacklistPermissions             pulumi.StringArrayOutput                     `pulumi:"blacklistPermissions"`
	BlacklistPermissionsEnabled      pulumi.BoolPtrOutput                         `pulumi:"blacklistPermissionsEnabled"`
	BlacklistedLicenses              pulumi.StringArrayOutput                     `pulumi:"blacklistedLicenses"`
	BlacklistedLicensesEnabled       pulumi.BoolPtrOutput                         `pulumi:"blacklistedLicensesEnabled"`
	BlockFailed                      pulumi.BoolPtrOutput                         `pulumi:"blockFailed"`
	ControlExcludeNoFix              pulumi.BoolPtrOutput                         `pulumi:"controlExcludeNoFix"`
	CustomChecks                     HostAssurancePolicyCustomCheckArrayOutput    `pulumi:"customChecks"`
	CustomChecksEnabled              pulumi.BoolPtrOutput                         `pulumi:"customChecksEnabled"`
	CustomSeverityEnabled            pulumi.BoolPtrOutput                         `pulumi:"customSeverityEnabled"`
	CvesBlackListEnabled             pulumi.BoolPtrOutput                         `pulumi:"cvesBlackListEnabled"`
	CvesBlackLists                   pulumi.StringArrayOutput                     `pulumi:"cvesBlackLists"`
	CvesWhiteListEnabled             pulumi.BoolPtrOutput                         `pulumi:"cvesWhiteListEnabled"`
	CvesWhiteLists                   pulumi.StringArrayOutput                     `pulumi:"cvesWhiteLists"`
	CvssSeverity                     pulumi.StringPtrOutput                       `pulumi:"cvssSeverity"`
	CvssSeverityEnabled              pulumi.BoolPtrOutput                         `pulumi:"cvssSeverityEnabled"`
	CvssSeverityExcludeNoFix         pulumi.BoolPtrOutput                         `pulumi:"cvssSeverityExcludeNoFix"`
	Description                      pulumi.StringPtrOutput                       `pulumi:"description"`
	DisallowMalware                  pulumi.BoolPtrOutput                         `pulumi:"disallowMalware"`
	DockerCisEnabled                 pulumi.BoolPtrOutput                         `pulumi:"dockerCisEnabled"`
	Domain                           pulumi.StringPtrOutput                       `pulumi:"domain"`
	DomainName                       pulumi.StringPtrOutput                       `pulumi:"domainName"`
	DtaEnabled                       pulumi.BoolPtrOutput                         `pulumi:"dtaEnabled"`
	DtaSeverity                      pulumi.StringPtrOutput                       `pulumi:"dtaSeverity"`
	Enabled                          pulumi.BoolPtrOutput                         `pulumi:"enabled"`
	Enforce                          pulumi.BoolPtrOutput                         `pulumi:"enforce"`
	EnforceAfterDays                 pulumi.IntPtrOutput                          `pulumi:"enforceAfterDays"`
	EnforceExcessivePermissions      pulumi.BoolPtrOutput                         `pulumi:"enforceExcessivePermissions"`
	ExceptionalMonitoredMalwarePaths pulumi.StringArrayOutput                     `pulumi:"exceptionalMonitoredMalwarePaths"`
	FailCicd                         pulumi.BoolPtrOutput                         `pulumi:"failCicd"`
	ForbiddenLabels                  HostAssurancePolicyForbiddenLabelArrayOutput `pulumi:"forbiddenLabels"`
	ForbiddenLabelsEnabled           pulumi.BoolPtrOutput                         `pulumi:"forbiddenLabelsEnabled"`
	ForceMicroenforcer               pulumi.BoolPtrOutput                         `pulumi:"forceMicroenforcer"`
	FunctionIntegrityEnabled         pulumi.BoolPtrOutput                         `pulumi:"functionIntegrityEnabled"`
	// The ID of this resource.
	Id                               pulumi.StringOutput                             `pulumi:"id"`
	IgnoreRecentlyPublishedVln       pulumi.BoolPtrOutput                            `pulumi:"ignoreRecentlyPublishedVln"`
	IgnoreRecentlyPublishedVlnPeriod pulumi.IntOutput                                `pulumi:"ignoreRecentlyPublishedVlnPeriod"`
	IgnoreRiskResourcesEnabled       pulumi.BoolPtrOutput                            `pulumi:"ignoreRiskResourcesEnabled"`
	IgnoredRiskResources             pulumi.StringArrayOutput                        `pulumi:"ignoredRiskResources"`
	Images                           pulumi.StringArrayOutput                        `pulumi:"images"`
	KubeCisEnabled                   pulumi.BoolPtrOutput                            `pulumi:"kubeCisEnabled"`
	Labels                           pulumi.StringArrayOutput                        `pulumi:"labels"`
	MalwareAction                    pulumi.StringPtrOutput                          `pulumi:"malwareAction"`
	MaximumScore                     pulumi.Float64PtrOutput                         `pulumi:"maximumScore"`
	MaximumScoreEnabled              pulumi.BoolPtrOutput                            `pulumi:"maximumScoreEnabled"`
	MaximumScoreExcludeNoFix         pulumi.BoolPtrOutput                            `pulumi:"maximumScoreExcludeNoFix"`
	MonitoredMalwarePaths            pulumi.StringArrayOutput                        `pulumi:"monitoredMalwarePaths"`
	Name                             pulumi.StringOutput                             `pulumi:"name"`
	OnlyNoneRootUsers                pulumi.BoolPtrOutput                            `pulumi:"onlyNoneRootUsers"`
	PackagesBlackListEnabled         pulumi.BoolPtrOutput                            `pulumi:"packagesBlackListEnabled"`
	PackagesBlackLists               HostAssurancePolicyPackagesBlackListArrayOutput `pulumi:"packagesBlackLists"`
	PackagesWhiteListEnabled         pulumi.BoolPtrOutput                            `pulumi:"packagesWhiteListEnabled"`
	PackagesWhiteLists               HostAssurancePolicyPackagesWhiteListArrayOutput `pulumi:"packagesWhiteLists"`
	PartialResultsImageFail          pulumi.BoolPtrOutput                            `pulumi:"partialResultsImageFail"`
	ReadOnly                         pulumi.BoolPtrOutput                            `pulumi:"readOnly"`
	Registries                       pulumi.StringArrayOutput                        `pulumi:"registries"`
	Registry                         pulumi.StringPtrOutput                          `pulumi:"registry"`
	RequiredLabels                   HostAssurancePolicyRequiredLabelArrayOutput     `pulumi:"requiredLabels"`
	RequiredLabelsEnabled            pulumi.BoolPtrOutput                            `pulumi:"requiredLabelsEnabled"`
	ScanNfsMounts                    pulumi.BoolPtrOutput                            `pulumi:"scanNfsMounts"`
	ScanSensitiveData                pulumi.BoolPtrOutput                            `pulumi:"scanSensitiveData"`
	ScapEnabled                      pulumi.BoolPtrOutput                            `pulumi:"scapEnabled"`
	ScapFiles                        pulumi.StringArrayOutput                        `pulumi:"scapFiles"`
	Scopes                           HostAssurancePolicyScopeArrayOutput             `pulumi:"scopes"`
	TrustedBaseImages                HostAssurancePolicyTrustedBaseImageArrayOutput  `pulumi:"trustedBaseImages"`
	TrustedBaseImagesEnabled         pulumi.BoolPtrOutput                            `pulumi:"trustedBaseImagesEnabled"`
	WhitelistedLicenses              pulumi.StringArrayOutput                        `pulumi:"whitelistedLicenses"`
	WhitelistedLicensesEnabled       pulumi.BoolPtrOutput                            `pulumi:"whitelistedLicensesEnabled"`
}

// NewHostAssurancePolicy registers a new resource with the given unique name, arguments, and options.
func NewHostAssurancePolicy(ctx *pulumi.Context,
	name string, args *HostAssurancePolicyArgs, opts ...pulumi.ResourceOption) (*HostAssurancePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationScopes == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationScopes'")
	}
	var resource HostAssurancePolicy
	err := ctx.RegisterResource("aquasec:index/hostAssurancePolicy:HostAssurancePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostAssurancePolicy gets an existing HostAssurancePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostAssurancePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostAssurancePolicyState, opts ...pulumi.ResourceOption) (*HostAssurancePolicy, error) {
	var resource HostAssurancePolicy
	err := ctx.ReadResource("aquasec:index/hostAssurancePolicy:HostAssurancePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostAssurancePolicy resources.
type hostAssurancePolicyState struct {
	AllowedImages                    []string                            `pulumi:"allowedImages"`
	ApplicationScopes                []string                            `pulumi:"applicationScopes"`
	AssuranceType                    *string                             `pulumi:"assuranceType"`
	AuditOnFailure                   *bool                               `pulumi:"auditOnFailure"`
	Author                           *string                             `pulumi:"author"`
	AutoScanConfigured               *bool                               `pulumi:"autoScanConfigured"`
	AutoScanEnabled                  *bool                               `pulumi:"autoScanEnabled"`
	AutoScanTimes                    []HostAssurancePolicyAutoScanTime   `pulumi:"autoScanTimes"`
	BlacklistPermissions             []string                            `pulumi:"blacklistPermissions"`
	BlacklistPermissionsEnabled      *bool                               `pulumi:"blacklistPermissionsEnabled"`
	BlacklistedLicenses              []string                            `pulumi:"blacklistedLicenses"`
	BlacklistedLicensesEnabled       *bool                               `pulumi:"blacklistedLicensesEnabled"`
	BlockFailed                      *bool                               `pulumi:"blockFailed"`
	ControlExcludeNoFix              *bool                               `pulumi:"controlExcludeNoFix"`
	CustomChecks                     []HostAssurancePolicyCustomCheck    `pulumi:"customChecks"`
	CustomChecksEnabled              *bool                               `pulumi:"customChecksEnabled"`
	CustomSeverityEnabled            *bool                               `pulumi:"customSeverityEnabled"`
	CvesBlackListEnabled             *bool                               `pulumi:"cvesBlackListEnabled"`
	CvesBlackLists                   []string                            `pulumi:"cvesBlackLists"`
	CvesWhiteListEnabled             *bool                               `pulumi:"cvesWhiteListEnabled"`
	CvesWhiteLists                   []string                            `pulumi:"cvesWhiteLists"`
	CvssSeverity                     *string                             `pulumi:"cvssSeverity"`
	CvssSeverityEnabled              *bool                               `pulumi:"cvssSeverityEnabled"`
	CvssSeverityExcludeNoFix         *bool                               `pulumi:"cvssSeverityExcludeNoFix"`
	Description                      *string                             `pulumi:"description"`
	DisallowMalware                  *bool                               `pulumi:"disallowMalware"`
	DockerCisEnabled                 *bool                               `pulumi:"dockerCisEnabled"`
	Domain                           *string                             `pulumi:"domain"`
	DomainName                       *string                             `pulumi:"domainName"`
	DtaEnabled                       *bool                               `pulumi:"dtaEnabled"`
	DtaSeverity                      *string                             `pulumi:"dtaSeverity"`
	Enabled                          *bool                               `pulumi:"enabled"`
	Enforce                          *bool                               `pulumi:"enforce"`
	EnforceAfterDays                 *int                                `pulumi:"enforceAfterDays"`
	EnforceExcessivePermissions      *bool                               `pulumi:"enforceExcessivePermissions"`
	ExceptionalMonitoredMalwarePaths []string                            `pulumi:"exceptionalMonitoredMalwarePaths"`
	FailCicd                         *bool                               `pulumi:"failCicd"`
	ForbiddenLabels                  []HostAssurancePolicyForbiddenLabel `pulumi:"forbiddenLabels"`
	ForbiddenLabelsEnabled           *bool                               `pulumi:"forbiddenLabelsEnabled"`
	ForceMicroenforcer               *bool                               `pulumi:"forceMicroenforcer"`
	FunctionIntegrityEnabled         *bool                               `pulumi:"functionIntegrityEnabled"`
	// The ID of this resource.
	Id                               *string                                `pulumi:"id"`
	IgnoreRecentlyPublishedVln       *bool                                  `pulumi:"ignoreRecentlyPublishedVln"`
	IgnoreRecentlyPublishedVlnPeriod *int                                   `pulumi:"ignoreRecentlyPublishedVlnPeriod"`
	IgnoreRiskResourcesEnabled       *bool                                  `pulumi:"ignoreRiskResourcesEnabled"`
	IgnoredRiskResources             []string                               `pulumi:"ignoredRiskResources"`
	Images                           []string                               `pulumi:"images"`
	KubeCisEnabled                   *bool                                  `pulumi:"kubeCisEnabled"`
	Labels                           []string                               `pulumi:"labels"`
	MalwareAction                    *string                                `pulumi:"malwareAction"`
	MaximumScore                     *float64                               `pulumi:"maximumScore"`
	MaximumScoreEnabled              *bool                                  `pulumi:"maximumScoreEnabled"`
	MaximumScoreExcludeNoFix         *bool                                  `pulumi:"maximumScoreExcludeNoFix"`
	MonitoredMalwarePaths            []string                               `pulumi:"monitoredMalwarePaths"`
	Name                             *string                                `pulumi:"name"`
	OnlyNoneRootUsers                *bool                                  `pulumi:"onlyNoneRootUsers"`
	PackagesBlackListEnabled         *bool                                  `pulumi:"packagesBlackListEnabled"`
	PackagesBlackLists               []HostAssurancePolicyPackagesBlackList `pulumi:"packagesBlackLists"`
	PackagesWhiteListEnabled         *bool                                  `pulumi:"packagesWhiteListEnabled"`
	PackagesWhiteLists               []HostAssurancePolicyPackagesWhiteList `pulumi:"packagesWhiteLists"`
	PartialResultsImageFail          *bool                                  `pulumi:"partialResultsImageFail"`
	ReadOnly                         *bool                                  `pulumi:"readOnly"`
	Registries                       []string                               `pulumi:"registries"`
	Registry                         *string                                `pulumi:"registry"`
	RequiredLabels                   []HostAssurancePolicyRequiredLabel     `pulumi:"requiredLabels"`
	RequiredLabelsEnabled            *bool                                  `pulumi:"requiredLabelsEnabled"`
	ScanNfsMounts                    *bool                                  `pulumi:"scanNfsMounts"`
	ScanSensitiveData                *bool                                  `pulumi:"scanSensitiveData"`
	ScapEnabled                      *bool                                  `pulumi:"scapEnabled"`
	ScapFiles                        []string                               `pulumi:"scapFiles"`
	Scopes                           []HostAssurancePolicyScope             `pulumi:"scopes"`
	TrustedBaseImages                []HostAssurancePolicyTrustedBaseImage  `pulumi:"trustedBaseImages"`
	TrustedBaseImagesEnabled         *bool                                  `pulumi:"trustedBaseImagesEnabled"`
	WhitelistedLicenses              []string                               `pulumi:"whitelistedLicenses"`
	WhitelistedLicensesEnabled       *bool                                  `pulumi:"whitelistedLicensesEnabled"`
}

type HostAssurancePolicyState struct {
	AllowedImages                    pulumi.StringArrayInput
	ApplicationScopes                pulumi.StringArrayInput
	AssuranceType                    pulumi.StringPtrInput
	AuditOnFailure                   pulumi.BoolPtrInput
	Author                           pulumi.StringPtrInput
	AutoScanConfigured               pulumi.BoolPtrInput
	AutoScanEnabled                  pulumi.BoolPtrInput
	AutoScanTimes                    HostAssurancePolicyAutoScanTimeArrayInput
	BlacklistPermissions             pulumi.StringArrayInput
	BlacklistPermissionsEnabled      pulumi.BoolPtrInput
	BlacklistedLicenses              pulumi.StringArrayInput
	BlacklistedLicensesEnabled       pulumi.BoolPtrInput
	BlockFailed                      pulumi.BoolPtrInput
	ControlExcludeNoFix              pulumi.BoolPtrInput
	CustomChecks                     HostAssurancePolicyCustomCheckArrayInput
	CustomChecksEnabled              pulumi.BoolPtrInput
	CustomSeverityEnabled            pulumi.BoolPtrInput
	CvesBlackListEnabled             pulumi.BoolPtrInput
	CvesBlackLists                   pulumi.StringArrayInput
	CvesWhiteListEnabled             pulumi.BoolPtrInput
	CvesWhiteLists                   pulumi.StringArrayInput
	CvssSeverity                     pulumi.StringPtrInput
	CvssSeverityEnabled              pulumi.BoolPtrInput
	CvssSeverityExcludeNoFix         pulumi.BoolPtrInput
	Description                      pulumi.StringPtrInput
	DisallowMalware                  pulumi.BoolPtrInput
	DockerCisEnabled                 pulumi.BoolPtrInput
	Domain                           pulumi.StringPtrInput
	DomainName                       pulumi.StringPtrInput
	DtaEnabled                       pulumi.BoolPtrInput
	DtaSeverity                      pulumi.StringPtrInput
	Enabled                          pulumi.BoolPtrInput
	Enforce                          pulumi.BoolPtrInput
	EnforceAfterDays                 pulumi.IntPtrInput
	EnforceExcessivePermissions      pulumi.BoolPtrInput
	ExceptionalMonitoredMalwarePaths pulumi.StringArrayInput
	FailCicd                         pulumi.BoolPtrInput
	ForbiddenLabels                  HostAssurancePolicyForbiddenLabelArrayInput
	ForbiddenLabelsEnabled           pulumi.BoolPtrInput
	ForceMicroenforcer               pulumi.BoolPtrInput
	FunctionIntegrityEnabled         pulumi.BoolPtrInput
	// The ID of this resource.
	Id                               pulumi.StringPtrInput
	IgnoreRecentlyPublishedVln       pulumi.BoolPtrInput
	IgnoreRecentlyPublishedVlnPeriod pulumi.IntPtrInput
	IgnoreRiskResourcesEnabled       pulumi.BoolPtrInput
	IgnoredRiskResources             pulumi.StringArrayInput
	Images                           pulumi.StringArrayInput
	KubeCisEnabled                   pulumi.BoolPtrInput
	Labels                           pulumi.StringArrayInput
	MalwareAction                    pulumi.StringPtrInput
	MaximumScore                     pulumi.Float64PtrInput
	MaximumScoreEnabled              pulumi.BoolPtrInput
	MaximumScoreExcludeNoFix         pulumi.BoolPtrInput
	MonitoredMalwarePaths            pulumi.StringArrayInput
	Name                             pulumi.StringPtrInput
	OnlyNoneRootUsers                pulumi.BoolPtrInput
	PackagesBlackListEnabled         pulumi.BoolPtrInput
	PackagesBlackLists               HostAssurancePolicyPackagesBlackListArrayInput
	PackagesWhiteListEnabled         pulumi.BoolPtrInput
	PackagesWhiteLists               HostAssurancePolicyPackagesWhiteListArrayInput
	PartialResultsImageFail          pulumi.BoolPtrInput
	ReadOnly                         pulumi.BoolPtrInput
	Registries                       pulumi.StringArrayInput
	Registry                         pulumi.StringPtrInput
	RequiredLabels                   HostAssurancePolicyRequiredLabelArrayInput
	RequiredLabelsEnabled            pulumi.BoolPtrInput
	ScanNfsMounts                    pulumi.BoolPtrInput
	ScanSensitiveData                pulumi.BoolPtrInput
	ScapEnabled                      pulumi.BoolPtrInput
	ScapFiles                        pulumi.StringArrayInput
	Scopes                           HostAssurancePolicyScopeArrayInput
	TrustedBaseImages                HostAssurancePolicyTrustedBaseImageArrayInput
	TrustedBaseImagesEnabled         pulumi.BoolPtrInput
	WhitelistedLicenses              pulumi.StringArrayInput
	WhitelistedLicensesEnabled       pulumi.BoolPtrInput
}

func (HostAssurancePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostAssurancePolicyState)(nil)).Elem()
}

type hostAssurancePolicyArgs struct {
	AllowedImages                    []string                               `pulumi:"allowedImages"`
	ApplicationScopes                []string                               `pulumi:"applicationScopes"`
	AssuranceType                    *string                                `pulumi:"assuranceType"`
	AuditOnFailure                   *bool                                  `pulumi:"auditOnFailure"`
	AutoScanConfigured               *bool                                  `pulumi:"autoScanConfigured"`
	AutoScanEnabled                  *bool                                  `pulumi:"autoScanEnabled"`
	AutoScanTimes                    []HostAssurancePolicyAutoScanTime      `pulumi:"autoScanTimes"`
	BlacklistPermissions             []string                               `pulumi:"blacklistPermissions"`
	BlacklistPermissionsEnabled      *bool                                  `pulumi:"blacklistPermissionsEnabled"`
	BlacklistedLicenses              []string                               `pulumi:"blacklistedLicenses"`
	BlacklistedLicensesEnabled       *bool                                  `pulumi:"blacklistedLicensesEnabled"`
	BlockFailed                      *bool                                  `pulumi:"blockFailed"`
	ControlExcludeNoFix              *bool                                  `pulumi:"controlExcludeNoFix"`
	CustomChecks                     []HostAssurancePolicyCustomCheck       `pulumi:"customChecks"`
	CustomChecksEnabled              *bool                                  `pulumi:"customChecksEnabled"`
	CustomSeverityEnabled            *bool                                  `pulumi:"customSeverityEnabled"`
	CvesBlackListEnabled             *bool                                  `pulumi:"cvesBlackListEnabled"`
	CvesBlackLists                   []string                               `pulumi:"cvesBlackLists"`
	CvesWhiteListEnabled             *bool                                  `pulumi:"cvesWhiteListEnabled"`
	CvesWhiteLists                   []string                               `pulumi:"cvesWhiteLists"`
	CvssSeverity                     *string                                `pulumi:"cvssSeverity"`
	CvssSeverityEnabled              *bool                                  `pulumi:"cvssSeverityEnabled"`
	CvssSeverityExcludeNoFix         *bool                                  `pulumi:"cvssSeverityExcludeNoFix"`
	Description                      *string                                `pulumi:"description"`
	DisallowMalware                  *bool                                  `pulumi:"disallowMalware"`
	DockerCisEnabled                 *bool                                  `pulumi:"dockerCisEnabled"`
	Domain                           *string                                `pulumi:"domain"`
	DomainName                       *string                                `pulumi:"domainName"`
	DtaEnabled                       *bool                                  `pulumi:"dtaEnabled"`
	DtaSeverity                      *string                                `pulumi:"dtaSeverity"`
	Enabled                          *bool                                  `pulumi:"enabled"`
	Enforce                          *bool                                  `pulumi:"enforce"`
	EnforceAfterDays                 *int                                   `pulumi:"enforceAfterDays"`
	EnforceExcessivePermissions      *bool                                  `pulumi:"enforceExcessivePermissions"`
	ExceptionalMonitoredMalwarePaths []string                               `pulumi:"exceptionalMonitoredMalwarePaths"`
	FailCicd                         *bool                                  `pulumi:"failCicd"`
	ForbiddenLabels                  []HostAssurancePolicyForbiddenLabel    `pulumi:"forbiddenLabels"`
	ForbiddenLabelsEnabled           *bool                                  `pulumi:"forbiddenLabelsEnabled"`
	ForceMicroenforcer               *bool                                  `pulumi:"forceMicroenforcer"`
	FunctionIntegrityEnabled         *bool                                  `pulumi:"functionIntegrityEnabled"`
	IgnoreRecentlyPublishedVln       *bool                                  `pulumi:"ignoreRecentlyPublishedVln"`
	IgnoreRiskResourcesEnabled       *bool                                  `pulumi:"ignoreRiskResourcesEnabled"`
	IgnoredRiskResources             []string                               `pulumi:"ignoredRiskResources"`
	Images                           []string                               `pulumi:"images"`
	KubeCisEnabled                   *bool                                  `pulumi:"kubeCisEnabled"`
	Labels                           []string                               `pulumi:"labels"`
	MalwareAction                    *string                                `pulumi:"malwareAction"`
	MaximumScore                     *float64                               `pulumi:"maximumScore"`
	MaximumScoreEnabled              *bool                                  `pulumi:"maximumScoreEnabled"`
	MaximumScoreExcludeNoFix         *bool                                  `pulumi:"maximumScoreExcludeNoFix"`
	MonitoredMalwarePaths            []string                               `pulumi:"monitoredMalwarePaths"`
	Name                             *string                                `pulumi:"name"`
	OnlyNoneRootUsers                *bool                                  `pulumi:"onlyNoneRootUsers"`
	PackagesBlackListEnabled         *bool                                  `pulumi:"packagesBlackListEnabled"`
	PackagesBlackLists               []HostAssurancePolicyPackagesBlackList `pulumi:"packagesBlackLists"`
	PackagesWhiteListEnabled         *bool                                  `pulumi:"packagesWhiteListEnabled"`
	PackagesWhiteLists               []HostAssurancePolicyPackagesWhiteList `pulumi:"packagesWhiteLists"`
	PartialResultsImageFail          *bool                                  `pulumi:"partialResultsImageFail"`
	ReadOnly                         *bool                                  `pulumi:"readOnly"`
	Registries                       []string                               `pulumi:"registries"`
	Registry                         *string                                `pulumi:"registry"`
	RequiredLabels                   []HostAssurancePolicyRequiredLabel     `pulumi:"requiredLabels"`
	RequiredLabelsEnabled            *bool                                  `pulumi:"requiredLabelsEnabled"`
	ScanNfsMounts                    *bool                                  `pulumi:"scanNfsMounts"`
	ScanSensitiveData                *bool                                  `pulumi:"scanSensitiveData"`
	ScapEnabled                      *bool                                  `pulumi:"scapEnabled"`
	ScapFiles                        []string                               `pulumi:"scapFiles"`
	Scopes                           []HostAssurancePolicyScope             `pulumi:"scopes"`
	TrustedBaseImages                []HostAssurancePolicyTrustedBaseImage  `pulumi:"trustedBaseImages"`
	TrustedBaseImagesEnabled         *bool                                  `pulumi:"trustedBaseImagesEnabled"`
	WhitelistedLicenses              []string                               `pulumi:"whitelistedLicenses"`
	WhitelistedLicensesEnabled       *bool                                  `pulumi:"whitelistedLicensesEnabled"`
}

// The set of arguments for constructing a HostAssurancePolicy resource.
type HostAssurancePolicyArgs struct {
	AllowedImages                    pulumi.StringArrayInput
	ApplicationScopes                pulumi.StringArrayInput
	AssuranceType                    pulumi.StringPtrInput
	AuditOnFailure                   pulumi.BoolPtrInput
	AutoScanConfigured               pulumi.BoolPtrInput
	AutoScanEnabled                  pulumi.BoolPtrInput
	AutoScanTimes                    HostAssurancePolicyAutoScanTimeArrayInput
	BlacklistPermissions             pulumi.StringArrayInput
	BlacklistPermissionsEnabled      pulumi.BoolPtrInput
	BlacklistedLicenses              pulumi.StringArrayInput
	BlacklistedLicensesEnabled       pulumi.BoolPtrInput
	BlockFailed                      pulumi.BoolPtrInput
	ControlExcludeNoFix              pulumi.BoolPtrInput
	CustomChecks                     HostAssurancePolicyCustomCheckArrayInput
	CustomChecksEnabled              pulumi.BoolPtrInput
	CustomSeverityEnabled            pulumi.BoolPtrInput
	CvesBlackListEnabled             pulumi.BoolPtrInput
	CvesBlackLists                   pulumi.StringArrayInput
	CvesWhiteListEnabled             pulumi.BoolPtrInput
	CvesWhiteLists                   pulumi.StringArrayInput
	CvssSeverity                     pulumi.StringPtrInput
	CvssSeverityEnabled              pulumi.BoolPtrInput
	CvssSeverityExcludeNoFix         pulumi.BoolPtrInput
	Description                      pulumi.StringPtrInput
	DisallowMalware                  pulumi.BoolPtrInput
	DockerCisEnabled                 pulumi.BoolPtrInput
	Domain                           pulumi.StringPtrInput
	DomainName                       pulumi.StringPtrInput
	DtaEnabled                       pulumi.BoolPtrInput
	DtaSeverity                      pulumi.StringPtrInput
	Enabled                          pulumi.BoolPtrInput
	Enforce                          pulumi.BoolPtrInput
	EnforceAfterDays                 pulumi.IntPtrInput
	EnforceExcessivePermissions      pulumi.BoolPtrInput
	ExceptionalMonitoredMalwarePaths pulumi.StringArrayInput
	FailCicd                         pulumi.BoolPtrInput
	ForbiddenLabels                  HostAssurancePolicyForbiddenLabelArrayInput
	ForbiddenLabelsEnabled           pulumi.BoolPtrInput
	ForceMicroenforcer               pulumi.BoolPtrInput
	FunctionIntegrityEnabled         pulumi.BoolPtrInput
	IgnoreRecentlyPublishedVln       pulumi.BoolPtrInput
	IgnoreRiskResourcesEnabled       pulumi.BoolPtrInput
	IgnoredRiskResources             pulumi.StringArrayInput
	Images                           pulumi.StringArrayInput
	KubeCisEnabled                   pulumi.BoolPtrInput
	Labels                           pulumi.StringArrayInput
	MalwareAction                    pulumi.StringPtrInput
	MaximumScore                     pulumi.Float64PtrInput
	MaximumScoreEnabled              pulumi.BoolPtrInput
	MaximumScoreExcludeNoFix         pulumi.BoolPtrInput
	MonitoredMalwarePaths            pulumi.StringArrayInput
	Name                             pulumi.StringPtrInput
	OnlyNoneRootUsers                pulumi.BoolPtrInput
	PackagesBlackListEnabled         pulumi.BoolPtrInput
	PackagesBlackLists               HostAssurancePolicyPackagesBlackListArrayInput
	PackagesWhiteListEnabled         pulumi.BoolPtrInput
	PackagesWhiteLists               HostAssurancePolicyPackagesWhiteListArrayInput
	PartialResultsImageFail          pulumi.BoolPtrInput
	ReadOnly                         pulumi.BoolPtrInput
	Registries                       pulumi.StringArrayInput
	Registry                         pulumi.StringPtrInput
	RequiredLabels                   HostAssurancePolicyRequiredLabelArrayInput
	RequiredLabelsEnabled            pulumi.BoolPtrInput
	ScanNfsMounts                    pulumi.BoolPtrInput
	ScanSensitiveData                pulumi.BoolPtrInput
	ScapEnabled                      pulumi.BoolPtrInput
	ScapFiles                        pulumi.StringArrayInput
	Scopes                           HostAssurancePolicyScopeArrayInput
	TrustedBaseImages                HostAssurancePolicyTrustedBaseImageArrayInput
	TrustedBaseImagesEnabled         pulumi.BoolPtrInput
	WhitelistedLicenses              pulumi.StringArrayInput
	WhitelistedLicensesEnabled       pulumi.BoolPtrInput
}

func (HostAssurancePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostAssurancePolicyArgs)(nil)).Elem()
}

type HostAssurancePolicyInput interface {
	pulumi.Input

	ToHostAssurancePolicyOutput() HostAssurancePolicyOutput
	ToHostAssurancePolicyOutputWithContext(ctx context.Context) HostAssurancePolicyOutput
}

func (*HostAssurancePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**HostAssurancePolicy)(nil)).Elem()
}

func (i *HostAssurancePolicy) ToHostAssurancePolicyOutput() HostAssurancePolicyOutput {
	return i.ToHostAssurancePolicyOutputWithContext(context.Background())
}

func (i *HostAssurancePolicy) ToHostAssurancePolicyOutputWithContext(ctx context.Context) HostAssurancePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostAssurancePolicyOutput)
}

// HostAssurancePolicyArrayInput is an input type that accepts HostAssurancePolicyArray and HostAssurancePolicyArrayOutput values.
// You can construct a concrete instance of `HostAssurancePolicyArrayInput` via:
//
//          HostAssurancePolicyArray{ HostAssurancePolicyArgs{...} }
type HostAssurancePolicyArrayInput interface {
	pulumi.Input

	ToHostAssurancePolicyArrayOutput() HostAssurancePolicyArrayOutput
	ToHostAssurancePolicyArrayOutputWithContext(context.Context) HostAssurancePolicyArrayOutput
}

type HostAssurancePolicyArray []HostAssurancePolicyInput

func (HostAssurancePolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostAssurancePolicy)(nil)).Elem()
}

func (i HostAssurancePolicyArray) ToHostAssurancePolicyArrayOutput() HostAssurancePolicyArrayOutput {
	return i.ToHostAssurancePolicyArrayOutputWithContext(context.Background())
}

func (i HostAssurancePolicyArray) ToHostAssurancePolicyArrayOutputWithContext(ctx context.Context) HostAssurancePolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostAssurancePolicyArrayOutput)
}

// HostAssurancePolicyMapInput is an input type that accepts HostAssurancePolicyMap and HostAssurancePolicyMapOutput values.
// You can construct a concrete instance of `HostAssurancePolicyMapInput` via:
//
//          HostAssurancePolicyMap{ "key": HostAssurancePolicyArgs{...} }
type HostAssurancePolicyMapInput interface {
	pulumi.Input

	ToHostAssurancePolicyMapOutput() HostAssurancePolicyMapOutput
	ToHostAssurancePolicyMapOutputWithContext(context.Context) HostAssurancePolicyMapOutput
}

type HostAssurancePolicyMap map[string]HostAssurancePolicyInput

func (HostAssurancePolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostAssurancePolicy)(nil)).Elem()
}

func (i HostAssurancePolicyMap) ToHostAssurancePolicyMapOutput() HostAssurancePolicyMapOutput {
	return i.ToHostAssurancePolicyMapOutputWithContext(context.Background())
}

func (i HostAssurancePolicyMap) ToHostAssurancePolicyMapOutputWithContext(ctx context.Context) HostAssurancePolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostAssurancePolicyMapOutput)
}

type HostAssurancePolicyOutput struct{ *pulumi.OutputState }

func (HostAssurancePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostAssurancePolicy)(nil)).Elem()
}

func (o HostAssurancePolicyOutput) ToHostAssurancePolicyOutput() HostAssurancePolicyOutput {
	return o
}

func (o HostAssurancePolicyOutput) ToHostAssurancePolicyOutputWithContext(ctx context.Context) HostAssurancePolicyOutput {
	return o
}

type HostAssurancePolicyArrayOutput struct{ *pulumi.OutputState }

func (HostAssurancePolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostAssurancePolicy)(nil)).Elem()
}

func (o HostAssurancePolicyArrayOutput) ToHostAssurancePolicyArrayOutput() HostAssurancePolicyArrayOutput {
	return o
}

func (o HostAssurancePolicyArrayOutput) ToHostAssurancePolicyArrayOutputWithContext(ctx context.Context) HostAssurancePolicyArrayOutput {
	return o
}

func (o HostAssurancePolicyArrayOutput) Index(i pulumi.IntInput) HostAssurancePolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HostAssurancePolicy {
		return vs[0].([]*HostAssurancePolicy)[vs[1].(int)]
	}).(HostAssurancePolicyOutput)
}

type HostAssurancePolicyMapOutput struct{ *pulumi.OutputState }

func (HostAssurancePolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostAssurancePolicy)(nil)).Elem()
}

func (o HostAssurancePolicyMapOutput) ToHostAssurancePolicyMapOutput() HostAssurancePolicyMapOutput {
	return o
}

func (o HostAssurancePolicyMapOutput) ToHostAssurancePolicyMapOutputWithContext(ctx context.Context) HostAssurancePolicyMapOutput {
	return o
}

func (o HostAssurancePolicyMapOutput) MapIndex(k pulumi.StringInput) HostAssurancePolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HostAssurancePolicy {
		return vs[0].(map[string]*HostAssurancePolicy)[vs[1].(string)]
	}).(HostAssurancePolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HostAssurancePolicyInput)(nil)).Elem(), &HostAssurancePolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostAssurancePolicyArrayInput)(nil)).Elem(), HostAssurancePolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostAssurancePolicyMapInput)(nil)).Elem(), HostAssurancePolicyMap{})
	pulumi.RegisterOutputType(HostAssurancePolicyOutput{})
	pulumi.RegisterOutputType(HostAssurancePolicyArrayOutput{})
	pulumi.RegisterOutputType(HostAssurancePolicyMapOutput{})
}
