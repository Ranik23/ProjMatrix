// Инициализация переменных для коэффициентов
let coefficients = [];
let currentCoefficientIndex = 0;

// Показать меню
function showMenu() {
    document.getElementById('start-container').style.display = 'none';
    document.getElementById('menu-container').style.display = 'flex';
}

// Показать форму ввода матрицы или линейной формы
function showMatrixInput() {
    const dataInput = document.getElementById('data-input').value;
    const calculationType = document.getElementById('calculation-type').value;

    if (dataInput === 'manual') {
        if (calculationType === 'polynomial') {
            setupPolynomialMode();
        } else if (calculationType === 'linear-form') {
            setupLinearFormMode();
        }
    }
}

// Настройки для полинома
function setupPolynomialMode() {
    document.getElementById('menu-container').style.display = 'none';
    document.getElementById('matrix-input-container').style.display = 'flex';
    document.getElementById('linear-form-next-button').style.display = 'none'; // Hide the linear form "Next" button

    // Устанавливаем значения по умолчанию для размера матрицы на 3x3
    document.getElementById('matrix-rows').value = "3";
    document.getElementById('matrix-columns').value = "3";

    // Создаем матрицу 3x3 при отображении формы
    createMatrixInputs();

    // Показываем ввод для степени и коэффициентов
    document.getElementById('polynomial-degree-container').style.display = 'block';
    document.getElementById('coefficient-entry-container').style.display = 'block';
    document.getElementById('coefficient-display-container').style.display = 'none';

    // Запускаем ввод коэффициентов
    initializeCoefficientInput();
}

// Настройки для линейной формы
function setupLinearFormMode() {
    document.getElementById('menu-container').style.display = 'none';
    document.getElementById('matrix-input-container').style.display = 'none';
    document.getElementById('linear-form-next-button').style.display = 'block';
    document.getElementById('matrix-count-selection').style.display = 'block';
    document.getElementById('linear-matrix-size-selection').style.display = 'block';
}

// Показать поля для выбора размера матриц и для ввода матриц
function setupMatrixSizeSelection() {
    document.getElementById('linear-matrix-size-selection').style.display = 'flex';
    createMatrixFields();
}

// Создать поля для ввода матриц в зависимости от количества и размера
function createMatrixFields() {
    const matrixCount = parseInt(document.getElementById('matrix-count').value);
    const rows = parseInt(document.getElementById('linear-matrix-rows').value);
    const columns = parseInt(document.getElementById('linear-matrix-columns').value);
    const container = document.getElementById('matrix-input-fields-container');
    container.innerHTML = ''; // Очищаем контейнер

    // Создаем поля ввода для каждой матрицы
    for (let i = 1; i <= matrixCount; i++) {
        const matrixLabel = document.createElement('h3');
        matrixLabel.textContent = `Заполните матрицу X${i}`;
        container.appendChild(matrixLabel);

        // Создаем контейнер для текущей матрицы
        const matrixContainer = document.createElement('div');
        matrixContainer.classList.add('matrix-border');
        matrixContainer.style.display = 'grid';
        matrixContainer.style.gridTemplateColumns = `repeat(${columns}, 50px)`;

        // Создаем отдельные ячейки ввода для каждой позиции в матрице
        for (let j = 0; j < rows * columns; j++) {
            const input = document.createElement('input');
            input.type = 'number';
            input.value = 0;
            input.min = 0;
            input.max = 10;

            input.addEventListener('focus', (event) => {
                if (event.target.value === "0") {
                    event.target.value = ''; // Убираем ноль при фокусе
                }
            });

            input.addEventListener('blur', (event) => {
                if (event.target.value === '') {
                    event.target.value = 0; // Если пусто, ставим ноль
                }
            });

            matrixContainer.appendChild(input);
        }

        container.appendChild(matrixContainer);
    }

    container.style.display = 'block';
}

// Proceed to the next step for the linear form mode
function proceedToNextStep() {
    // Implement any further actions needed based on the filled matrices
    console.log("Proceeding to the next step in linear form mode");
}

// Proceed to the next step for the linear form mode
function proceedToNextStep() {
    const matrixCount = document.getElementById('matrix-count').value;
    console.log(`Number of matrices selected: ${matrixCount}`);
    // Implement any further actions needed based on matrix count
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

        input.addEventListener('focus', (event) => {
            if (event.target.value === "0") {
                event.target.value = ''; // Убираем ноль при фокусе
            }
        });

        input.addEventListener('blur', (event) => {
            if (event.target.value === '') {
                event.target.value = 0; // Если пусто, ставим ноль
            }
        });

        matrixContainer.appendChild(input);
    }

    enableArrowNavigation();
}

// Включить навигацию по стрелкам
function enableArrowNavigation() {
    const matrixInputs = document.querySelectorAll('#matrix-container input');
    const columns = parseInt(document.getElementById('matrix-columns').value);

    matrixInputs.forEach((input, index) => {
        input.addEventListener('keydown', (event) => {
            switch (event.key) {
                case 'ArrowRight':
                    if (index % columns < columns - 1) {
                        matrixInputs[index + 1].focus();
                    }
                    event.preventDefault();
                    break;
                case 'ArrowLeft':
                    if (index % columns > 0) {
                        matrixInputs[index - 1].focus();
                    }
                    event.preventDefault();
                    break;
                case 'ArrowDown':
                    if (index + columns < matrixInputs.length) {
                        matrixInputs[index + columns].focus();
                    }
                    event.preventDefault();
                    break;
                case 'ArrowUp':
                    if (index - columns >= 0) {
                        matrixInputs[index - columns].focus();
                    }
                    event.preventDefault();
                    break;
                default:
                    break;
            }
        });
    });
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

    // Заполняем коэффициенты до нужной степени значениями 1, если они не введены
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
