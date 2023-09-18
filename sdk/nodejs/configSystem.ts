// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 */
export class ConfigSystem extends pulumi.CustomResource {
    /**
     * Get an existing ConfigSystem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConfigSystemState, opts?: pulumi.CustomResourceOptions): ConfigSystem {
        return new ConfigSystem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'harbor:index/configSystem:ConfigSystem';

    /**
     * Returns true if the given object is an instance of ConfigSystem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConfigSystem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConfigSystem.__pulumiType;
    }

    public readonly projectCreationRestriction!: pulumi.Output<string | undefined>;
    public readonly readOnly!: pulumi.Output<boolean | undefined>;
    public readonly robotNamePrefix!: pulumi.Output<string | undefined>;
    public readonly robotTokenExpiration!: pulumi.Output<number | undefined>;
    public readonly scannerSkipUpdatePulltime!: pulumi.Output<boolean | undefined>;
    public readonly storagePerProject!: pulumi.Output<number | undefined>;

    /**
     * Create a ConfigSystem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ConfigSystemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConfigSystemArgs | ConfigSystemState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConfigSystemState | undefined;
            resourceInputs["projectCreationRestriction"] = state ? state.projectCreationRestriction : undefined;
            resourceInputs["readOnly"] = state ? state.readOnly : undefined;
            resourceInputs["robotNamePrefix"] = state ? state.robotNamePrefix : undefined;
            resourceInputs["robotTokenExpiration"] = state ? state.robotTokenExpiration : undefined;
            resourceInputs["scannerSkipUpdatePulltime"] = state ? state.scannerSkipUpdatePulltime : undefined;
            resourceInputs["storagePerProject"] = state ? state.storagePerProject : undefined;
        } else {
            const args = argsOrState as ConfigSystemArgs | undefined;
            resourceInputs["projectCreationRestriction"] = args ? args.projectCreationRestriction : undefined;
            resourceInputs["readOnly"] = args ? args.readOnly : undefined;
            resourceInputs["robotNamePrefix"] = args ? args.robotNamePrefix : undefined;
            resourceInputs["robotTokenExpiration"] = args ? args.robotTokenExpiration : undefined;
            resourceInputs["scannerSkipUpdatePulltime"] = args ? args.scannerSkipUpdatePulltime : undefined;
            resourceInputs["storagePerProject"] = args ? args.storagePerProject : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ConfigSystem.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConfigSystem resources.
 */
export interface ConfigSystemState {
    projectCreationRestriction?: pulumi.Input<string>;
    readOnly?: pulumi.Input<boolean>;
    robotNamePrefix?: pulumi.Input<string>;
    robotTokenExpiration?: pulumi.Input<number>;
    scannerSkipUpdatePulltime?: pulumi.Input<boolean>;
    storagePerProject?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ConfigSystem resource.
 */
export interface ConfigSystemArgs {
    projectCreationRestriction?: pulumi.Input<string>;
    readOnly?: pulumi.Input<boolean>;
    robotNamePrefix?: pulumi.Input<string>;
    robotTokenExpiration?: pulumi.Input<number>;
    scannerSkipUpdatePulltime?: pulumi.Input<boolean>;
    storagePerProject?: pulumi.Input<number>;
}
