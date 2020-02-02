var longestConsecutive = function(nums) {
    if (!nums || nums.length === 0) {
        return 0
    }
    var cache = new Cache()
    for (var i = 0; i < nums.length; i++) {
        var num = nums[i]
        cache.add(num)
    }
    return cache.max
};

function Cache() {
    this.m = {}
    this.max = -1000000000
}

Cache.prototype.add = function(num) {
    if (this.m[num]) {
        return
    }
    var target = this.max
    var i
    this.m[num] = 1
    if (this.m[num - 1] && !this.m[num + 1]) {
        target = this.m[num - 1] + 1
        for (i = num; i > num - target ; i--) {
            this.m[i] = target
        } 
    } else if (!this.m[num - 1] && this.m[num + 1]) {
        target = this.m[num + 1] + 1
        for (i = num; i < num + target ; i++) {
            this.m[i] = target
        } 
    } else if (this.m[num - 1] && this.m[num + 1]) {
        target = this.m[num + 1] + this.m[num - 1] + 1
        var start = num - this.m[num - 1]
        var end = num + this.m[num + 1]
        for (i = start; i <= end; i++) {
            this.m[i] = target
        } 
    } else {
        target = 1
    }
    if (target > this.max) {
        this.max = target
    }
}

function test() {
    var nums = [100, 4, 200, 1, 3, 2]
    var nums = [100, 10, 200, -1, 3, 7, 201, 5, 101, 13, 11, 102]
    var nums = [1 ,3, 5, 2, 4]
    var result = longestConsecutive(nums)
    console.log(result)
}

test()