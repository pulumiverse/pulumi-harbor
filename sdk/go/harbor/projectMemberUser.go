// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package harbor

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-harbor/sdk/go/harbor"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			mainProject, err := harbor.NewProject(ctx, "mainProject", nil)
//			if err != nil {
//				return err
//			}
//			_, err = harbor.NewProjectMemberUser(ctx, "mainProjectMemberUser", &harbor.ProjectMemberUserArgs{
//				ProjectId: mainProject.ID(),
//				UserName:  pulumi.String("testing1"),
//				Role:      pulumi.String("projectadmin"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Harbor project member user can be imported using the `project id` and `member id` eg, `
//
// ```sh
//
//	$ pulumi import harbor:index/projectMemberUser:ProjectMemberUser main /projects/10/members/200
//
// ```
//
//	`
type ProjectMemberUser struct {
	pulumi.CustomResourceState

	MemberId  pulumi.IntOutput    `pulumi:"memberId"`
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	Role      pulumi.StringOutput `pulumi:"role"`
	UserName  pulumi.StringOutput `pulumi:"userName"`
}

// NewProjectMemberUser registers a new resource with the given unique name, arguments, and options.
func NewProjectMemberUser(ctx *pulumi.Context,
	name string, args *ProjectMemberUserArgs, opts ...pulumi.ResourceOption) (*ProjectMemberUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ProjectMemberUser
	err := ctx.RegisterResource("harbor:index/projectMemberUser:ProjectMemberUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectMemberUser gets an existing ProjectMemberUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectMemberUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectMemberUserState, opts ...pulumi.ResourceOption) (*ProjectMemberUser, error) {
	var resource ProjectMemberUser
	err := ctx.ReadResource("harbor:index/projectMemberUser:ProjectMemberUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectMemberUser resources.
type projectMemberUserState struct {
	MemberId  *int    `pulumi:"memberId"`
	ProjectId *string `pulumi:"projectId"`
	Role      *string `pulumi:"role"`
	UserName  *string `pulumi:"userName"`
}

type ProjectMemberUserState struct {
	MemberId  pulumi.IntPtrInput
	ProjectId pulumi.StringPtrInput
	Role      pulumi.StringPtrInput
	UserName  pulumi.StringPtrInput
}

func (ProjectMemberUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectMemberUserState)(nil)).Elem()
}

type projectMemberUserArgs struct {
	ProjectId string `pulumi:"projectId"`
	Role      string `pulumi:"role"`
	UserName  string `pulumi:"userName"`
}

// The set of arguments for constructing a ProjectMemberUser resource.
type ProjectMemberUserArgs struct {
	ProjectId pulumi.StringInput
	Role      pulumi.StringInput
	UserName  pulumi.StringInput
}

func (ProjectMemberUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectMemberUserArgs)(nil)).Elem()
}

type ProjectMemberUserInput interface {
	pulumi.Input

	ToProjectMemberUserOutput() ProjectMemberUserOutput
	ToProjectMemberUserOutputWithContext(ctx context.Context) ProjectMemberUserOutput
}

func (*ProjectMemberUser) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectMemberUser)(nil)).Elem()
}

func (i *ProjectMemberUser) ToProjectMemberUserOutput() ProjectMemberUserOutput {
	return i.ToProjectMemberUserOutputWithContext(context.Background())
}

func (i *ProjectMemberUser) ToProjectMemberUserOutputWithContext(ctx context.Context) ProjectMemberUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMemberUserOutput)
}

// ProjectMemberUserArrayInput is an input type that accepts ProjectMemberUserArray and ProjectMemberUserArrayOutput values.
// You can construct a concrete instance of `ProjectMemberUserArrayInput` via:
//
//	ProjectMemberUserArray{ ProjectMemberUserArgs{...} }
type ProjectMemberUserArrayInput interface {
	pulumi.Input

	ToProjectMemberUserArrayOutput() ProjectMemberUserArrayOutput
	ToProjectMemberUserArrayOutputWithContext(context.Context) ProjectMemberUserArrayOutput
}

type ProjectMemberUserArray []ProjectMemberUserInput

func (ProjectMemberUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectMemberUser)(nil)).Elem()
}

