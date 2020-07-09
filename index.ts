class TreeNode {
  val: number;
  left: TreeNode | null;
  right: TreeNode | null;
  constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
    this.val = val === undefined ? 0 : val;
    this.left = left === undefined ? null : left;
    this.right = right === undefined ? null : right;
  }
}

function deleteNode(root: TreeNode | null, key: number): TreeNode | null {
  if (root === null) {
    return null;
  }
  if (root.val === key) {
    return null;
  }
  let parent: TreeNode = root;
  let cur: TreeNode;
  let isFound: boolean = false;
  if (parent.val < key) {
    cur = parent.right;
  } else {
    cur = parent.left;
  }
  while (cur !== null) {
    if (cur.val == key) {
      isFound = true;
      break;
    } else if (cur.val < key) {
      parent = cur;
      cur = cur.right;
    } else {
      parent = cur;
      cur = cur.left;
    }
  }
  if (!isFound) {
    return root;
  } else {
    transform(cur, parent);
    return root;
  }
}

function transform(cur: TreeNode, parent: TreeNode) {
  const isLeft: boolean = parent.left === cur;
  let replaceNode: TreeNode;
  if (cur.left === null && cur.right === null) {
    return;
  } else if (cur.left !== null && cur.right === null) {
    replaceNode = cur.left
  } else if (cur.left === null && cur.right !== null) {
    replaceNode = cur.right
  } else {
    if (cur.left.right === null) {
      replaceNode = cur.left
    } else {
      let rightParent: TreeNode = cur.left
      let right: TreeNode = cur.left.right
      while (right.right !== null) {
        rightParent = right
        right = right.right
        rightParent.left = 
      }
    }
  }
  if (isLeft) {
    parent.left = replaceNode;
  } else {
    parent.right = replaceNode;
  }
}
