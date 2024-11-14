let linearFormCoefficients = [];
let currentLinearCoefficientIndex = 0;

// Функция для инициализации ввода коэффициентов при изменении количества матриц
function resetLinearFormCoefficientInput() {
    const matrixCount = parseInt(document.getElementById('matrix-count').value);

    // Сбросить массив коэффициентов и индекс
    linearFormCoefficients = [];
    currentLinearCoefficientIndex = 0;

    // Обновляем интерфейс: скрываем отображение старых коэффициентов, очищаем ввод
    const coefficientContainer = document.getElementById('linear-coefficient-entry-container');
    const coefficientDisplayContainer = document.getElementById('linear-coefficient-display-container');
    const coefficientInput = document.getElementById('linear-coefficient-input');

    // Скрываем старое отображение и показываем контейнер для ввода
    coefficientContainer.style.display = 'block';
    coefficientDisplayContainer.style.display = 'none';
    coefficientDisplayContainer.textContent = ''; // Очищаем старое отображение
    coefficientInput.value = ''; // Очищаем текущее значение ввода

    // Запустить ввод заново
    updateLinearCoefficientLabel();
}

function initializeLinearFormCoefficientInput() {
    const matrixCount = parseInt(document.getElementById('matrix-count').value);

    if (matrixCount >= 1) {
        resetLinearFormCoefficientInput();
    }
}

// Обработчик ввода коэффициента с помощью Enter
function handleLinearCoefficientEnter(event) {
    if (event.key === 'Enter') {
        const input = document.getElementById('linear-coefficient-input');
        const value = parseInt(input.value);

        if (!isNaN(value)) {
            linearFormCoefficients.push(value);
            currentLinearCoefficientIndex++;

            const matrixCount = parseInt(document.getElementById('matrix-count').value);
            if (currentLinearCoefficientIndex < matrixCount) {
                input.value = '';
                updateLinearCoefficientLabel();
            } else {
                displayLinearFormCoefficients();
            }
        }
    }
}

// Обновить метку для ввода коэффициентов
function updateLinearCoefficientLabel() {
    const label = document.getElementById('linear-coefficient-label');
    label.textContent = `Введите a${currentLinearCoefficientIndex}:`;
}

// Показать введённые коэффициенты
function displayLinearFormCoefficients() {
    const coefficientDisplayContainer = document.getElementById('linear-coefficient-display-container');
    const coefficientContainer = document.getElementById('linear-coefficient-entry-container');

    coefficientContainer.style.display = 'none';
    coefficientDisplayContainer.style.display = 'block';
    coefficientDisplayContainer.textContent = `Коэффициенты: [${linearFormCoefficients.join(', ')}]`;
}

// Функция для перехода на страницу результатов
function proceedToNextStep() {
    // Вся необходимая логика для сбора данных может быть здесь
    // После завершения переход на страницу с результатами
    window.location.href = "/results";
}

// Экспортируем функции для глобального доступа
window.initializeLinearFormCoefficientInput = initializeLinearFormCoefficientInput;
window.handleLinearCoefficientEnter = handleLinearCoefficientEnter;
window.proceedToNextStep = proceedToNextStep;

function createMatrixFields() {
    const matrixCount = parseInt(document.getElementById('matrix-count').value);
    const rows = parseInt(document.getElementById('linear-matrix-rows').value);
    const columns = parseInt(document.getElementById('linear-matrix-columns').value);
    const container = document.getElementById('matrix-input-fields-container');
    container.innerHTML = ''; // Очищаем контейнер

    for (let i = 1; i <= matrixCount; i++) {
        const matrixLabel = document.createElement('h3');
        matrixLabel.classList.add('matrix-title'); // Применяем класс
        matrixLabel.textContent = `Заполните матрицу X${i}`;
        container.appendChild(matrixLabel);

        const matrixContainer = document.createElement('div');
        matrixContainer.classList.add('matrix-border');
        matrixContainer.style.gridTemplateColumns = `repeat(${columns}, 50px)`; // Ширина input

        for (let j = 0; j < rows * columns; j++) {
            const input = document.createElement('input');
            input.type = 'number';
            input.value = 0;
            input.min = 0;
            input.max = 10;

            input.addEventListener('focus', (event) => {
                if (event.target.value === "0") {
                    event.target.value = '';
                }
            });

            input.addEventListener('blur', (event) => {
                if (event.target.value === '') {
                    event.target.value = 0;
                }
            });

            matrixContainer.appendChild(input);
        }

        // Включаем навигацию по стрелкам для текущей матрицы
        enableArrowNavigation(matrixContainer, columns);

        container.appendChild(matrixContainer);
    }

    container.style.display = 'block';
}

// Экспортируем для глобального доступа
window.createMatrixFields = createMatrixFields;


// Создать поля ввода матрицы
function createMatrixInputs() {
    const rows = parseInt(document.getElementById('matrix-rows').value);
    const columns = parseInt(document.getElementById('matrix-columns').value);
    const matrixContainer = document.getElementById('matrix-container');
    matrixContainer.innerHTML = ''; // Очищаем контейнер

    matrixContainer.style.gridTemplateColumns = `repeat(${columns}, 50px)`; // Используем ширину для input

    for (let i = 0; i < rows * columns; i++) {
        const input = document.createElement('input');
        input.type = 'number';
        input.value = 0;
        input.min = 0;
        input.max = 10;

        input.addEventListener('focus', (event) => {
            if (event.target.value === "0") {
                event.target.value = '';
            }
        });

        input.addEventListener('blur', (event) => {
            if (event.target.value === '') {
                event.target.value = 0;
            }
        });

        matrixContainer.appendChild(input);
    }

    // Включаем навигацию по стрелкам для матрицы на странице полинома
    enableArrowNavigation(matrixContainer, columns);
}

// Экспортируем для глобального доступа
window.createMatrixInputs = createMatrixInputs;

// Включить навигацию по стрелкам
function enableArrowNavigation(container, columns) {
    const matrixInputs = container.querySelectorAll('input');

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

// matrix.js

function setupLinearFormMode() {
    document.getElementById('menu-container').style.display = 'none';
    document.getElementById('matrix-input-container').style.display = 'none';
    document.getElementById('linear-form-next-button').style.display = 'block';
    document.getElementById('matrix-count-selection').style.display = 'block';
    document.getElementById('linear-matrix-size-selection').style.display = 'block';

    // Устанавливаем значение по умолчанию для количества матриц и размера
    document.getElementById('matrix-count').value = "1";
    document.getElementById('linear-matrix-rows').value = "3";
    document.getElementById('linear-matrix-columns').value = "3";

    // Создаем матрицу по умолчанию
    setupMatrixSizeSelection();
    initializeLinearFormCoefficientInput();
}

// Экспортируем функцию для глобального доступа
window.setupLinearFormMode = setupLinearFormMode;

// Показать поля для выбора размера матриц и для ввода матриц
function setupMatrixSizeSelection() {
    document.getElementById('linear-matrix-size-selection').style.display = 'flex';
    createMatrixFields(); // Создаем поля для одной матрицы
}

// Экспортируем функцию для глобального доступа
window.setupMatrixSizeSelection = setupMatrixSizeSelection;