func (i ProjectMemberUserArray) ToProjectMemberUserArrayOutput() ProjectMemberUserArrayOutput {
	return i.ToProjectMemberUserArrayOutputWithContext(context.Background())
}

func (i ProjectMemberUserArray) ToProjectMemberUserArrayOutputWithContext(ctx context.Context) ProjectMemberUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMemberUserArrayOutput)
}

// ProjectMemberUserMapInput is an input type that accepts ProjectMemberUserMap and ProjectMemberUserMapOutput values.
// You can construct a concrete instance of `ProjectMemberUserMapInput` via:
//
//	ProjectMemberUserMap{ "key": ProjectMemberUserArgs{...} }
type ProjectMemberUserMapInput interface {
	pulumi.Input

	ToProjectMemberUserMapOutput() ProjectMemberUserMapOutput
	ToProjectMemberUserMapOutputWithContext(context.Context) ProjectMemberUserMapOutput
}

type ProjectMemberUserMap map[string]ProjectMemberUserInput

func (ProjectMemberUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectMemberUser)(nil)).Elem()
}

func (i ProjectMemberUserMap) ToProjectMemberUserMapOutput() ProjectMemberUserMapOutput {
	return i.ToProjectMemberUserMapOutputWithContext(context.Background())
}

func (i ProjectMemberUserMap) ToProjectMemberUserMapOutputWithContext(ctx context.Context) ProjectMemberUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMemberUserMapOutput)
}

type ProjectMemberUserOutput struct{ *pulumi.OutputState }

func (ProjectMemberUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectMemberUser)(nil)).Elem()
}

func (o ProjectMemberUserOutput) ToProjectMemberUserOutput() ProjectMemberUserOutput {
	return o
}

func (o ProjectMemberUserOutput) ToProjectMemberUserOutputWithContext(ctx context.Context) ProjectMemberUserOutput {
	return o
}

func (o ProjectMemberUserOutput) MemberId() pulumi.IntOutput {
	return o.ApplyT(func(v *ProjectMemberUser) pulumi.IntOutput { return v.MemberId }).(pulumi.IntOutput)
}

func (o ProjectMemberUserOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectMemberUser) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

func (o ProjectMemberUserOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectMemberUser) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o ProjectMemberUserOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectMemberUser) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

type ProjectMemberUserArrayOutput struct{ *pulumi.OutputState }

func (ProjectMemberUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectMemberUser)(nil)).Elem()
}

func (o ProjectMemberUserArrayOutput) ToProjectMemberUserArrayOutput() ProjectMemberUserArrayOutput {
	return o
}

func (o ProjectMemberUserArrayOutput) ToProjectMemberUserArrayOutputWithContext(ctx context.Context) ProjectMemberUserArrayOutput {
	return o
}

func (o ProjectMemberUserArrayOutput) Index(i pulumi.IntInput) ProjectMemberUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectMemberUser {
		return vs[0].([]*ProjectMemberUser)[vs[1].(int)]
	}).(ProjectMemberUserOutput)
}

type ProjectMemberUserMapOutput struct{ *pulumi.OutputState }

func (ProjectMemberUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectMemberUser)(nil)).Elem()
}

func (o ProjectMemberUserMapOutput) ToProjectMemberUserMapOutput() ProjectMemberUserMapOutput {
	return o
}

func (o ProjectMemberUserMapOutput) ToProjectMemberUserMapOutputWithContext(ctx context.Context) ProjectMemberUserMapOutput {
	return o
}

func (o ProjectMemberUserMapOutput) MapIndex(k pulumi.StringInput) ProjectMemberUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectMemberUser {
		return vs[0].(map[string]*ProjectMemberUser)[vs[1].(string)]
	}).(ProjectMemberUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMemberUserInput)(nil)).Elem(), &ProjectMemberUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMemberUserArrayInput)(nil)).Elem(), ProjectMemberUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMemberUserMapInput)(nil)).Elem(), ProjectMemberUserMap{})
	pulumi.RegisterOutputType(ProjectMemberUserOutput{})
	pulumi.RegisterOutputType(ProjectMemberUserArrayOutput{})
	pulumi.RegisterOutputType(ProjectMemberUserMapOutput{})
}
