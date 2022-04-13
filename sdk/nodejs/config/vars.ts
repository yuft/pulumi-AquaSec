// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("aquasec");

/**
 * This is the base URL of your Aqua instance. Can alternatively be sourced from the `AQUA_URL` environment variable.
 */
export declare const aquaUrl: string | undefined;
Object.defineProperty(exports, "aquaUrl", {
    get() {
        return __config.get("aquaUrl") ?? utilities.getEnv("AQUA_URL");
    },
    enumerable: true,
});

/**
 * This is the file path for server CA certificates if they are not available on the host OS. Can alternatively be sourced
 * from the `AQUA_CA_CERT_PATH` environment variable.
 */
export declare const caCertificatePath: string | undefined;
Object.defineProperty(exports, "caCertificatePath", {
    get() {
        return __config.get("caCertificatePath");
    },
    enumerable: true,
});

/**
 * This is the file path for Aqua provider configuration. The default configuration path is `~/.aqua/tf.config`. Can
 * alternatively be sourced from the `AQUA_CONFIG` environment variable.
 */
export declare const configPath: string | undefined;
Object.defineProperty(exports, "configPath", {
    get() {
        return __config.get("configPath");
    },
    enumerable: true,
});

/**
 * This is the password that should be used to make the connection. Can alternatively be sourced from the `AQUA_PASSWORD`
 * environment variable.
 */
export declare const password: string | undefined;
Object.defineProperty(exports, "password", {
    get() {
        return __config.get("password") ?? utilities.getEnv("AUQA_PASSWORD");
    },
    enumerable: true,
});

/**
 * This is the user id that should be used to make the connection. Can alternatively be sourced from the `AQUA_USER`
 * environment variable.
 */
export declare const username: string | undefined;
Object.defineProperty(exports, "username", {
    get() {
        return __config.get("username") ?? utilities.getEnv("AQUA_USERNAME");
    },
    enumerable: true,
});

/**
 * If true, server tls certificates will be verified by the client before making a connection. Defaults to true. Can
 * alternatively be sourced from the `AQUA_TLS_VERIFY` environment variable.
 */
export declare const verifyTls: boolean | undefined;
Object.defineProperty(exports, "verifyTls", {
    get() {
        return __config.getObject<boolean>("verifyTls");
    },
    enumerable: true,
});

