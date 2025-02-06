# Fetch Interview Task
Please visit the [challenge repository](https://github.com/fetch-rewards/receipt-processor-challenge/tree/main?tab=readme-ov-file) if you want more information on this challenge!

## Getting started
There are no dependencies other than Go. The version used for this task is `1.23.6`.

The following commands below should get you going:
```
go mod download
go mod verify
go run .
```

## Closing Thoughts
I'm wondering if there is a better approach than implementing the rules in a functional manner. Theoretically, if more rules are added overtime, my solution has the potential to not scale very well.

But it was important to me to also keep things simple and not over-engineer a solution.