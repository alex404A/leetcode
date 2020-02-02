function Node(val,neighbors) {
   this.val = val;
   this.neighbors = neighbors;
};

function Queue() {
    this.list = []
    this.length = 0
}

Queue.prototype.offer = function(node) {
    this.list.push(node)
    this.length++
}

Queue.prototype.pop = function(node) {
    var node = this.list.shift()
    if (node !== undefined) {
        this.length--
    }
    return node
}

function Pair(node) {
    this.before = node
    this.after = new Node(node.val, [])
}

function Cache() {
    this.m = {}
}

Cache.prototype.add = function(node) {
    var list = this.m[node.val]
    if (!list) {
        this.m[node.val] = []
    }
    this.m[node.val].push(new Pair(node))
}

Cache.prototype.check = function(node) {
    var list = this.m[node.val]
    if (!list) {
        return false
    }
    var i = 0
    var pair
    for (; i < list.length; i++) {
        pair = list[i] 
        if (pair.before === node) {
            return true
        }
    }
    return false
}

Cache.prototype.getPair = function(node) {
    var list = this.m[node.val]
    if (!list) {
        return null
    }
    var i = 0
    var pair
    for (; i < list.length; i++) {
        pair = list[i] 
        if (pair.before === node) {
            return pair
        }
    }
    return null
}

var cloneGraph = function(node) {
    if (node === null || node === undefined) {
        return node
    }
    var queue = new Queue()
    queue.offer(node)
    var cache = new Cache()
    cache.add(node)
    while (queue.length > 0) {
        var before = queue.pop()
        var i = 0
        var beforePair = cache.getPair(before)
        for (; i< before.neighbors.length; i++) {
            var neighbor = before.neighbors[i]
            if (!cache.check(neighbor)) {
                cache.add(neighbor)
                queue.offer(neighbor)
            }
            var neighborPair = cache.getPair(neighbor)
            beforePair.after.neighbors.push(neighborPair.after)
        }
    }
    var rootPair = cache.getPair(node)
    return rootPair.after
};

function test() {
    var n1 = new Node(1, [])
    var n2 = new Node(2, [])
    var n3 = new Node(3, [])
    var n4 = new Node(4, [])
    n1.neighbors.push(n2)
    n1.neighbors.push(n4)
    n2.neighbors.push(n1)
    n2.neighbors.push(n3)
    n3.neighbors.push(n2)
    n3.neighbors.push(n4)
    n4.neighbors.push(n3)
    n4.neighbors.push(n1)
    var node = cloneGraph(node)
    console.log(1)
}

test()