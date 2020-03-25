package main

import "io"

type SpeakWriter struct {
	w io.Writer
}
