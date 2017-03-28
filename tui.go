/*
<!--
Copyright (c) 2017 Christoph Berger. Some rights reserved.

Use of the text in this file is governed by a Creative Commons Attribution Non-Commercial
Share-Alike License that can be found in the LICENSE.txt file.

Use of the code in this file is governed by a BSD 3-clause license that can be found
in the LICENSE.txt file.

The source code contained in this file may import third-party source code
whose licenses are provided in the respective license files.
-->

<!--
NOTE: The comments in this file are NOT godoc compliant. This is not an oversight.

Comments and code in this file are used for describing and explaining a particular topic to the reader. While this file is a syntactically valid Go source file, its main purpose is to get converted into a blog article. The comments were created for learning and not for code documentation.
-->

+++
title = ""
description = ""
author = "Christoph Berger"
email = "chris@appliedgo.net"
date = "2017-03-28"
publishdate = "2017-03-28"
draft = "true"
domains = ["User Interface"]
tags = ["TUI", "Console", "Terminal", "UI"]
categories = [""]
+++

Want to equip your terminal application with a nice visual user interface? TUI libraries are here to help.

<!--more-->

Console applications usually take some parameters at start, and maybe some more input through basic console I/O. And that's ok in most cases, though sometimes it would be great to have a visual user interface but without the code size and complexity of a full-blown Web app.

[Text-Based User Interface](https://en.wikipedia.org/wiki/Text-based_user_interface libraries (or TUI libraries) meet this need. They bring panes, input, output, mouse support, graphics, and audio, to your terminal.


## OK, so can we do a side-by-side test, please?

TL;DR: Unfortunately, no.

Originally, I planned to evaluate three to four TUI libraries for Go side-by-side, using a sample app definition with a specific set of features to be tested. I had a few requirements on the libraries: They should -

* be fairly complete TUI libraries (as opposed to say, just drawing a progress bar and nothing else),
* provide high-level abstractions,
* be active projects (that is, the last commit should not date back to months ago),
* be past the alpha stage, and
* have reasonable documentation or sample code. (No, API docs don't count. I wanted at least a simple how-to-start tutorial that covers the basics.)

Turned out that none of the libraries passed all requirements.


## So that's it? No single all-in-one, high-level GUI lib with decent documentation?

Well, not if you really want all those nifty UI features at once. However, if you can forego some of the requirements, there is indeed a small selection of capable TUI libraries at your disposal.


### Bare bone TUI libraries

If you are fine with building your UI on top of a handful of UI primitives, there are two libraries that look fairly stable and complete:


#### [nsf/termbox-go](https://github.com/nsf/termbox-go)

`termbox` has a small and straightforward [API](https://godoc.org/github.com/nsf/termbox-go). It uses a two-dimensional cell buffer to represent the terminal screen, with a few methods like setting cells, clearing the screen, or setting the cursor. All input, whether keyboard, mouse, or resize events, is tracked through a unified Event type.
As simple as the API might seem, it has not prevented others from creating interesting projects on top of `termbox`, including high-level TUI libraries (`termui`, `gocui`, and `termloop``), simple games, [godit](https://github.com/nsf/godit) (a text editor) and [hecate](https://github.com/evanmiller/hecate) (a hex editor), and more.

#### [gdamore/tcell](https://github.com/gdamore/tcell)

According to the author, `tcell` was born out of the need for certain features that `termbox-go` does not provide, and patching `termbox-go` turned out to not be a suitable way to go forward. `tcell` claims to have a couple of advantages over `termbox-go`, including [added functionality](https://godoc.org/github.com/gdamore/tcell), better portability, better support for mouse, Unicode, colors, and more. (I cannot comment on these claims, as I have tested neither `termbox` nor `tcell` in great depth.)
Tcell is the TUI library behind the [Micro editor](https://github.com/zyedidia/micro/), and is also used by [godu](https://github.com/viktomas/godu).

I should note here that both libraries require `cgo`.

### Higher-level libraries

#### [gizak/termui](https://github.com/gizak/termui)

`termui` is one of the high-level TUI's that are built on top of `termbox-go`. It is specialized on displaying information in graphical form - as bar chart, line diagram, spark line, or as a gauge. Text output is also possible of course. The UI is built on the concept of widgets, or blocks, that contain exactly one type of data representation - text, charts, or diagrams.

![termui dashboard demo](dashboard.gif)

`termui` provides a static layout and a grid layout. The grid layout can respond dynamically to resizing of the terminal window.

On the input side, `termui` offers a simple event handling system that can react to key presses, window resizes, or timers.

What is missing from `termui` to make it a "full" TUI toolikt is high-level input handling (input boxes, buttons, menus, drop-down lists etc). I also missed some sort of resizeable panes, although the grid layout probably provides a suitable alternative to that.

A couple of projects have already been built with `termui`, including [ctop](https://github.com/bcicen/ctop).


#### [jroimartin/gocui](https://github.com/jroimartin/gocui)

This library almost seems like the opposite of `termui`: No fancy widgets anywhere, but instead, building blocks for split views, overlapping views, editable views, and event handling for keyboard and mouse events.

![gocui real-life screenshot - httplab](httplab.gif)


#### [clui](https://github.com/VladimirMarkelov/clui)

From the screenshots, this is perhaps the best-looking TUI library in this list. It also looks fairly advanced already (featuring a rich set of UI widgets like buttons, drop-down lists, gauges, etc. as well as a window manager and theming), however, the lack of documentation (except the API docs) indicates that this is still a work in progress.

![clui screencast](clui.gif)

#### [wm](https://github.com/cznic/wm)

Definitely a work in progress, `wm` manages overlapping, resizable windows with scrollable content.

![wm screenshot](tk.png)

#### [termloop](https://github.com/JoelOtter/termloop)

This is maybe an edge case. `termloop` is a very specialized TUI - it is a game toolkit for the terminal. So perhaps you won't be able to build a "standard" application UI on top of it, but you surely can make your window rectangles make a sound when colliding! 😉

![termloop screenshot](termloop.png)

##

## The code
*/

