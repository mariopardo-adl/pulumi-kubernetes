// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2alpha1

import (
	"context"
	"reflect"

	batchv1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/batch/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// CronJob represents the configuration of a single cron job.
type CronJobType struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMeta `pulumi:"metadata"`
	// Specification of the desired behavior of a cron job, including the schedule. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec CronJobSpec `pulumi:"spec"`
	// Current status of a cron job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status CronJobStatus `pulumi:"status"`
}

// CronJobTypeInput is an input type that accepts CronJobTypeArgs and CronJobTypeOutput values.
// You can construct a concrete instance of `CronJobTypeInput` via:
//
// 		 CronJobTypeArgs{...}
//
type CronJobTypeInput interface {
	pulumi.Input

	ToCronJobTypeOutput() CronJobTypeOutput
	ToCronJobTypeOutputWithContext(context.Context) CronJobTypeOutput
}

// CronJob represents the configuration of a single cron job.
type CronJobTypeArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringInput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringInput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaInput `pulumi:"metadata"`
	// Specification of the desired behavior of a cron job, including the schedule. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec CronJobSpecInput `pulumi:"spec"`
	// Current status of a cron job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Status CronJobStatusInput `pulumi:"status"`
}

func (CronJobTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CronJobType)(nil)).Elem()
}

func (i CronJobTypeArgs) ToCronJobTypeOutput() CronJobTypeOutput {
	return i.ToCronJobTypeOutputWithContext(context.Background())
}

func (i CronJobTypeArgs) ToCronJobTypeOutputWithContext(ctx context.Context) CronJobTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobTypeOutput)
}

// CronJobTypeArrayInput is an input type that accepts CronJobTypeArray and CronJobTypeArrayOutput values.
// You can construct a concrete instance of `CronJobTypeArrayInput` via:
//
// 		 CronJobTypeArray{ CronJobTypeArgs{...} }
//
type CronJobTypeArrayInput interface {
	pulumi.Input

	ToCronJobTypeArrayOutput() CronJobTypeArrayOutput
	ToCronJobTypeArrayOutputWithContext(context.Context) CronJobTypeArrayOutput
}

type CronJobTypeArray []CronJobTypeInput

func (CronJobTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CronJobType)(nil)).Elem()
}

func (i CronJobTypeArray) ToCronJobTypeArrayOutput() CronJobTypeArrayOutput {
	return i.ToCronJobTypeArrayOutputWithContext(context.Background())
}

func (i CronJobTypeArray) ToCronJobTypeArrayOutputWithContext(ctx context.Context) CronJobTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobTypeArrayOutput)
}

// CronJob represents the configuration of a single cron job.
type CronJobTypeOutput struct{ *pulumi.OutputState }

func (CronJobTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CronJobType)(nil)).Elem()
}

func (o CronJobTypeOutput) ToCronJobTypeOutput() CronJobTypeOutput {
	return o
}

func (o CronJobTypeOutput) ToCronJobTypeOutputWithContext(ctx context.Context) CronJobTypeOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o CronJobTypeOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v CronJobType) string { return v.ApiVersion }).(pulumi.StringOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o CronJobTypeOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v CronJobType) string { return v.Kind }).(pulumi.StringOutput)
}

// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o CronJobTypeOutput) Metadata() metav1.ObjectMetaOutput {
	return o.ApplyT(func(v CronJobType) metav1.ObjectMeta { return v.Metadata }).(metav1.ObjectMetaOutput)
}

// Specification of the desired behavior of a cron job, including the schedule. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o CronJobTypeOutput) Spec() CronJobSpecOutput {
	return o.ApplyT(func(v CronJobType) CronJobSpec { return v.Spec }).(CronJobSpecOutput)
}

// Current status of a cron job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o CronJobTypeOutput) Status() CronJobStatusOutput {
	return o.ApplyT(func(v CronJobType) CronJobStatus { return v.Status }).(CronJobStatusOutput)
}

type CronJobTypeArrayOutput struct{ *pulumi.OutputState }

func (CronJobTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CronJobType)(nil)).Elem()
}

func (o CronJobTypeArrayOutput) ToCronJobTypeArrayOutput() CronJobTypeArrayOutput {
	return o
}

