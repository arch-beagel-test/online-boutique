Shipping quotes
===============

``shippingservice`` (Go) gives a quote and ships an order.

Quote
-----

The quote is a fixed 8.99 USD per order in this demo, whatever the number of items. A real carrier integration is
planned behind the same ``GetQuote`` call.

Ship
----

``ShipOrder`` returns a tracking ID made from the address and a random part. It never fails for a valid address.
