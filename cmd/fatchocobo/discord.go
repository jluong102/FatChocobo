package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

import (
	"github.com/gorilla/websocket"
)

const DISCORD_URL string = "https://discord.com/api"

// Gateway Opcodes
const (
	GATEWAY_OPCODE_DISPATCH              = 0
	GATEWAY_OPCODE_HEARTBEAT             = 1
	GATEWAY_OPCODE_IDENTIFY              = 2
	GATEWAY_OPCODE_PRESENCE_UPDATE       = 3
	GATEWAY_OPCODE_VOICE_STATE_UPDATE    = 4
	GATEWAY_OPCODE_RESUME                = 6
	GATEWAY_OPCODE_RECONNECT             = 7
	GATEWAY_OPCODE_REQUEST_GUILD_MEMBERS = 8
	GATEWAY_OPCODE_INVALID_SESSION       = 9
	GATEWAY_OPCODE_HELLO                 = 10
	GATEWAY_OPCODE_HEARTBEAT_ACK         = 11
)

// Gateway Close Event Codes
const (
	GATEWAY_CLOSE_UNKNOWN_ERROR         = 4000
	GATEWAY_CLOSE_UNKNOWN_OPCODE        = 4001
	GATEWAY_CLOSE_DECODE_ERROR          = 4002
	GATEWAY_CLOSE_NOT_AUTHENTICATED     = 4003
	GATEWAY_CLOSE_AUTHENTICATION_FAILED = 4004
	GATEWAY_CLOSE_ALREADY_AUTHENTICATED = 4005
	GATEWAY_CLOSE_INVALID_SEQ           = 4007
	GATEWAY_CLOSE_RATE_LIMITED          = 4008
	GATEWAY_CLOSE_SESSION_TIMED_OUT     = 4009
	GATEWAY_CLOSE_INVALID_SHARD         = 4010
	GATEWAY_CLOSE_SHARDING_REQUIRED     = 4011
	GATEWAY_CLOSE_INVALID_API_VERSION   = 4012
	GATEWAY_CLOSE_INVALID_INTENT        = 4013
	GATEWAY_CLOSE_DISALLOWED_INTENT     = 4014
)

// Gateway Intents
const (
	GATEWAY_INTENT_GUILDS                        = 1 << 0
	GATEWAY_INTENT_GUILD_MEMBERS                 = 1 << 1
	GATEWAY_INTENT_GUILD_MODERATION              = 1 << 2
	GATEWAY_INTENT_GUILD_EMOJIS_AND_STICKERS     = 1 << 3
	GATEWAY_INTENT_GUILD_INTEGRATIONS            = 1 << 4
	GATEWAY_INTENT_GUILD_WEBHOOKS                = 1 << 5
	GATEWAY_INTENT_GUILD_INVITES                 = 1 << 6
	GATEWAY_INTENT_GUILD_VOICE_STATES            = 1 << 7
	GATEWAY_INTENT_GUILD_PRESENCES               = 1 << 8
	GATEWAY_INTENT_GUILD_MESSAGES                = 1 << 9
	GATEWAY_INTENT_GUILD_MESSAGE_REACTIONS       = 1 << 10
	GATEWAY_INTENT_GUILD_MESSAGE_TYPING          = 1 << 11
	GATEWAY_INTENT_DIRECT_MESSAGES               = 1 << 12
	GATEWAY_INTENT_DIRECT_MESSAGE_REACTIONS      = 1 << 13
	GATEWAY_INTENT_DIRECT_MESSAGE_TYPES          = 1 << 14
	GATEWAY_INTENT_MESSAGE_CONTENT               = 1 << 15
	GATEWAY_INTENT_GUILD_SCHEDULED_EVENTS        = 1 << 16
	GATEWAY_INTENT_AUTO_MODERATION_CONFIGURATION = 1 << 20
	GATEWAY_INTENT_AUTO_MODERATION_EXECUTION     = 1 << 21
)

