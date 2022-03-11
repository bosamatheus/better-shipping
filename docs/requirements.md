# Better Shipping

## Description

Shipping costs are expensive.
Our team is working to validate a business concept to offer **better** options of shipping for our merchants and their consumers.
Your challenge is to build a PoC and validate the usability of our business concept.

## Functional requirements

We want a RESTful API that given different shipping options from our **external API** is able to recommend the same shipping options ordered by the best combination of cost and time.

## Technical requirements

Write down unit tests for the main logic.

The **response** must have the shipping option fields:

```
name: required, text
type: required, text
cost: required, decimal
estimated_days: required, numeric
```

## Considerations

- If there’s no shipping options available, the response should be an empty list;
- We can use a external API (e.g. `/api/v1/shipping-options`) endpoint in your implementation.
- We **must** use the acceptance criteria to write down **unit tests**, mocking the shipping options and asserting the expected response.

## Acceptance Criteria

1. Same shipping costs and estimated delivery dates.

a. Shipping options:

```json
[
  { "name": "Option 1", "type": "Delivery", "cost": 10.0, "estimated_days": 3 },
  { "name": "Option 2", "type": "Custom", "cost": 10.0, "estimated_days": 3 },
  { "name": "Option 3", "type": "Pickup", "cost": 10.0, "estimated_days": 3 }
]
```

b. Expected response:

```json
[
  { "name": "Option 1", "type": "Delivery", "cost": 10.0, "estimated_days": 3 },
  { "name": "Option 2", "type": "Custom", "cost": 10.0, "estimated_days": 3 },
  { "name": "Option 3", "type": "Pickup", "cost": 10.0, "estimated_days": 3 }
]
```

2. Same shipping costs, different estimated delivery dates.

a. Shipping options:

```json
[
  { "name": "Option 1", "type": "Delivery", "cost": 10.0, "estimated_days": 5 },
  { "name": "Option 2", "type": "Custom", "cost": 10.0, "estimated_days": 2 },
  { "name": "Option 3", "type": "Pickup", "cost": 10.0, "estimated_days": 3 }
]
```

b. Expected response:

```json
[
  { "name": "Option 2", "type": "Custom", "cost": 10.0, "estimated_days": 2 },
  { "name": "Option 3", "type": "Pickup", "cost": 10.0, "estimated_days": 3 },
  { "name": "Option 1", "type": "Delivery", "cost": 10.0, "estimated_days": 5 }
]
```

3. Same estimated delivery dates, different shipping costs.

a. Shipping options:

```json
[
  { "name": "Option 1", "type": "Delivery", "cost": 6.0, "estimated_days": 3 },
  { "name": "Option 2", "type": "Custom", "cost": 5.0, "estimated_days": 3 },
  { "name": "Option 3", "type": "Pickup", "cost": 10.0, "estimated_days": 3 }
]
```

b. Expected response:

```json
[
  { "name": "Option 2", "type": "Custom", "cost": 5.0, "estimated_days": 3 },
  { "name": "Option 1", "type": "Delivery", "cost": 6.0, "estimated_days": 3 },
  { "name": "Option 3", "type": "Pickup", "cost": 10.0, "estimated_days": 3 }
]
```

4. Both shipping costs and estimated delivery dates are different.

a. Shipping options:

```json
[
  { "name": "Option 4", "type": "Delivery", "cost": 10.0, "estimated_days": 3 },
  { "name": "Option 1", "type": "Delivery", "cost": 10.0, "estimated_days": 5 },
  { "name": "Option 2", "type": "Custom", "cost": 5.0, "estimated_days": 4 },
  { "name": "Option 3", "type": "Pickup", "cost": 7.0, "estimated_days": 1 }
]
```

b. Expected response:

```json
[
  { "name": "Option 2", "type": "Custom", "cost": 5.0, "estimated_days": 4 },
  { "name": "Option 3", "type": "Pickup", "cost": 7.0, "estimated_days": 1 },
  { "name": "Option 4", "type": "Delivery", "cost": 10.0, "estimated_days": 3 },
  { "name": "Option 1", "type": "Delivery", "cost": 10.0, "estimated_days": 5 }
]
```

5. No shipping options available.

a. Shipping options:

```json
[]
```

b. Expected response:

```json
[]
```
