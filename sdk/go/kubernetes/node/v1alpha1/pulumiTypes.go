// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Overhead structure represents the resource overhead associated with running a pod.
type Overhead struct {
	// PodFixed represents the fixed resource overhead associated with running a pod.
	PodFixed map[string]string `pulumi:"podFixed"`
}

// OverheadInput is an input type that accepts OverheadArgs and OverheadOutput values.
// You can construct a concrete instance of `OverheadInput` via:
//
//          OverheadArgs{...}
type OverheadInput interface {
	pulumi.Input

	ToOverheadOutput() OverheadOutput
	ToOverheadOutputWithContext(context.Context) OverheadOutput
}

// Overhead structure represents the resource overhead associated with running a pod.
type OverheadArgs struct {
	// PodFixed represents the fixed resource overhead associated with running a pod.
	PodFixed pulumi.StringMapInput `pulumi:"podFixed"`
}

func (OverheadArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Overhead)(nil)).Elem()
}

func (i OverheadArgs) ToOverheadOutput() OverheadOutput {
	return i.ToOverheadOutputWithContext(context.Background())
}

func (i OverheadArgs) ToOverheadOutputWithContext(ctx context.Context) OverheadOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OverheadOutput)
}

func (i OverheadArgs) ToOverheadPtrOutput() OverheadPtrOutput {
	return i.ToOverheadPtrOutputWithContext(context.Background())
}

func (i OverheadArgs) ToOverheadPtrOutputWithContext(ctx context.Context) OverheadPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OverheadOutput).ToOverheadPtrOutputWithContext(ctx)
}

// OverheadPtrInput is an input type that accepts OverheadArgs, OverheadPtr and OverheadPtrOutput values.
// You can construct a concrete instance of `OverheadPtrInput` via:
//
//          OverheadArgs{...}
//
//  or:
//
//          nil
type OverheadPtrInput interface {
	pulumi.Input

	ToOverheadPtrOutput() OverheadPtrOutput
	ToOverheadPtrOutputWithContext(context.Context) OverheadPtrOutput
}

type overheadPtrType OverheadArgs

func OverheadPtr(v *OverheadArgs) OverheadPtrInput {
	return (*overheadPtrType)(v)
}

func (*overheadPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Overhead)(nil)).Elem()
}

func (i *overheadPtrType) ToOverheadPtrOutput() OverheadPtrOutput {
	return i.ToOverheadPtrOutputWithContext(context.Background())
}

func (i *overheadPtrType) ToOverheadPtrOutputWithContext(ctx context.Context) OverheadPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OverheadPtrOutput)
}

// Overhead structure represents the resource overhead associated with running a pod.
type OverheadOutput struct{ *pulumi.OutputState }

func (OverheadOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Overhead)(nil)).Elem()
}

func (o OverheadOutput) ToOverheadOutput() OverheadOutput {
	return o
}

func (o OverheadOutput) ToOverheadOutputWithContext(ctx context.Context) OverheadOutput {
	return o
}

func (o OverheadOutput) ToOverheadPtrOutput() OverheadPtrOutput {
	return o.ToOverheadPtrOutputWithContext(context.Background())
}

func (o OverheadOutput) ToOverheadPtrOutputWithContext(ctx context.Context) OverheadPtrOutput {
	return o.ApplyT(func(v Overhead) *Overhead {
		return &v
	}).(OverheadPtrOutput)
}

// PodFixed represents the fixed resource overhead associated with running a pod.
func (o OverheadOutput) PodFixed() pulumi.StringMapOutput {
	return o.ApplyT(func(v Overhead) map[string]string { return v.PodFixed }).(pulumi.StringMapOutput)
}

type OverheadPtrOutput struct{ *pulumi.OutputState }

func (OverheadPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Overhead)(nil)).Elem()
}

func (o OverheadPtrOutput) ToOverheadPtrOutput() OverheadPtrOutput {
	return o
}

