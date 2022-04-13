// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aquasec
{
    /// <summary>
    /// The `aquasec.Group` resource manages your groups within Aqua.
    /// 
    /// The Groups created must have at least one Role that is already present within Aqua.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aquasec = Pulumi.Aquasec;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @group = new Aquasec.Group("group", new Aquasec.GroupArgs
    ///         {
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [AquasecResourceType("aquasec:index/group:Group")]
    public partial class Group : Pulumi.CustomResource
    {
        [Output("created")]
        public Output<string> Created { get; private set; } = null!;

        [Output("groupId")]
        public Output<int> GroupId { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a Group resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Group(string name, GroupArgs? args = null, CustomResourceOptions? options = null)
            : base("aquasec:index/group:Group", name, args ?? new GroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Group(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
            : base("aquasec:index/group:Group", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Group resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Group Get(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
        {
            return new Group(name, id, state, options);
        }
    }

    public sealed class GroupArgs : Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GroupArgs()
        {
        }
    }

    public sealed class GroupState : Pulumi.ResourceArgs
    {
        [Input("created")]
        public Input<string>? Created { get; set; }

        [Input("groupId")]
        public Input<int>? GroupId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        public GroupState()
        {
        }
    }
}