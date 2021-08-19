import dog from "./dog"

dog();

const obj = {
    foo: 'bar',
}

const extendedObj = {
    ...obj,
    name: 'j',
}
console.log(extendedObj);
console.log("es8 object.values", Object.values);