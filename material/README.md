# Material Components For GopherJS

Material is GopherJS bindings for the [material-components-web][] JavaScript
library ([MDC](#Terminology)).

## Project Status

Material is under active development, and the API is likely to change often, at
any time. Input is encouraged regarding Material's design at this early stage.

The core Material [component package][Material component] is usable. Work is
underway to implement the specific Material component types, which is tracked in
[an issue][gl-issue-1].

## Getting Started

There is a quickstart guide in the [material package][material component]
documentation, and every component package's documentation page has a basic
example of how to set up and use it.

## Contributing

Any and all ideas, changes, bug reports, etc. are encouraged.

Get involved by:
- Opening issues and submitting merge requests at our [Gitlab project][Material]
  .
- Starting a discussion in the [Gopher Slack][] #GopherJS channel.

## Project Goals

- Expose an API that supports idiomatic Go programming while maintaining
  the general terminology of MDC resources.
- Minimal dependencies to keep generated JavaScript sizes as small as possible.
- Modular package organization so you can pick-and-choose components just like
  the MDC library.
- Limited scope. Only focus on wrapping the JavaScript functionality of the MDC
  library.

## Project Non-Goals

- HTML/CSS. These are areas of the MDC library that are best handled by projects
  that consume Material's packages. Check out
  [vecty-material][] as an example.

## Terminology

To help avoid ambiguity in code and documentation, here's a quick list of
definitions we are trying to enforce in Material.

| Term                        | Definition                                           |
| :---                        | :---------                                           |
| [material-components-web][] | The official upstream JavaScript library.            |
| MDC or MCW                  | Synonyms/abbreviations of `material-components-web`. |
| [Material][]                | This project.                                        |
| [Material component][]      | Our `github.com/vecty-material/material/*` Go packages.             |
| [Material menu][]           | An example of a specific type of Material component. |

[material-components-web]: https://github.com/material-components/material-components-web
[Material]: https://gitlab.com/agamigo/material
[Material component]: https://godoc.org/github.com/vecty-material/material
[Material menu]: https://godoc.org/github.com/vecty-material/material/menu
[vecty-material]: https://gitlab.com/agamigo/vecty-material
[Gopher Slack]: https://gophers.slack.com
[gl-issue-1]: https://gitlab.com/agamigo/material/issues/1