func (o OverheadPtrOutput) ToOverheadPtrOutputWithContext(ctx context.Context) OverheadPtrOutput {
	return o
}

func (o OverheadPtrOutput) Elem() OverheadOutput {
	return o.ApplyT(func(v *Overhead) Overhead { return *v }).(OverheadOutput)
}

// PodFixed represents the fixed resource overhead associated with running a pod.
func (o OverheadPtrOutput) PodFixed() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Overhead) map[string]string {
		if v == nil {
			return nil
		}
		return v.PodFixed
	}).(pulumi.StringMapOutput)
}

// RuntimeClass defines a class of container runtime supported in the cluster. The RuntimeClass is used to determine which container runtime is used to run all containers in a pod. RuntimeClasses are (currently) manually defined by a user or cluster provisioner, and referenced in the PodSpec. The Kubelet is responsible for resolving the RuntimeClassName reference before running the pod.  For more details, see https://git.k8s.io/enhancements/keps/sig-node/runtime-class.md
type RuntimeClassType struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Specification of the RuntimeClass More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec RuntimeClassSpec `pulumi:"spec"`
}

// RuntimeClassTypeInput is an input type that accepts RuntimeClassTypeArgs and RuntimeClassTypeOutput values.
// You can construct a concrete instance of `RuntimeClassTypeInput` via:
//
//          RuntimeClassTypeArgs{...}
type RuntimeClassTypeInput interface {
	pulumi.Input

	ToRuntimeClassTypeOutput() RuntimeClassTypeOutput
	ToRuntimeClassTypeOutputWithContext(context.Context) RuntimeClassTypeOutput
}

// RuntimeClass defines a class of container runtime supported in the cluster. The RuntimeClass is used to determine which container runtime is used to run all containers in a pod. RuntimeClasses are (currently) manually defined by a user or cluster provisioner, and referenced in the PodSpec. The Kubelet is responsible for resolving the RuntimeClassName reference before running the pod.  For more details, see https://git.k8s.io/enhancements/keps/sig-node/runtime-class.md
type RuntimeClassTypeArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput `pulumi:"kind"`
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput `pulumi:"metadata"`
	// Specification of the RuntimeClass More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec RuntimeClassSpecInput `pulumi:"spec"`
}

func (RuntimeClassTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuntimeClassType)(nil)).Elem()
}

func (i RuntimeClassTypeArgs) ToRuntimeClassTypeOutput() RuntimeClassTypeOutput {
	return i.ToRuntimeClassTypeOutputWithContext(context.Background())
}

func (i RuntimeClassTypeArgs) ToRuntimeClassTypeOutputWithContext(ctx context.Context) RuntimeClassTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeClassTypeOutput)
}

// RuntimeClassTypeArrayInput is an input type that accepts RuntimeClassTypeArray and RuntimeClassTypeArrayOutput values.
// You can construct a concrete instance of `RuntimeClassTypeArrayInput` via:
//
//          RuntimeClassTypeArray{ RuntimeClassTypeArgs{...} }
type RuntimeClassTypeArrayInput interface {
	pulumi.Input

	ToRuntimeClassTypeArrayOutput() RuntimeClassTypeArrayOutput
	ToRuntimeClassTypeArrayOutputWithContext(context.Context) RuntimeClassTypeArrayOutput
}

type RuntimeClassTypeArray []RuntimeClassTypeInput

func (RuntimeClassTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuntimeClassType)(nil)).Elem()
}

func (i RuntimeClassTypeArray) ToRuntimeClassTypeArrayOutput() RuntimeClassTypeArrayOutput {
	return i.ToRuntimeClassTypeArrayOutputWithContext(context.Background())
}

func (i RuntimeClassTypeArray) ToRuntimeClassTypeArrayOutputWithContext(ctx context.Context) RuntimeClassTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeClassTypeArrayOutput)
}

