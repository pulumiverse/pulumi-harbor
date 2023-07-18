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
type ProjectWebhook struct {
	pulumi.CustomResourceState

	// The address of the webhook
	Address pulumi.StringOutput `pulumi:"address"`
	// authentication header for you the webhook
	AuthHeader pulumi.StringPtrOutput `pulumi:"authHeader"`
	// _ (Optional, string) A description of the webhook
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// , To enable / disable the webhook. Default `true`
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// ) The type events you want to subscript to can be `DELETE_ARTIFACT`, `PULL_ARTIFACT`, `PUSH_ARTIFACT`, `QUOTA_EXCEED`, `QUOTA_WARNING`, `REPLICATION`, `SCANNING_FAILED`, `SCANNING_COMPLETED`, `TAG_RETENTION`
	EventsTypes pulumi.StringArrayOutput `pulumi:"eventsTypes"`
	// The name of the webhook that will be created in harbor.
	Name pulumi.StringOutput `pulumi:"name"`
	// The notification type either `http` or `slack`
	NotifyType pulumi.StringOutput `pulumi:"notifyType"`
	// The project id (**/projects/ID**) of the harbor that webhook related to.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// checks the for validate SSL certificate.
	SkipCertVerify pulumi.BoolPtrOutput `pulumi:"skipCertVerify"`
}

// NewProjectWebhook registers a new resource with the given unique name, arguments, and options.
func NewProjectWebhook(ctx *pulumi.Context,
	name string, args *ProjectWebhookArgs, opts ...pulumi.ResourceOption) (*ProjectWebhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address == nil {
		return nil, errors.New("invalid value for required argument 'Address'")
	}
	if args.EventsTypes == nil {
		return nil, errors.New("invalid value for required argument 'EventsTypes'")
	}
	if args.NotifyType == nil {
		return nil, errors.New("invalid value for required argument 'NotifyType'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectWebhook
	err := ctx.RegisterResource("harbor:index/projectWebhook:ProjectWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectWebhook gets an existing ProjectWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectWebhookState, opts ...pulumi.ResourceOption) (*ProjectWebhook, error) {
	var resource ProjectWebhook
	err := ctx.ReadResource("harbor:index/projectWebhook:ProjectWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectWebhook resources.
type projectWebhookState struct {
	// The address of the webhook
	Address *string `pulumi:"address"`
	// authentication header for you the webhook
	AuthHeader *string `pulumi:"authHeader"`
	// _ (Optional, string) A description of the webhook
	Description *string `pulumi:"description"`
	// , To enable / disable the webhook. Default `true`
	Enabled *bool `pulumi:"enabled"`
	// ) The type events you want to subscript to can be `DELETE_ARTIFACT`, `PULL_ARTIFACT`, `PUSH_ARTIFACT`, `QUOTA_EXCEED`, `QUOTA_WARNING`, `REPLICATION`, `SCANNING_FAILED`, `SCANNING_COMPLETED`, `TAG_RETENTION`
	EventsTypes []string `pulumi:"eventsTypes"`
	// The name of the webhook that will be created in harbor.
	Name *string `pulumi:"name"`
	// The notification type either `http` or `slack`
	NotifyType *string `pulumi:"notifyType"`
	// The project id (**/projects/ID**) of the harbor that webhook related to.
	ProjectId *string `pulumi:"projectId"`
	// checks the for validate SSL certificate.
	SkipCertVerify *bool `pulumi:"skipCertVerify"`
}

type ProjectWebhookState struct {
	// The address of the webhook
	Address pulumi.StringPtrInput
	// authentication header for you the webhook
	AuthHeader pulumi.StringPtrInput
	// _ (Optional, string) A description of the webhook
	Description pulumi.StringPtrInput
	// , To enable / disable the webhook. Default `true`
	Enabled pulumi.BoolPtrInput
	// ) The type events you want to subscript to can be `DELETE_ARTIFACT`, `PULL_ARTIFACT`, `PUSH_ARTIFACT`, `QUOTA_EXCEED`, `QUOTA_WARNING`, `REPLICATION`, `SCANNING_FAILED`, `SCANNING_COMPLETED`, `TAG_RETENTION`
	EventsTypes pulumi.StringArrayInput
	// The name of the webhook that will be created in harbor.
	Name pulumi.StringPtrInput
	// The notification type either `http` or `slack`
	NotifyType pulumi.StringPtrInput
	// The project id (**/projects/ID**) of the harbor that webhook related to.
	ProjectId pulumi.StringPtrInput
	// checks the for validate SSL certificate.
	SkipCertVerify pulumi.BoolPtrInput
}

func (ProjectWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectWebhookState)(nil)).Elem()
}

type projectWebhookArgs struct {
	// The address of the webhook
	Address string `pulumi:"address"`
	// authentication header for you the webhook
	AuthHeader *string `pulumi:"authHeader"`
	// _ (Optional, string) A description of the webhook
	Description *string `pulumi:"description"`
	// , To enable / disable the webhook. Default `true`
	Enabled *bool `pulumi:"enabled"`
	// ) The type events you want to subscript to can be `DELETE_ARTIFACT`, `PULL_ARTIFACT`, `PUSH_ARTIFACT`, `QUOTA_EXCEED`, `QUOTA_WARNING`, `REPLICATION`, `SCANNING_FAILED`, `SCANNING_COMPLETED`, `TAG_RETENTION`
	EventsTypes []string `pulumi:"eventsTypes"`
	// The name of the webhook that will be created in harbor.
	Name *string `pulumi:"name"`
	// The notification type either `http` or `slack`
	NotifyType string `pulumi:"notifyType"`
	// The project id (**/projects/ID**) of the harbor that webhook related to.
	ProjectId string `pulumi:"projectId"`
	// checks the for validate SSL certificate.
	SkipCertVerify *bool `pulumi:"skipCertVerify"`
}

// The set of arguments for constructing a ProjectWebhook resource.
type ProjectWebhookArgs struct {
	// The address of the webhook
	Address pulumi.StringInput
	// authentication header for you the webhook
	AuthHeader pulumi.StringPtrInput
	// _ (Optional, string) A description of the webhook
	Description pulumi.StringPtrInput
	// , To enable / disable the webhook. Default `true`
	Enabled pulumi.BoolPtrInput
	// ) The type events you want to subscript to can be `DELETE_ARTIFACT`, `PULL_ARTIFACT`, `PUSH_ARTIFACT`, `QUOTA_EXCEED`, `QUOTA_WARNING`, `REPLICATION`, `SCANNING_FAILED`, `SCANNING_COMPLETED`, `TAG_RETENTION`
	EventsTypes pulumi.StringArrayInput
	// The name of the webhook that will be created in harbor.
	Name pulumi.StringPtrInput
	// The notification type either `http` or `slack`
	NotifyType pulumi.StringInput
	// The project id (**/projects/ID**) of the harbor that webhook related to.
	ProjectId pulumi.StringInput
	// checks the for validate SSL certificate.
	SkipCertVerify pulumi.BoolPtrInput
}

func (ProjectWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectWebhookArgs)(nil)).Elem()
}

type ProjectWebhookInput interface {
	pulumi.Input

	ToProjectWebhookOutput() ProjectWebhookOutput
	ToProjectWebhookOutputWithContext(ctx context.Context) ProjectWebhookOutput
}

func (*ProjectWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectWebhook)(nil)).Elem()
}

