# Alert, a very simple CLI reminder

Created so I don't forget whatever I have on the stove or in the oven.
Will send a terminal bell (window highlight) when a timer elapses.
May or may not work with your desktop environment or window manager.

## Installation

```
$ go get -u github.com/fgahr/alert
```

## Usage

```
$ alert help
usage: alert <cmd> [args...]

commands:
        help              Print this message
        in    <duration>  Alert when the given duration has elapsed
        at    <time>      Alert at the given time
$ alert in 2s
Alerting in 2s at 2021-03-21T21:24:54
Timer elapsed
```

Format for the duration specifier is the one used by
[Golang's stdlib](https://golang.org/pkg/time/#ParseDuration) so most
combinations of `h`, `m`, `s`, `ms`, `ns` (for whatever reason) should work,
e.g. `2h35m8s1ms`.

The `at` command is working with the same format as outputted on announcing
a timer, so e.g.

```
$ alert at 2021-03-21T21:24:54
Alerting in 2s at 2021-03-21T21:24:54
Timer elapsed
```

I don't expect much need for this one and I might update it for better
usability when the need arises.
