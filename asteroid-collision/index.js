var Stack = /** @class */ (function () {
    function Stack() {
        this.list = [];
        this.length = 0;
    }
    Stack.prototype.push = function (value) {
        this.list.push(value);
        this.length++;
    };
    Stack.prototype.retrieve = function () {
        return this.length > 0 ? this.list[this.length - 1] : null;
    };
    Stack.prototype.pop = function () {
        if (this.list.length > 0) {
            var result = this.list[this.list.length - 1];
            this.list = this.list.slice(0, this.list.length - 1);
            this.length--;
            return result;
        }
        else {
            throw new Error('try to pop when stack is empty');
        }
    };
    Stack.prototype.toList = function () {
        return this.list.map(function (a) { return a; });
    };
    return Stack;
}());
function asteroidCollision(asteroids) {
    var results = new Stack();
    for (var _i = 0, asteroids_1 = asteroids; _i < asteroids_1.length; _i++) {
        var asteroid = asteroids_1[_i];
        if (asteroid >= 0) {
            results.push(asteroid);
        }
        else {
            var collision = asteroid;
            while (results.length > 0 && collision !== null && collision < 0) {
                var before = results.retrieve();
                if (before === null || before < 0) {
                    break;
                }
                before = results.pop();
                collision = collide(before, asteroid);
            }
            if (collision !== null) {
                results.push(collision);
            }
        }
    }
    return results.toList();
}
function collide(a, b) {
    if (a + b === 0) {
        return null;
    }
    var size = Math.max(Math.abs(a), Math.abs(b));
    var isPositive = a < 0 ? Math.abs(a) < Math.abs(b) : Math.abs(a) > Math.abs(b);
    return isPositive ? size : 0 - size;
}
function main() {
    var before = [5, 10, -5, -2, -11, 4, 3, -3, 3, -2, -1];
    var after = asteroidCollision(before);
    console.log(after);
}
main();
