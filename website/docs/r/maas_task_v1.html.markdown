---
layout: "telefonicaopencloud"
page_title: "TelefonicaOpenCloud: telefonicaopencloud_maas_task_v1"
sidebar_current: "docs-telefonicaopencloud-resource-maas-task-v1"
description: |-
  Manages resource task within TelefonicaOpenCloud MAAS.
---

# telefonicaopencloud\_maas\_task\_v1

Manages resource task within TelefonicaOpenCloud MAAS.

## Example Usage:  Creating a MAAS task

```hcl
resource "telefonicaopencloud_maas_task_v1" "task_1" {
  description = "migration task"
  enable_kms = true
  thread_num = 1
  src_node {
    region = "ap-northeast-1"
    ak = "AK",
	sk = "SK",
    object_key = "123.txt",
    bucket = "tommy-bucket",
  }
  dst_node {
    region = "eu-de",
    ak = "AK",
    sk = "SK",
    object_key = "maas/",
    bucket = "test-maas",
  }
}
```

## Argument Reference

The following arguments are supported:

* `src_node` - (Required) Specifies the source node information.

* `dst_node` - (Required) Specifies the destination node information.

* `enable_kms` - (Required) Specifies whether to use KMS encryption.

* `thread_num` - (Required) Specifies the number of threads used by the migration
	task. The value cannot exceed 50.

* `description` - (Optional) Specifies tasks description, which cannot exceed 255
	characters. The following special characters are not allowed: <>()"&

* `smn_info` - (Optional) Specifies the field used for sending messages using the
	Simple Message Notification (SMN) service.


The `src_node` block supports:

* `region` - (Required) Specifies the region where the source bucket locates.
* `ak` - (Required) Specifies the source bucket Access Key.
* `sk` - (Required) Specifies the source bucket Secret Key.
* `object_key` - (Required) Specifies the name of the object to be selected in the
    source bucket.
* `bucket` - (Required) Specifies the name of the source bucket.
* `cloud_type` - (Optional) Specifies the source cloud vendor. The default value is
	AWS and only AWS is supported now.

The `dst_node` block supports:

* `region` - (Required) Specifies the region where the destination bucket locates.
* `ak` - (Required) Specifies the destination bucket Access Key.
* `sk` - (Required) Specifies the destination bucket Secret Key.
* `object_key` - (Optional) Specifies the name of the object to be selected in the
    destination bucket.
* `bucket` - (Required) Specifies the name of the destination bucket.

The `smn_info` block supports:

* `topic_urn` - (Required) Specifies the SMN message topic URN bound to a migration
	task.
* `language` - (Optional) Specifies the management console language used by the
	current users. Users can select en-us.
* `trigger_conditions` - (Required) Specifies the trigger conditions of sending messages
	using SMN. The value depending on the state of a migration task. The migration task
	status can be SUCCESS or FAIL.

## Attributes Reference

The following attributes are exported:

* `src_node` - See Argument Reference above.
* `dst_node` - See Argument Reference above.
* `enable_kms` - See Argument Reference above.
* `thread_num` - See Argument Reference above.
* `description` - See Argument Reference above.
* `smn_info` - See Argument Reference above.
* `name` - Specifies the name for a task.
* `status` - Specifies the task status as follows: 0: Not started, 1: Waiting to migrate,
	2: Migrating, 3: Migration paused, 4: Migration failed, 5: Migration succeeded.
