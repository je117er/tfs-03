const isPalindrome = require("../src/isPalindrome")

test('test palindrome', () => {
    expect(isPalindrome("_eye")).toBe(true)
})

test('test palindrome', () => {
    expect(isPalindrome("race car")).toBe(true)
})

test('test palindrome', () => {
    expect(isPalindrome("_eye")).toBe(true)
})


test('test palindrome', () => {
    expect(isPalindrome("0_0 (: /-\\ :) 0-0")).toBe(true)
})

test('test palindrome', () => {
    expect(isPalindrome("five|\\_/|four")).toBe(false)
})

test('test palindrome', () => {
    expect(isPalindrome("A man, a plan, a canal. Panama")).toBe(true)
})

test('test palindrome', () => {
    expect(isPalindrome("My age is 0, 0 si ega ym.")).toBe(true)
})

test('test palindrome', () => {
    expect(isPalindrome("jodi is not palindromic")).toBe(false)
})