// ## Imports and globals
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gizak/termui"
	"github.com/jroimartin/gocui"
)

func runTermui() {
	err := termui.Init()
	if err != nil {
		log.Fatalln("Cannot initialize termui")
	}
	defer termui.Close()

	p := termui.NewPar("\nPress Q to quit")
	p.Height = 5
	p.Width = 50
	p.BorderLabel = "Info"
	p.TextFgColor = termui.ColorWhite
	p.BorderFg = termui.ColorCyan

	termui.Render(p)

	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})

	termui.Loop()
}

func layout(g *gocui.Gui) error {
	width, height := g.Size()
	view, err := g.SetView("main", 4, 4, width-4, height-4)
	return err
}

func quit(g *gocui.Gui) error {
	return gocui.ErrQuit
}

func runGocui() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		fmt.Println("Failed to create a GUI:", err)
		return
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	err = g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit)
	if err != nil {
		fmt.Println("Could not set key binding:", err)
		return
	}

	err := g.MainLoop()
	fmt.Println("Main loop has finished: ", err)
}

func main() {
	if len(os.Args) == 0 {
		fmt.Println("Usage: go run tui.go [termui|gocui]")
	}
	if os.Args == "termui" {
		runTermui()
		return
	}
	if os.Args == "gocui" {
		runGocui()
		return
	}
	fmt.Println("No such option:", os.Args)
}

/*
## How to get and run the code

Step 1: `go get` the code. Note the `-d` flag that prevents auto-installing
the binary into `$GOPATH/bin`.

    go get -d github.com/appliedgo/TODO:

Step 2: `cd` to the source code directory.

    cd $GOPATH/src/github.com/appliedgo/TODO:

Step 3. Run the binary.

    go run TODO:.go


## Odds and ends
## Some remarks
## Tips
## Links


**Happy coding!**

*/
