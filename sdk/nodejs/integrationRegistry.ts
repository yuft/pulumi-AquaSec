// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class IntegrationRegistry extends pulumi.CustomResource {
    /**
     * Get an existing IntegrationRegistry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IntegrationRegistryState, opts?: pulumi.CustomResourceOptions): IntegrationRegistry {
        return new IntegrationRegistry(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aquasec:index/integrationRegistry:IntegrationRegistry';

    /**
     * Returns true if the given object is an instance of IntegrationRegistry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IntegrationRegistry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IntegrationRegistry.__pulumiType;
    }

    public readonly author!: pulumi.Output<string>;
    public readonly autoPull!: pulumi.Output<boolean | undefined>;
    public readonly autoPullMax!: pulumi.Output<number | undefined>;
    public readonly autoPullTime!: pulumi.Output<string | undefined>;
    public readonly lastUpdated!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly password!: pulumi.Output<string | undefined>;
    public readonly prefixes!: pulumi.Output<string[]>;
    public readonly type!: pulumi.Output<string>;
    public readonly url!: pulumi.Output<string | undefined>;
    public readonly username!: pulumi.Output<string | undefined>;

    /**
     * Create a IntegrationRegistry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationRegistryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IntegrationRegistryArgs | IntegrationRegistryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IntegrationRegistryState | undefined;
            resourceInputs["author"] = state ? state.author : undefined;
            resourceInputs["autoPull"] = state ? state.autoPull : undefined;
            resourceInputs["autoPullMax"] = state ? state.autoPullMax : undefined;
            resourceInputs["autoPullTime"] = state ? state.autoPullTime : undefined;
            resourceInputs["lastUpdated"] = state ? state.lastUpdated : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["prefixes"] = state ? state.prefixes : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as IntegrationRegistryArgs | undefined;
            if ((!args || args.prefixes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'prefixes'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["author"] = args ? args.author : undefined;
            resourceInputs["autoPull"] = args ? args.autoPull : undefined;
            resourceInputs["autoPullMax"] = args ? args.autoPullMax : undefined;
            resourceInputs["autoPullTime"] = args ? args.autoPullTime : undefined;
            resourceInputs["lastUpdated"] = args ? args.lastUpdated : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["prefixes"] = args ? args.prefixes : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IntegrationRegistry.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IntegrationRegistry resources.
 */
export interface IntegrationRegistryState {
    author?: pulumi.Input<string>;
    autoPull?: pulumi.Input<boolean>;
    autoPullMax?: pulumi.Input<number>;
    autoPullTime?: pulumi.Input<string>;
    lastUpdated?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    prefixes?: pulumi.Input<pulumi.Input<string>[]>;
    type?: pulumi.Input<string>;
    url?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IntegrationRegistry resource.
 */
export interface IntegrationRegistryArgs {
    author?: pulumi.Input<string>;
    autoPull?: pulumi.Input<boolean>;
    autoPullMax?: pulumi.Input<number>;
    autoPullTime?: pulumi.Input<string>;
    lastUpdated?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    prefixes: pulumi.Input<pulumi.Input<string>[]>;
    type: pulumi.Input<string>;
    url?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
}
