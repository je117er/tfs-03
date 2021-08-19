// checks if a string is palindromic
const isPalindrome = (str) => {
    // removes all non-alphanumeric characters
    const newStr = str.replace(/[^a-zA-Z0-9]/g, '').toLowerCase();
    return newStr === newStr.split('').reverse().join('');
}

module.exports = isPalindrome

