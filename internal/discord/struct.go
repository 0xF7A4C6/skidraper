package discord

import "time"

type Messages []struct {
	ID        string `json:"id"`
	Type      int    `json:"type"`
	Content   string `json:"content"`
	ChannelID string `json:"channel_id"`
	Author    struct {
		ID               string `json:"id"`
		Username         string `json:"username"`
		GlobalName       string `json:"global_name"`
		Avatar           string `json:"avatar"`
		Discriminator    string `json:"discriminator"`
		PublicFlags      int    `json:"public_flags"`
		AvatarDecoration any    `json:"avatar_decoration"`
	} `json:"author"`
	Attachments      []any     `json:"attachments"`
	Embeds           []any     `json:"embeds"`
	Mentions         []any     `json:"mentions"`
	MentionRoles     []any     `json:"mention_roles"`
	Pinned           bool      `json:"pinned"`
	MentionEveryone  bool      `json:"mention_everyone"`
	Tts              bool      `json:"tts"`
	Timestamp        time.Time `json:"timestamp"`
	EditedTimestamp  any       `json:"edited_timestamp"`
	Flags            int       `json:"flags"`
	Components       []any     `json:"components"`
	MessageReference struct {
		ChannelID string `json:"channel_id"`
		GuildID   string `json:"guild_id"`
		MessageID string `json:"message_id"`
	} `json:"message_reference,omitempty"`
	ReferencedMessage struct {
		ID        string `json:"id"`
		Type      int    `json:"type"`
		Content   string `json:"content"`
		ChannelID string `json:"channel_id"`
		Author    struct {
			ID               string `json:"id"`
			Username         string `json:"username"`
			GlobalName       any    `json:"global_name"`
			Avatar           any    `json:"avatar"`
			Discriminator    string `json:"discriminator"`
			PublicFlags      int    `json:"public_flags"`
			AvatarDecoration any    `json:"avatar_decoration"`
		} `json:"author"`
		Attachments     []any     `json:"attachments"`
		Embeds          []any     `json:"embeds"`
		Mentions        []any     `json:"mentions"`
		MentionRoles    []any     `json:"mention_roles"`
		Pinned          bool      `json:"pinned"`
		MentionEveryone bool      `json:"mention_everyone"`
		Tts             bool      `json:"tts"`
		Timestamp       time.Time `json:"timestamp"`
		EditedTimestamp any       `json:"edited_timestamp"`
		Flags           int       `json:"flags"`
		Components      []any     `json:"components"`
	} `json:"referenced_message,omitempty"`
	ReferencedMessage0 struct {
		ID        string `json:"id"`
		Type      int    `json:"type"`
		Content   string `json:"content"`
		ChannelID string `json:"channel_id"`
		Author    struct {
			ID               string `json:"id"`
			Username         string `json:"username"`
			GlobalName       any    `json:"global_name"`
			Avatar           string `json:"avatar"`
			Discriminator    string `json:"discriminator"`
			PublicFlags      int    `json:"public_flags"`
			AvatarDecoration any    `json:"avatar_decoration"`
		} `json:"author"`
		Attachments []any `json:"attachments"`
		Embeds      []any `json:"embeds"`
		Mentions    []struct {
			ID               string `json:"id"`
			Username         string `json:"username"`
			Avatar           any    `json:"avatar"`
			Discriminator    string `json:"discriminator"`
			PublicFlags      int    `json:"public_flags"`
			Flags            int    `json:"flags"`
			Banner           any    `json:"banner"`
			AccentColor      any    `json:"accent_color"`
			GlobalName       any    `json:"global_name"`
			AvatarDecoration any    `json:"avatar_decoration"`
			DisplayName      any    `json:"display_name"`
			BannerColor      any    `json:"banner_color"`
		} `json:"mentions"`
		MentionRoles     []any     `json:"mention_roles"`
		Pinned           bool      `json:"pinned"`
		MentionEveryone  bool      `json:"mention_everyone"`
		Tts              bool      `json:"tts"`
		Timestamp        time.Time `json:"timestamp"`
		EditedTimestamp  any       `json:"edited_timestamp"`
		Flags            int       `json:"flags"`
		Components       []any     `json:"components"`
		MessageReference struct {
			ChannelID string `json:"channel_id"`
			GuildID   string `json:"guild_id"`
			MessageID string `json:"message_id"`
		} `json:"message_reference"`
	} `json:"referenced_message,omitempty"`
	ReferencedMessage1 struct {
		ID        string `json:"id"`
		Type      int    `json:"type"`
		Content   string `json:"content"`
		ChannelID string `json:"channel_id"`
		Author    struct {
			ID               string `json:"id"`
			Username         string `json:"username"`
			GlobalName       any    `json:"global_name"`
			Avatar           any    `json:"avatar"`
			Discriminator    string `json:"discriminator"`
			PublicFlags      int    `json:"public_flags"`
			AvatarDecoration any    `json:"avatar_decoration"`
		} `json:"author"`
		Attachments []any `json:"attachments"`
		Embeds      []any `json:"embeds"`
		Mentions    []struct {
			ID               string `json:"id"`
			Username         string `json:"username"`
			Avatar           string `json:"avatar"`
			Discriminator    string `json:"discriminator"`
			PublicFlags      int    `json:"public_flags"`
			Flags            int    `json:"flags"`
			Banner           any    `json:"banner"`
			AccentColor      any    `json:"accent_color"`
			GlobalName       any    `json:"global_name"`
			AvatarDecoration any    `json:"avatar_decoration"`
			DisplayName      any    `json:"display_name"`
			BannerColor      any    `json:"banner_color"`
		} `json:"mentions"`
		MentionRoles     []any     `json:"mention_roles"`
		Pinned           bool      `json:"pinned"`
		MentionEveryone  bool      `json:"mention_everyone"`
		Tts              bool      `json:"tts"`
		Timestamp        time.Time `json:"timestamp"`
		EditedTimestamp  time.Time `json:"edited_timestamp"`
		Flags            int       `json:"flags"`
		Components       []any     `json:"components"`
		MessageReference struct {
			ChannelID string `json:"channel_id"`
			GuildID   string `json:"guild_id"`
			MessageID string `json:"message_id"`
		} `json:"message_reference"`
	} `json:"referenced_message,omitempty"`
	ReferencedMessage2 struct {
		ID        string `json:"id"`
		Type      int    `json:"type"`
		Content   string `json:"content"`
		ChannelID string `json:"channel_id"`
		Author    struct {
			ID               string `json:"id"`
			Username         string `json:"username"`
			GlobalName       any    `json:"global_name"`
			Avatar           any    `json:"avatar"`
			Discriminator    string `json:"discriminator"`
			PublicFlags      int    `json:"public_flags"`
			AvatarDecoration any    `json:"avatar_decoration"`
		} `json:"author"`
		Attachments []any `json:"attachments"`
		Embeds      []any `json:"embeds"`
		Mentions    []struct {
			ID               string `json:"id"`
			Username         string `json:"username"`
			Avatar           string `json:"avatar"`
			Discriminator    string `json:"discriminator"`
			PublicFlags      int    `json:"public_flags"`
			Flags            int    `json:"flags"`
			Banner           any    `json:"banner"`
			AccentColor      any    `json:"accent_color"`
			GlobalName       any    `json:"global_name"`
			AvatarDecoration any    `json:"avatar_decoration"`
			DisplayName      any    `json:"display_name"`
			BannerColor      any    `json:"banner_color"`
		} `json:"mentions"`
		MentionRoles     []any     `json:"mention_roles"`
		Pinned           bool      `json:"pinned"`
		MentionEveryone  bool      `json:"mention_everyone"`
		Tts              bool      `json:"tts"`
		Timestamp        time.Time `json:"timestamp"`
		EditedTimestamp  time.Time `json:"edited_timestamp"`
		Flags            int       `json:"flags"`
		Components       []any     `json:"components"`
		MessageReference struct {
			ChannelID string `json:"channel_id"`
			GuildID   string `json:"guild_id"`
			MessageID string `json:"message_id"`
		} `json:"message_reference"`
	} `json:"referenced_message,omitempty"`
	ReferencedMessage3 struct {
		ID        string `json:"id"`
		Type      int    `json:"type"`
		Content   string `json:"content"`
		ChannelID string `json:"channel_id"`
		Author    struct {
			ID               string `json:"id"`
			Username         string `json:"username"`
			GlobalName       any    `json:"global_name"`
			Avatar           string `json:"avatar"`
			Discriminator    string `json:"discriminator"`
			PublicFlags      int    `json:"public_flags"`
			AvatarDecoration any    `json:"avatar_decoration"`
		} `json:"author"`
		Attachments []any `json:"attachments"`
		Embeds      []any `json:"embeds"`
		Mentions    []struct {
			ID               string `json:"id"`
			Username         string `json:"username"`
			Avatar           any    `json:"avatar"`
			Discriminator    string `json:"discriminator"`
			PublicFlags      int    `json:"public_flags"`
			Flags            int    `json:"flags"`
			Banner           any    `json:"banner"`
			AccentColor      any    `json:"accent_color"`
			GlobalName       any    `json:"global_name"`
			AvatarDecoration any    `json:"avatar_decoration"`
			DisplayName      any    `json:"display_name"`
			BannerColor      any    `json:"banner_color"`
		} `json:"mentions"`
		MentionRoles     []any     `json:"mention_roles"`
		Pinned           bool      `json:"pinned"`
		MentionEveryone  bool      `json:"mention_everyone"`
		Tts              bool      `json:"tts"`
		Timestamp        time.Time `json:"timestamp"`
		EditedTimestamp  any       `json:"edited_timestamp"`
		Flags            int       `json:"flags"`
		Components       []any     `json:"components"`
		MessageReference struct {
			ChannelID string `json:"channel_id"`
			GuildID   string `json:"guild_id"`
			MessageID string `json:"message_id"`
		} `json:"message_reference"`
	} `json:"referenced_message,omitempty"`
	ReferencedMessage4 struct {
		ID        string `json:"id"`
		Type      int    `json:"type"`
		Content   string `json:"content"`
		ChannelID string `json:"channel_id"`
		Author    struct {
			ID               string `json:"id"`
			Username         string `json:"username"`
			GlobalName       string `json:"global_name"`
			Avatar           string `json:"avatar"`
			Discriminator    string `json:"discriminator"`
			PublicFlags      int    `json:"public_flags"`
			AvatarDecoration string `json:"avatar_decoration"`
		} `json:"author"`
		Attachments []any `json:"attachments"`
		Embeds      []any `json:"embeds"`
		Mentions    []struct {
			ID               string `json:"id"`
			Username         string `json:"username"`
			Avatar           string `json:"avatar"`
			Discriminator    string `json:"discriminator"`
			PublicFlags      int    `json:"public_flags"`
			Flags            int    `json:"flags"`
			Banner           any    `json:"banner"`
			AccentColor      any    `json:"accent_color"`
			GlobalName       string `json:"global_name"`
			AvatarDecoration any    `json:"avatar_decoration"`
			DisplayName      string `json:"display_name"`
			BannerColor      any    `json:"banner_color"`
		} `json:"mentions"`
		MentionRoles     []any     `json:"mention_roles"`
		Pinned           bool      `json:"pinned"`
		MentionEveryone  bool      `json:"mention_everyone"`
		Tts              bool      `json:"tts"`
		Timestamp        time.Time `json:"timestamp"`
		EditedTimestamp  any       `json:"edited_timestamp"`
		Flags            int       `json:"flags"`
		Components       []any     `json:"components"`
		MessageReference struct {
			ChannelID string `json:"channel_id"`
			GuildID   string `json:"guild_id"`
			MessageID string `json:"message_id"`
		} `json:"message_reference"`
	} `json:"referenced_message,omitempty"`
	ReferencedMessage5 struct {
		ID        string `json:"id"`
		Type      int    `json:"type"`
		Content   string `json:"content"`
		ChannelID string `json:"channel_id"`
		Author    struct {
			ID               string `json:"id"`
			Username         string `json:"username"`
			GlobalName       any    `json:"global_name"`
			Avatar           any    `json:"avatar"`
			Discriminator    string `json:"discriminator"`
			PublicFlags      int    `json:"public_flags"`
			AvatarDecoration any    `json:"avatar_decoration"`
		} `json:"author"`
		Attachments []any `json:"attachments"`
		Embeds      []any `json:"embeds"`
		Mentions    []struct {
			ID               string `json:"id"`
			Username         string `json:"username"`
			Avatar           string `json:"avatar"`
			Discriminator    string `json:"discriminator"`
			PublicFlags      int    `json:"public_flags"`
			Flags            int    `json:"flags"`
			Banner           any    `json:"banner"`
			AccentColor      any    `json:"accent_color"`
			GlobalName       any    `json:"global_name"`
			AvatarDecoration any    `json:"avatar_decoration"`
			DisplayName      any    `json:"display_name"`
			BannerColor      any    `json:"banner_color"`
		} `json:"mentions"`
		MentionRoles     []any     `json:"mention_roles"`
		Pinned           bool      `json:"pinned"`
		MentionEveryone  bool      `json:"mention_everyone"`
		Tts              bool      `json:"tts"`
		Timestamp        time.Time `json:"timestamp"`
		EditedTimestamp  any       `json:"edited_timestamp"`
		Flags            int       `json:"flags"`
		Components       []any     `json:"components"`
		MessageReference struct {
			ChannelID string `json:"channel_id"`
			GuildID   string `json:"guild_id"`
			MessageID string `json:"message_id"`
		} `json:"message_reference"`
	} `json:"referenced_message,omitempty"`
	ReferencedMessage6 struct {
		ID        string `json:"id"`
		Type      int    `json:"type"`
		Content   string `json:"content"`
		ChannelID string `json:"channel_id"`
		Author    struct {
			ID               string `json:"id"`
			Username         string `json:"username"`
			GlobalName       any    `json:"global_name"`
			Avatar           string `json:"avatar"`
			Discriminator    string `json:"discriminator"`
			PublicFlags      int    `json:"public_flags"`
			AvatarDecoration any    `json:"avatar_decoration"`
		} `json:"author"`
		Attachments []any `json:"attachments"`
		Embeds      []any `json:"embeds"`
		Mentions    []struct {
			ID               string `json:"id"`
			Username         string `json:"username"`
			Avatar           string `json:"avatar"`
			Discriminator    string `json:"discriminator"`
			PublicFlags      int    `json:"public_flags"`
			Flags            int    `json:"flags"`
			Banner           any    `json:"banner"`
			AccentColor      any    `json:"accent_color"`
			GlobalName       string `json:"global_name"`
			AvatarDecoration string `json:"avatar_decoration"`
			DisplayName      string `json:"display_name"`
			BannerColor      any    `json:"banner_color"`
		} `json:"mentions"`
		MentionRoles     []any     `json:"mention_roles"`
		Pinned           bool      `json:"pinned"`
		MentionEveryone  bool      `json:"mention_everyone"`
		Tts              bool      `json:"tts"`
		Timestamp        time.Time `json:"timestamp"`
		EditedTimestamp  any       `json:"edited_timestamp"`
		Flags            int       `json:"flags"`
		Components       []any     `json:"components"`
		MessageReference struct {
			ChannelID string `json:"channel_id"`
			GuildID   string `json:"guild_id"`
			MessageID string `json:"message_id"`
		} `json:"message_reference"`
	} `json:"referenced_message,omitempty"`
	ReferencedMessage7 struct {
		ID        string `json:"id"`
		Type      int    `json:"type"`
		Content   string `json:"content"`
		ChannelID string `json:"channel_id"`
		Author    struct {
			ID               string `json:"id"`
			Username         string `json:"username"`
			GlobalName       any    `json:"global_name"`
			Avatar           string `json:"avatar"`
			Discriminator    string `json:"discriminator"`
			PublicFlags      int    `json:"public_flags"`
			AvatarDecoration any    `json:"avatar_decoration"`
		} `json:"author"`
		Attachments []any `json:"attachments"`
		Embeds      []any `json:"embeds"`
		Mentions    []struct {
			ID               string `json:"id"`
			Username         string `json:"username"`
			Avatar           any    `json:"avatar"`
			Discriminator    string `json:"discriminator"`
			PublicFlags      int    `json:"public_flags"`
			Flags            int    `json:"flags"`
			Banner           any    `json:"banner"`
			AccentColor      any    `json:"accent_color"`
			GlobalName       any    `json:"global_name"`
			AvatarDecoration any    `json:"avatar_decoration"`
			DisplayName      any    `json:"display_name"`
			BannerColor      any    `json:"banner_color"`
		} `json:"mentions"`
		MentionRoles     []any     `json:"mention_roles"`
		Pinned           bool      `json:"pinned"`
		MentionEveryone  bool      `json:"mention_everyone"`
		Tts              bool      `json:"tts"`
		Timestamp        time.Time `json:"timestamp"`
		EditedTimestamp  time.Time `json:"edited_timestamp"`
		Flags            int       `json:"flags"`
		Components       []any     `json:"components"`
		MessageReference struct {
			ChannelID string `json:"channel_id"`
			GuildID   string `json:"guild_id"`
			MessageID string `json:"message_id"`
		} `json:"message_reference"`
	} `json:"referenced_message,omitempty"`
	ReferencedMessage8 struct {
		ID        string `json:"id"`
		Type      int    `json:"type"`
		Content   string `json:"content"`
		ChannelID string `json:"channel_id"`
		Author    struct {
			ID               string `json:"id"`
			Username         string `json:"username"`
			GlobalName       string `json:"global_name"`
			Avatar           string `json:"avatar"`
			Discriminator    string `json:"discriminator"`
			PublicFlags      int    `json:"public_flags"`
			AvatarDecoration string `json:"avatar_decoration"`
		} `json:"author"`
		Attachments []any `json:"attachments"`
		Embeds      []any `json:"embeds"`
		Mentions    []struct {
			ID               string `json:"id"`
			Username         string `json:"username"`
			Avatar           string `json:"avatar"`
			Discriminator    string `json:"discriminator"`
			PublicFlags      int    `json:"public_flags"`
			Flags            int    `json:"flags"`
			Banner           any    `json:"banner"`
			AccentColor      any    `json:"accent_color"`
			GlobalName       any    `json:"global_name"`
			AvatarDecoration any    `json:"avatar_decoration"`
			DisplayName      any    `json:"display_name"`
			BannerColor      any    `json:"banner_color"`
		} `json:"mentions"`
		MentionRoles     []any     `json:"mention_roles"`
		Pinned           bool      `json:"pinned"`
		MentionEveryone  bool      `json:"mention_everyone"`
		Tts              bool      `json:"tts"`
		Timestamp        time.Time `json:"timestamp"`
		EditedTimestamp  any       `json:"edited_timestamp"`
		Flags            int       `json:"flags"`
		Components       []any     `json:"components"`
		MessageReference struct {
			ChannelID string `json:"channel_id"`
			GuildID   string `json:"guild_id"`
			MessageID string `json:"message_id"`
		} `json:"message_reference"`
	} `json:"referenced_message,omitempty"`
}