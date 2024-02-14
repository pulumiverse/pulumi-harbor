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
// ## Harbor project example as proxy cache
//
// ## Import
//
// Harbor project can be imported using the `project id` eg,<break><break> ` <break><break> ```sh<break> $ pulumi import harbor:index/project:Project main /projects/1 <break>```<break><break>  `<break><break>
type Project struct {
	pulumi.CustomResourceState

	// Project allowlist allows vulnerabilities in this list to be ignored in this project when pushing and pulling images. Should be in the format or `["CVE-123", "CVE-145"]` or `["CVE-123"]`
	CveAllowlists pulumi.StringArrayOutput `pulumi:"cveAllowlists"`
	// Prevent deployment of images with vulnerability severity equal or higher than the specified value. Images must be scanned before this takes effect. Possible values: `critical`, `high`, `medium`, `low`, `none`. (Default: `""` - empty)
	DeploymentSecurity pulumi.StringPtrOutput `pulumi:"deploymentSecurity"`
	// Enables Content Trust for project. When enabled it queries the embedded docker notary server. Can be set to `"true"` or `"false"` (Default: false)
	EnableContentTrust pulumi.BoolPtrOutput `pulumi:"enableContentTrust"`
	// Enables Content Trust Cosign for project. When enabled it queries Cosign. Can be set to `"true"` or `"false"` (Default: false)
	EnableContentTrustCosign pulumi.BoolPtrOutput `pulumi:"enableContentTrustCosign"`
	// A boolean that indicates all repositories should be deleted from the project so that the project can be destroyed without error. These repositories are *not* recoverable.
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// The name of the project that will be created in harbor.
	Name pulumi.StringOutput `pulumi:"name"`
	// The id of the project with harbor.
	ProjectId pulumi.IntOutput `pulumi:"projectId"`
	// The project will be public accessibility. Can be set to `"true"` or `"false"` (Default: false)
	Public pulumi.BoolPtrOutput `pulumi:"public"`
	// To enable project as Proxy Cache
	RegistryId pulumi.IntOutput `pulumi:"registryId"`
	// The storage quota of the project in GB's
	StorageQuota pulumi.IntPtrOutput `pulumi:"storageQuota"`
	// Images will be scanned for vulnerabilities when push to harbor. Can be set to `"true"` or `"false"` (Default: true)
	VulnerabilityScanning pulumi.BoolPtrOutput `pulumi:"vulnerabilityScanning"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		args = &ProjectArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Project
	err := ctx.RegisterResource("harbor:index/project:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("harbor:index/project:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
	// Project allowlist allows vulnerabilities in this list to be ignored in this project when pushing and pulling images. Should be in the format or `["CVE-123", "CVE-145"]` or `["CVE-123"]`
	CveAllowlists []string `pulumi:"cveAllowlists"`
	// Prevent deployment of images with vulnerability severity equal or higher than the specified value. Images must be scanned before this takes effect. Possible values: `critical`, `high`, `medium`, `low`, `none`. (Default: `""` - empty)
	DeploymentSecurity *string `pulumi:"deploymentSecurity"`
	// Enables Content Trust for project. When enabled it queries the embedded docker notary server. Can be set to `"true"` or `"false"` (Default: false)
	EnableContentTrust *bool `pulumi:"enableContentTrust"`
	// Enables Content Trust Cosign for project. When enabled it queries Cosign. Can be set to `"true"` or `"false"` (Default: false)
	EnableContentTrustCosign *bool `pulumi:"enableContentTrustCosign"`
	// A boolean that indicates all repositories should be deleted from the project so that the project can be destroyed without error. These repositories are *not* recoverable.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// The name of the project that will be created in harbor.
	Name *string `pulumi:"name"`
	// The id of the project with harbor.
	ProjectId *int `pulumi:"projectId"`
	// The project will be public accessibility. Can be set to `"true"` or `"false"` (Default: false)
	Public *bool `pulumi:"public"`
	// To enable project as Proxy Cache
	RegistryId *int `pulumi:"registryId"`
	// The storage quota of the project in GB's
	StorageQuota *int `pulumi:"storageQuota"`
	// Images will be scanned for vulnerabilities when push to harbor. Can be set to `"true"` or `"false"` (Default: true)
	VulnerabilityScanning *bool `pulumi:"vulnerabilityScanning"`
}

type ProjectState struct {
	// Project allowlist allows vulnerabilities in this list to be ignored in this project when pushing and pulling images. Should be in the format or `["CVE-123", "CVE-145"]` or `["CVE-123"]`
	CveAllowlists pulumi.StringArrayInput
	// Prevent deployment of images with vulnerability severity equal or higher than the specified value. Images must be scanned before this takes effect. Possible values: `critical`, `high`, `medium`, `low`, `none`. (Default: `""` - empty)
	DeploymentSecurity pulumi.StringPtrInput
	// Enables Content Trust for project. When enabled it queries the embedded docker notary server. Can be set to `"true"` or `"false"` (Default: false)
	EnableContentTrust pulumi.BoolPtrInput
	// Enables Content Trust Cosign for project. When enabled it queries Cosign. Can be set to `"true"` or `"false"` (Default: false)
	EnableContentTrustCosign pulumi.BoolPtrInput
	// A boolean that indicates all repositories should be deleted from the project so that the project can be destroyed without error. These repositories are *not* recoverable.
	ForceDestroy pulumi.BoolPtrInput
	// The name of the project that will be created in harbor.
	Name pulumi.StringPtrInput
	// The id of the project with harbor.
	ProjectId pulumi.IntPtrInput
	// The project will be public accessibility. Can be set to `"true"` or `"false"` (Default: false)
	Public pulumi.BoolPtrInput
	// To enable project as Proxy Cache
	RegistryId pulumi.IntPtrInput
	// The storage quota of the project in GB's
	StorageQuota pulumi.IntPtrInput
	// Images will be scanned for vulnerabilities when push to harbor. Can be set to `"true"` or `"false"` (Default: true)
	VulnerabilityScanning pulumi.BoolPtrInput
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// Project allowlist allows vulnerabilities in this list to be ignored in this project when pushing and pulling images. Should be in the format or `["CVE-123", "CVE-145"]` or `["CVE-123"]`
	CveAllowlists []string `pulumi:"cveAllowlists"`
	// Prevent deployment of images with vulnerability severity equal or higher than the specified value. Images must be scanned before this takes effect. Possible values: `critical`, `high`, `medium`, `low`, `none`. (Default: `""` - empty)
	DeploymentSecurity *string `pulumi:"deploymentSecurity"`
	// Enables Content Trust for project. When enabled it queries the embedded docker notary server. Can be set to `"true"` or `"false"` (Default: false)
	EnableContentTrust *bool `pulumi:"enableContentTrust"`
	// Enables Content Trust Cosign for project. When enabled it queries Cosign. Can be set to `"true"` or `"false"` (Default: false)
	EnableContentTrustCosign *bool `pulumi:"enableContentTrustCosign"`
	// A boolean that indicates all repositories should be deleted from the project so that the project can be destroyed without error. These repositories are *not* recoverable.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// The name of the project that will be created in harbor.
	Name *string `pulumi:"name"`
	// The project will be public accessibility. Can be set to `"true"` or `"false"` (Default: false)
	Public *bool `pulumi:"public"`
	// To enable project as Proxy Cache
	RegistryId *int `pulumi:"registryId"`
	// The storage quota of the project in GB's
	StorageQuota *int `pulumi:"storageQuota"`
	// Images will be scanned for vulnerabilities when push to harbor. Can be set to `"true"` or `"false"` (Default: true)
	VulnerabilityScanning *bool `pulumi:"vulnerabilityScanning"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// Project allowlist allows vulnerabilities in this list to be ignored in this project when pushing and pulling images. Should be in the format or `["CVE-123", "CVE-145"]` or `["CVE-123"]`
	CveAllowlists pulumi.StringArrayInput
	// Prevent deployment of images with vulnerability severity equal or higher than the specified value. Images must be scanned before this takes effect. Possible values: `critical`, `high`, `medium`, `low`, `none`. (Default: `""` - empty)
	DeploymentSecurity pulumi.StringPtrInput
	// Enables Content Trust for project. When enabled it queries the embedded docker notary server. Can be set to `"true"` or `"false"` (Default: false)
	EnableContentTrust pulumi.BoolPtrInput
	// Enables Content Trust Cosign for project. When enabled it queries Cosign. Can be set to `"true"` or `"false"` (Default: false)
	EnableContentTrustCosign pulumi.BoolPtrInput
	// A boolean that indicates all repositories should be deleted from the project so that the project can be destroyed without error. These repositories are *not* recoverable.
	ForceDestroy pulumi.BoolPtrInput
	// The name of the project that will be created in harbor.
	Name pulumi.StringPtrInput
	// The project will be public accessibility. Can be set to `"true"` or `"false"` (Default: false)
	Public pulumi.BoolPtrInput
	// To enable project as Proxy Cache
	RegistryId pulumi.IntPtrInput
	// The storage quota of the project in GB's
	StorageQuota pulumi.IntPtrInput
	// Images will be scanned for vulnerabilities when push to harbor. Can be set to `"true"` or `"false"` (Default: true)
	VulnerabilityScanning pulumi.BoolPtrInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}

