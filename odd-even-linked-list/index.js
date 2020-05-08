var oddEvenList = function(head) {
  if (head === null || head.next === null || head.next.next === null) {
    return head
  }
  var source = head
  var head = head.next
  var end = head
  var previous = head.next
  source.next = previous
  head.next = previous.next
  previous.next = head
  var result = {
    head,
    end,
    previous
  }
  while (true) {
    result = exchange(result.previous, result.head, result.end)
    if (result.isEnd) {
      break
    }
    console.log(`previous: ${result.previous.value}, head: ${result.head.value}, end: ${result.end.value}`)
  }
  return source
}

/**
 *
 * @return head, tail, isEnd
**/
function exchange(previous, head, end) {
  var result = {
    head,
    end: null,
    previous: null,
    isEnd: true
  }
  if (end.next === null || end.next.next === null) {
    return result
  } else {
    result.isEnd = false
    result.end = end.next
    previous.next = result.end.next 
    result.end.next = result.end.next.next
    previous.next.next = result.head
    result.previous = previous.next
    return result
  }
}

class LinkNode {
  constructor(value) {
    this.value = value
    this.next = null
  }
}

function build(arr) {
  var previous = new LinkNode(arr[0])
  var head = previous
  for (var i = 1; i < arr.length; i++) {
    var node = new LinkNode(arr[i])
    previous.next = node
    previous = node
  }
  return head
}

function print(head) {
  var values = []
  while (head != null) {
    values.push(head.value)
    head = head.next
  }
  console.log(values)
}

function main() {
  var head = build([1,2,3,4,5,6,7,8,9,10,11,12,13])
  print(head)
  head = oddEvenList(head)
  print(head)
}

main()
