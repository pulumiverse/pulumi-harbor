// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.harbor.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetProjectResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String name;
    private Integer projectId;
    private Boolean public_;
    private Boolean vulnerabilityScanning;

    private GetProjectResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String name() {
        return this.name;
    }
    public Integer projectId() {
        return this.projectId;
    }
    public Boolean public_() {
        return this.public_;
    }
    public Boolean vulnerabilityScanning() {
        return this.vulnerabilityScanning;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String name;
        private Integer projectId;
        private Boolean public_;
        private Boolean vulnerabilityScanning;
        public Builder() {}
        public Builder(GetProjectResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.projectId = defaults.projectId;
    	      this.public_ = defaults.public_;
    	      this.vulnerabilityScanning = defaults.vulnerabilityScanning;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(Integer projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter("public")
        public Builder public_(Boolean public_) {
            this.public_ = Objects.requireNonNull(public_);
            return this;
        }
        @CustomType.Setter
        public Builder vulnerabilityScanning(Boolean vulnerabilityScanning) {
            this.vulnerabilityScanning = Objects.requireNonNull(vulnerabilityScanning);
            return this;
        }
        public GetProjectResult build() {
            final var o = new GetProjectResult();
            o.id = id;
            o.name = name;
            o.projectId = projectId;
            o.public_ = public_;
            o.vulnerabilityScanning = vulnerabilityScanning;
            return o;
        }
    }
}
