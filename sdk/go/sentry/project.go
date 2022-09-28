// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sentry

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Project struct {
	pulumi.CustomResourceState

	Color pulumi.StringOutput `pulumi:"color"`
	// The maximum amount of time (in seconds) to wait between scheduling digests for delivery.
	DigestsMaxDelay pulumi.IntOutput `pulumi:"digestsMaxDelay"`
	// The minimum amount of time (in seconds) to wait between scheduling digests for delivery after the initial scheduling.
	DigestsMinDelay pulumi.IntOutput         `pulumi:"digestsMinDelay"`
	Features        pulumi.StringArrayOutput `pulumi:"features"`
	// The internal ID for this project.
	InternalId pulumi.StringOutput `pulumi:"internalId"`
	// Deprecated: is_bookmarked is no longer used
	IsBookmarked pulumi.BoolOutput `pulumi:"isBookmarked"`
	IsPublic     pulumi.BoolOutput `pulumi:"isPublic"`
	// The name for the project.
	Name pulumi.StringOutput `pulumi:"name"`
	// The slug of the organization the project belongs to.
	Organization pulumi.StringOutput `pulumi:"organization"`
	// The optional platform for this project.
	Platform pulumi.StringOutput `pulumi:"platform"`
	// Use `internal_id` instead.
	//
	// Deprecated: Use `internal_id` instead.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Hours in which an issue is automatically resolve if not seen after this amount of time.
	ResolveAge pulumi.IntOutput `pulumi:"resolveAge"`
	// The optional slug for this project.
	Slug   pulumi.StringOutput `pulumi:"slug"`
	Status pulumi.StringOutput `pulumi:"status"`
	// The slug of the team to create the project for.
	Team pulumi.StringOutput `pulumi:"team"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Organization == nil {
		return nil, errors.New("invalid value for required argument 'Organization'")
	}
	if args.Team == nil {
		return nil, errors.New("invalid value for required argument 'Team'")
	}
	var resource Project
	err := ctx.RegisterResource("sentry:index/project:Project", name, args, &resource, opts...)
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
	err := ctx.ReadResource("sentry:index/project:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
	Color *string `pulumi:"color"`
	// The maximum amount of time (in seconds) to wait between scheduling digests for delivery.
	DigestsMaxDelay *int `pulumi:"digestsMaxDelay"`
	// The minimum amount of time (in seconds) to wait between scheduling digests for delivery after the initial scheduling.
	DigestsMinDelay *int     `pulumi:"digestsMinDelay"`
	Features        []string `pulumi:"features"`
	// The internal ID for this project.
	InternalId *string `pulumi:"internalId"`
	// Deprecated: is_bookmarked is no longer used
	IsBookmarked *bool `pulumi:"isBookmarked"`
	IsPublic     *bool `pulumi:"isPublic"`
	// The name for the project.
	Name *string `pulumi:"name"`
	// The slug of the organization the project belongs to.
	Organization *string `pulumi:"organization"`
	// The optional platform for this project.
	Platform *string `pulumi:"platform"`
	// Use `internal_id` instead.
	//
	// Deprecated: Use `internal_id` instead.
	ProjectId *string `pulumi:"projectId"`
	// Hours in which an issue is automatically resolve if not seen after this amount of time.
	ResolveAge *int `pulumi:"resolveAge"`
	// The optional slug for this project.
	Slug   *string `pulumi:"slug"`
	Status *string `pulumi:"status"`
	// The slug of the team to create the project for.
	Team *string `pulumi:"team"`
}

type ProjectState struct {
	Color pulumi.StringPtrInput
	// The maximum amount of time (in seconds) to wait between scheduling digests for delivery.
	DigestsMaxDelay pulumi.IntPtrInput
	// The minimum amount of time (in seconds) to wait between scheduling digests for delivery after the initial scheduling.
	DigestsMinDelay pulumi.IntPtrInput
	Features        pulumi.StringArrayInput
	// The internal ID for this project.
	InternalId pulumi.StringPtrInput
	// Deprecated: is_bookmarked is no longer used
	IsBookmarked pulumi.BoolPtrInput
	IsPublic     pulumi.BoolPtrInput
	// The name for the project.
	Name pulumi.StringPtrInput
	// The slug of the organization the project belongs to.
	Organization pulumi.StringPtrInput
	// The optional platform for this project.
	Platform pulumi.StringPtrInput
	// Use `internal_id` instead.
	//
	// Deprecated: Use `internal_id` instead.
	ProjectId pulumi.StringPtrInput
	// Hours in which an issue is automatically resolve if not seen after this amount of time.
	ResolveAge pulumi.IntPtrInput
	// The optional slug for this project.
	Slug   pulumi.StringPtrInput
	Status pulumi.StringPtrInput
	// The slug of the team to create the project for.
	Team pulumi.StringPtrInput
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// The maximum amount of time (in seconds) to wait between scheduling digests for delivery.
	DigestsMaxDelay *int `pulumi:"digestsMaxDelay"`
	// The minimum amount of time (in seconds) to wait between scheduling digests for delivery after the initial scheduling.
	DigestsMinDelay *int `pulumi:"digestsMinDelay"`
	// The name for the project.
	Name *string `pulumi:"name"`
	// The slug of the organization the project belongs to.
	Organization string `pulumi:"organization"`
	// The optional platform for this project.
	Platform *string `pulumi:"platform"`
	// Hours in which an issue is automatically resolve if not seen after this amount of time.
	ResolveAge *int `pulumi:"resolveAge"`
	// The optional slug for this project.
	Slug *string `pulumi:"slug"`
	// The slug of the team to create the project for.
	Team string `pulumi:"team"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// The maximum amount of time (in seconds) to wait between scheduling digests for delivery.
	DigestsMaxDelay pulumi.IntPtrInput
	// The minimum amount of time (in seconds) to wait between scheduling digests for delivery after the initial scheduling.
	DigestsMinDelay pulumi.IntPtrInput
	// The name for the project.
	Name pulumi.StringPtrInput
	// The slug of the organization the project belongs to.
	Organization pulumi.StringInput
	// The optional platform for this project.
	Platform pulumi.StringPtrInput
	// Hours in which an issue is automatically resolve if not seen after this amount of time.
	ResolveAge pulumi.IntPtrInput
	// The optional slug for this project.
	Slug pulumi.StringPtrInput
	// The slug of the team to create the project for.
	Team pulumi.StringInput
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

func (o ProjectOutput) Color() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Color }).(pulumi.StringOutput)
}

