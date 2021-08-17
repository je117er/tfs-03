// removes all provided elements from the array
const seekAndDestroy = (arr, ...args) => {
    return arr.filter(e => args.indexOf(e) === -1);
}

module.exports = seekAndDestroy