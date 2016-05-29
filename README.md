CSV Explorer
============

The CSV Explorer is a command-line utility written in Go for dealing with CSV
files. Anyone who deals with data munging will know the importance of
inspecting data files for consistency, and this is a tool that can help.

# Installation
```bash
$ go get github.com/wyc/csvxp
$ go install github.com/wyc/csvxp
```

## Features and Examples

Please refer to this CSV data set in `file.csv` for the examples below:
```bash
Header1,Header2,Header3
a,b,c
1,2,3
```

Column headers of a CSV file
```bash
$ csvxp -column-headers file.csv
Header1
Header2
Header3
```

Values of a single column
```bash
$ csvxp -column-name "Header1" file.csv
a
1
$ csvxp -column-name "Header3" file.csv
c
3
```

Values of multiple columns
```bash
$ csvxp -column-names "Header1","Header3" file.csv
a,c
1,3
```

Calculating `title min mean max` column value length statistics across all rows:
```bash
$ csvxp -column-stats file.csv
"Header1" 1 1 1
"Header2" 1 1 1
"Header3" 1 1 1
```
