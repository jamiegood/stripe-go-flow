$( document ).ready(function() {
 
  // Your code here.
  console.log('in index.j');
  var stripe = Stripe('pk_test_Vt9YD59W6N1U93C1PtkIQpfx00qc0R9JDZ');
  var elements = stripe.elements();
  var cardElement = elements.create('card');
  cardElement.mount('#card-element');
  
  var cardholderName = document.getElementById('cardholder-name');
  var cardButton = document.getElementById('card-button');
  var clientSecret = cardButton.dataset.secret;

  //4000008260000000
  cardButton.addEventListener('click', function(ev) {
    console.log("cardholderName.value:: " + cardholderName.value);
    console.log(clientSecret);

    stripe.handleCardPayment(
      clientSecret, cardElement, {
        payment_method_data: {
          billing_details: {name: cardholderName.value}
        }
      }
    ).then(function(result) {
      if (result.error) {
        // Display error.message in your UI.
        console.log("ERROR FROM Stripe");

        console.log(result.err);
      } else {
        // The payment has succeeded. Display a success message.
        console.log("payment succeeded");

      }
    });    
  });

});


/*
cardButton.addEventListener('click', function(ev) {
  stripe.handleCardPayment(
    clientSecret, cardElement, {
      payment_method_data: {
        billing_details: {name: cardholderName.value}
      }
    }
  ).then(function(result) {
    if (result.error) {
      // Display error.message in your UI.

    } else {

      // The payment has succeeded. Display a success message.
    }
  });
*/