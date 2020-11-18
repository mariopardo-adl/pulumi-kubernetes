// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Storage.V1Beta1
{

    /// <summary>
    /// TokenRequest contains parameters of a service account token.
    /// </summary>
    public class TokenRequestArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Audience is the intended audience of the token in "TokenRequestSpec". It will default to the audiences of kube apiserver.
        /// </summary>
        [Input("audience", required: true)]
        public Input<string> Audience { get; set; } = null!;

        /// <summary>
        /// ExpirationSeconds is the duration of validity of the token in "TokenRequestSpec". It has the same default value of "ExpirationSeconds" in "TokenRequestSpec"
        /// </summary>
        [Input("expirationSeconds")]
        public Input<int>? ExpirationSeconds { get; set; }

        public TokenRequestArgs()
        {
        }
    }
}