// User Flags
const (
	USER_FLAG_STAFF                    = 1 << 0
	USER_FLAG_PARTNER                  = 1 << 1
	USER_FLAG_HYPESQUAD                = 1 << 2
	USER_FLAG_BUG_HUNTER_LEVEL_1       = 1 << 3
	USER_FLAG_HYPESQUAD_ONLINE_HOUSE_1 = 1 << 6
	USER_FLAG_HYPESQUAD_ONLINE_HOUSE_2 = 1 << 7
	USER_FLAG_HYPESQUAD_ONLINE_HOUSE_3 = 1 << 8
	USER_FLAG_PREMIUM_EARLY_SUPPORTER  = 1 << 9
	USER_FLAG_TEAM_PSEUDO_USER         = 1 << 10
	USER_FLAG_BUG_HUNTER_LEVEL_2       = 1 << 14
	USER_FLAG_VERIFIED_BOT             = 1 << 16
	USER_FLAG_VERIFIED_DEVELOPER       = 1 << 17
	USER_FLAG_CERTIFIED_MODERATOR      = 1 << 18
	USER_FLAG_BOT_HTTP_INTERACTIONS    = 1 << 19
	USER_FLAG_ACTIVE_DEVELOPER         = 1 << 22
)

// Activity Flags
const (
	ACTIVITY_FLAG_INSTANCE              = 1 << 0
	ACTIVITY_FLAG_JOIN                  = 1 << 1
	ACTIVITY_FLAG_SPECTATE              = 1 << 2
	ACTIVITY_FLAG_JOIN_REQUEST          = 1 << 3
	ACTIVITY_FLAG_SYNC                  = 1 << 4
	ACTIVITY_FLAG_PLAY                  = 1 << 5
	ACTIVITY_FLAG_PARTY_PRIVACY_FRIENDS = 1 << 6
	ACTIVITY_FLAG_PRIVACY_VOICE_CHANNEL = 1 << 7
	ACTIVITY_FLAG_EMBEDDED              = 1 << 8
)

// Message Flags
const (
	MESSAGE_FLAGS_CROSSPOSTED                            = 1 << 0
	MESSAGE_FLAGS_IS_CROSSPOST                           = 1 << 1
	MESSAGE_FLAGS_SUPPRESS_EMBEDS                        = 1 << 2
	MESSAGE_FLAGS_SOURCE_MESSAGE_DELETED                 = 1 << 3
	MESSAGE_FLAGS_URGENT                                 = 1 << 4
	MESSAGE_FLAGS_HAS_THREAD                             = 1 << 5
	MESSAGE_FLAGS_EPHEMERAL                              = 1 << 6
	MESSAGE_FLAGS_LOADING                                = 1 << 7
	MESSAGE_FLAGS_FAILED_TO_MENTION_SOME_ROLES_IN_THREAD = 1 << 8
	MESSAGE_FLAGS_SUPRESS_NOTIFICATIONS                  = 1 << 12
	MESSAGE_FLAGS_IS_VOICE_MESSAGE                       = 1 << 13
)

// Premium types
const (
	PREMIUM_TYPES_NONE          = 0
	PREMIUM_TYPES_NITRO_CLASSIC = 1
	PREMIUM_TYPES_NITRO         = 2
	PREMIUM_TYPES_NITRO_BASIC   = 3
)

// Membership State Enum
const (
	MEMBERSHIP_STATE_INVITED  = 1
	MEMBERSHIP_STATE_ACCEPTED = 2
)

// Sticker types
const (
	STICKER_TYPE_STANDARD = 1
	STICKER_TYPE_GUILD    = 2
)

// Message Activity Types
const (
	MESSAGE_ACTIVITY_TYPE_JOIN         = 1
	MESSAGE_ACTIVITY_TYPE_SPECTATE     = 2
	MESSAGE_ACTIVITY_TYPE_LISTEN       = 3
	MESSAGE_ACTIVITY_TYPE_JOIN_REQUEST = 4
)

// Component Types
const (
	COMPONENT_TYPE_ACTION_ROW         = 1
	COMPONENT_TYPE_BUTTON             = 2
	COMPONENT_TYPE_STRING_SELECT      = 3
	COMPONENT_TYPE_TEXT_INPUT         = 4
	COMPONENT_TYPE_USER_SELECT        = 5
	COMPONENT_TYPE_ROLE_SELECT        = 6
	COMPONENT_TYPE_MENTIONABLE_SELECT = 7
	COMPONENT_TYPE_CHANNEL_SELECT     = 8
)

type Discord struct {
	Websocket *websocket.Conn
	Heartbeat int
	token     string
}

type Snowflake string

// HTTP Responses
type GatewayBotResponse struct {
	Url               string                  `json:"url"`
	Shards            int                     `json:"shards"`
	SessionStartLimit SessionStartLimitObject `json:"session_start_limit"`
}

