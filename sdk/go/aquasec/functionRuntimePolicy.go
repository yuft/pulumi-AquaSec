// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aquasec

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FunctionRuntimePolicy struct {
	pulumi.CustomResourceState

	// Indicates the application scope of the service.
	ApplicationScopes pulumi.StringArrayOutput `pulumi:"applicationScopes"`
	// Username of the account that created the service.
	Author pulumi.StringOutput `pulumi:"author"`
	// If true, prevent creation of malicious executables in functions during their runtime post invocation.
	BlockMaliciousExecutables pulumi.BoolPtrOutput `pulumi:"blockMaliciousExecutables"`
	// List of executables that are prevented from running in containers.
	BlockedExecutables pulumi.StringArrayOutput `pulumi:"blockedExecutables"`
	// The description of the function runtime policy
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Indicates if the runtime policy is enabled or not.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Indicates that policy should effect container execution (not just for audit).
	Enforce pulumi.BoolPtrOutput `pulumi:"enforce"`
	// Honeypot User ID (Access Key)
	HoneypotAccessKey pulumi.StringPtrOutput `pulumi:"honeypotAccessKey"`
	// List of options to apply the honeypot on (Environment Vairable, Layer, File)
	HoneypotApplyOns pulumi.StringArrayOutput `pulumi:"honeypotApplyOns"`
	// Honeypot User Password (Secret Key)
	HoneypotSecretKey pulumi.StringPtrOutput `pulumi:"honeypotSecretKey"`
	// Serverless application name
	HoneypotServerlessAppName pulumi.StringPtrOutput `pulumi:"honeypotServerlessAppName"`
	// Name of the function runtime policy
	Name pulumi.StringOutput `pulumi:"name"`
	// Logical expression of how to compute the dependency of the scope variables.
	ScopeExpression pulumi.StringOutput `pulumi:"scopeExpression"`
	// List of scope attributes.
	ScopeVariables FunctionRuntimePolicyScopeVariableArrayOutput `pulumi:"scopeVariables"`
}

