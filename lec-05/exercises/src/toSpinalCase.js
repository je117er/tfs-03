// converts a string to spinal case
// e.g., "USAToday" to "usa-today"
const toSpinalCase = (str) => {
    // puts a space among words
    return str.replace(/([a-z]|[A-Z]+)([A-Z])/g, "$1 $2")
        // replaces spaces and underscores with an en dash
        .replace(/\s+|_+/g, '-')
        // converts to lower case
        .toLowerCase();
}

module.exports = toSpinalCase