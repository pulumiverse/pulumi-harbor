// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Harbor
{
    /// <summary>
    /// ## Example Usage
    /// </summary>
    [HarborResourceType("harbor:index/garbageCollection:GarbageCollection")]
    public partial class GarbageCollection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Allow garbage collection on untagged artifacts.
        /// </summary>
        [Output("deleteUntagged")]
        public Output<bool?> DeleteUntagged { get; private set; } = null!;

        /// <summary>
        /// Sets the schedule how often the Garbage Collection will run.  Can be to `"hourly"`, `"daily"`, `"weekly"` or can be a custom cron string ie, `"0 5 4 * * *"`
        /// </summary>
        [Output("schedule")]
        public Output<string> Schedule { get; private set; } = null!;

        /// <summary>
        /// Number of workers to run the garbage collection, value must be between 1 and 5.
        /// </summary>
        [Output("workers")]
        public Output<int?> Workers { get; private set; } = null!;


        /// <summary>
        /// Create a GarbageCollection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GarbageCollection(string name, GarbageCollectionArgs args, CustomResourceOptions? options = null)
            : base("harbor:index/garbageCollection:GarbageCollection", name, args ?? new GarbageCollectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GarbageCollection(string name, Input<string> id, GarbageCollectionState? state = null, CustomResourceOptions? options = null)
            : base("harbor:index/garbageCollection:GarbageCollection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GarbageCollection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GarbageCollection Get(string name, Input<string> id, GarbageCollectionState? state = null, CustomResourceOptions? options = null)
        {
            return new GarbageCollection(name, id, state, options);
        }
    }

    public sealed class GarbageCollectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allow garbage collection on untagged artifacts.
        /// </summary>
        [Input("deleteUntagged")]
        public Input<bool>? DeleteUntagged { get; set; }

        /// <summary>
        /// Sets the schedule how often the Garbage Collection will run.  Can be to `"hourly"`, `"daily"`, `"weekly"` or can be a custom cron string ie, `"0 5 4 * * *"`
        /// </summary>
        [Input("schedule", required: true)]
        public Input<string> Schedule { get; set; } = null!;

        /// <summary>
        /// Number of workers to run the garbage collection, value must be between 1 and 5.
        /// </summary>
        [Input("workers")]
        public Input<int>? Workers { get; set; }

        public GarbageCollectionArgs()
        {
        }
        public static new GarbageCollectionArgs Empty => new GarbageCollectionArgs();
    }

    public sealed class GarbageCollectionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allow garbage collection on untagged artifacts.
        /// </summary>
        [Input("deleteUntagged")]
        public Input<bool>? DeleteUntagged { get; set; }

        /// <summary>
        /// Sets the schedule how often the Garbage Collection will run.  Can be to `"hourly"`, `"daily"`, `"weekly"` or can be a custom cron string ie, `"0 5 4 * * *"`
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        /// <summary>
        /// Number of workers to run the garbage collection, value must be between 1 and 5.
        /// </summary>
        [Input("workers")]
        public Input<int>? Workers { get; set; }

        public GarbageCollectionState()
        {
        }
        public static new GarbageCollectionState Empty => new GarbageCollectionState();
    }
}
