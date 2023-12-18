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
// Harbor project member group can be imported using the `project id` and `member id` eg, `<break><break>```sh<break> $ pulumi import harbor:index/projectMemberGroup:ProjectMemberGroup main /projects/10/members/200 <break>```<break><break>`
type ProjectMemberGroup struct {
	pulumi.CustomResourceState

	GroupId     pulumi.IntPtrOutput    `pulumi:"groupId"`
	GroupName   pulumi.StringPtrOutput `pulumi:"groupName"`
	LdapGroupDn pulumi.StringPtrOutput `pulumi:"ldapGroupDn"`
	MemberId    pulumi.IntOutput       `pulumi:"memberId"`
	ProjectId   pulumi.StringOutput    `pulumi:"projectId"`
	Role        pulumi.StringOutput    `pulumi:"role"`
	Type        pulumi.StringOutput    `pulumi:"type"`
}

// NewProjectMemberGroup registers a new resource with the given unique name, arguments, and options.
func NewProjectMemberGroup(ctx *pulumi.Context,
	name string, args *ProjectMemberGroupArgs, opts ...pulumi.ResourceOption) (*ProjectMemberGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProjectMemberGroup
	err := ctx.RegisterResource("harbor:index/projectMemberGroup:ProjectMemberGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectMemberGroup gets an existing ProjectMemberGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectMemberGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectMemberGroupState, opts ...pulumi.ResourceOption) (*ProjectMemberGroup, error) {
	var resource ProjectMemberGroup
	err := ctx.ReadResource("harbor:index/projectMemberGroup:ProjectMemberGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectMemberGroup resources.
type projectMemberGroupState struct {
	GroupId     *int    `pulumi:"groupId"`
	GroupName   *string `pulumi:"groupName"`
	LdapGroupDn *string `pulumi:"ldapGroupDn"`
	MemberId    *int    `pulumi:"memberId"`
	ProjectId   *string `pulumi:"projectId"`
	Role        *string `pulumi:"role"`
	Type        *string `pulumi:"type"`
}

type ProjectMemberGroupState struct {
	GroupId     pulumi.IntPtrInput
	GroupName   pulumi.StringPtrInput
	LdapGroupDn pulumi.StringPtrInput
	MemberId    pulumi.IntPtrInput
	ProjectId   pulumi.StringPtrInput
	Role        pulumi.StringPtrInput
	Type        pulumi.StringPtrInput
}

func (ProjectMemberGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectMemberGroupState)(nil)).Elem()
}

type projectMemberGroupArgs struct {
	GroupId     *int    `pulumi:"groupId"`
	GroupName   *string `pulumi:"groupName"`
	LdapGroupDn *string `pulumi:"ldapGroupDn"`
	ProjectId   string  `pulumi:"projectId"`
	Role        string  `pulumi:"role"`
	Type        string  `pulumi:"type"`
}

// The set of arguments for constructing a ProjectMemberGroup resource.
type ProjectMemberGroupArgs struct {
	GroupId     pulumi.IntPtrInput
	GroupName   pulumi.StringPtrInput
	LdapGroupDn pulumi.StringPtrInput
	ProjectId   pulumi.StringInput
	Role        pulumi.StringInput
	Type        pulumi.StringInput
}

func (ProjectMemberGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectMemberGroupArgs)(nil)).Elem()
}

type ProjectMemberGroupInput interface {
	pulumi.Input

	ToProjectMemberGroupOutput() ProjectMemberGroupOutput
	ToProjectMemberGroupOutputWithContext(ctx context.Context) ProjectMemberGroupOutput
}

func (*ProjectMemberGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectMemberGroup)(nil)).Elem()
}

func (i *ProjectMemberGroup) ToProjectMemberGroupOutput() ProjectMemberGroupOutput {
	return i.ToProjectMemberGroupOutputWithContext(context.Background())
}

func (i *ProjectMemberGroup) ToProjectMemberGroupOutputWithContext(ctx context.Context) ProjectMemberGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMemberGroupOutput)
}

// ProjectMemberGroupArrayInput is an input type that accepts ProjectMemberGroupArray and ProjectMemberGroupArrayOutput values.
// You can construct a concrete instance of `ProjectMemberGroupArrayInput` via:
//
//	ProjectMemberGroupArray{ ProjectMemberGroupArgs{...} }
type ProjectMemberGroupArrayInput interface {
	pulumi.Input

	ToProjectMemberGroupArrayOutput() ProjectMemberGroupArrayOutput
	ToProjectMemberGroupArrayOutputWithContext(context.Context) ProjectMemberGroupArrayOutput
}

type ProjectMemberGroupArray []ProjectMemberGroupInput

func (ProjectMemberGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectMemberGroup)(nil)).Elem()
}