func (i *ProjectWebhook) ToProjectWebhookOutput() ProjectWebhookOutput {
	return i.ToProjectWebhookOutputWithContext(context.Background())
}

func (i *ProjectWebhook) ToProjectWebhookOutputWithContext(ctx context.Context) ProjectWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectWebhookOutput)
}

// ProjectWebhookArrayInput is an input type that accepts ProjectWebhookArray and ProjectWebhookArrayOutput values.
// You can construct a concrete instance of `ProjectWebhookArrayInput` via:
//
//	ProjectWebhookArray{ ProjectWebhookArgs{...} }
type ProjectWebhookArrayInput interface {
	pulumi.Input

	ToProjectWebhookArrayOutput() ProjectWebhookArrayOutput
	ToProjectWebhookArrayOutputWithContext(context.Context) ProjectWebhookArrayOutput
}

type ProjectWebhookArray []ProjectWebhookInput

func (ProjectWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectWebhook)(nil)).Elem()
}

func (i ProjectWebhookArray) ToProjectWebhookArrayOutput() ProjectWebhookArrayOutput {
	return i.ToProjectWebhookArrayOutputWithContext(context.Background())
}

func (i ProjectWebhookArray) ToProjectWebhookArrayOutputWithContext(ctx context.Context) ProjectWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectWebhookArrayOutput)
}

// ProjectWebhookMapInput is an input type that accepts ProjectWebhookMap and ProjectWebhookMapOutput values.
// You can construct a concrete instance of `ProjectWebhookMapInput` via:
//
//	ProjectWebhookMap{ "key": ProjectWebhookArgs{...} }
type ProjectWebhookMapInput interface {
	pulumi.Input

	ToProjectWebhookMapOutput() ProjectWebhookMapOutput
	ToProjectWebhookMapOutputWithContext(context.Context) ProjectWebhookMapOutput
}

