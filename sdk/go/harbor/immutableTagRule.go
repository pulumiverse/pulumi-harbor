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
// Harbor immutable tag rule can be imported using the `project and immutabletagrule ids` eg, `<break><break>```sh<break> $ pulumi import harbor:index/immutableTagRule:ImmutableTagRule main /projects/4/immutabletagrules/25 <break>```<break><break>`
type ImmutableTagRule struct {
	pulumi.CustomResourceState

	// Specify if the rule is disable or not. Defaults to `false`
	Disabled  pulumi.BoolPtrOutput `pulumi:"disabled"`
	ProjectId pulumi.StringOutput  `pulumi:"projectId"`
	// For the repositories excuding.
	RepoExcluding pulumi.StringPtrOutput `pulumi:"repoExcluding"`
	// For the repositories matching.
	RepoMatching pulumi.StringPtrOutput `pulumi:"repoMatching"`
	// For the tag excuding.
	TagExcluding pulumi.StringPtrOutput `pulumi:"tagExcluding"`
	// For the tag matching.
	TagMatching pulumi.StringPtrOutput `pulumi:"tagMatching"`
}

// NewImmutableTagRule registers a new resource with the given unique name, arguments, and options.
func NewImmutableTagRule(ctx *pulumi.Context,
	name string, args *ImmutableTagRuleArgs, opts ...pulumi.ResourceOption) (*ImmutableTagRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ImmutableTagRule
	err := ctx.RegisterResource("harbor:index/immutableTagRule:ImmutableTagRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImmutableTagRule gets an existing ImmutableTagRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImmutableTagRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImmutableTagRuleState, opts ...pulumi.ResourceOption) (*ImmutableTagRule, error) {
	var resource ImmutableTagRule
	err := ctx.ReadResource("harbor:index/immutableTagRule:ImmutableTagRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImmutableTagRule resources.
type immutableTagRuleState struct {
	// Specify if the rule is disable or not. Defaults to `false`
	Disabled  *bool   `pulumi:"disabled"`
	ProjectId *string `pulumi:"projectId"`
	// For the repositories excuding.
	RepoExcluding *string `pulumi:"repoExcluding"`
	// For the repositories matching.
	RepoMatching *string `pulumi:"repoMatching"`
	// For the tag excuding.
	TagExcluding *string `pulumi:"tagExcluding"`
	// For the tag matching.
	TagMatching *string `pulumi:"tagMatching"`
}

type ImmutableTagRuleState struct {
	// Specify if the rule is disable or not. Defaults to `false`
	Disabled  pulumi.BoolPtrInput
	ProjectId pulumi.StringPtrInput
	// For the repositories excuding.
	RepoExcluding pulumi.StringPtrInput
	// For the repositories matching.
	RepoMatching pulumi.StringPtrInput
	// For the tag excuding.
	TagExcluding pulumi.StringPtrInput
	// For the tag matching.
	TagMatching pulumi.StringPtrInput
}

func (ImmutableTagRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*immutableTagRuleState)(nil)).Elem()
}

type immutableTagRuleArgs struct {
	// Specify if the rule is disable or not. Defaults to `false`
	Disabled  *bool  `pulumi:"disabled"`
	ProjectId string `pulumi:"projectId"`
	// For the repositories excuding.
	RepoExcluding *string `pulumi:"repoExcluding"`
	// For the repositories matching.
	RepoMatching *string `pulumi:"repoMatching"`
	// For the tag excuding.
	TagExcluding *string `pulumi:"tagExcluding"`
	// For the tag matching.
	TagMatching *string `pulumi:"tagMatching"`
}

// The set of arguments for constructing a ImmutableTagRule resource.
type ImmutableTagRuleArgs struct {
	// Specify if the rule is disable or not. Defaults to `false`
	Disabled  pulumi.BoolPtrInput
	ProjectId pulumi.StringInput
	// For the repositories excuding.
	RepoExcluding pulumi.StringPtrInput
	// For the repositories matching.
	RepoMatching pulumi.StringPtrInput
	// For the tag excuding.
	TagExcluding pulumi.StringPtrInput
	// For the tag matching.
	TagMatching pulumi.StringPtrInput
}

func (ImmutableTagRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*immutableTagRuleArgs)(nil)).Elem()
}

type ImmutableTagRuleInput interface {
	pulumi.Input

	ToImmutableTagRuleOutput() ImmutableTagRuleOutput
	ToImmutableTagRuleOutputWithContext(ctx context.Context) ImmutableTagRuleOutput
}

func (*ImmutableTagRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ImmutableTagRule)(nil)).Elem()
}

func (i *ImmutableTagRule) ToImmutableTagRuleOutput() ImmutableTagRuleOutput {
	return i.ToImmutableTagRuleOutputWithContext(context.Background())
}

func (i *ImmutableTagRule) ToImmutableTagRuleOutputWithContext(ctx context.Context) ImmutableTagRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImmutableTagRuleOutput)
}

// ImmutableTagRuleArrayInput is an input type that accepts ImmutableTagRuleArray and ImmutableTagRuleArrayOutput values.
// You can construct a concrete instance of `ImmutableTagRuleArrayInput` via:
//
//	ImmutableTagRuleArray{ ImmutableTagRuleArgs{...} }
type ImmutableTagRuleArrayInput interface {
	pulumi.Input

	ToImmutableTagRuleArrayOutput() ImmutableTagRuleArrayOutput
	ToImmutableTagRuleArrayOutputWithContext(context.Context) ImmutableTagRuleArrayOutput
}