// The maximum amount of time (in seconds) to wait between scheduling digests for delivery.
func (o ProjectOutput) DigestsMaxDelay() pulumi.IntOutput {
	return o.ApplyT(func(v *Project) pulumi.IntOutput { return v.DigestsMaxDelay }).(pulumi.IntOutput)
}

// The minimum amount of time (in seconds) to wait between scheduling digests for delivery after the initial scheduling.
func (o ProjectOutput) DigestsMinDelay() pulumi.IntOutput {
	return o.ApplyT(func(v *Project) pulumi.IntOutput { return v.DigestsMinDelay }).(pulumi.IntOutput)
}

func (o ProjectOutput) Features() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Project) pulumi.StringArrayOutput { return v.Features }).(pulumi.StringArrayOutput)
}

// The internal ID for this project.
func (o ProjectOutput) InternalId() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.InternalId }).(pulumi.StringOutput)
}

// Deprecated: is_bookmarked is no longer used
func (o ProjectOutput) IsBookmarked() pulumi.BoolOutput {
	return o.ApplyT(func(v *Project) pulumi.BoolOutput { return v.IsBookmarked }).(pulumi.BoolOutput)
}

func (o ProjectOutput) IsPublic() pulumi.BoolOutput {
	return o.ApplyT(func(v *Project) pulumi.BoolOutput { return v.IsPublic }).(pulumi.BoolOutput)
}

// The name for the project.
func (o ProjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The slug of the organization the project belongs to.
func (o ProjectOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Organization }).(pulumi.StringOutput)
}

// The optional platform for this project.
func (o ProjectOutput) Platform() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Platform }).(pulumi.StringOutput)
}

// Use `internal_id` instead.
//
// Deprecated: Use `internal_id` instead.
func (o ProjectOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Hours in which an issue is automatically resolve if not seen after this amount of time.
func (o ProjectOutput) ResolveAge() pulumi.IntOutput {
	return o.ApplyT(func(v *Project) pulumi.IntOutput { return v.ResolveAge }).(pulumi.IntOutput)
}

// The optional slug for this project.
func (o ProjectOutput) Slug() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Slug }).(pulumi.StringOutput)
}

func (o ProjectOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The slug of the team to create the project for.
func (o ProjectOutput) Team() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Team }).(pulumi.StringOutput)
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