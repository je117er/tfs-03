// drops elements from the beginning of an array
// not satisfying a given condition until reaching the one that does
const drop = (arr, func) => {
    // returns the first index of an element
    // statisfying the function
    const index = arr.findIndex(func);
    return index === -1 ? [] : arr.slice(index);
};

module.exports = drop