// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package harbor

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-harbor/sdk/v3/go/harbor/internal"
)

// ## Example Usage
//
// ### OIDC
//
// ### LDAP
type ConfigAuth struct {
	pulumi.CustomResourceState

	// Harbor authentication mode. Can be `"oidcAuth"`, `"dbAuth"` or `"ldapAuth"`. (Default: `"dbAuth"`)
	AuthMode            pulumi.StringOutput    `pulumi:"authMode"`
	LdapBaseDn          pulumi.StringPtrOutput `pulumi:"ldapBaseDn"`
	LdapFilter          pulumi.StringPtrOutput `pulumi:"ldapFilter"`
	LdapGroupAdminDn    pulumi.StringPtrOutput `pulumi:"ldapGroupAdminDn"`
	LdapGroupBaseDn     pulumi.StringPtrOutput `pulumi:"ldapGroupBaseDn"`
	LdapGroupFilter     pulumi.StringPtrOutput `pulumi:"ldapGroupFilter"`
	LdapGroupGid        pulumi.StringPtrOutput `pulumi:"ldapGroupGid"`
	LdapGroupMembership pulumi.StringPtrOutput `pulumi:"ldapGroupMembership"`
	LdapGroupScope      pulumi.StringPtrOutput `pulumi:"ldapGroupScope"`
	LdapGroupUid        pulumi.StringPtrOutput `pulumi:"ldapGroupUid"`
	LdapScope           pulumi.StringPtrOutput `pulumi:"ldapScope"`
	LdapSearchDn        pulumi.StringPtrOutput `pulumi:"ldapSearchDn"`
	LdapSearchPassword  pulumi.StringPtrOutput `pulumi:"ldapSearchPassword"`
	LdapUid             pulumi.StringPtrOutput `pulumi:"ldapUid"`
	LdapUrl             pulumi.StringPtrOutput `pulumi:"ldapUrl"`
	LdapVerifyCert      pulumi.BoolPtrOutput   `pulumi:"ldapVerifyCert"`
	OidcAdminGroup      pulumi.StringPtrOutput `pulumi:"oidcAdminGroup"`
	OidcAutoOnboard     pulumi.BoolPtrOutput   `pulumi:"oidcAutoOnboard"`
	OidcClientId        pulumi.StringPtrOutput `pulumi:"oidcClientId"`
	OidcClientSecret    pulumi.StringPtrOutput `pulumi:"oidcClientSecret"`
	OidcEndpoint        pulumi.StringPtrOutput `pulumi:"oidcEndpoint"`
	OidcGroupFilter     pulumi.StringPtrOutput `pulumi:"oidcGroupFilter"`
	OidcGroupsClaim     pulumi.StringPtrOutput `pulumi:"oidcGroupsClaim"`
	OidcName            pulumi.StringPtrOutput `pulumi:"oidcName"`
	OidcScope           pulumi.StringPtrOutput `pulumi:"oidcScope"`
	OidcUserClaim       pulumi.StringPtrOutput `pulumi:"oidcUserClaim"`
	OidcVerifyCert      pulumi.BoolPtrOutput   `pulumi:"oidcVerifyCert"`
	PrimaryAuthMode     pulumi.BoolPtrOutput   `pulumi:"primaryAuthMode"`
}

