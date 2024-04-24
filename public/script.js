const id = Math.random().toString()
const source = new EventSource(`http://localhost:8080/notify?id=${id}`)

source.addEventListener("open", () => console.log("OPEN: ", id))

source.addEventListener("greeting", (event) => console.log("greeting: ", event.data))

source.addEventListener("jumping", (event) => console.log("jumping: ", event.data))