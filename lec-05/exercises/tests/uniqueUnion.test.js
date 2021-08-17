const uniqueUnion = require("../src/uniqueUnion")

test('test unique union', () => {
    expect(uniqueUnion([1, 3, 2], [5, 2, 1, 4], [2, 1])).toEqual([1, 3, 2, 5, 4])
})

test('test unique union', () => {
    expect(uniqueUnion([1, 2, 3], [5, 2, 1])).toEqual([1, 2, 3, 5])
})


test('test unique union', () => {
    expect(uniqueUnion([1, 2, 3], [5, 2, 1, 4], [2, 1], [6, 7, 8]) ).toEqual([1, 2, 3, 5, 4, 6, 7, 8])
})
