// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aquasec

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AquasecHostRuntimePolicy struct {
	pulumi.CustomResourceState

	// Indicates the application scope of the service.
	ApplicationScopes pulumi.StringArrayOutput `pulumi:"applicationScopes"`
	// If true, all process activity will be audited.
	AuditAllOsUserActivity pulumi.BoolPtrOutput `pulumi:"auditAllOsUserActivity"`
	// If true, full command arguments will be audited.
	AuditFullCommandArguments pulumi.BoolPtrOutput `pulumi:"auditFullCommandArguments"`
	// Username of the account that created the service.
	Author pulumi.StringOutput `pulumi:"author"`
	// List of files that are prevented from being read, modified and executed in the containers.
	BlockedFiles pulumi.StringArrayOutput `pulumi:"blockedFiles"`
	// The description of the host runtime policy
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// If true, detect and prevent communication from containers to IP addresses known to have a bad reputation.
	EnableIpReputationSecurity pulumi.BoolPtrOutput `pulumi:"enableIpReputationSecurity"`
	// Indicates if the runtime policy is enabled or not.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Indicates that policy should effect container execution (not just for audit).
	Enforce pulumi.BoolPtrOutput `pulumi:"enforce"`
	// Indicates the number of days after which the runtime policy will be changed to enforce mode.
	EnforceAfterDays pulumi.IntPtrOutput `pulumi:"enforceAfterDays"`
	// Configuration for file integrity monitoring.
	FileIntegrityMonitoring AquasecHostRuntimePolicyFileIntegrityMonitoringPtrOutput `pulumi:"fileIntegrityMonitoring"`
	// If true, system time changes will be monitored.
	MonitorSystemTimeChanges pulumi.BoolPtrOutput `pulumi:"monitorSystemTimeChanges"`
	// If true, windows service operations will be monitored.
	MonitorWindowsServices pulumi.BoolPtrOutput `pulumi:"monitorWindowsServices"`
	// Name of the host runtime policy
	Name pulumi.StringOutput `pulumi:"name"`
	// List of OS (Linux or Windows) groups that are allowed to authenticate to the host, and block authentication requests from all others. Groups can be either Linux groups or Windows AD groups.
	OsGroupsAlloweds pulumi.StringArrayOutput `pulumi:"osGroupsAlloweds"`
	// List of OS (Linux or Windows) groups that are not allowed to authenticate to the host, and block authentication requests from all others. Groups can be either Linux groups or Windows AD groups.
	OsGroupsBlockeds pulumi.StringArrayOutput `pulumi:"osGroupsBlockeds"`
	// List of OS (Linux or Windows) users that are allowed to authenticate to the host, and block authentication requests from all others.
	OsUsersAlloweds pulumi.StringArrayOutput `pulumi:"osUsersAlloweds"`
	// List of OS (Linux or Windows) users that are not allowed to authenticate to the host, and block authentication requests from all others.
	OsUsersBlockeds pulumi.StringArrayOutput `pulumi:"osUsersBlockeds"`
	// Logical expression of how to compute the dependency of the scope variables.
	ScopeExpression pulumi.StringOutput `pulumi:"scopeExpression"`
	// List of scope attributes.
	ScopeVariables AquasecHostRuntimePolicyScopeVariableArrayOutput `pulumi:"scopeVariables"`
	// Configuration for windows registry monitoring.
	WindowsRegistryMonitoring AquasecHostRuntimePolicyWindowsRegistryMonitoringPtrOutput `pulumi:"windowsRegistryMonitoring"`
	// Configuration for windows registry protection.
	WindowsRegistryProtection AquasecHostRuntimePolicyWindowsRegistryProtectionPtrOutput `pulumi:"windowsRegistryProtection"`
}

