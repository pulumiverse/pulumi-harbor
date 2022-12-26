// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.harbor.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProjectMemberGroupState extends com.pulumi.resources.ResourceArgs {

    public static final ProjectMemberGroupState Empty = new ProjectMemberGroupState();

    @Import(name="groupId")
    private @Nullable Output<Integer> groupId;

    public Optional<Output<Integer>> groupId() {
        return Optional.ofNullable(this.groupId);
    }

    @Import(name="groupName")
    private @Nullable Output<String> groupName;

    public Optional<Output<String>> groupName() {
        return Optional.ofNullable(this.groupName);
    }

    @Import(name="ldapGroupDn")
    private @Nullable Output<String> ldapGroupDn;

    public Optional<Output<String>> ldapGroupDn() {
        return Optional.ofNullable(this.ldapGroupDn);
    }

    @Import(name="memberId")
    private @Nullable Output<Integer> memberId;

    public Optional<Output<Integer>> memberId() {
        return Optional.ofNullable(this.memberId);
    }

    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    @Import(name="role")
    private @Nullable Output<String> role;

    public Optional<Output<String>> role() {
        return Optional.ofNullable(this.role);
    }

    @Import(name="type")
    private @Nullable Output<String> type;

    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private ProjectMemberGroupState() {}

    private ProjectMemberGroupState(ProjectMemberGroupState $) {
        this.groupId = $.groupId;
        this.groupName = $.groupName;
        this.ldapGroupDn = $.ldapGroupDn;
        this.memberId = $.memberId;
        this.projectId = $.projectId;
        this.role = $.role;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectMemberGroupState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectMemberGroupState $;

        public Builder() {
            $ = new ProjectMemberGroupState();
        }

        public Builder(ProjectMemberGroupState defaults) {
            $ = new ProjectMemberGroupState(Objects.requireNonNull(defaults));
        }

        public Builder groupId(@Nullable Output<Integer> groupId) {
            $.groupId = groupId;
            return this;
        }

        public Builder groupId(Integer groupId) {
            return groupId(Output.of(groupId));
        }

        public Builder groupName(@Nullable Output<String> groupName) {
            $.groupName = groupName;
            return this;
        }

        public Builder groupName(String groupName) {
            return groupName(Output.of(groupName));
        }

        public Builder ldapGroupDn(@Nullable Output<String> ldapGroupDn) {
            $.ldapGroupDn = ldapGroupDn;
            return this;
        }

        public Builder ldapGroupDn(String ldapGroupDn) {
            return ldapGroupDn(Output.of(ldapGroupDn));
        }

        public Builder memberId(@Nullable Output<Integer> memberId) {
            $.memberId = memberId;
            return this;
        }

        public Builder memberId(Integer memberId) {
            return memberId(Output.of(memberId));
        }

        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        public Builder role(@Nullable Output<String> role) {
            $.role = role;
            return this;
        }

        public Builder role(String role) {
            return role(Output.of(role));
        }

        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        public Builder type(String type) {
            return type(Output.of(type));
        }

        public ProjectMemberGroupState build() {
            return $;
        }
    }

}
