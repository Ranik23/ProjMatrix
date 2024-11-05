// utils.js

// Сборка JSON и отправка на сервер
function submitPolynomialData() {
    const rows = parseInt(document.getElementById('matrix-rows').value);
    const columns = parseInt(document.getElementById('matrix-columns').value);
    const matrix = [];
    const degree = parseInt(document.getElementById('polynomial-degree').value);

    const matrixInputs = document.querySelectorAll('#matrix-container input');
    for (let input of matrixInputs) {
        matrix.push(parseInt(input.value));
    }

    for (let i = coefficients.length; i <= degree; i++) {
        coefficients.push(1);
    }

    const data = {
        matrixSize: { rows, columns },
        matrix: matrix,
        degree: degree,
        coefficients: coefficients,
    };

    console.log(JSON.stringify(data)); // Вывод JSON в консоль для проверки

    // Здесь можно отправить JSON на сервер
}
