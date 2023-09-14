package main

import (
    "fmt"
    "log"
    "strconv"
    "strings"
)



func getInput() string {
    return `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`;
}

type Point struct {
    x int
    y int
}

type Line struct {
    p1 Point
    p2 Point
}

func parsePoint(line string) (*Point, error) {
    parts := strings.Split(line, ",");
    x, err := strconv.Atoi(parts[0]);
    if err != nil {
        return nil, err
    }
    y, err := strconv.Atoi(parts[1]);
    if err != nil {
        return nil, err
    }
    
    return &Point{
        x: x,
        y: y,
    }, nil
}

func parseLine(line string) (*Line, error) {
    parts := strings.Split(line, " -> ");
    p1, err := parsePoint(parts[0]);
    if err != nil {
        return nil, err
    }
    p2, err := parsePoint(parts[1]);
    if err != nil {
        return nil, err
    }

    return &Line{
        p1: *p1,
        p2: *p2,
    }, nil
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func isOrthogonal(line Line) bool {
    return line.p1.x == line.p2.x || line.p1.y == line.p2.y   
}

func isHorizontal(line Line) bool {
    return line.p1.y == line.p2.y
}

func main() {
    lines := []Line{}
    for _, l := range strings.Split(getInput(), "\n") {
        line, err := parseLine(l)
        if err != nil {
            log.Fatal("Hey we couldn't parse the line")
        }
        if isOrthogonal(*line) {
            lines = append(lines, *line)
        }
    }

    matrix := [10][10]int{}

    for _, line := range lines {
        if isHorizontal(line) {
            y := line.p1.y
            startX := min(line.p1.x, line.p2.x)
            endX := max(line.p1.x, line.p2.x)
            for x := startX; x <= endX; x++ {
                matrix[y][x]++
            }
        } else {
            x := line.p1.x
            startY := min(line.p1.y, line.p2.y)
            endY := max(line.p1.y, line.p2.y)
            for y := startY; y <= endY; y++ {
                matrix[y][x]++
            }
        }
    }

    overlaps := 0
    for i := 0; i < 10; i++ {
        for j := 0; j < 10; j++ {
            if matrix[i][j] > 1 {
                overlaps++
            }
        }
    }

    fmt.Printf("Number of overlaps: %d\n", overlaps)
}
