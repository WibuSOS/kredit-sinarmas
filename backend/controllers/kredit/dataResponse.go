package kredit

type ResponseChecklistPencairan struct {
	Records     []RecordChecklistPencairan `json:"records"`
	CountRecord int                        `json:"count_record"`
	CountPage   int                        `json:"count_page"`
}
