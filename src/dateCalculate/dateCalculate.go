package dateCalculate

type  duration struct {
    Form  	 	string
	To 	   	 	string
	Days   	 	string
	Years    	string
	Seconds	 	string
	Minutes  	string
	Hours 	    string
	Weeks 	    string
	RatioOfYear string
}

func MakeJSON(startDate, endDate string) duration{
	return duration{Form : "Thursday, 4 January 2018",
		To : "Monday, 4 June 2018",
		Days : "152 days",
		Years : "5 months, 1 days",
		Seconds : "13,132,800 seconds",
		Minutes : "218,880 minutes",
		Hours : "3,648 hours",
		Weeks : "1 weeks 7 days",
		RatioOfYear : "41.64% of 2018",}
}