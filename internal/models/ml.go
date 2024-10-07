package models

type Datafeed struct {
	FeedID                 string                         `json:"-"`
	Aggregations           interface{}                    `json:"aggregations,omitempty"`
	ChunkingConfig         DatafeedChunkingConfig         `json:"chunking_config,omitempty"`
	DelayedDataCheckConfig DatafeedDelayedDataCheckConfig `json:"delayed_data_check_config,omitempty"`
	// TODO: Time units? - https://www.elastic.co/guide/en/elasticsearch/reference/current/api-conventions.html#time-units
	Frequency string   `json:"frequency,omitempty"`
	Indices   []string `json:"indices"`
	// TODO: indices_options
	JobID            string `json:"job_id"`
	MaxEmptySearches int    `json:"max_empty_searches,omitempty"`
	// TODO: query
	// TODO: Time units? - https://www.elastic.co/guide/en/elasticsearch/reference/current/api-conventions.html#time-units
	QueryDelay string `json:"query_delay,omitempty"`
	// TODO: runtime_mappings
	// TODO: script_fields
	ScrollSize int `json:"scroll_size,omitempty"`
}

type DatafeedDelayedDataCheckConfig struct {
	// TODO: Time units? - https://www.elastic.co/guide/en/elasticsearch/reference/current/api-conventions.html#time-units
	CheckWindow string `json:"check_window"`
	Enabled     bool   `json:"enabled,omitempty"`
}

type DatafeedChunkingConfig struct {
	Mode string `json:"mode"`
	// TODO: Time units? - https://www.elastic.co/guide/en/elasticsearch/reference/current/api-conventions.html#time-units
	TimeSpan string `json:"time_span"`
}

type PutDatafeedParams struct {
	AllowNoIndices    bool   `json:"allow_no_indices,omitempty"`
	ExpandWildcards   string `json:"expand_wildcards,omitempty"`
	IgnoreThrottled   bool   `json:"ignore_throttled,omitempty"`
	IgnoreUnavailable bool   `json:"ignore_unavailable,omitempty"`
}
