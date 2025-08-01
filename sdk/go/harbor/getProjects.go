// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package harbor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-harbor/sdk/v3/go/harbor/internal"
)

// ## Example Usage
func GetProjects(ctx *pulumi.Context, args *GetProjectsArgs, opts ...pulumi.InvokeOption) (*GetProjectsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetProjectsResult
	err := ctx.Invoke("harbor:index/getProjects:getProjects", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjects.
type GetProjectsArgs struct {
	// The name of the project.
	Name *string `pulumi:"name"`
	// If the project has public accessibility.
	Public *bool `pulumi:"public"`
	// The type of the project : Project or ProxyCache.
	Type *string `pulumi:"type"`
	// If the images will be scanned for vulnerabilities when push to harbor.
	VulnerabilityScanning *bool `pulumi:"vulnerabilityScanning"`
}

// A collection of values returned by getProjects.
type GetProjectsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the project.
	Name     *string              `pulumi:"name"`
	Projects []GetProjectsProject `pulumi:"projects"`
	// If the project has public accessibility.
	Public *bool `pulumi:"public"`
	// The type of the project : Project or ProxyCache.
	Type *string `pulumi:"type"`
	// If the images will be scanned for vulnerabilities when push to harbor.
	VulnerabilityScanning *bool `pulumi:"vulnerabilityScanning"`
}

func GetProjectsOutput(ctx *pulumi.Context, args GetProjectsOutputArgs, opts ...pulumi.InvokeOption) GetProjectsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetProjectsResultOutput, error) {
			args := v.(GetProjectsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("harbor:index/getProjects:getProjects", args, GetProjectsResultOutput{}, options).(GetProjectsResultOutput), nil
		}).(GetProjectsResultOutput)
}

// A collection of arguments for invoking getProjects.
type GetProjectsOutputArgs struct {
	// The name of the project.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// If the project has public accessibility.
	Public pulumi.BoolPtrInput `pulumi:"public"`
	// The type of the project : Project or ProxyCache.
	Type pulumi.StringPtrInput `pulumi:"type"`
	// If the images will be scanned for vulnerabilities when push to harbor.
	VulnerabilityScanning pulumi.BoolPtrInput `pulumi:"vulnerabilityScanning"`
}

func (GetProjectsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectsArgs)(nil)).Elem()
}

// A collection of values returned by getProjects.
type GetProjectsResultOutput struct{ *pulumi.OutputState }

func (GetProjectsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectsResult)(nil)).Elem()
}

func (o GetProjectsResultOutput) ToGetProjectsResultOutput() GetProjectsResultOutput {
	return o
}

func (o GetProjectsResultOutput) ToGetProjectsResultOutputWithContext(ctx context.Context) GetProjectsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetProjectsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the project.
func (o GetProjectsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetProjectsResultOutput) Projects() GetProjectsProjectArrayOutput {
	return o.ApplyT(func(v GetProjectsResult) []GetProjectsProject { return v.Projects }).(GetProjectsProjectArrayOutput)
}

// If the project has public accessibility.
func (o GetProjectsResultOutput) Public() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *bool { return v.Public }).(pulumi.BoolPtrOutput)
}

// The type of the project : Project or ProxyCache.
func (o GetProjectsResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// If the images will be scanned for vulnerabilities when push to harbor.
func (o GetProjectsResultOutput) VulnerabilityScanning() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetProjectsResult) *bool { return v.VulnerabilityScanning }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectsResultOutput{})
}
