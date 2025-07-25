// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 */
export function getRobotAccounts(args?: GetRobotAccountsArgs, opts?: pulumi.InvokeOptions): Promise<GetRobotAccountsResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("harbor:index/getRobotAccounts:getRobotAccounts", {
        "level": args.level,
        "name": args.name,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getRobotAccounts.
 */
export interface GetRobotAccountsArgs {
    /**
     * Level of the robot account, currently either `system` or `project`. Default is `system`.
     */
    level?: string;
    /**
     * The name of the robot account to filter by.
     */
    name?: string;
    /**
     * The id of the project within harbor.
     */
    projectId?: number;
}

/**
 * A collection of values returned by getRobotAccounts.
 */
export interface GetRobotAccountsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Level of the robot account, currently either `system` or `project`. Default is `system`.
     */
    readonly level?: string;
    /**
     * The name of the robot account to filter by.
     */
    readonly name?: string;
    /**
     * The id of the project within harbor.
     */
    readonly projectId?: number;
    readonly robotAccounts: outputs.GetRobotAccountsRobotAccount[];
}
/**
 * ## Example Usage
 */
export function getRobotAccountsOutput(args?: GetRobotAccountsOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRobotAccountsResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("harbor:index/getRobotAccounts:getRobotAccounts", {
        "level": args.level,
        "name": args.name,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getRobotAccounts.
 */
export interface GetRobotAccountsOutputArgs {
    /**
     * Level of the robot account, currently either `system` or `project`. Default is `system`.
     */
    level?: pulumi.Input<string>;
    /**
     * The name of the robot account to filter by.
     */
    name?: pulumi.Input<string>;
    /**
     * The id of the project within harbor.
     */
    projectId?: pulumi.Input<number>;
}
