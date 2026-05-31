package krunkgo

type RankedProfile struct {
	RankedRegion     int32 `json:"region"`
	RankedMMR        int32 `json:"mmr"`
	RankedWins       int32 `json:"wins"`
	RankedLosses     int32 `json:"losses"`
	RankedKills      int32 `json:"kills"`
	RankedDeaths     int32 `json:"deaths"`
	RankedAssists    int32 `json:"assists"`
	RankedScore      int64 `json:"score"`
	RankedDamageDone int64 `json:"damage_done"`
	RankedTimePlayed int64 `json:"time_played"`
}

type Player struct {
	PlayerName              string          `json:"player_name"` // No rename macro in Rust; keys match directly
	PlayerClan              string          `json:"clan"`
	PlayerVerified          bool            `json:"verified"`
	PlayerFlag              int32           `json:"flag"`
	PlayerBadges            []int32         `json:"badges"`
	PlayerFollowing         int32           `json:"following"`
	PlayerFollowers         int32           `json:"followers"`
	PlayerRanked            []RankedProfile `json:"ranked"`
	PlayerKR                uint64          `json:"kr"`
	PlayerLevel             int32           `json:"level"`
	PlayerJunk              float64         `json:"junk"`
	PlayerInventory         uint64          `json:"inventory"`
	PlayerScore             uint64          `json:"score"`
	PlayerSPK               float64         `json:"spk"`
	PlayerKills             uint64          `json:"kills"`
	PlayerDeaths            uint64          `json:"deaths"`
	PlayerKDR               float64         `json:"kdr"`
	PlayerKPG               float64         `json:"kpg"`
	PlayerGames             int32           `json:"games"`
	PlayerWins              int32           `json:"wins"`
	PlayerLosses            int32           `json:"losses"`
	PlayerAssists           int32           `json:"assists"`
	PlayerMelees            int32           `json:"melees"`
	PlayerBeatdowns         int32           `json:"beatdowns"`
	PlayerBullseyes         int32           `json:"bullseyes"`
	PlayerHeadshots         int32           `json:"headshots"`
	PlayerLegshots          int32           `json:"legshots"`
	PlayerWallbangs         int32           `json:"wallbangs"`
	PlayerShots             uint64          `json:"shots"`
	PlayerHits              uint64          `json:"hits"`
	PlayerMisses            uint64          `json:"misses"`
	PlayerTimePlayed        int64           `json:"time_played"`
	PlayerNukes             int32           `json:"nukes"`
	PlayerAirdrops          int32           `json:"airdrops"`
	PlayerAirdropsStolen    int32           `json:"airdrops_stolen"`
	PlayerSlimes            int32           `json:"slimes"`
	PlayerJuggernauts       int32           `json:"juggernauts"`
	PlayerJuggernautsKilled int32           `json:"juggernauts_killed"`
	PlayerWarmachines       int32           `json:"warmachines"`
	PlayerHackerTagged      bool            `json:"hacker_tagged"`
	PlayerCreatedAt         string          `json:"created_at"`
}

type InventoryItem struct {
	InventorySkinIndex int32 `json:"skin_index"`
	InventoryCount     int32 `json:"count"`
}

type PlayerMatch struct {
	PMMatchId        int64  `json:"match_id"`
	PMDate           string `json:"date"`
	PMMap            int32  `json:"map"`
	PMDuration       int64  `json:"duration"`
	PMSeason         int32  `json:"season"`
	PMRegion         int32  `json:"region"`
	PMKills          int64  `json:"kills"`
	PMDeaths         int64  `json:"deaths"`
	PMAssists        int64  `json:"assists"`
	PMScore          int64  `json:"score"`
	PMDamageDone     int64  `json:"damage_done"`
	PMHeadshots      int32  `json:"headshots"`
	PMAccuracy       int32  `json:"accuracy"`
	PMObjectiveScore int32  `json:"objective_score"`
	PMKR             int32  `json:"kr"`
	PMVictory        int32  `json:"victory"`
	PMRoundsWon      int32  `json:"rounds_won"`
	PMTeam           int32  `json:"team"`
	PMPlayTime       int64  `json:"play_time"`
}

type PlayerMatchesResponse struct {
	PMRPage    int32         `json:"page"`
	PMRPerPage int32         `json:"per_page"`
	PMRMatches []PlayerMatch `json:"matches"`
}

type Post struct {
	PostDate         string `json:"date"`
	PostText         string `json:"text"`
	PostVotes        int32  `json:"votes"`
	PostCommentCount int32  `json:"comment_count"`
}

type PostsResponse struct {
	PostsPage    int32  `json:"page"`
	PostsPerPage int32  `json:"per_page"`
	PostsPosts   []Post `json:"posts"`
}

type MatchParticipant struct {
	MPPlayerName     string `json:"player_name"`
	MPKills          int32  `json:"kills"`
	MPDeaths         int32  `json:"deaths"`
	MPAssists        int32  `json:"assists"`
	MPScore          int64  `json:"score"`
	MPDamageDone     int64  `json:"damage_done"`
	MPHeadshots      int32  `json:"headshots"`
	MPAccuracy       int32  `json:"accuracy"`
	MPObjectiveScore int32  `json:"objective_score"`
	MPVictory        int32  `json:"victory"`
	MPRoundsWon      int32  `json:"rounds_won"`
	MPTeam           int32  `json:"team"`
	MPPlayTime       int64  `json:"play_time"`
}

