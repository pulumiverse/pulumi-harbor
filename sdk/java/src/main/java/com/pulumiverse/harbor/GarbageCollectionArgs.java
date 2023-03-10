// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.harbor;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GarbageCollectionArgs extends com.pulumi.resources.ResourceArgs {

    public static final GarbageCollectionArgs Empty = new GarbageCollectionArgs();

    @Import(name="deleteUntagged")
    private @Nullable Output<Boolean> deleteUntagged;

    public Optional<Output<Boolean>> deleteUntagged() {
        return Optional.ofNullable(this.deleteUntagged);
    }

    @Import(name="schedule", required=true)
    private Output<String> schedule;

    public Output<String> schedule() {
        return this.schedule;
    }

    private GarbageCollectionArgs() {}

    private GarbageCollectionArgs(GarbageCollectionArgs $) {
        this.deleteUntagged = $.deleteUntagged;
        this.schedule = $.schedule;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GarbageCollectionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GarbageCollectionArgs $;

        public Builder() {
            $ = new GarbageCollectionArgs();
        }

        public Builder(GarbageCollectionArgs defaults) {
            $ = new GarbageCollectionArgs(Objects.requireNonNull(defaults));
        }

        public Builder deleteUntagged(@Nullable Output<Boolean> deleteUntagged) {
            $.deleteUntagged = deleteUntagged;
            return this;
        }

        public Builder deleteUntagged(Boolean deleteUntagged) {
            return deleteUntagged(Output.of(deleteUntagged));
        }

        public Builder schedule(Output<String> schedule) {
            $.schedule = schedule;
            return this;
        }

        public Builder schedule(String schedule) {
            return schedule(Output.of(schedule));
        }

        public GarbageCollectionArgs build() {
            $.schedule = Objects.requireNonNull($.schedule, "expected parameter 'schedule' to be non-null");
            return $;
        }
    }

}
