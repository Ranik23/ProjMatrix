function sendDataToServer(url, payload) {
    fetch(url, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(payload),
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
        })
        .then(result => {
            console.log("Ответ от сервера:", result);
            // После успешной отправки переходим на страницу результатов
            proceedToNextStep();
        })
        .catch(error => {
            console.error("Ошибка отправки данных:", error);
        });
}
