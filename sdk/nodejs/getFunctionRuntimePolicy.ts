// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getFunctionRuntimePolicy(args: GetFunctionRuntimePolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetFunctionRuntimePolicyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aquasec:index/getFunctionRuntimePolicy:getFunctionRuntimePolicy", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getFunctionRuntimePolicy.
 */
export interface GetFunctionRuntimePolicyArgs {
    /**
     * Name of the container runtime policy
     */
    name: string;
}

/**
 * A collection of values returned by getFunctionRuntimePolicy.
 */
export interface GetFunctionRuntimePolicyResult {
    /**
     * Indicates the application scope of the service.
     */
    readonly applicationScopes: string[];
    /**
     * Username of the account that created the service.
     */
    readonly author: string;
    /**
     * If true, prevent creation of malicious executables in functions during their runtime post invocation.
     */
    readonly blockMaliciousExecutables: boolean;
    /**
     * List of executables that are prevented from running in containers.
     */
    readonly blockedExecutables: string[];
    /**
     * The description of the container runtime policy
     */
    readonly description: string;
    /**
     * Indicates if the runtime policy is enabled or not.
     */
    readonly enabled: boolean;
    /**
     * Indicates that policy should effect container execution (not just for audit).
     */
    readonly enforce: boolean;
    /**
     * Honeypot User ID (Access Key)
     */
    readonly honeypotAccessKey: string;
    /**
     * List of options to apply the honeypot on (Environment Vairable, Layer, File)
     */
    readonly honeypotApplyOns: string[];
    /**
     * Honeypot User Password (Secret Key)
     */
    readonly honeypotSecretKey: string;
    /**
     * Serverless application name
     */
    readonly honeypotServerlessAppName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Name of the container runtime policy
     */
    readonly name: string;
    /**
     * Logical expression of how to compute the dependency of the scope variables.
     */
    readonly scopeExpression: string;
    /**
     * List of scope attributes.
     */
    readonly scopeVariables: outputs.GetFunctionRuntimePolicyScopeVariable[];
}

export function getFunctionRuntimePolicyOutput(args: GetFunctionRuntimePolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFunctionRuntimePolicyResult> {
    return pulumi.output(args).apply(a => getFunctionRuntimePolicy(a, opts))
}

/**
 * A collection of arguments for invoking getFunctionRuntimePolicy.
 */
export interface GetFunctionRuntimePolicyOutputArgs {
    /**
     * Name of the container runtime policy
     */
    name: pulumi.Input<string>;
}