type ProjectInput interface {
	pulumi.Input

	ToProjectOutput() ProjectOutput
	ToProjectOutputWithContext(ctx context.Context) ProjectOutput
}

func (*Project) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (i *Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i *Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

// ProjectArrayInput is an input type that accepts ProjectArray and ProjectArrayOutput values.
// You can construct a concrete instance of `ProjectArrayInput` via:
//
//	ProjectArray{ ProjectArgs{...} }
type ProjectArrayInput interface {
	pulumi.Input

	ToProjectArrayOutput() ProjectArrayOutput
	ToProjectArrayOutputWithContext(context.Context) ProjectArrayOutput
}

type ProjectArray []ProjectInput

func (ProjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Project)(nil)).Elem()
}

func (i ProjectArray) ToProjectArrayOutput() ProjectArrayOutput {
	return i.ToProjectArrayOutputWithContext(context.Background())
}

func (i ProjectArray) ToProjectArrayOutputWithContext(ctx context.Context) ProjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectArrayOutput)
}

// ProjectMapInput is an input type that accepts ProjectMap and ProjectMapOutput values.
// You can construct a concrete instance of `ProjectMapInput` via:
//
//	ProjectMap{ "key": ProjectArgs{...} }
type ProjectMapInput interface {
	pulumi.Input

	ToProjectMapOutput() ProjectMapOutput
	ToProjectMapOutputWithContext(context.Context) ProjectMapOutput
}