// Discord JSON sub-objects
type SessionStartLimitObject struct {
	Total          int `json:"total"`
	Remaining      int `json:"remaining"`
	ResetAfter     int `json:"reset_after"`
	MaxConcurrency int `json:"max_concurrency"`
}

type UserObject struct {
	Id               Snowflake `json:"id"`
	Username         string    `json:"username"`
	Discriminator    string    `json:"discriminator"`
	GlobalName       string    `json:"global_name"`
	Avatar           string    `json:"avatar"`
	Bot              bool      `json:"bot"`
	System           bool      `json:"system"`
	MFAEnabled       bool      `json:"mfa_enabled"`
	Banner           string    `json:"banner"`
	AccentColor      int       `json:"accent_color"`
	Locale           string    `json:"locale"`
	Verified         bool      `json:"verified"`
	Email            string    `json:"email"`
	Flags            int       `json:"flag"`
	PremiumType      int       `json:"premium_type"`
	PublicFlags      int       `json:"public_flags"`
	AvatarDecoration string    `json:"avatar_decoration"`
}

type TeamObject struct {
	Icon        string             `json:"icon"`
	Id          Snowflake          `json:"id"`
	Members     []TeamMemberObject `json:"members"`
	Name        string             `json:"name"`
	OwnerUserId Snowflake          `json:"owner_user_id"`
}

type TeamMemberObject struct {
	MembershipState int        `json:"membership_state"`
	TeamId          Snowflake  `json:"team_id"`
	User            UserObject `json:"user"`
	Role            string     `json:"role"`
}

type GuildObject struct {
	Id                          Snowflake           `json:"id"`
	Name                        string              `json:"name"`
	Icon                        string              `json:"icon"`
	IconHash                    string              `json:"icon_hash"`
	Splash                      string              `json:"splash"`
	DiscoverySplash             string              `json:"string"`
	Owner                       bool                `json:"owner"`
	OwnerId                     Snowflake           `json:"owner_id"`
	Permissions                 string              `json:"permissions"`
	Region                      string              `json:"region"`
	AfkChannelId                Snowflake           `json:"afk_channel_id"`
	AfkTimeout                  int                 `json:"afk_timeout"`
	WidgetEnabled               bool                `json:"widget_enabled"`
	WidgetChannelId             Snowflake           `json:"widget_channel_id"`
	VerificationLevel           int                 `json:"verification_level"`
	DefaultMessageNotifications int                 `json:"default_message_notifications"`
	ExplicitContentFilter       int                 `json:"explicit_content_filter"`
	Roles                       []RoleObject        `json:"roles"`
	Emojis                      []EmojisObject      `json:"emojis"`
	Features                    []string            `json:"features"`
	MFALevel                    int                 `json:"mfa_level"`
	ApplicationId               Snowflake           `json:"application_id"`
	SystemChannelId             Snowflake           `json:"system_channel_id"`
	SystemChannelFlags          int                 `json:"system_channel_flags"`
	RulesChannelId              Snowflake           `json:"rules_channel_id"`
	MaxPresences                int                 `json:"max_presences"`
	MaxMembers                  int                 `json:"max_members"`
	VanityUrlCode               string              `json:"vanity_url_code"`
	Description                 string              `json:"description"`
	Banner                      string              `json:"banner"`
	PremiumTier                 int                 `json:"premium_tier"`
	PremiumSubscriptionCount    int                 `json:"premium_subscription_count"`
	PreferredLocale             string              `json:"preferred_locale"`
	PublicUpdatesChannelId      Snowflake           `json:"public_updates_channel_id"`
	MaxVideoChannelUsers        int                 `json:"max_video_channel_users"`
	MaxStageVideoChannelUsers   int                 `json:"max_stage_video_channel_users"`
	ApproximateMemberCount      int                 `json:"approximate_member_count"`
	WelcomeScreen               WelcomeScreenObject `json:"welcome_screen"`
	NsfwLevel                   int                 `json:"nsfw_level"`
	Stickers                    []StickerObject     `json:"stickers"`
	PremiumProgressBarEnabled   bool                `json:"premium_progress_bar_enabled"`
	SafetyAlertsChannelId       Snowflake           `json:"safety_alerts_channel_id"`
}

type UnavailableGuildObject struct {
	Id          Snowflake `json:"id"`
	Unavailable bool      `json:"unavailble"`
}

