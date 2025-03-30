# Wikiroute

Returns the [Six Degrees of Separation](https://en.wikipedia.org/wiki/Six_degrees_of_separation) path between two Wikipedia articles.

<br />

## Algorithm

### Single-directional BFS

```
        Depth = 0        →        S
        Depth = 1        →       ○○○
        Depth = 2        →      ○○○○○
        Depth = 3        →     ○○○○○○○
        Depth = 4        →    ○○○○○○○○○
        Depth = 5        →   ○○○○○○○○○○○
        Depth = 6        →  ○○○○○○○○○○○○○
```
Time complexity: O(b^6)

### Bidirectional BFS
From Start Side: From Goal Side:
```
  From Start Side:              From Goal Side:

        S                           G
       ○○○                        ○○○
      ○○○○○                      ○○○○○
     ○○○○○○○                    ○○○○○○○
        ↓                          ↑
        ↓←←←←←←←←←←→→→→→→→→→→→→→→↑
            Meet in the middle!
```
Time complexity: O(b^3 + b^3) = O(2b^3)

<br />

## Setup

### 1. Download Wikipedia dump files
```sh
make download-all DATE=20250301
```

### 2. Generate parsed data
```sh
make generate-all
```

<br />

## Run the application
```sh
docker compose up
```
