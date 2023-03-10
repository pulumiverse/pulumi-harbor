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


public final class GroupState extends com.pulumi.resources.ResourceArgs {

    public static final GroupState Empty = new GroupState();

    @Import(name="groupName")
    private @Nullable Output<String> groupName;

    public Optional<Output<String>> groupName() {
        return Optional.ofNullable(this.groupName);
    }

    @Import(name="groupType")
    private @Nullable Output<Integer> groupType;

    public Optional<Output<Integer>> groupType() {
        return Optional.ofNullable(this.groupType);
    }

    private GroupState() {}

    private GroupState(GroupState $) {
        this.groupName = $.groupName;
        this.groupType = $.groupType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupState $;

        public Builder() {
            $ = new GroupState();
        }

        public Builder(GroupState defaults) {
            $ = new GroupState(Objects.requireNonNull(defaults));
        }

        public Builder groupName(@Nullable Output<String> groupName) {
            $.groupName = groupName;
            return this;
        }

        public Builder groupName(String groupName) {
            return groupName(Output.of(groupName));
        }

        public Builder groupType(@Nullable Output<Integer> groupType) {
            $.groupType = groupType;
            return this;
        }

        public Builder groupType(Integer groupType) {
            return groupType(Output.of(groupType));
        }

        public GroupState build() {
            return $;
        }
    }

}
