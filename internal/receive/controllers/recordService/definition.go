package recordService

const (
	Database = "database"
)

type RecordQuery struct {
	ID string `json:"id" xml:"id" form:"id" binding:"required"`
}
