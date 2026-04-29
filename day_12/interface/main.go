package main

import "fmt"

// ------------ payment gateway --------------

type razorpay struct{}

func (r razorpay) pay(amount float32){
	// login to make payment
	fmt.Println("making payment using razorpay")
}

type stripe struct {}

func (s stripe) pay(amount float32){
	fmt.Println("making payment using stripe")	
}

// ------------- user flow ---------------

type payment struct {}

func (p payment) makePayment(amount float32){
	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)

	stripePaymentGw := stripe{}
	stripePaymentGw.pay(amount)
}

func main(){
	newPayment := payment{}
	newPayment.makePayment(100.0)
}

