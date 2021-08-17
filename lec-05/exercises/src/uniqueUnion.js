// aggregates all arrays into one
// and removes duplicates
const uniqueUnion = (...arr) => {
    // concatenates elements from all arrays,
    // puts in a set, and transforms it back
    // to an array
    return [...new Set([].concat(...arr))];
}

module.exports = uniqueUnion
