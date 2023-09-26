function sendData() {
    const amount = document.getElementById("inputField").value;
    const curOld = document.getElementById("currencyOld").value;
    const curNew = document.getElementById("currencyNew").value;

    // Отправляем GET-запрос на сервер и обрабатываем ответ
    fetch(`http://localhost:8000/?amount=${amount}&currencyOld=${curOld}&currencyNew=${curNew}`)
        .then(response => response.text())
        .then(data => {
            // Отображаем полученные данные во втором поле
            document.getElementById("outputField").value = data;
        })
        .catch(error => console.error(error));
}