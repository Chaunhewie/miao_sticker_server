package models

type UserInfo struct {
	Name       string `json:"name"`
	ProfileUrl string `json:"profile_url"`
}

type RepoInfo struct {
	StargazersCount int32 `json:"stargazers_count"`
	WatchersCount   int32 `json:"watchers_count"`
	Forks           int32 `json:"forks_count"`
}