// NewAquasecHostRuntimePolicy registers a new resource with the given unique name, arguments, and options.
func NewAquasecHostRuntimePolicy(ctx *pulumi.Context,
	name string, args *AquasecHostRuntimePolicyArgs, opts ...pulumi.ResourceOption) (*AquasecHostRuntimePolicy, error) {
	if args == nil {
		args = &AquasecHostRuntimePolicyArgs{}
	}

	var resource AquasecHostRuntimePolicy
	err := ctx.RegisterResource("aquasec:index/aquasecHostRuntimePolicy:AquasecHostRuntimePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAquasecHostRuntimePolicy gets an existing AquasecHostRuntimePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAquasecHostRuntimePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AquasecHostRuntimePolicyState, opts ...pulumi.ResourceOption) (*AquasecHostRuntimePolicy, error) {
	var resource AquasecHostRuntimePolicy
	err := ctx.ReadResource("aquasec:index/aquasecHostRuntimePolicy:AquasecHostRuntimePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AquasecHostRuntimePolicy resources.
type aquasecHostRuntimePolicyState struct {
	// Indicates the application scope of the service.
	ApplicationScopes []string `pulumi:"applicationScopes"`
	// If true, all process activity will be audited.
	AuditAllOsUserActivity *bool `pulumi:"auditAllOsUserActivity"`
	// If true, full command arguments will be audited.
	AuditFullCommandArguments *bool `pulumi:"auditFullCommandArguments"`
	// Username of the account that created the service.
	Author *string `pulumi:"author"`
	// List of files that are prevented from being read, modified and executed in the containers.
	BlockedFiles []string `pulumi:"blockedFiles"`
	// The description of the host runtime policy
	Description *string `pulumi:"description"`
	// If true, detect and prevent communication from containers to IP addresses known to have a bad reputation.
	EnableIpReputationSecurity *bool `pulumi:"enableIpReputationSecurity"`
	// Indicates if the runtime policy is enabled or not.
	Enabled *bool `pulumi:"enabled"`
	// Indicates that policy should effect container execution (not just for audit).
	Enforce *bool `pulumi:"enforce"`
	// Indicates the number of days after which the runtime policy will be changed to enforce mode.
	EnforceAfterDays *int `pulumi:"enforceAfterDays"`
	// Configuration for file integrity monitoring.
	FileIntegrityMonitoring *AquasecHostRuntimePolicyFileIntegrityMonitoring `pulumi:"fileIntegrityMonitoring"`
	// If true, system time changes will be monitored.
	MonitorSystemTimeChanges *bool `pulumi:"monitorSystemTimeChanges"`
	// If true, windows service operations will be monitored.
	MonitorWindowsServices *bool `pulumi:"monitorWindowsServices"`
	// Name of the host runtime policy
	Name *string `pulumi:"name"`
	// List of OS (Linux or Windows) groups that are allowed to authenticate to the host, and block authentication requests from all others. Groups can be either Linux groups or Windows AD groups.
	OsGroupsAlloweds []string `pulumi:"osGroupsAlloweds"`
	// List of OS (Linux or Windows) groups that are not allowed to authenticate to the host, and block authentication requests from all others. Groups can be either Linux groups or Windows AD groups.
	OsGroupsBlockeds []string `pulumi:"osGroupsBlockeds"`
	// List of OS (Linux or Windows) users that are allowed to authenticate to the host, and block authentication requests from all others.
	OsUsersAlloweds []string `pulumi:"osUsersAlloweds"`
	// List of OS (Linux or Windows) users that are not allowed to authenticate to the host, and block authentication requests from all others.
	OsUsersBlockeds []string `pulumi:"osUsersBlockeds"`
	// Logical expression of how to compute the dependency of the scope variables.
	ScopeExpression *string `pulumi:"scopeExpression"`
	// List of scope attributes.
	ScopeVariables []AquasecHostRuntimePolicyScopeVariable `pulumi:"scopeVariables"`
	// Configuration for windows registry monitoring.
	WindowsRegistryMonitoring *AquasecHostRuntimePolicyWindowsRegistryMonitoring `pulumi:"windowsRegistryMonitoring"`
	// Configuration for windows registry protection.
	WindowsRegistryProtection *AquasecHostRuntimePolicyWindowsRegistryProtection `pulumi:"windowsRegistryProtection"`
}

type AquasecHostRuntimePolicyState struct {
	// Indicates the application scope of the service.
	ApplicationScopes pulumi.StringArrayInput
	// If true, all process activity will be audited.
	AuditAllOsUserActivity pulumi.BoolPtrInput
	// If true, full command arguments will be audited.
	AuditFullCommandArguments pulumi.BoolPtrInput
	// Username of the account that created the service.
	Author pulumi.StringPtrInput
	// List of files that are prevented from being read, modified and executed in the containers.
	BlockedFiles pulumi.StringArrayInput
	// The description of the host runtime policy
	Description pulumi.StringPtrInput
	// If true, detect and prevent communication from containers to IP addresses known to have a bad reputation.
	EnableIpReputationSecurity pulumi.BoolPtrInput
	// Indicates if the runtime policy is enabled or not.
	Enabled pulumi.BoolPtrInput
	// Indicates that policy should effect container execution (not just for audit).
	Enforce pulumi.BoolPtrInput
	// Indicates the number of days after which the runtime policy will be changed to enforce mode.
	EnforceAfterDays pulumi.IntPtrInput
	// Configuration for file integrity monitoring.
	FileIntegrityMonitoring AquasecHostRuntimePolicyFileIntegrityMonitoringPtrInput
	// If true, system time changes will be monitored.
	MonitorSystemTimeChanges pulumi.BoolPtrInput
	// If true, windows service operations will be monitored.
	MonitorWindowsServices pulumi.BoolPtrInput
	// Name of the host runtime policy
	Name pulumi.StringPtrInput
	// List of OS (Linux or Windows) groups that are allowed to authenticate to the host, and block authentication requests from all others. Groups can be either Linux groups or Windows AD groups.
	OsGroupsAlloweds pulumi.StringArrayInput
	// List of OS (Linux or Windows) groups that are not allowed to authenticate to the host, and block authentication requests from all others. Groups can be either Linux groups or Windows AD groups.
	OsGroupsBlockeds pulumi.StringArrayInput
	// List of OS (Linux or Windows) users that are allowed to authenticate to the host, and block authentication requests from all others.
	OsUsersAlloweds pulumi.StringArrayInput
	// List of OS (Linux or Windows) users that are not allowed to authenticate to the host, and block authentication requests from all others.
	OsUsersBlockeds pulumi.StringArrayInput
	// Logical expression of how to compute the dependency of the scope variables.
	ScopeExpression pulumi.StringPtrInput
	// List of scope attributes.
	ScopeVariables AquasecHostRuntimePolicyScopeVariableArrayInput
	// Configuration for windows registry monitoring.
	WindowsRegistryMonitoring AquasecHostRuntimePolicyWindowsRegistryMonitoringPtrInput
	// Configuration for windows registry protection.
	WindowsRegistryProtection AquasecHostRuntimePolicyWindowsRegistryProtectionPtrInput
}

func (AquasecHostRuntimePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*aquasecHostRuntimePolicyState)(nil)).Elem()
}

type aquasecHostRuntimePolicyArgs struct {
	// Indicates the application scope of the service.
	ApplicationScopes []string `pulumi:"applicationScopes"`
	// If true, all process activity will be audited.
	AuditAllOsUserActivity *bool `pulumi:"auditAllOsUserActivity"`
	// If true, full command arguments will be audited.
	AuditFullCommandArguments *bool `pulumi:"auditFullCommandArguments"`
	// List of files that are prevented from being read, modified and executed in the containers.
	BlockedFiles []string `pulumi:"blockedFiles"`
	// The description of the host runtime policy
	Description *string `pulumi:"description"`
	// If true, detect and prevent communication from containers to IP addresses known to have a bad reputation.
	EnableIpReputationSecurity *bool `pulumi:"enableIpReputationSecurity"`
	// Indicates if the runtime policy is enabled or not.
	Enabled *bool `pulumi:"enabled"`
	// Indicates that policy should effect container execution (not just for audit).
	Enforce *bool `pulumi:"enforce"`
	// Indicates the number of days after which the runtime policy will be changed to enforce mode.
	EnforceAfterDays *int `pulumi:"enforceAfterDays"`
	// Configuration for file integrity monitoring.
	FileIntegrityMonitoring *AquasecHostRuntimePolicyFileIntegrityMonitoring `pulumi:"fileIntegrityMonitoring"`
	// If true, system time changes will be monitored.
	MonitorSystemTimeChanges *bool `pulumi:"monitorSystemTimeChanges"`
	// If true, windows service operations will be monitored.
	MonitorWindowsServices *bool `pulumi:"monitorWindowsServices"`
	// Name of the host runtime policy
	Name *string `pulumi:"name"`
	// List of OS (Linux or Windows) groups that are allowed to authenticate to the host, and block authentication requests from all others. Groups can be either Linux groups or Windows AD groups.
	OsGroupsAlloweds []string `pulumi:"osGroupsAlloweds"`
	// List of OS (Linux or Windows) groups that are not allowed to authenticate to the host, and block authentication requests from all others. Groups can be either Linux groups or Windows AD groups.
	OsGroupsBlockeds []string `pulumi:"osGroupsBlockeds"`
	// List of OS (Linux or Windows) users that are allowed to authenticate to the host, and block authentication requests from all others.
	OsUsersAlloweds []string `pulumi:"osUsersAlloweds"`
	// List of OS (Linux or Windows) users that are not allowed to authenticate to the host, and block authentication requests from all others.
	OsUsersBlockeds []string `pulumi:"osUsersBlockeds"`
	// Logical expression of how to compute the dependency of the scope variables.
	ScopeExpression *string `pulumi:"scopeExpression"`
	// List of scope attributes.
	ScopeVariables []AquasecHostRuntimePolicyScopeVariable `pulumi:"scopeVariables"`
	// Configuration for windows registry monitoring.
	WindowsRegistryMonitoring *AquasecHostRuntimePolicyWindowsRegistryMonitoring `pulumi:"windowsRegistryMonitoring"`
	// Configuration for windows registry protection.
	WindowsRegistryProtection *AquasecHostRuntimePolicyWindowsRegistryProtection `pulumi:"windowsRegistryProtection"`
}

// The set of arguments for constructing a AquasecHostRuntimePolicy resource.
type AquasecHostRuntimePolicyArgs struct {
	// Indicates the application scope of the service.
	ApplicationScopes pulumi.StringArrayInput
	// If true, all process activity will be audited.
	AuditAllOsUserActivity pulumi.BoolPtrInput
	// If true, full command arguments will be audited.
	AuditFullCommandArguments pulumi.BoolPtrInput
	// List of files that are prevented from being read, modified and executed in the containers.
	BlockedFiles pulumi.StringArrayInput
	// The description of the host runtime policy
	Description pulumi.StringPtrInput
	// If true, detect and prevent communication from containers to IP addresses known to have a bad reputation.
	EnableIpReputationSecurity pulumi.BoolPtrInput
	// Indicates if the runtime policy is enabled or not.
	Enabled pulumi.BoolPtrInput
	// Indicates that policy should effect container execution (not just for audit).
	Enforce pulumi.BoolPtrInput
	// Indicates the number of days after which the runtime policy will be changed to enforce mode.
	EnforceAfterDays pulumi.IntPtrInput
	// Configuration for file integrity monitoring.
	FileIntegrityMonitoring AquasecHostRuntimePolicyFileIntegrityMonitoringPtrInput
	// If true, system time changes will be monitored.
	MonitorSystemTimeChanges pulumi.BoolPtrInput
	// If true, windows service operations will be monitored.
	MonitorWindowsServices pulumi.BoolPtrInput
	// Name of the host runtime policy
	Name pulumi.StringPtrInput
	// List of OS (Linux or Windows) groups that are allowed to authenticate to the host, and block authentication requests from all others. Groups can be either Linux groups or Windows AD groups.
	OsGroupsAlloweds pulumi.StringArrayInput
	// List of OS (Linux or Windows) groups that are not allowed to authenticate to the host, and block authentication requests from all others. Groups can be either Linux groups or Windows AD groups.
	OsGroupsBlockeds pulumi.StringArrayInput
	// List of OS (Linux or Windows) users that are allowed to authenticate to the host, and block authentication requests from all others.
	OsUsersAlloweds pulumi.StringArrayInput
	// List of OS (Linux or Windows) users that are not allowed to authenticate to the host, and block authentication requests from all others.
	OsUsersBlockeds pulumi.StringArrayInput
	// Logical expression of how to compute the dependency of the scope variables.
	ScopeExpression pulumi.StringPtrInput
	// List of scope attributes.
	ScopeVariables AquasecHostRuntimePolicyScopeVariableArrayInput
	// Configuration for windows registry monitoring.
	WindowsRegistryMonitoring AquasecHostRuntimePolicyWindowsRegistryMonitoringPtrInput
	// Configuration for windows registry protection.
	WindowsRegistryProtection AquasecHostRuntimePolicyWindowsRegistryProtectionPtrInput
}

func (AquasecHostRuntimePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aquasecHostRuntimePolicyArgs)(nil)).Elem()
}

