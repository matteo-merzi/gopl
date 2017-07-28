package column

import (
	"sort"
	"testing"
)

func TestByArtist(t *testing.T) {
	var tracks = []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, Length("3m38s")},
		{"Go", "Moby", "Moby", 1992, Length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, Length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, Length("4m24s")},
	}
	c := NewByColumns(tracks)
	c.Select(ByArtist)
	sort.Sort(c)
	PrintTracks(tracks)
	Cmp(tracks, []*Track{
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, Length("4m36s")},
		{"Go", "Delilah", "From the Roots Up", 2012, Length("3m38s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, Length("4m24s")},
		{"Go", "Moby", "Moby", 1992, Length("3m37s")},
	}, t)
}

func TestByTitleAndByYear(t *testing.T) {
	var tracks = []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, Length("3m38s")},
		{"Go", "Moby", "Moby", 1992, Length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, Length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, Length("4m24s")},
	}
	c := NewByColumns(tracks)
	c.Select(ByTitle)
	c.Select(ByYear)
	sort.Sort(c)
	PrintTracks(tracks)
	Cmp(tracks, []*Track{
		{"Go", "Moby", "Moby", 1992, Length("3m37s")},
		{"Go", "Delilah", "From the Roots Up", 2012, Length("3m38s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, Length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, Length("4m24s")},
	}, t)
}

func TestByTitleAndByLength(t *testing.T) {
	var tracks = []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, Length("3m38s")},
		{"Go", "Moby", "Moby", 1992, Length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, Length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, Length("4m24s")},
	}
	c := NewByColumns(tracks)
	c.Select(ByTitle)
	c.Select(ByLength)
	sort.Sort(c)
	PrintTracks(tracks)
	Cmp(tracks, []*Track{
		{"Go", "Moby", "Moby", 1992, Length("3m37s")},
		{"Go", "Delilah", "From the Roots Up", 2012, Length("3m38s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, Length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, Length("4m24s")},
	}, t)
}

func Cmp(a, b []*Track, t *testing.T) {
	if len(a) != len(b) {
		t.Log("different lengths")
		t.Logf("%s\n%s", a, b)
		t.Fail()
		return
	}
	for i := 0; i < len(a); i++ {
		if a[i].Album != b[i].Album || a[i].Artist != b[i].Artist || a[i].Year != b[i].Year || a[i].Title != b[i].Title {
			t.Logf("different elements, starting at %d", i)
			t.Logf("%s\n%s", a, b)
			t.Fail()
			return
		}
	}
}