func (o CronJobTypeArrayOutput) ToCronJobTypeArrayOutputWithContext(ctx context.Context) CronJobTypeArrayOutput {
	return o
}

func (o CronJobTypeArrayOutput) Index(i pulumi.IntInput) CronJobTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CronJobType {
		return vs[0].([]CronJobType)[vs[1].(int)]
	}).(CronJobTypeOutput)
}

// CronJobList is a collection of cron jobs.
type CronJobListType struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion string `pulumi:"apiVersion"`
	// items is the list of CronJobs.
	Items []CronJobType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMeta `pulumi:"metadata"`
}

// CronJobListTypeInput is an input type that accepts CronJobListTypeArgs and CronJobListTypeOutput values.
// You can construct a concrete instance of `CronJobListTypeInput` via:
//
// 		 CronJobListTypeArgs{...}
//
type CronJobListTypeInput interface {
	pulumi.Input

	ToCronJobListTypeOutput() CronJobListTypeOutput
	ToCronJobListTypeOutputWithContext(context.Context) CronJobListTypeOutput
}

// CronJobList is a collection of cron jobs.
type CronJobListTypeArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringInput `pulumi:"apiVersion"`
	// items is the list of CronJobs.
	Items CronJobTypeArrayInput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringInput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaInput `pulumi:"metadata"`
}

func (CronJobListTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CronJobListType)(nil)).Elem()
}

func (i CronJobListTypeArgs) ToCronJobListTypeOutput() CronJobListTypeOutput {
	return i.ToCronJobListTypeOutputWithContext(context.Background())
}

func (i CronJobListTypeArgs) ToCronJobListTypeOutputWithContext(ctx context.Context) CronJobListTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobListTypeOutput)
}

// CronJobList is a collection of cron jobs.
type CronJobListTypeOutput struct{ *pulumi.OutputState }

func (CronJobListTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CronJobListType)(nil)).Elem()
}

func (o CronJobListTypeOutput) ToCronJobListTypeOutput() CronJobListTypeOutput {
	return o
}

func (o CronJobListTypeOutput) ToCronJobListTypeOutputWithContext(ctx context.Context) CronJobListTypeOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o CronJobListTypeOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v CronJobListType) string { return v.ApiVersion }).(pulumi.StringOutput)
}

// items is the list of CronJobs.
func (o CronJobListTypeOutput) Items() CronJobTypeArrayOutput {
	return o.ApplyT(func(v CronJobListType) []CronJobType { return v.Items }).(CronJobTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o CronJobListTypeOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v CronJobListType) string { return v.Kind }).(pulumi.StringOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o CronJobListTypeOutput) Metadata() metav1.ListMetaOutput {
	return o.ApplyT(func(v CronJobListType) metav1.ListMeta { return v.Metadata }).(metav1.ListMetaOutput)
}

// CronJobSpec describes how the job execution will look like and when it will actually run.
type CronJobSpec struct {
	// Specifies how to treat concurrent executions of a Job. Valid values are: - "Allow" (default): allows CronJobs to run concurrently; - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - "Replace": cancels currently running job and replaces it with a new one
	ConcurrencyPolicy *string `pulumi:"concurrencyPolicy"`
	// The number of failed finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified.
	FailedJobsHistoryLimit *int `pulumi:"failedJobsHistoryLimit"`
	// Specifies the job that will be created when executing a CronJob.
	JobTemplate JobTemplateSpec `pulumi:"jobTemplate"`
	// The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
	Schedule string `pulumi:"schedule"`
	// Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones.
	StartingDeadlineSeconds *int `pulumi:"startingDeadlineSeconds"`
	// The number of successful finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified.
	SuccessfulJobsHistoryLimit *int `pulumi:"successfulJobsHistoryLimit"`
	// This flag tells the controller to suspend subsequent executions, it does not apply to already started executions.  Defaults to false.
	Suspend *bool `pulumi:"suspend"`
}

// CronJobSpecInput is an input type that accepts CronJobSpecArgs and CronJobSpecOutput values.
// You can construct a concrete instance of `CronJobSpecInput` via:
//
// 		 CronJobSpecArgs{...}
//
type CronJobSpecInput interface {
	pulumi.Input

	ToCronJobSpecOutput() CronJobSpecOutput
	ToCronJobSpecOutputWithContext(context.Context) CronJobSpecOutput
}

