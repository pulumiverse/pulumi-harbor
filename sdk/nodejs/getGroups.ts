// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 */
export function getGroups(args?: GetGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupsResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("harbor:index/getGroups:getGroups", {
        "groupName": args.groupName,
        "ldapGroupDn": args.ldapGroupDn,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroups.
 */
export interface GetGroupsArgs {
    /**
     * The name of the group to filter by.
     */
    groupName?: string;
    /**
     * The LDAP group DN to filter by.
     */
    ldapGroupDn?: string;
}

/**
 * A collection of values returned by getGroups.
 */
export interface GetGroupsResult {
    /**
     * The name of the group to filter by.
     */
    readonly groupName?: string;
    readonly groups: outputs.GetGroupsGroup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The LDAP group DN to filter by.
     */
    readonly ldapGroupDn?: string;
}
/**
 * ## Example Usage
 */
export function getGroupsOutput(args?: GetGroupsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetGroupsResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("harbor:index/getGroups:getGroups", {
        "groupName": args.groupName,
        "ldapGroupDn": args.ldapGroupDn,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroups.
 */
export interface GetGroupsOutputArgs {
    /**
     * The name of the group to filter by.
     */
    groupName?: pulumi.Input<string>;
    /**
     * The LDAP group DN to filter by.
     */
    ldapGroupDn?: pulumi.Input<string>;
}
