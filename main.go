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
	stripe.Key = ""

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(1099),
		Currency: stripe.String(string(stripe.CurrencyEUR)),
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

	http.ListenAndServe(":3000", nil)

}
