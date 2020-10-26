---
subcategory: "Test Full"
layout: "test"
page_title: "Test: test_full"
description: |-
  Manages a Test Full
---

# Resource: test_full

Manages a Test Full.

## Example Usage

```hcl
resource "test_full" "example" {
  name = "example"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of thing.
* `tags` - (Optional) Key-value map of resource tags.
* `type` - (Optional) Type of thing.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Name of thing.

## Timeouts

`test_full` provides the following [Timeouts](/docs/configuration/resources.html#timeouts)
configuration options:

* `create` - (Default `10m`) How long to wait for the thing to be created.

## Import

Test Fulls can be imported using the `name`, e.g.

```
$ terraform import test_full.example example
```
