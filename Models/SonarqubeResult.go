package Models

// The response pattern that returns after querying the Sonarqube api.
type SonarQubeResult struct {
	Components []struct {
		Enabled     bool   `json:"enabled"`
		Key         string `json:"key"`
		LongName    string `json:"longName"`
		Name        string `json:"name"`
		Path        string `json:"path"`
		PullRequest string `json:"pullRequest"`
		Qualifier   string `json:"qualifier"`
		UUID        string `json:"uuid"`
	} `json:"components"`
	DebtTotal   int64         `json:"debtTotal"`
	EffortTotal int64         `json:"effortTotal"`
	Facets      []interface{} `json:"facets"`
	Issues      []struct {
		Assignee     string        `json:"assignee"`
		Author       string        `json:"author"`
		Component    string        `json:"component"`
		CreationDate string        `json:"creationDate"`
		Debt         string        `json:"debt"`
		Effort       string        `json:"effort"`
		Flows        []interface{} `json:"flows"`
		Hash         string        `json:"hash"`
		Key          string        `json:"key"`
		Line         int           `json:"line"`
		Message      string        `json:"message"`
		Project      string        `json:"project"`
		PullRequest  string        `json:"pullRequest"`
		Rule         string        `json:"rule"`
		Scope        string        `json:"scope"`
		Severity     string        `json:"severity"`
		Status       string        `json:"status"`
		Tags         []string      `json:"tags"`
		TextRange    struct {
			EndLine     int64 `json:"endLine"`
			EndOffset   int64 `json:"endOffset"`
			StartLine   int64 `json:"startLine"`
			StartOffset int64 `json:"startOffset"`
		} `json:"textRange"`
		Type       string `json:"type"`
		UpdateDate string `json:"updateDate"`
	} `json:"issues"`
	P      int64 `json:"p"`
	Paging struct {
		PageIndex int64 `json:"pageIndex"`
		PageSize  int64 `json:"pageSize"`
		Total     int64 `json:"total"`
	} `json:"paging"`
	Ps    int64 `json:"ps"`
	Total int64 `json:"total"`
}
