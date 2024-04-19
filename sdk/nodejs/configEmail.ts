// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 */
export class ConfigEmail extends pulumi.CustomResource {
    /**
     * Get an existing ConfigEmail resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConfigEmailState, opts?: pulumi.CustomResourceOptions): ConfigEmail {
        return new ConfigEmail(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'harbor:index/configEmail:ConfigEmail';

    /**
     * Returns true if the given object is an instance of ConfigEmail.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConfigEmail {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConfigEmail.__pulumiType;
    }

    /**
     * The email from address ie, `dont_reply@acme.com`
     */
    public readonly emailFrom!: pulumi.Output<string>;
    /**
     * The FQDN of the email server
     */
    public readonly emailHost!: pulumi.Output<string>;
    /**
     * Disables validation of email server certificate `Default: false`
     */
    public readonly emailInsecure!: pulumi.Output<boolean | undefined>;
    /**
     * The password for the email server
     */
    public readonly emailPassword!: pulumi.Output<string | undefined>;
    /**
     * The smtp port for the email server `Default: 25`
     */
    public readonly emailPort!: pulumi.Output<number | undefined>;
    /**
     * Enable SSL for email server connection
     */
    public readonly emailSsl!: pulumi.Output<boolean | undefined>;
    /**
     * The username for the email server
     */
    public readonly emailUsername!: pulumi.Output<string | undefined>;

    /**
     * Create a ConfigEmail resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConfigEmailArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConfigEmailArgs | ConfigEmailState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConfigEmailState | undefined;
            resourceInputs["emailFrom"] = state ? state.emailFrom : undefined;
            resourceInputs["emailHost"] = state ? state.emailHost : undefined;
            resourceInputs["emailInsecure"] = state ? state.emailInsecure : undefined;
            resourceInputs["emailPassword"] = state ? state.emailPassword : undefined;
            resourceInputs["emailPort"] = state ? state.emailPort : undefined;
            resourceInputs["emailSsl"] = state ? state.emailSsl : undefined;
            resourceInputs["emailUsername"] = state ? state.emailUsername : undefined;
        } else {
            const args = argsOrState as ConfigEmailArgs | undefined;
            if ((!args || args.emailFrom === undefined) && !opts.urn) {
                throw new Error("Missing required property 'emailFrom'");
            }
            if ((!args || args.emailHost === undefined) && !opts.urn) {
                throw new Error("Missing required property 'emailHost'");
            }
            resourceInputs["emailFrom"] = args ? args.emailFrom : undefined;
            resourceInputs["emailHost"] = args ? args.emailHost : undefined;
            resourceInputs["emailInsecure"] = args ? args.emailInsecure : undefined;
            resourceInputs["emailPassword"] = args?.emailPassword ? pulumi.secret(args.emailPassword) : undefined;
            resourceInputs["emailPort"] = args ? args.emailPort : undefined;
            resourceInputs["emailSsl"] = args ? args.emailSsl : undefined;
            resourceInputs["emailUsername"] = args ? args.emailUsername : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["emailPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ConfigEmail.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConfigEmail resources.
 */
export interface ConfigEmailState {
    /**
     * The email from address ie, `dont_reply@acme.com`
     */
    emailFrom?: pulumi.Input<string>;
    /**
     * The FQDN of the email server
     */
    emailHost?: pulumi.Input<string>;
    /**
     * Disables validation of email server certificate `Default: false`
     */
    emailInsecure?: pulumi.Input<boolean>;
    /**
     * The password for the email server
     */
    emailPassword?: pulumi.Input<string>;
    /**
     * The smtp port for the email server `Default: 25`
     */
    emailPort?: pulumi.Input<number>;
    /**
     * Enable SSL for email server connection
     */
    emailSsl?: pulumi.Input<boolean>;
    /**
     * The username for the email server
     */
    emailUsername?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ConfigEmail resource.
 */
export interface ConfigEmailArgs {
    /**
     * The email from address ie, `dont_reply@acme.com`
     */
    emailFrom: pulumi.Input<string>;
    /**
     * The FQDN of the email server
     */
    emailHost: pulumi.Input<string>;
    /**
     * Disables validation of email server certificate `Default: false`
     */
    emailInsecure?: pulumi.Input<boolean>;
    /**
     * The password for the email server
     */
    emailPassword?: pulumi.Input<string>;
    /**
     * The smtp port for the email server `Default: 25`
     */
    emailPort?: pulumi.Input<number>;
    /**
     * Enable SSL for email server connection
     */
    emailSsl?: pulumi.Input<boolean>;
    /**
     * The username for the email server
     */
    emailUsername?: pulumi.Input<string>;
}
