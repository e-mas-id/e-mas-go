package response

type Error struct {
	ErrorCode 		int			`json:"error_code"`
	ErrorMessage	string		`json:"error_message"`
	ErrorData		interface{}	`json:"error_data"`
	Timestamp		string		`json:"timestamp"`
}