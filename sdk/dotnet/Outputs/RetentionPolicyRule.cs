// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Harbor.Outputs
{

    [OutputType]
    public sealed class RetentionPolicyRule
    {
        /// <summary>
        /// retain always.
        /// </summary>
        public readonly bool? AlwaysRetain;
        /// <summary>
        /// Specify if the rule is disable or not. Defaults to `false`
        /// </summary>
        public readonly bool? Disabled;
        /// <summary>
        /// retain the most recently pulled n artifacts.
        /// </summary>
        public readonly int? MostRecentlyPulled;
        /// <summary>
        /// retain the most recently pushed n artifacts.
        /// </summary>
        public readonly int? MostRecentlyPushed;
        /// <summary>
        /// retains the artifacts pulled within the lasts n days.
        /// </summary>
        public readonly int? NDaysSinceLastPull;
        /// <summary>
        /// retains the artifacts pushed within the lasts n days.
        /// </summary>
        public readonly int? NDaysSinceLastPush;
        /// <summary>
        /// For the repositories excluding.
        /// </summary>
        public readonly string? RepoExcluding;
        /// <summary>
        /// For the repositories matching.
        /// </summary>
        public readonly string? RepoMatching;
        /// <summary>
        /// For the tag excluding.
        /// </summary>
        public readonly string? TagExcluding;
        /// <summary>
        /// For the tag matching.
        /// </summary>
        public readonly string? TagMatching;
        /// <summary>
        /// with untagged artifacts. Defaults to `true`
        /// </summary>
        public readonly bool? UntaggedArtifacts;

        [OutputConstructor]
        private RetentionPolicyRule(
            bool? alwaysRetain,

            bool? disabled,

            int? mostRecentlyPulled,

            int? mostRecentlyPushed,

            int? nDaysSinceLastPull,

            int? nDaysSinceLastPush,

            string? repoExcluding,

            string? repoMatching,

            string? tagExcluding,

            string? tagMatching,

            bool? untaggedArtifacts)
        {
            AlwaysRetain = alwaysRetain;
            Disabled = disabled;
            MostRecentlyPulled = mostRecentlyPulled;
            MostRecentlyPushed = mostRecentlyPushed;
            NDaysSinceLastPull = nDaysSinceLastPull;
            NDaysSinceLastPush = nDaysSinceLastPush;
            RepoExcluding = repoExcluding;
            RepoMatching = repoMatching;
            TagExcluding = tagExcluding;
            TagMatching = tagMatching;
            UntaggedArtifacts = untaggedArtifacts;
        }
    }
}
