# Parse AOC '23 Day 2

This is practice parsing using
[participle](https://github.com/alecthomas/participle) to parse the
[Advent of Code Day 2 problem](https://adventofcode.com/2023/day/2) for 2023.

## Running the tool

This practice includes a sample file `sample.aoc` that can be passed to the main
file. This can be run as:

> go run main.go sample.aoc

The output is a JSON representation of what participle parsed.

## Format description

The Advent of Code website describes the format as:

> You play several games and record the information from each game (your puzzle
> input). Each game is listed with its ID number (like the 11 in Game 11: ...)
> followed by a semicolon-separated list of subsets of cubes that were revealed
> from the bag (like 3 red, 5 green, 4 blue).
>
> For example, the record of a few games might look like this:

```
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green Game 2: 1 blue, 2
green; 3 green, 4 blue, 1 red; 1 green, 1 blue Game 3: 8 green, 6 blue, 20
red; 5 blue, 4 red, 13 green; 5 green, 1 red Game 4: 1 green, 3 red, 6 blue; 3
green, 6 red; 3 green, 15 blue, 14 red Game 5: 6 red, 1 blue, 3 green; 2 blue,
1 red, 2 green
```

> In game 1, three sets of cubes are revealed from the bag (and then put back
> again). The first set is 3 blue cubes and 4 red cubes; the second set is 1 red
> cube, 2 green cubes, and 6 blue cubes; the third set is only 2 green cubes.
