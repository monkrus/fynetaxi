package main

import "strings"

type note struct {
	content string
}

func (n *note) title() string {
	if n.content == "" {
		return "Untitled"
	}
	return strings.SplitN(n.content, "\n", 2)[0]
}

type notelist struct {
	notes []*note
}

func (l *notelist) add() *note {
	newNote := &note{}
	l.notes = append([]*note{newNote}, l.notes...)
	return newNote
}

//func (l *notelist) remove() *note {

//}
