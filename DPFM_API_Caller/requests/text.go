package requests

type Text struct {
	Language     			string  `json:"Language"`
	CorrespondenceLanguage	string  `json:"CorrespondenceLanguage"`
	LanguageName			string  `json:"LanguageName"`
	CreationDate			string	`json:"CreationDate"`
	LastChangeDate			string	`json:"LastChangeDate"`
	IsMarkedForDeletion		*bool	`json:"IsMarkedForDeletion"`
}
