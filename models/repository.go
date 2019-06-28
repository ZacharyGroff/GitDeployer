package models

import (
	"time"
)

type Repository struct {
    ArchiveUrl       string    `json:"archive_url"`
    AssigneesUrl     string    `json:"assignees_url"`
    BlobsUrl         string    `json:"blobs_url"`
    BranchesUrl      string    `json:"branches_url"`
    CloneUrl         string    `json:"clone_url"`
    CollaboratorsUrl string    `json:"collaborators_url"`
    CommentsUrl      string    `json:"comments_url"`
    CommitsUrl       string    `json:"commits_url"`
    CompareUrl       string    `json:"compare_url"`
    ContentsUrl      string    `json:"contents_url"`
    ContributorsUrl  string    `json:"contributors_url"`
    CreatedAt        time.Time `json:"created_at"`
    DefaultBranch    string    `json:"default_branch"`
    Description      string    `json:"description"`
    DownloadsUrl     string    `json:"downloads_url"`
    EventsUrl        string    `json:"events_url"`
    Fork             bool      `json:"fork"`
    Forks            int       `json:"forks"`
    ForksCount       int       `json:"forks_count"`
    ForksUrl         string    `json:"forks_url"`
    FullName         string    `json:"full_name"`
    GitCommitsUrl    string    `json:"git_commits_url"`
    GitRefsUrl       string    `json:"git_refs_url"`
    GitTagsUrl       string    `json:"git_tags_url"`
    GitUrl           string    `json:"git_url"`
    HtmlUrl          string    `json:"html_url"`
    HasDownloads     bool      `json:"has_downloads"`
    HasIssues        bool      `json:"has_issues"`
    HasPages         bool      `json:"has_pages"`
    HasWiki          bool      `json:"has_wiki"`
    Homepage         string    `json:"homepage"`
    HooksUrl         string    `json:"hooks_url"`
    Id               int       `json:"id"`
    IssueCommentUrl  string    `json:"issue_comment_url"`
    IssueEventsUrl   string    `json:"issue_events_url"`
    IssuesUrl        string    `json:"issues_url"`
    KeysUrl          string    `json:"keys_url"`
    LabelsUrl        string    `json:"labels_url"`
    Language         string    `json:"language"`
    LanguagesUrl     string    `json:"languages_url"`
    MasterBranch     string    `json:"master_branch"`
    MergesUrl        string    `json:"merges_url"`
    MilestonesUrl    string    `json:"milestones_url"`
    MirrorUrl        string    `json:"mirror_url"`
    Name             string    `json:"name"`
    NotificationsUrl string    `json:"notifications_url"`
    OpenIssues       int       `json:"open_issues"`
    OpenIssuesCount  int       `json:"open_issues_count"`
    Owner            Owner     `json:"owner"`
    Private          bool      `json:"private"`
    PullsUrl         string    `json:"pulls_url"`
    PushedAt         time.Time `json:"pushed_at"`
    ReleasesUrl      string    `json:"releases_url"`
    SshUrl           string    `json:"ssh_url"`
    Size             int       `json:"size"`
    Stargazers       int       `json:"stargazers"`
    StargazersCount  int       `json:"stargazers_count"`
    StargazersUrl    string    `json:"stargazers_url"`
    StatusesUrl      string    `json:"statuses_url"`
    SubscribersUrl   string    `json:"subscribers_url"`
    SubscriptionUrl  string    `json:"subscription_url"`
    SvnUrl           string    `json:"svn_url"`
    TagsUrl          string    `json:"tags_url"`
    TeamsUrl         string    `json:"teams_url"`
    TreesUrl         string    `json:"trees_url"`
    Url              string    `json:"url"`
    UpdatedAt        time.Time `json:"updated_at"`
    Watchers         int       `json:"watchers"`
    WatchersCount    int       `json:"watchers_count"`
}
