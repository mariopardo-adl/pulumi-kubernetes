// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Inputs.Discovery.V1Beta1
{

    /// <summary>
    /// EndpointConditions represents the current condition of an endpoint.
    /// </summary>
    public class EndpointConditionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ready indicates that this endpoint is prepared to receive traffic, according to whatever system is managing the endpoint. A nil value indicates an unknown state. In most cases consumers should interpret this unknown state as ready.
        /// </summary>
        [Input("ready")]
        public Input<bool>? Ready { get; set; }

        public EndpointConditionsArgs()
        {
        }
    }
}
