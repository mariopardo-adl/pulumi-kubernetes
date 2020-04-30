// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// SelfSubjectAccessReview checks whether or the current user can perform an action.  Not filling in a spec.namespace means "in all namespaces".  Self is a special case, because users should always be able to check whether they can perform an action
type SelfSubjectAccessReview struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringOutput     `pulumi:"kind"`
	Metadata metav1.ObjectMetaOutput `pulumi:"metadata"`
	// Spec holds information about the request being evaluated.  user and groups must be empty
	Spec SelfSubjectAccessReviewSpecOutput `pulumi:"spec"`
	// Status is filled in by the server and indicates whether the request is allowed or not
	Status SubjectAccessReviewStatusOutput `pulumi:"status"`
}

// NewSelfSubjectAccessReview registers a new resource with the given unique name, arguments, and options.
func NewSelfSubjectAccessReview(ctx *pulumi.Context,
	name string, args *SelfSubjectAccessReviewArgs, opts ...pulumi.ResourceOption) (*SelfSubjectAccessReview, error) {
	if args == nil || args.Spec == nil {
		return nil, errors.New("missing required argument 'Spec'")
	}
	if args == nil {
		args = &SelfSubjectAccessReviewArgs{}
	}
	args.ApiVersion = pulumi.StringPtr("authorization.k8s.io/v1")
	args.Kind = pulumi.StringPtr("SelfSubjectAccessReview")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:authorization.k8s.io/v1beta1:SelfSubjectAccessReview"),
		},
	})
	opts = append(opts, aliases)
	var resource SelfSubjectAccessReview
	err := ctx.RegisterResource("kubernetes:authorization.k8s.io/v1:SelfSubjectAccessReview", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSelfSubjectAccessReview gets an existing SelfSubjectAccessReview resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSelfSubjectAccessReview(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SelfSubjectAccessReviewState, opts ...pulumi.ResourceOption) (*SelfSubjectAccessReview, error) {
	var resource SelfSubjectAccessReview
	err := ctx.ReadResource("kubernetes:authorization.k8s.io/v1:SelfSubjectAccessReview", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SelfSubjectAccessReview resources.
type selfSubjectAccessReviewState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec holds information about the request being evaluated.  user and groups must be empty
	Spec *SelfSubjectAccessReviewSpec `pulumi:"spec"`
	// Status is filled in by the server and indicates whether the request is allowed or not
	Status *SubjectAccessReviewStatus `pulumi:"status"`
}

type SelfSubjectAccessReviewState struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// Spec holds information about the request being evaluated.  user and groups must be empty
	Spec SelfSubjectAccessReviewSpecPtrInput
	// Status is filled in by the server and indicates whether the request is allowed or not
	Status SubjectAccessReviewStatusPtrInput
}

func (SelfSubjectAccessReviewState) ElementType() reflect.Type {
	return reflect.TypeOf((*selfSubjectAccessReviewState)(nil)).Elem()
}

type selfSubjectAccessReviewArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     *string            `pulumi:"kind"`
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec holds information about the request being evaluated.  user and groups must be empty
	Spec SelfSubjectAccessReviewSpec `pulumi:"spec"`
}

// The set of arguments for constructing a SelfSubjectAccessReview resource.
type SelfSubjectAccessReviewArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind     pulumi.StringPtrInput
	Metadata metav1.ObjectMetaPtrInput
	// Spec holds information about the request being evaluated.  user and groups must be empty
	Spec SelfSubjectAccessReviewSpecInput
}

func (SelfSubjectAccessReviewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*selfSubjectAccessReviewArgs)(nil)).Elem()
}
