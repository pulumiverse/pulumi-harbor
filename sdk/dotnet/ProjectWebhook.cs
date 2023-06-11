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
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Harbor = Pulumiverse.Harbor;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mainProject = new Harbor.Project("mainProject");
    /// 
    ///     var mainProjectWebhook = new Harbor.ProjectWebhook("mainProjectWebhook", new()
    ///     {
    ///         Address = "https://webhook.domain.com",
    ///         ProjectId = mainProject.Id,
    ///         NotifyType = "http",
    ///         EventsTypes = new[]
    ///         {
    ///             "DELETE_ARTIFACT",
    ///             "PULL_ARTIFACT",
    ///             "PUSH_ARTIFACT",
    ///             "QUOTA_EXCEED",
    ///             "QUOTA_WARNING",
    ///             "REPLICATION",
    ///             "SCANNING_FAILED",
    ///             "SCANNING_COMPLETED",
    ///             "TAG_RETENTION",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [HarborResourceType("harbor:index/projectWebhook:ProjectWebhook")]
    public partial class ProjectWebhook : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The address of the webhook
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// authentication header for you the webhook
        /// </summary>
        [Output("authHeader")]
        public Output<string?> AuthHeader { get; private set; } = null!;

        /// <summary>
        /// _ (Optional, string) A description of the webhook
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// , To enable / disable the webhook. Default `true`
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// ) The type events you want to subscript to can be `DELETE_ARTIFACT`, `PULL_ARTIFACT`, `PUSH_ARTIFACT`, `QUOTA_EXCEED`, `QUOTA_WARNING`, `REPLICATION`, `SCANNING_FAILED`, `SCANNING_COMPLETED`, `TAG_RETENTION`
        /// </summary>
        [Output("eventsTypes")]
        public Output<ImmutableArray<string>> EventsTypes { get; private set; } = null!;

        /// <summary>
        /// The name of the webhook that will be created in harbor.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The notification type either `http` or `slack`
        /// </summary>
        [Output("notifyType")]
        public Output<string> NotifyType { get; private set; } = null!;

        /// <summary>
        /// The project id of the harbor that webhook related to.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// checks the for validate SSL certificate.
        /// </summary>
        [Output("skipCertVerify")]
        public Output<bool?> SkipCertVerify { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectWebhook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectWebhook(string name, ProjectWebhookArgs args, CustomResourceOptions? options = null)
            : base("harbor:index/projectWebhook:ProjectWebhook", name, args ?? new ProjectWebhookArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectWebhook(string name, Input<string> id, ProjectWebhookState? state = null, CustomResourceOptions? options = null)
            : base("harbor:index/projectWebhook:ProjectWebhook", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectWebhook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectWebhook Get(string name, Input<string> id, ProjectWebhookState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectWebhook(name, id, state, options);
        }
    }

    public sealed class ProjectWebhookArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The address of the webhook
        /// </summary>
        [Input("address", required: true)]
        public Input<string> Address { get; set; } = null!;

        /// <summary>
        /// authentication header for you the webhook
        /// </summary>
        [Input("authHeader")]
        public Input<string>? AuthHeader { get; set; }

        /// <summary>
        /// _ (Optional, string) A description of the webhook
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// , To enable / disable the webhook. Default `true`
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("eventsTypes", required: true)]
        private InputList<string>? _eventsTypes;

        /// <summary>
        /// ) The type events you want to subscript to can be `DELETE_ARTIFACT`, `PULL_ARTIFACT`, `PUSH_ARTIFACT`, `QUOTA_EXCEED`, `QUOTA_WARNING`, `REPLICATION`, `SCANNING_FAILED`, `SCANNING_COMPLETED`, `TAG_RETENTION`
        /// </summary>
        public InputList<string> EventsTypes
        {
            get => _eventsTypes ?? (_eventsTypes = new InputList<string>());
            set => _eventsTypes = value;
        }

        /// <summary>
        /// The name of the webhook that will be created in harbor.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The notification type either `http` or `slack`
        /// </summary>
        [Input("notifyType", required: true)]
        public Input<string> NotifyType { get; set; } = null!;

        /// <summary>
        /// The project id of the harbor that webhook related to.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// checks the for validate SSL certificate.
        /// </summary>
        [Input("skipCertVerify")]
        public Input<bool>? SkipCertVerify { get; set; }

        public ProjectWebhookArgs()
        {
        }
        public static new ProjectWebhookArgs Empty => new ProjectWebhookArgs();
    }

    public sealed class ProjectWebhookState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The address of the webhook
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// authentication header for you the webhook
        /// </summary>
        [Input("authHeader")]
        public Input<string>? AuthHeader { get; set; }

        /// <summary>
        /// _ (Optional, string) A description of the webhook
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// , To enable / disable the webhook. Default `true`
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("eventsTypes")]
        private InputList<string>? _eventsTypes;

        /// <summary>
        /// ) The type events you want to subscript to can be `DELETE_ARTIFACT`, `PULL_ARTIFACT`, `PUSH_ARTIFACT`, `QUOTA_EXCEED`, `QUOTA_WARNING`, `REPLICATION`, `SCANNING_FAILED`, `SCANNING_COMPLETED`, `TAG_RETENTION`
        /// </summary>
        public InputList<string> EventsTypes
        {
            get => _eventsTypes ?? (_eventsTypes = new InputList<string>());
            set => _eventsTypes = value;
        }

        /// <summary>
        /// The name of the webhook that will be created in harbor.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The notification type either `http` or `slack`
        /// </summary>
        [Input("notifyType")]
        public Input<string>? NotifyType { get; set; }

        /// <summary>
        /// The project id of the harbor that webhook related to.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// checks the for validate SSL certificate.
        /// </summary>
        [Input("skipCertVerify")]
        public Input<bool>? SkipCertVerify { get; set; }

        public ProjectWebhookState()
        {
        }
        public static new ProjectWebhookState Empty => new ProjectWebhookState();
    }
}
