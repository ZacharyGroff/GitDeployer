package models

type Sender struct {
    AvatarUrl         string `json:"avatar_url"`
    EventsUrl         string `json:"events_url"`
    FollowersUrl      string `json:"followers_url"`
    FollowingUrl      string `json:"following_url"`
    GistsUrl          string `json:"gists_url"`
    GravatarId        string `json:"gravatar_id"`
    HtmlUrl           string `json:"html_url"`
    Id                int    `json:"id"`
    Login             string `json:"login"`
    OrganizationsUrl  string `json:"organizations_url"`
    ReceivedEventsUrl string `json:"received_events_url"`
    ReposUrl          string `json:"repos_url"`
    SiteAdmin         bool   `json:"site_admin"`
    StarredUrl        string `json:"starred_url"`
    SubscriptionsUrl  string `json:"subscriptions_url"`
    Type              string `json:"type"`
    Url               string `json:"url"`
}
