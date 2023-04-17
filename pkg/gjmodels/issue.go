package gjmodels

type IssueStatusCategory struct {
	ID        uint64 `json:"id"`
	Key       string `json:"key"`
	Self      string `json:"self"`
	Name      string `json:"name"`
	ColorName string `json:"colorName"`
}

type IssueStatus struct {
	ID             string              `json:"id"`
	Self           string              `json:"self"`
	Name           string              `json:"name"`
	Description    string              `json:"description"`
	IconURL        string              `json:"iconUrl"`
	StatusCategory IssueStatusCategory `json:"statusCategory"`
}

type IssueType struct {
	ID             string `json:"id"`
	Self           string `json:"self"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	IconURL        string `json:"iconUrl"`
	Subtask        bool   `json:"subtask"`
	AvatarID       int    `json:"avatarId"`
	EntityIdd      string `json:"entityId"`
	HierarchyLevel int    `json:"hierarchyLevel"`
}

type IssuePriority struct {
	ID      string `json:"id"`
	Self    string `json:"self"`
	Name    string `json:"name"`
	IconURL string `json:"iconUrl"`
}

type SubTasks struct {
	ID        string        `json:"id"`
	Key       string        `json:"key"`
	Self      string        `json:"self"`
	Summary   string        `json:"summary"`
	IssueType IssueType     `json:"issuetype"`
	Priority  IssuePriority `json:"priority"`
	Status    IssueStatus   `json:"status"`
}

type IssueFields struct {
	Created        string        `json:"created"`
	Updated        string        `json:"updated"`
	ResolutionDate *string       `json:"resolutiondate"`
	SubTasks       []SubTasks    `json:"subtasks"`
	Priority       IssuePriority `json:"priority"`
	IssueType      IssueType     `json:"issuetype"`
	Summary        string        `json:"summary"`
	Status         IssueStatus   `json:"status"`
	Project        Project       `json:"project"`
	Creator        User          `json:"creator"`
	Reporter       User          `json:"reporter"`
	Assignee       *User         `json:"assignee"`
	Developer      *User         `json:"developer"`
}

type Issue struct {
	Expand string      `json:"expand"`
	ID     string      `json:"id"`
	Key    string      `json:"key"`
	Self   string      `json:"self"`
	Fields IssueFields `json:"fields"`
}
