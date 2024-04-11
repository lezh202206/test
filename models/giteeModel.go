package models

import "time"

type GiteeResponse struct {
	Id                int    `json:"id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	DiffUrl           string `json:"diff_url"`
	PatchUrl          string `json:"patch_url"`
	IssueUrl          string `json:"issue_url"`
	CommitsUrl        string `json:"commits_url"`
	ReviewCommentsUrl string `json:"review_comments_url"`
	ReviewCommentUrl  string `json:"review_comment_url"`
	CommentsUrl       string `json:"comments_url"`
	Number            int    `json:"number"`
	CloseRelatedIssue int    `json:"close_related_issue"`
	PruneBranch       bool   `json:"prune_branch"`
	State             string `json:"state"`
	AssigneesNumber   int    `json:"assignees_number"`
	TestersNumber     int    `json:"testers_number"`
	Assignees         []struct {
		Id                int    `json:"id"`
		Login             string `json:"login"`
		Name              string `json:"name"`
		AvatarUrl         string `json:"avatar_url"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		Remark            string `json:"remark"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		Assignee          bool   `json:"assignee"`
		CodeOwner         bool   `json:"code_owner"`
		Accept            bool   `json:"accept"`
	} `json:"assignees"`
	Testers []struct {
		Id                int    `json:"id"`
		Login             string `json:"login"`
		Name              string `json:"name"`
		AvatarUrl         string `json:"avatar_url"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		Remark            string `json:"remark"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		Assignee          bool   `json:"assignee"`
		CodeOwner         bool   `json:"code_owner"`
		Accept            bool   `json:"accept"`
	} `json:"testers"`
	Milestone     interface{}   `json:"milestone"`
	Labels        []interface{} `json:"labels"`
	Locked        bool          `json:"locked"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	ClosedAt      interface{}   `json:"closed_at"`
	Draft         bool          `json:"draft"`
	MergedAt      interface{}   `json:"merged_at"`
	Mergeable     bool          `json:"mergeable"`
	CanMergeCheck bool          `json:"can_merge_check"`
	Links         struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Html struct {
			Href string `json:"href"`
		} `json:"html"`
		Issue struct {
			Href string `json:"href"`
		} `json:"issue"`
		Comments struct {
			Href string `json:"href"`
		} `json:"comments"`
		ReviewComments struct {
			Href string `json:"href"`
		} `json:"review_comments"`
		ReviewComment struct {
			Href string `json:"href"`
		} `json:"review_comment"`
		Commits struct {
			Href string `json:"href"`
		} `json:"commits"`
	} `json:"_links"`
	User struct {
		Id                int    `json:"id"`
		Login             string `json:"login"`
		Name              string `json:"name"`
		AvatarUrl         string `json:"avatar_url"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		Remark            string `json:"remark"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
	} `json:"user"`
	RefPullRequests []interface{} `json:"ref_pull_requests"`
	Title           string        `json:"title"`
	Body            interface{}   `json:"body"`
	Head            struct {
		Label string `json:"label"`
		Ref   string `json:"ref"`
		Sha   string `json:"sha"`
		User  struct {
			Id                int    `json:"id"`
			Login             string `json:"login"`
			Name              string `json:"name"`
			AvatarUrl         string `json:"avatar_url"`
			Url               string `json:"url"`
			HtmlUrl           string `json:"html_url"`
			Remark            string `json:"remark"`
			FollowersUrl      string `json:"followers_url"`
			FollowingUrl      string `json:"following_url"`
			GistsUrl          string `json:"gists_url"`
			StarredUrl        string `json:"starred_url"`
			SubscriptionsUrl  string `json:"subscriptions_url"`
			OrganizationsUrl  string `json:"organizations_url"`
			ReposUrl          string `json:"repos_url"`
			EventsUrl         string `json:"events_url"`
			ReceivedEventsUrl string `json:"received_events_url"`
			Type              string `json:"type"`
		} `json:"user"`
		Repo struct {
			Id        int    `json:"id"`
			FullName  string `json:"full_name"`
			HumanName string `json:"human_name"`
			Url       string `json:"url"`
			Namespace struct {
				Id      int    `json:"id"`
				Type    string `json:"type"`
				Name    string `json:"name"`
				Path    string `json:"path"`
				HtmlUrl string `json:"html_url"`
			} `json:"namespace"`
			Path  string `json:"path"`
			Name  string `json:"name"`
			Owner struct {
				Id                int    `json:"id"`
				Login             string `json:"login"`
				Name              string `json:"name"`
				AvatarUrl         string `json:"avatar_url"`
				Url               string `json:"url"`
				HtmlUrl           string `json:"html_url"`
				Remark            string `json:"remark"`
				FollowersUrl      string `json:"followers_url"`
				FollowingUrl      string `json:"following_url"`
				GistsUrl          string `json:"gists_url"`
				StarredUrl        string `json:"starred_url"`
				SubscriptionsUrl  string `json:"subscriptions_url"`
				OrganizationsUrl  string `json:"organizations_url"`
				ReposUrl          string `json:"repos_url"`
				EventsUrl         string `json:"events_url"`
				ReceivedEventsUrl string `json:"received_events_url"`
				Type              string `json:"type"`
			} `json:"owner"`
			Assigner struct {
				Id                int    `json:"id"`
				Login             string `json:"login"`
				Name              string `json:"name"`
				AvatarUrl         string `json:"avatar_url"`
				Url               string `json:"url"`
				HtmlUrl           string `json:"html_url"`
				Remark            string `json:"remark"`
				FollowersUrl      string `json:"followers_url"`
				FollowingUrl      string `json:"following_url"`
				GistsUrl          string `json:"gists_url"`
				StarredUrl        string `json:"starred_url"`
				SubscriptionsUrl  string `json:"subscriptions_url"`
				OrganizationsUrl  string `json:"organizations_url"`
				ReposUrl          string `json:"repos_url"`
				EventsUrl         string `json:"events_url"`
				ReceivedEventsUrl string `json:"received_events_url"`
				Type              string `json:"type"`
			} `json:"assigner"`
			Description string `json:"description"`
			Private     bool   `json:"private"`
			Public      bool   `json:"public"`
			Internal    bool   `json:"internal"`
			Fork        bool   `json:"fork"`
			HtmlUrl     string `json:"html_url"`
			SshUrl      string `json:"ssh_url"`
		} `json:"repo"`
	} `json:"head"`
	Base struct {
		Label string `json:"label"`
		Ref   string `json:"ref"`
		Sha   string `json:"sha"`
		User  struct {
			Id                int    `json:"id"`
			Login             string `json:"login"`
			Name              string `json:"name"`
			AvatarUrl         string `json:"avatar_url"`
			Url               string `json:"url"`
			HtmlUrl           string `json:"html_url"`
			Remark            string `json:"remark"`
			FollowersUrl      string `json:"followers_url"`
			FollowingUrl      string `json:"following_url"`
			GistsUrl          string `json:"gists_url"`
			StarredUrl        string `json:"starred_url"`
			SubscriptionsUrl  string `json:"subscriptions_url"`
			OrganizationsUrl  string `json:"organizations_url"`
			ReposUrl          string `json:"repos_url"`
			EventsUrl         string `json:"events_url"`
			ReceivedEventsUrl string `json:"received_events_url"`
			Type              string `json:"type"`
		} `json:"user"`
		Repo struct {
			Id        int    `json:"id"`
			FullName  string `json:"full_name"`
			HumanName string `json:"human_name"`
			Url       string `json:"url"`
			Namespace struct {
				Id      int    `json:"id"`
				Type    string `json:"type"`
				Name    string `json:"name"`
				Path    string `json:"path"`
				HtmlUrl string `json:"html_url"`
			} `json:"namespace"`
			Path  string `json:"path"`
			Name  string `json:"name"`
			Owner struct {
				Id                int    `json:"id"`
				Login             string `json:"login"`
				Name              string `json:"name"`
				AvatarUrl         string `json:"avatar_url"`
				Url               string `json:"url"`
				HtmlUrl           string `json:"html_url"`
				Remark            string `json:"remark"`
				FollowersUrl      string `json:"followers_url"`
				FollowingUrl      string `json:"following_url"`
				GistsUrl          string `json:"gists_url"`
				StarredUrl        string `json:"starred_url"`
				SubscriptionsUrl  string `json:"subscriptions_url"`
				OrganizationsUrl  string `json:"organizations_url"`
				ReposUrl          string `json:"repos_url"`
				EventsUrl         string `json:"events_url"`
				ReceivedEventsUrl string `json:"received_events_url"`
				Type              string `json:"type"`
			} `json:"owner"`
			Assigner struct {
				Id                int    `json:"id"`
				Login             string `json:"login"`
				Name              string `json:"name"`
				AvatarUrl         string `json:"avatar_url"`
				Url               string `json:"url"`
				HtmlUrl           string `json:"html_url"`
				Remark            string `json:"remark"`
				FollowersUrl      string `json:"followers_url"`
				FollowingUrl      string `json:"following_url"`
				GistsUrl          string `json:"gists_url"`
				StarredUrl        string `json:"starred_url"`
				SubscriptionsUrl  string `json:"subscriptions_url"`
				OrganizationsUrl  string `json:"organizations_url"`
				ReposUrl          string `json:"repos_url"`
				EventsUrl         string `json:"events_url"`
				ReceivedEventsUrl string `json:"received_events_url"`
				Type              string `json:"type"`
			} `json:"assigner"`
			Description string `json:"description"`
			Private     bool   `json:"private"`
			Public      bool   `json:"public"`
			Internal    bool   `json:"internal"`
			Fork        bool   `json:"fork"`
			HtmlUrl     string `json:"html_url"`
			SshUrl      string `json:"ssh_url"`
		} `json:"repo"`
	} `json:"base"`
}
