package main

import "fmt"

// for reference agains hellow_world.nb file
func main() {
	plotLineSimple(1, 0, 2, 3)
	fmt.Println("-------------")
	plotLine(10, -3, 11, -6)
	fmt.Println("-------------")
	plotLine(-5, 2, 0, 1)
	fmt.Println("-------------")
}

func plotLineSimple(x0, y0, x1, y1 int) {
	dx := x1 - x0
	dy := y1 - y0
	D := 2*dy - dx
	y := y0

	for x := x0; x < x1+1; x++ {
		fmt.Println(x, y)
		if D > 0 {
			y = y + 1
			D = D - 2*dx
		}
		D = D + 2*dy
	}
}

// plotLine(x0, y0, x1, y1)
//     dx = x1 - x0
//     dy = y1 - y0
//     D = 2*dy - dx
//     y = y0

//     for x from x0 to x1
//         plot(x,y)
//         if D > 0
//             y = y + 1
//             D = D - 2*dx
//         end if
//         D = D + 2*dy

// a go implementation of Bresenham's line algorithm
func plotLine(x0, y0, x1, y1 int) {
	dx := abs(x1 - x0)
	sx := -1
	if x0 < x1 {
		sx = 1
	}

	dy := -abs(y1 - y0)

	sy := -1
	if y0 < y1 {
		sy = 1
	}

	err := dx + dy // error value e_xy
	for {
		fmt.Println(x0, y0)
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 >= dy { // e_xy+e_x > 0
			err += dy
			x0 += sx
		}
		if e2 <= dx { // e_xy+e_y < 0
			err += dx
			y0 += sy
		}
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// TODO: remove this ref
// plotLine(int x0, int y0, int x1, int y1)
//     dx =  abs(x1-x0);
//     sx = x0<x1 ? 1 : -1;
//     dy = -abs(y1-y0);
//     sy = y0<y1 ? 1 : -1;
//     err = dx+dy;  /* error value e_xy */
//     while (true)   /* loop */
//         plot(x0, y0);
//         if (x0 == x1 && y0 == y1) break;
//         e2 = 2*err;
//         if (e2 >= dy) /* e_xy+e_x > 0 */
//             err += dy;
//             x0 += sx;
//         end if
//         if (e2 <= dx) /* e_xy+e_y < 0 */
//             err += dx;
//             y0 += sy;
//         end if
//     end while
