package main

type MockPostInfo struct {
	author  string
	title   string
	content string
}

// ============================================
// Implements PostInfo interface
// ============================================

func (mpi *MockPostInfo) Author() string  { return mpi.author }
func (mpi *MockPostInfo) PostTitle() string   { return mpi.title }
func (mpi *MockPostInfo) Content() string { return mpi.content }
func (mpi *MockPostInfo) ExtraInfo() string { return "Yes some extra info" }

func NewMockPostInfo() *MockPostInfo {
	return &MockPostInfo{
		author: "John Doe",
		title:  "Introduction to Advanced Algorithms",
		content: `Welcome to the Advanced Algorithms course!

This semester, we will explore various algorithmic techniques and data structures that are fundamental to computer science. Here's what you can expect:

**Topics Covered:**
1. Dynamic Programming
   - Memoization techniques
   - Bottom-up approaches
   - Classic problems (Knapsack, Longest Common Subsequence, etc.)

2. Graph Algorithms
   - Breadth-First Search (BFS)
   - Depth-First Search (DFS)
   - Dijkstra's Algorithm
   - Minimum Spanning Trees (Prim's and Kruskal's)

3. Advanced Data Structures
   - Segment Trees
   - Fenwick Trees (Binary Indexed Trees)
   - Trie structures
   - Disjoint Set Union (Union-Find)

4. String Algorithms
   - KMP Pattern Matching
   - Rabin-Karp Algorithm
   - Z-Algorithm

**Course Requirements:**
- Weekly programming assignments
- Midterm exam (Week 8)
- Final project (Weeks 13-15)
- Active participation in class discussions

**Grading Breakdown:**
- Assignments: 40%
- Midterm: 25%
- Final Project: 30%
- Participation: 5%

**Office Hours:**
Monday and Wednesday, 2:00 PM - 4:00 PM
Room: CS Building, Office 301

**Important Dates:**
- First Assignment Due: February 1st
- Midterm Exam: March 15th
- Final Project Proposal: April 1st
- Final Project Presentation: May 10th

Please make sure to check the course website regularly for updates, lecture notes, and additional resources. If you have any questions, feel free to reach out during office hours or via email.

Looking forward to a great semester!

Best regards,
Professor John Doe
Department of Computer Science`,
	}
}
