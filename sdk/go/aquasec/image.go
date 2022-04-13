// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aquasec

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Image struct {
	pulumi.CustomResourceState

	// If this field is set to true, the image will be whitelisted.
	AllowImage pulumi.BoolPtrOutput `pulumi:"allowImage"`
	// The image architecture.
	Architecture pulumi.StringOutput `pulumi:"architecture"`
	// The list of image assurance checks performed on the image.
	AssuranceChecksPerformeds ImageAssuranceChecksPerformedArrayOutput `pulumi:"assuranceChecksPerformeds"`
	// The name of the user who registered the image.
	Author pulumi.StringOutput `pulumi:"author"`
	// Whether the image is blacklisted.
	Blacklisted pulumi.BoolOutput `pulumi:"blacklisted"`
	// If this field is set to true, the image will be blacklisted.
	BlockImage pulumi.BoolPtrOutput `pulumi:"blockImage"`
	// The image creation comment.
	Comment pulumi.StringOutput `pulumi:"comment"`
	// The date and time when the image was registered.
	Created pulumi.StringOutput `pulumi:"created"`
	// Number of critical severity vulnerabilities detected in the image.
	CriticalVulnerabilities pulumi.IntOutput `pulumi:"criticalVulnerabilities"`
	// The default user of the image.
	DefaultUser pulumi.StringOutput `pulumi:"defaultUser"`
	// The content digest of the image.
	Digest pulumi.StringOutput `pulumi:"digest"`
	// Whether the image is disallowed (non-compliant).
	Disallowed pulumi.BoolOutput `pulumi:"disallowed"`
	// Whether the image was disallowed because of Image Assurance Policies.
	DisallowedByAssuranceChecks pulumi.BoolOutput `pulumi:"disallowedByAssuranceChecks"`
	// The Docker image ID.
	DockerId pulumi.StringOutput `pulumi:"dockerId"`
	// Docker labels of the image.
	DockerLabels pulumi.StringArrayOutput `pulumi:"dockerLabels"`
	// The Docker version used when building the image.
	DockerVersion pulumi.StringOutput `pulumi:"dockerVersion"`
	// DTA severity score.
	DtaSeverityScore pulumi.StringOutput `pulumi:"dtaSeverityScore"`
	// If DTA was skipped.
	DtaSkipped pulumi.BoolOutput `pulumi:"dtaSkipped"`
	// The reason why DTA was skipped.
	DtaSkippedReason pulumi.StringOutput `pulumi:"dtaSkippedReason"`
	// Environment variables in the image.
	EnvironmentVariables pulumi.StringArrayOutput `pulumi:"environmentVariables"`
	// Number of high severity vulnerabilities detected in the image.
	HighVulnerabilities pulumi.IntOutput `pulumi:"highVulnerabilities"`
	// The Docker history of the image.
	Histories ImageHistoryArrayOutput `pulumi:"histories"`
	// The size of the image in bytes.
	ImageSize pulumi.IntOutput `pulumi:"imageSize"`
	// The type of the image.
	ImageType pulumi.StringOutput `pulumi:"imageType"`
	// Aqua labels of the image.
	Labels pulumi.StringArrayOutput `pulumi:"labels"`
	// Number of low severity vulnerabilities detected in the image.
	LowVulnerabilities pulumi.IntOutput `pulumi:"lowVulnerabilities"`
	// Number of malware found on the image.
	Malware pulumi.IntOutput `pulumi:"malware"`
	// Number of medium severity vulnerabilities detected in the image.
	MediumVulnerabilities pulumi.IntOutput `pulumi:"mediumVulnerabilities"`
	// The name of the image.
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of negligible severity vulnerabilities detected in the image.
	NegligibleVulnerabilities pulumi.IntOutput `pulumi:"negligibleVulnerabilities"`
	// Whether a new version of the image is available in the registry but is not scanned and registered yet.
	NewerImageExists pulumi.BoolOutput `pulumi:"newerImageExists"`
	// The operating system detected in the image
	Os pulumi.StringOutput `pulumi:"os"`
	// The version of the OS detected in the image.
	OsVersion pulumi.StringOutput `pulumi:"osVersion"`
	// The ID of the parent image.
	Parent pulumi.StringOutput `pulumi:"parent"`
	// Whether the image could only be partially scanned.
	PartialResults pulumi.BoolOutput `pulumi:"partialResults"`
	// Whether the image is non-compliant, but is pending this status due to running containers.
	PendingDisallowed pulumi.BoolOutput `pulumi:"pendingDisallowed"`
	// Permission of the image.
	Permission pulumi.StringOutput `pulumi:"permission"`
	// The name of the user who last modified the image permissions.
	PermissionAuthor pulumi.StringOutput `pulumi:"permissionAuthor"`
	// The comment provided when the image permissions were last modified
	PermissionComment pulumi.StringOutput `pulumi:"permissionComment"`
	// A comment on why the image was whitelisted or blacklisted
	PermissionModificationComment pulumi.StringPtrOutput `pulumi:"permissionModificationComment"`
	// The name of the registry where the image is stored.
	Registry pulumi.StringOutput `pulumi:"registry"`
	// Type of the registry.
	RegistryType pulumi.StringOutput `pulumi:"registryType"`
	// The repository digests.
	RepoDigests pulumi.StringArrayOutput `pulumi:"repoDigests"`
	// The name of the image's repository.
	Repository pulumi.StringOutput `pulumi:"repository"`
	// The date and time when the image was last scanned.
	ScanDate pulumi.StringOutput `pulumi:"scanDate"`
	// If the image scan failed, the failure message.
	ScanError pulumi.StringOutput `pulumi:"scanError"`
	// The scan status of the image (either 'pending', 'in*progress', 'finished', 'failed' or 'not*started').
	ScanStatus pulumi.StringOutput `pulumi:"scanStatus"`
	// Number of sensitive data detected in the image.
	SensitiveData pulumi.IntOutput `pulumi:"sensitiveData"`
	// The tag of the image.
	Tag pulumi.StringOutput `pulumi:"tag"`
	// The total number of vulnerabilities detected in the image.
	TotalVulnerabilities pulumi.IntOutput `pulumi:"totalVulnerabilities"`
	// The virtual size of the image.
	VirtualSize pulumi.IntOutput `pulumi:"virtualSize"`
	// A list of all the vulnerabilities found in the image
	Vulnerabilities ImageVulnerabilityArrayOutput `pulumi:"vulnerabilities"`
	// Whether the image is whitelisted.
	Whitelisted pulumi.BoolOutput `pulumi:"whitelisted"`
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Registry == nil {
		return nil, errors.New("invalid value for required argument 'Registry'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	if args.Tag == nil {
		return nil, errors.New("invalid value for required argument 'Tag'")
	}
	var resource Image
	err := ctx.RegisterResource("aquasec:index/image:Image", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImage gets an existing Image resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageState, opts ...pulumi.ResourceOption) (*Image, error) {
	var resource Image
	err := ctx.ReadResource("aquasec:index/image:Image", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Image resources.
type imageState struct {
	// If this field is set to true, the image will be whitelisted.
	AllowImage *bool `pulumi:"allowImage"`
	// The image architecture.
	Architecture *string `pulumi:"architecture"`
	// The list of image assurance checks performed on the image.
	AssuranceChecksPerformeds []ImageAssuranceChecksPerformed `pulumi:"assuranceChecksPerformeds"`
	// The name of the user who registered the image.
	Author *string `pulumi:"author"`
	// Whether the image is blacklisted.
	Blacklisted *bool `pulumi:"blacklisted"`
	// If this field is set to true, the image will be blacklisted.
	BlockImage *bool `pulumi:"blockImage"`
	// The image creation comment.
	Comment *string `pulumi:"comment"`
	// The date and time when the image was registered.
	Created *string `pulumi:"created"`
	// Number of critical severity vulnerabilities detected in the image.
	CriticalVulnerabilities *int `pulumi:"criticalVulnerabilities"`
	// The default user of the image.
	DefaultUser *string `pulumi:"defaultUser"`
	// The content digest of the image.
	Digest *string `pulumi:"digest"`
	// Whether the image is disallowed (non-compliant).
	Disallowed *bool `pulumi:"disallowed"`
	// Whether the image was disallowed because of Image Assurance Policies.
	DisallowedByAssuranceChecks *bool `pulumi:"disallowedByAssuranceChecks"`
	// The Docker image ID.
	DockerId *string `pulumi:"dockerId"`
	// Docker labels of the image.
	DockerLabels []string `pulumi:"dockerLabels"`
	// The Docker version used when building the image.
	DockerVersion *string `pulumi:"dockerVersion"`
	// DTA severity score.
	DtaSeverityScore *string `pulumi:"dtaSeverityScore"`
	// If DTA was skipped.
	DtaSkipped *bool `pulumi:"dtaSkipped"`
	// The reason why DTA was skipped.
	DtaSkippedReason *string `pulumi:"dtaSkippedReason"`
	// Environment variables in the image.
	EnvironmentVariables []string `pulumi:"environmentVariables"`
	// Number of high severity vulnerabilities detected in the image.
	HighVulnerabilities *int `pulumi:"highVulnerabilities"`
	// The Docker history of the image.
	Histories []ImageHistory `pulumi:"histories"`
	// The size of the image in bytes.
	ImageSize *int `pulumi:"imageSize"`
	// The type of the image.
	ImageType *string `pulumi:"imageType"`
	// Aqua labels of the image.
	Labels []string `pulumi:"labels"`
	// Number of low severity vulnerabilities detected in the image.
	LowVulnerabilities *int `pulumi:"lowVulnerabilities"`
	// Number of malware found on the image.
	Malware *int `pulumi:"malware"`
	// Number of medium severity vulnerabilities detected in the image.
	MediumVulnerabilities *int `pulumi:"mediumVulnerabilities"`
	// The name of the image.
	Name *string `pulumi:"name"`
	// Number of negligible severity vulnerabilities detected in the image.
	NegligibleVulnerabilities *int `pulumi:"negligibleVulnerabilities"`
	// Whether a new version of the image is available in the registry but is not scanned and registered yet.
	NewerImageExists *bool `pulumi:"newerImageExists"`
	// The operating system detected in the image
	Os *string `pulumi:"os"`
	// The version of the OS detected in the image.
	OsVersion *string `pulumi:"osVersion"`
	// The ID of the parent image.
	Parent *string `pulumi:"parent"`
	// Whether the image could only be partially scanned.
	PartialResults *bool `pulumi:"partialResults"`
	// Whether the image is non-compliant, but is pending this status due to running containers.
	PendingDisallowed *bool `pulumi:"pendingDisallowed"`
	// Permission of the image.
	Permission *string `pulumi:"permission"`
	// The name of the user who last modified the image permissions.
	PermissionAuthor *string `pulumi:"permissionAuthor"`
	// The comment provided when the image permissions were last modified
	PermissionComment *string `pulumi:"permissionComment"`
	// A comment on why the image was whitelisted or blacklisted
	PermissionModificationComment *string `pulumi:"permissionModificationComment"`
	// The name of the registry where the image is stored.
	Registry *string `pulumi:"registry"`
	// Type of the registry.
	RegistryType *string `pulumi:"registryType"`
	// The repository digests.
	RepoDigests []string `pulumi:"repoDigests"`
	// The name of the image's repository.
	Repository *string `pulumi:"repository"`
	// The date and time when the image was last scanned.
	ScanDate *string `pulumi:"scanDate"`
	// If the image scan failed, the failure message.
	ScanError *string `pulumi:"scanError"`
	// The scan status of the image (either 'pending', 'in*progress', 'finished', 'failed' or 'not*started').
	ScanStatus *string `pulumi:"scanStatus"`
	// Number of sensitive data detected in the image.
	SensitiveData *int `pulumi:"sensitiveData"`
	// The tag of the image.
	Tag *string `pulumi:"tag"`
	// The total number of vulnerabilities detected in the image.
	TotalVulnerabilities *int `pulumi:"totalVulnerabilities"`
	// The virtual size of the image.
	VirtualSize *int `pulumi:"virtualSize"`
	// A list of all the vulnerabilities found in the image
	Vulnerabilities []ImageVulnerability `pulumi:"vulnerabilities"`
	// Whether the image is whitelisted.
	Whitelisted *bool `pulumi:"whitelisted"`
}

type ImageState struct {
	// If this field is set to true, the image will be whitelisted.
	AllowImage pulumi.BoolPtrInput
	// The image architecture.
	Architecture pulumi.StringPtrInput
	// The list of image assurance checks performed on the image.
	AssuranceChecksPerformeds ImageAssuranceChecksPerformedArrayInput
	// The name of the user who registered the image.
	Author pulumi.StringPtrInput
	// Whether the image is blacklisted.
	Blacklisted pulumi.BoolPtrInput
	// If this field is set to true, the image will be blacklisted.
	BlockImage pulumi.BoolPtrInput
	// The image creation comment.
	Comment pulumi.StringPtrInput
	// The date and time when the image was registered.
	Created pulumi.StringPtrInput
	// Number of critical severity vulnerabilities detected in the image.
	CriticalVulnerabilities pulumi.IntPtrInput
	// The default user of the image.
	DefaultUser pulumi.StringPtrInput
	// The content digest of the image.
	Digest pulumi.StringPtrInput
	// Whether the image is disallowed (non-compliant).
	Disallowed pulumi.BoolPtrInput
	// Whether the image was disallowed because of Image Assurance Policies.
	DisallowedByAssuranceChecks pulumi.BoolPtrInput
	// The Docker image ID.
	DockerId pulumi.StringPtrInput
	// Docker labels of the image.
	DockerLabels pulumi.StringArrayInput
	// The Docker version used when building the image.
	DockerVersion pulumi.StringPtrInput
	// DTA severity score.
	DtaSeverityScore pulumi.StringPtrInput
	// If DTA was skipped.
	DtaSkipped pulumi.BoolPtrInput
	// The reason why DTA was skipped.
	DtaSkippedReason pulumi.StringPtrInput
	// Environment variables in the image.
	EnvironmentVariables pulumi.StringArrayInput
	// Number of high severity vulnerabilities detected in the image.
	HighVulnerabilities pulumi.IntPtrInput
	// The Docker history of the image.
	Histories ImageHistoryArrayInput
	// The size of the image in bytes.
	ImageSize pulumi.IntPtrInput
	// The type of the image.
	ImageType pulumi.StringPtrInput
	// Aqua labels of the image.
	Labels pulumi.StringArrayInput
	// Number of low severity vulnerabilities detected in the image.
	LowVulnerabilities pulumi.IntPtrInput
	// Number of malware found on the image.
	Malware pulumi.IntPtrInput
	// Number of medium severity vulnerabilities detected in the image.
	MediumVulnerabilities pulumi.IntPtrInput
	// The name of the image.
	Name pulumi.StringPtrInput
	// Number of negligible severity vulnerabilities detected in the image.
	NegligibleVulnerabilities pulumi.IntPtrInput
	// Whether a new version of the image is available in the registry but is not scanned and registered yet.
	NewerImageExists pulumi.BoolPtrInput
	// The operating system detected in the image
	Os pulumi.StringPtrInput
	// The version of the OS detected in the image.
	OsVersion pulumi.StringPtrInput
	// The ID of the parent image.
	Parent pulumi.StringPtrInput
	// Whether the image could only be partially scanned.
	PartialResults pulumi.BoolPtrInput
	// Whether the image is non-compliant, but is pending this status due to running containers.
	PendingDisallowed pulumi.BoolPtrInput
	// Permission of the image.
	Permission pulumi.StringPtrInput
	// The name of the user who last modified the image permissions.
	PermissionAuthor pulumi.StringPtrInput
	// The comment provided when the image permissions were last modified
	PermissionComment pulumi.StringPtrInput
	// A comment on why the image was whitelisted or blacklisted
	PermissionModificationComment pulumi.StringPtrInput
	// The name of the registry where the image is stored.
	Registry pulumi.StringPtrInput
	// Type of the registry.
	RegistryType pulumi.StringPtrInput
	// The repository digests.
	RepoDigests pulumi.StringArrayInput
	// The name of the image's repository.
	Repository pulumi.StringPtrInput
	// The date and time when the image was last scanned.
	ScanDate pulumi.StringPtrInput
	// If the image scan failed, the failure message.
	ScanError pulumi.StringPtrInput
	// The scan status of the image (either 'pending', 'in*progress', 'finished', 'failed' or 'not*started').
	ScanStatus pulumi.StringPtrInput
	// Number of sensitive data detected in the image.
	SensitiveData pulumi.IntPtrInput
	// The tag of the image.
	Tag pulumi.StringPtrInput
	// The total number of vulnerabilities detected in the image.
	TotalVulnerabilities pulumi.IntPtrInput
	// The virtual size of the image.
	VirtualSize pulumi.IntPtrInput
	// A list of all the vulnerabilities found in the image
	Vulnerabilities ImageVulnerabilityArrayInput
	// Whether the image is whitelisted.
	Whitelisted pulumi.BoolPtrInput
}

func (ImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageState)(nil)).Elem()
}

type imageArgs struct {
	// If this field is set to true, the image will be whitelisted.
	AllowImage *bool `pulumi:"allowImage"`
	// If this field is set to true, the image will be blacklisted.
	BlockImage *bool `pulumi:"blockImage"`
	// A comment on why the image was whitelisted or blacklisted
	PermissionModificationComment *string `pulumi:"permissionModificationComment"`
	// The name of the registry where the image is stored.
	Registry string `pulumi:"registry"`
	// The name of the image's repository.
	Repository string `pulumi:"repository"`
	// The tag of the image.
	Tag string `pulumi:"tag"`
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	// If this field is set to true, the image will be whitelisted.
	AllowImage pulumi.BoolPtrInput
	// If this field is set to true, the image will be blacklisted.
	BlockImage pulumi.BoolPtrInput
	// A comment on why the image was whitelisted or blacklisted
	PermissionModificationComment pulumi.StringPtrInput
	// The name of the registry where the image is stored.
	Registry pulumi.StringInput
	// The name of the image's repository.
	Repository pulumi.StringInput
	// The tag of the image.
	Tag pulumi.StringInput
}

func (ImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageArgs)(nil)).Elem()
}

