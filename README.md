# Bill Generator
A cli tool to generate Bills.

## Problem statement

Create a program that will ask for name input on bill.
### Give 3 options:
- Add item: Ask for item name and amount
- Save bill: save the items, tip and total into a txt file on fs
- Add tip: ask for tip amount

### Hints:
- Use package bufio to read user input
- Use os.WriteFile to save bill into a txt file