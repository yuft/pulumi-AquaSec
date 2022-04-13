// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aquasec

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("aquasec:index/getService:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getService.
type LookupServiceArgs struct {
	// The name of the service. It is recommended not to use whitespace characters in the name.
	Name string `pulumi:"name"`
}

// A collection of values returned by getService.
type LookupServiceResult struct {
	// Indicates the application scope of the service.
	ApplicationScopes []string `pulumi:"applicationScopes"`
	// Username of the account that created the service.
	Author string `pulumi:"author"`
	// The number of containers associated with the service.
	ContainersCount int `pulumi:"containersCount"`
	// A textual description of the service record; maximum 500 characters.
	Description string `pulumi:"description"`
	// Enforcement status of the service.
	Enforce bool `pulumi:"enforce"`
	// Whether the service has been evaluated for security vulnerabilities.
	Evaluated bool `pulumi:"evaluated"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Indicates if registered or not.
	IsRegistered bool `pulumi:"isRegistered"`
	// Timestamp of the last update in Unix time format.
	Lastupdate int `pulumi:"lastupdate"`
	// Indicates if monitoring is enabled or not
	Monitoring bool `pulumi:"monitoring"`
	// The name of the service. It is recommended not to use whitespace characters in the name.
	Name string `pulumi:"name"`
	// The number of container that are not evaluated.
	NotEvaluatedCount int `pulumi:"notEvaluatedCount"`
	// The service's policies; an array of container firewall policy names.
	Policies []string `pulumi:"policies"`
	// Rules priority, must be between 1-100.
	Priority int `pulumi:"priority"`
	// Logical expression of how to compute the dependency of the scope variables.
	ScopeExpression string `pulumi:"scopeExpression"`
	// List of scope attributes.
	ScopeVariables []GetServiceScopeVariable `pulumi:"scopeVariables"`
	// Type of the workload. container or host.
	Target string `pulumi:"target"`
	// The number of containers allocated to the service that are not registered.
	UnregisteredCount int `pulumi:"unregisteredCount"`
	// Number of high severity vulnerabilities.
	VulnerabilitiesHigh int `pulumi:"vulnerabilitiesHigh"`
	// Number of low severity vulnerabilities.
	VulnerabilitiesLow int `pulumi:"vulnerabilitiesLow"`
	// Number of malware.
	VulnerabilitiesMalware int `pulumi:"vulnerabilitiesMalware"`
	// Number of medium severity vulnerabilities.
	VulnerabilitiesMedium int `pulumi:"vulnerabilitiesMedium"`
	// Number of negligible vulnerabilities.
	VulnerabilitiesNegligible int `pulumi:"vulnerabilitiesNegligible"`
	// The CVSS average vulnerabilities score.
	VulnerabilitiesScoreAverage int `pulumi:"vulnerabilitiesScoreAverage"`
	// Number of sensitive vulnerabilities.
	VulnerabilitiesSensitive int `pulumi:"vulnerabilitiesSensitive"`
	// Total number of vulnerabilities.
	VulnerabilitiesTotal int `pulumi:"vulnerabilitiesTotal"`
}

func LookupServiceOutput(ctx *pulumi.Context, args LookupServiceOutputArgs, opts ...pulumi.InvokeOption) LookupServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceResult, error) {
			args := v.(LookupServiceArgs)
			r, err := LookupService(ctx, &args, opts...)
			return *r, err
		}).(LookupServiceResultOutput)
}

// A collection of arguments for invoking getService.
type LookupServiceOutputArgs struct {
	// The name of the service. It is recommended not to use whitespace characters in the name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceArgs)(nil)).Elem()
}

// A collection of values returned by getService.
type LookupServiceResultOutput struct{ *pulumi.OutputState }

func (LookupServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceResult)(nil)).Elem()
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutput() LookupServiceResultOutput {
	return o
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutputWithContext(ctx context.Context) LookupServiceResultOutput {
	return o
}

// Indicates the application scope of the service.
func (o LookupServiceResultOutput) ApplicationScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServiceResult) []string { return v.ApplicationScopes }).(pulumi.StringArrayOutput)
}

// Username of the account that created the service.
func (o LookupServiceResultOutput) Author() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Author }).(pulumi.StringOutput)
}

// The number of containers associated with the service.
func (o LookupServiceResultOutput) ContainersCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupServiceResult) int { return v.ContainersCount }).(pulumi.IntOutput)
}

// A textual description of the service record; maximum 500 characters.
func (o LookupServiceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Description }).(pulumi.StringOutput)
}

// Enforcement status of the service.
func (o LookupServiceResultOutput) Enforce() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServiceResult) bool { return v.Enforce }).(pulumi.BoolOutput)
}

// Whether the service has been evaluated for security vulnerabilities.
func (o LookupServiceResultOutput) Evaluated() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServiceResult) bool { return v.Evaluated }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Indicates if registered or not.
func (o LookupServiceResultOutput) IsRegistered() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServiceResult) bool { return v.IsRegistered }).(pulumi.BoolOutput)
}

// Timestamp of the last update in Unix time format.
func (o LookupServiceResultOutput) Lastupdate() pulumi.IntOutput {
	return o.ApplyT(func(v LookupServiceResult) int { return v.Lastupdate }).(pulumi.IntOutput)
}

// Indicates if monitoring is enabled or not
func (o LookupServiceResultOutput) Monitoring() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServiceResult) bool { return v.Monitoring }).(pulumi.BoolOutput)
}

// The name of the service. It is recommended not to use whitespace characters in the name.
func (o LookupServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The number of container that are not evaluated.
func (o LookupServiceResultOutput) NotEvaluatedCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupServiceResult) int { return v.NotEvaluatedCount }).(pulumi.IntOutput)
}

// The service's policies; an array of container firewall policy names.
func (o LookupServiceResultOutput) Policies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServiceResult) []string { return v.Policies }).(pulumi.StringArrayOutput)
}

// Rules priority, must be between 1-100.
func (o LookupServiceResultOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v LookupServiceResult) int { return v.Priority }).(pulumi.IntOutput)
}

// Logical expression of how to compute the dependency of the scope variables.
func (o LookupServiceResultOutput) ScopeExpression() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.ScopeExpression }).(pulumi.StringOutput)
}

// List of scope attributes.
func (o LookupServiceResultOutput) ScopeVariables() GetServiceScopeVariableArrayOutput {
	return o.ApplyT(func(v LookupServiceResult) []GetServiceScopeVariable { return v.ScopeVariables }).(GetServiceScopeVariableArrayOutput)
}

// Type of the workload. container or host.
func (o LookupServiceResultOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Target }).(pulumi.StringOutput)
}

// The number of containers allocated to the service that are not registered.
func (o LookupServiceResultOutput) UnregisteredCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupServiceResult) int { return v.UnregisteredCount }).(pulumi.IntOutput)
}

// Number of high severity vulnerabilities.
func (o LookupServiceResultOutput) VulnerabilitiesHigh() pulumi.IntOutput {
	return o.ApplyT(func(v LookupServiceResult) int { return v.VulnerabilitiesHigh }).(pulumi.IntOutput)
}

// Number of low severity vulnerabilities.
func (o LookupServiceResultOutput) VulnerabilitiesLow() pulumi.IntOutput {
	return o.ApplyT(func(v LookupServiceResult) int { return v.VulnerabilitiesLow }).(pulumi.IntOutput)
}

// Number of malware.
func (o LookupServiceResultOutput) VulnerabilitiesMalware() pulumi.IntOutput {
	return o.ApplyT(func(v LookupServiceResult) int { return v.VulnerabilitiesMalware }).(pulumi.IntOutput)
}

// Number of medium severity vulnerabilities.
func (o LookupServiceResultOutput) VulnerabilitiesMedium() pulumi.IntOutput {
	return o.ApplyT(func(v LookupServiceResult) int { return v.VulnerabilitiesMedium }).(pulumi.IntOutput)
}

// Number of negligible vulnerabilities.
func (o LookupServiceResultOutput) VulnerabilitiesNegligible() pulumi.IntOutput {
	return o.ApplyT(func(v LookupServiceResult) int { return v.VulnerabilitiesNegligible }).(pulumi.IntOutput)
}

// The CVSS average vulnerabilities score.
func (o LookupServiceResultOutput) VulnerabilitiesScoreAverage() pulumi.IntOutput {
	return o.ApplyT(func(v LookupServiceResult) int { return v.VulnerabilitiesScoreAverage }).(pulumi.IntOutput)
}

// Number of sensitive vulnerabilities.
func (o LookupServiceResultOutput) VulnerabilitiesSensitive() pulumi.IntOutput {
	return o.ApplyT(func(v LookupServiceResult) int { return v.VulnerabilitiesSensitive }).(pulumi.IntOutput)
}

// Total number of vulnerabilities.
func (o LookupServiceResultOutput) VulnerabilitiesTotal() pulumi.IntOutput {
	return o.ApplyT(func(v LookupServiceResult) int { return v.VulnerabilitiesTotal }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceResultOutput{})
}