// NewFunctionRuntimePolicy registers a new resource with the given unique name, arguments, and options.
func NewFunctionRuntimePolicy(ctx *pulumi.Context,
	name string, args *FunctionRuntimePolicyArgs, opts ...pulumi.ResourceOption) (*FunctionRuntimePolicy, error) {
	if args == nil {
		args = &FunctionRuntimePolicyArgs{}
	}

	var resource FunctionRuntimePolicy
	err := ctx.RegisterResource("aquasec:index/functionRuntimePolicy:FunctionRuntimePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunctionRuntimePolicy gets an existing FunctionRuntimePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionRuntimePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionRuntimePolicyState, opts ...pulumi.ResourceOption) (*FunctionRuntimePolicy, error) {
	var resource FunctionRuntimePolicy
	err := ctx.ReadResource("aquasec:index/functionRuntimePolicy:FunctionRuntimePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FunctionRuntimePolicy resources.
type functionRuntimePolicyState struct {
	// Indicates the application scope of the service.
	ApplicationScopes []string `pulumi:"applicationScopes"`
	// Username of the account that created the service.
	Author *string `pulumi:"author"`
	// If true, prevent creation of malicious executables in functions during their runtime post invocation.
	BlockMaliciousExecutables *bool `pulumi:"blockMaliciousExecutables"`
	// List of executables that are prevented from running in containers.
	BlockedExecutables []string `pulumi:"blockedExecutables"`
	// The description of the function runtime policy
	Description *string `pulumi:"description"`
	// Indicates if the runtime policy is enabled or not.
	Enabled *bool `pulumi:"enabled"`
	// Indicates that policy should effect container execution (not just for audit).
	Enforce *bool `pulumi:"enforce"`
	// Honeypot User ID (Access Key)
	HoneypotAccessKey *string `pulumi:"honeypotAccessKey"`
	// List of options to apply the honeypot on (Environment Vairable, Layer, File)
	HoneypotApplyOns []string `pulumi:"honeypotApplyOns"`
	// Honeypot User Password (Secret Key)
	HoneypotSecretKey *string `pulumi:"honeypotSecretKey"`
	// Serverless application name
	HoneypotServerlessAppName *string `pulumi:"honeypotServerlessAppName"`
	// Name of the function runtime policy
	Name *string `pulumi:"name"`
	// Logical expression of how to compute the dependency of the scope variables.
	ScopeExpression *string `pulumi:"scopeExpression"`
	// List of scope attributes.
	ScopeVariables []FunctionRuntimePolicyScopeVariable `pulumi:"scopeVariables"`
}

type FunctionRuntimePolicyState struct {
	// Indicates the application scope of the service.
	ApplicationScopes pulumi.StringArrayInput
	// Username of the account that created the service.
	Author pulumi.StringPtrInput
	// If true, prevent creation of malicious executables in functions during their runtime post invocation.
	BlockMaliciousExecutables pulumi.BoolPtrInput
	// List of executables that are prevented from running in containers.
	BlockedExecutables pulumi.StringArrayInput
	// The description of the function runtime policy
	Description pulumi.StringPtrInput
	// Indicates if the runtime policy is enabled or not.
	Enabled pulumi.BoolPtrInput
	// Indicates that policy should effect container execution (not just for audit).
	Enforce pulumi.BoolPtrInput
	// Honeypot User ID (Access Key)
	HoneypotAccessKey pulumi.StringPtrInput
	// List of options to apply the honeypot on (Environment Vairable, Layer, File)
	HoneypotApplyOns pulumi.StringArrayInput
	// Honeypot User Password (Secret Key)
	HoneypotSecretKey pulumi.StringPtrInput
	// Serverless application name
	HoneypotServerlessAppName pulumi.StringPtrInput
	// Name of the function runtime policy
	Name pulumi.StringPtrInput
	// Logical expression of how to compute the dependency of the scope variables.
	ScopeExpression pulumi.StringPtrInput
	// List of scope attributes.
	ScopeVariables FunctionRuntimePolicyScopeVariableArrayInput
}

func (FunctionRuntimePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionRuntimePolicyState)(nil)).Elem()
}

type functionRuntimePolicyArgs struct {
	// Indicates the application scope of the service.
	ApplicationScopes []string `pulumi:"applicationScopes"`
	// If true, prevent creation of malicious executables in functions during their runtime post invocation.
	BlockMaliciousExecutables *bool `pulumi:"blockMaliciousExecutables"`
	// List of executables that are prevented from running in containers.
	BlockedExecutables []string `pulumi:"blockedExecutables"`
	// The description of the function runtime policy
	Description *string `pulumi:"description"`
	// Indicates if the runtime policy is enabled or not.
	Enabled *bool `pulumi:"enabled"`
	// Indicates that policy should effect container execution (not just for audit).
	Enforce *bool `pulumi:"enforce"`
	// Honeypot User ID (Access Key)
	HoneypotAccessKey *string `pulumi:"honeypotAccessKey"`
	// List of options to apply the honeypot on (Environment Vairable, Layer, File)
	HoneypotApplyOns []string `pulumi:"honeypotApplyOns"`
	// Honeypot User Password (Secret Key)
	HoneypotSecretKey *string `pulumi:"honeypotSecretKey"`
	// Serverless application name
	HoneypotServerlessAppName *string `pulumi:"honeypotServerlessAppName"`
	// Name of the function runtime policy
	Name *string `pulumi:"name"`
	// Logical expression of how to compute the dependency of the scope variables.
	ScopeExpression *string `pulumi:"scopeExpression"`
	// List of scope attributes.
	ScopeVariables []FunctionRuntimePolicyScopeVariable `pulumi:"scopeVariables"`
}

// The set of arguments for constructing a FunctionRuntimePolicy resource.
type FunctionRuntimePolicyArgs struct {
	// Indicates the application scope of the service.
	ApplicationScopes pulumi.StringArrayInput
	// If true, prevent creation of malicious executables in functions during their runtime post invocation.
	BlockMaliciousExecutables pulumi.BoolPtrInput
	// List of executables that are prevented from running in containers.
	BlockedExecutables pulumi.StringArrayInput
	// The description of the function runtime policy
	Description pulumi.StringPtrInput
	// Indicates if the runtime policy is enabled or not.
	Enabled pulumi.BoolPtrInput
	// Indicates that policy should effect container execution (not just for audit).
	Enforce pulumi.BoolPtrInput
	// Honeypot User ID (Access Key)
	HoneypotAccessKey pulumi.StringPtrInput
	// List of options to apply the honeypot on (Environment Vairable, Layer, File)
	HoneypotApplyOns pulumi.StringArrayInput
	// Honeypot User Password (Secret Key)
	HoneypotSecretKey pulumi.StringPtrInput
	// Serverless application name
	HoneypotServerlessAppName pulumi.StringPtrInput
	// Name of the function runtime policy
	Name pulumi.StringPtrInput
	// Logical expression of how to compute the dependency of the scope variables.
	ScopeExpression pulumi.StringPtrInput
	// List of scope attributes.
	ScopeVariables FunctionRuntimePolicyScopeVariableArrayInput
}

func (FunctionRuntimePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionRuntimePolicyArgs)(nil)).Elem()
}

type FunctionRuntimePolicyInput interface {
	pulumi.Input

	ToFunctionRuntimePolicyOutput() FunctionRuntimePolicyOutput
	ToFunctionRuntimePolicyOutputWithContext(ctx context.Context) FunctionRuntimePolicyOutput
}

func (*FunctionRuntimePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionRuntimePolicy)(nil)).Elem()
}

func (i *FunctionRuntimePolicy) ToFunctionRuntimePolicyOutput() FunctionRuntimePolicyOutput {
	return i.ToFunctionRuntimePolicyOutputWithContext(context.Background())
}

func (i *FunctionRuntimePolicy) ToFunctionRuntimePolicyOutputWithContext(ctx context.Context) FunctionRuntimePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionRuntimePolicyOutput)
}