// NewConfigAuth registers a new resource with the given unique name, arguments, and options.
func NewConfigAuth(ctx *pulumi.Context,
	name string, args *ConfigAuthArgs, opts ...pulumi.ResourceOption) (*ConfigAuth, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthMode == nil {
		return nil, errors.New("invalid value for required argument 'AuthMode'")
	}
	if args.LdapSearchPassword != nil {
		args.LdapSearchPassword = pulumi.ToSecret(args.LdapSearchPassword).(pulumi.StringPtrInput)
	}
	if args.OidcClientSecret != nil {
		args.OidcClientSecret = pulumi.ToSecret(args.OidcClientSecret).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"ldapSearchPassword",
		"oidcClientSecret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ConfigAuth
	err := ctx.RegisterResource("harbor:index/configAuth:ConfigAuth", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigAuth gets an existing ConfigAuth resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigAuth(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigAuthState, opts ...pulumi.ResourceOption) (*ConfigAuth, error) {
	var resource ConfigAuth
	err := ctx.ReadResource("harbor:index/configAuth:ConfigAuth", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigAuth resources.
type configAuthState struct {
	// Harbor authentication mode. Can be `"oidcAuth"`, `"dbAuth"` or `"ldapAuth"`. (Default: `"dbAuth"`)
	AuthMode            *string `pulumi:"authMode"`
	LdapBaseDn          *string `pulumi:"ldapBaseDn"`
	LdapFilter          *string `pulumi:"ldapFilter"`
	LdapGroupAdminDn    *string `pulumi:"ldapGroupAdminDn"`
	LdapGroupBaseDn     *string `pulumi:"ldapGroupBaseDn"`
	LdapGroupFilter     *string `pulumi:"ldapGroupFilter"`
	LdapGroupGid        *string `pulumi:"ldapGroupGid"`
	LdapGroupMembership *string `pulumi:"ldapGroupMembership"`
	LdapGroupScope      *string `pulumi:"ldapGroupScope"`
	LdapGroupUid        *string `pulumi:"ldapGroupUid"`
	LdapScope           *string `pulumi:"ldapScope"`
	LdapSearchDn        *string `pulumi:"ldapSearchDn"`
	LdapSearchPassword  *string `pulumi:"ldapSearchPassword"`
	LdapUid             *string `pulumi:"ldapUid"`
	LdapUrl             *string `pulumi:"ldapUrl"`
	LdapVerifyCert      *bool   `pulumi:"ldapVerifyCert"`
	OidcAdminGroup      *string `pulumi:"oidcAdminGroup"`
	OidcAutoOnboard     *bool   `pulumi:"oidcAutoOnboard"`
	OidcClientId        *string `pulumi:"oidcClientId"`
	OidcClientSecret    *string `pulumi:"oidcClientSecret"`
	OidcEndpoint        *string `pulumi:"oidcEndpoint"`
	OidcGroupFilter     *string `pulumi:"oidcGroupFilter"`
	OidcGroupsClaim     *string `pulumi:"oidcGroupsClaim"`
	OidcName            *string `pulumi:"oidcName"`
	OidcScope           *string `pulumi:"oidcScope"`
	OidcUserClaim       *string `pulumi:"oidcUserClaim"`
	OidcVerifyCert      *bool   `pulumi:"oidcVerifyCert"`
	PrimaryAuthMode     *bool   `pulumi:"primaryAuthMode"`
}

type ConfigAuthState struct {
	// Harbor authentication mode. Can be `"oidcAuth"`, `"dbAuth"` or `"ldapAuth"`. (Default: `"dbAuth"`)
	AuthMode            pulumi.StringPtrInput
	LdapBaseDn          pulumi.StringPtrInput
	LdapFilter          pulumi.StringPtrInput
	LdapGroupAdminDn    pulumi.StringPtrInput
	LdapGroupBaseDn     pulumi.StringPtrInput
	LdapGroupFilter     pulumi.StringPtrInput
	LdapGroupGid        pulumi.StringPtrInput
	LdapGroupMembership pulumi.StringPtrInput
	LdapGroupScope      pulumi.StringPtrInput
	LdapGroupUid        pulumi.StringPtrInput
	LdapScope           pulumi.StringPtrInput
	LdapSearchDn        pulumi.StringPtrInput
	LdapSearchPassword  pulumi.StringPtrInput
	LdapUid             pulumi.StringPtrInput
	LdapUrl             pulumi.StringPtrInput
	LdapVerifyCert      pulumi.BoolPtrInput
	OidcAdminGroup      pulumi.StringPtrInput
	OidcAutoOnboard     pulumi.BoolPtrInput
	OidcClientId        pulumi.StringPtrInput
	OidcClientSecret    pulumi.StringPtrInput
	OidcEndpoint        pulumi.StringPtrInput
	OidcGroupFilter     pulumi.StringPtrInput
	OidcGroupsClaim     pulumi.StringPtrInput
	OidcName            pulumi.StringPtrInput
	OidcScope           pulumi.StringPtrInput
	OidcUserClaim       pulumi.StringPtrInput
	OidcVerifyCert      pulumi.BoolPtrInput
	PrimaryAuthMode     pulumi.BoolPtrInput
}

func (ConfigAuthState) ElementType() reflect.Type {
	return reflect.TypeOf((*configAuthState)(nil)).Elem()
}

type configAuthArgs struct {
	// Harbor authentication mode. Can be `"oidcAuth"`, `"dbAuth"` or `"ldapAuth"`. (Default: `"dbAuth"`)
	AuthMode            string  `pulumi:"authMode"`
	LdapBaseDn          *string `pulumi:"ldapBaseDn"`
	LdapFilter          *string `pulumi:"ldapFilter"`
	LdapGroupAdminDn    *string `pulumi:"ldapGroupAdminDn"`
	LdapGroupBaseDn     *string `pulumi:"ldapGroupBaseDn"`
	LdapGroupFilter     *string `pulumi:"ldapGroupFilter"`
	LdapGroupGid        *string `pulumi:"ldapGroupGid"`
	LdapGroupMembership *string `pulumi:"ldapGroupMembership"`
	LdapGroupScope      *string `pulumi:"ldapGroupScope"`
	LdapGroupUid        *string `pulumi:"ldapGroupUid"`
	LdapScope           *string `pulumi:"ldapScope"`
	LdapSearchDn        *string `pulumi:"ldapSearchDn"`
	LdapSearchPassword  *string `pulumi:"ldapSearchPassword"`
	LdapUid             *string `pulumi:"ldapUid"`
	LdapUrl             *string `pulumi:"ldapUrl"`
	LdapVerifyCert      *bool   `pulumi:"ldapVerifyCert"`
	OidcAdminGroup      *string `pulumi:"oidcAdminGroup"`
	OidcAutoOnboard     *bool   `pulumi:"oidcAutoOnboard"`
	OidcClientId        *string `pulumi:"oidcClientId"`
	OidcClientSecret    *string `pulumi:"oidcClientSecret"`
	OidcEndpoint        *string `pulumi:"oidcEndpoint"`
	OidcGroupFilter     *string `pulumi:"oidcGroupFilter"`
	OidcGroupsClaim     *string `pulumi:"oidcGroupsClaim"`
	OidcName            *string `pulumi:"oidcName"`
	OidcScope           *string `pulumi:"oidcScope"`
	OidcUserClaim       *string `pulumi:"oidcUserClaim"`
	OidcVerifyCert      *bool   `pulumi:"oidcVerifyCert"`
	PrimaryAuthMode     *bool   `pulumi:"primaryAuthMode"`
}

// The set of arguments for constructing a ConfigAuth resource.
type ConfigAuthArgs struct {
	// Harbor authentication mode. Can be `"oidcAuth"`, `"dbAuth"` or `"ldapAuth"`. (Default: `"dbAuth"`)
	AuthMode            pulumi.StringInput
	LdapBaseDn          pulumi.StringPtrInput
	LdapFilter          pulumi.StringPtrInput
	LdapGroupAdminDn    pulumi.StringPtrInput
	LdapGroupBaseDn     pulumi.StringPtrInput
	LdapGroupFilter     pulumi.StringPtrInput
	LdapGroupGid        pulumi.StringPtrInput
	LdapGroupMembership pulumi.StringPtrInput
	LdapGroupScope      pulumi.StringPtrInput
	LdapGroupUid        pulumi.StringPtrInput
	LdapScope           pulumi.StringPtrInput
	LdapSearchDn        pulumi.StringPtrInput
	LdapSearchPassword  pulumi.StringPtrInput
	LdapUid             pulumi.StringPtrInput
	LdapUrl             pulumi.StringPtrInput
	LdapVerifyCert      pulumi.BoolPtrInput
	OidcAdminGroup      pulumi.StringPtrInput
	OidcAutoOnboard     pulumi.BoolPtrInput
	OidcClientId        pulumi.StringPtrInput
	OidcClientSecret    pulumi.StringPtrInput
	OidcEndpoint        pulumi.StringPtrInput
	OidcGroupFilter     pulumi.StringPtrInput
	OidcGroupsClaim     pulumi.StringPtrInput
	OidcName            pulumi.StringPtrInput
	OidcScope           pulumi.StringPtrInput
	OidcUserClaim       pulumi.StringPtrInput
	OidcVerifyCert      pulumi.BoolPtrInput
	PrimaryAuthMode     pulumi.BoolPtrInput
}

func (ConfigAuthArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configAuthArgs)(nil)).Elem()
}

type ConfigAuthInput interface {
	pulumi.Input

	ToConfigAuthOutput() ConfigAuthOutput
	ToConfigAuthOutputWithContext(ctx context.Context) ConfigAuthOutput
}

func (*ConfigAuth) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigAuth)(nil)).Elem()
}

