package discord

import (
	"github.com/andersfylling/disgord"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
)

func resourceDiscordSystemChannel() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSystemChannelCreate,
		ReadContext:   resourceSystemChannelRead,
		UpdateContext: resourceSystemChannelUpdate,
		DeleteContext: resourceSystemChannelDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"server_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"system_channel_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSystemChannelCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*Context).Client

	serverId := disgord.Snowflake(0)
	if v, ok := d.GetOk("server_id"); ok {
		serverId = disgord.ParseSnowflakeString(v.(string))
	}

	_, err := client.Guild(serverId).Get()
	if err != nil {
		return diag.Errorf("Failed to find server: %s", err.Error())
	}

	var systemChannelId disgord.Snowflake
	if v, ok := d.GetOk("system_channel_id"); ok {
		systemChannelId = disgord.ParseSnowflakeString(v.(string))
	} else {
		return diag.Errorf("Failed to parse system channel id: %s", err.Error())
	}

	if _, err := client.Guild(serverId).Update(&disgord.UpdateGuild{
		SystemChannelID: &systemChannelId,
	}); err != nil {
		return diag.Errorf("Failed to edit server: %s", err.Error())
	}

	d.SetId(d.Get("server_id").(string))

	return diags
}

func resourceSystemChannelRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*Context).Client

	serverId := disgord.ParseSnowflakeString(d.Id())

	server, err := client.Guild(serverId).Get()
	if err != nil {
		return diag.Errorf("Error fetching server: %s", err.Error())
	}

	d.Set("system_channel_id", server.SystemChannelID.String())

	return diags
}

func resourceSystemChannelUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*Context).Client

	serverId := disgord.ParseSnowflakeString(d.Get("server_id").(string))

	if _, err := client.Guild(serverId).Get(); err != nil {
		return diag.Errorf("Error fetching server: %s", err.Error())
	}

	if d.HasChange("system_channel_id") {
		id := disgord.ParseSnowflakeString(d.Get("system_channel_id").(string))

		if _, err := client.Guild(serverId).Update(&disgord.UpdateGuild{
			SystemChannelID: &id,
		}); err != nil {
			return diag.Errorf("Failed to edit server: %s", err.Error())
		}
	}

	return diags
}

func resourceSystemChannelDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*Context).Client

	serverId := disgord.ParseSnowflakeString(d.Get("server_id").(string))

	if _, err := client.Guild(serverId).Get(); err != nil {
		return diag.Errorf("Error fetching server: %s", err.Error())
	}

	// TODO: 本当にこれで動くのか？
	systemChannelId := disgord.Snowflake(1)
	if _, err := client.Guild(serverId).Update(&disgord.UpdateGuild{
		SystemChannelID: &systemChannelId,
	}); err != nil {
		return diag.Errorf("Failed to edit server: %s", err.Error())
	}

	return diags
}
