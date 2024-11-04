// Показать меню
function showMenu() {
    document.getElementById('start-container').style.display = 'none';
    document.getElementById('menu-container').style.display = 'flex';
}

// Показать форму ввода матрицы
function showMatrixInput() {
    const dataInput = document.getElementById('data-input').value;
    const calculationType = document.getElementById('calculation-type').value;

    if (dataInput === 'manual' && calculationType === 'polynomial') {
        document.getElementById('menu-container').style.display = 'none';
        document.getElementById('matrix-input-container').style.display = 'flex';

        // Устанавливаем значения по умолчанию для размера матрицы на 3x3
        document.getElementById('matrix-rows').value = "3";
        document.getElementById('matrix-columns').value = "3";

        // Создаем матрицу 3x3 при отображении формы
        createMatrixInputs();

        // Запускаем ввод коэффициентов сразу, если значение m установлено по умолчанию
        initializeCoefficientInput();
    }
}

// Создать поля ввода матрицы
function createMatrixInputs() {
    const rows = parseInt(document.getElementById('matrix-rows').value);
    const columns = parseInt(document.getElementById('matrix-columns').value);
    const matrixContainer = document.getElementById('matrix-container');
    matrixContainer.innerHTML = ''; // Очищаем контейнер

    // Устанавливаем количество столбцов в CSS-сетке для матрицы
    matrixContainer.style.gridTemplateColumns = `repeat(${columns}, 50px)`;

    // Создаем поля для каждой ячейки матрицы в зависимости от количества строк и столбцов
    for (let i = 0; i < rows * columns; i++) {
        const input = document.createElement('input');
        input.type = 'number';
        input.value = 0;
        input.min = 0;
        input.max = 10;
        matrixContainer.appendChild(input);
    }
}

// Инициализировать ввод коэффициентов
function initializeCoefficientInput() {
    const degree = parseInt(document.getElementById('polynomial-degree').value);
    const coefficientContainer = document.getElementById('coefficient-entry-container');
    const coefficientDisplayContainer = document.getElementById('coefficient-display-container');

    if (degree >= 1) {
        coefficientContainer.style.display = 'block';
        coefficientDisplayContainer.style.display = 'none';
        currentCoefficientIndex = 0;
        coefficients = [];
        updateCoefficientLabel();
    }
}

// Обработчик для ввода коэффициента
function handleCoefficientEnter(event) {
    if (event.key === 'Enter') {
        const input = document.getElementById('coefficient-input');
        const value = parseInt(input.value);

        if (!isNaN(value)) {
            coefficients.push(value);
            currentCoefficientIndex++;

            if (currentCoefficientIndex <= parseInt(document.getElementById('polynomial-degree').value)) {
                input.value = '';
                updateCoefficientLabel();
            } else {
                displayCoefficients();
            }
        }
    }
}

// Обновить текст метки для коэффициента
function updateCoefficientLabel() {
    const label = document.getElementById('coefficient-label');
    label.textContent = `Введите a${currentCoefficientIndex}:`;
}

// Показать введенные коэффициенты
function displayCoefficients() {
    const coefficientDisplayContainer = document.getElementById('coefficient-display-container');
    const coefficientContainer = document.getElementById('coefficient-entry-container');

    coefficientContainer.style.display = 'none';
    coefficientDisplayContainer.style.display = 'block';
    coefficientDisplayContainer.textContent = `Коэффициенты: [${coefficients.join(', ')}]`;
}

// Сборка JSON и отправка на сервер
function submitPolynomialData() {
    const rows = parseInt(document.getElementById('matrix-rows').value);
    const columns = parseInt(document.getElementById('matrix-columns').value);
    const matrix = [];
    const degree = parseInt(document.getElementById('polynomial-degree').value);

    // Сборка значений матрицы
    const matrixInputs = document.querySelectorAll('#matrix-container input');
    for (let input of matrixInputs) {
        matrix.push(parseInt(input.value));
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

// Инициализация переменных
let coefficients = [];
let currentCoefficientIndex = 0;
