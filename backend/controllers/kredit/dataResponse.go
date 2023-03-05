package kredit

type ResponseChecklistPencairan struct {
	Records     []RecordChecklistPencairan `json:"records"`
	CountRecord int                        `json:"count_record"`
	CountPage   int                        `json:"count_page"`
}

type ResponseDrawdownReport struct {
	ResponseChecklistPencairan
	Companies        []RecordCompany `json:"companies"`
	Branches         []RecordBranch  `json:"branches"`
	ApprovalStatuses []string        `json:"approval_statuses"`
}
