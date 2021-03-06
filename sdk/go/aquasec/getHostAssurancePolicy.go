// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aquasec

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHostAssurancePolicy(ctx *pulumi.Context, args *LookupHostAssurancePolicyArgs, opts ...pulumi.InvokeOption) (*LookupHostAssurancePolicyResult, error) {
	var rv LookupHostAssurancePolicyResult
	err := ctx.Invoke("aquasec:index/getHostAssurancePolicy:getHostAssurancePolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getHostAssurancePolicy.
type LookupHostAssurancePolicyArgs struct {
	MaximumScoreExcludeNoFix *bool  `pulumi:"maximumScoreExcludeNoFix"`
	Name                     string `pulumi:"name"`
}

// A collection of values returned by getHostAssurancePolicy.
type LookupHostAssurancePolicyResult struct {
	AllowedImages                    []string                               `pulumi:"allowedImages"`
	ApplicationScopes                []string                               `pulumi:"applicationScopes"`
	AssuranceType                    string                                 `pulumi:"assuranceType"`
	AuditOnFailure                   bool                                   `pulumi:"auditOnFailure"`
	Author                           string                                 `pulumi:"author"`
	AutoScanConfigured               bool                                   `pulumi:"autoScanConfigured"`
	AutoScanEnabled                  bool                                   `pulumi:"autoScanEnabled"`
	AutoScanTimes                    []GetHostAssurancePolicyAutoScanTime   `pulumi:"autoScanTimes"`
	BlacklistPermissions             []string                               `pulumi:"blacklistPermissions"`
	BlacklistPermissionsEnabled      bool                                   `pulumi:"blacklistPermissionsEnabled"`
	BlacklistedLicenses              []string                               `pulumi:"blacklistedLicenses"`
	BlacklistedLicensesEnabled       bool                                   `pulumi:"blacklistedLicensesEnabled"`
	BlockFailed                      bool                                   `pulumi:"blockFailed"`
	ControlExcludeNoFix              bool                                   `pulumi:"controlExcludeNoFix"`
	CustomChecks                     []GetHostAssurancePolicyCustomCheck    `pulumi:"customChecks"`
	CustomChecksEnabled              bool                                   `pulumi:"customChecksEnabled"`
	CustomSeverityEnabled            bool                                   `pulumi:"customSeverityEnabled"`
	CvesBlackListEnabled             bool                                   `pulumi:"cvesBlackListEnabled"`
	CvesBlackLists                   []string                               `pulumi:"cvesBlackLists"`
	CvesWhiteListEnabled             bool                                   `pulumi:"cvesWhiteListEnabled"`
	CvesWhiteLists                   []string                               `pulumi:"cvesWhiteLists"`
	CvssSeverity                     string                                 `pulumi:"cvssSeverity"`
	CvssSeverityEnabled              bool                                   `pulumi:"cvssSeverityEnabled"`
	CvssSeverityExcludeNoFix         bool                                   `pulumi:"cvssSeverityExcludeNoFix"`
	Description                      string                                 `pulumi:"description"`
	DisallowMalware                  bool                                   `pulumi:"disallowMalware"`
	DockerCisEnabled                 bool                                   `pulumi:"dockerCisEnabled"`
	Domain                           string                                 `pulumi:"domain"`
	DomainName                       string                                 `pulumi:"domainName"`
	DtaEnabled                       bool                                   `pulumi:"dtaEnabled"`
	DtaSeverity                      string                                 `pulumi:"dtaSeverity"`
	Enabled                          bool                                   `pulumi:"enabled"`
	Enforce                          bool                                   `pulumi:"enforce"`
	EnforceAfterDays                 int                                    `pulumi:"enforceAfterDays"`
	EnforceExcessivePermissions      bool                                   `pulumi:"enforceExcessivePermissions"`
	ExceptionalMonitoredMalwarePaths []string                               `pulumi:"exceptionalMonitoredMalwarePaths"`
	FailCicd                         bool                                   `pulumi:"failCicd"`
	ForbiddenLabels                  []GetHostAssurancePolicyForbiddenLabel `pulumi:"forbiddenLabels"`
	ForbiddenLabelsEnabled           bool                                   `pulumi:"forbiddenLabelsEnabled"`
	ForceMicroenforcer               bool                                   `pulumi:"forceMicroenforcer"`
	FunctionIntegrityEnabled         bool                                   `pulumi:"functionIntegrityEnabled"`
	// The ID of this resource.
	Id                               string                                    `pulumi:"id"`
	IgnoreRecentlyPublishedVln       bool                                      `pulumi:"ignoreRecentlyPublishedVln"`
	IgnoreRecentlyPublishedVlnPeriod int                                       `pulumi:"ignoreRecentlyPublishedVlnPeriod"`
	IgnoreRiskResourcesEnabled       bool                                      `pulumi:"ignoreRiskResourcesEnabled"`
	IgnoredRiskResources             []string                                  `pulumi:"ignoredRiskResources"`
	Images                           []string                                  `pulumi:"images"`
	KubeCisEnabled                   bool                                      `pulumi:"kubeCisEnabled"`
	Labels                           []string                                  `pulumi:"labels"`
	MalwareAction                    string                                    `pulumi:"malwareAction"`
	MaximumScore                     float64                                   `pulumi:"maximumScore"`
	MaximumScoreEnabled              bool                                      `pulumi:"maximumScoreEnabled"`
	MaximumScoreExcludeNoFix         *bool                                     `pulumi:"maximumScoreExcludeNoFix"`
	MonitoredMalwarePaths            []string                                  `pulumi:"monitoredMalwarePaths"`
	Name                             string                                    `pulumi:"name"`
	OnlyNoneRootUsers                bool                                      `pulumi:"onlyNoneRootUsers"`
	PackagesBlackListEnabled         bool                                      `pulumi:"packagesBlackListEnabled"`
	PackagesBlackLists               []GetHostAssurancePolicyPackagesBlackList `pulumi:"packagesBlackLists"`
	PackagesWhiteListEnabled         bool                                      `pulumi:"packagesWhiteListEnabled"`
	PackagesWhiteLists               []GetHostAssurancePolicyPackagesWhiteList `pulumi:"packagesWhiteLists"`
	PartialResultsImageFail          bool                                      `pulumi:"partialResultsImageFail"`
	ReadOnly                         bool                                      `pulumi:"readOnly"`
	Registries                       []string                                  `pulumi:"registries"`
	Registry                         string                                    `pulumi:"registry"`
	RequiredLabels                   []GetHostAssurancePolicyRequiredLabel     `pulumi:"requiredLabels"`
	RequiredLabelsEnabled            bool                                      `pulumi:"requiredLabelsEnabled"`
	ScanNfsMounts                    bool                                      `pulumi:"scanNfsMounts"`
	ScanSensitiveData                bool                                      `pulumi:"scanSensitiveData"`
	ScapEnabled                      bool                                      `pulumi:"scapEnabled"`
	ScapFiles                        []string                                  `pulumi:"scapFiles"`
	Scopes                           []GetHostAssurancePolicyScope             `pulumi:"scopes"`
	TrustedBaseImages                []GetHostAssurancePolicyTrustedBaseImage  `pulumi:"trustedBaseImages"`
	TrustedBaseImagesEnabled         bool                                      `pulumi:"trustedBaseImagesEnabled"`
	WhitelistedLicenses              []string                                  `pulumi:"whitelistedLicenses"`
	WhitelistedLicensesEnabled       bool                                      `pulumi:"whitelistedLicensesEnabled"`
}

func LookupHostAssurancePolicyOutput(ctx *pulumi.Context, args LookupHostAssurancePolicyOutputArgs, opts ...pulumi.InvokeOption) LookupHostAssurancePolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHostAssurancePolicyResult, error) {
			args := v.(LookupHostAssurancePolicyArgs)
			r, err := LookupHostAssurancePolicy(ctx, &args, opts...)
			return *r, err
		}).(LookupHostAssurancePolicyResultOutput)
}

