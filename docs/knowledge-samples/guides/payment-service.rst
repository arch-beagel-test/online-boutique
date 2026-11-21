Payment service
===============

``paymentservice`` is written in Node.js and charges a credit card for an order.

What it accepts
---------------

* VISA and MasterCard only. Any other card type raises ``UnacceptedCreditCard``.
* Cards that have expired raise ``ExpiredCreditCard``.
* The amount and currency come from ``checkoutservice`` after currency conversion.

What it returns
---------------

A transaction ID, a version 4 UUID. No card number is stored or logged; only the last four digits appear in logs.

Configuration
-------------

``PORT`` (default 50051) and ``DISABLE_PROFILER``.