type Match struct {
	MatchId           int64              `json:"match_id"`
	MatchDate         string             `json:"date"`
	MatchMap          int32              `json:"map"`
	MatchDuration     int32              `json:"duration"`
	MatchSeason       int32              `json:"season"`
	MatchRegion       int32              `json:"region"`
	MatchParticipants []MatchParticipant `json:"participants"`
}

type Clan struct {
	ClanName        string `json:"name"`
	ClanOwnerName   string `json:"owner_name"`
	ClanScore       int64  `json:"score"`
	ClanRank        int32  `json:"rank"`
	ClanMemberCount int32  `json:"member_count"`
	ClanCreatedAt   string `json:"created_at"`
	ClanDiscord     string `json:"discord"`
}

type ClanMember struct {
	CMPlayerName string `json:"player_name"`
	CMRole       int32  `json:"role"`
}

type ClanMembersResponse struct {
	CMRPage    int32        `json:"page"`
	CMRPerPage int32        `json:"per_page"`
	CMRMembers []ClanMember `json:"members"`
}

type LeaderboardEntry struct {
	LEPosition   int32  `json:"position"`
	LEPlayerName string `json:"player_name"`
	LEMMR        int32  `json:"mmr"`
	LEWins       int32  `json:"wins"`
	LELosses     int32  `json:"losses"`
	LEKills      int32  `json:"kills"`
	LEDeaths     int32  `json:"deaths"`
	LEAssists    int32  `json:"assists"`
	LEScore      int64  `json:"score"`
	LEDamageDone int64  `json:"damage_done"`
}

type LeaderboardResponse struct {
	LRPage    int32              `json:"page"`
	LRPerPage int32              `json:"per_page"`
	LRSeason  int32              `json:"season"`
	LRRegion  int32              `json:"region"`
	LREntries []LeaderboardEntry `json:"entries"`
}

type GameMap struct {
	GMMapId            int32  `json:"map_id"`
	GMName             string `json:"name"`
	GMDescription      string `json:"description"`
	GMCreatorName      string `json:"creator_name"`
	GMVotes            int32  `json:"votes"`
	GMGameplays        int32  `json:"gameplays"`
	GMPlaytime         int64  `json:"playtime"`
	GMCategory         int32  `json:"category"`
	GMCreatedAt        string `json:"created_at"`
	GMUpdatedAt        string `json:"updated_at"`
	GMLeaderboardType  string `json:"leaderboard_type"`
	GMLeaderboardOrder int32  `json:"leaderboard_order"`
}

type MapLeaderboardEntry struct {
	MLEPosition   int32  `json:"position"`
	MLEPlayerName string `json:"player_name"`
	MLEValue      int32  `json:"value"`
	MLEDate       string `json:"date"`
}

type MapLeaderboardResponse struct {
	MLRPage             int32                 `json:"page"`
	MLRPerPage          int32                 `json:"per_page"`
	MLRMapName          string                `json:"map_name"`
	MLRLeaderboardType  string                `json:"leaderboard_type"`
	MLRLeaderboardOrder int32                 `json:"leaderboard_order"`
	MLREntries          []MapLeaderboardEntry `json:"entries"`
}

type Mod struct {
	ModId          int32  `json:"mod_id"`
	ModName        string `json:"name"`
	ModDescription string `json:"description"`
	ModCreatorName string `json:"creator_name"`
	ModVotes       int32  `json:"votes"`
	ModFeatured    bool   `json:"featured"`
	ModVersion     int32  `json:"version"`
	ModCreatedAt   string `json:"created_at"`
	ModUpdatedAt   string `json:"updated_at"`
}

type ModsResponse struct {
	ModsPage    int32 `json:"page"`
	ModsPerPage int32 `json:"per_page"`
	ModsMods    []Mod `json:"mods"`
}

type MarketListing struct {
	MLPrice      int32  `json:"price"`
	MLSellerName string `json:"seller_name"`
	MLListedAt   string `json:"listed_at"`
}

type MarketOwner struct {
	MOPlayerName string `json:"player_name"`
	MOCount      int32  `json:"count"`
}

type PriceHistory struct {
	PHDate         string  `json:"date"`
	PHAveragePrice float64 `json:"average_price"`
	PHSales        int32   `json:"sales"`
}

type MarketResponse struct {
	MRSkinIndex        int32           `json:"skin_index"`
	MRTotalListings    int32           `json:"total_listings"`
	MRLowestPrice      int32           `json:"lowest_price"`
	MRAveragePrice     float64         `json:"average_price"`
	MRTotalCirculating int32           `json:"total_circulating"`
	MRListings         []MarketListing `json:"listings"`
	MROwners           []MarketOwner   `json:"owners"`
	MRPriceHistory     []PriceHistory  `json:"price_history"`
}

type RateLimitResponse struct {
	Error      string `json:"error"`
	RetryAfter uint64 `json:"retry_after"`
}

type GenericErrorResponse struct {
	Error string `json:"error"`
}

type RateLimitInfo struct {
	Limit     uint32 `json:"limit"`
	Remaining uint32 `json:"remaining"`
	Reset     uint64 `json:"reset"`
}