type AquasecHostRuntimePolicyInput interface {
	pulumi.Input

	ToAquasecHostRuntimePolicyOutput() AquasecHostRuntimePolicyOutput
	ToAquasecHostRuntimePolicyOutputWithContext(ctx context.Context) AquasecHostRuntimePolicyOutput
}

func (*AquasecHostRuntimePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**AquasecHostRuntimePolicy)(nil)).Elem()
}

func (i *AquasecHostRuntimePolicy) ToAquasecHostRuntimePolicyOutput() AquasecHostRuntimePolicyOutput {
	return i.ToAquasecHostRuntimePolicyOutputWithContext(context.Background())
}

func (i *AquasecHostRuntimePolicy) ToAquasecHostRuntimePolicyOutputWithContext(ctx context.Context) AquasecHostRuntimePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AquasecHostRuntimePolicyOutput)
}

// AquasecHostRuntimePolicyArrayInput is an input type that accepts AquasecHostRuntimePolicyArray and AquasecHostRuntimePolicyArrayOutput values.
// You can construct a concrete instance of `AquasecHostRuntimePolicyArrayInput` via:
//
//          AquasecHostRuntimePolicyArray{ AquasecHostRuntimePolicyArgs{...} }
type AquasecHostRuntimePolicyArrayInput interface {
	pulumi.Input

	ToAquasecHostRuntimePolicyArrayOutput() AquasecHostRuntimePolicyArrayOutput
	ToAquasecHostRuntimePolicyArrayOutputWithContext(context.Context) AquasecHostRuntimePolicyArrayOutput
}

