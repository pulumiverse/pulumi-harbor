// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.harbor;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InterrogationServicesArgs extends com.pulumi.resources.ResourceArgs {

    public static final InterrogationServicesArgs Empty = new InterrogationServicesArgs();

    /**
     * Sets the default interrogation service **Clair**
     * 
     */
    @Import(name="defaultScanner")
    private @Nullable Output<String> defaultScanner;

    /**
     * @return Sets the default interrogation service **Clair**
     * 
     */
    public Optional<Output<String>> defaultScanner() {
        return Optional.ofNullable(this.defaultScanner);
    }

    /**
     * The frequency of the vulnerability scanning is done. This can be `daily`, `weekly`, `monthly` or can be a custom cron string.
     * 
     */
    @Import(name="vulnerabilityScanPolicy", required=true)
    private Output<String> vulnerabilityScanPolicy;

    /**
     * @return The frequency of the vulnerability scanning is done. This can be `daily`, `weekly`, `monthly` or can be a custom cron string.
     * 
     */
    public Output<String> vulnerabilityScanPolicy() {
        return this.vulnerabilityScanPolicy;
    }

    private InterrogationServicesArgs() {}

    private InterrogationServicesArgs(InterrogationServicesArgs $) {
        this.defaultScanner = $.defaultScanner;
        this.vulnerabilityScanPolicy = $.vulnerabilityScanPolicy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InterrogationServicesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InterrogationServicesArgs $;

        public Builder() {
            $ = new InterrogationServicesArgs();
        }

        public Builder(InterrogationServicesArgs defaults) {
            $ = new InterrogationServicesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param defaultScanner Sets the default interrogation service **Clair**
         * 
         * @return builder
         * 
         */
        public Builder defaultScanner(@Nullable Output<String> defaultScanner) {
            $.defaultScanner = defaultScanner;
            return this;
        }

        /**
         * @param defaultScanner Sets the default interrogation service **Clair**
         * 
         * @return builder
         * 
         */
        public Builder defaultScanner(String defaultScanner) {
            return defaultScanner(Output.of(defaultScanner));
        }

        /**
         * @param vulnerabilityScanPolicy The frequency of the vulnerability scanning is done. This can be `daily`, `weekly`, `monthly` or can be a custom cron string.
         * 
         * @return builder
         * 
         */
        public Builder vulnerabilityScanPolicy(Output<String> vulnerabilityScanPolicy) {
            $.vulnerabilityScanPolicy = vulnerabilityScanPolicy;
            return this;
        }

        /**
         * @param vulnerabilityScanPolicy The frequency of the vulnerability scanning is done. This can be `daily`, `weekly`, `monthly` or can be a custom cron string.
         * 
         * @return builder
         * 
         */
        public Builder vulnerabilityScanPolicy(String vulnerabilityScanPolicy) {
            return vulnerabilityScanPolicy(Output.of(vulnerabilityScanPolicy));
        }

        public InterrogationServicesArgs build() {
            $.vulnerabilityScanPolicy = Objects.requireNonNull($.vulnerabilityScanPolicy, "expected parameter 'vulnerabilityScanPolicy' to be non-null");
            return $;
        }
    }

}
