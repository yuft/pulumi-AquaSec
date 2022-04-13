// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aquasec from "@pulumi/aquasec";
 *
 * const terraformiap = new aquasec.ApplicationScope("terraformiap", {
 *     // Categories is a nested block of artifacts, workloads and infrastructure
 *     categories: [{
 *         // Artifacts is a nested block of Image, Function, CF
 *         artifacts: [{
 *             // Every object requires expression(logical combinations of variables v1, v2, v3...) and list of variables consists of attribute(pre-defined) and value
 *             images: [{
 *                 expression: "v1 && v2",
 *                 variables: [
 *                     {
 *                         attribute: "aqua.registry",
 *                         value: "test-registry",
 *                     },
 *                     {
 *                         attribute: "image.repo",
 *                         value: "nginx",
 *                     },
 *                 ],
 *             }],
 *         }],
 *         // Infrastructure is a nested block of Kubernetes, OS
 *         infrastructures: [{
 *             // Every object requires expression and list of variables consists of attribute(pre-defined) and value
 *             kubernetes: [{
 *                 expression: "v1",
 *                 variables: [{
 *                     attribute: "kubernetes.cluster",
 *                     value: "aqua",
 *                 }],
 *             }],
 *         }],
 *         // Workloads is a nested block of Kubernetes, OS, CF
 *         workloads: [{
 *             // Every object requires expression(logical combinations of variables v1, v2, v3...) and list of variables consists of attribute(pre-defined) and value
 *             kubernetes: [{
 *                 expression: "v1 && v2",
 *                 variables: [
 *                     {
 *                         attribute: "kubernetes.cluster",
 *                         value: "aqua",
 *                     },
 *                     {
 *                         attribute: "kubernetes.namespace",
 *                         value: "aqua",
 *                     },
 *                 ],
 *             }],
 *         }],
 *     }],
 *     description: "test123",
 * });
 * ```
 */
export class ApplicationScope extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationScope resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationScopeState, opts?: pulumi.CustomResourceOptions): ApplicationScope {
        return new ApplicationScope(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aquasec:index/applicationScope:ApplicationScope';

    /**
     * Returns true if the given object is an instance of ApplicationScope.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationScope {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationScope.__pulumiType;
    }

    public /*out*/ readonly author!: pulumi.Output<string>;
    public readonly categories!: pulumi.Output<outputs.ApplicationScopeCategory[] | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly ownerEmail!: pulumi.Output<string | undefined>;

    /**
     * Create a ApplicationScope resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ApplicationScopeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationScopeArgs | ApplicationScopeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationScopeState | undefined;
            resourceInputs["author"] = state ? state.author : undefined;
            resourceInputs["categories"] = state ? state.categories : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["ownerEmail"] = state ? state.ownerEmail : undefined;
        } else {
            const args = argsOrState as ApplicationScopeArgs | undefined;
            resourceInputs["categories"] = args ? args.categories : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ownerEmail"] = args ? args.ownerEmail : undefined;
            resourceInputs["author"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApplicationScope.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApplicationScope resources.
 */
export interface ApplicationScopeState {
    author?: pulumi.Input<string>;
    categories?: pulumi.Input<pulumi.Input<inputs.ApplicationScopeCategory>[]>;
    description?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    ownerEmail?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApplicationScope resource.
 */
export interface ApplicationScopeArgs {
    categories?: pulumi.Input<pulumi.Input<inputs.ApplicationScopeCategory>[]>;
    description?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    ownerEmail?: pulumi.Input<string>;
}