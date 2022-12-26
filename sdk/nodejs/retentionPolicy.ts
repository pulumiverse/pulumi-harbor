// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as harbor from "@pulumiverse/harbor";
 *
 * const mainProject = new harbor.Project("mainProject", {});
 * const mainRetentionPolicy = new harbor.RetentionPolicy("mainRetentionPolicy", {
 *     scope: mainProject.id,
 *     schedule: "daily",
 *     rules: [
 *         {
 *             nDaysSinceLastPull: 5,
 *             repoMatching: "**",
 *             tagMatching: "latest",
 *         },
 *         {
 *             nDaysSinceLastPush: 10,
 *             repoMatching: "**",
 *             tagMatching: "latest",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Harbor retention policy can be imported using the `retention_policy id` eg, `
 *
 * ```sh
 *  $ pulumi import harbor:index/retentionPolicy:RetentionPolicy main /retentions/10
 * ```
 *
 *  `
 */
export class RetentionPolicy extends pulumi.CustomResource {
    /**
     * Get an existing RetentionPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RetentionPolicyState, opts?: pulumi.CustomResourceOptions): RetentionPolicy {
        return new RetentionPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'harbor:index/retentionPolicy:RetentionPolicy';

    /**
     * Returns true if the given object is an instance of RetentionPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RetentionPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RetentionPolicy.__pulumiType;
    }

    /**
     * Al collection of rule blocks as documented below.
     */
    public readonly rules!: pulumi.Output<outputs.RetentionPolicyRule[]>;
    /**
     * The schedule of when you would like the policy to run. This can be `hourly`, `daily`, `weekly` or can be a custom cron string.
     */
    public readonly schedule!: pulumi.Output<string | undefined>;
    /**
     * The project id of which you would like to apply this policy.
     */
    public readonly scope!: pulumi.Output<string>;

    /**
     * Create a RetentionPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RetentionPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RetentionPolicyArgs | RetentionPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RetentionPolicyState | undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["schedule"] = state ? state.schedule : undefined;
            resourceInputs["scope"] = state ? state.scope : undefined;
        } else {
            const args = argsOrState as RetentionPolicyArgs | undefined;
            if ((!args || args.rules === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rules'");
            }
            if ((!args || args.scope === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scope'");
            }
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["schedule"] = args ? args.schedule : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RetentionPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RetentionPolicy resources.
 */
export interface RetentionPolicyState {
    /**
     * Al collection of rule blocks as documented below.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.RetentionPolicyRule>[]>;
    /**
     * The schedule of when you would like the policy to run. This can be `hourly`, `daily`, `weekly` or can be a custom cron string.
     */
    schedule?: pulumi.Input<string>;
    /**
     * The project id of which you would like to apply this policy.
     */
    scope?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RetentionPolicy resource.
 */
export interface RetentionPolicyArgs {
    /**
     * Al collection of rule blocks as documented below.
     */
    rules: pulumi.Input<pulumi.Input<inputs.RetentionPolicyRule>[]>;
    /**
     * The schedule of when you would like the policy to run. This can be `hourly`, `daily`, `weekly` or can be a custom cron string.
     */
    schedule?: pulumi.Input<string>;
    /**
     * The project id of which you would like to apply this policy.
     */
    scope: pulumi.Input<string>;
}
