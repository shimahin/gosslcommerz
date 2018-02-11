## Integration Steps
#Initiate Payment
Provide Information about your customer and order to SSLCommerz along with your store id to initiate the payment.
Rest of the payment process will be done by SSLCommerz

#Validate Payment
After successfully taking the payment SSLCommerz will send the request back to you as SUCCESS, FAILED or CANCEL status.
You must validate with our validation API using transaction ID, amount and currency

#Update your transaction
After validation of the transaction that you have received, Depending on the status you have to update your
transaction in your Database. The status will SUCCESS, FAILED, CANCEL depending on BANK Status

#Enable most advanced IPN
This is an important and interesting part of integration. If somehow your consumer pays your payable amount to
BANK Side and SSLCommerz accept it as SUCCESS but your website/Connectivity/Customer Network got downtime and unable to
update the payment at your side you can use IPN ( Instant Payment Notification ).
It will send an notification to your set up URL in SSLCommerz Merchant Dashboard to notify you and your database even if
your user unable to return back to your website