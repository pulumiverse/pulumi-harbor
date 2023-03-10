// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface ReplicationFilter {
    decoration?: string;
    labels?: string[];
    name?: string;
    resource?: string;
    tag?: string;
}

export interface RetentionPolicyRule {
    /**
     * retain always.
     */
    alwaysRetain?: boolean;
    /**
     * Specify if the rule is disable or not. Defaults to `false`
     */
    disabled?: boolean;
    /**
     * retain the most recently pulled n artifacts.
     */
    mostRecentlyPulled?: number;
    /**
     * retain the most recently pushed n artifacts.
     */
    mostRecentlyPushed?: number;
    /**
     * retains the artifacts pulled within the lasts n days.
     */
    nDaysSinceLastPull?: number;
    /**
     * retains the artifacts pushed within the lasts n days.
     */
    nDaysSinceLastPush?: number;
    /**
     * For the repositories excuding.
     */
    repoExcluding?: string;
    /**
     * For the repositories matching.
     */
    repoMatching?: string;
    /**
     * For the tag excuding.
     */
    tagExcluding?: string;
    /**
     * For the tag matching.
     */
    tagMatching?: string;
    /**
     * with untagged artifacts. Defaults to `true`
     */
    untaggedArtifacts?: boolean;
}

export interface RobotAccountPermission {
    accesses: outputs.RobotAccountPermissionAccess[];
    kind: string;
    namespace: string;
}

export interface RobotAccountPermissionAccess {
    action: string;
    effect?: string;
    resource: string;
}

