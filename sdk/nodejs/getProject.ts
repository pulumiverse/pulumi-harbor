// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 */
export function getProject(args: GetProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("harbor:index/getProject:getProject", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getProject.
 */
export interface GetProjectArgs {
    /**
     * The name of the project.
     */
    name: string;
}

/**
 * A collection of values returned by getProject.
 */
export interface GetProjectResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the project.
     */
    readonly name: string;
    /**
     * The id of the project within harbor.
     */
    readonly projectId: number;
    /**
     * If the project has public accessibility.
     */
    readonly public: boolean;
    /**
     * The type of the project : Project or ProxyCache.
     */
    readonly type: string;
    /**
     * If the images is scanned for vulnerabilities when push to harbor.
     */
    readonly vulnerabilityScanning: boolean;
}
/**
 * ## Example Usage
 */
export function getProjectOutput(args: GetProjectOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectResult> {
    return pulumi.output(args).apply((a: any) => getProject(a, opts))
}

/**
 * A collection of arguments for invoking getProject.
 */
export interface GetProjectOutputArgs {
    /**
     * The name of the project.
     */
    name: pulumi.Input<string>;
}