type ImageInput interface {
	pulumi.Input

	ToImageOutput() ImageOutput
	ToImageOutputWithContext(ctx context.Context) ImageOutput
}

func (*Image) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil)).Elem()
}

func (i *Image) ToImageOutput() ImageOutput {
	return i.ToImageOutputWithContext(context.Background())
}

func (i *Image) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOutput)
}

// ImageArrayInput is an input type that accepts ImageArray and ImageArrayOutput values.
// You can construct a concrete instance of `ImageArrayInput` via:
//
//          ImageArray{ ImageArgs{...} }
type ImageArrayInput interface {
	pulumi.Input

	ToImageArrayOutput() ImageArrayOutput
	ToImageArrayOutputWithContext(context.Context) ImageArrayOutput
}

type ImageArray []ImageInput

func (ImageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Image)(nil)).Elem()
}

func (i ImageArray) ToImageArrayOutput() ImageArrayOutput {
	return i.ToImageArrayOutputWithContext(context.Background())
}

func (i ImageArray) ToImageArrayOutputWithContext(ctx context.Context) ImageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageArrayOutput)
}

// ImageMapInput is an input type that accepts ImageMap and ImageMapOutput values.
// You can construct a concrete instance of `ImageMapInput` via:
//
//          ImageMap{ "key": ImageArgs{...} }
type ImageMapInput interface {
	pulumi.Input

	ToImageMapOutput() ImageMapOutput
	ToImageMapOutputWithContext(context.Context) ImageMapOutput
}

