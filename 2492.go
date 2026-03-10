func minScore(n int, roads [][]int) int {
    cities, absMin := buildGraph(roads, n)
    minLength := 2147483647
    paths := make([]*path, 0, n)
    paths = append(paths, &path{map[int]struct{}{1: struct{}{}}, 1, -1})

    for len(paths) > 0 {
        p := paths[0]
        paths = paths[1:]

        for _, r := range cities[p.currentCity].roads {
            if _, ok := p.visited[r.endCity]; !ok {
                paths = append(paths, &path{newVisited(p.visited, r.endCity), r.endCity, min(p.length, r.length)})
            } else if minLength > p.length {
                minLength = p.length
            }

            if minLength == absMin {
                return absMin
            }
        }
    }

    return minLength
}

func min(a, b int) int {
    if a != -1 && a < b {
        return a
    }
    return b
}

func newVisited(visited map[int]struct{}, newCity int) map[int]struct{} {
    nv := make(map[int]struct{}, len(visited)+1)
    for i := range visited {
        nv[i] = struct{}{}
    }

    nv[newCity] = struct{}{}
    return nv
}

type path struct {
    visited map[int]struct{}
    currentCity int
    length int
}

type city struct {
    roads []road
}

type road struct {
    endCity, length int
}

func buildGraph(roads [][]int, n int) (map[int]*city, int) {
    minLength := 2147483647
    cities := make(map[int]*city, n)
    for i := 1; i <= n; i++ {
        cities[i] = new(city)
    }
    for _, r := range roads {
        cities[r[0]].roads = append(cities[r[0]].roads, road{r[1], r[2]})
        cities[r[1]].roads = append(cities[r[1]].roads, road{r[0], r[2]})
        if r[2] < minLength {
            minLength = r[2]
        }
    }

    return cities, minLength
}