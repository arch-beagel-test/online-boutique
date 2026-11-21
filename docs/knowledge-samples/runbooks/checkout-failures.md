# Runbook: checkout failures

## Symptoms

- The "Place order" button returns an error page.
- `checkoutservice` logs `PlaceOrder failed` with a downstream service name.

## Diagnosis

1. Check which downstream call failed: payment, shipping, email or currency.
2. If `paymentservice` returns `InvalidCreditCard`, the card is rejected by design (the mock rejects cards that
   aren't VISA or MasterCard).
3. If `emailservice` is down, the order still succeeds; confirmation emails are best effort.

## Mitigation

Restart the failing deployment. If `shippingservice` is failing, set the shipping fallback quote to 8.99 USD.

## Rollback

Roll back `checkoutservice` to the previous image tag.
