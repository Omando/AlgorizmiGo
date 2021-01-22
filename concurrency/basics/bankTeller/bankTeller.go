package bankTeller

/* Implements a back teller as a monitor goroutine
Two channels are used to receive requests to query and update a specific account
A third channel is used to send results back to clients */

// AccountQuery type is used with chAccountQuery channel to send a request to query an account
type AccountQuery struct{
	AccountNo string
}
var chAccountQuery = make(chan AccountQuery)

// AccountReply type is used with accountReply channel to send a reply to the query
type AccountReply struct {
	AccountNo string
	Balance float32
}
var ChAccountReply = make(chan AccountReply)

// AccountUpdate type is used with chAccountUpdate channel to send a request to update an account
type AccountUpdate struct {
	AccountNo string
	Deposit float32
}
var chAccountUpdate = make(chan AccountUpdate)

func Query(query AccountQuery) {
	chAccountQuery <- query
}

func Deposit(depositRequest AccountUpdate) {
	chAccountUpdate <- depositRequest
}

func RunTeller() {
	// Create a map of accounts and balances (get from db for example)
	var accounts = make(map[string]float32)
	accounts["AB12"] = 100
	accounts["CD34"] = 200
	accounts["EF"] = 300

	go func() {
		// todo
	}()
}