type ApplicationObject struct {
	Id                             Snowflake           `json:"id"`
	Name                           string              `json:"name"`
	Icon                           string              `json:"icon"`
	Description                    string              `json:"description"`
	RpcOrigin                      []string            `json:"rpc_origin"`
	BotPublic                      bool                `json:"bot_public"`
	BotRequireCodeGrant            bool                `json:"bot_require_code_grant"`
	Bot                            UserObject          `json:"bot"`
	TermsOfServiceUrl              string              `json:"terms_of_service_url"`
	PrivacyPolicyUrl               string              `json:"privacy_policy_url"`
	Owner                          UserObject          `json:"owner"`
	Summary                        string              `json:"summary"` // Deprecated
	VerifyKey                      string              `json:"verify_key"`
	Team                           TeamObject          `json:"team"`
	GuildId                        Snowflake           `json:"guild_id"`
	Guild                          GuildObject         `json:"guild"`
	PrimarySkuId                   Snowflake           `json:"primary_sku_id"`
	Slug                           string              `json:"slug"`
	CoverImage                     string              `json:"cover_image"`
	Flags                          int                 `json:"flags"`
	ApproximateGuildCount          int                 `json:"approximate_guild_count"`
	RedirectUris                   []string            `json:"redirect_uris"`
	InteractionsEndpointUrl        string              `json:"interactions_endpoint_url"`
	RoleConnectionsVerificationUrl string              `json:"role_connections_verification_url"`
	Tags                           []string            `json:"tags"`
	InstallParams                  InstallParamsObject `json:"install_params"`
}

type RoleObject struct {
	Id           Snowflake     `json:"id"`
	Name         string        `json:"name"`
	Color        int           `json:"color"`
	Hoist        bool          `json:"hoist"`
	Icon         string        `json:"icon"`
	UnicodeEmoji string        `json:"unicode_emoji"`
	Position     int           `json:"position"`
	Permissions  string        `json:"permissions"`
	Managed      bool          `json:"managed"`
	Mentionable  bool          `json:"mentionable"`
	Tags         RoleTagObject `json:"tags"`
	Flags        int           `json:"flags"`
}

type RoleTagObject struct {
	BotId                 Snowflake `json:"bot_id"`
	IntergrationId        Snowflake `json:"intergration_id"`
	PremiumSubscriber     bool      `json:"premium_subscriber"`
	SubscriptionListingId Snowflake `json:"subscription_listing_id"`
	AvailableForPurchase  bool      `json:"available_for_purchase"`
	GuildConnections      bool      `json:"guild_connections"`
}

type EmojisObject struct {
	Id            Snowflake    `json:"id"`
	Name          string       `json:"name"`
	Roles         []RoleObject `json:"roles"`
	User          UserObject   `json:"user"`
	RequireColons bool         `json:"require_colons"`
	Managed       bool         `json:"managed"`
	Animated      bool         `json:"animated"`
	Available     bool         `json:"available"`
}

type WelcomeScreenObject struct {
	Description     string                 `json:"description"`
	WelcomeChannels []WelcomeScreenChannel `json:"welcome_channels"`
}

type WelcomeScreenChannel struct {
	ChannelId   Snowflake `json:"channel_id"`
	Description string    `json:"description"`
	EmojiId     Snowflake `json:"emoji_id"`
	EmojiName   string    `json:"emoji_name"`
}

