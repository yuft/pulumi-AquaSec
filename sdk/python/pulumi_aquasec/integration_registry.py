# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IntegrationRegistryArgs', 'IntegrationRegistry']

@pulumi.input_type
class IntegrationRegistryArgs:
    def __init__(__self__, *,
                 prefixes: pulumi.Input[Sequence[pulumi.Input[str]]],
                 type: pulumi.Input[str],
                 author: Optional[pulumi.Input[str]] = None,
                 auto_pull: Optional[pulumi.Input[bool]] = None,
                 auto_pull_max: Optional[pulumi.Input[int]] = None,
                 auto_pull_time: Optional[pulumi.Input[str]] = None,
                 last_updated: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IntegrationRegistry resource.
        """
        pulumi.set(__self__, "prefixes", prefixes)
        pulumi.set(__self__, "type", type)
        if author is not None:
            pulumi.set(__self__, "author", author)
        if auto_pull is not None:
            pulumi.set(__self__, "auto_pull", auto_pull)
        if auto_pull_max is not None:
            pulumi.set(__self__, "auto_pull_max", auto_pull_max)
        if auto_pull_time is not None:
            pulumi.set(__self__, "auto_pull_time", auto_pull_time)
        if last_updated is not None:
            pulumi.set(__self__, "last_updated", last_updated)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def prefixes(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "prefixes")

    @prefixes.setter
    def prefixes(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "prefixes", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def author(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "author")

    @author.setter
    def author(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "author", value)

    @property
    @pulumi.getter(name="autoPull")
    def auto_pull(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "auto_pull")

    @auto_pull.setter
    def auto_pull(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_pull", value)

    @property
    @pulumi.getter(name="autoPullMax")
    def auto_pull_max(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "auto_pull_max")

    @auto_pull_max.setter
    def auto_pull_max(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "auto_pull_max", value)

    @property
    @pulumi.getter(name="autoPullTime")
    def auto_pull_time(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "auto_pull_time")

    @auto_pull_time.setter
    def auto_pull_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_pull_time", value)

    @property
    @pulumi.getter(name="lastUpdated")
    def last_updated(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "last_updated")

    @last_updated.setter
    def last_updated(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_updated", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class _IntegrationRegistryState:
    def __init__(__self__, *,
                 author: Optional[pulumi.Input[str]] = None,
                 auto_pull: Optional[pulumi.Input[bool]] = None,
                 auto_pull_max: Optional[pulumi.Input[int]] = None,
                 auto_pull_time: Optional[pulumi.Input[str]] = None,
                 last_updated: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IntegrationRegistry resources.
        """
        if author is not None:
            pulumi.set(__self__, "author", author)
        if auto_pull is not None:
            pulumi.set(__self__, "auto_pull", auto_pull)
        if auto_pull_max is not None:
            pulumi.set(__self__, "auto_pull_max", auto_pull_max)
        if auto_pull_time is not None:
            pulumi.set(__self__, "auto_pull_time", auto_pull_time)
        if last_updated is not None:
            pulumi.set(__self__, "last_updated", last_updated)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if prefixes is not None:
            pulumi.set(__self__, "prefixes", prefixes)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def author(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "author")

    @author.setter
    def author(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "author", value)

    @property
    @pulumi.getter(name="autoPull")
    def auto_pull(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "auto_pull")

    @auto_pull.setter
    def auto_pull(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_pull", value)

    @property
    @pulumi.getter(name="autoPullMax")
    def auto_pull_max(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "auto_pull_max")

    @auto_pull_max.setter
    def auto_pull_max(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "auto_pull_max", value)

    @property
    @pulumi.getter(name="autoPullTime")
    def auto_pull_time(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "auto_pull_time")

    @auto_pull_time.setter
    def auto_pull_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_pull_time", value)

    @property
    @pulumi.getter(name="lastUpdated")
    def last_updated(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "last_updated")

    @last_updated.setter
    def last_updated(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_updated", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def prefixes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "prefixes")

    @prefixes.setter
    def prefixes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "prefixes", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class IntegrationRegistry(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 author: Optional[pulumi.Input[str]] = None,
                 auto_pull: Optional[pulumi.Input[bool]] = None,
                 auto_pull_max: Optional[pulumi.Input[int]] = None,
                 auto_pull_time: Optional[pulumi.Input[str]] = None,
                 last_updated: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a IntegrationRegistry resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IntegrationRegistryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a IntegrationRegistry resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param IntegrationRegistryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IntegrationRegistryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 author: Optional[pulumi.Input[str]] = None,
                 auto_pull: Optional[pulumi.Input[bool]] = None,
                 auto_pull_max: Optional[pulumi.Input[int]] = None,
                 auto_pull_time: Optional[pulumi.Input[str]] = None,
                 last_updated: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
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
            __props__ = IntegrationRegistryArgs.__new__(IntegrationRegistryArgs)

            __props__.__dict__["author"] = author
            __props__.__dict__["auto_pull"] = auto_pull
            __props__.__dict__["auto_pull_max"] = auto_pull_max
            __props__.__dict__["auto_pull_time"] = auto_pull_time
            __props__.__dict__["last_updated"] = last_updated
            __props__.__dict__["name"] = name
            __props__.__dict__["password"] = password
            if prefixes is None and not opts.urn:
                raise TypeError("Missing required property 'prefixes'")
            __props__.__dict__["prefixes"] = prefixes
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["url"] = url
            __props__.__dict__["username"] = username
        super(IntegrationRegistry, __self__).__init__(
            'aquasec:index/integrationRegistry:IntegrationRegistry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            author: Optional[pulumi.Input[str]] = None,
            auto_pull: Optional[pulumi.Input[bool]] = None,
            auto_pull_max: Optional[pulumi.Input[int]] = None,
            auto_pull_time: Optional[pulumi.Input[str]] = None,
            last_updated: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            type: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'IntegrationRegistry':
        """
        Get an existing IntegrationRegistry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IntegrationRegistryState.__new__(_IntegrationRegistryState)

        __props__.__dict__["author"] = author
        __props__.__dict__["auto_pull"] = auto_pull
        __props__.__dict__["auto_pull_max"] = auto_pull_max
        __props__.__dict__["auto_pull_time"] = auto_pull_time
        __props__.__dict__["last_updated"] = last_updated
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["prefixes"] = prefixes
        __props__.__dict__["type"] = type
        __props__.__dict__["url"] = url
        __props__.__dict__["username"] = username
        return IntegrationRegistry(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def author(self) -> pulumi.Output[str]:
        return pulumi.get(self, "author")

    @property
    @pulumi.getter(name="autoPull")
    def auto_pull(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "auto_pull")

    @property
    @pulumi.getter(name="autoPullMax")
    def auto_pull_max(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "auto_pull_max")

    @property
    @pulumi.getter(name="autoPullTime")
    def auto_pull_time(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "auto_pull_time")

    @property
    @pulumi.getter(name="lastUpdated")
    def last_updated(self) -> pulumi.Output[str]:
        return pulumi.get(self, "last_updated")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def prefixes(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "prefixes")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "username")

