package column

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func (t Track) String() string {
	return fmt.Sprintf(`Title: %s
	Artist: %s
	Album: %s
	Year: %d
	Length: %v\n`, t.Title, t.Artist, t.Album, t.Year, t.Length)
}

type comparison int

const (
	lt comparison = iota
	eq
	gt
)

type columnCmp func(a, b *Track) comparison

type ByColumns struct {
	t    []*Track
	cmps []columnCmp
}

func NewByColumns(t []*Track) *ByColumns {
	return &ByColumns{t, nil}
}

func ByTitle(a, b *Track) comparison {
	switch {
	case a.Title == b.Title:
		return eq
	case a.Title < b.Title:
		return lt
	default:
		return gt
	}
}

func ByArtist(a, b *Track) comparison {
	switch {
	case a.Artist == b.Artist:
		return eq
	case a.Artist < b.Artist:
		return lt
	default:
		return gt
	}
}

func ByAlbum(a, b *Track) comparison {
	switch {
	case a.Album == b.Album:
		return eq
	case a.Album < b.Album:
		return lt
	default:
		return gt
	}
}

func ByYear(a, b *Track) comparison {
	switch {
	case a.Year == b.Year:
		return eq
	case a.Year < b.Year:
		return lt
	default:
		return gt
	}
}

func ByLength(a, b *Track) comparison {
	switch {
	case a.Length == b.Length:
		return eq
	case a.Length < b.Length:
		return lt
	default:
		return gt
	}
}

func Length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func (c *ByColumns) Len() int      { return len(c.t) }
func (c *ByColumns) Swap(i, j int) { c.t[i], c.t[j] = c.t[j], c.t[i] }

func (c *ByColumns) Less(i, j int) bool {
	for _, f := range c.cmps {
		cmp := f(c.t[i], c.t[j])
		switch cmp {
		case eq:
			continue
		case lt:
			return true
		case gt:
			return false
		}
	}
	return false
}

func (c *ByColumns) Select(cmp columnCmp) {
	c.cmps = append(c.cmps, cmp)
}

func PrintTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}
