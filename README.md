# csvToCamt
Convert CSV correction files to CAMT format

Get binaries from [release page](https://github.com/krishnaprasadmg/csvToCamt/releases)

# Example

Execute the command with config file (see [sample](config.yaml)) and CSV correction files

```
$./csvToCamt -c config.yaml q3.csv q4.csv
```

Optionally use the `-s` flag to skip the headers from CSV files (If the CSVs have headers)