// RuntimeClass defines a class of container runtime supported in the cluster. The RuntimeClass is used to determine which container runtime is used to run all containers in a pod. RuntimeClasses are (currently) manually defined by a user or cluster provisioner, and referenced in the PodSpec. The Kubelet is responsible for resolving the RuntimeClassName reference before running the pod.  For more details, see https://git.k8s.io/enhancements/keps/sig-node/runtime-class.md
type RuntimeClassTypeOutput struct{ *pulumi.OutputState }

func (RuntimeClassTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuntimeClassType)(nil)).Elem()
}

func (o RuntimeClassTypeOutput) ToRuntimeClassTypeOutput() RuntimeClassTypeOutput {
	return o
}

func (o RuntimeClassTypeOutput) ToRuntimeClassTypeOutputWithContext(ctx context.Context) RuntimeClassTypeOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o RuntimeClassTypeOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuntimeClassType) *string { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o RuntimeClassTypeOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuntimeClassType) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o RuntimeClassTypeOutput) Metadata() metav1.ObjectMetaPtrOutput {
	return o.ApplyT(func(v RuntimeClassType) *metav1.ObjectMeta { return v.Metadata }).(metav1.ObjectMetaPtrOutput)
}

// Specification of the RuntimeClass More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
func (o RuntimeClassTypeOutput) Spec() RuntimeClassSpecOutput {
	return o.ApplyT(func(v RuntimeClassType) RuntimeClassSpec { return v.Spec }).(RuntimeClassSpecOutput)
}

type RuntimeClassTypeArrayOutput struct{ *pulumi.OutputState }

func (RuntimeClassTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuntimeClassType)(nil)).Elem()
}

func (o RuntimeClassTypeArrayOutput) ToRuntimeClassTypeArrayOutput() RuntimeClassTypeArrayOutput {
	return o
}

func (o RuntimeClassTypeArrayOutput) ToRuntimeClassTypeArrayOutputWithContext(ctx context.Context) RuntimeClassTypeArrayOutput {
	return o
}

func (o RuntimeClassTypeArrayOutput) Index(i pulumi.IntInput) RuntimeClassTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RuntimeClassType {
		return vs[0].([]RuntimeClassType)[vs[1].(int)]
	}).(RuntimeClassTypeOutput)
}

// RuntimeClassList is a list of RuntimeClass objects.
type RuntimeClassListType struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is a list of schema objects.
	Items []RuntimeClassType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// RuntimeClassListTypeInput is an input type that accepts RuntimeClassListTypeArgs and RuntimeClassListTypeOutput values.
// You can construct a concrete instance of `RuntimeClassListTypeInput` via:
//
//          RuntimeClassListTypeArgs{...}
type RuntimeClassListTypeInput interface {
	pulumi.Input

	ToRuntimeClassListTypeOutput() RuntimeClassListTypeOutput
	ToRuntimeClassListTypeOutputWithContext(context.Context) RuntimeClassListTypeOutput
}

// RuntimeClassList is a list of RuntimeClass objects.
type RuntimeClassListTypeArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput `pulumi:"apiVersion"`
	// Items is a list of schema objects.
	Items RuntimeClassTypeArrayInput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput `pulumi:"metadata"`
}

func (RuntimeClassListTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuntimeClassListType)(nil)).Elem()
}

func (i RuntimeClassListTypeArgs) ToRuntimeClassListTypeOutput() RuntimeClassListTypeOutput {
	return i.ToRuntimeClassListTypeOutputWithContext(context.Background())
}

func (i RuntimeClassListTypeArgs) ToRuntimeClassListTypeOutputWithContext(ctx context.Context) RuntimeClassListTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeClassListTypeOutput)
}

// RuntimeClassList is a list of RuntimeClass objects.
type RuntimeClassListTypeOutput struct{ *pulumi.OutputState }

func (RuntimeClassListTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuntimeClassListType)(nil)).Elem()
}

func (o RuntimeClassListTypeOutput) ToRuntimeClassListTypeOutput() RuntimeClassListTypeOutput {
	return o
}

