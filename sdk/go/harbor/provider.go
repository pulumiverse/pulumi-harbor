// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package harbor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-harbor/sdk/v3/go/harbor/internal"
)

// The provider type for the harbor package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	BearerToken pulumi.StringPtrOutput `pulumi:"bearerToken"`
	Password    pulumi.StringPtrOutput `pulumi:"password"`
	RobotPrefix pulumi.StringPtrOutput `pulumi:"robotPrefix"`
	Url         pulumi.StringPtrOutput `pulumi:"url"`
	Username    pulumi.StringPtrOutput `pulumi:"username"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	if args.ApiVersion == nil {
		args.ApiVersion = pulumi.IntPtr(2)
	}
	if args.Insecure == nil {
		if d := internal.GetEnvOrDefault(true, internal.ParseEnvBool, "HARBOR_IGNORE_CERT"); d != nil {
			args.Insecure = pulumi.BoolPtr(d.(bool))
		}
	}
	if args.Password == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "HARBOR_PASSWORD"); d != nil {
			args.Password = pulumi.StringPtr(d.(string))
		}
	}
	if args.Url == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "HARBOR_URL"); d != nil {
			args.Url = pulumi.StringPtr(d.(string))
		}
	}
	if args.Username == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "HARBOR_USERNAME"); d != nil {
			args.Username = pulumi.StringPtr(d.(string))
		}
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:harbor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	ApiVersion  *int    `pulumi:"apiVersion"`
	BearerToken *string `pulumi:"bearerToken"`
	Insecure    *bool   `pulumi:"insecure"`
	Password    *string `pulumi:"password"`
	RobotPrefix *string `pulumi:"robotPrefix"`
	Url         *string `pulumi:"url"`
	Username    *string `pulumi:"username"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	ApiVersion  pulumi.IntPtrInput
	BearerToken pulumi.StringPtrInput
	Insecure    pulumi.BoolPtrInput
	Password    pulumi.StringPtrInput
	RobotPrefix pulumi.StringPtrInput
	Url         pulumi.StringPtrInput
	Username    pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

func (o ProviderOutput) BearerToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.BearerToken }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) RobotPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.RobotPrefix }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Url }).(pulumi.StringPtrOutput)
}

func (o ProviderOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