func (i *ConfigAuth) ToConfigAuthOutput() ConfigAuthOutput {
	return i.ToConfigAuthOutputWithContext(context.Background())
}

func (i *ConfigAuth) ToConfigAuthOutputWithContext(ctx context.Context) ConfigAuthOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigAuthOutput)
}

// ConfigAuthArrayInput is an input type that accepts ConfigAuthArray and ConfigAuthArrayOutput values.
// You can construct a concrete instance of `ConfigAuthArrayInput` via:
//
//	ConfigAuthArray{ ConfigAuthArgs{...} }
type ConfigAuthArrayInput interface {
	pulumi.Input

	ToConfigAuthArrayOutput() ConfigAuthArrayOutput
	ToConfigAuthArrayOutputWithContext(context.Context) ConfigAuthArrayOutput
}

type ConfigAuthArray []ConfigAuthInput

func (ConfigAuthArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigAuth)(nil)).Elem()
}

func (i ConfigAuthArray) ToConfigAuthArrayOutput() ConfigAuthArrayOutput {
	return i.ToConfigAuthArrayOutputWithContext(context.Background())
}

func (i ConfigAuthArray) ToConfigAuthArrayOutputWithContext(ctx context.Context) ConfigAuthArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigAuthArrayOutput)
}

// ConfigAuthMapInput is an input type that accepts ConfigAuthMap and ConfigAuthMapOutput values.
// You can construct a concrete instance of `ConfigAuthMapInput` via:
//
//	ConfigAuthMap{ "key": ConfigAuthArgs{...} }
type ConfigAuthMapInput interface {
	pulumi.Input

	ToConfigAuthMapOutput() ConfigAuthMapOutput
	ToConfigAuthMapOutputWithContext(context.Context) ConfigAuthMapOutput
}

