// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.harbor;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumiverse.harbor.inputs.RetentionPolicyRuleArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RetentionPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final RetentionPolicyArgs Empty = new RetentionPolicyArgs();

    /**
     * Al collection of rule blocks as documented below.
     * 
     */
    @Import(name="rules", required=true)
    private Output<List<RetentionPolicyRuleArgs>> rules;

    /**
     * @return Al collection of rule blocks as documented below.
     * 
     */
    public Output<List<RetentionPolicyRuleArgs>> rules() {
        return this.rules;
    }

    /**
     * The schedule of when you would like the policy to run. This can be `hourly`, `daily`, `weekly` or can be a custom cron string.
     * 
     */
    @Import(name="schedule")
    private @Nullable Output<String> schedule;

    /**
     * @return The schedule of when you would like the policy to run. This can be `hourly`, `daily`, `weekly` or can be a custom cron string.
     * 
     */
    public Optional<Output<String>> schedule() {
        return Optional.ofNullable(this.schedule);
    }

    /**
     * The project id of which you would like to apply this policy.
     * 
     */
    @Import(name="scope", required=true)
    private Output<String> scope;

    /**
     * @return The project id of which you would like to apply this policy.
     * 
     */
    public Output<String> scope() {
        return this.scope;
    }

    private RetentionPolicyArgs() {}

    private RetentionPolicyArgs(RetentionPolicyArgs $) {
        this.rules = $.rules;
        this.schedule = $.schedule;
        this.scope = $.scope;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RetentionPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RetentionPolicyArgs $;

        public Builder() {
            $ = new RetentionPolicyArgs();
        }

        public Builder(RetentionPolicyArgs defaults) {
            $ = new RetentionPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param rules Al collection of rule blocks as documented below.
         * 
         * @return builder
         * 
         */
        public Builder rules(Output<List<RetentionPolicyRuleArgs>> rules) {
            $.rules = rules;
            return this;
        }

        /**
         * @param rules Al collection of rule blocks as documented below.
         * 
         * @return builder
         * 
         */
        public Builder rules(List<RetentionPolicyRuleArgs> rules) {
            return rules(Output.of(rules));
        }

        /**
         * @param rules Al collection of rule blocks as documented below.
         * 
         * @return builder
         * 
         */
        public Builder rules(RetentionPolicyRuleArgs... rules) {
            return rules(List.of(rules));
        }

        /**
         * @param schedule The schedule of when you would like the policy to run. This can be `hourly`, `daily`, `weekly` or can be a custom cron string.
         * 
         * @return builder
         * 
         */
        public Builder schedule(@Nullable Output<String> schedule) {
            $.schedule = schedule;
            return this;
        }

        /**
         * @param schedule The schedule of when you would like the policy to run. This can be `hourly`, `daily`, `weekly` or can be a custom cron string.
         * 
         * @return builder
         * 
         */
        public Builder schedule(String schedule) {
            return schedule(Output.of(schedule));
        }

        /**
         * @param scope The project id of which you would like to apply this policy.
         * 
         * @return builder
         * 
         */
        public Builder scope(Output<String> scope) {
            $.scope = scope;
            return this;
        }

        /**
         * @param scope The project id of which you would like to apply this policy.
         * 
         * @return builder
         * 
         */
        public Builder scope(String scope) {
            return scope(Output.of(scope));
        }

        public RetentionPolicyArgs build() {
            $.rules = Objects.requireNonNull($.rules, "expected parameter 'rules' to be non-null");
            $.scope = Objects.requireNonNull($.scope, "expected parameter 'scope' to be non-null");
            return $;
        }
    }

}
