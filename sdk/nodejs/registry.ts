// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ## Import
 *
 * Harbor project can be imported using the `registry id` eg, `<break><break>```sh<break> $ pulumi import harbor:index/registry:Registry main /registries/7 <break>```<break><break>`
 */
export class Registry extends pulumi.CustomResource {
    /**
     * Get an existing Registry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegistryState, opts?: pulumi.CustomResourceOptions): Registry {
        return new Registry(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'harbor:index/registry:Registry';

    /**
     * Returns true if the given object is an instance of Registry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Registry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Registry.__pulumiType;
    }

    public readonly accessId!: pulumi.Output<string | undefined>;
    public readonly accessSecret!: pulumi.Output<string | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly endpointUrl!: pulumi.Output<string>;
    public readonly insecure!: pulumi.Output<boolean | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly providerName!: pulumi.Output<string>;
    public /*out*/ readonly registryId!: pulumi.Output<number>;
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a Registry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegistryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegistryArgs | RegistryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RegistryState | undefined;
            resourceInputs["accessId"] = state ? state.accessId : undefined;
            resourceInputs["accessSecret"] = state ? state.accessSecret : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["endpointUrl"] = state ? state.endpointUrl : undefined;
            resourceInputs["insecure"] = state ? state.insecure : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["providerName"] = state ? state.providerName : undefined;
            resourceInputs["registryId"] = state ? state.registryId : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as RegistryArgs | undefined;
            if ((!args || args.endpointUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpointUrl'");
            }
            if ((!args || args.providerName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'providerName'");
            }
            resourceInputs["accessId"] = args ? args.accessId : undefined;
            resourceInputs["accessSecret"] = args?.accessSecret ? pulumi.secret(args.accessSecret) : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["endpointUrl"] = args ? args.endpointUrl : undefined;
            resourceInputs["insecure"] = args ? args.insecure : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["providerName"] = args ? args.providerName : undefined;
            resourceInputs["registryId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["accessSecret"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Registry.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Registry resources.
 */
export interface RegistryState {
    accessId?: pulumi.Input<string>;
    accessSecret?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    endpointUrl?: pulumi.Input<string>;
    insecure?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    providerName?: pulumi.Input<string>;
    registryId?: pulumi.Input<number>;
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Registry resource.
 */
export interface RegistryArgs {
    accessId?: pulumi.Input<string>;
    accessSecret?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    endpointUrl: pulumi.Input<string>;
    insecure?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    providerName: pulumi.Input<string>;
}