func (o RuntimeClassListTypeOutput) ToRuntimeClassListTypeOutputWithContext(ctx context.Context) RuntimeClassListTypeOutput {
	return o
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (o RuntimeClassListTypeOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuntimeClassListType) *string { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

// Items is a list of schema objects.
func (o RuntimeClassListTypeOutput) Items() RuntimeClassTypeArrayOutput {
	return o.ApplyT(func(v RuntimeClassListType) []RuntimeClassType { return v.Items }).(RuntimeClassTypeArrayOutput)
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (o RuntimeClassListTypeOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuntimeClassListType) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
func (o RuntimeClassListTypeOutput) Metadata() metav1.ListMetaPtrOutput {
	return o.ApplyT(func(v RuntimeClassListType) *metav1.ListMeta { return v.Metadata }).(metav1.ListMetaPtrOutput)
}

// RuntimeClassSpec is a specification of a RuntimeClass. It contains parameters that are required to describe the RuntimeClass to the Container Runtime Interface (CRI) implementation, as well as any other components that need to understand how the pod will be run. The RuntimeClassSpec is immutable.
type RuntimeClassSpec struct {
	// Overhead represents the resource overhead associated with running a pod for a given RuntimeClass. For more details, see https://git.k8s.io/enhancements/keps/sig-node/20190226-pod-overhead.md This field is alpha-level as of Kubernetes v1.15, and is only honored by servers that enable the PodOverhead feature.
	Overhead *Overhead `pulumi:"overhead"`
	// RuntimeHandler specifies the underlying runtime and configuration that the CRI implementation will use to handle pods of this class. The possible values are specific to the node & CRI configuration.  It is assumed that all handlers are available on every node, and handlers of the same name are equivalent on every node. For example, a handler called "runc" might specify that the runc OCI runtime (using native Linux containers) will be used to run the containers in a pod. The RuntimeHandler must be lowercase, conform to the DNS Label (RFC 1123) requirements, and is immutable.
	RuntimeHandler string `pulumi:"runtimeHandler"`
	// Scheduling holds the scheduling constraints to ensure that pods running with this RuntimeClass are scheduled to nodes that support it. If scheduling is nil, this RuntimeClass is assumed to be supported by all nodes.
	Scheduling *Scheduling `pulumi:"scheduling"`
}

// RuntimeClassSpecInput is an input type that accepts RuntimeClassSpecArgs and RuntimeClassSpecOutput values.
// You can construct a concrete instance of `RuntimeClassSpecInput` via:
//
//          RuntimeClassSpecArgs{...}
type RuntimeClassSpecInput interface {
	pulumi.Input

	ToRuntimeClassSpecOutput() RuntimeClassSpecOutput
	ToRuntimeClassSpecOutputWithContext(context.Context) RuntimeClassSpecOutput
}

// RuntimeClassSpec is a specification of a RuntimeClass. It contains parameters that are required to describe the RuntimeClass to the Container Runtime Interface (CRI) implementation, as well as any other components that need to understand how the pod will be run. The RuntimeClassSpec is immutable.
type RuntimeClassSpecArgs struct {
	// Overhead represents the resource overhead associated with running a pod for a given RuntimeClass. For more details, see https://git.k8s.io/enhancements/keps/sig-node/20190226-pod-overhead.md This field is alpha-level as of Kubernetes v1.15, and is only honored by servers that enable the PodOverhead feature.
	Overhead OverheadPtrInput `pulumi:"overhead"`
	// RuntimeHandler specifies the underlying runtime and configuration that the CRI implementation will use to handle pods of this class. The possible values are specific to the node & CRI configuration.  It is assumed that all handlers are available on every node, and handlers of the same name are equivalent on every node. For example, a handler called "runc" might specify that the runc OCI runtime (using native Linux containers) will be used to run the containers in a pod. The RuntimeHandler must be lowercase, conform to the DNS Label (RFC 1123) requirements, and is immutable.
	RuntimeHandler pulumi.StringInput `pulumi:"runtimeHandler"`
	// Scheduling holds the scheduling constraints to ensure that pods running with this RuntimeClass are scheduled to nodes that support it. If scheduling is nil, this RuntimeClass is assumed to be supported by all nodes.
	Scheduling SchedulingPtrInput `pulumi:"scheduling"`
}

func (RuntimeClassSpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuntimeClassSpec)(nil)).Elem()
}

func (i RuntimeClassSpecArgs) ToRuntimeClassSpecOutput() RuntimeClassSpecOutput {
	return i.ToRuntimeClassSpecOutputWithContext(context.Background())
}

func (i RuntimeClassSpecArgs) ToRuntimeClassSpecOutputWithContext(ctx context.Context) RuntimeClassSpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeClassSpecOutput)
}

