let display = value => {

    document.getElementById("result").value += value
};

const del = () => {
    let val = document.getElementById("result").value
    val = val.slice(0, -1);
    document.getElementById("result").value = val;
}
const equals = () => {
    let expression = document.getElementById("result").value;
    expression = expression.replace("+", "%2b");
    console.log(expression)

    fetch('http://localhost:8000/eval?exp=' + expression, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        },
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            console.log('Success:', data);
            console.log(Object.values(data)[0]);
            value = Object.values(data)[0]
            if (isNumber(value)) {
                document.getElementById('result').value = Object.values(data)[0];
            } else {
                document.getElementById('result').value = "Math ERROR";
            }
        })
        .catch((error) => {
            console.error('Error: ', error)
        })
};

function clr() {
    document.getElementById("result").value = ""
}
function isNumber(str) {
    return /^\-?[0-9]+(e[0-9]+)?(\.[0-9]+)?$/.test(str);
}

// async function eval(expression) {
//     let response = await fetch("http://localhost:8000/eval?exp=" + expression);
//     let responseText = await response.text();
//
//     document.getElementById('result').innerHTML = responseText;
// }
// (async() => {
//     await eval(expression);
// })();

