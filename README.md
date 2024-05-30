# International and Vague Decision Engine

Given an item, an age, and a purchase price return a monthly price.

## Maffs

Prices are calculated by:

```math
q = \left( { \left(l - a\right) / 12 \over lp * rp } \right) p
```

Where `l` is the expected lifetime of the product, `a` is the current age.

`lp` is the lastability probability, and `rp` is the reparability probablity. These are the probabilities that the product will be used for its lifetime, and the probability of an error being repairable respectively.

`p` is the profaibility percentage.

## Configuration

```bash
$ DECISION_ENGINE_PROFITABILITY=1.15    # Multiplier for profitability; set too high nobody buys. Too low, we lose money
```

## grpcurl

```bash
$ grpcurl -d '{"expected_lifetime":13,"age":2,"purchase_price":200,"lastability":0.75,"repairability":0.9}' -plaintext localhost:8888 Quotes.GetQuote
{
  "monthly": 1.21
}
```