type ImageMap map[string]ImageInput

func (ImageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Image)(nil)).Elem()
}

func (i ImageMap) ToImageMapOutput() ImageMapOutput {
	return i.ToImageMapOutputWithContext(context.Background())
}

func (i ImageMap) ToImageMapOutputWithContext(ctx context.Context) ImageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageMapOutput)
}

type ImageOutput struct{ *pulumi.OutputState }

func (ImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil)).Elem()
}

func (o ImageOutput) ToImageOutput() ImageOutput {
	return o
}

func (o ImageOutput) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return o
}

type ImageArrayOutput struct{ *pulumi.OutputState }

func (ImageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Image)(nil)).Elem()
}

func (o ImageArrayOutput) ToImageArrayOutput() ImageArrayOutput {
	return o
}

func (o ImageArrayOutput) ToImageArrayOutputWithContext(ctx context.Context) ImageArrayOutput {
	return o
}

func (o ImageArrayOutput) Index(i pulumi.IntInput) ImageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Image {
		return vs[0].([]*Image)[vs[1].(int)]
	}).(ImageOutput)
}

type ImageMapOutput struct{ *pulumi.OutputState }

func (ImageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Image)(nil)).Elem()
}

func (o ImageMapOutput) ToImageMapOutput() ImageMapOutput {
	return o
}

func (o ImageMapOutput) ToImageMapOutputWithContext(ctx context.Context) ImageMapOutput {
	return o
}

func (o ImageMapOutput) MapIndex(k pulumi.StringInput) ImageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Image {
		return vs[0].(map[string]*Image)[vs[1].(string)]
	}).(ImageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageInput)(nil)).Elem(), &Image{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageArrayInput)(nil)).Elem(), ImageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageMapInput)(nil)).Elem(), ImageMap{})
	pulumi.RegisterOutputType(ImageOutput{})
	pulumi.RegisterOutputType(ImageArrayOutput{})
	pulumi.RegisterOutputType(ImageMapOutput{})
}