// CronJobSpec describes how the job execution will look like and when it will actually run.
type CronJobSpecArgs struct {
	// Specifies how to treat concurrent executions of a Job. Valid values are: - "Allow" (default): allows CronJobs to run concurrently; - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - "Replace": cancels currently running job and replaces it with a new one
	ConcurrencyPolicy pulumi.StringPtrInput `pulumi:"concurrencyPolicy"`
	// The number of failed finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified.
	FailedJobsHistoryLimit pulumi.IntPtrInput `pulumi:"failedJobsHistoryLimit"`
	// Specifies the job that will be created when executing a CronJob.
	JobTemplate JobTemplateSpecInput `pulumi:"jobTemplate"`
	// The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
	Schedule pulumi.StringInput `pulumi:"schedule"`
	// Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones.
	StartingDeadlineSeconds pulumi.IntPtrInput `pulumi:"startingDeadlineSeconds"`
	// The number of successful finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified.
	SuccessfulJobsHistoryLimit pulumi.IntPtrInput `pulumi:"successfulJobsHistoryLimit"`
	// This flag tells the controller to suspend subsequent executions, it does not apply to already started executions.  Defaults to false.
	Suspend pulumi.BoolPtrInput `pulumi:"suspend"`
}

func (CronJobSpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CronJobSpec)(nil)).Elem()
}

func (i CronJobSpecArgs) ToCronJobSpecOutput() CronJobSpecOutput {
	return i.ToCronJobSpecOutputWithContext(context.Background())
}

func (i CronJobSpecArgs) ToCronJobSpecOutputWithContext(ctx context.Context) CronJobSpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobSpecOutput)
}

func (i CronJobSpecArgs) ToCronJobSpecPtrOutput() CronJobSpecPtrOutput {
	return i.ToCronJobSpecPtrOutputWithContext(context.Background())
}

func (i CronJobSpecArgs) ToCronJobSpecPtrOutputWithContext(ctx context.Context) CronJobSpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobSpecOutput).ToCronJobSpecPtrOutputWithContext(ctx)
}

// CronJobSpecPtrInput is an input type that accepts CronJobSpecArgs, CronJobSpecPtr and CronJobSpecPtrOutput values.
// You can construct a concrete instance of `CronJobSpecPtrInput` via:
//
// 		 CronJobSpecArgs{...}
//
//  or:
//
// 		 nil
//
type CronJobSpecPtrInput interface {
	pulumi.Input

	ToCronJobSpecPtrOutput() CronJobSpecPtrOutput
	ToCronJobSpecPtrOutputWithContext(context.Context) CronJobSpecPtrOutput
}

type cronJobSpecPtrType CronJobSpecArgs

func CronJobSpecPtr(v *CronJobSpecArgs) CronJobSpecPtrInput {
	return (*cronJobSpecPtrType)(v)
}

func (*cronJobSpecPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CronJobSpec)(nil)).Elem()
}

func (i *cronJobSpecPtrType) ToCronJobSpecPtrOutput() CronJobSpecPtrOutput {
	return i.ToCronJobSpecPtrOutputWithContext(context.Background())
}

func (i *cronJobSpecPtrType) ToCronJobSpecPtrOutputWithContext(ctx context.Context) CronJobSpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobSpecPtrOutput)
}

// CronJobSpec describes how the job execution will look like and when it will actually run.
type CronJobSpecOutput struct{ *pulumi.OutputState }

func (CronJobSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CronJobSpec)(nil)).Elem()
}

func (o CronJobSpecOutput) ToCronJobSpecOutput() CronJobSpecOutput {
	return o
}

func (o CronJobSpecOutput) ToCronJobSpecOutputWithContext(ctx context.Context) CronJobSpecOutput {
	return o
}

func (o CronJobSpecOutput) ToCronJobSpecPtrOutput() CronJobSpecPtrOutput {
	return o.ToCronJobSpecPtrOutputWithContext(context.Background())
}

func (o CronJobSpecOutput) ToCronJobSpecPtrOutputWithContext(ctx context.Context) CronJobSpecPtrOutput {
	return o.ApplyT(func(v CronJobSpec) *CronJobSpec {
		return &v
	}).(CronJobSpecPtrOutput)
}