// A collection of arguments for invoking getHostAssurancePolicy.
type LookupHostAssurancePolicyOutputArgs struct {
	MaximumScoreExcludeNoFix pulumi.BoolPtrInput `pulumi:"maximumScoreExcludeNoFix"`
	Name                     pulumi.StringInput  `pulumi:"name"`
}

func (LookupHostAssurancePolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostAssurancePolicyArgs)(nil)).Elem()
}

// A collection of values returned by getHostAssurancePolicy.
type LookupHostAssurancePolicyResultOutput struct{ *pulumi.OutputState }

func (LookupHostAssurancePolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostAssurancePolicyResult)(nil)).Elem()
}

func (o LookupHostAssurancePolicyResultOutput) ToLookupHostAssurancePolicyResultOutput() LookupHostAssurancePolicyResultOutput {
	return o
}

func (o LookupHostAssurancePolicyResultOutput) ToLookupHostAssurancePolicyResultOutputWithContext(ctx context.Context) LookupHostAssurancePolicyResultOutput {
	return o
}

func (o LookupHostAssurancePolicyResultOutput) AllowedImages() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) []string { return v.AllowedImages }).(pulumi.StringArrayOutput)
}

func (o LookupHostAssurancePolicyResultOutput) ApplicationScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) []string { return v.ApplicationScopes }).(pulumi.StringArrayOutput)
}

