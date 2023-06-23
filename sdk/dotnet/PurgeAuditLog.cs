// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Harbor
{
    [HarborResourceType("harbor:index/purgeAuditLog:PurgeAuditLog")]
    public partial class PurgeAuditLog : global::Pulumi.CustomResource
    {
        [Output("auditRetentionHour")]
        public Output<int> AuditRetentionHour { get; private set; } = null!;

        [Output("includeOperations")]
        public Output<string> IncludeOperations { get; private set; } = null!;

        [Output("schedule")]
        public Output<string> Schedule { get; private set; } = null!;


        /// <summary>
        /// Create a PurgeAuditLog resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PurgeAuditLog(string name, PurgeAuditLogArgs args, CustomResourceOptions? options = null)
            : base("harbor:index/purgeAuditLog:PurgeAuditLog", name, args ?? new PurgeAuditLogArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PurgeAuditLog(string name, Input<string> id, PurgeAuditLogState? state = null, CustomResourceOptions? options = null)
            : base("harbor:index/purgeAuditLog:PurgeAuditLog", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-harbor",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PurgeAuditLog resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PurgeAuditLog Get(string name, Input<string> id, PurgeAuditLogState? state = null, CustomResourceOptions? options = null)
        {
            return new PurgeAuditLog(name, id, state, options);
        }
    }

    public sealed class PurgeAuditLogArgs : global::Pulumi.ResourceArgs
    {
        [Input("auditRetentionHour", required: true)]
        public Input<int> AuditRetentionHour { get; set; } = null!;

        [Input("includeOperations", required: true)]
        public Input<string> IncludeOperations { get; set; } = null!;

        [Input("schedule", required: true)]
        public Input<string> Schedule { get; set; } = null!;

        public PurgeAuditLogArgs()
        {
        }
        public static new PurgeAuditLogArgs Empty => new PurgeAuditLogArgs();
    }

    public sealed class PurgeAuditLogState : global::Pulumi.ResourceArgs
    {
        [Input("auditRetentionHour")]
        public Input<int>? AuditRetentionHour { get; set; }

        [Input("includeOperations")]
        public Input<string>? IncludeOperations { get; set; }

        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        public PurgeAuditLogState()
        {
        }
        public static new PurgeAuditLogState Empty => new PurgeAuditLogState();
    }
}
