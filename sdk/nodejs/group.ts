// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ## Import
 *
 * An OIDC group can be imported using the `group id` eg, `<break><break>```sh<break> $ pulumi import harbor:index/group:Group storage-group /usergroups/19 <break>```<break><break>`
 */
export class Group extends pulumi.CustomResource {
    /**
     * Get an existing Group resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupState, opts?: pulumi.CustomResourceOptions): Group {
        return new Group(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'harbor:index/group:Group';

    /**
     * Returns true if the given object is an instance of Group.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Group {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Group.__pulumiType;
    }

    public readonly groupName!: pulumi.Output<string>;
    public readonly groupType!: pulumi.Output<number>;
    public readonly ldapGroupDn!: pulumi.Output<string | undefined>;

    /**
     * Create a Group resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupArgs | GroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupState | undefined;
            resourceInputs["groupName"] = state ? state.groupName : undefined;
            resourceInputs["groupType"] = state ? state.groupType : undefined;
            resourceInputs["ldapGroupDn"] = state ? state.ldapGroupDn : undefined;
        } else {
            const args = argsOrState as GroupArgs | undefined;
            if ((!args || args.groupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupName'");
            }
            if ((!args || args.groupType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupType'");
            }
            resourceInputs["groupName"] = args ? args.groupName : undefined;
            resourceInputs["groupType"] = args ? args.groupType : undefined;
            resourceInputs["ldapGroupDn"] = args ? args.ldapGroupDn : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Group.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Group resources.
 */
export interface GroupState {
    groupName?: pulumi.Input<string>;
    groupType?: pulumi.Input<number>;
    ldapGroupDn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Group resource.
 */
export interface GroupArgs {
    groupName: pulumi.Input<string>;
    groupType: pulumi.Input<number>;
    ldapGroupDn?: pulumi.Input<string>;
}
