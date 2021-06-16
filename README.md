# Roman [![Test Status](https://github.com/imbue11235/roman/workflows/Go/badge.svg)](https://github.com/imbue11235/roman/actions?query=workflow:Go)
Package `roman` provides functions for converting arabic to roman numerals and back

## Motivation
Needed for a work project, and with the lack of a solution that suited my exact needs, I decided to publish this micro package

## Installation
```sh
$ go get github.com/imbue11235/roman
```

## Usage

### Converting arabic to roman numerals

```go
numeral := roman.FromArabic(4) // => "IV"
```

### Converting roman to arabic numerals

```go
numeral := roman.ToArabic("IV") // => 4
```

## Limitations

- Arabic numerals lower than or equal to `0` cannot be converted and will return an empty string.
- Arabic numerals higher than `3999` cannot be converted and will be return an empty string, as:
```
The largest number that can be represented in this notation is 3,999 (MMMCMXCIX)
```
([Wiki for more information](https://en.wikipedia.org/wiki/Roman_numerals))

## License

This project is licensed under the [MIT license](LICENSE).