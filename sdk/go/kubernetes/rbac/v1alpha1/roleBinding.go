// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// RoleBinding references a role, but does not contain it.  It can reference a Role in the same namespace or a ClusterRole in the global namespace. It adds who information via Subjects and namespace information by which namespace it exists in.  RoleBindings in a given namespace only have effect in that namespace. Deprecated in v1.17 in favor of rbac.authorization.k8s.io/v1 RoleBinding, and will no longer be served in v1.20.
type RoleBinding struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Standard object's metadata.
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	// RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
	RoleRef RoleRefOutput `pulumi:"roleRef"`
	// Subjects holds references to the objects the role applies to.
	Subjects SubjectArrayOutput `pulumi:"subjects"`
}

// NewRoleBinding registers a new resource with the given unique name, arguments, and options.
func NewRoleBinding(ctx *pulumi.Context,
	name string, args *RoleBindingArgs, opts ...pulumi.ResourceOption) (*RoleBinding, error) {
	if args == nil || args.RoleRef == nil {
		return nil, errors.New("missing required argument 'RoleRef'")
	}
	if args == nil {
		args = &RoleBindingArgs{}
	}
	args.ApiVersion = pulumi.StringPtr("rbac.authorization.k8s.io/v1alpha1")
	args.Kind = pulumi.StringPtr("RoleBinding")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:rbac.authorization.k8s.io/v1:RoleBinding"),
		},
		{
			Type: pulumi.String("kubernetes:rbac.authorization.k8s.io/v1beta1:RoleBinding"),
		},
	})
	opts = append(opts, aliases)
	var resource RoleBinding
	err := ctx.RegisterResource("kubernetes:rbac.authorization.k8s.io/v1alpha1:RoleBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoleBinding gets an existing RoleBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoleBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleBindingState, opts ...pulumi.ResourceOption) (*RoleBinding, error) {
	var resource RoleBinding
	err := ctx.ReadResource("kubernetes:rbac.authorization.k8s.io/v1alpha1:RoleBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoleBinding resources.
type roleBindingState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata.
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
	RoleRef *RoleRef `pulumi:"roleRef"`
	// Subjects holds references to the objects the role applies to.
	Subjects []Subject `pulumi:"subjects"`
}

type RoleBindingState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata.
	Metadata metav1.ObjectMetaPtrInput
	// RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
	RoleRef RoleRefPtrInput
	// Subjects holds references to the objects the role applies to.
	Subjects SubjectArrayInput
}

func (RoleBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleBindingState)(nil)).Elem()
}

type roleBindingArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata.
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
	RoleRef RoleRef `pulumi:"roleRef"`
	// Subjects holds references to the objects the role applies to.
	Subjects []Subject `pulumi:"subjects"`
}

// The set of arguments for constructing a RoleBinding resource.
type RoleBindingArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata.
	Metadata metav1.ObjectMetaPtrInput
	// RoleRef can reference a Role in the current namespace or a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
	RoleRef RoleRefInput
	// Subjects holds references to the objects the role applies to.
	Subjects SubjectArrayInput
}

func (RoleBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleBindingArgs)(nil)).Elem()
}
