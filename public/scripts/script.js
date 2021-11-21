import GetRandomName from './names-generator.js'

new lc_emoji_picker('input', {
  trigger_position: {
    top: '6px',
    right: '6px',
  },
})

let conn

// chat elements
const chatTalk = document.getElementById('chat-talk')
const chatForm = document.getElementById('chat-form')
const chatInput = document.getElementById('chat-input')

const NAME = GetRandomName()

window.addEventListener('load', event => {
  if (!window['WebSocket']) {
    chatTalk.insertAdjacentHTML('beforeend', `<b>Error el navegador no soporta websocket</b>`)
    return
  }

  conn = new WebSocket('ws://' + document.location.host + `/ws?name=${NAME}`)

  conn.onclose = function (event) {
    chatTalk.insertAdjacentHTML('beforeend', `<b>Conexión Cerrada</b>`)
  }

  conn.onmessage = function (event) {
    const message = JSON.parse(event.data)
    appendMessage(message.name, 'received', message.message)
  }
})


chatForm.addEventListener('submit', event => {
  event.preventDefault()

  const message = chatInput.value
  if (!message) return

  appendMessage(NAME, 'send', message)
  chatInput.value = ''

  conn.send(JSON.stringify({ name: NAME, message: message }))
})

const appendMessage =(name, type, message) => {
  const msg = `
    <div class="msg msg-${type}">
        <div class="msg-info">
            <span class="msg-name">${name}</span>
            <span class="msg-time">${formatDate(new Date())}</span>
        </div>
    
        <div class="msg-text">
            ${message}
        </div>
    </div>
  `

  const doScroll = chatTalk.scrollTop > chatTalk.scrollHeight - chatTalk.clientHeight - 1
  chatTalk.insertAdjacentHTML('beforeend', msg)

  if (doScroll) {
    chatTalk.scrollTop = chatTalk.scrollHeight - chatTalk.clientHeight
  }
}

// ---------------------------------------------------------
// hora
// ---------------------------------------------------------

const formatDate =(date) => {
  const hours = '0' + date.getHours()
  const minutes = '0' + date.getMinutes()

  return `${hours.slice(-2)}:${minutes.slice(-2)}`
}