type StickerObject struct {
	Id          Snowflake  `json:"id"`
	PackId      Snowflake  `json:"pack_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Tags        string     `json:"tags"`
	Asset       string     `json:"asset"`
	Type        int        `json:"type"`
	FormatType  int        `json:"format_type"`
	Available   bool       `json:"available"`
	GuildId     Snowflake  `json:"guild_id"`
	User        UserObject `json:"user"`
	SortValue   int        `json:"sort_value"`
}

type InstallParamsObject struct {
	Scopes       []string `json:"scopes"`
	Permissiosns string   `json:"permissions"`
}

type IdentifyObject struct {
	Token          string                             `json:"token"`
	Properties     IdentifyConnectionPropertiesObject `json:"properties"`
	Compress       bool                               `json:"compress,omitempty"`
	LargeThreshold int                                `json:"large_threshold"`
	Shard          []int                              `json:"shard,omitempty"`
	Presence       UpdatePresenceObject               `json:"presence,omitempty"`
	Intents        int                                `json:"intents"`
}

type IdentifyConnectionPropertiesObject struct {
	Os      string `json:"os"`
	Browser string `json:"browser"`
	Device  string `json:"device"`
}

type UpdatePresenceObject struct {
	Since     int              `json:"since"`
	Activites []ActivityObject `json:"activites"`
	Status    string           `json:"status"`
	Afk       bool             `json:"afk"`
}

type ActivityObject struct {
	Name          string                   `json:"name"`
	Type          int                      `json:"type"`
	Url           string                   `json:"url,omitempty"`
	CreatedAt     int                      `json:"created_at"`
	Timestamps    ActivityTimestampsObject `json:"timestamps"`
	ApplicationId Snowflake                `json:"application_id"`
	Details       string                   `json:"details,omitempty"`
	State         string                   `json:"state,omitempty"`
	Emoji         ActivityEmojiObject      `json:"emoji,omitempty"`
	Party         ActivityPartyObject      `json:"party,omitempty"`
	Assets        ActivityAssetsObject     `json:"assets,omitempty"`
	Secrets       ActivitySecretsObject    `json:"secrets,omitempty"`
	Instance      bool                     `json:"instance,omitempty"`
	Flags         int                      `json:"flags,omitempty"`
	Buttons       []ActivityButtonsObject  `json:"buttons,omitempty"`
}

type ActivityTimestampsObject struct {
	Start int `json:"start,omitempty"`
	End   int `json:"end,omitempty"`
}

type ActivityEmojiObject struct {
	Name     string    `json:"name"`
	Id       Snowflake `json:"id,omitempty"`
	Animated bool      `json:"animated,omitempty"`
}

type ActivityPartyObject struct {
	Id   string `json:"id,omitempty"`
	Size []int  `json:"size,omitempty"`
}

type ActivityAssetsObject struct {
	LargeImage string `json:"large_image,omitempty"`
	LargeText  string `json:"large_text,omitempty"`
	SmallImage string `json:"small_image,omitempty"`
	SmallText  string `json:"small_text,omitempty"`
}

type ActivitySecretsObject struct {
	Join     string `json:"join,omitempty"`
	Spectate string `json:"spectate,omitempty"`
	Match    string `json:"match,omitempty"`
}

type ActivityButtonsObject struct {
	Label string `json:"label"`
	Url   string `json:"url"`
}

type MessageObject struct {
	Id                   Snowflake                        `json:"id"`
	ChannelId            Snowflake                        `json:"channel_id"`
	Author               UserObject                       `json:"author"`
	Content              string                           `json:"content"`
	Timestamp            string                           `json:"timestamp"`
	EditedTimestamp      string                           `json:"edited_timestamp"`
	TTS                  bool                             `json:"tts"`
	MentionEveryone      bool                             `json:"mention_everyone"`
	Mentions             []UserObject                     `json:"mentions"`
	MentionRoles         []RoleObject                     `json:"mention_roles"`
	MentionChannels      []ChannelMentionObject           `json:"mention_channels,omitempty"`
	Attachments          []AttachmentObject               `json:"attachments"`
	Embeds               []EmbedObject                    `json:"embeds"`
	Reactions            []ReactionObject                 `json:"reactions"`
	Nonce                int                              `json:"nonce,omitempty"`
	Pinned               bool                             `json:"pinned"`
	WebhookId            Snowflake                        `json:"webhook_id,omitempty"`
	Type                 int                              `json:"type"`
	Activity             MessageActivityObject            `json:"activity,omitempty"`
	Application          ApplicationObject                `json:"application,omitempty"`
	ApplicationId        Snowflake                        `json:"application_id,omitempty"`
	MessageReference     MessageReferenceObject           `json:"message_reference,omitempty"`
	Flags                int                              `json:"flags,omitempty"`
	ReferencedMessage    MessageObject                    `json:"referenced_message,omitempty"`
	InteractionMetadata  MessageInteractionMetadataObject `json:"interaction_metadata,omitempty"`
	Interaction          MessageInteractionObject         `json:"interaction,omitempty"`
	Thread               ChannelObject                    `json:"thread,omitempty"`
	Components           []MessageComponentObject         `json:"components,omitempty"`
	StickerItems         []StickerItemObject              `json:"sticker_items,omitempty"`
	Stickers             []StickerObject                  `json:"stickers,omitempty"`
	Position             int                              `json:"position,omitempty"`
	RoleSubscriptionData RoleSubscriptionDataObject       `json:"role_subscription_data"`
	Resolved             ResolvedDataObject               `json:"resolved,omitempty"`
}

type ChannelMentionObject struct {
	Id      Snowflake `json:"id"`
	GuildId Snowflake `json:"guild_id"`
	Type    int       `json:"type"`
	Name    string    `json:"name"`
}

type AttachmentObject struct {
	Id           Snowflake `json:"id"`
	Filename     string    `json:"filename"`
	Description  string    `json:"description,omitempty"`
	ContentType  string    `json:"content_type,omitempty"`
	Size         int       `json:"size"`
	Url          string    `json:"url"`
	ProxyUrl     string    `json:"proxy_url"`
	Height       int       `json:"height,omitempty"`
	Width        int       `json:"width,omitempty"`
	Ephermeral   bool      `json:"ephemeral,omitempty"`
	DurationSecs float32   `json:"duration_secs,omitempty"`
	Waveform     string    `json:"waveform,omitempty"`
	Flags        int       `json:"flags,omitempty"`
}

type EmbedObject struct {
	Title       string              `json:"title,omitempty"`
	Type        string              `json:"type,omitempty"`
	Description string              `json:"description,omitempty"`
	Url         string              `json:"url,omitempty"`
	Timestamp   string              `json:"timestamp,omitempty"`
	Color       int                 `json:"color,omitempty"`
	Footer      EmbedFooterObject   `json:"footer,omitempty"`
	Image       EmbedImageObject    `json:"image,omitempty"`
	Thumbnail   EmbedThumnailObject `json:"thumbnail,omitempty"`
	Video       EmbedVideoObject    `json:"video,omitempty"`
	Provider    EmbedProviderObject `json:"provider"omiteempty"`
	Author      EmbedAuthorObject   `json:"author,omitempty"`
	Fields      []EmbedFieldObject  `json:"fields,omitempty"`
}

