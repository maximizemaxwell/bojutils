~~미친 백준에서 go get을 지원하지 않아서 import를 못한다는 사실을 깨달아버림~~

# bojutils

- ~~백준에서 채점이 안되겠지만~~
- 만들때는 분명 고랭도 Cpp처럼 유사 STL이 존재한다면 많은 사람들이 고랭으로 백준을 풀지 않을까 라는 생각을 하고 만들었슴미다...

---

## Contents

1. [설치 방법](#설치-방법)
2. [제공 기능](#제공-기능)
   - [자료구조](#자료구조)
   - [알고리즘](#알고리즘)
3. [사용 방법](#사용-방법)
   - [Stack](#stack)
   - [Queue](#queue)
   - [Deque](#deque)
   - [DFS](#dfs)
   - [BFS](#bfs)

---

## **설치 방법**

```bash
go get github.com/maximizemaxwell/bojutils
```

- 위 명령어로 설치하고 개별적으로 임포트해서 사용할 수 있어요

```bash

import "github.com/maximizemaxwell/bojutils/datastructures/stack"
import "github.com/maximizemaxwell/bojutils/algorithms/bfs"
```

- 아니 너무 긴데 진지하게 잘못 만든 프로젝트 같아요

---

## 제공 기능

일단은 별 거 없고

### 자료구조

- 스택
    - Push(item T)
    - Pop() 
    - IsEmpty()

- 큐
    - Enqueue(item T)
    - Dequeue()
    - IsEmpty()

- 덱
    - PushFront(item T)
    - PushBack(item T)
    - PopFront()
    - PopBack()
    - IsEmpty()


### 알고리즘

- BFS

- DFS

---

## 사용 방법

### stack

```go
package main

import (
	"fmt"
	"github.com/maximizemaxwell/bojutils/datastructures/stack"
)

func main() {
	s := stack.NewStack[int]() // 정수형 스택 생성
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Pop())       // 출력: 2
	fmt.Println(s.IsEmpty())   // 출력: false
	fmt.Println(s.Pop())       // 출력: 1
	fmt.Println(s.IsEmpty())   // 출력: true
}
```

### queue

```go
package main

import (
	"fmt"
	"github.com/maximizemaxwell/bojutils/datastructures/queue"
)

func main() {
	q := queue.NewQueue[int]() // 정수형 큐 생성
	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println(q.Dequeue())   // 출력: 1
	fmt.Println(q.IsEmpty())   // 출력: false
	fmt.Println(q.Dequeue())   // 출력: 2
	fmt.Println(q.IsEmpty())   // 출력: true
}
```

### deque

```go
package main

import (
	"fmt"
	"github.com/maximizemaxwell/bojutils/datastructures/deque"
)

func main() {
	d := deque.NewDeque[int]() // 정수형 덱 생성
	d.PushBack(1)
	d.PushFront(2)
	fmt.Println(d.PopFront())  // 출력: 2
	d.PushBack(3)
	fmt.Println(d.PopBack())   // 출력: 3
	fmt.Println(d.PopBack())   // 출력: 1
	fmt.Println(d.IsEmpty())   // 출력: true
}
```

### DFS
~~쓸데없다~~

### BFS

[백준 미로탐색](https://www.acmicpc.net/problem/2178) 문제를 원래
```go

package main

import (
	"fmt"
	"github.com/maximizemaxwell/bojutils/algorithms/bfs"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	// 미로 입력 받기
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		var row string
		fmt.Scan(&row)
		grid[i] = make([]int, m)
		for j := 0; j < m; j++ {
			grid[i][j] = int(row[j] - '0')
		}
	}

	// BFS를 이용한 최단 거리 탐색
	start := [2]int{0, 0}
	end := [2]int{n - 1, m - 1}
	fmt.Println(bfs.BFS(grid, start, end))
}
```

이렇게 풀고 싶었으나...현실은 저렇게 임포트해도 백준에서 외부 라이브러리 못 써서 못함.


