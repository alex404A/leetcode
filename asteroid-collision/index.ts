class Stack {
  list: number[]
  length: number

  constructor() {
    this.list = []
    this.length = 0
  }

  push(value: number) {
    this.list.push(value)
    this.length++
  }

  retrieve(): number | null {
    return this.length > 0 ? this.list[this.length - 1] : null
  }

  pop(): number {
    if (this.list.length > 0) {
      const result = this.list[this.list.length - 1]
      this.list = this.list.slice(0, this.list.length - 1)
      this.length--
      return result
    } else {
      throw new Error('try to pop when stack is empty')
    }
  }

  toList(): number[] {
    return this.list.map(a => a)
  }

}

function asteroidCollision(asteroids: number[]): number[] {
  const results: Stack = new Stack()
  for (let asteroid of asteroids) {
    if (asteroid >= 0) {
      results.push(asteroid)
    } else {
      let collision : number | null = asteroid
      while (results.length > 0 && collision !== null && collision < 0) {
        let before = results.retrieve()
        if (before === null || before < 0) {
          break
        }
        before = results.pop()
        collision = collide(before, asteroid)
      }
      if (collision !== null) {
        results.push(collision)
      }
    }
  }
  return results.toList()
}

function collide(a: number, b: number): number | null {
  if (a + b === 0) {
    return null
  }
  const size = Math.max(Math.abs(a), Math.abs(b))
  const isPositive = a < 0 ? Math.abs(a) < Math.abs(b) : Math.abs(a) > Math.abs(b)
  return isPositive ? size : 0 - size
}

function main(): void {
  const before = [5,10,-5, -2, -11, 4, 3, -3, 3, -2, -1]
  const after = asteroidCollision(before)
  console.log(after)
}

main()
