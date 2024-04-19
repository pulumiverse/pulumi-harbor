// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package harbor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-harbor/sdk/v3/go/harbor/internal"
)

// ## Example Usage
//
// ### Global
//
// ### Project
//
// ## Import
//
// ```sh
// $ pulumi import harbor:index/label:Label main /labels/1
// ```
type Label struct {
	pulumi.CustomResourceState

	// The color of the label within harbor (Default: #FFFFF)
	Color pulumi.StringPtrOutput `pulumi:"color"`
	// The Description of the label within harbor
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The of name of the label within harbor.
	Name pulumi.StringOutput `pulumi:"name"`
	// The id of the project with harbor.
	ProjectId pulumi.StringPtrOutput `pulumi:"projectId"`
	Scope     pulumi.StringOutput    `pulumi:"scope"`
}

// NewLabel registers a new resource with the given unique name, arguments, and options.
func NewLabel(ctx *pulumi.Context,
	name string, args *LabelArgs, opts ...pulumi.ResourceOption) (*Label, error) {
	if args == nil {
		args = &LabelArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Label
	err := ctx.RegisterResource("harbor:index/label:Label", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLabel gets an existing Label resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLabel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LabelState, opts ...pulumi.ResourceOption) (*Label, error) {
	var resource Label
	err := ctx.ReadResource("harbor:index/label:Label", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Label resources.
type labelState struct {
	// The color of the label within harbor (Default: #FFFFF)
	Color *string `pulumi:"color"`
	// The Description of the label within harbor
	Description *string `pulumi:"description"`
	// The of name of the label within harbor.
	Name *string `pulumi:"name"`
	// The id of the project with harbor.
	ProjectId *string `pulumi:"projectId"`
	Scope     *string `pulumi:"scope"`
}

type LabelState struct {
	// The color of the label within harbor (Default: #FFFFF)
	Color pulumi.StringPtrInput
	// The Description of the label within harbor
	Description pulumi.StringPtrInput
	// The of name of the label within harbor.
	Name pulumi.StringPtrInput
	// The id of the project with harbor.
	ProjectId pulumi.StringPtrInput
	Scope     pulumi.StringPtrInput
}

func (LabelState) ElementType() reflect.Type {
	return reflect.TypeOf((*labelState)(nil)).Elem()
}

type labelArgs struct {
	// The color of the label within harbor (Default: #FFFFF)
	Color *string `pulumi:"color"`
	// The Description of the label within harbor
	Description *string `pulumi:"description"`
	// The of name of the label within harbor.
	Name *string `pulumi:"name"`
	// The id of the project with harbor.
	ProjectId *string `pulumi:"projectId"`
}

// The set of arguments for constructing a Label resource.
type LabelArgs struct {
	// The color of the label within harbor (Default: #FFFFF)
	Color pulumi.StringPtrInput
	// The Description of the label within harbor
	Description pulumi.StringPtrInput
	// The of name of the label within harbor.
	Name pulumi.StringPtrInput
	// The id of the project with harbor.
	ProjectId pulumi.StringPtrInput
}

func (LabelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*labelArgs)(nil)).Elem()
}

type LabelInput interface {
	pulumi.Input

	ToLabelOutput() LabelOutput
	ToLabelOutputWithContext(ctx context.Context) LabelOutput
}

func (*Label) ElementType() reflect.Type {
	return reflect.TypeOf((**Label)(nil)).Elem()
}

func (i *Label) ToLabelOutput() LabelOutput {
	return i.ToLabelOutputWithContext(context.Background())
}

func (i *Label) ToLabelOutputWithContext(ctx context.Context) LabelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelOutput)
}

// LabelArrayInput is an input type that accepts LabelArray and LabelArrayOutput values.
// You can construct a concrete instance of `LabelArrayInput` via:
//
//	LabelArray{ LabelArgs{...} }
type LabelArrayInput interface {
	pulumi.Input

	ToLabelArrayOutput() LabelArrayOutput
	ToLabelArrayOutputWithContext(context.Context) LabelArrayOutput
}

type LabelArray []LabelInput

func (LabelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Label)(nil)).Elem()
}

func (i LabelArray) ToLabelArrayOutput() LabelArrayOutput {
	return i.ToLabelArrayOutputWithContext(context.Background())
}

func (i LabelArray) ToLabelArrayOutputWithContext(ctx context.Context) LabelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelArrayOutput)
}

// LabelMapInput is an input type that accepts LabelMap and LabelMapOutput values.
// You can construct a concrete instance of `LabelMapInput` via:
//
//	LabelMap{ "key": LabelArgs{...} }
type LabelMapInput interface {
	pulumi.Input

	ToLabelMapOutput() LabelMapOutput
	ToLabelMapOutputWithContext(context.Context) LabelMapOutput
}

type LabelMap map[string]LabelInput

func (LabelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Label)(nil)).Elem()
}

func (i LabelMap) ToLabelMapOutput() LabelMapOutput {
	return i.ToLabelMapOutputWithContext(context.Background())
}

func (i LabelMap) ToLabelMapOutputWithContext(ctx context.Context) LabelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelMapOutput)
}

type LabelOutput struct{ *pulumi.OutputState }

func (LabelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Label)(nil)).Elem()
}

func (o LabelOutput) ToLabelOutput() LabelOutput {
	return o
}

func (o LabelOutput) ToLabelOutputWithContext(ctx context.Context) LabelOutput {
	return o
}

// The color of the label within harbor (Default: #FFFFF)
func (o LabelOutput) Color() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Label) pulumi.StringPtrOutput { return v.Color }).(pulumi.StringPtrOutput)
}

// The Description of the label within harbor
func (o LabelOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Label) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The of name of the label within harbor.
func (o LabelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Label) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The id of the project with harbor.
func (o LabelOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Label) pulumi.StringPtrOutput { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LabelOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v *Label) pulumi.StringOutput { return v.Scope }).(pulumi.StringOutput)
}

type LabelArrayOutput struct{ *pulumi.OutputState }

func (LabelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Label)(nil)).Elem()
}

func (o LabelArrayOutput) ToLabelArrayOutput() LabelArrayOutput {
	return o
}

func (o LabelArrayOutput) ToLabelArrayOutputWithContext(ctx context.Context) LabelArrayOutput {
	return o
}

func (o LabelArrayOutput) Index(i pulumi.IntInput) LabelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Label {
		return vs[0].([]*Label)[vs[1].(int)]
	}).(LabelOutput)
}

type LabelMapOutput struct{ *pulumi.OutputState }

func (LabelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Label)(nil)).Elem()
}

func (o LabelMapOutput) ToLabelMapOutput() LabelMapOutput {
	return o
}

func (o LabelMapOutput) ToLabelMapOutputWithContext(ctx context.Context) LabelMapOutput {
	return o
}

func (o LabelMapOutput) MapIndex(k pulumi.StringInput) LabelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Label {
		return vs[0].(map[string]*Label)[vs[1].(string)]
	}).(LabelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LabelInput)(nil)).Elem(), &Label{})
	pulumi.RegisterInputType(reflect.TypeOf((*LabelArrayInput)(nil)).Elem(), LabelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LabelMapInput)(nil)).Elem(), LabelMap{})
	pulumi.RegisterOutputType(LabelOutput{})
	pulumi.RegisterOutputType(LabelArrayOutput{})
	pulumi.RegisterOutputType(LabelMapOutput{})
}
