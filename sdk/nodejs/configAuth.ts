// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ### OIDC
 *
 * ### LDAP
 */
export class ConfigAuth extends pulumi.CustomResource {
    /**
     * Get an existing ConfigAuth resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConfigAuthState, opts?: pulumi.CustomResourceOptions): ConfigAuth {
        return new ConfigAuth(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'harbor:index/configAuth:ConfigAuth';

    /**
     * Returns true if the given object is an instance of ConfigAuth.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConfigAuth {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConfigAuth.__pulumiType;
    }

    /**
     * Harbor authentication mode. Can be `"oidcAuth"`, `"dbAuth"` or `"ldapAuth"`. (Default: `"dbAuth"`)
     */
    public readonly authMode!: pulumi.Output<string>;
    public readonly ldapBaseDn!: pulumi.Output<string | undefined>;
    public readonly ldapFilter!: pulumi.Output<string | undefined>;
    public readonly ldapGroupAdminDn!: pulumi.Output<string | undefined>;
    public readonly ldapGroupBaseDn!: pulumi.Output<string | undefined>;
    public readonly ldapGroupFilter!: pulumi.Output<string | undefined>;
    public readonly ldapGroupGid!: pulumi.Output<string | undefined>;
    public readonly ldapGroupMembership!: pulumi.Output<string | undefined>;
    public readonly ldapGroupScope!: pulumi.Output<string | undefined>;
    public readonly ldapGroupUid!: pulumi.Output<string | undefined>;
    public readonly ldapScope!: pulumi.Output<string | undefined>;
    public readonly ldapSearchDn!: pulumi.Output<string | undefined>;
    public readonly ldapSearchPassword!: pulumi.Output<string | undefined>;
    public readonly ldapUid!: pulumi.Output<string | undefined>;
    public readonly ldapUrl!: pulumi.Output<string | undefined>;
    public readonly ldapVerifyCert!: pulumi.Output<boolean | undefined>;
    public readonly oidcAdminGroup!: pulumi.Output<string | undefined>;
    public readonly oidcAutoOnboard!: pulumi.Output<boolean | undefined>;
    public readonly oidcClientId!: pulumi.Output<string | undefined>;
    public readonly oidcClientSecret!: pulumi.Output<string | undefined>;
    public readonly oidcEndpoint!: pulumi.Output<string | undefined>;
    public readonly oidcGroupFilter!: pulumi.Output<string | undefined>;
    public readonly oidcGroupsClaim!: pulumi.Output<string | undefined>;
    public readonly oidcName!: pulumi.Output<string | undefined>;
    public readonly oidcScope!: pulumi.Output<string | undefined>;
    public readonly oidcUserClaim!: pulumi.Output<string | undefined>;
    public readonly oidcVerifyCert!: pulumi.Output<boolean | undefined>;
    public readonly primaryAuthMode!: pulumi.Output<boolean | undefined>;

