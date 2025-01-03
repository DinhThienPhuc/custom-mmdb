# Custom MMDB Generator

A Go-based tool for enriching MaxMind DB files with custom IP range data.

## Overview

This tool extends MaxMind's GeoLite2 Country database by adding custom department and environment access data for IP ranges. It demonstrates how to:

- Load and modify existing MMDB files
- Add custom data to specific IP ranges
- Generate enriched MMDB files with combined data

## Features

- Load existing GeoLite2 Country MMDB files
- Add custom department data to IP ranges
- Define environment access levels
- Generate hybrid MMDB files
- Fast IP address range queries

## Prerequisites

- Go 1.14+
- Git
- [mmdbinspect](https://github.com/maxmind/mmdbinspect)
- [GeoLite2 Country](https://dev.maxmind.com/geoip/geoip2/geolite2/) database

## Installation

```bash
# Clone the repository
git clone https://github.com/DinhThienPhuc/Purr.git

# Navigate to the app directory
cd apps/custom-mmdb

# Install dependencies
go mod download
```

## Usage

```bash
# Build the application
go build

# Run the generator
./custom-mmdb
```

## How It Works

1. Loads the GeoLite2 Country MMDB
2. Adds custom department data for specified IP ranges
3. Creates a new enriched MMDB file
4. Validates the enriched data using mmdbinspect

For more details about MMDB file format and manipulation, see the [MaxMind DB Specification](https://github.com/maxmind/MaxMind-DB/blob/main/MaxMind-DB-spec.md).

## License

MIT

## Author

@DinhThienPhuc