func (i ProjectMemberGroupArray) ToProjectMemberGroupArrayOutput() ProjectMemberGroupArrayOutput {
	return i.ToProjectMemberGroupArrayOutputWithContext(context.Background())
}

func (i ProjectMemberGroupArray) ToProjectMemberGroupArrayOutputWithContext(ctx context.Context) ProjectMemberGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMemberGroupArrayOutput)
}

// ProjectMemberGroupMapInput is an input type that accepts ProjectMemberGroupMap and ProjectMemberGroupMapOutput values.
// You can construct a concrete instance of `ProjectMemberGroupMapInput` via:
//
//	ProjectMemberGroupMap{ "key": ProjectMemberGroupArgs{...} }
type ProjectMemberGroupMapInput interface {
	pulumi.Input

	ToProjectMemberGroupMapOutput() ProjectMemberGroupMapOutput
	ToProjectMemberGroupMapOutputWithContext(context.Context) ProjectMemberGroupMapOutput
}

type ProjectMemberGroupMap map[string]ProjectMemberGroupInput

func (ProjectMemberGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectMemberGroup)(nil)).Elem()
}

func (i ProjectMemberGroupMap) ToProjectMemberGroupMapOutput() ProjectMemberGroupMapOutput {
	return i.ToProjectMemberGroupMapOutputWithContext(context.Background())
}

func (i ProjectMemberGroupMap) ToProjectMemberGroupMapOutputWithContext(ctx context.Context) ProjectMemberGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMemberGroupMapOutput)
}

type ProjectMemberGroupOutput struct{ *pulumi.OutputState }

func (ProjectMemberGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectMemberGroup)(nil)).Elem()
}

func (o ProjectMemberGroupOutput) ToProjectMemberGroupOutput() ProjectMemberGroupOutput {
	return o
}

func (o ProjectMemberGroupOutput) ToProjectMemberGroupOutputWithContext(ctx context.Context) ProjectMemberGroupOutput {
	return o
}

func (o ProjectMemberGroupOutput) GroupId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ProjectMemberGroup) pulumi.IntPtrOutput { return v.GroupId }).(pulumi.IntPtrOutput)
}

func (o ProjectMemberGroupOutput) GroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectMemberGroup) pulumi.StringPtrOutput { return v.GroupName }).(pulumi.StringPtrOutput)
}

func (o ProjectMemberGroupOutput) LdapGroupDn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectMemberGroup) pulumi.StringPtrOutput { return v.LdapGroupDn }).(pulumi.StringPtrOutput)
}

func (o ProjectMemberGroupOutput) MemberId() pulumi.IntOutput {
	return o.ApplyT(func(v *ProjectMemberGroup) pulumi.IntOutput { return v.MemberId }).(pulumi.IntOutput)
}

func (o ProjectMemberGroupOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectMemberGroup) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

func (o ProjectMemberGroupOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectMemberGroup) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o ProjectMemberGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectMemberGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type ProjectMemberGroupArrayOutput struct{ *pulumi.OutputState }

func (ProjectMemberGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectMemberGroup)(nil)).Elem()
}

func (o ProjectMemberGroupArrayOutput) ToProjectMemberGroupArrayOutput() ProjectMemberGroupArrayOutput {
	return o
}

func (o ProjectMemberGroupArrayOutput) ToProjectMemberGroupArrayOutputWithContext(ctx context.Context) ProjectMemberGroupArrayOutput {
	return o
}

func (o ProjectMemberGroupArrayOutput) Index(i pulumi.IntInput) ProjectMemberGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectMemberGroup {
		return vs[0].([]*ProjectMemberGroup)[vs[1].(int)]
	}).(ProjectMemberGroupOutput)
}

type ProjectMemberGroupMapOutput struct{ *pulumi.OutputState }

func (ProjectMemberGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectMemberGroup)(nil)).Elem()
}

func (o ProjectMemberGroupMapOutput) ToProjectMemberGroupMapOutput() ProjectMemberGroupMapOutput {
	return o
}

func (o ProjectMemberGroupMapOutput) ToProjectMemberGroupMapOutputWithContext(ctx context.Context) ProjectMemberGroupMapOutput {
	return o
}

func (o ProjectMemberGroupMapOutput) MapIndex(k pulumi.StringInput) ProjectMemberGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectMemberGroup {
		return vs[0].(map[string]*ProjectMemberGroup)[vs[1].(string)]
	}).(ProjectMemberGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMemberGroupInput)(nil)).Elem(), &ProjectMemberGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMemberGroupArrayInput)(nil)).Elem(), ProjectMemberGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMemberGroupMapInput)(nil)).Elem(), ProjectMemberGroupMap{})
	pulumi.RegisterOutputType(ProjectMemberGroupOutput{})
	pulumi.RegisterOutputType(ProjectMemberGroupArrayOutput{})
	pulumi.RegisterOutputType(ProjectMemberGroupMapOutput{})
}