    /**
     * Create a ConfigAuth resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConfigAuthArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConfigAuthArgs | ConfigAuthState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConfigAuthState | undefined;
            resourceInputs["authMode"] = state ? state.authMode : undefined;
            resourceInputs["ldapBaseDn"] = state ? state.ldapBaseDn : undefined;
            resourceInputs["ldapFilter"] = state ? state.ldapFilter : undefined;
            resourceInputs["ldapGroupAdminDn"] = state ? state.ldapGroupAdminDn : undefined;
            resourceInputs["ldapGroupBaseDn"] = state ? state.ldapGroupBaseDn : undefined;
            resourceInputs["ldapGroupFilter"] = state ? state.ldapGroupFilter : undefined;
            resourceInputs["ldapGroupGid"] = state ? state.ldapGroupGid : undefined;
            resourceInputs["ldapGroupMembership"] = state ? state.ldapGroupMembership : undefined;
            resourceInputs["ldapGroupScope"] = state ? state.ldapGroupScope : undefined;
            resourceInputs["ldapGroupUid"] = state ? state.ldapGroupUid : undefined;
            resourceInputs["ldapScope"] = state ? state.ldapScope : undefined;
            resourceInputs["ldapSearchDn"] = state ? state.ldapSearchDn : undefined;
            resourceInputs["ldapSearchPassword"] = state ? state.ldapSearchPassword : undefined;
            resourceInputs["ldapUid"] = state ? state.ldapUid : undefined;
            resourceInputs["ldapUrl"] = state ? state.ldapUrl : undefined;
            resourceInputs["ldapVerifyCert"] = state ? state.ldapVerifyCert : undefined;
            resourceInputs["oidcAdminGroup"] = state ? state.oidcAdminGroup : undefined;
            resourceInputs["oidcAutoOnboard"] = state ? state.oidcAutoOnboard : undefined;
            resourceInputs["oidcClientId"] = state ? state.oidcClientId : undefined;
            resourceInputs["oidcClientSecret"] = state ? state.oidcClientSecret : undefined;
            resourceInputs["oidcEndpoint"] = state ? state.oidcEndpoint : undefined;
            resourceInputs["oidcGroupFilter"] = state ? state.oidcGroupFilter : undefined;
            resourceInputs["oidcGroupsClaim"] = state ? state.oidcGroupsClaim : undefined;
            resourceInputs["oidcName"] = state ? state.oidcName : undefined;
            resourceInputs["oidcScope"] = state ? state.oidcScope : undefined;
            resourceInputs["oidcUserClaim"] = state ? state.oidcUserClaim : undefined;
            resourceInputs["oidcVerifyCert"] = state ? state.oidcVerifyCert : undefined;
            resourceInputs["primaryAuthMode"] = state ? state.primaryAuthMode : undefined;
        } else {
            const args = argsOrState as ConfigAuthArgs | undefined;
            if ((!args || args.authMode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authMode'");
            }
            resourceInputs["authMode"] = args ? args.authMode : undefined;
            resourceInputs["ldapBaseDn"] = args ? args.ldapBaseDn : undefined;
            resourceInputs["ldapFilter"] = args ? args.ldapFilter : undefined;
            resourceInputs["ldapGroupAdminDn"] = args ? args.ldapGroupAdminDn : undefined;
            resourceInputs["ldapGroupBaseDn"] = args ? args.ldapGroupBaseDn : undefined;
            resourceInputs["ldapGroupFilter"] = args ? args.ldapGroupFilter : undefined;
            resourceInputs["ldapGroupGid"] = args ? args.ldapGroupGid : undefined;
            resourceInputs["ldapGroupMembership"] = args ? args.ldapGroupMembership : undefined;
            resourceInputs["ldapGroupScope"] = args ? args.ldapGroupScope : undefined;
            resourceInputs["ldapGroupUid"] = args ? args.ldapGroupUid : undefined;
            resourceInputs["ldapScope"] = args ? args.ldapScope : undefined;
            resourceInputs["ldapSearchDn"] = args ? args.ldapSearchDn : undefined;
            resourceInputs["ldapSearchPassword"] = args?.ldapSearchPassword ? pulumi.secret(args.ldapSearchPassword) : undefined;
            resourceInputs["ldapUid"] = args ? args.ldapUid : undefined;
            resourceInputs["ldapUrl"] = args ? args.ldapUrl : undefined;
            resourceInputs["ldapVerifyCert"] = args ? args.ldapVerifyCert : undefined;
            resourceInputs["oidcAdminGroup"] = args ? args.oidcAdminGroup : undefined;
            resourceInputs["oidcAutoOnboard"] = args ? args.oidcAutoOnboard : undefined;
            resourceInputs["oidcClientId"] = args ? args.oidcClientId : undefined;
            resourceInputs["oidcClientSecret"] = args?.oidcClientSecret ? pulumi.secret(args.oidcClientSecret) : undefined;
            resourceInputs["oidcEndpoint"] = args ? args.oidcEndpoint : undefined;
            resourceInputs["oidcGroupFilter"] = args ? args.oidcGroupFilter : undefined;
            resourceInputs["oidcGroupsClaim"] = args ? args.oidcGroupsClaim : undefined;
            resourceInputs["oidcName"] = args ? args.oidcName : undefined;
            resourceInputs["oidcScope"] = args ? args.oidcScope : undefined;
            resourceInputs["oidcUserClaim"] = args ? args.oidcUserClaim : undefined;
            resourceInputs["oidcVerifyCert"] = args ? args.oidcVerifyCert : undefined;
            resourceInputs["primaryAuthMode"] = args ? args.primaryAuthMode : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["ldapSearchPassword", "oidcClientSecret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ConfigAuth.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConfigAuth resources.
 */
export interface ConfigAuthState {
    /**
     * Harbor authentication mode. Can be `"oidcAuth"`, `"dbAuth"` or `"ldapAuth"`. (Default: `"dbAuth"`)
     */
    authMode?: pulumi.Input<string>;
    ldapBaseDn?: pulumi.Input<string>;
    ldapFilter?: pulumi.Input<string>;
    ldapGroupAdminDn?: pulumi.Input<string>;
    ldapGroupBaseDn?: pulumi.Input<string>;
    ldapGroupFilter?: pulumi.Input<string>;
    ldapGroupGid?: pulumi.Input<string>;
    ldapGroupMembership?: pulumi.Input<string>;
    ldapGroupScope?: pulumi.Input<string>;
    ldapGroupUid?: pulumi.Input<string>;
    ldapScope?: pulumi.Input<string>;
    ldapSearchDn?: pulumi.Input<string>;
    ldapSearchPassword?: pulumi.Input<string>;
    ldapUid?: pulumi.Input<string>;
    ldapUrl?: pulumi.Input<string>;
    ldapVerifyCert?: pulumi.Input<boolean>;
    oidcAdminGroup?: pulumi.Input<string>;
    oidcAutoOnboard?: pulumi.Input<boolean>;
    oidcClientId?: pulumi.Input<string>;
    oidcClientSecret?: pulumi.Input<string>;
    oidcEndpoint?: pulumi.Input<string>;
    oidcGroupFilter?: pulumi.Input<string>;
    oidcGroupsClaim?: pulumi.Input<string>;
    oidcName?: pulumi.Input<string>;
    oidcScope?: pulumi.Input<string>;
    oidcUserClaim?: pulumi.Input<string>;
    oidcVerifyCert?: pulumi.Input<boolean>;
    primaryAuthMode?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ConfigAuth resource.
 */
export interface ConfigAuthArgs {
    /**
     * Harbor authentication mode. Can be `"oidcAuth"`, `"dbAuth"` or `"ldapAuth"`. (Default: `"dbAuth"`)
     */
    authMode: pulumi.Input<string>;
    ldapBaseDn?: pulumi.Input<string>;
    ldapFilter?: pulumi.Input<string>;
    ldapGroupAdminDn?: pulumi.Input<string>;
    ldapGroupBaseDn?: pulumi.Input<string>;
    ldapGroupFilter?: pulumi.Input<string>;
    ldapGroupGid?: pulumi.Input<string>;
    ldapGroupMembership?: pulumi.Input<string>;
    ldapGroupScope?: pulumi.Input<string>;
    ldapGroupUid?: pulumi.Input<string>;
    ldapScope?: pulumi.Input<string>;
    ldapSearchDn?: pulumi.Input<string>;
    ldapSearchPassword?: pulumi.Input<string>;
    ldapUid?: pulumi.Input<string>;
    ldapUrl?: pulumi.Input<string>;
    ldapVerifyCert?: pulumi.Input<boolean>;
    oidcAdminGroup?: pulumi.Input<string>;
    oidcAutoOnboard?: pulumi.Input<boolean>;
    oidcClientId?: pulumi.Input<string>;
    oidcClientSecret?: pulumi.Input<string>;
    oidcEndpoint?: pulumi.Input<string>;
    oidcGroupFilter?: pulumi.Input<string>;
    oidcGroupsClaim?: pulumi.Input<string>;
    oidcName?: pulumi.Input<string>;
    oidcScope?: pulumi.Input<string>;
    oidcUserClaim?: pulumi.Input<string>;
    oidcVerifyCert?: pulumi.Input<boolean>;
    primaryAuthMode?: pulumi.Input<boolean>;
}
