package gjmodels

type Project struct {
	Expand         string  `json:"expand"`
	ID             string  `json:"id"`
	Key            string  `json:"key"`
	Self           string  `json:"self"`
	Name           string  `json:"name"`
	ProjectTypeKey string  `json:"project_type_key"`
	Simplified     bool    `json:"simplified"`
	AvatarUrls     Avatars `json:"avatarUrls"`
}
