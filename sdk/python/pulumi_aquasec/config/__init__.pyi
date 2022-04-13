# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

aquaUrl: Optional[str]
"""
This is the base URL of your Aqua instance. Can alternatively be sourced from the `AQUA_URL` environment variable.
"""

caCertificatePath: Optional[str]
"""
This is the file path for server CA certificates if they are not available on the host OS. Can alternatively be sourced
from the `AQUA_CA_CERT_PATH` environment variable.
"""

configPath: Optional[str]
"""
This is the file path for Aqua provider configuration. The default configuration path is `~/.aqua/tf.config`. Can
alternatively be sourced from the `AQUA_CONFIG` environment variable.
"""

password: Optional[str]
"""
This is the password that should be used to make the connection. Can alternatively be sourced from the `AQUA_PASSWORD`
environment variable.
"""

username: Optional[str]
"""
This is the user id that should be used to make the connection. Can alternatively be sourced from the `AQUA_USER`
environment variable.
"""

verifyTls: Optional[bool]
"""
If true, server tls certificates will be verified by the client before making a connection. Defaults to true. Can
alternatively be sourced from the `AQUA_TLS_VERIFY` environment variable.
"""

