// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the aquasec package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'aquasec';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provider.__pulumiType;
    }

    /**
     * This is the base URL of your Aqua instance. Can alternatively be sourced from the `AQUA_URL` environment variable.
     */
    public readonly aquaUrl!: pulumi.Output<string | undefined>;
    /**
     * This is the file path for server CA certificates if they are not available on the host OS. Can alternatively be sourced
     * from the `AQUA_CA_CERT_PATH` environment variable.
     */
    public readonly caCertificatePath!: pulumi.Output<string | undefined>;
    /**
     * This is the file path for Aqua provider configuration. The default configuration path is `~/.aqua/tf.config`. Can
     * alternatively be sourced from the `AQUA_CONFIG` environment variable.
     */
    public readonly configPath!: pulumi.Output<string | undefined>;
    /**
     * This is the password that should be used to make the connection. Can alternatively be sourced from the `AQUA_PASSWORD`
     * environment variable.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * This is the user id that should be used to make the connection. Can alternatively be sourced from the `AQUA_USER`
     * environment variable.
     */
    public readonly username!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            resourceInputs["aquaUrl"] = (args ? args.aquaUrl : undefined) ?? utilities.getEnv("AQUA_URL");
            resourceInputs["caCertificatePath"] = args ? args.caCertificatePath : undefined;
            resourceInputs["configPath"] = args ? args.configPath : undefined;
            resourceInputs["password"] = (args ? args.password : undefined) ?? utilities.getEnv("AUQA_PASSWORD");
            resourceInputs["username"] = (args ? args.username : undefined) ?? utilities.getEnv("AQUA_USERNAME");
            resourceInputs["verifyTls"] = pulumi.output(args ? args.verifyTls : undefined).apply(JSON.stringify);
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * This is the base URL of your Aqua instance. Can alternatively be sourced from the `AQUA_URL` environment variable.
     */
    aquaUrl?: pulumi.Input<string>;
    /**
     * This is the file path for server CA certificates if they are not available on the host OS. Can alternatively be sourced
     * from the `AQUA_CA_CERT_PATH` environment variable.
     */
    caCertificatePath?: pulumi.Input<string>;
    /**
     * This is the file path for Aqua provider configuration. The default configuration path is `~/.aqua/tf.config`. Can
     * alternatively be sourced from the `AQUA_CONFIG` environment variable.
     */
    configPath?: pulumi.Input<string>;
    /**
     * This is the password that should be used to make the connection. Can alternatively be sourced from the `AQUA_PASSWORD`
     * environment variable.
     */
    password?: pulumi.Input<string>;
    /**
     * This is the user id that should be used to make the connection. Can alternatively be sourced from the `AQUA_USER`
     * environment variable.
     */
    username?: pulumi.Input<string>;
    /**
     * If true, server tls certificates will be verified by the client before making a connection. Defaults to true. Can
     * alternatively be sourced from the `AQUA_TLS_VERIFY` environment variable.
     */
    verifyTls?: pulumi.Input<boolean>;
}
