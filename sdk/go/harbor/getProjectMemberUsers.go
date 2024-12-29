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
func GetProjectMemberUsers(ctx *pulumi.Context, args *GetProjectMemberUsersArgs, opts ...pulumi.InvokeOption) (*GetProjectMemberUsersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetProjectMemberUsersResult
	err := ctx.Invoke("harbor:index/getProjectMemberUsers:getProjectMemberUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjectMemberUsers.
type GetProjectMemberUsersArgs struct {
	// The id of the project within harbor.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getProjectMemberUsers.
type GetProjectMemberUsersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The id of the project within harbor.
	ProjectId          string                                   `pulumi:"projectId"`
	ProjectMemberUsers []GetProjectMemberUsersProjectMemberUser `pulumi:"projectMemberUsers"`
}

func GetProjectMemberUsersOutput(ctx *pulumi.Context, args GetProjectMemberUsersOutputArgs, opts ...pulumi.InvokeOption) GetProjectMemberUsersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProjectMemberUsersResultOutput, error) {
			args := v.(GetProjectMemberUsersArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetProjectMemberUsersResult
			secret, err := ctx.InvokePackageRaw("harbor:index/getProjectMemberUsers:getProjectMemberUsers", args, &rv, "", opts...)
			if err != nil {
				return GetProjectMemberUsersResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetProjectMemberUsersResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetProjectMemberUsersResultOutput), nil
			}
			return output, nil
		}).(GetProjectMemberUsersResultOutput)
}

// A collection of arguments for invoking getProjectMemberUsers.
type GetProjectMemberUsersOutputArgs struct {
	// The id of the project within harbor.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (GetProjectMemberUsersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectMemberUsersArgs)(nil)).Elem()
}

// A collection of values returned by getProjectMemberUsers.
type GetProjectMemberUsersResultOutput struct{ *pulumi.OutputState }

func (GetProjectMemberUsersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectMemberUsersResult)(nil)).Elem()
}

func (o GetProjectMemberUsersResultOutput) ToGetProjectMemberUsersResultOutput() GetProjectMemberUsersResultOutput {
	return o
}

func (o GetProjectMemberUsersResultOutput) ToGetProjectMemberUsersResultOutputWithContext(ctx context.Context) GetProjectMemberUsersResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetProjectMemberUsersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectMemberUsersResult) string { return v.Id }).(pulumi.StringOutput)
}

// The id of the project within harbor.
func (o GetProjectMemberUsersResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectMemberUsersResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o GetProjectMemberUsersResultOutput) ProjectMemberUsers() GetProjectMemberUsersProjectMemberUserArrayOutput {
	return o.ApplyT(func(v GetProjectMemberUsersResult) []GetProjectMemberUsersProjectMemberUser {
		return v.ProjectMemberUsers
	}).(GetProjectMemberUsersProjectMemberUserArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectMemberUsersResultOutput{})
}