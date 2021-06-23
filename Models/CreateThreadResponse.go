package Models

// The pattern of the response returned when a thread is created on Azure Devops. The id here will help us add a comment.
type CreateTheradResponse struct {
	Links struct {
		Repository struct {
			Href string `json:"href"`
		} `json:"repository"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Comments []struct {
		Links struct {
			PullRequests struct {
				Href string `json:"href"`
			} `json:"pullRequests"`
			Repository struct {
				Href string `json:"href"`
			} `json:"repository"`
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Threads struct {
				Href string `json:"href"`
			} `json:"threads"`
		} `json:"_links"`
		Author struct {
			Links struct {
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"_links"`
			Descriptor  string `json:"descriptor"`
			DisplayName string `json:"displayName"`
			ID          string `json:"id"`
			ImageURL    string `json:"imageUrl"`
			UniqueName  string `json:"uniqueName"`
			URL         string `json:"url"`
		} `json:"author"`
		CommentType            string `json:"commentType"`
		Content                string `json:"content"`
		ID                     int64  `json:"id"`
		LastContentUpdatedDate string `json:"lastContentUpdatedDate"`
		LastUpdatedDate        string `json:"lastUpdatedDate"`
		ParentCommentID        int64  `json:"parentCommentId"`
		PublishedDate          string `json:"publishedDate"`
	} `json:"comments"`
	ID                       int 	     `json:"id"`
	Identities               interface{} `json:"identities"`
	IsDeleted                bool        `json:"isDeleted"`
	LastUpdatedDate          string      `json:"lastUpdatedDate"`
	Properties               interface{} `json:"properties"`
	PublishedDate            string      `json:"publishedDate"`
	PullRequestThreadContext interface{} `json:"pullRequestThreadContext"`
	ThreadContext            interface{} `json:"threadContext"`
}

