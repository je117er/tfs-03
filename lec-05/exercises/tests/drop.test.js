const drop = require('../src/drop')

test('test drop', () => {
    expect(drop([1, 2, 3, 4], function(n) {return n >= 3;})).toEqual([3, 4])
})

test('test drop', () => {
    expect(drop([0, 1, 0, 1], function(n) {return n === 1;})).toEqual([1, 0, 1])
})
test('test drop', () => {
    expect(drop([1, 2, 3, 4], function(n) {return n > 5;})).toEqual([])
})