func (o LookupHostAssurancePolicyResultOutput) AssuranceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) string { return v.AssuranceType }).(pulumi.StringOutput)
}

func (o LookupHostAssurancePolicyResultOutput) AuditOnFailure() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.AuditOnFailure }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) Author() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) string { return v.Author }).(pulumi.StringOutput)
}

func (o LookupHostAssurancePolicyResultOutput) AutoScanConfigured() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.AutoScanConfigured }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) AutoScanEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.AutoScanEnabled }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) AutoScanTimes() GetHostAssurancePolicyAutoScanTimeArrayOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) []GetHostAssurancePolicyAutoScanTime { return v.AutoScanTimes }).(GetHostAssurancePolicyAutoScanTimeArrayOutput)
}

func (o LookupHostAssurancePolicyResultOutput) BlacklistPermissions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) []string { return v.BlacklistPermissions }).(pulumi.StringArrayOutput)
}

func (o LookupHostAssurancePolicyResultOutput) BlacklistPermissionsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.BlacklistPermissionsEnabled }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) BlacklistedLicenses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) []string { return v.BlacklistedLicenses }).(pulumi.StringArrayOutput)
}

func (o LookupHostAssurancePolicyResultOutput) BlacklistedLicensesEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.BlacklistedLicensesEnabled }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) BlockFailed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.BlockFailed }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) ControlExcludeNoFix() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.ControlExcludeNoFix }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) CustomChecks() GetHostAssurancePolicyCustomCheckArrayOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) []GetHostAssurancePolicyCustomCheck { return v.CustomChecks }).(GetHostAssurancePolicyCustomCheckArrayOutput)
}

func (o LookupHostAssurancePolicyResultOutput) CustomChecksEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.CustomChecksEnabled }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) CustomSeverityEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.CustomSeverityEnabled }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) CvesBlackListEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.CvesBlackListEnabled }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) CvesBlackLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) []string { return v.CvesBlackLists }).(pulumi.StringArrayOutput)
}

func (o LookupHostAssurancePolicyResultOutput) CvesWhiteListEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.CvesWhiteListEnabled }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) CvesWhiteLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) []string { return v.CvesWhiteLists }).(pulumi.StringArrayOutput)
}

func (o LookupHostAssurancePolicyResultOutput) CvssSeverity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) string { return v.CvssSeverity }).(pulumi.StringOutput)
}

func (o LookupHostAssurancePolicyResultOutput) CvssSeverityEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.CvssSeverityEnabled }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) CvssSeverityExcludeNoFix() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.CvssSeverityExcludeNoFix }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupHostAssurancePolicyResultOutput) DisallowMalware() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.DisallowMalware }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) DockerCisEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.DockerCisEnabled }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) string { return v.Domain }).(pulumi.StringOutput)
}

func (o LookupHostAssurancePolicyResultOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) string { return v.DomainName }).(pulumi.StringOutput)
}

func (o LookupHostAssurancePolicyResultOutput) DtaEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.DtaEnabled }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) DtaSeverity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) string { return v.DtaSeverity }).(pulumi.StringOutput)
}

func (o LookupHostAssurancePolicyResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) Enforce() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.Enforce }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) EnforceAfterDays() pulumi.IntOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) int { return v.EnforceAfterDays }).(pulumi.IntOutput)
}

func (o LookupHostAssurancePolicyResultOutput) EnforceExcessivePermissions() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.EnforceExcessivePermissions }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) ExceptionalMonitoredMalwarePaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) []string { return v.ExceptionalMonitoredMalwarePaths }).(pulumi.StringArrayOutput)
}

func (o LookupHostAssurancePolicyResultOutput) FailCicd() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.FailCicd }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) ForbiddenLabels() GetHostAssurancePolicyForbiddenLabelArrayOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) []GetHostAssurancePolicyForbiddenLabel {
		return v.ForbiddenLabels
	}).(GetHostAssurancePolicyForbiddenLabelArrayOutput)
}

func (o LookupHostAssurancePolicyResultOutput) ForbiddenLabelsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.ForbiddenLabelsEnabled }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) ForceMicroenforcer() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.ForceMicroenforcer }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) FunctionIntegrityEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.FunctionIntegrityEnabled }).(pulumi.BoolOutput)
}