type AquasecHostRuntimePolicyArray []AquasecHostRuntimePolicyInput

func (AquasecHostRuntimePolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AquasecHostRuntimePolicy)(nil)).Elem()
}

func (i AquasecHostRuntimePolicyArray) ToAquasecHostRuntimePolicyArrayOutput() AquasecHostRuntimePolicyArrayOutput {
	return i.ToAquasecHostRuntimePolicyArrayOutputWithContext(context.Background())
}

func (i AquasecHostRuntimePolicyArray) ToAquasecHostRuntimePolicyArrayOutputWithContext(ctx context.Context) AquasecHostRuntimePolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AquasecHostRuntimePolicyArrayOutput)
}

// AquasecHostRuntimePolicyMapInput is an input type that accepts AquasecHostRuntimePolicyMap and AquasecHostRuntimePolicyMapOutput values.
// You can construct a concrete instance of `AquasecHostRuntimePolicyMapInput` via:
//
//          AquasecHostRuntimePolicyMap{ "key": AquasecHostRuntimePolicyArgs{...} }
type AquasecHostRuntimePolicyMapInput interface {
	pulumi.Input

	ToAquasecHostRuntimePolicyMapOutput() AquasecHostRuntimePolicyMapOutput
	ToAquasecHostRuntimePolicyMapOutputWithContext(context.Context) AquasecHostRuntimePolicyMapOutput
}

