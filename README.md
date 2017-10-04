# Markov Chain Text Generator

## Summary

Basic implementation of [Markov chain](https://en.wikipedia.org/wiki/Markov_chain) text generator. Bit on a sloppy side, but I needed something quick to generate bunch of dummy text for makeup and layout, so here we are.

## Installation

```bash
$ go get github.com/eiri/mchain
$ go install
```

## Usage

```
$ mchain -h
Usage of mchain:
  -count int
        number of words to generate (default 25)
  -input string
        path to source text

$ mchain --input ~/Downloads/samples/rrc.txt --count 120
When this kidnapping idea struck us. It was, as Bill was, and as good a runner
as I am, he was a fire burning behind the big rock at the rising of the
external outward surface of Alabama that lay exposed to my view. "Perhaps,"
says I to myself, "it has not yet been discovered that the wolves have borne
away the tender lambkin from the box under the tree, got the knife away from
the boy amused and quiet till I tell you. About two miles from Summit was a
sling that Red Chief had said I could about how the kidnapping had been
kidnapped and chained to a bed in Bedlam. Besides being a thorough gentleman,
I think can.
```

## Licence

[MIT](https://github.com/eiri/mchain/blob/master/LICENSE)
