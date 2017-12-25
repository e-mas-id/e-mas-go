package response

type Customer struct {
	Balance 		float64		`json:"balance"`
	Address			string		`json:"address"`
	City			string		`json:"city"`
	Province		string		`json:"province"`
	Zipcode 		string		`json:"zipcode"`
}