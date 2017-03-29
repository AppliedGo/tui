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

Well, not if you really want all those nifty UI features at once. However, if you can forego some of the requirements, there is indeed a small selection of capable TUI libraries at your disposal that I'll briefly introduce below.

Just one thing before we start.

The list is (and cannot) be complete. I used GitHub.com, awesome-go.com, GolangLibs.com, and LibHunt.com to find TUI libraries but still there is no guarantee that I found all relevant projects out there.

Without the tips of a couple of redditors in [this `/r/golang` thread](https://www.reddit.com/r/golang/comments/5zx4xb/suggest_me_your_favourite_terminal_gui_libs/), `clui` and `wm` would perhaps have slipped through.

And then, of course, the above list of criteria is a very fuzzy filter, and you might have included a library that I ruled out because of this list.

But now let's move on to the TUI library overview.


### Bare bone TUI libraries

If you are fine with building your UI on top of a handful of UI primitives, there are two libraries that look fairly stable and complete:


#### [nsf/termbox-go](https://github.com/nsf/termbox-go)

`termbox` has a small and straightforward [API](https://godoc.org/github.com/nsf/termbox-go). It uses a two-dimensional cell buffer to represent the terminal screen, with a few methods like setting cells, clearing the screen, or setting the cursor. All input, whether keyboard, mouse, or resize events, is tracked through a unified Event type.

As simple as the API might seem, it has not prevented others from creating interesting projects on top of `termbox`, including higher-level TUI libraries (`termui`, `gocui`, `tui-go` and `termloop``), simple games, [godit](https://github.com/nsf/godit) (a text editor) and [hecate](https://github.com/evanmiller/hecate) (a hex editor), and more.

The README makes no claims about supported platforms. The code includes some `syscall_*.go` files, indicating that `termbox-go` runs on Linux, macOS, Windows, OpenBSD, NetBSD, FreeBSD, and Dragonfly BSD.

`termbox` requires `cgo` for doing low-level calls to system libraries.


#### [gdamore/tcell](https://github.com/gdamore/tcell)

According to the author, `tcell` was born out of the need for certain features that `termbox-go` does not provide, and patching `termbox-go` turned out to not be a suitable way to go forward. `tcell` claims to have a couple of advantages over `termbox-go`, including [added functionality](https://godoc.org/github.com/gdamore/tcell), better portability, better support for mouse, Unicode, colors, and more. (I cannot comment on these claims, as I have tested neither `termbox` nor `tcell` in great depth.)
Tcell is the TUI library behind the [Micro editor](https://github.com/zyedidia/micro/), and is also used by [godu](https://github.com/viktomas/godu).

According to the README, `tcell` works on POSIX systems that provide a POSIX termios implementation with `/dev/tty`, as well as on Windows.

The "POSIX" requirement includes Linux, macOS, FreeBSD, and Solaris, and certainly more systems that are not explicitly listed in the README.

Like `termbox`, `tcell` also requires `cgo`.


### Higher-level libraries


#### [gizak/termui](https://github.com/gizak/termui)

`termui` is one of the high-level TUI's that are built on top of `termbox-go`. It is specialized on displaying information in graphical form - as bar chart, line diagram, spark line, or as a gauge. Text output is also possible of course. The UI is built on the concept of widgets, or blocks, that contain exactly one type of data representation - text, charts, or diagrams.

![termui dashboard demo](dashboard.gif)

`termui` provides a static layout and a grid layout. The grid layout can respond dynamically to resizing of the terminal window.

On the input side, `termui` offers a simple event handling system that can react to key presses, window resizes, or timers.

What is missing from `termui` to make it a "feature-complete" TUI toolkit is high-level input handling (input boxes, buttons, menus, drop-down lists etc). I also missed some sort of resizeable panes, although the grid layout probably provides a suitable alternative to that.

A couple of projects have already been built with `termui`, including [ctop](https://github.com/bcicen/ctop).


#### [jroimartin/gocui](https://github.com/jroimartin/gocui)

This library almost seems like the opposite of `termui`: No fancy widgets anywhere, but instead, building blocks for split views, overlapping views, editable views, and event handling for keyboard and mouse events.

![gocui real-life screenshot - httplab](httplab.gif)


### Promising projects

I couldn't resist including the following projects here, although they do not meet  some of the inclusion criteria. Still, they look very promising and may offer one or the other feature that you might have missed in this list so far.


#### [VladimirMarkelov/clui](https://github.com/VladimirMarkelov/clui)

From the screenshots, this is perhaps the best-looking TUI library in this list. It also looks fairly advanced already (featuring a rich set of UI widgets like buttons, drop-down lists, gauges, etc. as well as a window manager and theming), however, the lack of documentation (except the API docs) indicates that this is still a work in progress.

![clui screencast](clui.gif)


#### [cznic/wm](https://github.com/cznic/wm)

Definitely a work in progress, `wm` manages overlapping, resizeable, decorated windows with scrollable content.

![wm screenshot](tk.png)


#### [marcusolsson/tui-go](https://github.com/marcusolsson/tui-go)

The building blocks of `tui-go` are widgets and layout boxes. There seems no support for mouse-based resizing of layout boxes, but each layout box can assume one of (currently) two automatic resizing behaviors.


#### [termloop](https://github.com/JoelOtter/termloop)

This is maybe an edge case. `termloop` is a very specialized TUI - it is a game toolkit for the terminal. So perhaps you won't be able to build a "standard" application UI on top of it, but you surely can make your window rectangles make a sound when colliding! 😉

![termloop screenshot](termloop.png)


## The code

As always, there is some code included to try out at your end. This time, the code uses the two libraries "termui" and "gocui" for creating a minimal UI with a simple layout:

* A list box with a fixed width at the left side
* A text entry box with a fixed height at the bottom
* A general-purpose output pane in the remaining area

If the libary provides a text entry widget, text entered there shall appear in the output pane.
*/

// ## Imports and globals
//
// Importing two UI libraries at the same time is probably not a good idea, as
// each has its own event loop. This demo code takes care of using only one
// of these libraries at a time.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gizak/termui"
	"github.com/jroimartin/gocui"
	"github.com/pkg/errors"
)

const (
	// List box width.
	lw = 20
	// Input box height.
	ih = 3
)

// Items to fill the list with.
var listItems = []string{
	"Line 1",
	"Line 2",
	"Line 3",
	"Line 4",
	"Line 5",
}

/*
## termui

Let's start with `termui`.

*/

//
func runTermui() {
	// Initialize termui.
	err := termui.Init()
	if err != nil {
		log.Fatalln("Cannot initialize termui")
	}
	// termui needs some cleanup when terminating.
	defer termui.Close()

	// Get width and height of the terminal.
	//tw := termui.TermWidth()
	th := termui.TermHeight()

	// The list pane
	lp := termui.NewList()
	lp.Height = th
	lp.BorderLabel = "List"
	lp.BorderLabelFg = termui.ColorGreen
	lp.BorderFg = termui.ColorGreen
	lp.ItemFgColor = termui.ColorWhite
	lp.Items = listItems

	// The input pane. termui has no edit box yet, but at the time of
	// this writing, there is an open pull request for adding
	// a text input widget.
	ip := termui.NewPar("")
	ip.Height = ih
	ip.BorderLabel = "Input"
	ip.BorderLabelFg = termui.ColorYellow
	ip.BorderFg = termui.ColorYellow
	ip.TextFgColor = termui.ColorWhite

	// The Output pane
	op := termui.NewPar("\nPress Ctrl-C to quit")
	op.Height = th - ih
	op.BorderLabel = "Output"
	op.BorderLabelFg = termui.ColorCyan
	op.BorderFg = termui.ColorCyan
	op.TextFgColor = termui.ColorWhite

	// Now we need to create the layout. The panes have gotten a size
	// but no position. A grid layout puts everything into place.
	// termui.Body is a pre-defined grid. We add one row that contains
	// two columns.
	// The grid uses a 12-column system, so we have to give a "span"
	termui.Body.AddRows(
		termui.NewRow(
			termui.NewCol(3, 0, lp),
			termui.NewCol(9, 0, op, ip)))

	// Render the grid.
	termui.Body.Align()
	termui.Render(termui.Body)

	// When the window resizes, the grid must adopt to the new size.
	// We use a hander func for this.
	termui.Handle("/sys/wnd/resize", func(termui.Event) {
		termui.Body.Align()
		termui.Render(termui.Body)
	})
	// We need a way out. Ctrl-C shall stop the event loop.
	termui.Handle("/sys/kbd/C-c", func(termui.Event) {
		termui.StopLoop()
	})

	// start the event loop.
	termui.Loop()
}

/*
That wasn't too difficult, was it? Text entry is missing here, as `termui`
has no input controls yet. Still, implementing an input box is possible with
the available API methods; see the `_example` subdirectory in `termui`'s repository.

## gocui

Now let's see how `gocui` solves the same task.
*/

// Set up the widgets and run the event loop.
func runGocui() {
	// Create a new GUI.
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Println("Failed to create a GUI:", err)
		return
	}
	defer g.Close()

	// Activate the cursor for the current view.
	g.Cursor = true

	// The GUI object wants to know how to manage the layout.
	// Unlike `termui`, `gocui` does not use
	// a grid layout. Instead, it relies on a custom layout handler function
	// to manage the layout.

	// Here we set the layout manager to a function named `layout`
	// that is defined further down.
	g.SetManagerFunc(layout)

	// Bind the `quit` handler function (also defined further down) to Ctrl-C,
	// so that we can leave the application at any time.
	err = g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit)
	if err != nil {
		log.Println("Could not set key binding:", err)
		return
	}

	// Now let's define the views.

	// First, some dimensions.
	tw, th := g.Size()

	// First, the list view.
	lv, err := g.SetView("list", 0, 0, lw, th-1)
	if err != nil && err != gocui.ErrUnknownView {
		log.Println("Failed to create main view:", err)
		return
	}
	lv.Title = "List"
	lv.FgColor = gocui.ColorCyan

	// Then the output view.
	ov, err := g.SetView("output", lw+1, 0, tw-1, th-ih-1)
	if err != nil && err != gocui.ErrUnknownView {
		log.Println("Failed to create output view:", err)
		return
	}
	ov.Title = "Output"
	ov.FgColor = gocui.ColorGreen
	// Let the view scroll if the output exceeds the visible area.
	ov.Autoscroll = true

	// And finally the input view.
	iv, err := g.SetView("input", lw+1, th-ih, tw-1, th-1)
	if err != nil && err != gocui.ErrUnknownView {
		log.Println("Failed to create input view:", err)
		return
	}
	iv.Title = "Input"
	iv.FgColor = gocui.ColorYellow
	// The input view is editable.
	iv.Editable = true
	err = iv.SetCursor(0, 0)
	if err != nil {
		log.Println("Failed to set cursor:", err)
		return
	}

	// Make the enter key copy the input to the output.
	g.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, iv *gocui.View) error {
		iv.Rewind()

		// Get the output view via its name.
		ov, err := g.View("output")
		if err != nil {
			log.Println("Cannot get output view:", err)
			return err
		}
		fmt.Fprint(ov, iv.Buffer())
		iv.Clear()
		err = iv.SetCursor(0, 0)
		if err != nil {
			log.Println("Failed to set cursor:", err)
		}
		return err

	})

	// Fill the list view.
	for _, s := range listItems {
		_, err = fmt.Fprintln(lv, s)
		if err != nil {
			log.Println("Error writing to the list view:", err)
			return
		}
	}

	// Set the focus to the input view.
	g.SetCurrentView("input")

	// Start the main loop.
	err = g.MainLoop()
	log.Println("Main loop has finished:", err)
}

// The layout handler calculates all sizes depending
// on the current terminal size.
func layout(g *gocui.Gui) error {
	// Get the current terminal size.
	tw, th := g.Size()

	// Update the views according to the new terminal size.
	_, err := g.SetView("list", 0, 0, lw, th-1)
	if err != nil {
		return errors.Wrap(err, "Cannot update list view")
	}
	_, err = g.SetView("output", lw+1, 0, tw-1, th-ih-1)
	if err != nil {
		return errors.Wrap(err, "Cannot update output view")
	}
	_, err = g.SetView("input", lw+1, th-ih, tw-1, th-1)
	if err != nil {
		return errors.Wrap(err, "Cannot update input view.")
	}
	return nil
}

// `quit` is a handler that gets bound to Ctrl-C.
// It signals the main loop to exit.
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

/*
Our main func just needs to read the name from the TUI lib from the command line
and execute the respective code.
*/

//
func main() {
	if len(os.Args) <= 1 {
		log.Println("Usage: go run tui.go [termui|gocui]")
		return
	}
	if os.Args[1] == "termui" {
		runTermui()
		return
	}
	if os.Args[1] == "gocui" {
		runGocui()
		return
	}
	log.Println("No such option:", os.Args[0])
}

/*
## How to get and run the code

Step 1: `go get` the code. Note the `-d` flag that prevents auto-installing
the binary into `$GOPATH/bin`.

    go get -d github.com/appliedgo/tui

Step 2: `cd` to the source code directory.

    cd $GOPATH/src/github.com/appliedgo/tui

Step 3. Run the binary.

    go run tui.go termui
    go run tui.go gocui


## Odds and ends
## Some remarks
## Tips
## Links


**Happy coding!**

*/