// Specifies how to treat concurrent executions of a Job. Valid values are: - "Allow" (default): allows CronJobs to run concurrently; - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - "Replace": cancels currently running job and replaces it with a new one
func (o CronJobSpecOutput) ConcurrencyPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CronJobSpec) *string { return v.ConcurrencyPolicy }).(pulumi.StringPtrOutput)
}

// The number of failed finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified.
func (o CronJobSpecOutput) FailedJobsHistoryLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CronJobSpec) *int { return v.FailedJobsHistoryLimit }).(pulumi.IntPtrOutput)
}

// Specifies the job that will be created when executing a CronJob.
func (o CronJobSpecOutput) JobTemplate() JobTemplateSpecOutput {
	return o.ApplyT(func(v CronJobSpec) JobTemplateSpec { return v.JobTemplate }).(JobTemplateSpecOutput)
}

// The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
func (o CronJobSpecOutput) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v CronJobSpec) string { return v.Schedule }).(pulumi.StringOutput)
}

// Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones.
func (o CronJobSpecOutput) StartingDeadlineSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CronJobSpec) *int { return v.StartingDeadlineSeconds }).(pulumi.IntPtrOutput)
}

// The number of successful finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified.
func (o CronJobSpecOutput) SuccessfulJobsHistoryLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CronJobSpec) *int { return v.SuccessfulJobsHistoryLimit }).(pulumi.IntPtrOutput)
}

// This flag tells the controller to suspend subsequent executions, it does not apply to already started executions.  Defaults to false.
func (o CronJobSpecOutput) Suspend() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CronJobSpec) *bool { return v.Suspend }).(pulumi.BoolPtrOutput)
}

type CronJobSpecPtrOutput struct{ *pulumi.OutputState }

func (CronJobSpecPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CronJobSpec)(nil)).Elem()
}

func (o CronJobSpecPtrOutput) ToCronJobSpecPtrOutput() CronJobSpecPtrOutput {
	return o
}

func (o CronJobSpecPtrOutput) ToCronJobSpecPtrOutputWithContext(ctx context.Context) CronJobSpecPtrOutput {
	return o
}

func (o CronJobSpecPtrOutput) Elem() CronJobSpecOutput {
	return o.ApplyT(func(v *CronJobSpec) CronJobSpec { return *v }).(CronJobSpecOutput)
}

// Specifies how to treat concurrent executions of a Job. Valid values are: - "Allow" (default): allows CronJobs to run concurrently; - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - "Replace": cancels currently running job and replaces it with a new one
func (o CronJobSpecPtrOutput) ConcurrencyPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CronJobSpec) *string {
		if v == nil {
			return nil
		}
		return v.ConcurrencyPolicy
	}).(pulumi.StringPtrOutput)
}

// The number of failed finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified.
func (o CronJobSpecPtrOutput) FailedJobsHistoryLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CronJobSpec) *int {
		if v == nil {
			return nil
		}
		return v.FailedJobsHistoryLimit
	}).(pulumi.IntPtrOutput)
}

// Specifies the job that will be created when executing a CronJob.
func (o CronJobSpecPtrOutput) JobTemplate() JobTemplateSpecPtrOutput {
	return o.ApplyT(func(v *CronJobSpec) *JobTemplateSpec {
		if v == nil {
			return nil
		}
		return &v.JobTemplate
	}).(JobTemplateSpecPtrOutput)
}

// The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
func (o CronJobSpecPtrOutput) Schedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CronJobSpec) *string {
		if v == nil {
			return nil
		}
		return &v.Schedule
	}).(pulumi.StringPtrOutput)
}

// Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones.
func (o CronJobSpecPtrOutput) StartingDeadlineSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CronJobSpec) *int {
		if v == nil {
			return nil
		}
		return v.StartingDeadlineSeconds
	}).(pulumi.IntPtrOutput)
}

// The number of successful finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified.
func (o CronJobSpecPtrOutput) SuccessfulJobsHistoryLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CronJobSpec) *int {
		if v == nil {
			return nil
		}
		return v.SuccessfulJobsHistoryLimit
	}).(pulumi.IntPtrOutput)
}

