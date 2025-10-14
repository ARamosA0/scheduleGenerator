import { defineStore } from 'pinia'

interface ProcessEventMessage {
  message: string
  progress: number
  status: string
  scheduleId: string
}

export const useSseStore = defineStore('sse', {
    state: () => ({
        eventSource: null as EventSource | null,
        progress: 0,
        scheduleId: "",
        messages: [] as ProcessEventMessage[],
        connected: false,
    }),
    
    actions: {
        connect(userId: number) {
            console.log('Intentando conectar SSE con ID:', userId)
            
            if (this.eventSource) {
                console.log('️Ya existe una conexión, cerrando...')
                this.disconnect()
            }
            
            const url = `http://localhost:8080/api/events?id=${userId}`
            this.eventSource = new EventSource(url)
            
            this.eventSource.onopen = () => {
                this.connected = true
                console.log('SSE conectado exitosamente')
            }
            
            this.eventSource.onmessage = (event) => {
                console.log('Evento recibido:', event.data)
                const data = JSON.parse(event.data) as ProcessEventMessage
                this.progress = data.progress
                this.scheduleId = data.scheduleId
                this.messages.push(data)
                
                // Cerrar si recibe señal de cierre
                if (event.data === '__CLOSE__') {
                    console.log('Recibido señal de cierre')
                    this.disconnect()
                }
            }
            
            this.eventSource.onerror = (error) => {
                console.error('Error SSE:', error)
                console.log('Estado de conexión:', this.eventSource?.readyState)
                this.connected = false
                this.disconnect()
            }
        },
        
        disconnect() {
            if (this.eventSource) {
                this.eventSource.close()
                this.eventSource = null
                this.connected = false
                console.log('SSE desconectado')
            }
        },
        
        clearMessages() {
            this.messages = []
        },
    },
})