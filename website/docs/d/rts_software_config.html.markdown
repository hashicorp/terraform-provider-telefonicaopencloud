---
layout: "telefonicaopencloud"
page_title: "TelefonicaOpenCloud: telefonicaopencloud_rts_software_config_v1"
sidebar_current: "docs-telefonicaopencloud-datasource-rts-software-config-v1"
description: |-
  Provides details about a specific RTS Software Config.
---

# Data Source: telefonicaopencloud_rts_software_config_v1

The RTS Software Config data source provides details about a specific RTS Software Config.

## Example Usage


```hcl
variable "config_id" {}

variable "server_id" {}

data "telefonicaopencloud_rts_software_config_v1" "myconfig" {
  id = "${var.config_id}"
}

```

## Argument Reference
The following arguments are supported:

* `id` - (Optional) The id of the software configuration.


## Attributes Reference
In addition to all arguments above, the following attributes are exported:

* `name` - The name of the software configuration.

* `group` - The namespace that groups this software configuration by when it is delivered to a server.

* `inputs` -  A list of software configuration inputs.

* `outputs` - A list of software configuration outputs.

* `config` - The software configuration code.

* `options` - The software configuration options.