// The ID of this resource.
func (o LookupHostAssurancePolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupHostAssurancePolicyResultOutput) IgnoreRecentlyPublishedVln() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.IgnoreRecentlyPublishedVln }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) IgnoreRecentlyPublishedVlnPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) int { return v.IgnoreRecentlyPublishedVlnPeriod }).(pulumi.IntOutput)
}

func (o LookupHostAssurancePolicyResultOutput) IgnoreRiskResourcesEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.IgnoreRiskResourcesEnabled }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) IgnoredRiskResources() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) []string { return v.IgnoredRiskResources }).(pulumi.StringArrayOutput)
}

func (o LookupHostAssurancePolicyResultOutput) Images() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) []string { return v.Images }).(pulumi.StringArrayOutput)
}

func (o LookupHostAssurancePolicyResultOutput) KubeCisEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.KubeCisEnabled }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

func (o LookupHostAssurancePolicyResultOutput) MalwareAction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) string { return v.MalwareAction }).(pulumi.StringOutput)
}

func (o LookupHostAssurancePolicyResultOutput) MaximumScore() pulumi.Float64Output {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) float64 { return v.MaximumScore }).(pulumi.Float64Output)
}

func (o LookupHostAssurancePolicyResultOutput) MaximumScoreEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.MaximumScoreEnabled }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) MaximumScoreExcludeNoFix() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) *bool { return v.MaximumScoreExcludeNoFix }).(pulumi.BoolPtrOutput)
}

func (o LookupHostAssurancePolicyResultOutput) MonitoredMalwarePaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) []string { return v.MonitoredMalwarePaths }).(pulumi.StringArrayOutput)
}

func (o LookupHostAssurancePolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupHostAssurancePolicyResultOutput) OnlyNoneRootUsers() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.OnlyNoneRootUsers }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) PackagesBlackListEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.PackagesBlackListEnabled }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) PackagesBlackLists() GetHostAssurancePolicyPackagesBlackListArrayOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) []GetHostAssurancePolicyPackagesBlackList {
		return v.PackagesBlackLists
	}).(GetHostAssurancePolicyPackagesBlackListArrayOutput)
}

func (o LookupHostAssurancePolicyResultOutput) PackagesWhiteListEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.PackagesWhiteListEnabled }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) PackagesWhiteLists() GetHostAssurancePolicyPackagesWhiteListArrayOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) []GetHostAssurancePolicyPackagesWhiteList {
		return v.PackagesWhiteLists
	}).(GetHostAssurancePolicyPackagesWhiteListArrayOutput)
}

func (o LookupHostAssurancePolicyResultOutput) PartialResultsImageFail() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.PartialResultsImageFail }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) ReadOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.ReadOnly }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) Registries() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) []string { return v.Registries }).(pulumi.StringArrayOutput)
}

func (o LookupHostAssurancePolicyResultOutput) Registry() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) string { return v.Registry }).(pulumi.StringOutput)
}

func (o LookupHostAssurancePolicyResultOutput) RequiredLabels() GetHostAssurancePolicyRequiredLabelArrayOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) []GetHostAssurancePolicyRequiredLabel { return v.RequiredLabels }).(GetHostAssurancePolicyRequiredLabelArrayOutput)
}

func (o LookupHostAssurancePolicyResultOutput) RequiredLabelsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.RequiredLabelsEnabled }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) ScanNfsMounts() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.ScanNfsMounts }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) ScanSensitiveData() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.ScanSensitiveData }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) ScapEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.ScapEnabled }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) ScapFiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) []string { return v.ScapFiles }).(pulumi.StringArrayOutput)
}

func (o LookupHostAssurancePolicyResultOutput) Scopes() GetHostAssurancePolicyScopeArrayOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) []GetHostAssurancePolicyScope { return v.Scopes }).(GetHostAssurancePolicyScopeArrayOutput)
}

func (o LookupHostAssurancePolicyResultOutput) TrustedBaseImages() GetHostAssurancePolicyTrustedBaseImageArrayOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) []GetHostAssurancePolicyTrustedBaseImage {
		return v.TrustedBaseImages
	}).(GetHostAssurancePolicyTrustedBaseImageArrayOutput)
}

func (o LookupHostAssurancePolicyResultOutput) TrustedBaseImagesEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.TrustedBaseImagesEnabled }).(pulumi.BoolOutput)
}

func (o LookupHostAssurancePolicyResultOutput) WhitelistedLicenses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) []string { return v.WhitelistedLicenses }).(pulumi.StringArrayOutput)
}

func (o LookupHostAssurancePolicyResultOutput) WhitelistedLicensesEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHostAssurancePolicyResult) bool { return v.WhitelistedLicensesEnabled }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHostAssurancePolicyResultOutput{})
}
