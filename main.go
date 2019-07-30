package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

// CheckoutData ...
type CheckoutData struct {
	ClientSecret string
}

func main() {
	stripe.Key = "sk_test_VFPOtLv8MJnV1a56b4UswQfh00HdXjfEVF"

	//acct_1F0WADCRCcEx9KXQ

	/*	params := &stripe.PaymentIntentParams{
			Amount:   stripe.Int64(1099),
			Currency: stripe.String(string(stripe.CurrencyEUR)),
		}

	*/

	params := &stripe.PaymentIntentParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		Amount:   stripe.Int64(1000),
		Currency: stripe.String(string(stripe.CurrencyEUR)),
		TransferData: &stripe.PaymentIntentTransferDataParams{
			Amount:      stripe.Int64(877),
			Destination: stripe.String("acct_1F0WADCRCcEx9KXQ"),
		},
	}
	intentA, err := paymentintent.New(params)

	if err != nil {
		fmt.Println("there was an ERROR")
		panic(err)
	} else {
		fmt.Println(intentA.Amount)
		fmt.Println(intentA.ClientSecret)
	}

	checkoutTmpl := template.Must(template.ParseFiles("views/checkout.html"))

	http.HandleFunc("/checkout", func(w http.ResponseWriter, r *http.Request) {
		intent := intentA
		data := CheckoutData{
			ClientSecret: intent.ClientSecret,
		}
		checkoutTmpl.Execute(w, data)
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.ListenAndServe(":3666", nil)

}