type EmbedFooterObject struct {
	Text         string `json:"text"`
	IconUrl      string `json:"icon_url"`
	ProxyIconUrl string `json:"proxy_icon_url"`
}

type EmbedImageObject struct {
	Url      string `json:"url"`
	ProxyUrl string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedThumnailObject struct {
	Url      string `json:"url"`
	ProxyUrl string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedVideoObject struct {
	Url      string `json:"url,omitempty"`
	ProxyUrl string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedProviderObject struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type EmbedAuthorObject struct {
	Name         string `json:"name"`
	Url          string `json:"url,omitempty"`
	IconUrl      string `json:"icon_url,omitempty"`
	ProxyIconUrl string `json:"proxy_icon_url,omitempty"`
}

type EmbedFieldObject struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type ReactionObject struct {
	Count        int         `json:"count"`
	CountDetails interface{} `json:"count_details"`
	Me           bool        `json:"me"`
	MeBurst      bool        `json:"me_burst"`
	Emoji        EmojiObject `json:"emoji"`
	BurstColors  []int       `json:"burst_colors"`
}

type ReactionCountDetailsObject struct {
	Burst  int `json:"burst"`
	Normal int `jsonh:"normal"`
}

type EmojiObject struct {
	Id            Snowflake    `json:"id"`
	Name          string       `json:"name"`
	Roles         []RoleObject `json:"role,omitempty"`
	User          UserObject   `json:"user,omitempty"`
	RequireColons bool         `json:"require_colons,omitempty"`
	Managed       bool         `json:"managed,omitempty"`
	Animated      bool         `json:"animated,omitempty"`
	Available     bool         `json:"available,omitempty"`
}

type MessageActivityObject struct {
	Type    int    `json:"type"`
	PartyId string `json:"party_id,omitempty"`
}

type MessageReferenceObject struct {
	MessageId       Snowflake `json:"message_id,omitempty"`
	ChannelId       Snowflake `json:"channel_id,omitempty"`
	GuildId         Snowflake `json:"guild_id,omitempty"`
	FailIfNotExists bool      `json:"fail_if_not_exists,omitempty"`
}

type MessageInteractionMetadataObject struct {
	Id                            Snowflake                        `json:"id"`
	Type                          int                              `json:"type"`
	UserId                        Snowflake                        `json:"user_id"`
	AuthorizingIntergrationOwners map[string]int                   `json:"authorizing_intergration_owners"`
	OriginalResponseMessageId     Snowflake                        `json:"original_response_message_id,omitempty"`
	InteractedMessageId           Snowflake                        `json:"interacted_message_id,omitempty"`
	TriggeringInteractionMetadata MessageInteractionMetadataObject `json:"triggering_interaction_metadata,omitempty"`
}

type MessageInteractionObject struct {
	Id     Snowflake         `json:"id"`
	Type   int               `json:"type"`
	Name   string            `json:"name"`
	User   UserObject        `json:"user"`
	Member GuildMemberObject `json:"member,omitempty"`
}

type GuildMemberObject struct {
	User                       UserObject  `json:"user,omitempty"`
	Nick                       string      `json:"nick,omitempty"`
	Avatar                     string      `json:"avatar,omitempty"`
	Roles                      []Snowflake `json:"roles"`
	JoinedAt                   string      `json:"joined_at"`
	PremiumSince               string      `json:"premium_since,omitempty"`
	Deaf                       bool        `json:"deaf"`
	Mute                       bool        `json:"mute"`
	Flags                      int         `json:"flags"`
	Pending                    bool        `json:"pending,omitempty"`
	Permissions                string      `json:"permissions,omitempty"`
	CommunicationDisabledUntil string      `json:"communication_disabled_until,omitempty"`
}

type ChannelObject struct {
	Id                            Snowflake             `json:"id"`
	Type                          int                   `json:"type"`
	GuildId                       Snowflake             `json:"guild_id,omitempty"`
	Position                      int                   `json:"position,omitempty"`
	PermissionOverwrite           []OverwriteObject     `json:"permission_overwrite"`
	Name                          string                `json:"name,omitempty"`
	Topic                         string                `json:"topic,omitempty"`
	Nswf                          bool                  `json:"nswf,omitempty"`
	LastMessageId                 Snowflake             `json:"last_message_id,omitempty"`
	Bitrate                       int                   `json:"bitrate,omitempty"`
	UserLimit                     int                   `json:"user_limit,omitempty"`
	RateLimitPerUser              int                   `json:"rate_limit_per_user,omitempty"`
	Recipients                    []UserObject          `json:"recipients,omitempty"`
	Icon                          string                `json:"icon,omitempty"`
	OwnerId                       Snowflake             `json:"owner_id,omitempty"`
	ApplicationId                 Snowflake             `json:"application_id,omitempty"`
	Managed                       bool                  `json:"managed,omitempty"`
	ParentId                      Snowflake             `json:"parent_id,omitempty"`
	LastPinTimestamp              string                `json:"last_pin_timestampe,omitempty"`
	RTCRegion                     string                `json:"rtc_region,omitempty"`
	VideoQualityMode              int                   `json:"video_quality_mode,omitempty"`
	MessageCount                  int                   `json:"message_count,omitempty"`
	MemberCount                   int                   `json:"member_count,omitempty"`
	ThreadMetadata                ThreadMetadataObject  `json:"thread_metadata,omitempty"`
	Member                        ThreadMemberObject    `json:"member,omitempty"`
	DefaultAutoArchiveDuration    int                   `json:"default_auto_archive_duration,omitempty"`
	Permissions                   string                `json:"permissions,omitempty"`
	Flags                         int                   `json:"flags,omitempty"`
	TotalMessageSent              int                   `json:"totla_message_sent,omitempty"`
	AvailableTags                 []TagObject           `json:"availble_tags,omitempty"`
	AppliedTags                   []Snowflake           `json:"applied_tags"`
	DefaultReactionEmoji          DefaultReactionObject `json:"default_reaction_emoji,omitempty"`
	DefaultThreadRateLimitPerUser int                   `json:"default_thread_rate_limit_per_user,omitempty"`
	DefaultSortOrder              int                   `json:"default_sort_order,omitempty"`
	DefaultForumLayout            int                   `json:"default_forum_layout,omitempty"`
}

type OverwriteObject struct {
	Id    Snowflake `json:"id"`
	Type  int       `json:"type"`
	Allow string    `json:"allow"`
	Deny  string    `json:"deny"`
}

type ThreadMetadataObject struct {
	Archived            bool   `json:"archived"`
	AutoArchiveDuration int    `json:"auto_archive_duration"`
	ArchiveTimestamp    string `json:"archive_timestamp"`
	Locked              bool   `json:"locked"`
	Invitable           bool   `json:"invitable,omitempty"`
	CreateTimestamp     string `json:"create_timestamp,omitempty"`
}

type ThreadMemberObject struct {
	Id            Snowflake         `json:"id,omitempty"`
	UserId        Snowflake         `json:"user_id,omitempty"`
	JoinTimestamp string            `json:"join_timestamp"`
	Flags         int               `json:"flags"`
	Memberw       GuildMemberObject `json:"member,omitempty"`
}

type TagObject struct {
	Id        Snowflake `json:"id"`
	Name      string    `json:"name"`
	Moderated bool      `json:"moderated"`
	EmojiId   Snowflake `json:"emoji_id"`
	EmojiName string    `json:"emoji_name"`
}

type DefaultReactionObject struct {
	EmojiId   Snowflake `json:"emoji_id"`
	EmojiName string    `json:"emoji_name"`
}

type MessageComponentObject interface{}

type StickerItemObject struct {
	Id     Snowflake `json:"id"`
	Name   string    `json:"name"`
	Format int       `json:"format_type"`
}

type RoleSubscriptionDataObject struct {
	RoleSubscriptionListingId Snowflake `json:"role_subscription_listing_id"`
	TierName                  string    `json:"tier_name"`
	TotalMonthsSubscribed     int       `json:"total_months_subscribed"`
	IsRenewal                 bool      `json:"is_renewal"`
}

// Gateway stuff
type GatewayEventPayload struct {
	Op int         `json:"op"` // Gateway Opcode
	D  interface{} `json:"d"`  // Event Data
	S  int         `json:"s"`  // Sequence number
	T  string      `json:"t"`  // Event name
}

type HelloEvent struct {
	HeartbeatInterval int `json:"heartbeat_interval"`
}

type ReadyEvent struct {
	V                int                      `json:"v"`
	User             UserObject               `json:"user"`
	Guilds           []UnavailableGuildObject `json:"guilds"`
	SessionId        string                   `json:"session_id"`
	ResumeGatewayUrl string                   `json:"resume_gateway_url"`
	Shard            []int                    `json:"shard,omitempty"`
	Application      ApplicationObject        `json:"application"`
}

type ResumeEvent struct {
	Token     string `json:"token"`
	SessionId string `json:"session_id"`
	Seq       int    `json:"seq"`
}

// HTTP Requests
func (this Discord) GetGatewayBot() (*http.Response, error) {
	endpoint := DISCORD_URL + "/gateway/bot"

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)

	if err != nil {
		return nil, err
	}

	return this.MakeRequest(request)
}

// General HTTP request for discord]
func (this Discord) MakeRequest(request *http.Request) (*http.Response, error) {
	request.Header["Authorization"] = []string{fmt.Sprintf("Bot %s", this.token)}
	request.Header["User-Agent"] = []string{"DiscordBot (Fat Chocobo, 0)"}

	client := http.Client{}
	return client.Do(request)
}

// Send websocket events
func (this Discord) SendHeartbeat(seq int) {
	payload := &GatewayEventPayload{
		Op: GATEWAY_OPCODE_HEARTBEAT,
		D:  seq,
	}

	this.Websocket.WriteJSON(payload)
}

func (this Discord) InitGatewayHandshake(intents int) {
	connectionProperties := IdentifyConnectionPropertiesObject{
		Os:      "linux",
		Browser: "Fat Chocobo",
		Device:  "Fat Chocobo",
	}

	data := &IdentifyObject{
		Token:      this.token,
		Properties: connectionProperties,
		Intents:    intents,
	}

	payload := &GatewayEventPayload{
		Op: GATEWAY_OPCODE_IDENTIFY,
		D:  data,
	}

	this.Websocket.WriteJSON(payload)
}

// Constructor
func CreateDiscord(token string) *Discord {
	discord := new(Discord)
	discord.token = token

	return discord
}

// Parsing through HTTP respones
func GetDiscordGatewayBot(discord *Discord) (*GatewayBotResponse, error) {
	data := new(GatewayBotResponse)

	response, err := discord.GetGatewayBot()

	if err != nil {
		return nil, err
	} else if response.StatusCode != 200 {
		return nil, fmt.Errorf("Bad status code: %s", response.Status)
	}

	raw, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(raw, data)
	defer response.Body.Close()

	return data, err
}