func (i RuntimeClassSpecArgs) ToRuntimeClassSpecPtrOutput() RuntimeClassSpecPtrOutput {
	return i.ToRuntimeClassSpecPtrOutputWithContext(context.Background())
}

func (i RuntimeClassSpecArgs) ToRuntimeClassSpecPtrOutputWithContext(ctx context.Context) RuntimeClassSpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeClassSpecOutput).ToRuntimeClassSpecPtrOutputWithContext(ctx)
}

// RuntimeClassSpecPtrInput is an input type that accepts RuntimeClassSpecArgs, RuntimeClassSpecPtr and RuntimeClassSpecPtrOutput values.
// You can construct a concrete instance of `RuntimeClassSpecPtrInput` via:
//
//          RuntimeClassSpecArgs{...}
//
//  or:
//
//          nil
type RuntimeClassSpecPtrInput interface {
	pulumi.Input

	ToRuntimeClassSpecPtrOutput() RuntimeClassSpecPtrOutput
	ToRuntimeClassSpecPtrOutputWithContext(context.Context) RuntimeClassSpecPtrOutput
}

type runtimeClassSpecPtrType RuntimeClassSpecArgs

func RuntimeClassSpecPtr(v *RuntimeClassSpecArgs) RuntimeClassSpecPtrInput {
	return (*runtimeClassSpecPtrType)(v)
}

func (*runtimeClassSpecPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RuntimeClassSpec)(nil)).Elem()
}

func (i *runtimeClassSpecPtrType) ToRuntimeClassSpecPtrOutput() RuntimeClassSpecPtrOutput {
	return i.ToRuntimeClassSpecPtrOutputWithContext(context.Background())
}

func (i *runtimeClassSpecPtrType) ToRuntimeClassSpecPtrOutputWithContext(ctx context.Context) RuntimeClassSpecPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeClassSpecPtrOutput)
}

// RuntimeClassSpec is a specification of a RuntimeClass. It contains parameters that are required to describe the RuntimeClass to the Container Runtime Interface (CRI) implementation, as well as any other components that need to understand how the pod will be run. The RuntimeClassSpec is immutable.
type RuntimeClassSpecOutput struct{ *pulumi.OutputState }

func (RuntimeClassSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuntimeClassSpec)(nil)).Elem()
}

func (o RuntimeClassSpecOutput) ToRuntimeClassSpecOutput() RuntimeClassSpecOutput {
	return o
}

func (o RuntimeClassSpecOutput) ToRuntimeClassSpecOutputWithContext(ctx context.Context) RuntimeClassSpecOutput {
	return o
}

func (o RuntimeClassSpecOutput) ToRuntimeClassSpecPtrOutput() RuntimeClassSpecPtrOutput {
	return o.ToRuntimeClassSpecPtrOutputWithContext(context.Background())
}

func (o RuntimeClassSpecOutput) ToRuntimeClassSpecPtrOutputWithContext(ctx context.Context) RuntimeClassSpecPtrOutput {
	return o.ApplyT(func(v RuntimeClassSpec) *RuntimeClassSpec {
		return &v
	}).(RuntimeClassSpecPtrOutput)
}

