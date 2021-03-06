// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `aquasec.PermissionsSets` resource manages your Permission Set within Aqua.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aquasec from "@pulumi/aquasec";
 *
 * const myTerraformPermSet = new aquasec.PermissionsSets("my_terraform_perm_set", {
 *     actions: [
 *         "dashboard.read",
 *         "risks.vulnerabilities.read",
 *         "risks.vulnerabilities.write",
 *         "risks.host_images.read",
 *         "risks.benchmark.read",
 *         "risk_explorer.read",
 *         "images.read",
 *         "image_profiles.read",
 *         "image_assurance.read",
 *         "image_assurance.write",
 *         "runtime_policies.read",
 *         "runtime_policies.write",
 *         "functions.read",
 *         "gateways.read",
 *         "secrets.read",
 *         "audits.read",
 *         "containers.read",
 *         "enforcers.read",
 *         "infrastructure.read",
 *         "consoles.read",
 *         "settings.read",
 *         "network_policies.read",
 *         "acl_policies.read",
 *         "acl_policies.write",
 *         "services.read",
 *         "integrations.read",
 *         "registries_integrations.read",
 *         "web_hook.read",
 *         "incidents.read",
 *     ],
 *     author: "system",
 *     description: "created from terraform",
 *     isSuper: false,
 *     uiAccess: true,
 * });
 * ```
 */
export class PermissionsSets extends pulumi.CustomResource {
    /**
     * Get an existing PermissionsSets resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PermissionsSetsState, opts?: pulumi.CustomResourceOptions): PermissionsSets {
        return new PermissionsSets(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aquasec:index/permissionsSets:PermissionsSets';

    /**
     * Returns true if the given object is an instance of PermissionsSets.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PermissionsSets {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PermissionsSets.__pulumiType;
    }

    public readonly actions!: pulumi.Output<string[]>;
    public readonly author!: pulumi.Output<string | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of this resource.
     */
    public readonly id!: pulumi.Output<string>;
    public readonly isSuper!: pulumi.Output<boolean | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly uiAccess!: pulumi.Output<boolean>;
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a PermissionsSets resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PermissionsSetsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PermissionsSetsArgs | PermissionsSetsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PermissionsSetsState | undefined;
            resourceInputs["actions"] = state ? state.actions : undefined;
            resourceInputs["author"] = state ? state.author : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["id"] = state ? state.id : undefined;
            resourceInputs["isSuper"] = state ? state.isSuper : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["uiAccess"] = state ? state.uiAccess : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as PermissionsSetsArgs | undefined;
            if ((!args || args.actions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'actions'");
            }
            if ((!args || args.uiAccess === undefined) && !opts.urn) {
                throw new Error("Missing required property 'uiAccess'");
            }
            resourceInputs["actions"] = args ? args.actions : undefined;
            resourceInputs["author"] = args ? args.author : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["id"] = args ? args.id : undefined;
            resourceInputs["isSuper"] = args ? args.isSuper : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["uiAccess"] = args ? args.uiAccess : undefined;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PermissionsSets.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PermissionsSets resources.
 */
export interface PermissionsSetsState {
    actions?: pulumi.Input<pulumi.Input<string>[]>;
    author?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    /**
     * The ID of this resource.
     */
    id?: pulumi.Input<string>;
    isSuper?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    uiAccess?: pulumi.Input<boolean>;
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PermissionsSets resource.
 */
export interface PermissionsSetsArgs {
    actions: pulumi.Input<pulumi.Input<string>[]>;
    author?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    /**
     * The ID of this resource.
     */
    id?: pulumi.Input<string>;
    isSuper?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    uiAccess: pulumi.Input<boolean>;
}
