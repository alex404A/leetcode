package main

import (
	"fmt"
	"strings"
)

func isMatch(s string, p string) bool {
	fsm := parse(p)
	return check(s, fsm)
}

type State struct {
	name  string
	index int
	edges []*Edge
}

func (state *State) addEdge(target *Edge) {
	for _, edge := range state.edges {
		if edge.val == target.val && edge.next == target.next {
			return
		}
	}
	state.edges = append(state.edges, target)
}

func (state *State) findNext(val string) []*State {
	result := make([]*State, 0)
	for _, edge := range state.edges {
		if edge.val == val {
			result = append(result, edge.next)
		}
	}
	return result
}

type Edge struct {
	val  string
	next *State
}

type StateMachine struct {
	start *State
	m     map[int]*State
	end   *State
}

func (fsm *StateMachine) find(index int) (*State, bool) {
	state, ok := fsm.m[index]
	if ok {
		return state, true
	} else {
		return nil, false
	}
}

type StateSet struct {
	m map[int]*State
}

func (set *StateSet) add(state *State) {
	set.m[state.index] = state
}

func (set *StateSet) pop(state *State) {
	_, ok := set.m[state.index]
	if ok {
		delete(set.m, state.index)
	}
}

func (set *StateSet) toList() []*State {
	list := make([]*State, 0)
	for _, state := range set.m {
		list = append(list, state)
	}
	return list
}

func (set *StateSet) findLast() (*State, bool) {
	if len(set.m) == 0 {
		return nil, false
	}
	var maxState *State
	max := -2
	for _, state := range set.m {
		if state.index > max {
			max = state.index
			maxState = state
		}
	}
	return maxState, true
}

func parse(p string) *StateMachine {
	tokens := make([]string, 0)
	for i := 0; i < len(p); i++ {
		b := p[i]
		if i < len(p)-1 && p[i+1] == '*' {
			tokens = append(tokens, string([]byte{b, p[i+1]}))
			i++
			continue
		}
		tokens = append(tokens, string([]byte{b}))
	}
	cur := &State{"idle", -1, make([]*Edge, 0)}
	fsm := &StateMachine{cur, map[int]*State{-1: cur}, nil}
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		nextState := &State{token, i, make([]*Edge, 0)}
		fsm.m[i] = nextState
		if len(token) > 1 {
			bytes := strings.Split(token, "")
			b1 := bytes[0]
			nextState.addEdge(&Edge{b1, nextState})
			cur.addEdge(&Edge{b1, nextState})
		} else {
			cur.addEdge(&Edge{token, nextState})
		}
		backward(nextState, fsm)
		cur = nextState
	}
	fsm.end = cur
	return fsm
}

func check(s string, fsm *StateMachine) bool {
	set := &StateSet{map[int]*State{fsm.start.index: fsm.start}}
	for i := 0; i < len(s); i++ {
		c := string([]byte{s[i]})
		m := make(map[int]bool)
		states := make([]*State, 0)
		list := set.toList()
		for _, state := range list {
			nextStates1 := state.findNext(c)
			nextStates2 := state.findNext(".")
			states = append(states, state)
			if len(nextStates1) > 0 {
				for _, nextState := range nextStates1 {
					m[nextState.index] = true
					set.add(nextState)
				}
			}
			if len(nextStates2) > 0 {
				for _, nextState := range nextStates2 {
					m[nextState.index] = true
					set.add(nextState)
				}
			}
		}
		for _, state := range states {
			if _, ok := m[state.index]; !ok {
				set.pop(state)
			}
		}
		// delete(set.m, fsm.start.index)
		if len(set.m) == 0 {
			return false
		}
	}
	if _, ok := set.m[fsm.end.index]; ok {
		return true
	} else {
		end := fsm.end
		if len(end.name) == 1 {
			return false
		}
		edgeVal := strings.Split(end.name, "")[0]
		maxState, ok := set.findLast()
		if !ok {
			return false
		}
		cur := maxState
		for _, edge := range cur.edges {
			if edge.val == edgeVal && edge.next == end {
				return true
			}
		}
		return false
	}
}

func backward(nextState *State, fsm *StateMachine) {
	if nextState.index == 0 {
		return
	}
	var edge string
	if len(nextState.name) == 1 {
		edge = nextState.name
	} else {
		edge = strings.Split(nextState.name, "")[0]
	}
	for i := nextState.index - 1; i >= -1; i-- {
		previous, _ := fsm.m[i]
		previous.addEdge(&Edge{edge, nextState})
		if len(previous.name) == 1 {
			break
		}
	}
	return
}

func test(s string, p string, expected bool) {
	actual := isMatch(s, p)
	if expected != actual {
		fmt.Printf("%s %s\n", s, p)
	}
}

func main() {
	test("aaa", "a*a", true)
	test("mississippi", "mis*is*p*.", false)
	test("aab", "c*a*b", true)
	test("ab", ".*", true)
	test("aa", "a*", true)
	test("aa", "a", false)
	test("a", "ab*", true)
	test("aaa", "ab*a*c*a", true)
	test("a", ".*..a*", false)
	test("abcdede", "ab.*de", true)
	test("aasdfasdfasdfasdfas", "aasdf.*asdf.*asdf.*asdf.*s", true)
	test("abbaaaabaabbcba", "a*.*ba.*c*..a*.a*.", true)
}