// Overhead represents the resource overhead associated with running a pod for a given RuntimeClass. For more details, see https://git.k8s.io/enhancements/keps/sig-node/20190226-pod-overhead.md This field is alpha-level as of Kubernetes v1.15, and is only honored by servers that enable the PodOverhead feature.
func (o RuntimeClassSpecOutput) Overhead() OverheadPtrOutput {
	return o.ApplyT(func(v RuntimeClassSpec) *Overhead { return v.Overhead }).(OverheadPtrOutput)
}

// RuntimeHandler specifies the underlying runtime and configuration that the CRI implementation will use to handle pods of this class. The possible values are specific to the node & CRI configuration.  It is assumed that all handlers are available on every node, and handlers of the same name are equivalent on every node. For example, a handler called "runc" might specify that the runc OCI runtime (using native Linux containers) will be used to run the containers in a pod. The RuntimeHandler must be lowercase, conform to the DNS Label (RFC 1123) requirements, and is immutable.
func (o RuntimeClassSpecOutput) RuntimeHandler() pulumi.StringOutput {
	return o.ApplyT(func(v RuntimeClassSpec) string { return v.RuntimeHandler }).(pulumi.StringOutput)
}

// Scheduling holds the scheduling constraints to ensure that pods running with this RuntimeClass are scheduled to nodes that support it. If scheduling is nil, this RuntimeClass is assumed to be supported by all nodes.
func (o RuntimeClassSpecOutput) Scheduling() SchedulingPtrOutput {
	return o.ApplyT(func(v RuntimeClassSpec) *Scheduling { return v.Scheduling }).(SchedulingPtrOutput)
}

type RuntimeClassSpecPtrOutput struct{ *pulumi.OutputState }

func (RuntimeClassSpecPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuntimeClassSpec)(nil)).Elem()
}

func (o RuntimeClassSpecPtrOutput) ToRuntimeClassSpecPtrOutput() RuntimeClassSpecPtrOutput {
	return o
}

func (o RuntimeClassSpecPtrOutput) ToRuntimeClassSpecPtrOutputWithContext(ctx context.Context) RuntimeClassSpecPtrOutput {
	return o
}

func (o RuntimeClassSpecPtrOutput) Elem() RuntimeClassSpecOutput {
	return o.ApplyT(func(v *RuntimeClassSpec) RuntimeClassSpec { return *v }).(RuntimeClassSpecOutput)
}

// Overhead represents the resource overhead associated with running a pod for a given RuntimeClass. For more details, see https://git.k8s.io/enhancements/keps/sig-node/20190226-pod-overhead.md This field is alpha-level as of Kubernetes v1.15, and is only honored by servers that enable the PodOverhead feature.
func (o RuntimeClassSpecPtrOutput) Overhead() OverheadPtrOutput {
	return o.ApplyT(func(v *RuntimeClassSpec) *Overhead {
		if v == nil {
			return nil
		}
		return v.Overhead
	}).(OverheadPtrOutput)
}

// RuntimeHandler specifies the underlying runtime and configuration that the CRI implementation will use to handle pods of this class. The possible values are specific to the node & CRI configuration.  It is assumed that all handlers are available on every node, and handlers of the same name are equivalent on every node. For example, a handler called "runc" might specify that the runc OCI runtime (using native Linux containers) will be used to run the containers in a pod. The RuntimeHandler must be lowercase, conform to the DNS Label (RFC 1123) requirements, and is immutable.
func (o RuntimeClassSpecPtrOutput) RuntimeHandler() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuntimeClassSpec) *string {
		if v == nil {
			return nil
		}
		return &v.RuntimeHandler
	}).(pulumi.StringPtrOutput)
}

