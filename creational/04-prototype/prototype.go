package main

import "time"

type Replicable interface {
	Clone() Replicable
}

type Homework struct {
	Subject  string
	Deadline time.Time
	Content  string
}

func (h *Homework) Clone() Replicable {
	return &Homework{
		Subject:  h.Subject,
		Deadline: h.Deadline.Add(time.Hour * 24),
		Content:  h.Content,
	}
}

type Dream struct {
	Content string
}

func (d *Dream) Clone() Replicable {
	return &Dream{Content: "You should find your own dream, this is mine, not yours"}
}
