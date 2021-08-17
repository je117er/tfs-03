const toSpinalCase = require("../src/toSpinalCase")

test('test spinal case', () => {
    expect(toSpinalCase("This Is Spinal Tap")).toBe('this-is-spinal-tap')
})

test('test spinal case', () => {
    expect(toSpinalCase("This Is Spinal Tap")).toBe('this-is-spinal-tap')
})
test('test spinal case', () => {
    expect(toSpinalCase("thisIsSpinalTap")).toBe('this-is-spinal-tap')
})
test('test spinal case', () => {
    expect(toSpinalCase("AllThe-small Things")).toBe('all-the-small-things')
})
test('test spinal case', () => {
    expect(toSpinalCase("This Is Spinal Tap")).toBe('this-is-spinal-tap')
})
test('test spinal case', () => {
    expect(toSpinalCase("I no longer read USAToday since renouncing my American citizenship")).toBe("i-no-longer-read-usa-today-since-renouncing-my-american-citizenship")
})

test('test spinal case', () => {
    expect(toSpinalCase("Simon say oh-no-no")).toBe("simon-say-oh-no-no")
})

test('test spinal case', () => {
    expect(toSpinalCase("I_wanna_escape")).toBe("i-wanna-escape")
})