// FunctionRuntimePolicyArrayInput is an input type that accepts FunctionRuntimePolicyArray and FunctionRuntimePolicyArrayOutput values.
// You can construct a concrete instance of `FunctionRuntimePolicyArrayInput` via:
//
//          FunctionRuntimePolicyArray{ FunctionRuntimePolicyArgs{...} }
type FunctionRuntimePolicyArrayInput interface {
	pulumi.Input

	ToFunctionRuntimePolicyArrayOutput() FunctionRuntimePolicyArrayOutput
	ToFunctionRuntimePolicyArrayOutputWithContext(context.Context) FunctionRuntimePolicyArrayOutput
}

type FunctionRuntimePolicyArray []FunctionRuntimePolicyInput

func (FunctionRuntimePolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionRuntimePolicy)(nil)).Elem()
}

func (i FunctionRuntimePolicyArray) ToFunctionRuntimePolicyArrayOutput() FunctionRuntimePolicyArrayOutput {
	return i.ToFunctionRuntimePolicyArrayOutputWithContext(context.Background())
}

func (i FunctionRuntimePolicyArray) ToFunctionRuntimePolicyArrayOutputWithContext(ctx context.Context) FunctionRuntimePolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionRuntimePolicyArrayOutput)
}

// FunctionRuntimePolicyMapInput is an input type that accepts FunctionRuntimePolicyMap and FunctionRuntimePolicyMapOutput values.
// You can construct a concrete instance of `FunctionRuntimePolicyMapInput` via:
//
//          FunctionRuntimePolicyMap{ "key": FunctionRuntimePolicyArgs{...} }
type FunctionRuntimePolicyMapInput interface {
	pulumi.Input

	ToFunctionRuntimePolicyMapOutput() FunctionRuntimePolicyMapOutput
	ToFunctionRuntimePolicyMapOutputWithContext(context.Context) FunctionRuntimePolicyMapOutput
}

type FunctionRuntimePolicyMap map[string]FunctionRuntimePolicyInput

func (FunctionRuntimePolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionRuntimePolicy)(nil)).Elem()
}

func (i FunctionRuntimePolicyMap) ToFunctionRuntimePolicyMapOutput() FunctionRuntimePolicyMapOutput {
	return i.ToFunctionRuntimePolicyMapOutputWithContext(context.Background())
}

func (i FunctionRuntimePolicyMap) ToFunctionRuntimePolicyMapOutputWithContext(ctx context.Context) FunctionRuntimePolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionRuntimePolicyMapOutput)
}

type FunctionRuntimePolicyOutput struct{ *pulumi.OutputState }

func (FunctionRuntimePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionRuntimePolicy)(nil)).Elem()
}

func (o FunctionRuntimePolicyOutput) ToFunctionRuntimePolicyOutput() FunctionRuntimePolicyOutput {
	return o
}

func (o FunctionRuntimePolicyOutput) ToFunctionRuntimePolicyOutputWithContext(ctx context.Context) FunctionRuntimePolicyOutput {
	return o
}

type FunctionRuntimePolicyArrayOutput struct{ *pulumi.OutputState }

func (FunctionRuntimePolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionRuntimePolicy)(nil)).Elem()
}

func (o FunctionRuntimePolicyArrayOutput) ToFunctionRuntimePolicyArrayOutput() FunctionRuntimePolicyArrayOutput {
	return o
}

func (o FunctionRuntimePolicyArrayOutput) ToFunctionRuntimePolicyArrayOutputWithContext(ctx context.Context) FunctionRuntimePolicyArrayOutput {
	return o
}

func (o FunctionRuntimePolicyArrayOutput) Index(i pulumi.IntInput) FunctionRuntimePolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FunctionRuntimePolicy {
		return vs[0].([]*FunctionRuntimePolicy)[vs[1].(int)]
	}).(FunctionRuntimePolicyOutput)
}

type FunctionRuntimePolicyMapOutput struct{ *pulumi.OutputState }

func (FunctionRuntimePolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionRuntimePolicy)(nil)).Elem()
}

func (o FunctionRuntimePolicyMapOutput) ToFunctionRuntimePolicyMapOutput() FunctionRuntimePolicyMapOutput {
	return o
}

func (o FunctionRuntimePolicyMapOutput) ToFunctionRuntimePolicyMapOutputWithContext(ctx context.Context) FunctionRuntimePolicyMapOutput {
	return o
}

func (o FunctionRuntimePolicyMapOutput) MapIndex(k pulumi.StringInput) FunctionRuntimePolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FunctionRuntimePolicy {
		return vs[0].(map[string]*FunctionRuntimePolicy)[vs[1].(string)]
	}).(FunctionRuntimePolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionRuntimePolicyInput)(nil)).Elem(), &FunctionRuntimePolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionRuntimePolicyArrayInput)(nil)).Elem(), FunctionRuntimePolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionRuntimePolicyMapInput)(nil)).Elem(), FunctionRuntimePolicyMap{})
	pulumi.RegisterOutputType(FunctionRuntimePolicyOutput{})
	pulumi.RegisterOutputType(FunctionRuntimePolicyArrayOutput{})
	pulumi.RegisterOutputType(FunctionRuntimePolicyMapOutput{})
}