type ProjectWebhookMap map[string]ProjectWebhookInput

func (ProjectWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectWebhook)(nil)).Elem()
}

func (i ProjectWebhookMap) ToProjectWebhookMapOutput() ProjectWebhookMapOutput {
	return i.ToProjectWebhookMapOutputWithContext(context.Background())
}

func (i ProjectWebhookMap) ToProjectWebhookMapOutputWithContext(ctx context.Context) ProjectWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectWebhookMapOutput)
}

type ProjectWebhookOutput struct{ *pulumi.OutputState }

func (ProjectWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectWebhook)(nil)).Elem()
}

func (o ProjectWebhookOutput) ToProjectWebhookOutput() ProjectWebhookOutput {
	return o
}

func (o ProjectWebhookOutput) ToProjectWebhookOutputWithContext(ctx context.Context) ProjectWebhookOutput {
	return o
}

// The address of the webhook
func (o ProjectWebhookOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectWebhook) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// authentication header for you the webhook
func (o ProjectWebhookOutput) AuthHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectWebhook) pulumi.StringPtrOutput { return v.AuthHeader }).(pulumi.StringPtrOutput)
}

// _ (Optional, string) A description of the webhook
func (o ProjectWebhookOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectWebhook) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// , To enable / disable the webhook. Default `true`
func (o ProjectWebhookOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProjectWebhook) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// ) The type events you want to subscript to can be `DELETE_ARTIFACT`, `PULL_ARTIFACT`, `PUSH_ARTIFACT`, `QUOTA_EXCEED`, `QUOTA_WARNING`, `REPLICATION`, `SCANNING_FAILED`, `SCANNING_COMPLETED`, `TAG_RETENTION`
func (o ProjectWebhookOutput) EventsTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProjectWebhook) pulumi.StringArrayOutput { return v.EventsTypes }).(pulumi.StringArrayOutput)
}

// The name of the webhook that will be created in harbor.
func (o ProjectWebhookOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectWebhook) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The notification type either `http` or `slack`
func (o ProjectWebhookOutput) NotifyType() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectWebhook) pulumi.StringOutput { return v.NotifyType }).(pulumi.StringOutput)
}

// The project id (**/projects/ID**) of the harbor that webhook related to.
func (o ProjectWebhookOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectWebhook) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// checks the for validate SSL certificate.
func (o ProjectWebhookOutput) SkipCertVerify() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProjectWebhook) pulumi.BoolPtrOutput { return v.SkipCertVerify }).(pulumi.BoolPtrOutput)
}

type ProjectWebhookArrayOutput struct{ *pulumi.OutputState }

func (ProjectWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectWebhook)(nil)).Elem()
}

func (o ProjectWebhookArrayOutput) ToProjectWebhookArrayOutput() ProjectWebhookArrayOutput {
	return o
}

func (o ProjectWebhookArrayOutput) ToProjectWebhookArrayOutputWithContext(ctx context.Context) ProjectWebhookArrayOutput {
	return o
}

func (o ProjectWebhookArrayOutput) Index(i pulumi.IntInput) ProjectWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectWebhook {
		return vs[0].([]*ProjectWebhook)[vs[1].(int)]
	}).(ProjectWebhookOutput)
}

type ProjectWebhookMapOutput struct{ *pulumi.OutputState }

func (ProjectWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectWebhook)(nil)).Elem()
}

func (o ProjectWebhookMapOutput) ToProjectWebhookMapOutput() ProjectWebhookMapOutput {
	return o
}

func (o ProjectWebhookMapOutput) ToProjectWebhookMapOutputWithContext(ctx context.Context) ProjectWebhookMapOutput {
	return o
}

func (o ProjectWebhookMapOutput) MapIndex(k pulumi.StringInput) ProjectWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectWebhook {
		return vs[0].(map[string]*ProjectWebhook)[vs[1].(string)]
	}).(ProjectWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectWebhookInput)(nil)).Elem(), &ProjectWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectWebhookArrayInput)(nil)).Elem(), ProjectWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectWebhookMapInput)(nil)).Elem(), ProjectWebhookMap{})
	pulumi.RegisterOutputType(ProjectWebhookOutput{})
	pulumi.RegisterOutputType(ProjectWebhookArrayOutput{})
	pulumi.RegisterOutputType(ProjectWebhookMapOutput{})
}