// This flag tells the controller to suspend subsequent executions, it does not apply to already started executions.  Defaults to false.
func (o CronJobSpecPtrOutput) Suspend() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CronJobSpec) *bool {
		if v == nil {
			return nil
		}
		return v.Suspend
	}).(pulumi.BoolPtrOutput)
}

// CronJobStatus represents the current state of a cron job.
type CronJobStatus struct {
	// A list of pointers to currently running jobs.
	Active []corev1.ObjectReference `pulumi:"active"`
	// Information when was the last time the job was successfully scheduled.
	LastScheduleTime *string `pulumi:"lastScheduleTime"`
}

// CronJobStatusInput is an input type that accepts CronJobStatusArgs and CronJobStatusOutput values.
// You can construct a concrete instance of `CronJobStatusInput` via:
//
// 		 CronJobStatusArgs{...}
//
type CronJobStatusInput interface {
	pulumi.Input

	ToCronJobStatusOutput() CronJobStatusOutput
	ToCronJobStatusOutputWithContext(context.Context) CronJobStatusOutput
}

// CronJobStatus represents the current state of a cron job.
type CronJobStatusArgs struct {
	// A list of pointers to currently running jobs.
	Active corev1.ObjectReferenceArrayInput `pulumi:"active"`
	// Information when was the last time the job was successfully scheduled.
	LastScheduleTime pulumi.StringPtrInput `pulumi:"lastScheduleTime"`
}

func (CronJobStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CronJobStatus)(nil)).Elem()
}

func (i CronJobStatusArgs) ToCronJobStatusOutput() CronJobStatusOutput {
	return i.ToCronJobStatusOutputWithContext(context.Background())
}

func (i CronJobStatusArgs) ToCronJobStatusOutputWithContext(ctx context.Context) CronJobStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobStatusOutput)
}

func (i CronJobStatusArgs) ToCronJobStatusPtrOutput() CronJobStatusPtrOutput {
	return i.ToCronJobStatusPtrOutputWithContext(context.Background())
}

func (i CronJobStatusArgs) ToCronJobStatusPtrOutputWithContext(ctx context.Context) CronJobStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobStatusOutput).ToCronJobStatusPtrOutputWithContext(ctx)
}

// CronJobStatusPtrInput is an input type that accepts CronJobStatusArgs, CronJobStatusPtr and CronJobStatusPtrOutput values.
// You can construct a concrete instance of `CronJobStatusPtrInput` via:
//
// 		 CronJobStatusArgs{...}
//
//  or:
//
// 		 nil
//
type CronJobStatusPtrInput interface {
	pulumi.Input

	ToCronJobStatusPtrOutput() CronJobStatusPtrOutput
	ToCronJobStatusPtrOutputWithContext(context.Context) CronJobStatusPtrOutput
}

type cronJobStatusPtrType CronJobStatusArgs

func CronJobStatusPtr(v *CronJobStatusArgs) CronJobStatusPtrInput {
	return (*cronJobStatusPtrType)(v)
}

func (*cronJobStatusPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CronJobStatus)(nil)).Elem()
}

func (i *cronJobStatusPtrType) ToCronJobStatusPtrOutput() CronJobStatusPtrOutput {
	return i.ToCronJobStatusPtrOutputWithContext(context.Background())
}

func (i *cronJobStatusPtrType) ToCronJobStatusPtrOutputWithContext(ctx context.Context) CronJobStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CronJobStatusPtrOutput)
}

// CronJobStatus represents the current state of a cron job.
type CronJobStatusOutput struct{ *pulumi.OutputState }

func (CronJobStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CronJobStatus)(nil)).Elem()
}

func (o CronJobStatusOutput) ToCronJobStatusOutput() CronJobStatusOutput {
	return o
}

func (o CronJobStatusOutput) ToCronJobStatusOutputWithContext(ctx context.Context) CronJobStatusOutput {
	return o
}

func (o CronJobStatusOutput) ToCronJobStatusPtrOutput() CronJobStatusPtrOutput {
	return o.ToCronJobStatusPtrOutputWithContext(context.Background())
}

func (o CronJobStatusOutput) ToCronJobStatusPtrOutputWithContext(ctx context.Context) CronJobStatusPtrOutput {
	return o.ApplyT(func(v CronJobStatus) *CronJobStatus {
		return &v
	}).(CronJobStatusPtrOutput)
}

