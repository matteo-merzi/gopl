// ex3.4 serves SVG rendering of a 3-D surface function over http.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func svg(w io.Writer, peak string, valley string) {
	zmin, zmax := minmax()
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}
			fmt.Fprintf(w, "<polygon style='stroke: %s; fill: #222222' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color(i, j, zmin, zmax, peak, valley), ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("GET params were:", r.URL.Query())
		cs := [6]string{"red", "green", "blue", "yellow", "cyan", "violet"}
		p := checkColor(r.URL.Query().Get("peak"), cs, "red")
		v := checkColor(r.URL.Query().Get("valley"), cs, "blue")

		w.Header().Set("Content-Type", "image/svg+xml")
		svg(w, p, v)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func checkColor(color string, colors [6]string, defColor string) string {
	for _, c := range colors {
		if color == c {
			return color
		}
	}

	return defColor
}

func minmax() (min float64, max float64) {
	min = math.NaN()
	max = math.NaN()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			for xoff := 0; xoff <= 1; xoff++ {
				for yoff := 0; yoff <= 1; yoff++ {
					x := xyrange * (float64(i+xoff)/cells - 0.5)
					y := xyrange * (float64(j+yoff)/cells - 0.5)
					z := f(x, y)
					if math.IsNaN(min) || z < min {
						min = z
					}
					if math.IsNaN(max) || z > max {
						max = z
					}
				}
			}
		}
	}
	return
}

func color(i, j int, zmin, zmax float64, peak, valley string) string {
	min := math.NaN()
	max := math.NaN()
	for xoff := 0; xoff <= 1; xoff++ {
		for yoff := 0; yoff <= 1; yoff++ {
			x := xyrange * (float64(i+xoff)/cells - 0.5)
			y := xyrange * (float64(j+yoff)/cells - 0.5)
			z := f(x, y)
			if math.IsNaN(min) || z < min {
				min = z
			}
			if math.IsNaN(max) || z > max {
				max = z
			}
		}
	}

	color := ""
	if math.Abs(max) > math.Abs(min) {
		p := math.Exp(math.Abs(max)) / math.Exp(math.Abs(zmax)) * 255
		if p > 255 {
			p = 255
		}
		color = code(peak, int(p))
	} else {
		v := math.Exp(math.Abs(min)) / math.Exp(math.Abs(zmin)) * 255
		if v > 255 {
			v = 255
		}
		color = code(valley, int(v))
	}
	return color
}

func code(color string, op int) string {
	code := ""
	switch color {
	case "red":
		code = fmt.Sprintf("#%02x0000", op)
	case "green":
		code = fmt.Sprintf("#00%02x00", op)
	case "yellow":
		code = fmt.Sprintf("#%02x%02x00", op, op)
	case "cyan":
		code = fmt.Sprintf("#00%02x%02x", op, op)
	case "violet":
		code = fmt.Sprintf("#%02x00%02x", op, op)
	case "blue":
		code = fmt.Sprintf("#0000%02x", op)
	}
	return code
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
