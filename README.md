This repository is an experiment for me to create a "shorter" ulid.

Why I want this:
- I want some of the features of ULID
  - Lexicographically Sortable Identifier
- ULID is long at 26 characters
  - I want something shorter, say 9 characters
  - To be used as "order ids"
- I don't need millisecond support
- I don't need as much entropy as ULID
- I don't need support till 10889 AD

Do not use this if you need a lot of entropy.

|                                 | ULID        | shortulid |
|---------------------------------|-------------|----------:|
| Characters                      |          26 |         9 |
| Characters reserved for time    |          10 |         6 |
| Characters reserved for entropy |          16 |         3 |
| Support till Year               |       10889 |      2054 |
| Bits of entropy                 |          80 |        15 |
| Precision                       | millisecond |    second |

Also See:

- https://github.com/ulid/spec
- https://github.com/oklog/ulid