// Scheduling holds the scheduling constraints to ensure that pods running with this RuntimeClass are scheduled to nodes that support it. If scheduling is nil, this RuntimeClass is assumed to be supported by all nodes.
func (o RuntimeClassSpecPtrOutput) Scheduling() SchedulingPtrOutput {
	return o.ApplyT(func(v *RuntimeClassSpec) *Scheduling {
		if v == nil {
			return nil
		}
		return v.Scheduling
	}).(SchedulingPtrOutput)
}

// Scheduling specifies the scheduling constraints for nodes supporting a RuntimeClass.
type Scheduling struct {
	// nodeSelector lists labels that must be present on nodes that support this RuntimeClass. Pods using this RuntimeClass can only be scheduled to a node matched by this selector. The RuntimeClass nodeSelector is merged with a pod's existing nodeSelector. Any conflicts will cause the pod to be rejected in admission.
	NodeSelector map[string]string `pulumi:"nodeSelector"`
	// tolerations are appended (excluding duplicates) to pods running with this RuntimeClass during admission, effectively unioning the set of nodes tolerated by the pod and the RuntimeClass.
	Tolerations []corev1.Toleration `pulumi:"tolerations"`
}

// SchedulingInput is an input type that accepts SchedulingArgs and SchedulingOutput values.
// You can construct a concrete instance of `SchedulingInput` via:
//
//          SchedulingArgs{...}
type SchedulingInput interface {
	pulumi.Input

	ToSchedulingOutput() SchedulingOutput
	ToSchedulingOutputWithContext(context.Context) SchedulingOutput
}

// Scheduling specifies the scheduling constraints for nodes supporting a RuntimeClass.
type SchedulingArgs struct {
	// nodeSelector lists labels that must be present on nodes that support this RuntimeClass. Pods using this RuntimeClass can only be scheduled to a node matched by this selector. The RuntimeClass nodeSelector is merged with a pod's existing nodeSelector. Any conflicts will cause the pod to be rejected in admission.
	NodeSelector pulumi.StringMapInput `pulumi:"nodeSelector"`
	// tolerations are appended (excluding duplicates) to pods running with this RuntimeClass during admission, effectively unioning the set of nodes tolerated by the pod and the RuntimeClass.
	Tolerations corev1.TolerationArrayInput `pulumi:"tolerations"`
}

func (SchedulingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Scheduling)(nil)).Elem()
}

func (i SchedulingArgs) ToSchedulingOutput() SchedulingOutput {
	return i.ToSchedulingOutputWithContext(context.Background())
}

func (i SchedulingArgs) ToSchedulingOutputWithContext(ctx context.Context) SchedulingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchedulingOutput)
}

func (i SchedulingArgs) ToSchedulingPtrOutput() SchedulingPtrOutput {
	return i.ToSchedulingPtrOutputWithContext(context.Background())
}

func (i SchedulingArgs) ToSchedulingPtrOutputWithContext(ctx context.Context) SchedulingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchedulingOutput).ToSchedulingPtrOutputWithContext(ctx)
}

// SchedulingPtrInput is an input type that accepts SchedulingArgs, SchedulingPtr and SchedulingPtrOutput values.
// You can construct a concrete instance of `SchedulingPtrInput` via:
//
//          SchedulingArgs{...}
//
//  or:
//
//          nil
type SchedulingPtrInput interface {
	pulumi.Input

	ToSchedulingPtrOutput() SchedulingPtrOutput
	ToSchedulingPtrOutputWithContext(context.Context) SchedulingPtrOutput
}

type schedulingPtrType SchedulingArgs

func SchedulingPtr(v *SchedulingArgs) SchedulingPtrInput {
	return (*schedulingPtrType)(v)
}

func (*schedulingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Scheduling)(nil)).Elem()
}

func (i *schedulingPtrType) ToSchedulingPtrOutput() SchedulingPtrOutput {
	return i.ToSchedulingPtrOutputWithContext(context.Background())
}

func (i *schedulingPtrType) ToSchedulingPtrOutputWithContext(ctx context.Context) SchedulingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchedulingPtrOutput)
}

