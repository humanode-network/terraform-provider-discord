# Discord News Channel Resource

A resource to create a news channel

## Example Usage

```hcl-terraform
resource discord_news_channel general {
  name = "general"
  server_id = var.server_id
  position = 0
}
```

## Argument Reference

* `name` (Required) Name of the category
* `server_id` (Required) ID of server this category is in
* `position` (Optional) Position of the channel, 0-indexed
* `topic` (Optional) Topic of the channel
* `category` (Optional) ID of category to place this channel in
* `sync_perms_with_category` (Optional) Whether channel permissions should be synced or not with the category this channel is in