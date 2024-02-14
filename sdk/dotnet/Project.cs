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
    /// ## Harbor project example as proxy cache
    /// 
    /// ## Import
    /// 
    /// Harbor project can be imported using the `project id` eg,&lt;break&gt;&lt;break&gt; ` &lt;break&gt;&lt;break&gt; ```sh&lt;break&gt; $ pulumi import harbor:index/project:Project main /projects/1 &lt;break&gt;```&lt;break&gt;&lt;break&gt;  `&lt;break&gt;&lt;break&gt;
    /// </summary>
    [HarborResourceType("harbor:index/project:Project")]
    public partial class Project : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Project allowlist allows vulnerabilities in this list to be ignored in this project when pushing and pulling images. Should be in the format or `["CVE-123", "CVE-145"]` or `["CVE-123"]`
        /// </summary>
        [Output("cveAllowlists")]
        public Output<ImmutableArray<string>> CveAllowlists { get; private set; } = null!;

        /// <summary>
        /// Prevent deployment of images with vulnerability severity equal or higher than the specified value. Images must be scanned before this takes effect. Possible values: `critical`, `high`, `medium`, `low`, `none`. (Default: `""` - empty)
        /// </summary>
        [Output("deploymentSecurity")]
        public Output<string?> DeploymentSecurity { get; private set; } = null!;

        /// <summary>
        /// Enables Content Trust for project. When enabled it queries the embedded docker notary server. Can be set to `"true"` or `"false"` (Default: false)
        /// </summary>
        [Output("enableContentTrust")]
        public Output<bool?> EnableContentTrust { get; private set; } = null!;

        /// <summary>
        /// Enables Content Trust Cosign for project. When enabled it queries Cosign. Can be set to `"true"` or `"false"` (Default: false)
        /// </summary>
        [Output("enableContentTrustCosign")]
        public Output<bool?> EnableContentTrustCosign { get; private set; } = null!;

        /// <summary>
        /// A boolean that indicates all repositories should be deleted from the project so that the project can be destroyed without error. These repositories are *not* recoverable.
        /// </summary>
        [Output("forceDestroy")]
        public Output<bool?> ForceDestroy { get; private set; } = null!;

        /// <summary>
        /// The name of the project that will be created in harbor.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The id of the project with harbor.
        /// </summary>
        [Output("projectId")]
        public Output<int> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The project will be public accessibility. Can be set to `"true"` or `"false"` (Default: false)
        /// </summary>
        [Output("public")]
        public Output<bool?> Public { get; private set; } = null!;

        /// <summary>
        /// To enable project as Proxy Cache
        /// </summary>
        [Output("registryId")]
        public Output<int> RegistryId { get; private set; } = null!;

        /// <summary>
        /// The storage quota of the project in GB's
        /// </summary>
        [Output("storageQuota")]
        public Output<int?> StorageQuota { get; private set; } = null!;

        /// <summary>
        /// Images will be scanned for vulnerabilities when push to harbor. Can be set to `"true"` or `"false"` (Default: true)
        /// </summary>
        [Output("vulnerabilityScanning")]
        public Output<bool?> VulnerabilityScanning { get; private set; } = null!;


        /// <summary>
        /// Create a Project resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Project(string name, ProjectArgs? args = null, CustomResourceOptions? options = null)
            : base("harbor:index/project:Project", name, args ?? new ProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Project(string name, Input<string> id, ProjectState? state = null, CustomResourceOptions? options = null)
            : base("harbor:index/project:Project", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Project resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Project Get(string name, Input<string> id, ProjectState? state = null, CustomResourceOptions? options = null)
        {
            return new Project(name, id, state, options);
        }
    }

    public sealed class ProjectArgs : global::Pulumi.ResourceArgs
    {
        [Input("cveAllowlists")]
        private InputList<string>? _cveAllowlists;

        /// <summary>
        /// Project allowlist allows vulnerabilities in this list to be ignored in this project when pushing and pulling images. Should be in the format or `["CVE-123", "CVE-145"]` or `["CVE-123"]`
        /// </summary>
        public InputList<string> CveAllowlists
        {
            get => _cveAllowlists ?? (_cveAllowlists = new InputList<string>());
            set => _cveAllowlists = value;
        }

        /// <summary>
        /// Prevent deployment of images with vulnerability severity equal or higher than the specified value. Images must be scanned before this takes effect. Possible values: `critical`, `high`, `medium`, `low`, `none`. (Default: `""` - empty)
        /// </summary>
        [Input("deploymentSecurity")]
        public Input<string>? DeploymentSecurity { get; set; }

        /// <summary>
        /// Enables Content Trust for project. When enabled it queries the embedded docker notary server. Can be set to `"true"` or `"false"` (Default: false)
        /// </summary>
        [Input("enableContentTrust")]
        public Input<bool>? EnableContentTrust { get; set; }

        /// <summary>
        /// Enables Content Trust Cosign for project. When enabled it queries Cosign. Can be set to `"true"` or `"false"` (Default: false)
        /// </summary>
        [Input("enableContentTrustCosign")]
        public Input<bool>? EnableContentTrustCosign { get; set; }

        /// <summary>
        /// A boolean that indicates all repositories should be deleted from the project so that the project can be destroyed without error. These repositories are *not* recoverable.
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        /// <summary>
        /// The name of the project that will be created in harbor.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project will be public accessibility. Can be set to `"true"` or `"false"` (Default: false)
        /// </summary>
        [Input("public")]
        public Input<bool>? Public { get; set; }

        /// <summary>
        /// To enable project as Proxy Cache
        /// </summary>
        [Input("registryId")]
        public Input<int>? RegistryId { get; set; }

        /// <summary>
        /// The storage quota of the project in GB's
        /// </summary>
        [Input("storageQuota")]
        public Input<int>? StorageQuota { get; set; }

        /// <summary>
        /// Images will be scanned for vulnerabilities when push to harbor. Can be set to `"true"` or `"false"` (Default: true)
        /// </summary>
        [Input("vulnerabilityScanning")]
        public Input<bool>? VulnerabilityScanning { get; set; }

        public ProjectArgs()
        {
        }
        public static new ProjectArgs Empty => new ProjectArgs();
    }

    public sealed class ProjectState : global::Pulumi.ResourceArgs
    {
        [Input("cveAllowlists")]
        private InputList<string>? _cveAllowlists;

        /// <summary>
        /// Project allowlist allows vulnerabilities in this list to be ignored in this project when pushing and pulling images. Should be in the format or `["CVE-123", "CVE-145"]` or `["CVE-123"]`
        /// </summary>
        public InputList<string> CveAllowlists
        {
            get => _cveAllowlists ?? (_cveAllowlists = new InputList<string>());
            set => _cveAllowlists = value;
        }

        /// <summary>
        /// Prevent deployment of images with vulnerability severity equal or higher than the specified value. Images must be scanned before this takes effect. Possible values: `critical`, `high`, `medium`, `low`, `none`. (Default: `""` - empty)
        /// </summary>
        [Input("deploymentSecurity")]
        public Input<string>? DeploymentSecurity { get; set; }

        /// <summary>
        /// Enables Content Trust for project. When enabled it queries the embedded docker notary server. Can be set to `"true"` or `"false"` (Default: false)
        /// </summary>
        [Input("enableContentTrust")]
        public Input<bool>? EnableContentTrust { get; set; }

        /// <summary>
        /// Enables Content Trust Cosign for project. When enabled it queries Cosign. Can be set to `"true"` or `"false"` (Default: false)
        /// </summary>
        [Input("enableContentTrustCosign")]
        public Input<bool>? EnableContentTrustCosign { get; set; }

        /// <summary>
        /// A boolean that indicates all repositories should be deleted from the project so that the project can be destroyed without error. These repositories are *not* recoverable.
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        /// <summary>
        /// The name of the project that will be created in harbor.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The id of the project with harbor.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// The project will be public accessibility. Can be set to `"true"` or `"false"` (Default: false)
        /// </summary>
        [Input("public")]
        public Input<bool>? Public { get; set; }

        /// <summary>
        /// To enable project as Proxy Cache
        /// </summary>
        [Input("registryId")]
        public Input<int>? RegistryId { get; set; }

        /// <summary>
        /// The storage quota of the project in GB's
        /// </summary>
        [Input("storageQuota")]
        public Input<int>? StorageQuota { get; set; }

        /// <summary>
        /// Images will be scanned for vulnerabilities when push to harbor. Can be set to `"true"` or `"false"` (Default: true)
        /// </summary>
        [Input("vulnerabilityScanning")]
        public Input<bool>? VulnerabilityScanning { get; set; }

        public ProjectState()
        {
        }
        public static new ProjectState Empty => new ProjectState();
    }
}
