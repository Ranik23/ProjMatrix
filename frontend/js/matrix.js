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

        // Создаем контейнер для текущей матрицы и применяем нужные классы
        const matrixContainer = document.createElement('div');
        matrixContainer.classList.add('matrix-border');
        matrixContainer.style.display = 'grid';
        matrixContainer.style.gridTemplateColumns = `repeat(${columns}, 50px)`; // Используем ширину из стилей

        // Создаем поля ввода для каждой ячейки матрицы
        for (let j = 0; j < rows * columns; j++) {
            const input = document.createElement('input');
            input.type = 'number';
            input.value = 0;
            input.min = 0;
            input.max = 10;

            // Применяем стили непосредственно для соответствия полиномной странице
            input.style.width = '50px';
            input.style.height = '40px';
            input.style.textAlign = 'center';
            input.style.fontSize = '1rem';
            input.style.border = '1px solid #ccc';
            input.style.borderRadius = '5px';

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

        container.appendChild(matrixContainer);
    }

    container.style.display = 'block';
}


// Экспортируем функцию для глобального доступа
window.createMatrixFields = createMatrixFields;


// Создать поля ввода матрицы
function createMatrixInputs() {
    const rows = parseInt(document.getElementById('matrix-rows').value);
    const columns = parseInt(document.getElementById('matrix-columns').value);
    const matrixContainer = document.getElementById('matrix-container');
    matrixContainer.innerHTML = ''; // Очищаем контейнер

    matrixContainer.style.gridTemplateColumns = `repeat(${columns}, 50px)`;

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