type ProjectMap map[string]ProjectInput

func (ProjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Project)(nil)).Elem()
}

func (i ProjectMap) ToProjectMapOutput() ProjectMapOutput {
	return i.ToProjectMapOutputWithContext(context.Background())
}

func (i ProjectMap) ToProjectMapOutputWithContext(ctx context.Context) ProjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMapOutput)
}

type ProjectOutput struct{ *pulumi.OutputState }

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

// Project allowlist allows vulnerabilities in this list to be ignored in this project when pushing and pulling images. Should be in the format or `["CVE-123", "CVE-145"]` or `["CVE-123"]`
func (o ProjectOutput) CveAllowlists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Project) pulumi.StringArrayOutput { return v.CveAllowlists }).(pulumi.StringArrayOutput)
}

// Prevent deployment of images with vulnerability severity equal or higher than the specified value. Images must be scanned before this takes effect. Possible values: `critical`, `high`, `medium`, `low`, `none`. (Default: `""` - empty)
func (o ProjectOutput) DeploymentSecurity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.DeploymentSecurity }).(pulumi.StringPtrOutput)
}

// Enables Content Trust for project. When enabled it queries the embedded docker notary server. Can be set to `"true"` or `"false"` (Default: false)
func (o ProjectOutput) EnableContentTrust() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.BoolPtrOutput { return v.EnableContentTrust }).(pulumi.BoolPtrOutput)
}

// Enables Content Trust Cosign for project. When enabled it queries Cosign. Can be set to `"true"` or `"false"` (Default: false)
func (o ProjectOutput) EnableContentTrustCosign() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.BoolPtrOutput { return v.EnableContentTrustCosign }).(pulumi.BoolPtrOutput)
}

// A boolean that indicates all repositories should be deleted from the project so that the project can be destroyed without error. These repositories are *not* recoverable.
func (o ProjectOutput) ForceDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.BoolPtrOutput { return v.ForceDestroy }).(pulumi.BoolPtrOutput)
}

// The name of the project that will be created in harbor.
func (o ProjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The id of the project with harbor.
func (o ProjectOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v *Project) pulumi.IntOutput { return v.ProjectId }).(pulumi.IntOutput)
}

// The project will be public accessibility. Can be set to `"true"` or `"false"` (Default: false)
func (o ProjectOutput) Public() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.BoolPtrOutput { return v.Public }).(pulumi.BoolPtrOutput)
}

// To enable project as Proxy Cache
func (o ProjectOutput) RegistryId() pulumi.IntOutput {
	return o.ApplyT(func(v *Project) pulumi.IntOutput { return v.RegistryId }).(pulumi.IntOutput)
}

// The storage quota of the project in GB's
func (o ProjectOutput) StorageQuota() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.IntPtrOutput { return v.StorageQuota }).(pulumi.IntPtrOutput)
}

// Images will be scanned for vulnerabilities when push to harbor. Can be set to `"true"` or `"false"` (Default: true)
func (o ProjectOutput) VulnerabilityScanning() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.BoolPtrOutput { return v.VulnerabilityScanning }).(pulumi.BoolPtrOutput)
}

type ProjectArrayOutput struct{ *pulumi.OutputState }

func (ProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Project)(nil)).Elem()
}

func (o ProjectArrayOutput) ToProjectArrayOutput() ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) ToProjectArrayOutputWithContext(ctx context.Context) ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) Index(i pulumi.IntInput) ProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Project {
		return vs[0].([]*Project)[vs[1].(int)]
	}).(ProjectOutput)
}

type ProjectMapOutput struct{ *pulumi.OutputState }

func (ProjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Project)(nil)).Elem()
}

func (o ProjectMapOutput) ToProjectMapOutput() ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) ToProjectMapOutputWithContext(ctx context.Context) ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) MapIndex(k pulumi.StringInput) ProjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Project {
		return vs[0].(map[string]*Project)[vs[1].(string)]
	}).(ProjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectInput)(nil)).Elem(), &Project{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectArrayInput)(nil)).Elem(), ProjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMapInput)(nil)).Elem(), ProjectMap{})
	pulumi.RegisterOutputType(ProjectOutput{})
	pulumi.RegisterOutputType(ProjectArrayOutput{})
	pulumi.RegisterOutputType(ProjectMapOutput{})
}
