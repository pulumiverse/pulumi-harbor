// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface GetGroupsGroup {
    /**
     * The name of the group to filter by.
     */
    groupName: string;
    groupType: number;
    /**
     * The ID of this resource.
     */
    id: number;
    /**
     * The LDAP group DN to filter by.
     */
    ldapGroupDn: string;
}

export interface GetProjectsProject {
    name: string;
    projectId: number;
    public: boolean;
    type: string;
    vulnerabilityScanning: boolean;
}

export interface ReplicationFilter {
    /**
     * Matches or excludes the result. Can be one of the following. `matches`, `excludes`
     */
    decoration?: string;
    /**
     * Filter on the resource according to labels.
     */
    labels?: string[];
    /**
     * Filter on the name of the resource.
     */
    name?: string;
    /**
     * Filter on the resource type. Can be one of the following types. `chart`, `artifact`
     */
    resource?: string;
    /**
     * Filter on the tag/version of the resource.
     */
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
     * For the repositories excluding.
     */
    repoExcluding?: string;
    /**
     * For the repositories matching.
     */
    repoMatching?: string;
    /**
     * For the tag excluding.
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
    /**
     * Either `system` or `project`.
     */
    kind: string;
    /**
     * namespace is the name of your project. For kind `system` permissions, always use `/` as namespace. Use `*` to match all projects.
     */
    namespace: string;
}

export interface RobotAccountPermissionAccess {
    /**
     * Eg. `push`, `pull`, `read`, etc. Check [available actions](https://github.com/goharbor/harbor/blob/-/src/common/rbac/const.go).
     */
    action: string;
    /**
     * Either `allow` or `deny`. Defaults to `allow`.
     */
    effect?: string;
    /**
     * Eg. `repository`, `labels`, etc. Check [available resources](https://github.com/goharbor/harbor/blob/-/src/common/rbac/const.go).
     */
    resource: string;
}