type AquasecHostRuntimePolicyMap map[string]AquasecHostRuntimePolicyInput

func (AquasecHostRuntimePolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AquasecHostRuntimePolicy)(nil)).Elem()
}

func (i AquasecHostRuntimePolicyMap) ToAquasecHostRuntimePolicyMapOutput() AquasecHostRuntimePolicyMapOutput {
	return i.ToAquasecHostRuntimePolicyMapOutputWithContext(context.Background())
}

func (i AquasecHostRuntimePolicyMap) ToAquasecHostRuntimePolicyMapOutputWithContext(ctx context.Context) AquasecHostRuntimePolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AquasecHostRuntimePolicyMapOutput)
}

type AquasecHostRuntimePolicyOutput struct{ *pulumi.OutputState }

func (AquasecHostRuntimePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AquasecHostRuntimePolicy)(nil)).Elem()
}

func (o AquasecHostRuntimePolicyOutput) ToAquasecHostRuntimePolicyOutput() AquasecHostRuntimePolicyOutput {
	return o
}

func (o AquasecHostRuntimePolicyOutput) ToAquasecHostRuntimePolicyOutputWithContext(ctx context.Context) AquasecHostRuntimePolicyOutput {
	return o
}

type AquasecHostRuntimePolicyArrayOutput struct{ *pulumi.OutputState }

