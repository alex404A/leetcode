package main

import "fmt"

type Queue struct {
	length int
	list   []string
}

func (queue *Queue) offer(s string) {
	queue.list = append(queue.list, s)
	queue.length++
}

func (queue *Queue) poll() (s string, ok bool) {
	if queue.length > 0 {
		s = queue.list[0]
		ok = true
		queue.list = queue.list[1:]
		queue.length--
		return
	} else {
		s = ""
		ok = false
		return
	}
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordList = append(wordList, beginWord)
	relation := build(wordList)
	neighbours, _ := relation[beginWord]
	if len(neighbours) == 0 {
		return make([][]string, 0)
	}
	visited, layer, isFound := bfs(beginWord, endWord, relation)
	if !isFound {
		return make([][]string, 0)
	}
	solution := make([]string, 1)
	solution[0] = beginWord
	return dfs(beginWord, endWord, relation, visited, layer, solution)
}

func bfs(beginWord string, endWord string, relation map[string][]string) (map[string]int, int, bool) {
	visited := make(map[string]int)
	queue := Queue{0, make([]string, 0)}
	queue.offer(beginWord)
	layer := 1
	visited[beginWord] = layer
	isFound := false
	for queue.length > 0 {
		layer++
		length := queue.length
		for i := 0; i < length; i++ {
			word, _ := queue.poll()
			neighbours := relation[word]
			for j := 0; j < len(neighbours); j++ {
				next := neighbours[j]
				if _, ok := visited[next]; !ok {
					visited[next] = layer
					if next == endWord {
						isFound = true
					} else {
						queue.offer(next)
					}
				}
			}
		}
		if isFound {
			break
		}
	}
	return visited, layer, isFound
}

func copy(solution []string) []string {
	other := make([]string, len(solution))
	for i := 0; i < len(solution); i++ {
		other[i] = solution[i]
	}
	return other
}

func dfs(beginWord string, endWord string, relation map[string][]string, visited map[string]int, layer int, solution []string) [][]string {
	if beginWord == endWord {
		res := make([][]string, 1)
		res[0] = solution
		return res
	}
	curLayer, _ := visited[beginWord]
	if layer <= curLayer {
		return make([][]string, 0)
	}
	res := make([][]string, 0)
	neighbours := relation[beginWord]
	for _, next := range neighbours {
		nextLayer, ok := visited[next]
		if ok && nextLayer == curLayer+1 {
			other := copy(solution)
			other = append(other, next)
			x := dfs(next, endWord, relation, visited, layer, other)
			if len(x) > 0 {
				for _, candidate := range x {
					res = append(res, candidate)
				}
			}
		}
	}
	return res
}

func build(wordList []string) map[string][]string {
	m := make(map[string][]string)
	for i := 0; i < len(wordList); i++ {
		candidate := wordList[i]
		m[candidate] = make([]string, 0)
		for j := 0; j < len(wordList); j++ {
			if i == j {
				continue
			} else {
				if isAdjacent(candidate, wordList[j]) {
					m[candidate] = append(m[candidate], wordList[j])
				}
			}
		}
	}
	return m
}

func isAdjacent(w1 string, w2 string) bool {
	cnt := 0
	for i := 0; i < len(w1); i++ {
		if w1[i] != w2[i] {
			cnt++
		}
		if cnt > 1 {
			return false
		}
	}
	return true
}

