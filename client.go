package krunkgo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type Client struct {
	baseURL   string
	http      *http.Client
	apiKey    string
	debug     atomic.Bool
	rateMu    sync.RWMutex
	rateLimit *RateLimitInfo
}

func NewClient(apiKey string) *Client {
	return &Client{
		baseURL: "https://gapi.svc.krunker.io/api",
		http:    &http.Client{Timeout: 15 * time.Second},
		apiKey:  apiKey,
	}
}

func (c *Client) SetDebug(debug bool) {
	c.debug.Store(debug)
}

func (c *Client) LastRateLimit() *RateLimitInfo {
	c.rateMu.RLock()
	defer c.rateMu.RUnlock()
	if c.rateLimit == nil {
		return nil
	}
	cp := *c.rateLimit
	return &cp
}

type queryParam struct {
	key   string
	value string
}

func request[T any](c *Client, path string, params []queryParam) (*T, error) {
	url := c.baseURL + path

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("X-Developer-API-Key", c.apiKey)

	// 100% sure i dont need this req header, but for some fucking reason it doesn't work without it.....
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	if len(params) > 0 {
		q := req.URL.Query()
		for _, p := range params {
			q.Add(p.key, p.value)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	limitStr := resp.Header.Get("X-RateLimit-Limit")
	remStr := resp.Header.Get("X-RateLimit-Remaining")
	resetStr := resp.Header.Get("X-RateLimit-Reset")

	if limitStr != "" && remStr != "" && resetStr != "" {
		limit, _ := strconv.ParseUint(limitStr, 10, 32)
		remaining, _ := strconv.ParseUint(remStr, 10, 32)
		reset, _ := strconv.ParseUint(resetStr, 10, 64)

		c.rateMu.Lock()
		c.rateLimit = &RateLimitInfo{
			Limit:     uint32(limit),
			Remaining: uint32(remaining),
			Reset:     reset,
		}
		c.rateMu.Unlock()
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading body bytes: %w", err)
	}

	if c.debug.Load() {
		fmt.Printf("DEBUG: Raw response body:\n%s\n", string(bodyBytes))
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var target T
		if err := json.Unmarshal(bodyBytes, &target); err != nil {
			var fieldName string
			if typeErr, ok := err.(*json.UnmarshalTypeError); ok {
				fieldName = typeErr.Field
			}
			return nil, &DecodeError{
				Message: err.Error(),
				Body:    string(bodyBytes),
				Field:   fieldName,
			}
		}
		return &target, nil
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		var rlRate RateLimitResponse
		if err := json.Unmarshal(bodyBytes, &rlRate); err == nil {
			return nil, &RateLimitError{RetryAfter: rlRate.RetryAfter}
		}
	}

	var genErr GenericErrorResponse
	var errMsg string
	if err := json.Unmarshal(bodyBytes, &genErr); err == nil && genErr.Error != "" {
		errMsg = genErr.Error
	} else {
		errMsg = string(bodyBytes)
	}

	return nil, &APIError{
		StatusCode: resp.StatusCode,
		Message:    errMsg,
	}
}

func (c *Client) GetPlayer(name string) (*Player, error) {
	return request[Player](c, "/player/"+name, nil)
}

func (c *Client) GetPlayerInventory(name string) (*[]InventoryItem, error) {
	return request[[]InventoryItem](c, "/player/"+name+"/inventory", nil)
}

func (c *Client) GetPlayerMatches(name string, page int32, season int32) (*PlayerMatchesResponse, error) {
	var params []queryParam
	if page > 0 {
		params = append(params, queryParam{"page", strconv.Itoa(int(page))})
	}
	if season > 0 {
		params = append(params, queryParam{"season", strconv.Itoa(int(season))})
	}
	return request[PlayerMatchesResponse](c, "/player/"+name+"/matches", params)
}

func (c *Client) GetPlayerPosts(name string, page int32) (*PostsResponse, error) {
	var params []queryParam
	if page > 0 {
		params = append(params, queryParam{"page", strconv.Itoa(int(page))})
	}
	return request[PostsResponse](c, "/player/"+name+"/posts", params)
}

func (c *Client) GetMatch(matchID int64) (*Match, error) {
	path := fmt.Sprintf("/match/%d", matchID)
	return request[Match](c, path, nil)
}

func (c *Client) GetClan(name string) (*Clan, error) {
	return request[Clan](c, "/clan/"+name, nil)
}

func (c *Client) GetClanMembers(name string, page int32) (*ClanMembersResponse, error) {
	var params []queryParam
	if page > 0 {
		params = append(params, queryParam{"page", strconv.Itoa(int(page))})
	}
	return request[ClanMembersResponse](c, "/clan/"+name+"/members", params)
}

func (c *Client) GetLeaderboard(region int32, page int32) (*LeaderboardResponse, error) {
	var params []queryParam
	if page > 0 {
		params = append(params, queryParam{"page", strconv.Itoa(int(page))})
	}
	path := fmt.Sprintf("/leaderboard/%d", region)
	return request[LeaderboardResponse](c, path, params)
}

func (c *Client) GetMap(name string) (*GameMap, error) {
	return request[GameMap](c, "/map/"+name, nil)
}

func (c *Client) GetMapLeaderboard(name string, page int32) (*MapLeaderboardResponse, error) {
	var params []queryParam
	if page > 0 {
		params = append(params, queryParam{"page", strconv.Itoa(int(page))})
	}
	return request[MapLeaderboardResponse](c, "/map/"+name+"/leaderboard", params)
}

func (c *Client) GetMods(page int32) (*ModsResponse, error) {
	var params []queryParam
	if page > 0 {
		params = append(params, queryParam{"page", strconv.Itoa(int(page))})
	}
	return request[ModsResponse](c, "/mods", params)
}

func (c *Client) GetMod(name string) (*Mod, error) {
	return request[Mod](c, "/mods/"+name, nil)
}

func (c *Client) GetMarketSkin(skinIndex int32, page int32) (*MarketResponse, error) {
	var params []queryParam
	if page > 0 {
		params = append(params, queryParam{"page", strconv.Itoa(int(page))})
	}
	path := fmt.Sprintf("/market/skin/%d", skinIndex)
	return request[MarketResponse](c, path, params)
}
