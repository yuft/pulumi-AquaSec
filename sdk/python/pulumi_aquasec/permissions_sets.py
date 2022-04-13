# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['PermissionsSetsArgs', 'PermissionsSets']

@pulumi.input_type
class PermissionsSetsArgs:
    def __init__(__self__, *,
                 actions: pulumi.Input[Sequence[pulumi.Input[str]]],
                 ui_access: pulumi.Input[bool],
                 author: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 is_super: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PermissionsSets resource.
        :param pulumi.Input[str] id: The ID of this resource.
        """
        pulumi.set(__self__, "actions", actions)
        pulumi.set(__self__, "ui_access", ui_access)
        if author is not None:
            pulumi.set(__self__, "author", author)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if is_super is not None:
            pulumi.set(__self__, "is_super", is_super)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def actions(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "actions")

    @actions.setter
    def actions(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "actions", value)

    @property
    @pulumi.getter(name="uiAccess")
    def ui_access(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "ui_access")

    @ui_access.setter
    def ui_access(self, value: pulumi.Input[bool]):
        pulumi.set(self, "ui_access", value)

    @property
    @pulumi.getter
    def author(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "author")

    @author.setter
    def author(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "author", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="isSuper")
    def is_super(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "is_super")

    @is_super.setter
    def is_super(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_super", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _PermissionsSetsState:
    def __init__(__self__, *,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 author: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 is_super: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ui_access: Optional[pulumi.Input[bool]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PermissionsSets resources.
        :param pulumi.Input[str] id: The ID of this resource.
        """
        if actions is not None:
            pulumi.set(__self__, "actions", actions)
        if author is not None:
            pulumi.set(__self__, "author", author)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if is_super is not None:
            pulumi.set(__self__, "is_super", is_super)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if ui_access is not None:
            pulumi.set(__self__, "ui_access", ui_access)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter
    def actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "actions")

    @actions.setter
    def actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "actions", value)

    @property
    @pulumi.getter
    def author(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "author")

    @author.setter
    def author(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "author", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="isSuper")
    def is_super(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "is_super")

    @is_super.setter
    def is_super(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_super", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="uiAccess")
    def ui_access(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "ui_access")

    @ui_access.setter
    def ui_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ui_access", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class PermissionsSets(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 author: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 is_super: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ui_access: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        The `PermissionsSets` resource manages your Permission Set within Aqua.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aquasec as aquasec

        my_terraform_perm_set = aquasec.PermissionsSets("myTerraformPermSet",
            actions=[
                "dashboard.read",
                "risks.vulnerabilities.read",
                "risks.vulnerabilities.write",
                "risks.host_images.read",
                "risks.benchmark.read",
                "risk_explorer.read",
                "images.read",
                "image_profiles.read",
                "image_assurance.read",
                "image_assurance.write",
                "runtime_policies.read",
                "runtime_policies.write",
                "functions.read",
                "gateways.read",
                "secrets.read",
                "audits.read",
                "containers.read",
                "enforcers.read",
                "infrastructure.read",
                "consoles.read",
                "settings.read",
                "network_policies.read",
                "acl_policies.read",
                "acl_policies.write",
                "services.read",
                "integrations.read",
                "registries_integrations.read",
                "web_hook.read",
                "incidents.read",
            ],
            author="system",
            description="created from terraform",
            is_super=False,
            ui_access=True)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] id: The ID of this resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PermissionsSetsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `PermissionsSets` resource manages your Permission Set within Aqua.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aquasec as aquasec

        my_terraform_perm_set = aquasec.PermissionsSets("myTerraformPermSet",
            actions=[
                "dashboard.read",
                "risks.vulnerabilities.read",
                "risks.vulnerabilities.write",
                "risks.host_images.read",
                "risks.benchmark.read",
                "risk_explorer.read",
                "images.read",
                "image_profiles.read",
                "image_assurance.read",
                "image_assurance.write",
                "runtime_policies.read",
                "runtime_policies.write",
                "functions.read",
                "gateways.read",
                "secrets.read",
                "audits.read",
                "containers.read",
                "enforcers.read",
                "infrastructure.read",
                "consoles.read",
                "settings.read",
                "network_policies.read",
                "acl_policies.read",
                "acl_policies.write",
                "services.read",
                "integrations.read",
                "registries_integrations.read",
                "web_hook.read",
                "incidents.read",
            ],
            author="system",
            description="created from terraform",
            is_super=False,
            ui_access=True)
        ```

        :param str resource_name: The name of the resource.
        :param PermissionsSetsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PermissionsSetsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 author: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 is_super: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ui_access: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PermissionsSetsArgs.__new__(PermissionsSetsArgs)

            if actions is None and not opts.urn:
                raise TypeError("Missing required property 'actions'")
            __props__.__dict__["actions"] = actions
            __props__.__dict__["author"] = author
            __props__.__dict__["description"] = description
            __props__.__dict__["id"] = id
            __props__.__dict__["is_super"] = is_super
            __props__.__dict__["name"] = name
            if ui_access is None and not opts.urn:
                raise TypeError("Missing required property 'ui_access'")
            __props__.__dict__["ui_access"] = ui_access
            __props__.__dict__["updated_at"] = None
        super(PermissionsSets, __self__).__init__(
            'aquasec:index/permissionsSets:PermissionsSets',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            author: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            id: Optional[pulumi.Input[str]] = None,
            is_super: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            ui_access: Optional[pulumi.Input[bool]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'PermissionsSets':
        """
        Get an existing PermissionsSets resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] id: The ID of this resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PermissionsSetsState.__new__(_PermissionsSetsState)

        __props__.__dict__["actions"] = actions
        __props__.__dict__["author"] = author
        __props__.__dict__["description"] = description
        __props__.__dict__["id"] = id
        __props__.__dict__["is_super"] = is_super
        __props__.__dict__["name"] = name
        __props__.__dict__["ui_access"] = ui_access
        __props__.__dict__["updated_at"] = updated_at
        return PermissionsSets(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def actions(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter
    def author(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "author")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> pulumi.Output[str]:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isSuper")
    def is_super(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "is_super")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="uiAccess")
    def ui_access(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "ui_access")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "updated_at")
