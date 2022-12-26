// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.harbor;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class TasksArgs extends com.pulumi.resources.ResourceArgs {

    public static final TasksArgs Empty = new TasksArgs();

    @Import(name="vulnerabilityScanPolicy", required=true)
    private Output<String> vulnerabilityScanPolicy;

    public Output<String> vulnerabilityScanPolicy() {
        return this.vulnerabilityScanPolicy;
    }

    private TasksArgs() {}

    private TasksArgs(TasksArgs $) {
        this.vulnerabilityScanPolicy = $.vulnerabilityScanPolicy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TasksArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TasksArgs $;

        public Builder() {
            $ = new TasksArgs();
        }

        public Builder(TasksArgs defaults) {
            $ = new TasksArgs(Objects.requireNonNull(defaults));
        }

        public Builder vulnerabilityScanPolicy(Output<String> vulnerabilityScanPolicy) {
            $.vulnerabilityScanPolicy = vulnerabilityScanPolicy;
            return this;
        }

        public Builder vulnerabilityScanPolicy(String vulnerabilityScanPolicy) {
            return vulnerabilityScanPolicy(Output.of(vulnerabilityScanPolicy));
        }

        public TasksArgs build() {
            $.vulnerabilityScanPolicy = Objects.requireNonNull($.vulnerabilityScanPolicy, "expected parameter 'vulnerabilityScanPolicy' to be non-null");
            return $;
        }
    }

}
