window.addEventListener("load", function () {
    let ws = new WebSocket(`ws://${window.location.host}/feed`)
    ws.binaryType = "blob"

    ws.onmessage = function (event) {
        if (event.data instanceof Blob) {
            console.log("Received binary data.")

            let img = document.createElement('img')
            img.style.maxHeight = '500px'
            img.style.maxWidth = '100%'

            let reader = new FileReader()
            reader.onload = function (e) {
                img.src = e.target.result
            }
            reader.readAsDataURL(event.data)
            showAlert(img)
            return
        }

        showAlert(event.data)
    }
})

function demo() {
    showAlert("Demo Nachricht")
}

function showAlert(message) {
    const messageElement = document.createElement("div")
    messageElement.className = "message"

    const bubbleElement = document.createElement("div")
    bubbleElement.className = "bubble slide-in"

    messageElement.appendChild(bubbleElement)

    bubbleElement.append(message)
    document.getElementById("alert-container").appendChild(messageElement)

    setTimeout(function () {
        bubbleElement.className = "bubble slide-out"

        setTimeout(function () {
            messageElement.remove()
        }, 500)
    }, 7000)
}
