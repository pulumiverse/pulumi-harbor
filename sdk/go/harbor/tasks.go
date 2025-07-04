// Code generated by pulumi-language-go DO NOT EDIT.
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
type Tasks struct {
	pulumi.CustomResourceState

	// The frequency of the vulnerability scanning is done. Can be to **"hourly"**, **"daily"** or **"weekly"**
	VulnerabilityScanPolicy pulumi.StringOutput `pulumi:"vulnerabilityScanPolicy"`
}

// NewTasks registers a new resource with the given unique name, arguments, and options.
func NewTasks(ctx *pulumi.Context,
	name string, args *TasksArgs, opts ...pulumi.ResourceOption) (*Tasks, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VulnerabilityScanPolicy == nil {
		return nil, errors.New("invalid value for required argument 'VulnerabilityScanPolicy'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Tasks
	err := ctx.RegisterResource("harbor:index/tasks:Tasks", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTasks gets an existing Tasks resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTasks(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TasksState, opts ...pulumi.ResourceOption) (*Tasks, error) {
	var resource Tasks
	err := ctx.ReadResource("harbor:index/tasks:Tasks", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Tasks resources.
type tasksState struct {
	// The frequency of the vulnerability scanning is done. Can be to **"hourly"**, **"daily"** or **"weekly"**
	VulnerabilityScanPolicy *string `pulumi:"vulnerabilityScanPolicy"`
}

type TasksState struct {
	// The frequency of the vulnerability scanning is done. Can be to **"hourly"**, **"daily"** or **"weekly"**
	VulnerabilityScanPolicy pulumi.StringPtrInput
}

func (TasksState) ElementType() reflect.Type {
	return reflect.TypeOf((*tasksState)(nil)).Elem()
}

type tasksArgs struct {
	// The frequency of the vulnerability scanning is done. Can be to **"hourly"**, **"daily"** or **"weekly"**
	VulnerabilityScanPolicy string `pulumi:"vulnerabilityScanPolicy"`
}

// The set of arguments for constructing a Tasks resource.
type TasksArgs struct {
	// The frequency of the vulnerability scanning is done. Can be to **"hourly"**, **"daily"** or **"weekly"**
	VulnerabilityScanPolicy pulumi.StringInput
}

func (TasksArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tasksArgs)(nil)).Elem()
}

type TasksInput interface {
	pulumi.Input

	ToTasksOutput() TasksOutput
	ToTasksOutputWithContext(ctx context.Context) TasksOutput
}

func (*Tasks) ElementType() reflect.Type {
	return reflect.TypeOf((**Tasks)(nil)).Elem()
}

func (i *Tasks) ToTasksOutput() TasksOutput {
	return i.ToTasksOutputWithContext(context.Background())
}

func (i *Tasks) ToTasksOutputWithContext(ctx context.Context) TasksOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TasksOutput)
}

// TasksArrayInput is an input type that accepts TasksArray and TasksArrayOutput values.
// You can construct a concrete instance of `TasksArrayInput` via:
//
//	TasksArray{ TasksArgs{...} }
type TasksArrayInput interface {
	pulumi.Input

	ToTasksArrayOutput() TasksArrayOutput
	ToTasksArrayOutputWithContext(context.Context) TasksArrayOutput
}

type TasksArray []TasksInput

func (TasksArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Tasks)(nil)).Elem()
}

func (i TasksArray) ToTasksArrayOutput() TasksArrayOutput {
	return i.ToTasksArrayOutputWithContext(context.Background())
}

func (i TasksArray) ToTasksArrayOutputWithContext(ctx context.Context) TasksArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TasksArrayOutput)
}

// TasksMapInput is an input type that accepts TasksMap and TasksMapOutput values.
// You can construct a concrete instance of `TasksMapInput` via:
//
//	TasksMap{ "key": TasksArgs{...} }
type TasksMapInput interface {
	pulumi.Input

	ToTasksMapOutput() TasksMapOutput
	ToTasksMapOutputWithContext(context.Context) TasksMapOutput
}

type TasksMap map[string]TasksInput

func (TasksMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Tasks)(nil)).Elem()
}

func (i TasksMap) ToTasksMapOutput() TasksMapOutput {
	return i.ToTasksMapOutputWithContext(context.Background())
}

func (i TasksMap) ToTasksMapOutputWithContext(ctx context.Context) TasksMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TasksMapOutput)
}

type TasksOutput struct{ *pulumi.OutputState }

func (TasksOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Tasks)(nil)).Elem()
}

func (o TasksOutput) ToTasksOutput() TasksOutput {
	return o
}

func (o TasksOutput) ToTasksOutputWithContext(ctx context.Context) TasksOutput {
	return o
}

// The frequency of the vulnerability scanning is done. Can be to **"hourly"**, **"daily"** or **"weekly"**
func (o TasksOutput) VulnerabilityScanPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *Tasks) pulumi.StringOutput { return v.VulnerabilityScanPolicy }).(pulumi.StringOutput)
}

type TasksArrayOutput struct{ *pulumi.OutputState }

func (TasksArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Tasks)(nil)).Elem()
}

func (o TasksArrayOutput) ToTasksArrayOutput() TasksArrayOutput {
	return o
}

func (o TasksArrayOutput) ToTasksArrayOutputWithContext(ctx context.Context) TasksArrayOutput {
	return o
}

func (o TasksArrayOutput) Index(i pulumi.IntInput) TasksOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Tasks {
		return vs[0].([]*Tasks)[vs[1].(int)]
	}).(TasksOutput)
}

type TasksMapOutput struct{ *pulumi.OutputState }

func (TasksMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Tasks)(nil)).Elem()
}

func (o TasksMapOutput) ToTasksMapOutput() TasksMapOutput {
	return o
}

func (o TasksMapOutput) ToTasksMapOutputWithContext(ctx context.Context) TasksMapOutput {
	return o
}

func (o TasksMapOutput) MapIndex(k pulumi.StringInput) TasksOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Tasks {
		return vs[0].(map[string]*Tasks)[vs[1].(string)]
	}).(TasksOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TasksInput)(nil)).Elem(), &Tasks{})
	pulumi.RegisterInputType(reflect.TypeOf((*TasksArrayInput)(nil)).Elem(), TasksArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TasksMapInput)(nil)).Elem(), TasksMap{})
	pulumi.RegisterOutputType(TasksOutput{})
	pulumi.RegisterOutputType(TasksArrayOutput{})
	pulumi.RegisterOutputType(TasksMapOutput{})
}