type ImmutableTagRuleArray []ImmutableTagRuleInput

func (ImmutableTagRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImmutableTagRule)(nil)).Elem()
}

func (i ImmutableTagRuleArray) ToImmutableTagRuleArrayOutput() ImmutableTagRuleArrayOutput {
	return i.ToImmutableTagRuleArrayOutputWithContext(context.Background())
}

func (i ImmutableTagRuleArray) ToImmutableTagRuleArrayOutputWithContext(ctx context.Context) ImmutableTagRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImmutableTagRuleArrayOutput)
}

// ImmutableTagRuleMapInput is an input type that accepts ImmutableTagRuleMap and ImmutableTagRuleMapOutput values.
// You can construct a concrete instance of `ImmutableTagRuleMapInput` via:
//
//	ImmutableTagRuleMap{ "key": ImmutableTagRuleArgs{...} }
type ImmutableTagRuleMapInput interface {
	pulumi.Input

	ToImmutableTagRuleMapOutput() ImmutableTagRuleMapOutput
	ToImmutableTagRuleMapOutputWithContext(context.Context) ImmutableTagRuleMapOutput
}

type ImmutableTagRuleMap map[string]ImmutableTagRuleInput

func (ImmutableTagRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImmutableTagRule)(nil)).Elem()
}

func (i ImmutableTagRuleMap) ToImmutableTagRuleMapOutput() ImmutableTagRuleMapOutput {
	return i.ToImmutableTagRuleMapOutputWithContext(context.Background())
}

func (i ImmutableTagRuleMap) ToImmutableTagRuleMapOutputWithContext(ctx context.Context) ImmutableTagRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImmutableTagRuleMapOutput)
}

type ImmutableTagRuleOutput struct{ *pulumi.OutputState }

func (ImmutableTagRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImmutableTagRule)(nil)).Elem()
}

func (o ImmutableTagRuleOutput) ToImmutableTagRuleOutput() ImmutableTagRuleOutput {
	return o
}

func (o ImmutableTagRuleOutput) ToImmutableTagRuleOutputWithContext(ctx context.Context) ImmutableTagRuleOutput {
	return o
}

// Specify if the rule is disable or not. Defaults to `false`
func (o ImmutableTagRuleOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ImmutableTagRule) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

func (o ImmutableTagRuleOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ImmutableTagRule) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// For the repositories excuding.
func (o ImmutableTagRuleOutput) RepoExcluding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImmutableTagRule) pulumi.StringPtrOutput { return v.RepoExcluding }).(pulumi.StringPtrOutput)
}

// For the repositories matching.
func (o ImmutableTagRuleOutput) RepoMatching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImmutableTagRule) pulumi.StringPtrOutput { return v.RepoMatching }).(pulumi.StringPtrOutput)
}

// For the tag excuding.
func (o ImmutableTagRuleOutput) TagExcluding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImmutableTagRule) pulumi.StringPtrOutput { return v.TagExcluding }).(pulumi.StringPtrOutput)
}

// For the tag matching.
func (o ImmutableTagRuleOutput) TagMatching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImmutableTagRule) pulumi.StringPtrOutput { return v.TagMatching }).(pulumi.StringPtrOutput)
}

type ImmutableTagRuleArrayOutput struct{ *pulumi.OutputState }

func (ImmutableTagRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImmutableTagRule)(nil)).Elem()
}

func (o ImmutableTagRuleArrayOutput) ToImmutableTagRuleArrayOutput() ImmutableTagRuleArrayOutput {
	return o
}

func (o ImmutableTagRuleArrayOutput) ToImmutableTagRuleArrayOutputWithContext(ctx context.Context) ImmutableTagRuleArrayOutput {
	return o
}

func (o ImmutableTagRuleArrayOutput) Index(i pulumi.IntInput) ImmutableTagRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ImmutableTagRule {
		return vs[0].([]*ImmutableTagRule)[vs[1].(int)]
	}).(ImmutableTagRuleOutput)
}

type ImmutableTagRuleMapOutput struct{ *pulumi.OutputState }

func (ImmutableTagRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImmutableTagRule)(nil)).Elem()
}

func (o ImmutableTagRuleMapOutput) ToImmutableTagRuleMapOutput() ImmutableTagRuleMapOutput {
	return o
}

func (o ImmutableTagRuleMapOutput) ToImmutableTagRuleMapOutputWithContext(ctx context.Context) ImmutableTagRuleMapOutput {
	return o
}

func (o ImmutableTagRuleMapOutput) MapIndex(k pulumi.StringInput) ImmutableTagRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ImmutableTagRule {
		return vs[0].(map[string]*ImmutableTagRule)[vs[1].(string)]
	}).(ImmutableTagRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImmutableTagRuleInput)(nil)).Elem(), &ImmutableTagRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImmutableTagRuleArrayInput)(nil)).Elem(), ImmutableTagRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImmutableTagRuleMapInput)(nil)).Elem(), ImmutableTagRuleMap{})
	pulumi.RegisterOutputType(ImmutableTagRuleOutput{})
	pulumi.RegisterOutputType(ImmutableTagRuleArrayOutput{})
	pulumi.RegisterOutputType(ImmutableTagRuleMapOutput{})
}
