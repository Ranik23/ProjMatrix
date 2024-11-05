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

function initializeDefaultInputs() {
    const rowsInput = document.getElementById('matrix-rows-generate');
    const columnsInput = document.getElementById('matrix-columns-generate');
    const degreeInput = document.getElementById('polynomial-degree-generate');

    // Устанавливаем значения по умолчанию при загрузке
    rowsInput.value = 3;
    columnsInput.value = 3;
    degreeInput.value = 1;

    // Устанавливаем обработчики для автоматической замены значений
    [rowsInput, columnsInput].forEach(input => {
        const defaultValue = input.value;

        input.addEventListener('focus', (event) => {
            if (event.target.value === defaultValue) {
                event.target.value = '';
            }
        });

        input.addEventListener('blur', (event) => {
            if (event.target.value === '') {
                event.target.value = defaultValue;
            }
        });
    });

    // Обработчики для степени полинома
    degreeInput.addEventListener('focus', (event) => {
        if (event.target.value === "1") {
            event.target.value = '';
        }
    });

    degreeInput.addEventListener('blur', (event) => {
        if (event.target.value === '') {
            event.target.value = 1;
        }
    });
}

// Вызываем функцию при загрузке страницы или переключении на страницу генерации
initializeDefaultInputs();

function initializeLinearFormGenerationInputs() {
    const matrixCountInput = document.getElementById('matrix-count-generate');
    const rowsInput = document.getElementById('matrix-rows-generate-linear');
    const columnsInput = document.getElementById('matrix-columns-generate-linear');

    // Устанавливаем значения по умолчанию при загрузке
    matrixCountInput.value = 1;
    rowsInput.value = 3;
    columnsInput.value = 3;

    // Устанавливаем обработчики для автоматической замены значений
    [matrixCountInput, rowsInput, columnsInput].forEach(input => {
        const defaultValue = input.value;

        input.addEventListener('focus', (event) => {
            if (event.target.value === defaultValue) {
                event.target.value = '';
            }
        });

        input.addEventListener('blur', (event) => {
            if (event.target.value === '') {
                event.target.value = defaultValue;
            }
        });
    });
}

// Вызываем функцию при загрузке страницы или переключении на страницу генерации линейной формы
initializeLinearFormGenerationInputs();