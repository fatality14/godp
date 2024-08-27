# Godp CLI Tool Usage

The `godp` CLI tool allows users to compare packages between two branches. It provides various options to load data from an API or local files, save data to local files, and log comparison statistics.

## Table of Contents
- [Installation](#installation)
- [Basic Usage](#basic-usage)
- [Options](#options)
- [Use Cases](#use-cases)
  - [Comparing Packages from the API](#comparing-packages-from-the-api)
  - [Comparing Packages from Local Files](#comparing-packages-from-local-files)
  - [Saving Data to Local Files](#saving-data-to-local-files)
  - [Logging Comparison Statistics](#logging-comparison-statistics)
- [Error Handling](#error-handling)

## Installation

To install the `godp` CLI tool, clone the repository and build the tool using Go:

```bash
git clone https://github.com/fatality14/godp.git
cd godp
go build -o godp ./main/main.go
```

## Basic Usage

```bash
./godp [OPTIONS] <branch1> <branch2>
```

branch1 and branch2 are the two branches whose packages you want to compare.

## Options

    -api-route: Specify the API route URL to fetch data from.
      (default: https://rdb.altlinux.org/api/export/branch_binary_packages)
    -from-files: Load packages from local files instead of the API.
    -save-files: Save the loaded packages to local files for future use.
    -log-stats: Log the branch comparison statistics.

## Use Cases

### Comparing Packages from the API

By default, the tool fetches package data from an API for the two specified branches and compares them. This is useful when you want to get the latest data from the server.

```bash
./godp sisyphus p10
```

### Comparing Packages from Local Files

Use the -from-files option to load package data from local files instead of the API. This is useful when you have previously saved data and want to avoid network calls.

```bash
./godp --from-files sisyphus p10
```

### Saving Data to Local Files

Use the -save-files option to save the fetched or loaded data to local files. This is useful if you want to cache the data for future use or for offline comparison.

```bash
./godp --save-files sisyphus p10
```

### Logging Comparison Statistics

Use the -log-stats option to log detailed statistics about the comparison. This includes the number of packages unique to each branch and packages with higher versions in each branch.

```bash
./godp --log-stats sisyphus p10
```

## Error Handling

The tool provides error handling for various scenarios:

    If the log file cannot be opened or created, the program exits with an error.
    If there is an error fetching data from the API or reading from local files, the tool logs the error and exits.
    If JSON marshaling of results fails, the tool logs the error and exits.

Make sure to check the log file app.log for detailed error messages and debugging information.
