def hehe():
    a = 2
    def lala():
        global a
        # print(a)
        a = 3
        print(a)
    lala()
    print(a)

hehe()
