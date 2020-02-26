# Compareto - A Mario Kart 8 Deluxe build comparator

Compareto uses the Pareto efficiency to determine all the optimal MK8DX builds based on the parts that you want to use, and the two criteria you want to optimize for.

- [Compareto - A Mario Kart 8 Deluxe build comparator](#compareto---a-mario-kart-8-deluxe-build-comparator)
- [Building](#building)
- [Running](#running)
  - [Basic usage](#basic-usage)
  - [Change optimization targets](#change-optimization-targets)
  - [Use a subset of characters or parts](#use-a-subset-of-characters-or-parts)

# Building

Run `make` from the root of this repo. This will bundle the files in `data/` into a `.go` file using [pkger](https://github.com/markbates/pkger/). The Go code is then compiled into the executable `compareto` at the root of the project.

# Running

Run `./compareto --help` for help:
```
Usage of ./compareto:
  -c string
    	Path the the characters CSV (default "/data/characters.csv")
  -c1 string
    	First criteria to use for the Pareto frontier (default "speed")
  -c2 string
    	Second criteria to use for the Pareto frontier (default "acceleration")
  -g string
    	Path to the gliders CSV (default "/data/gliders.csv")
  -t string
    	Path to the tires CSV (default "/data/tires.csv")
  -v string
    	Path to the vehicles CSV (default "/data/vehicles.csv")
```

## Basic usage

Running `compareto` with no arguments will compute the optimal builds for all characters and parts, by optimizing for speed and acceleration:

```
Character: Toadette. Wendy. Isabelle
Vehicle: Biddybuggy
Tires: Roller
Glider: Cloud Glider
Speed: 1.25	Acceleration: 5.75	Weight: 1.25
Handling: 5.00	Traction: 3.75		Miniturbo: 5.00
Total: 22.00

Character: Cat Peach. Inkling Girl. Villager Girl
Vehicle: Mr. Scooty
Tires: Roller
Glider: Paper Glider
Speed: 1.75	Acceleration: 5.50	Weight: 1.50
Handling: 4.75	Traction: 3.75		Miniturbo: 5.25
Total: 22.50

[...]

Character: Wario. Dry Bowser
Vehicle: Circuit Special
Tires: Slick
Glider: Waddle Wing
Speed: 5.75	Acceleration: 1.50	Weight: 4.75
Handling: 2.00	Traction: 1.50		Miniturbo: 1.25
Total: 16.75
```

## Change optimization targets

You can change which metrics are used for the optimization by using the `-c1` and `-c2` flags. Each one takes a string, that can be any of the following:

| Value          | Metric                                                                            |
| -------------- | --------------------------------------------------------------------------------- |
| `speed`        | Max speed                                                                         |
| `acceleration` | Acceleration (how quickly you reach your max speed)                               |
| `weight`       | Weight (how easily you will lose speed when hit)                                  |
| `handling`     | Handling (how easy is it to turn and drift)                                       |
| `traction`     | Traction (how much of your speed do you keep when off-road)                       |
| `miniturbo`    | Mini-turbo (how long does it take to get a mini-turbo, and how long it lasts for) |

**Note**: Some metrics are heavily correlated, for instance optimizing for Speed and Weight will only produce one result:
```
Character: Bowser. Morton
Vehicle: Standard ATV
Tires: Metal
Glider: Wario Wing
Speed: 5.75	Acceleration: 1.00	Weight: 5.75
Handling: 1.50	Traction: 2.50		Miniturbo: 1.00
Total: 17.50
```

## Use a subset of characters or parts


The executable produced by running `make` bundles the parts and characters data.

In order to limits which parts are used by the optimizer (if you don't have all the parts unlocked yet, or if you want to avoid certain characters for instance), make a copy of the files in `data/`, and remove the parts/characters you are not interested in. 

To have `compareto` take these modified files into account, use the following flags to give it the path to the modified file:

| Flag | Character/Part |
| ---- | -------------- |
| `-c` | Character      |
| `-v` | Vehicle        |
| `-t` | Tires          |
| `-g` | Glider         |

**Note**: not each character is unique, they are all part of a weight category. Characters with identical stats are listed on the same line in the `data/characters.csv` file. 