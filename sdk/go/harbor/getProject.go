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
func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProjectResult
	err := ctx.Invoke("harbor:index/getProject:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProject.
type LookupProjectArgs struct {
	// The name of the project.
	Name string `pulumi:"name"`
}

// A collection of values returned by getProject.
type LookupProjectResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the project.
	Name string `pulumi:"name"`
	// The id of the project within harbor.
	ProjectId int `pulumi:"projectId"`
	// If the project has public accessibility.
	Public bool `pulumi:"public"`
	// The type of the project : Project or ProxyCache.
	Type string `pulumi:"type"`
	// If the images is scanned for vulnerabilities when push to harbor.
	VulnerabilityScanning bool `pulumi:"vulnerabilityScanning"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectResult, error) {
			args := v.(LookupProjectArgs)
			r, err := LookupProject(ctx, &args, opts...)
			var s LookupProjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProjectResultOutput)
}

// A collection of arguments for invoking getProject.
type LookupProjectOutputArgs struct {
	// The name of the project.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}

// A collection of values returned by getProject.
type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the project.
func (o LookupProjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Name }).(pulumi.StringOutput)
}

// The id of the project within harbor.
func (o LookupProjectResultOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProjectResult) int { return v.ProjectId }).(pulumi.IntOutput)
}

// If the project has public accessibility.
func (o LookupProjectResultOutput) Public() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.Public }).(pulumi.BoolOutput)
}

// The type of the project : Project or ProxyCache.
func (o LookupProjectResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectResult) string { return v.Type }).(pulumi.StringOutput)
}

// If the images is scanned for vulnerabilities when push to harbor.
func (o LookupProjectResultOutput) VulnerabilityScanning() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProjectResult) bool { return v.VulnerabilityScanning }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}
