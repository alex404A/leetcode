def hehe():
    a = {1: 2}
    def lala(b):
        # print(a)
        b[1] = 3
    lala(a)
    print(a)

hehe()
