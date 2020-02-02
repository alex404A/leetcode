function Node(val,left,right,next) {
   this.val = val
   this.left = left
   this.right = right
   this.next = next
}

function WrapperNode(node, layer) {
    this.node = node
    this.layer = layer
}

function Queue() {
    this.list = []
}

Queue.prototype.offer = function(wrapperNode) {
    this.list.push(wrapperNode)
}

Queue.prototype.remove = function() {
    if (this.list.length === 0) {
        throw new Error('No more node in queue')
    }
    var wrapperNode = this.list[0]
    this.list = this.list.slice(1)
    return wrapperNode
}

Queue.prototype.peek = function() {
    if (this.list.length === 0) {
        return null
    }
    return this.list[0]
}

/**
 * @param {Node} root
 * @return {Node}
 */
var connect = function(root) {
    if (!root) {
        return root
    }
    var wrapper = new WrapperNode(root, 1)
    var next, left, right
    var queue = new Queue()
    queue.offer(wrapper)
    while (queue.list.length > 0) {
        wrapper = queue.remove()
        next = queue.peek()
        if (next && next.layer === wrapper.layer) {
            wrapper.node.next = next.node
        }
        if (wrapper.node.left) {
            left = new WrapperNode(wrapper.node.left, wrapper.layer + 1) 
            right = new WrapperNode(wrapper.node.right, wrapper.layer + 1) 
            queue.offer(left)
            queue.offer(right)
        }
    }
    return root
}

function build(nums) {
    var nodes = []
    var index
    var node = new Node(nums[0], null, null, null)
    nodes.push(node)
    for (var i = 1; i < nums.length; i++) {
        if (nums[i] === null) {
            continue
        }
        node = new Node(nums[i], null, null, null)
        nodes.push(node)
        var index
        if (i % 2 == 1) {
            index = Math.floor(i / 2)
            nodes[index].left = node
        } else {
            index = Math.floor(i / 2) - 1
            nodes[index].right = node
        }
    }
    return nodes[0]
}

function main() {
    var nums = [1,2,3,4,5,6,7]
    var root = build(nums)
    connect(root)
    console.log(1)
}

main()