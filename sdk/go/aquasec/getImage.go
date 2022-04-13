// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aquasec

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupImage(ctx *pulumi.Context, args *LookupImageArgs, opts ...pulumi.InvokeOption) (*LookupImageResult, error) {
	var rv LookupImageResult
	err := ctx.Invoke("aquasec:index/getImage:getImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getImage.
type LookupImageArgs struct {
	// The name of the registry where the image is stored.
	Registry string `pulumi:"registry"`
	// The name of the image's repository.
	Repository string `pulumi:"repository"`
	// The tag of the image.
	Tag string `pulumi:"tag"`
}

// A collection of values returned by getImage.
type LookupImageResult struct {
	// The image architecture.
	Architecture string `pulumi:"architecture"`
	// The list of image assurance checks performed on the image.
	AssuranceChecksPerformeds []GetImageAssuranceChecksPerformed `pulumi:"assuranceChecksPerformeds"`
	// The name of the user who registered the image.
	Author string `pulumi:"author"`
	// Whether the image is blacklisted.
	Blacklisted bool `pulumi:"blacklisted"`
	// The image creation comment.
	Comment string `pulumi:"comment"`
	// The date and time when the image was registered.
	Created string `pulumi:"created"`
	// Number of critical severity vulnerabilities detected in the image.
	CriticalVulnerabilities int `pulumi:"criticalVulnerabilities"`
	// The default user of the image.
	DefaultUser string `pulumi:"defaultUser"`
	// The content digest of the image.
	Digest string `pulumi:"digest"`
	// Whether the image is disallowed (non-compliant).
	Disallowed bool `pulumi:"disallowed"`
	// Whether the image was disallowed because of Image Assurance Policies.
	DisallowedByAssuranceChecks bool `pulumi:"disallowedByAssuranceChecks"`
	// The Docker image ID.
	DockerId string `pulumi:"dockerId"`
	// Docker labels of the image.
	DockerLabels []string `pulumi:"dockerLabels"`
	// The Docker version used when building the image.
	DockerVersion string `pulumi:"dockerVersion"`
	// DTA severity score.
	DtaSeverityScore string `pulumi:"dtaSeverityScore"`
	// If DTA was skipped.
	DtaSkipped bool `pulumi:"dtaSkipped"`
	// The reason why DTA was skipped.
	DtaSkippedReason string `pulumi:"dtaSkippedReason"`
	// Environment variables in the image.
	EnvironmentVariables []string `pulumi:"environmentVariables"`
	// Number of high severity vulnerabilities detected in the image.
	HighVulnerabilities int `pulumi:"highVulnerabilities"`
	// The Docker history of the image.
	Histories []GetImageHistory `pulumi:"histories"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The size of the image in bytes.
	ImageSize int `pulumi:"imageSize"`
	// The type of the image.
	ImageType string `pulumi:"imageType"`
	// Aqua labels of the image.
	Labels []string `pulumi:"labels"`
	// Number of low severity vulnerabilities detected in the image.
	LowVulnerabilities int `pulumi:"lowVulnerabilities"`
	// Number of malware found on the image.
	Malware int `pulumi:"malware"`
	// Number of medium severity vulnerabilities detected in the image.
	MediumVulnerabilities int `pulumi:"mediumVulnerabilities"`
	// The name of the image.
	Name string `pulumi:"name"`
	// Number of negligible severity vulnerabilities detected in the image.
	NegligibleVulnerabilities int `pulumi:"negligibleVulnerabilities"`
	// Whether a new version of the image is available in the registry but is not scanned and registered yet.
	NewerImageExists bool `pulumi:"newerImageExists"`
	// The operating system detected in the image
	Os string `pulumi:"os"`
	// The version of the OS detected in the image.
	OsVersion string `pulumi:"osVersion"`
	// The ID of the parent image.
	Parent string `pulumi:"parent"`
	// Whether the image could only be partially scanned.
	PartialResults bool `pulumi:"partialResults"`
	// Whether the image is non-compliant, but is pending this status due to running containers.
	PendingDisallowed bool `pulumi:"pendingDisallowed"`
	// Permission of the image.
	Permission string `pulumi:"permission"`
	// The name of the user who last modified the image permissions.
	PermissionAuthor string `pulumi:"permissionAuthor"`
	// The comment provided when the image permissions were last modified
	PermissionComment string `pulumi:"permissionComment"`
	// The name of the registry where the image is stored.
	Registry string `pulumi:"registry"`
	// Type of the registry.
	RegistryType string `pulumi:"registryType"`
	// The repository digests.
	RepoDigests []string `pulumi:"repoDigests"`
	// The name of the image's repository.
	Repository string `pulumi:"repository"`
	// The date and time when the image was last scanned.
	ScanDate string `pulumi:"scanDate"`
	// If the image scan failed, the failure message.
	ScanError string `pulumi:"scanError"`
	// The scan status of the image (either 'pending', 'in*progress', 'finished', 'failed' or 'not*started').
	ScanStatus string `pulumi:"scanStatus"`
	// Number of sensitive data detected in the image.
	SensitiveData int `pulumi:"sensitiveData"`
	// The tag of the image.
	Tag string `pulumi:"tag"`
	// The total number of vulnerabilities detected in the image.
	TotalVulnerabilities int `pulumi:"totalVulnerabilities"`
	// The virtual size of the image.
	VirtualSize int `pulumi:"virtualSize"`
	// A list of all the vulnerabilities found in the image
	Vulnerabilities []GetImageVulnerability `pulumi:"vulnerabilities"`
	// Whether the image is whitelisted.
	Whitelisted bool `pulumi:"whitelisted"`
}

func LookupImageOutput(ctx *pulumi.Context, args LookupImageOutputArgs, opts ...pulumi.InvokeOption) LookupImageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupImageResult, error) {
			args := v.(LookupImageArgs)
			r, err := LookupImage(ctx, &args, opts...)
			return *r, err
		}).(LookupImageResultOutput)
}

// A collection of arguments for invoking getImage.
type LookupImageOutputArgs struct {
	// The name of the registry where the image is stored.
	Registry pulumi.StringInput `pulumi:"registry"`
	// The name of the image's repository.
	Repository pulumi.StringInput `pulumi:"repository"`
	// The tag of the image.
	Tag pulumi.StringInput `pulumi:"tag"`
}

func (LookupImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageArgs)(nil)).Elem()
}

// A collection of values returned by getImage.
type LookupImageResultOutput struct{ *pulumi.OutputState }

func (LookupImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageResult)(nil)).Elem()
}

func (o LookupImageResultOutput) ToLookupImageResultOutput() LookupImageResultOutput {
	return o
}

func (o LookupImageResultOutput) ToLookupImageResultOutputWithContext(ctx context.Context) LookupImageResultOutput {
	return o
}

// The image architecture.
func (o LookupImageResultOutput) Architecture() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Architecture }).(pulumi.StringOutput)
}

// The list of image assurance checks performed on the image.
func (o LookupImageResultOutput) AssuranceChecksPerformeds() GetImageAssuranceChecksPerformedArrayOutput {
	return o.ApplyT(func(v LookupImageResult) []GetImageAssuranceChecksPerformed { return v.AssuranceChecksPerformeds }).(GetImageAssuranceChecksPerformedArrayOutput)
}

// The name of the user who registered the image.
func (o LookupImageResultOutput) Author() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Author }).(pulumi.StringOutput)
}

// Whether the image is blacklisted.
func (o LookupImageResultOutput) Blacklisted() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageResult) bool { return v.Blacklisted }).(pulumi.BoolOutput)
}

// The image creation comment.
func (o LookupImageResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Comment }).(pulumi.StringOutput)
}

// The date and time when the image was registered.
func (o LookupImageResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Created }).(pulumi.StringOutput)
}

// Number of critical severity vulnerabilities detected in the image.
func (o LookupImageResultOutput) CriticalVulnerabilities() pulumi.IntOutput {
	return o.ApplyT(func(v LookupImageResult) int { return v.CriticalVulnerabilities }).(pulumi.IntOutput)
}

// The default user of the image.
func (o LookupImageResultOutput) DefaultUser() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.DefaultUser }).(pulumi.StringOutput)
}

// The content digest of the image.
func (o LookupImageResultOutput) Digest() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Digest }).(pulumi.StringOutput)
}

// Whether the image is disallowed (non-compliant).
func (o LookupImageResultOutput) Disallowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageResult) bool { return v.Disallowed }).(pulumi.BoolOutput)
}

// Whether the image was disallowed because of Image Assurance Policies.
func (o LookupImageResultOutput) DisallowedByAssuranceChecks() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageResult) bool { return v.DisallowedByAssuranceChecks }).(pulumi.BoolOutput)
}

// The Docker image ID.
func (o LookupImageResultOutput) DockerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.DockerId }).(pulumi.StringOutput)
}

// Docker labels of the image.
func (o LookupImageResultOutput) DockerLabels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImageResult) []string { return v.DockerLabels }).(pulumi.StringArrayOutput)
}

// The Docker version used when building the image.
func (o LookupImageResultOutput) DockerVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.DockerVersion }).(pulumi.StringOutput)
}

// DTA severity score.
func (o LookupImageResultOutput) DtaSeverityScore() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.DtaSeverityScore }).(pulumi.StringOutput)
}

// If DTA was skipped.
func (o LookupImageResultOutput) DtaSkipped() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageResult) bool { return v.DtaSkipped }).(pulumi.BoolOutput)
}

// The reason why DTA was skipped.
func (o LookupImageResultOutput) DtaSkippedReason() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.DtaSkippedReason }).(pulumi.StringOutput)
}

// Environment variables in the image.
func (o LookupImageResultOutput) EnvironmentVariables() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImageResult) []string { return v.EnvironmentVariables }).(pulumi.StringArrayOutput)
}

// Number of high severity vulnerabilities detected in the image.
func (o LookupImageResultOutput) HighVulnerabilities() pulumi.IntOutput {
	return o.ApplyT(func(v LookupImageResult) int { return v.HighVulnerabilities }).(pulumi.IntOutput)
}

// The Docker history of the image.
func (o LookupImageResultOutput) Histories() GetImageHistoryArrayOutput {
	return o.ApplyT(func(v LookupImageResult) []GetImageHistory { return v.Histories }).(GetImageHistoryArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupImageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Id }).(pulumi.StringOutput)
}

// The size of the image in bytes.
func (o LookupImageResultOutput) ImageSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupImageResult) int { return v.ImageSize }).(pulumi.IntOutput)
}

// The type of the image.
func (o LookupImageResultOutput) ImageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.ImageType }).(pulumi.StringOutput)
}

// Aqua labels of the image.
func (o LookupImageResultOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImageResult) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

// Number of low severity vulnerabilities detected in the image.
func (o LookupImageResultOutput) LowVulnerabilities() pulumi.IntOutput {
	return o.ApplyT(func(v LookupImageResult) int { return v.LowVulnerabilities }).(pulumi.IntOutput)
}

// Number of malware found on the image.
func (o LookupImageResultOutput) Malware() pulumi.IntOutput {
	return o.ApplyT(func(v LookupImageResult) int { return v.Malware }).(pulumi.IntOutput)
}

// Number of medium severity vulnerabilities detected in the image.
func (o LookupImageResultOutput) MediumVulnerabilities() pulumi.IntOutput {
	return o.ApplyT(func(v LookupImageResult) int { return v.MediumVulnerabilities }).(pulumi.IntOutput)
}

// The name of the image.
func (o LookupImageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Name }).(pulumi.StringOutput)
}

// Number of negligible severity vulnerabilities detected in the image.
func (o LookupImageResultOutput) NegligibleVulnerabilities() pulumi.IntOutput {
	return o.ApplyT(func(v LookupImageResult) int { return v.NegligibleVulnerabilities }).(pulumi.IntOutput)
}

// Whether a new version of the image is available in the registry but is not scanned and registered yet.
func (o LookupImageResultOutput) NewerImageExists() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageResult) bool { return v.NewerImageExists }).(pulumi.BoolOutput)
}

// The operating system detected in the image
func (o LookupImageResultOutput) Os() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Os }).(pulumi.StringOutput)
}

// The version of the OS detected in the image.
func (o LookupImageResultOutput) OsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.OsVersion }).(pulumi.StringOutput)
}

// The ID of the parent image.
func (o LookupImageResultOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Parent }).(pulumi.StringOutput)
}

// Whether the image could only be partially scanned.
func (o LookupImageResultOutput) PartialResults() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageResult) bool { return v.PartialResults }).(pulumi.BoolOutput)
}

// Whether the image is non-compliant, but is pending this status due to running containers.
func (o LookupImageResultOutput) PendingDisallowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageResult) bool { return v.PendingDisallowed }).(pulumi.BoolOutput)
}

// Permission of the image.
func (o LookupImageResultOutput) Permission() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Permission }).(pulumi.StringOutput)
}

// The name of the user who last modified the image permissions.
func (o LookupImageResultOutput) PermissionAuthor() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.PermissionAuthor }).(pulumi.StringOutput)
}

// The comment provided when the image permissions were last modified
func (o LookupImageResultOutput) PermissionComment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.PermissionComment }).(pulumi.StringOutput)
}

// The name of the registry where the image is stored.
func (o LookupImageResultOutput) Registry() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Registry }).(pulumi.StringOutput)
}

// Type of the registry.
func (o LookupImageResultOutput) RegistryType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.RegistryType }).(pulumi.StringOutput)
}

// The repository digests.
func (o LookupImageResultOutput) RepoDigests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupImageResult) []string { return v.RepoDigests }).(pulumi.StringArrayOutput)
}

// The name of the image's repository.
func (o LookupImageResultOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Repository }).(pulumi.StringOutput)
}

// The date and time when the image was last scanned.
func (o LookupImageResultOutput) ScanDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.ScanDate }).(pulumi.StringOutput)
}

// If the image scan failed, the failure message.
func (o LookupImageResultOutput) ScanError() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.ScanError }).(pulumi.StringOutput)
}

// The scan status of the image (either 'pending', 'in*progress', 'finished', 'failed' or 'not*started').
func (o LookupImageResultOutput) ScanStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.ScanStatus }).(pulumi.StringOutput)
}

// Number of sensitive data detected in the image.
func (o LookupImageResultOutput) SensitiveData() pulumi.IntOutput {
	return o.ApplyT(func(v LookupImageResult) int { return v.SensitiveData }).(pulumi.IntOutput)
}

// The tag of the image.
func (o LookupImageResultOutput) Tag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageResult) string { return v.Tag }).(pulumi.StringOutput)
}

// The total number of vulnerabilities detected in the image.
func (o LookupImageResultOutput) TotalVulnerabilities() pulumi.IntOutput {
	return o.ApplyT(func(v LookupImageResult) int { return v.TotalVulnerabilities }).(pulumi.IntOutput)
}

// The virtual size of the image.
func (o LookupImageResultOutput) VirtualSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupImageResult) int { return v.VirtualSize }).(pulumi.IntOutput)
}

// A list of all the vulnerabilities found in the image
func (o LookupImageResultOutput) Vulnerabilities() GetImageVulnerabilityArrayOutput {
	return o.ApplyT(func(v LookupImageResult) []GetImageVulnerability { return v.Vulnerabilities }).(GetImageVulnerabilityArrayOutput)
}

// Whether the image is whitelisted.
func (o LookupImageResultOutput) Whitelisted() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImageResult) bool { return v.Whitelisted }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupImageResultOutput{})
}