type ConfigAuthMap map[string]ConfigAuthInput

func (ConfigAuthMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigAuth)(nil)).Elem()
}

func (i ConfigAuthMap) ToConfigAuthMapOutput() ConfigAuthMapOutput {
	return i.ToConfigAuthMapOutputWithContext(context.Background())
}

func (i ConfigAuthMap) ToConfigAuthMapOutputWithContext(ctx context.Context) ConfigAuthMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigAuthMapOutput)
}

type ConfigAuthOutput struct{ *pulumi.OutputState }

func (ConfigAuthOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigAuth)(nil)).Elem()
}

func (o ConfigAuthOutput) ToConfigAuthOutput() ConfigAuthOutput {
	return o
}

func (o ConfigAuthOutput) ToConfigAuthOutputWithContext(ctx context.Context) ConfigAuthOutput {
	return o
}

// Harbor authentication mode. Can be `"oidcAuth"`, `"dbAuth"` or `"ldapAuth"`. (Default: `"dbAuth"`)
func (o ConfigAuthOutput) AuthMode() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.StringOutput { return v.AuthMode }).(pulumi.StringOutput)
}

func (o ConfigAuthOutput) LdapBaseDn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.StringPtrOutput { return v.LdapBaseDn }).(pulumi.StringPtrOutput)
}

func (o ConfigAuthOutput) LdapFilter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.StringPtrOutput { return v.LdapFilter }).(pulumi.StringPtrOutput)
}

func (o ConfigAuthOutput) LdapGroupAdminDn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.StringPtrOutput { return v.LdapGroupAdminDn }).(pulumi.StringPtrOutput)
}

func (o ConfigAuthOutput) LdapGroupBaseDn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.StringPtrOutput { return v.LdapGroupBaseDn }).(pulumi.StringPtrOutput)
}

func (o ConfigAuthOutput) LdapGroupFilter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.StringPtrOutput { return v.LdapGroupFilter }).(pulumi.StringPtrOutput)
}

func (o ConfigAuthOutput) LdapGroupGid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.StringPtrOutput { return v.LdapGroupGid }).(pulumi.StringPtrOutput)
}

func (o ConfigAuthOutput) LdapGroupMembership() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.StringPtrOutput { return v.LdapGroupMembership }).(pulumi.StringPtrOutput)
}

func (o ConfigAuthOutput) LdapGroupScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.StringPtrOutput { return v.LdapGroupScope }).(pulumi.StringPtrOutput)
}

func (o ConfigAuthOutput) LdapGroupUid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.StringPtrOutput { return v.LdapGroupUid }).(pulumi.StringPtrOutput)
}

func (o ConfigAuthOutput) LdapScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.StringPtrOutput { return v.LdapScope }).(pulumi.StringPtrOutput)
}

func (o ConfigAuthOutput) LdapSearchDn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.StringPtrOutput { return v.LdapSearchDn }).(pulumi.StringPtrOutput)
}

