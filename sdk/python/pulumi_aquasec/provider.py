# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 aqua_url: Optional[pulumi.Input[str]] = None,
                 ca_certificate_path: Optional[pulumi.Input[str]] = None,
                 config_path: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 verify_tls: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] aqua_url: This is the base URL of your Aqua instance. Can alternatively be sourced from the `AQUA_URL` environment variable.
        :param pulumi.Input[str] ca_certificate_path: This is the file path for server CA certificates if they are not available on the host OS. Can alternatively be sourced
               from the `AQUA_CA_CERT_PATH` environment variable.
        :param pulumi.Input[str] config_path: This is the file path for Aqua provider configuration. The default configuration path is `~/.aqua/tf.config`. Can
               alternatively be sourced from the `AQUA_CONFIG` environment variable.
        :param pulumi.Input[str] password: This is the password that should be used to make the connection. Can alternatively be sourced from the `AQUA_PASSWORD`
               environment variable.
        :param pulumi.Input[str] username: This is the user id that should be used to make the connection. Can alternatively be sourced from the `AQUA_USER`
               environment variable.
        :param pulumi.Input[bool] verify_tls: If true, server tls certificates will be verified by the client before making a connection. Defaults to true. Can
               alternatively be sourced from the `AQUA_TLS_VERIFY` environment variable.
        """
        if aqua_url is None:
            aqua_url = _utilities.get_env('AQUA_URL')
        if aqua_url is not None:
            pulumi.set(__self__, "aqua_url", aqua_url)
        if ca_certificate_path is not None:
            pulumi.set(__self__, "ca_certificate_path", ca_certificate_path)
        if config_path is not None:
            pulumi.set(__self__, "config_path", config_path)
        if password is None:
            password = _utilities.get_env('AUQA_PASSWORD')
        if password is not None:
            pulumi.set(__self__, "password", password)
        if username is None:
            username = _utilities.get_env('AQUA_USERNAME')
        if username is not None:
            pulumi.set(__self__, "username", username)
        if verify_tls is not None:
            pulumi.set(__self__, "verify_tls", verify_tls)

    @property
    @pulumi.getter(name="aquaUrl")
    def aqua_url(self) -> Optional[pulumi.Input[str]]:
        """
        This is the base URL of your Aqua instance. Can alternatively be sourced from the `AQUA_URL` environment variable.
        """
        return pulumi.get(self, "aqua_url")

    @aqua_url.setter
    def aqua_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aqua_url", value)

    @property
    @pulumi.getter(name="caCertificatePath")
    def ca_certificate_path(self) -> Optional[pulumi.Input[str]]:
        """
        This is the file path for server CA certificates if they are not available on the host OS. Can alternatively be sourced
        from the `AQUA_CA_CERT_PATH` environment variable.
        """
        return pulumi.get(self, "ca_certificate_path")

    @ca_certificate_path.setter
    def ca_certificate_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_certificate_path", value)

    @property
    @pulumi.getter(name="configPath")
    def config_path(self) -> Optional[pulumi.Input[str]]:
        """
        This is the file path for Aqua provider configuration. The default configuration path is `~/.aqua/tf.config`. Can
        alternatively be sourced from the `AQUA_CONFIG` environment variable.
        """
        return pulumi.get(self, "config_path")

    @config_path.setter
    def config_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_path", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        This is the password that should be used to make the connection. Can alternatively be sourced from the `AQUA_PASSWORD`
        environment variable.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        This is the user id that should be used to make the connection. Can alternatively be sourced from the `AQUA_USER`
        environment variable.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter(name="verifyTls")
    def verify_tls(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, server tls certificates will be verified by the client before making a connection. Defaults to true. Can
        alternatively be sourced from the `AQUA_TLS_VERIFY` environment variable.
        """
        return pulumi.get(self, "verify_tls")

    @verify_tls.setter
    def verify_tls(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "verify_tls", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aqua_url: Optional[pulumi.Input[str]] = None,
                 ca_certificate_path: Optional[pulumi.Input[str]] = None,
                 config_path: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 verify_tls: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        The provider type for the aquasec package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aqua_url: This is the base URL of your Aqua instance. Can alternatively be sourced from the `AQUA_URL` environment variable.
        :param pulumi.Input[str] ca_certificate_path: This is the file path for server CA certificates if they are not available on the host OS. Can alternatively be sourced
               from the `AQUA_CA_CERT_PATH` environment variable.
        :param pulumi.Input[str] config_path: This is the file path for Aqua provider configuration. The default configuration path is `~/.aqua/tf.config`. Can
               alternatively be sourced from the `AQUA_CONFIG` environment variable.
        :param pulumi.Input[str] password: This is the password that should be used to make the connection. Can alternatively be sourced from the `AQUA_PASSWORD`
               environment variable.
        :param pulumi.Input[str] username: This is the user id that should be used to make the connection. Can alternatively be sourced from the `AQUA_USER`
               environment variable.
        :param pulumi.Input[bool] verify_tls: If true, server tls certificates will be verified by the client before making a connection. Defaults to true. Can
               alternatively be sourced from the `AQUA_TLS_VERIFY` environment variable.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProviderArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the aquasec package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aqua_url: Optional[pulumi.Input[str]] = None,
                 ca_certificate_path: Optional[pulumi.Input[str]] = None,
                 config_path: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 verify_tls: Optional[pulumi.Input[bool]] = None,
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
            __props__ = ProviderArgs.__new__(ProviderArgs)

            if aqua_url is None:
                aqua_url = _utilities.get_env('AQUA_URL')
            __props__.__dict__["aqua_url"] = aqua_url
            __props__.__dict__["ca_certificate_path"] = ca_certificate_path
            __props__.__dict__["config_path"] = config_path
            if password is None:
                password = _utilities.get_env('AUQA_PASSWORD')
            __props__.__dict__["password"] = password
            if username is None:
                username = _utilities.get_env('AQUA_USERNAME')
            __props__.__dict__["username"] = username
            __props__.__dict__["verify_tls"] = pulumi.Output.from_input(verify_tls).apply(pulumi.runtime.to_json) if verify_tls is not None else None
        super(Provider, __self__).__init__(
            'aquasec',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="aquaUrl")
    def aqua_url(self) -> pulumi.Output[Optional[str]]:
        """
        This is the base URL of your Aqua instance. Can alternatively be sourced from the `AQUA_URL` environment variable.
        """
        return pulumi.get(self, "aqua_url")

    @property
    @pulumi.getter(name="caCertificatePath")
    def ca_certificate_path(self) -> pulumi.Output[Optional[str]]:
        """
        This is the file path for server CA certificates if they are not available on the host OS. Can alternatively be sourced
        from the `AQUA_CA_CERT_PATH` environment variable.
        """
        return pulumi.get(self, "ca_certificate_path")

    @property
    @pulumi.getter(name="configPath")
    def config_path(self) -> pulumi.Output[Optional[str]]:
        """
        This is the file path for Aqua provider configuration. The default configuration path is `~/.aqua/tf.config`. Can
        alternatively be sourced from the `AQUA_CONFIG` environment variable.
        """
        return pulumi.get(self, "config_path")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        This is the password that should be used to make the connection. Can alternatively be sourced from the `AQUA_PASSWORD`
        environment variable.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[Optional[str]]:
        """
        This is the user id that should be used to make the connection. Can alternatively be sourced from the `AQUA_USER`
        environment variable.
        """
        return pulumi.get(self, "username")