// Scheduling specifies the scheduling constraints for nodes supporting a RuntimeClass.
type SchedulingOutput struct{ *pulumi.OutputState }

func (SchedulingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Scheduling)(nil)).Elem()
}

func (o SchedulingOutput) ToSchedulingOutput() SchedulingOutput {
	return o
}

func (o SchedulingOutput) ToSchedulingOutputWithContext(ctx context.Context) SchedulingOutput {
	return o
}

func (o SchedulingOutput) ToSchedulingPtrOutput() SchedulingPtrOutput {
	return o.ToSchedulingPtrOutputWithContext(context.Background())
}

func (o SchedulingOutput) ToSchedulingPtrOutputWithContext(ctx context.Context) SchedulingPtrOutput {
	return o.ApplyT(func(v Scheduling) *Scheduling {
		return &v
	}).(SchedulingPtrOutput)
}

// nodeSelector lists labels that must be present on nodes that support this RuntimeClass. Pods using this RuntimeClass can only be scheduled to a node matched by this selector. The RuntimeClass nodeSelector is merged with a pod's existing nodeSelector. Any conflicts will cause the pod to be rejected in admission.
func (o SchedulingOutput) NodeSelector() pulumi.StringMapOutput {
	return o.ApplyT(func(v Scheduling) map[string]string { return v.NodeSelector }).(pulumi.StringMapOutput)
}

// tolerations are appended (excluding duplicates) to pods running with this RuntimeClass during admission, effectively unioning the set of nodes tolerated by the pod and the RuntimeClass.
func (o SchedulingOutput) Tolerations() corev1.TolerationArrayOutput {
	return o.ApplyT(func(v Scheduling) []corev1.Toleration { return v.Tolerations }).(corev1.TolerationArrayOutput)
}

type SchedulingPtrOutput struct{ *pulumi.OutputState }

func (SchedulingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Scheduling)(nil)).Elem()
}

func (o SchedulingPtrOutput) ToSchedulingPtrOutput() SchedulingPtrOutput {
	return o
}

func (o SchedulingPtrOutput) ToSchedulingPtrOutputWithContext(ctx context.Context) SchedulingPtrOutput {
	return o
}

func (o SchedulingPtrOutput) Elem() SchedulingOutput {
	return o.ApplyT(func(v *Scheduling) Scheduling { return *v }).(SchedulingOutput)
}

// nodeSelector lists labels that must be present on nodes that support this RuntimeClass. Pods using this RuntimeClass can only be scheduled to a node matched by this selector. The RuntimeClass nodeSelector is merged with a pod's existing nodeSelector. Any conflicts will cause the pod to be rejected in admission.
func (o SchedulingPtrOutput) NodeSelector() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Scheduling) map[string]string {
		if v == nil {
			return nil
		}
		return v.NodeSelector
	}).(pulumi.StringMapOutput)
}

// tolerations are appended (excluding duplicates) to pods running with this RuntimeClass during admission, effectively unioning the set of nodes tolerated by the pod and the RuntimeClass.
func (o SchedulingPtrOutput) Tolerations() corev1.TolerationArrayOutput {
	return o.ApplyT(func(v *Scheduling) []corev1.Toleration {
		if v == nil {
			return nil
		}
		return v.Tolerations
	}).(corev1.TolerationArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(OverheadOutput{})
	pulumi.RegisterOutputType(OverheadPtrOutput{})
	pulumi.RegisterOutputType(RuntimeClassTypeOutput{})
	pulumi.RegisterOutputType(RuntimeClassTypeArrayOutput{})
	pulumi.RegisterOutputType(RuntimeClassListTypeOutput{})
	pulumi.RegisterOutputType(RuntimeClassSpecOutput{})
	pulumi.RegisterOutputType(RuntimeClassSpecPtrOutput{})
	pulumi.RegisterOutputType(SchedulingOutput{})
	pulumi.RegisterOutputType(SchedulingPtrOutput{})
}
