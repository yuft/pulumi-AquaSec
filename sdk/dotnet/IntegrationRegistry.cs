// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aquasec
{
    [AquasecResourceType("aquasec:index/integrationRegistry:IntegrationRegistry")]
    public partial class IntegrationRegistry : Pulumi.CustomResource
    {
        [Output("author")]
        public Output<string> Author { get; private set; } = null!;

        [Output("autoPull")]
        public Output<bool?> AutoPull { get; private set; } = null!;

        [Output("autoPullMax")]
        public Output<int?> AutoPullMax { get; private set; } = null!;

        [Output("autoPullTime")]
        public Output<string?> AutoPullTime { get; private set; } = null!;

        [Output("lastUpdated")]
        public Output<string> LastUpdated { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        [Output("prefixes")]
        public Output<ImmutableArray<string>> Prefixes { get; private set; } = null!;

        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        [Output("url")]
        public Output<string?> Url { get; private set; } = null!;

        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;


        /// <summary>
        /// Create a IntegrationRegistry resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IntegrationRegistry(string name, IntegrationRegistryArgs args, CustomResourceOptions? options = null)
            : base("aquasec:index/integrationRegistry:IntegrationRegistry", name, args ?? new IntegrationRegistryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IntegrationRegistry(string name, Input<string> id, IntegrationRegistryState? state = null, CustomResourceOptions? options = null)
            : base("aquasec:index/integrationRegistry:IntegrationRegistry", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IntegrationRegistry resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IntegrationRegistry Get(string name, Input<string> id, IntegrationRegistryState? state = null, CustomResourceOptions? options = null)
        {
            return new IntegrationRegistry(name, id, state, options);
        }
    }

    public sealed class IntegrationRegistryArgs : Pulumi.ResourceArgs
    {
        [Input("author")]
        public Input<string>? Author { get; set; }

        [Input("autoPull")]
        public Input<bool>? AutoPull { get; set; }

        [Input("autoPullMax")]
        public Input<int>? AutoPullMax { get; set; }

        [Input("autoPullTime")]
        public Input<string>? AutoPullTime { get; set; }

        [Input("lastUpdated")]
        public Input<string>? LastUpdated { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("prefixes", required: true)]
        private InputList<string>? _prefixes;
        public InputList<string> Prefixes
        {
            get => _prefixes ?? (_prefixes = new InputList<string>());
            set => _prefixes = value;
        }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("url")]
        public Input<string>? Url { get; set; }

        [Input("username")]
        public Input<string>? Username { get; set; }

        public IntegrationRegistryArgs()
        {
        }
    }

    public sealed class IntegrationRegistryState : Pulumi.ResourceArgs
    {
        [Input("author")]
        public Input<string>? Author { get; set; }

        [Input("autoPull")]
        public Input<bool>? AutoPull { get; set; }

        [Input("autoPullMax")]
        public Input<int>? AutoPullMax { get; set; }

        [Input("autoPullTime")]
        public Input<string>? AutoPullTime { get; set; }

        [Input("lastUpdated")]
        public Input<string>? LastUpdated { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("prefixes")]
        private InputList<string>? _prefixes;
        public InputList<string> Prefixes
        {
            get => _prefixes ?? (_prefixes = new InputList<string>());
            set => _prefixes = value;
        }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("url")]
        public Input<string>? Url { get; set; }

        [Input("username")]
        public Input<string>? Username { get; set; }

        public IntegrationRegistryState()
        {
        }
    }
}
