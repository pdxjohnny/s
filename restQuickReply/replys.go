package restQuickReply

var (
	// BlankResponse is and epty JSON
	BlankResponse = []byte("{}")
	// OKResponse sends ok = true
	OKResponse = []byte("{\"OK\":true}")
	// NotOKResponse sends ok = false
	NotOKResponse = []byte("{\"OK\":false}")
)
