package Models

// ConfigDatas Model that holds usernames, passwords, and other required parametric information
type ConfigDatas struct {
	ConfigDatas []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"configDatas"`
}