// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * The data source `aquasec.getGroups` provides a method to query all groups within the Aqua CSPMgroup database. The fields returned from this query are detailed in the Schema section below.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aquasec from "@pulumi/aquasec";
 *
 * const groups = aquasec.getGroups({});
 * export const firstGroupName = groups.then(groups => groups.groups?[0]?.name);
 * ```
 */
export function getGroups(opts?: pulumi.InvokeOptions): Promise<GetGroupsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aquasec:index/getGroups:getGroups", {
    }, opts);
}

/**
 * A collection of values returned by getGroups.
 */
export interface GetGroupsResult {
    readonly groups: outputs.GetGroupsGroup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}