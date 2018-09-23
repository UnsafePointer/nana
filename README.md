# Nana (ナナ)

> “The more my dreams are fulfilled, the quicklier they become realities losing their shine.”
>
> [Nana (manga), by Ai Yazawa](https://en.wikipedia.org/wiki/Nana_(manga))

Nana (ナナ) is a [Game Boy](https://en.wikipedia.org/wiki/Game_Boy) emulator written in Go. It's far from feature complete and there are still bugs (see screeshot) in what it's implemented but what I got here helped me a lot to understand how computers work at one of their lowest levels, I hope this too helps someone else.

![nana-screenshot.png](nana-screenshot.png)

### Dependencies

Linux dependencies are listed [here](./Dockerfile).

### Installation

```
$ go get github.com/Ruenzuo/nana
```

### Usage

```
$ nana path/to/rom.gb
```

### Environment variables

* `DEBUG`: enables buffered debug output to nana.log
* `MAX_CYCLES`: enabled automatic shut down to specific cycle threshold

### Building your own emulator

* Start with 8-bit system
* Performance matters, profile constantly against other emulators to check how you're doing so far
* Implement the [full instruction set](./emulator/instruction_set.go) + [jump table](./emulator/jump_table.go) **first**. Then try your work against an [emulation test collection](https://github.com/retrio/gb-test-roms), these are excelent to catch edge cases and implement the bugs the original systems had
* [Game Boy CPU Manual](http://www.codeslinger.co.uk/pages/projects/gameboy/files/GB.pdf)
* [Game Boy Programming Manual](https://archive.org/download/GameBoyProgManVer1.1/GameBoyProgManVer1.1.pdf)
