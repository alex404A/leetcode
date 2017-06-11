function hehe() {
  let a = 2
  function lala() {
    console.log(a)
    a = 3
    console.log(a)
  }
  lala()
  console.log(a)
}

a = 1
hehe()
