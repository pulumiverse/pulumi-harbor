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
// ## Import
//
// The list can be imported using the `id` eg, `<break><break>```sh<break> $ pulumi import harbor:index/configSecurity:ConfigSecurity main "7" <break>```<break><break>` > Note that at this point of time Harbor doesn't has any api endpoint for deleting this list. Only updating the records.
type ConfigSecurity struct {
	pulumi.CustomResourceState

	// Time of creation of the list.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// System allowlist. Vulnerabilities in this list will be ignored when pushing and pulling images. Should be in the format or `["CVE-123", "CVE-145"]` or `["CVE-123"]`
	CveAllowlists pulumi.StringArrayOutput `pulumi:"cveAllowlists"`
	// The time for expiration of the allowlist, in the form of seconds since epoch. This is an optional attribute, if it's not set the CVE allowlist does not expire.
	ExpiresAt pulumi.IntPtrOutput `pulumi:"expiresAt"`
	// Time of update of the list.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewConfigSecurity registers a new resource with the given unique name, arguments, and options.
func NewConfigSecurity(ctx *pulumi.Context,
	name string, args *ConfigSecurityArgs, opts ...pulumi.ResourceOption) (*ConfigSecurity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CveAllowlists == nil {
		return nil, errors.New("invalid value for required argument 'CveAllowlists'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ConfigSecurity
	err := ctx.RegisterResource("harbor:index/configSecurity:ConfigSecurity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigSecurity gets an existing ConfigSecurity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigSecurity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigSecurityState, opts ...pulumi.ResourceOption) (*ConfigSecurity, error) {
	var resource ConfigSecurity
	err := ctx.ReadResource("harbor:index/configSecurity:ConfigSecurity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigSecurity resources.
type configSecurityState struct {
	// Time of creation of the list.
	CreationTime *string `pulumi:"creationTime"`
	// System allowlist. Vulnerabilities in this list will be ignored when pushing and pulling images. Should be in the format or `["CVE-123", "CVE-145"]` or `["CVE-123"]`
	CveAllowlists []string `pulumi:"cveAllowlists"`
	// The time for expiration of the allowlist, in the form of seconds since epoch. This is an optional attribute, if it's not set the CVE allowlist does not expire.
	ExpiresAt *int `pulumi:"expiresAt"`
	// Time of update of the list.
	UpdateTime *string `pulumi:"updateTime"`
}

type ConfigSecurityState struct {
	// Time of creation of the list.
	CreationTime pulumi.StringPtrInput
	// System allowlist. Vulnerabilities in this list will be ignored when pushing and pulling images. Should be in the format or `["CVE-123", "CVE-145"]` or `["CVE-123"]`
	CveAllowlists pulumi.StringArrayInput
	// The time for expiration of the allowlist, in the form of seconds since epoch. This is an optional attribute, if it's not set the CVE allowlist does not expire.
	ExpiresAt pulumi.IntPtrInput
	// Time of update of the list.
	UpdateTime pulumi.StringPtrInput
}

func (ConfigSecurityState) ElementType() reflect.Type {
	return reflect.TypeOf((*configSecurityState)(nil)).Elem()
}

type configSecurityArgs struct {
	// System allowlist. Vulnerabilities in this list will be ignored when pushing and pulling images. Should be in the format or `["CVE-123", "CVE-145"]` or `["CVE-123"]`
	CveAllowlists []string `pulumi:"cveAllowlists"`
	// The time for expiration of the allowlist, in the form of seconds since epoch. This is an optional attribute, if it's not set the CVE allowlist does not expire.
	ExpiresAt *int `pulumi:"expiresAt"`
}

// The set of arguments for constructing a ConfigSecurity resource.
type ConfigSecurityArgs struct {
	// System allowlist. Vulnerabilities in this list will be ignored when pushing and pulling images. Should be in the format or `["CVE-123", "CVE-145"]` or `["CVE-123"]`
	CveAllowlists pulumi.StringArrayInput
	// The time for expiration of the allowlist, in the form of seconds since epoch. This is an optional attribute, if it's not set the CVE allowlist does not expire.
	ExpiresAt pulumi.IntPtrInput
}

func (ConfigSecurityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configSecurityArgs)(nil)).Elem()
}

type ConfigSecurityInput interface {
	pulumi.Input

	ToConfigSecurityOutput() ConfigSecurityOutput
	ToConfigSecurityOutputWithContext(ctx context.Context) ConfigSecurityOutput
}

func (*ConfigSecurity) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigSecurity)(nil)).Elem()
}

func (i *ConfigSecurity) ToConfigSecurityOutput() ConfigSecurityOutput {
	return i.ToConfigSecurityOutputWithContext(context.Background())
}

func (i *ConfigSecurity) ToConfigSecurityOutputWithContext(ctx context.Context) ConfigSecurityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigSecurityOutput)
}

