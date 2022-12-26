// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.harbor;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumiverse.harbor.inputs.RobotAccountPermissionArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RobotAccountArgs extends com.pulumi.resources.ResourceArgs {

    public static final RobotAccountArgs Empty = new RobotAccountArgs();

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="disable")
    private @Nullable Output<Boolean> disable;

    public Optional<Output<Boolean>> disable() {
        return Optional.ofNullable(this.disable);
    }

    @Import(name="duration")
    private @Nullable Output<Integer> duration;

    public Optional<Output<Integer>> duration() {
        return Optional.ofNullable(this.duration);
    }

    @Import(name="level", required=true)
    private Output<String> level;

    public Output<String> level() {
        return this.level;
    }

    @Import(name="name")
    private @Nullable Output<String> name;

    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="permissions", required=true)
    private Output<List<RobotAccountPermissionArgs>> permissions;

    public Output<List<RobotAccountPermissionArgs>> permissions() {
        return this.permissions;
    }

    @Import(name="secret")
    private @Nullable Output<String> secret;

    public Optional<Output<String>> secret() {
        return Optional.ofNullable(this.secret);
    }

    private RobotAccountArgs() {}

    private RobotAccountArgs(RobotAccountArgs $) {
        this.description = $.description;
        this.disable = $.disable;
        this.duration = $.duration;
        this.level = $.level;
        this.name = $.name;
        this.permissions = $.permissions;
        this.secret = $.secret;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RobotAccountArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RobotAccountArgs $;

        public Builder() {
            $ = new RobotAccountArgs();
        }

        public Builder(RobotAccountArgs defaults) {
            $ = new RobotAccountArgs(Objects.requireNonNull(defaults));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder disable(@Nullable Output<Boolean> disable) {
            $.disable = disable;
            return this;
        }

        public Builder disable(Boolean disable) {
            return disable(Output.of(disable));
        }

        public Builder duration(@Nullable Output<Integer> duration) {
            $.duration = duration;
            return this;
        }

        public Builder duration(Integer duration) {
            return duration(Output.of(duration));
        }

        public Builder level(Output<String> level) {
            $.level = level;
            return this;
        }

        public Builder level(String level) {
            return level(Output.of(level));
        }

        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder permissions(Output<List<RobotAccountPermissionArgs>> permissions) {
            $.permissions = permissions;
            return this;
        }

        public Builder permissions(List<RobotAccountPermissionArgs> permissions) {
            return permissions(Output.of(permissions));
        }

        public Builder permissions(RobotAccountPermissionArgs... permissions) {
            return permissions(List.of(permissions));
        }

        public Builder secret(@Nullable Output<String> secret) {
            $.secret = secret;
            return this;
        }

        public Builder secret(String secret) {
            return secret(Output.of(secret));
        }

        public RobotAccountArgs build() {
            $.level = Objects.requireNonNull($.level, "expected parameter 'level' to be non-null");
            $.permissions = Objects.requireNonNull($.permissions, "expected parameter 'permissions' to be non-null");
            return $;
        }
    }

}
