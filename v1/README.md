# Example version 1

This is the starting point of the exercise.
Let's see the code

Let's say we runs `a company of theatrical players`. and we have
these data

- playsJSON: We can perform three plays. Hamlet, As You Like It, and Othello
- Invoice: Our customer name is `BigCo`, and we had three performances for it.
  and we have audience number for each play.

```go
var (
	playsJSON string = `{
		"hamlet":  {"name": "Hamlet", "type": "tragedy"},
		"aslike":  {"name": "As You Like It", "type": "comedy"},
		"othello": {"name": "Othello", "type": "tragedy"}
	}`

	invoiceJSON string = `[
		{
			"customer": "BigCo",
			"performances": [
				{
					"playID": "hamlet",
					"audience": 55
				},
				{
					"playID": "aslike",
					"audience": 35
				},
				{
					"playID": "othello",
					"audience": 40
				}
			]
		}
	]`
)
```

This programs goal is printing the information for charging our customer,
`BigCo` about performing for them.

- Charging amount for each play and the sum
- Credits for the customer

```text
Statement for BigCo
	Hamlet: $650.00 (55 seats)
	As You Like It: $580.00 (35 seats)
	Othello: $500.00 (40 seats)
Amount owed is $1,730.00
You earned 47 credits
```

So, we parse the JSON, and calculate charging amount and credits
