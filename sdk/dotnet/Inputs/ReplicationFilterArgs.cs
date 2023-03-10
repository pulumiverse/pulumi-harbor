// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Harbor.Inputs
{

    public sealed class ReplicationFilterArgs : global::Pulumi.ResourceArgs
    {
        [Input("decoration")]
        public Input<string>? Decoration { get; set; }

        [Input("labels")]
        private InputList<string>? _labels;
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("resource")]
        public Input<string>? Resource { get; set; }

        [Input("tag")]
        public Input<string>? Tag { get; set; }

        public ReplicationFilterArgs()
        {
        }
        public static new ReplicationFilterArgs Empty => new ReplicationFilterArgs();
    }
}
