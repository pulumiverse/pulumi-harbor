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
    /// 
    /// ### OIDC
    /// 
    /// ### LDAP
    /// </summary>
    [HarborResourceType("harbor:index/configAuth:ConfigAuth")]
    public partial class ConfigAuth : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Harbor authentication mode. Can be `"oidc_auth"`, `"db_auth"` or `"ldap_auth"`. (Default: `"db_auth"`)
        /// </summary>
        [Output("authMode")]
        public Output<string> AuthMode { get; private set; } = null!;

        [Output("ldapBaseDn")]
        public Output<string?> LdapBaseDn { get; private set; } = null!;

        [Output("ldapFilter")]
        public Output<string?> LdapFilter { get; private set; } = null!;

        [Output("ldapGroupAdminDn")]
        public Output<string?> LdapGroupAdminDn { get; private set; } = null!;

        [Output("ldapGroupBaseDn")]
        public Output<string?> LdapGroupBaseDn { get; private set; } = null!;

        [Output("ldapGroupFilter")]
        public Output<string?> LdapGroupFilter { get; private set; } = null!;

        [Output("ldapGroupGid")]
        public Output<string?> LdapGroupGid { get; private set; } = null!;

        [Output("ldapGroupMembership")]
        public Output<string?> LdapGroupMembership { get; private set; } = null!;

        [Output("ldapGroupScope")]
        public Output<string?> LdapGroupScope { get; private set; } = null!;

        [Output("ldapGroupUid")]
        public Output<string?> LdapGroupUid { get; private set; } = null!;

        [Output("ldapScope")]
        public Output<string?> LdapScope { get; private set; } = null!;

        [Output("ldapSearchDn")]
        public Output<string?> LdapSearchDn { get; private set; } = null!;

        [Output("ldapSearchPassword")]
        public Output<string?> LdapSearchPassword { get; private set; } = null!;

        [Output("ldapUid")]
        public Output<string?> LdapUid { get; private set; } = null!;

        [Output("ldapUrl")]
        public Output<string?> LdapUrl { get; private set; } = null!;

        [Output("ldapVerifyCert")]
        public Output<bool?> LdapVerifyCert { get; private set; } = null!;

        [Output("oidcAdminGroup")]
        public Output<string?> OidcAdminGroup { get; private set; } = null!;

        [Output("oidcAutoOnboard")]
        public Output<bool?> OidcAutoOnboard { get; private set; } = null!;

        [Output("oidcClientId")]
        public Output<string?> OidcClientId { get; private set; } = null!;

        [Output("oidcClientSecret")]
        public Output<string?> OidcClientSecret { get; private set; } = null!;

        [Output("oidcEndpoint")]
        public Output<string?> OidcEndpoint { get; private set; } = null!;

        [Output("oidcGroupFilter")]
        public Output<string?> OidcGroupFilter { get; private set; } = null!;

        [Output("oidcGroupsClaim")]
        public Output<string?> OidcGroupsClaim { get; private set; } = null!;

        [Output("oidcName")]
        public Output<string?> OidcName { get; private set; } = null!;

        [Output("oidcScope")]
        public Output<string?> OidcScope { get; private set; } = null!;

        [Output("oidcUserClaim")]
        public Output<string?> OidcUserClaim { get; private set; } = null!;

        [Output("oidcVerifyCert")]
        public Output<bool?> OidcVerifyCert { get; private set; } = null!;

        [Output("primaryAuthMode")]
        public Output<bool?> PrimaryAuthMode { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigAuth resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigAuth(string name, ConfigAuthArgs args, CustomResourceOptions? options = null)
            : base("harbor:index/configAuth:ConfigAuth", name, args ?? new ConfigAuthArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConfigAuth(string name, Input<string> id, ConfigAuthState? state = null, CustomResourceOptions? options = null)
            : base("harbor:index/configAuth:ConfigAuth", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-harbor",
                AdditionalSecretOutputs =
                {
                    "ldapSearchPassword",
                    "oidcClientSecret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ConfigAuth resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigAuth Get(string name, Input<string> id, ConfigAuthState? state = null, CustomResourceOptions? options = null)
        {
            return new ConfigAuth(name, id, state, options);
        }
    }

    public sealed class ConfigAuthArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Harbor authentication mode. Can be `"oidc_auth"`, `"db_auth"` or `"ldap_auth"`. (Default: `"db_auth"`)
        /// </summary>
        [Input("authMode", required: true)]
        public Input<string> AuthMode { get; set; } = null!;

        [Input("ldapBaseDn")]
        public Input<string>? LdapBaseDn { get; set; }

        [Input("ldapFilter")]
        public Input<string>? LdapFilter { get; set; }

        [Input("ldapGroupAdminDn")]
        public Input<string>? LdapGroupAdminDn { get; set; }

        [Input("ldapGroupBaseDn")]
        public Input<string>? LdapGroupBaseDn { get; set; }

        [Input("ldapGroupFilter")]
        public Input<string>? LdapGroupFilter { get; set; }

        [Input("ldapGroupGid")]
        public Input<string>? LdapGroupGid { get; set; }

        [Input("ldapGroupMembership")]
        public Input<string>? LdapGroupMembership { get; set; }

        [Input("ldapGroupScope")]
        public Input<string>? LdapGroupScope { get; set; }

        [Input("ldapGroupUid")]
        public Input<string>? LdapGroupUid { get; set; }

        [Input("ldapScope")]
        public Input<string>? LdapScope { get; set; }

        [Input("ldapSearchDn")]
        public Input<string>? LdapSearchDn { get; set; }

        [Input("ldapSearchPassword")]
        private Input<string>? _ldapSearchPassword;
        public Input<string>? LdapSearchPassword
        {
            get => _ldapSearchPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _ldapSearchPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("ldapUid")]
        public Input<string>? LdapUid { get; set; }

        [Input("ldapUrl")]
        public Input<string>? LdapUrl { get; set; }

        [Input("ldapVerifyCert")]
        public Input<bool>? LdapVerifyCert { get; set; }

        [Input("oidcAdminGroup")]
        public Input<string>? OidcAdminGroup { get; set; }

        [Input("oidcAutoOnboard")]
        public Input<bool>? OidcAutoOnboard { get; set; }

        [Input("oidcClientId")]
        public Input<string>? OidcClientId { get; set; }

        [Input("oidcClientSecret")]
        private Input<string>? _oidcClientSecret;
        public Input<string>? OidcClientSecret
        {
            get => _oidcClientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _oidcClientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("oidcEndpoint")]
        public Input<string>? OidcEndpoint { get; set; }

        [Input("oidcGroupFilter")]
        public Input<string>? OidcGroupFilter { get; set; }

        [Input("oidcGroupsClaim")]
        public Input<string>? OidcGroupsClaim { get; set; }

        [Input("oidcName")]
        public Input<string>? OidcName { get; set; }

        [Input("oidcScope")]
        public Input<string>? OidcScope { get; set; }

        [Input("oidcUserClaim")]
        public Input<string>? OidcUserClaim { get; set; }

        [Input("oidcVerifyCert")]
        public Input<bool>? OidcVerifyCert { get; set; }

        [Input("primaryAuthMode")]
        public Input<bool>? PrimaryAuthMode { get; set; }

        public ConfigAuthArgs()
        {
        }
        public static new ConfigAuthArgs Empty => new ConfigAuthArgs();
    }

    public sealed class ConfigAuthState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Harbor authentication mode. Can be `"oidc_auth"`, `"db_auth"` or `"ldap_auth"`. (Default: `"db_auth"`)
        /// </summary>
        [Input("authMode")]
        public Input<string>? AuthMode { get; set; }

        [Input("ldapBaseDn")]
        public Input<string>? LdapBaseDn { get; set; }

        [Input("ldapFilter")]
        public Input<string>? LdapFilter { get; set; }

        [Input("ldapGroupAdminDn")]
        public Input<string>? LdapGroupAdminDn { get; set; }

        [Input("ldapGroupBaseDn")]
        public Input<string>? LdapGroupBaseDn { get; set; }

        [Input("ldapGroupFilter")]
        public Input<string>? LdapGroupFilter { get; set; }

        [Input("ldapGroupGid")]
        public Input<string>? LdapGroupGid { get; set; }

        [Input("ldapGroupMembership")]
        public Input<string>? LdapGroupMembership { get; set; }

        [Input("ldapGroupScope")]
        public Input<string>? LdapGroupScope { get; set; }

        [Input("ldapGroupUid")]
        public Input<string>? LdapGroupUid { get; set; }

        [Input("ldapScope")]
        public Input<string>? LdapScope { get; set; }

        [Input("ldapSearchDn")]
        public Input<string>? LdapSearchDn { get; set; }

        [Input("ldapSearchPassword")]
        private Input<string>? _ldapSearchPassword;
        public Input<string>? LdapSearchPassword
        {
            get => _ldapSearchPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _ldapSearchPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("ldapUid")]
        public Input<string>? LdapUid { get; set; }

        [Input("ldapUrl")]
        public Input<string>? LdapUrl { get; set; }

        [Input("ldapVerifyCert")]
        public Input<bool>? LdapVerifyCert { get; set; }

        [Input("oidcAdminGroup")]
        public Input<string>? OidcAdminGroup { get; set; }

        [Input("oidcAutoOnboard")]
        public Input<bool>? OidcAutoOnboard { get; set; }

        [Input("oidcClientId")]
        public Input<string>? OidcClientId { get; set; }

        [Input("oidcClientSecret")]
        private Input<string>? _oidcClientSecret;
        public Input<string>? OidcClientSecret
        {
            get => _oidcClientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _oidcClientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("oidcEndpoint")]
        public Input<string>? OidcEndpoint { get; set; }

        [Input("oidcGroupFilter")]
        public Input<string>? OidcGroupFilter { get; set; }

        [Input("oidcGroupsClaim")]
        public Input<string>? OidcGroupsClaim { get; set; }

        [Input("oidcName")]
        public Input<string>? OidcName { get; set; }

        [Input("oidcScope")]
        public Input<string>? OidcScope { get; set; }

        [Input("oidcUserClaim")]
        public Input<string>? OidcUserClaim { get; set; }

        [Input("oidcVerifyCert")]
        public Input<bool>? OidcVerifyCert { get; set; }

        [Input("primaryAuthMode")]
        public Input<bool>? PrimaryAuthMode { get; set; }

        public ConfigAuthState()
        {
        }
        public static new ConfigAuthState Empty => new ConfigAuthState();
    }
}
