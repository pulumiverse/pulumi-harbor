// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.harbor.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetProjectPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetProjectPlainArgs Empty = new GetProjectPlainArgs();

    @Import(name="name", required=true)
    private String name;

    public String name() {
        return this.name;
    }

    private GetProjectPlainArgs() {}

    private GetProjectPlainArgs(GetProjectPlainArgs $) {
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetProjectPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetProjectPlainArgs $;

        public Builder() {
            $ = new GetProjectPlainArgs();
        }

        public Builder(GetProjectPlainArgs defaults) {
            $ = new GetProjectPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder name(String name) {
            $.name = name;
            return this;
        }

        public GetProjectPlainArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