func main() {
	var beginWord string
	var endWord string
	var wordList []string
	var result [][]string
	beginWord = "hit"
	endWord = "cog"
	wordList = []string{"hot", "dot", "dog", "lot", "log", "cog"}
	result = findLadders(beginWord, endWord, wordList)
	fmt.Println(result)
	beginWord = "hit"
	endWord = "cog"
	wordList = []string{"hot", "dot", "dog", "lot", "log"}
	result = findLadders(beginWord, endWord, wordList)
	fmt.Println(result)
	beginWord = "qa"
	endWord = "sq"
	wordList = []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"}
	result = findLadders(beginWord, endWord, wordList)
	fmt.Println(result)
	beginWord = "cet"
	endWord = "ism"
	wordList = []string{"kid", "tag", "pup", "ail", "tun", "woo", "erg", "luz", "brr", "gay", "sip", "kay", "per", "val", "mes", "ohs", "now", "boa", "cet", "pal", "bar", "die", "war", "hay", "eco", "pub", "lob", "rue", "fry", "lit", "rex", "jan", "cot", "bid", "ali", "pay", "col", "gum", "ger", "row", "won", "dan", "rum", "fad", "tut", "sag", "yip", "sui", "ark", "has", "zip", "fez", "own", "ump", "dis", "ads", "max", "jaw", "out", "btu", "ana", "gap", "cry", "led", "abe", "box", "ore", "pig", "fie", "toy", "fat", "cal", "lie", "noh", "sew", "ono", "tam", "flu", "mgm", "ply", "awe", "pry", "tit", "tie", "yet", "too", "tax", "jim", "san", "pan", "map", "ski", "ova", "wed", "non", "wac", "nut", "why", "bye", "lye", "oct", "old", "fin", "feb", "chi", "sap", "owl", "log", "tod", "dot", "bow", "fob", "for", "joe", "ivy", "fan", "age", "fax", "hip", "jib", "mel", "hus", "sob", "ifs", "tab", "ara", "dab", "jag", "jar", "arm", "lot", "tom", "sax", "tex", "yum", "pei", "wen", "wry", "ire", "irk", "far", "mew", "wit", "doe", "gas", "rte", "ian", "pot", "ask", "wag", "hag", "amy", "nag", "ron", "soy", "gin", "don", "tug", "fay", "vic", "boo", "nam", "ave", "buy", "sop", "but", "orb", "fen", "paw", "his", "sub", "bob", "yea", "oft", "inn", "rod", "yam", "pew", "web", "hod", "hun", "gyp", "wei", "wis", "rob", "gad", "pie", "mon", "dog", "bib", "rub", "ere", "dig", "era", "cat", "fox", "bee", "mod", "day", "apr", "vie", "nev", "jam", "pam", "new", "aye", "ani", "and", "ibm", "yap", "can", "pyx", "tar", "kin", "fog", "hum", "pip", "cup", "dye", "lyx", "jog", "nun", "par", "wan", "fey", "bus", "oak", "bad", "ats", "set", "qom", "vat", "eat", "pus", "rev", "axe", "ion", "six", "ila", "lao", "mom", "mas", "pro", "few", "opt", "poe", "art", "ash", "oar", "cap", "lop", "may", "shy", "rid", "bat", "sum", "rim", "fee", "bmw", "sky", "maj", "hue", "thy", "ava", "rap", "den", "fla", "auk", "cox", "ibo", "hey", "saw", "vim", "sec", "ltd", "you", "its", "tat", "dew", "eva", "tog", "ram", "let", "see", "zit", "maw", "nix", "ate", "gig", "rep", "owe", "ind", "hog", "eve", "sam", "zoo", "any", "dow", "cod", "bed", "vet", "ham", "sis", "hex", "via", "fir", "nod", "mao", "aug", "mum", "hoe", "bah", "hal", "keg", "hew", "zed", "tow", "gog", "ass", "dem", "who", "bet", "gos", "son", "ear", "spy", "kit", "boy", "due", "sen", "oaf", "mix", "hep", "fur", "ada", "bin", "nil", "mia", "ewe", "hit", "fix", "sad", "rib", "eye", "hop", "haw", "wax", "mid", "tad", "ken", "wad", "rye", "pap", "bog", "gut", "ito", "woe", "our", "ado", "sin", "mad", "ray", "hon", "roy", "dip", "hen", "iva", "lug", "asp", "hui", "yak", "bay", "poi", "yep", "bun", "try", "lad", "elm", "nat", "wyo", "gym", "dug", "toe", "dee", "wig", "sly", "rip", "geo", "cog", "pas", "zen", "odd", "nan", "lay", "pod", "fit", "hem", "joy", "bum", "rio", "yon", "dec", "leg", "put", "sue", "dim", "pet", "yaw", "nub", "bit", "bur", "sid", "sun", "oil", "red", "doc", "moe", "caw", "eel", "dix", "cub", "end", "gem", "off", "yew", "hug", "pop", "tub", "sgt", "lid", "pun", "ton", "sol", "din", "yup", "jab", "pea", "bug", "gag", "mil", "jig", "hub", "low", "did", "tin", "get", "gte", "sox", "lei", "mig", "fig", "lon", "use", "ban", "flo", "nov", "jut", "bag", "mir", "sty", "lap", "two", "ins", "con", "ant", "net", "tux", "ode", "stu", "mug", "cad", "nap", "gun", "fop", "tot", "sow", "sal", "sic", "ted", "wot", "del", "imp", "cob", "way", "ann", "tan", "mci", "job", "wet", "ism", "err", "him", "all", "pad", "hah", "hie", "aim", "ike", "jed", "ego", "mac", "baa", "min", "com", "ill", "was", "cab", "ago", "ina", "big", "ilk", "gal", "tap", "duh", "ola", "ran", "lab", "top", "gob", "hot", "ora", "tia", "kip", "han", "met", "hut", "she", "sac", "fed", "goo", "tee", "ell", "not", "act", "gil", "rut", "ala", "ape", "rig", "cid", "god", "duo", "lin", "aid", "gel", "awl", "lag", "elf", "liz", "ref", "aha", "fib", "oho", "tho", "her", "nor", "ace", "adz", "fun", "ned", "coo", "win", "tao", "coy", "van", "man", "pit", "guy", "foe", "hid", "mai", "sup", "jay", "hob", "mow", "jot", "are", "pol", "arc", "lax", "aft", "alb", "len", "air", "pug", "pox", "vow", "got", "meg", "zoe", "amp", "ale", "bud", "gee", "pin", "dun", "pat", "ten", "mob"}
	result = findLadders(beginWord, endWord, wordList)
	fmt.Println(result)
}
