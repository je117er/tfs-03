import './top.css';
import './style.css';
function component() {
    const element = document.createElement('div');
    element.innerHTML = 'I can do it!';
    return element;
}
document.body.appendChild(component());