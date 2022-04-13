// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getHostRuntimePolicy(args: GetHostRuntimePolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetHostRuntimePolicyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aquasec:index/getHostRuntimePolicy:getHostRuntimePolicy", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getHostRuntimePolicy.
 */
export interface GetHostRuntimePolicyArgs {
    /**
     * Name of the host runtime policy
     */
    name: string;
}

/**
 * A collection of values returned by getHostRuntimePolicy.
 */
export interface GetHostRuntimePolicyResult {
    /**
     * Indicates the application scope of the service.
     */
    readonly applicationScopes: string[];
    /**
     * If true, all process activity will be audited.
     */
    readonly auditAllOsUserActivity: boolean;
    /**
     * If true, full command arguments will be audited.
     */
    readonly auditFullCommandArguments: boolean;
    /**
     * Username of the account that created the service.
     */
    readonly author: string;
    /**
     * List of files that are prevented from being read, modified and executed in the containers.
     */
    readonly blockedFiles: string[];
    /**
     * The description of the host runtime policy
     */
    readonly description: string;
    /**
     * If true, detect and prevent communication from containers to IP addresses known to have a bad reputation.
     */
    readonly enableIpReputationSecurity: boolean;
    /**
     * Indicates if the runtime policy is enabled or not.
     */
    readonly enabled: boolean;
    /**
     * Indicates that policy should effect container execution (not just for audit).
     */
    readonly enforce: boolean;
    /**
     * Indicates the number of days after which the runtime policy will be changed to enforce mode.
     */
    readonly enforceAfterDays: number;
    /**
     * Configuration for file integrity monitoring.
     */
    readonly fileIntegrityMonitorings: outputs.GetHostRuntimePolicyFileIntegrityMonitoring[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * If true, system time changes will be monitored.
     */
    readonly monitorSystemTimeChanges: boolean;
    /**
     * If true, windows service operations will be monitored.
     */
    readonly monitorWindowsServices: boolean;
    /**
     * Name of the host runtime policy
     */
    readonly name: string;
    /**
     * List of OS (Linux or Windows) groups that are allowed to authenticate to the host, and block authentication requests from all others. Groups can be either Linux groups or Windows AD groups.
     */
    readonly osGroupsAlloweds: string[];
    /**
     * List of OS (Linux or Windows) groups that are not allowed to authenticate to the host, and block authentication requests from all others. Groups can be either Linux groups or Windows AD groups.
     */
    readonly osGroupsBlockeds: string[];
    /**
     * List of OS (Linux or Windows) users that are allowed to authenticate to the host, and block authentication requests from all others.
     */
    readonly osUsersAlloweds: string[];
    /**
     * List of OS (Linux or Windows) users that are not allowed to authenticate to the host, and block authentication requests from all others.
     */
    readonly osUsersBlockeds: string[];
    /**
     * Logical expression of how to compute the dependency of the scope variables.
     */
    readonly scopeExpression: string;
    /**
     * List of scope attributes.
     */
    readonly scopeVariables: outputs.GetHostRuntimePolicyScopeVariable[];
    /**
     * Configuration for windows registry monitoring.
     */
    readonly windowsRegistryMonitorings: outputs.GetHostRuntimePolicyWindowsRegistryMonitoring[];
    /**
     * Configuration for windows registry protection.
     */
    readonly windowsRegistryProtections: outputs.GetHostRuntimePolicyWindowsRegistryProtection[];
}

export function getHostRuntimePolicyOutput(args: GetHostRuntimePolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetHostRuntimePolicyResult> {
    return pulumi.output(args).apply(a => getHostRuntimePolicy(a, opts))
}

/**
 * A collection of arguments for invoking getHostRuntimePolicy.
 */
export interface GetHostRuntimePolicyOutputArgs {
    /**
     * Name of the host runtime policy
     */
    name: pulumi.Input<string>;
}
