package golist

// Botblock api Client structure
type Client struct {
	// All the tokens supplied!
	Tokens map[string]string
}

// Botblock api's bot struct
type Bot struct {
	// ID of the discord bot
	ID string `json:"id"`
	// Username of the discord bot
	Username string `json:"username"`
	// Discriminator of the discord bot
	Discriminator string `json:"discriminator"`
	// Array of owner's discord id
	Owners []string `json:"owners"`
	// Server count of the bot
	ServerCount int64 `json:"server_count"`
	// Invite url of the bot
	Invite string `json:"invite"`
	// Prefix of the discord bot
	Prefix string `json:"prefix"`
	// Website url of the bot
	Website string `json:"website"`
	// Support server invite of the bot
	Support string `json:"support"`
	// Github repo url of the bot
	Github string `json:"github"`
	// The library which was used to make the bot
	Library string `json:"library"`
	// A map with the key as botlist id and the value as the data sent by each botlist!
	ListData map[string][]interface{} `json:"list_data"`
}

// Feature structure of the list
type ListFeature struct {
	// Name of the feature
	Name string `json:"name"`
	// Id of the feature
	ID int64 `json:"id"`
	// Order index of the features
	Display int64 `json:"display"`
	// Types of feature (0 = positive, 1 = neutral, 2 = bad
	Type int64 `json:"type"`
	// Description of the feature
	Description string `json:"description"`
	// Number stating is the list known to have this feature (0 = false, 1 = true)
	Value int64 `json:"value"`
}

// Botblock api's botlist structure
type List struct {
	// ID of the discord bot list
	ID string `json:"id"`
	// Timestamp when the botlist was added
	AddedAt int64 `json:"added"`
	// Name of the list
	Name string `json:"name"`
	// Web url of the list
	URL string `json:"url"`
	// Icon url of the list
	Icon string `json:"icon"`
	// Language which the list uses to communicate
	Language string `json:"language"`
	// Number stating is the list displayed in botblock or not (0 = hidden)
	Display int `json:"display"`
	// Number stating is the list defunct or not (0 = list active)
	Defunct int `json:"defunct"`
	// Boolean stating is the list discord related only or not
	DiscordOnly int `json:"discord_only"`
	// Description of the list
	Description string `json:"description"`
	// Api documentation url of the list
	ApiDocs string `json:"api_docs"`
	// Api url of the list where you can post stats
	ApiPost string `json:"api_post"`
	// Api field of the object returned by the list's api which contains the data
	ApiField string `json:"api_field"`
	// List's api shard id
	ApiShardID string `json:"api_shard_id"`
	// List's api shard count
	ApiShardCount string `json:"api_shard_count"`
	// List's api shards
	ApiShards string `json:"api_shards"`
	// List's web api url where you can get bot details
	ApiGet string `json:"api_get"`
	// List's web api url where you can get all the bots
	ApiAll string `json:"api_all"`
	// List's web url where you can view bot's details
	ViewBot string `json:"view_bot"`
	// List's bot widget api url
	BotWidget string `json:"bot_widget"`
	// Content of the botlist
	Content string `json:"content"`
	// Owners of the list
	Owners string `json:"owners"`
	// Discord invite of the list
	Discord string `json:"discord"`
	// Array of features of the list
	Features []ListFeature `json:"features"`
}

// Response snet by the post method of botblock api
type PostResponse struct {
	// A map with keys of botlist id which have successfully updated stats and the value is an array with the first item in the array is the HTTP status code of the response, the second is the string response body and the third item is the body data sent to the list.
	Success map[string][]interface{} `json:"success"`
	// A map with keys of botlist id which have failed updating stats and the value is an array with the first item in the array is the HTTP status code of the response, the second is the string response body and the third item is the body data sent to the list.
	Failure map[string][]interface{} `json:"failure"`
}

// Stats options structure to be sent in Client.PostStats
type Stats struct {
	// Server count of your bot to post
	ServerCount int64
	// Shard id
	ShardID int64
	// Shard count
	ShardCount int64
	// Array of shards
	Shards []int64
}

// Returns a new Client with default options!
func NewClient() Client {

	return Client{map[string]string{}}

}

// Add a paticular auth token for a paticular botlist after declaring client!
// The id paramater should be the botlist id and token should be the token of the botlist
func (self Client) AddToken(botlist string, token string) {
	self.Tokens[botlist] = token
}

// Post bot stats by sending the bot stats to all the lists registered in botblock!
// The id parameter should be your discord bot id and stats should be the Stats struct containing the stats details!
func (self Client) PostStats(ID string, stats Stats) (PostResponse, error) {
	body := map[string]interface{}{
		"server_count": stats.ServerCount,
		"bot_id":       ID,
		"shard_id":     stats.ShardID,
		"shard_count":  stats.ShardCount,
		"shards":       stats.Shards,
	}

	for list, token := range self.Tokens {
		body[list] = token
	}

	var response PostResponse
	err := Fetch("POST", "/count", &response, body)

	return response, err
}

// Returns the bot information by supplying the id parameter!
func (self Client) GetBot(ID string) (Bot, error) {
	var bot Bot
	err := Fetch("GET", "/bots/"+ID, &bot, map[string]interface{}{})
	return bot, err
}

// Returns the botlist information which is registered on botblock by supplying the id parameter!
func (self Client) GetList(ID string) (List, error) {
	var list List
	err := Fetch("GET", "/lists/"+ID, &list, map[string]interface{}{})
	return list, err
}

// Returns a map with botlist id as string and value as list containg all lists registered in botblock!
func (self Client) GetAllLists() (map[string]List, error) {
	var lists map[string]List
	err := Fetch("GET", "/lists", &lists, map[string]interface{}{})
	return lists, err
}

// Returns a map with the keys as old botlist if and values with the new one from the botblock api!
func (self Client) GetLegacyIDS() (map[string]string, error) {
	var ids map[string]string
	err := Fetch("GET", "/legacy-ids", &ids, map[string]interface{}{})
	return ids, err
}