func (o ConfigAuthOutput) LdapSearchPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.StringPtrOutput { return v.LdapSearchPassword }).(pulumi.StringPtrOutput)
}

func (o ConfigAuthOutput) LdapUid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.StringPtrOutput { return v.LdapUid }).(pulumi.StringPtrOutput)
}

func (o ConfigAuthOutput) LdapUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.StringPtrOutput { return v.LdapUrl }).(pulumi.StringPtrOutput)
}

func (o ConfigAuthOutput) LdapVerifyCert() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.BoolPtrOutput { return v.LdapVerifyCert }).(pulumi.BoolPtrOutput)
}

func (o ConfigAuthOutput) OidcAdminGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.StringPtrOutput { return v.OidcAdminGroup }).(pulumi.StringPtrOutput)
}

func (o ConfigAuthOutput) OidcAutoOnboard() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.BoolPtrOutput { return v.OidcAutoOnboard }).(pulumi.BoolPtrOutput)
}

func (o ConfigAuthOutput) OidcClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.StringPtrOutput { return v.OidcClientId }).(pulumi.StringPtrOutput)
}

func (o ConfigAuthOutput) OidcClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.StringPtrOutput { return v.OidcClientSecret }).(pulumi.StringPtrOutput)
}

func (o ConfigAuthOutput) OidcEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.StringPtrOutput { return v.OidcEndpoint }).(pulumi.StringPtrOutput)
}

func (o ConfigAuthOutput) OidcGroupFilter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.StringPtrOutput { return v.OidcGroupFilter }).(pulumi.StringPtrOutput)
}

func (o ConfigAuthOutput) OidcGroupsClaim() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.StringPtrOutput { return v.OidcGroupsClaim }).(pulumi.StringPtrOutput)
}

func (o ConfigAuthOutput) OidcName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.StringPtrOutput { return v.OidcName }).(pulumi.StringPtrOutput)
}

func (o ConfigAuthOutput) OidcScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.StringPtrOutput { return v.OidcScope }).(pulumi.StringPtrOutput)
}

func (o ConfigAuthOutput) OidcUserClaim() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.StringPtrOutput { return v.OidcUserClaim }).(pulumi.StringPtrOutput)
}

func (o ConfigAuthOutput) OidcVerifyCert() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.BoolPtrOutput { return v.OidcVerifyCert }).(pulumi.BoolPtrOutput)
}

func (o ConfigAuthOutput) PrimaryAuthMode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConfigAuth) pulumi.BoolPtrOutput { return v.PrimaryAuthMode }).(pulumi.BoolPtrOutput)
}

type ConfigAuthArrayOutput struct{ *pulumi.OutputState }

func (ConfigAuthArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigAuth)(nil)).Elem()
}

func (o ConfigAuthArrayOutput) ToConfigAuthArrayOutput() ConfigAuthArrayOutput {
	return o
}

func (o ConfigAuthArrayOutput) ToConfigAuthArrayOutputWithContext(ctx context.Context) ConfigAuthArrayOutput {
	return o
}

func (o ConfigAuthArrayOutput) Index(i pulumi.IntInput) ConfigAuthOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConfigAuth {
		return vs[0].([]*ConfigAuth)[vs[1].(int)]
	}).(ConfigAuthOutput)
}

type ConfigAuthMapOutput struct{ *pulumi.OutputState }

func (ConfigAuthMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigAuth)(nil)).Elem()
}

func (o ConfigAuthMapOutput) ToConfigAuthMapOutput() ConfigAuthMapOutput {
	return o
}

func (o ConfigAuthMapOutput) ToConfigAuthMapOutputWithContext(ctx context.Context) ConfigAuthMapOutput {
	return o
}

func (o ConfigAuthMapOutput) MapIndex(k pulumi.StringInput) ConfigAuthOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConfigAuth {
		return vs[0].(map[string]*ConfigAuth)[vs[1].(string)]
	}).(ConfigAuthOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigAuthInput)(nil)).Elem(), &ConfigAuth{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigAuthArrayInput)(nil)).Elem(), ConfigAuthArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigAuthMapInput)(nil)).Elem(), ConfigAuthMap{})
	pulumi.RegisterOutputType(ConfigAuthOutput{})
	pulumi.RegisterOutputType(ConfigAuthArrayOutput{})
	pulumi.RegisterOutputType(ConfigAuthMapOutput{})
}
