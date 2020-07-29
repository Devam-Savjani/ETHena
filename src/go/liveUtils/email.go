package liveUtils

import (
	"fmt"

	"github.com/luno/luno-go/decimal"
	"gopkg.in/gomail.v2" // perform go get <this> when initialising servers
	//  "time"
)

func Email(action string, yield decimal.Decimal) {
	//initialising email
	m := gomail.NewMessage()
	m.SetHeader("From", "profit.profit.profit.icl@gmail.com")
	m.SetHeader("To", "profit.profit.profit.icl@gmail.com") // can add multiple recievers
	var messageStr string

	switch action {
	// emailing graphed data
	case "GRAPH":
		messageStr = "Shivam's daily update: "
		if yield.Sign() == 1 {
			messageStr += "PROFIT! £" + yield.String()
		} else if yield.Sign() == -1 {
			yieldAbs := yield.Sub((decimal.NewFromFloat64(2, 2)).Mul(yield))
			messageStr += "LOSS! £" + yieldAbs.String()
		} else {
			messageStr += "FLAT! £" + yield.String()
		}
		//fileName := time.Now().Format("2006-01-02")
		m.Attach("../main/graph.png")
	//m.Attach("../main/" + fileName + ".xlsx")
	//emailing bot starting status
	case "START":
		messageStr = "NEWS! Shivam's bot has begun trading"
	}

	m.SetHeader("Subject", messageStr)
	m.SetBody("text/html", "")

	d := gomail.NewDialer("smtp.gmail.com", 587, "profit.profit.profit.icl@gmail.com", "Password123??")

	// function doesn't panic when error occurs
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("ERROR! graph.png doesn't exist ", err)
		fmt.Println("Update email NOT successfully sent")
		return
	}

	fmt.Println("Update email successfully sent")
}