func (AquasecHostRuntimePolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AquasecHostRuntimePolicy)(nil)).Elem()
}

func (o AquasecHostRuntimePolicyArrayOutput) ToAquasecHostRuntimePolicyArrayOutput() AquasecHostRuntimePolicyArrayOutput {
	return o
}

func (o AquasecHostRuntimePolicyArrayOutput) ToAquasecHostRuntimePolicyArrayOutputWithContext(ctx context.Context) AquasecHostRuntimePolicyArrayOutput {
	return o
}

func (o AquasecHostRuntimePolicyArrayOutput) Index(i pulumi.IntInput) AquasecHostRuntimePolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AquasecHostRuntimePolicy {
		return vs[0].([]*AquasecHostRuntimePolicy)[vs[1].(int)]
	}).(AquasecHostRuntimePolicyOutput)
}

type AquasecHostRuntimePolicyMapOutput struct{ *pulumi.OutputState }

func (AquasecHostRuntimePolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AquasecHostRuntimePolicy)(nil)).Elem()
}

func (o AquasecHostRuntimePolicyMapOutput) ToAquasecHostRuntimePolicyMapOutput() AquasecHostRuntimePolicyMapOutput {
	return o
}

func (o AquasecHostRuntimePolicyMapOutput) ToAquasecHostRuntimePolicyMapOutputWithContext(ctx context.Context) AquasecHostRuntimePolicyMapOutput {
	return o
}

func (o AquasecHostRuntimePolicyMapOutput) MapIndex(k pulumi.StringInput) AquasecHostRuntimePolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AquasecHostRuntimePolicy {
		return vs[0].(map[string]*AquasecHostRuntimePolicy)[vs[1].(string)]
	}).(AquasecHostRuntimePolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AquasecHostRuntimePolicyInput)(nil)).Elem(), &AquasecHostRuntimePolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*AquasecHostRuntimePolicyArrayInput)(nil)).Elem(), AquasecHostRuntimePolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AquasecHostRuntimePolicyMapInput)(nil)).Elem(), AquasecHostRuntimePolicyMap{})
	pulumi.RegisterOutputType(AquasecHostRuntimePolicyOutput{})
	pulumi.RegisterOutputType(AquasecHostRuntimePolicyArrayOutput{})
	pulumi.RegisterOutputType(AquasecHostRuntimePolicyMapOutput{})
}