// A list of pointers to currently running jobs.
func (o CronJobStatusOutput) Active() corev1.ObjectReferenceArrayOutput {
	return o.ApplyT(func(v CronJobStatus) []corev1.ObjectReference { return v.Active }).(corev1.ObjectReferenceArrayOutput)
}

// Information when was the last time the job was successfully scheduled.
func (o CronJobStatusOutput) LastScheduleTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CronJobStatus) *string { return v.LastScheduleTime }).(pulumi.StringPtrOutput)
}

type CronJobStatusPtrOutput struct{ *pulumi.OutputState }

func (CronJobStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CronJobStatus)(nil)).Elem()
}

func (o CronJobStatusPtrOutput) ToCronJobStatusPtrOutput() CronJobStatusPtrOutput {
	return o
}

func (o CronJobStatusPtrOutput) ToCronJobStatusPtrOutputWithContext(ctx context.Context) CronJobStatusPtrOutput {
	return o
}

func (o CronJobStatusPtrOutput) Elem() CronJobStatusOutput {
	return o.ApplyT(func(v *CronJobStatus) CronJobStatus { return *v }).(CronJobStatusOutput)
}

// A list of pointers to currently running jobs.
func (o CronJobStatusPtrOutput) Active() corev1.ObjectReferenceArrayOutput {
	return o.ApplyT(func(v *CronJobStatus) []corev1.ObjectReference {
		if v == nil {
			return nil
		}
		return v.Active
	}).(corev1.ObjectReferenceArrayOutput)
}

// Information when was the last time the job was successfully scheduled.
func (o CronJobStatusPtrOutput) LastScheduleTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CronJobStatus) *string {
		if v == nil {
			return nil
		}
		return v.LastScheduleTime
	}).(pulumi.StringPtrOutput)
}

// JobTemplateSpec describes the data a Job should have when created from a template
type JobTemplateSpec struct {
	// Standard object's metadata of the jobs created from this template. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Specification of the desired behavior of the job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *batchv1.JobSpec `pulumi:"spec"`
}

// JobTemplateSpecInput is an input type that accepts JobTemplateSpecArgs and JobTemplateSpecOutput values.
// You can construct a concrete instance of `JobTemplateSpecInput` via:
//
// 		 JobTemplateSpecArgs{...}
//
type JobTemplateSpecInput interface {
	pulumi.Input

	ToJobTemplateSpecOutput() JobTemplateSpecOutput
	ToJobTemplateSpecOutputWithContext(context.Context) JobTemplateSpecOutput
}

// JobTemplateSpec describes the data a Job should have when created from a template
type JobTemplateSpecArgs struct {
	// Standard object's metadata of the jobs created from this template. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput `pulumi:"metadata"`
	// Specification of the desired behavior of the job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec batchv1.JobSpecPtrInput `pulumi:"spec"`
}

func (JobTemplateSpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobTemplateSpec)(nil)).Elem()
}

func (i JobTemplateSpecArgs) ToJobTemplateSpecOutput() JobTemplateSpecOutput {
	return i.ToJobTemplateSpecOutputWithContext(context.Background())
}

func (i JobTemplateSpecArgs) ToJobTemplateSpecOutputWithContext(ctx context.Context) JobTemplateSpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobTemplateSpecOutput)
}

func (i JobTemplateSpecArgs) ToJobTemplateSpecPtrOutput() JobTemplateSpecPtrOutput {
	return i.ToJobTemplateSpecPtrOutputWithContext(context.Background())
}

func (i JobTemplateSpecArgs) ToJobTemplateSpecPtrOutputWithContext(ctx context.Context) JobTemplateSpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobTemplateSpecOutput).ToJobTemplateSpecPtrOutputWithContext(ctx)
}

// JobTemplateSpecPtrInput is an input type that accepts JobTemplateSpecArgs, JobTemplateSpecPtr and JobTemplateSpecPtrOutput values.
// You can construct a concrete instance of `JobTemplateSpecPtrInput` via:
//
// 		 JobTemplateSpecArgs{...}
//
//  or:
//
// 		 nil
//
type JobTemplateSpecPtrInput interface {
	pulumi.Input

	ToJobTemplateSpecPtrOutput() JobTemplateSpecPtrOutput
	ToJobTemplateSpecPtrOutputWithContext(context.Context) JobTemplateSpecPtrOutput
}

