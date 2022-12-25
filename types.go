package main

type DeletionReason struct {
	Name              string `json:"name"`
	ID                int    `json:"id"`
	Regex             string `json:"regex"`
	WithWarn          bool   `json:"withWarn"`
	TakePoints        bool   `json:"takePoints"`
	ReturnPoints      bool   `json:"returnPoints"`
	UserMustBeDeleted bool   `json:"userMustBeDeleted"`
}

type User struct {
	ID   int    `json:"id"`
	Nick string `json:"nick"`
}

type Action struct {
	Content        string         `json:"content"`
	TaskID         int            `json:"taskId"`
	Date           string         `json:"date"`
	User           User           `json:"user"`
	ContentType    string         `json:"contentType"`
	Type           string         `json:"type"`
	Reason         string         `json:"reason"`
	ModeratorID    int            `json:"moderatorId"`
	DeletionReason DeletionReason `json:"deletionReason"`
	Attachments    []string       `json:"attachments"`
}