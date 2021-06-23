package Models

// Model showing details of pull request
type PullRequest struct {
	ArtifactID   string `json:"artifactId"`
	CodeReviewID int64  `json:"codeReviewId"`
	CreatedBy    struct {
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
	} `json:"createdBy"`
	CreationDate    string `json:"creationDate"`
	Description     string `json:"description"`
	IsDraft         bool   `json:"isDraft"`
	LastMergeCommit struct {
		Author struct {
			Date  string `json:"date"`
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"author"`
		Comment   string `json:"comment"`
		CommitID  string `json:"commitId"`
		Committer struct {
			Date  string `json:"date"`
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"committer"`
		URL string `json:"url"`
	} `json:"lastMergeCommit"`
	LastMergeSourceCommit struct {
		CommitID string `json:"commitId"`
		URL      string `json:"url"`
	} `json:"lastMergeSourceCommit"`
	LastMergeTargetCommit struct {
		CommitID string `json:"commitId"`
		URL      string `json:"url"`
	} `json:"lastMergeTargetCommit"`
	MergeID       string `json:"mergeId"`
	MergeStatus   string `json:"mergeStatus"`
	PullRequestID int64  `json:"pullRequestId"`
	Repository    struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Project struct {
			Description    string `json:"description"`
			ID             string `json:"id"`
			LastUpdateTime string `json:"lastUpdateTime"`
			Name           string `json:"name"`
			Revision       int64  `json:"revision"`
			State          string `json:"state"`
			URL            string `json:"url"`
			Visibility     string `json:"visibility"`
		} `json:"project"`
		RemoteURL string `json:"remoteUrl"`
		Size      int64  `json:"size"`
		SSHURL    string `json:"sshUrl"`
		URL       string `json:"url"`
		WebURL    string `json:"webUrl"`
	} `json:"repository"`
	Reviewers []struct {
		Links struct {
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
		} `json:"_links"`
		DisplayName string `json:"displayName"`
		HasDeclined bool   `json:"hasDeclined"`
		ID          string `json:"id"`
		ImageURL    string `json:"imageUrl"`
		IsContainer bool   `json:"isContainer"`
		IsFlagged   bool   `json:"isFlagged"`
		IsRequired  bool   `json:"isRequired"`
		ReviewerURL string `json:"reviewerUrl"`
		UniqueName  string `json:"uniqueName"`
		URL         string `json:"url"`
		Vote        int64  `json:"vote"`
	} `json:"reviewers"`
	SourceRefName      string `json:"sourceRefName"`
	Status             string `json:"status"`
	SupportsIterations bool   `json:"supportsIterations"`
	TargetRefName      string `json:"targetRefName"`
	Title              string `json:"title"`
	URL                string `json:"url"`
}

