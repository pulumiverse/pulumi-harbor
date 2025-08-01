// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 */
export function getUsers(args?: GetUsersArgs, opts?: pulumi.InvokeOptions): Promise<GetUsersResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("harbor:index/getUsers:getUsers", {
        "email": args.email,
        "username": args.username,
    }, opts);
}

/**
 * A collection of arguments for invoking getUsers.
 */
export interface GetUsersArgs {
    /**
     * The email of the user to filter by.
     */
    email?: string;
    /**
     * The name of the user to filter by.
     */
    username?: string;
}

/**
 * A collection of values returned by getUsers.
 */
export interface GetUsersResult {
    /**
     * The email of the user to filter by.
     */
    readonly email?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the user to filter by.
     */
    readonly username?: string;
    readonly users: outputs.GetUsersUser[];
}
/**
 * ## Example Usage
 */
export function getUsersOutput(args?: GetUsersOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetUsersResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("harbor:index/getUsers:getUsers", {
        "email": args.email,
        "username": args.username,
    }, opts);
}

/**
 * A collection of arguments for invoking getUsers.
 */
export interface GetUsersOutputArgs {
    /**
     * The email of the user to filter by.
     */
    email?: pulumi.Input<string>;
    /**
     * The name of the user to filter by.
     */
    username?: pulumi.Input<string>;
}
