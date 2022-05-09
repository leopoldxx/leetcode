package main

type queue struct {
	data []int
}

func (q *queue) push(v int) {
	q.data = append(q.data, v)
}
func (q *queue) pop() *int {
	if len(q.data) == 0 {
		return nil
	}
	first := q.data[0]
	q.data = q.data[1:]
	return &first
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	preq := prerequisites
	var (
		degree = make([]int, numCourses)
		edges  = make([][]int, numCourses)
	)
	for i := 0; i < len(preq); i++ {
		degree[preq[i][0]]++
		edges[preq[i][1]] = append(edges[preq[i][1]], preq[i][0])
	}

	que := &queue{}
	for i := 0; i < numCourses; i++ {
		if degree[i] == 0 {
			que.push(i)
		}
	}
	count := 0
	for v := que.pop(); v != nil; v = que.pop() {
		count++
		for j := 0; j < len(edges[*v]); j++ {
			degree[edges[*v][j]]--
			if degree[edges[*v][j]] == 0 {
				que.push(edges[*v][j])
			}
		}
	}
	return count == numCourses

}
