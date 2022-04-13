// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class NotificationSlack extends pulumi.CustomResource {
    /**
     * Get an existing NotificationSlack resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NotificationSlackState, opts?: pulumi.CustomResourceOptions): NotificationSlack {
        return new NotificationSlack(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aquasec:index/notificationSlack:NotificationSlack';

    /**
     * Returns true if the given object is an instance of NotificationSlack.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NotificationSlack {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NotificationSlack.__pulumiType;
    }

    public readonly channel!: pulumi.Output<string>;
    public readonly enabled!: pulumi.Output<boolean>;
    public readonly icon!: pulumi.Output<string | undefined>;
    public readonly mainText!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly serviceKey!: pulumi.Output<string | undefined>;
    public readonly type!: pulumi.Output<string>;
    public readonly userName!: pulumi.Output<string>;
    public readonly webhookUrl!: pulumi.Output<string>;

    /**
     * Create a NotificationSlack resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NotificationSlackArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NotificationSlackArgs | NotificationSlackState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NotificationSlackState | undefined;
            resourceInputs["channel"] = state ? state.channel : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["icon"] = state ? state.icon : undefined;
            resourceInputs["mainText"] = state ? state.mainText : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["serviceKey"] = state ? state.serviceKey : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
            resourceInputs["webhookUrl"] = state ? state.webhookUrl : undefined;
        } else {
            const args = argsOrState as NotificationSlackArgs | undefined;
            if ((!args || args.channel === undefined) && !opts.urn) {
                throw new Error("Missing required property 'channel'");
            }
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            if ((!args || args.webhookUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'webhookUrl'");
            }
            resourceInputs["channel"] = args ? args.channel : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["icon"] = args ? args.icon : undefined;
            resourceInputs["mainText"] = args ? args.mainText : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["serviceKey"] = args ? args.serviceKey : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
            resourceInputs["webhookUrl"] = args ? args.webhookUrl : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NotificationSlack.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NotificationSlack resources.
 */
export interface NotificationSlackState {
    channel?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    icon?: pulumi.Input<string>;
    mainText?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    serviceKey?: pulumi.Input<string>;
    type?: pulumi.Input<string>;
    userName?: pulumi.Input<string>;
    webhookUrl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NotificationSlack resource.
 */
export interface NotificationSlackArgs {
    channel: pulumi.Input<string>;
    enabled: pulumi.Input<boolean>;
    icon?: pulumi.Input<string>;
    mainText?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    serviceKey?: pulumi.Input<string>;
    type: pulumi.Input<string>;
    userName: pulumi.Input<string>;
    webhookUrl: pulumi.Input<string>;
}