// matrix.js
// matrix.js

// matrix.js

// matrix.js

function createMatrixFields() {
    const matrixCount = parseInt(document.getElementById('matrix-count').value);
    const rows = parseInt(document.getElementById('linear-matrix-rows').value);
    const columns = parseInt(document.getElementById('linear-matrix-columns').value);
    const container = document.getElementById('matrix-input-fields-container');
    container.innerHTML = ''; // Очищаем контейнер

    for (let i = 1; i <= matrixCount; i++) {
        const matrixLabel = document.createElement('h3');
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