type jobTemplateSpecPtrType JobTemplateSpecArgs

func JobTemplateSpecPtr(v *JobTemplateSpecArgs) JobTemplateSpecPtrInput {
	return (*jobTemplateSpecPtrType)(v)
}

func (*jobTemplateSpecPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobTemplateSpec)(nil)).Elem()
}

func (i *jobTemplateSpecPtrType) ToJobTemplateSpecPtrOutput() JobTemplateSpecPtrOutput {
	return i.ToJobTemplateSpecPtrOutputWithContext(context.Background())
}

func (i *jobTemplateSpecPtrType) ToJobTemplateSpecPtrOutputWithContext(ctx context.Context) JobTemplateSpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobTemplateSpecPtrOutput)
}

// JobTemplateSpec describes the data a Job should have when created from a template
type JobTemplateSpecOutput struct{ *pulumi.OutputState }

func (JobTemplateSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobTemplateSpec)(nil)).Elem()
}

func (o JobTemplateSpecOutput) ToJobTemplateSpecOutput() JobTemplateSpecOutput {
	return o
}

func (o JobTemplateSpecOutput) ToJobTemplateSpecOutputWithContext(ctx context.Context) JobTemplateSpecOutput {
	return o
}

func (o JobTemplateSpecOutput) ToJobTemplateSpecPtrOutput() JobTemplateSpecPtrOutput {
	return o.ToJobTemplateSpecPtrOutputWithContext(context.Background())
}

func (o JobTemplateSpecOutput) ToJobTemplateSpecPtrOutputWithContext(ctx context.Context) JobTemplateSpecPtrOutput {
	return o.ApplyT(func(v JobTemplateSpec) *JobTemplateSpec {
		return &v
	}).(JobTemplateSpecPtrOutput)
}

// Standard object's metadata of the jobs created from this template. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o JobTemplateSpecOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v JobTemplateSpec) *metav1.ObjectMeta { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// Specification of the desired behavior of the job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o JobTemplateSpecOutput) Spec() batchv1.JobSpecPtrOutput {
	return o.ApplyT(func(v JobTemplateSpec) *batchv1.JobSpec { return v.Spec }).(batchv1.JobSpecPtrOutput)
}

type JobTemplateSpecPtrOutput struct{ *pulumi.OutputState }

func (JobTemplateSpecPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobTemplateSpec)(nil)).Elem()
}

func (o JobTemplateSpecPtrOutput) ToJobTemplateSpecPtrOutput() JobTemplateSpecPtrOutput {
	return o
}

func (o JobTemplateSpecPtrOutput) ToJobTemplateSpecPtrOutputWithContext(ctx context.Context) JobTemplateSpecPtrOutput {
	return o
}

func (o JobTemplateSpecPtrOutput) Elem() JobTemplateSpecOutput {
	return o.ApplyT(func(v *JobTemplateSpec) JobTemplateSpec { return *v }).(JobTemplateSpecOutput)
}

// Standard object's metadata of the jobs created from this template. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o JobTemplateSpecPtrOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v *JobTemplateSpec) *metav1.ObjectMeta {
		if v == nil {
			return nil
		}
		return v.Metadata
	}).(metav1.ObjectMetaPtrOutput)
}

// Specification of the desired behavior of the job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o JobTemplateSpecPtrOutput) Spec() batchv1.JobSpecPtrOutput {
	return o.ApplyT(func(v *JobTemplateSpec) *batchv1.JobSpec {
		if v == nil {
			return nil
		}
		return v.Spec
	}).(batchv1.JobSpecPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CronJobTypeOutput{})
	pulumi.RegisterOutputType(CronJobTypeArrayOutput{})
	pulumi.RegisterOutputType(CronJobListTypeOutput{})
	pulumi.RegisterOutputType(CronJobSpecOutput{})
	pulumi.RegisterOutputType(CronJobSpecPtrOutput{})
	pulumi.RegisterOutputType(CronJobStatusOutput{})
	pulumi.RegisterOutputType(CronJobStatusPtrOutput{})
	pulumi.RegisterOutputType(JobTemplateSpecOutput{})
	pulumi.RegisterOutputType(JobTemplateSpecPtrOutput{})
}
