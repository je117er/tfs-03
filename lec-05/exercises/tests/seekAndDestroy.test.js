const seekAndDestroy = require("../src/seekAndDestroy")

test('test seek and destroy', () => {
    expect(seekAndDestroy([1, 2, 3, 1, 2, 3], 2, 3)).toEqual([1, 1])
})

test('test seek and destroy', () => {
    expect(seekAndDestroy([1, 2, 3, 5, 1, 2, 3], 2, 3)).toEqual([1, 5, 1])
})

test('test seek and destroy', () => {
    expect(seekAndDestroy(["foo", "bar", 1], "foo", 1)).toEqual(["bar"])
})

