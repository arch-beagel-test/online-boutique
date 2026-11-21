# Features

What the boutique does for shoppers. Each feature has an ID that issues and pull requests cite.

| ID | Feature | Owner | Entry point |
|---|---|---|---|
| BTQ-BROWSE | Browse and search the product catalogue | Storefront team | `GET /` and `GET /product/{id}` |
| BTQ-CART | Add to cart, view and empty the cart | Storefront team | `POST /cart`, `GET /cart`, `POST /cart/empty` |
| BTQ-CHECKOUT | Place an order: pay, ship and confirm by email | Payments team | `POST /cart/checkout` |
| BTQ-CURRENCY | Show prices in the shopper's currency | Storefront team | `POST /setCurrency` |
| BTQ-RECO | Recommend products on the product and cart pages | Discovery team | `recommendationservice.ListRecommendations` |
| BTQ-ADS | Show a contextual ad | Discovery team | `adservice.GetAds` |
| BTQ-ASSIST | Shopping assistant answers questions about products | Discovery team | `POST /bot` |