// ConfigSecurityArrayInput is an input type that accepts ConfigSecurityArray and ConfigSecurityArrayOutput values.
// You can construct a concrete instance of `ConfigSecurityArrayInput` via:
//
//	ConfigSecurityArray{ ConfigSecurityArgs{...} }
type ConfigSecurityArrayInput interface {
	pulumi.Input

	ToConfigSecurityArrayOutput() ConfigSecurityArrayOutput
	ToConfigSecurityArrayOutputWithContext(context.Context) ConfigSecurityArrayOutput
}

type ConfigSecurityArray []ConfigSecurityInput

func (ConfigSecurityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigSecurity)(nil)).Elem()
}

func (i ConfigSecurityArray) ToConfigSecurityArrayOutput() ConfigSecurityArrayOutput {
	return i.ToConfigSecurityArrayOutputWithContext(context.Background())
}

func (i ConfigSecurityArray) ToConfigSecurityArrayOutputWithContext(ctx context.Context) ConfigSecurityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigSecurityArrayOutput)
}

// ConfigSecurityMapInput is an input type that accepts ConfigSecurityMap and ConfigSecurityMapOutput values.
// You can construct a concrete instance of `ConfigSecurityMapInput` via:
//
//	ConfigSecurityMap{ "key": ConfigSecurityArgs{...} }
type ConfigSecurityMapInput interface {
	pulumi.Input

	ToConfigSecurityMapOutput() ConfigSecurityMapOutput
	ToConfigSecurityMapOutputWithContext(context.Context) ConfigSecurityMapOutput
}

type ConfigSecurityMap map[string]ConfigSecurityInput

func (ConfigSecurityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigSecurity)(nil)).Elem()
}

func (i ConfigSecurityMap) ToConfigSecurityMapOutput() ConfigSecurityMapOutput {
	return i.ToConfigSecurityMapOutputWithContext(context.Background())
}

func (i ConfigSecurityMap) ToConfigSecurityMapOutputWithContext(ctx context.Context) ConfigSecurityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigSecurityMapOutput)
}

type ConfigSecurityOutput struct{ *pulumi.OutputState }

func (ConfigSecurityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigSecurity)(nil)).Elem()
}

func (o ConfigSecurityOutput) ToConfigSecurityOutput() ConfigSecurityOutput {
	return o
}

func (o ConfigSecurityOutput) ToConfigSecurityOutputWithContext(ctx context.Context) ConfigSecurityOutput {
	return o
}

// Time of creation of the list.
func (o ConfigSecurityOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigSecurity) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// System allowlist. Vulnerabilities in this list will be ignored when pushing and pulling images. Should be in the format or `["CVE-123", "CVE-145"]` or `["CVE-123"]`
func (o ConfigSecurityOutput) CveAllowlists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ConfigSecurity) pulumi.StringArrayOutput { return v.CveAllowlists }).(pulumi.StringArrayOutput)
}

// The time for expiration of the allowlist, in the form of seconds since epoch. This is an optional attribute, if it's not set the CVE allowlist does not expire.
func (o ConfigSecurityOutput) ExpiresAt() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConfigSecurity) pulumi.IntPtrOutput { return v.ExpiresAt }).(pulumi.IntPtrOutput)
}

// Time of update of the list.
func (o ConfigSecurityOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigSecurity) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type ConfigSecurityArrayOutput struct{ *pulumi.OutputState }

func (ConfigSecurityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigSecurity)(nil)).Elem()
}

func (o ConfigSecurityArrayOutput) ToConfigSecurityArrayOutput() ConfigSecurityArrayOutput {
	return o
}

func (o ConfigSecurityArrayOutput) ToConfigSecurityArrayOutputWithContext(ctx context.Context) ConfigSecurityArrayOutput {
	return o
}

func (o ConfigSecurityArrayOutput) Index(i pulumi.IntInput) ConfigSecurityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConfigSecurity {
		return vs[0].([]*ConfigSecurity)[vs[1].(int)]
	}).(ConfigSecurityOutput)
}

type ConfigSecurityMapOutput struct{ *pulumi.OutputState }

func (ConfigSecurityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigSecurity)(nil)).Elem()
}

func (o ConfigSecurityMapOutput) ToConfigSecurityMapOutput() ConfigSecurityMapOutput {
	return o
}

func (o ConfigSecurityMapOutput) ToConfigSecurityMapOutputWithContext(ctx context.Context) ConfigSecurityMapOutput {
	return o
}

func (o ConfigSecurityMapOutput) MapIndex(k pulumi.StringInput) ConfigSecurityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConfigSecurity {
		return vs[0].(map[string]*ConfigSecurity)[vs[1].(string)]
	}).(ConfigSecurityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigSecurityInput)(nil)).Elem(), &ConfigSecurity{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigSecurityArrayInput)(nil)).Elem(), ConfigSecurityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigSecurityMapInput)(nil)).Elem(), ConfigSecurityMap{})
	pulumi.RegisterOutputType(ConfigSecurityOutput{})
	pulumi.RegisterOutputType(ConfigSecurityArrayOutput{})
	pulumi.RegisterOutputType(ConfigSecurityMapOutput{})
}
