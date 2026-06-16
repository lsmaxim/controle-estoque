```vue
<script setup>

import { ref, onMounted } from 'vue'
import Swal from 'sweetalert2'
import api from '../services/api'

import FullCalendar from '@fullcalendar/vue3'
import dayGridPlugin from '@fullcalendar/daygrid'
import interactionPlugin from '@fullcalendar/interaction'

// EVENTOS
const eventos = ref([])

// CONFIGURAÇÃO DO CALENDÁRIO
const calendarOptions = ref({

  plugins: [
    dayGridPlugin,
    interactionPlugin
  ],

  initialView: 'dayGridMonth',

  locale: 'pt-br',

  contentHeight: 'auto',

  displayEventTime: false,

  dayMaxEvents: true,

  eventColor: '#2c3e50',

  events: eventos.value,

  // CLICAR NA DATA
  dateClick: async (info) => {

    const resultado = await Swal.fire({

      title: 'Novo agendamento',

      input: 'text',

      inputPlaceholder: 'Digite a tarefa',

      confirmButtonText: 'Salvar',

      cancelButtonText: 'Cancelar',

      showCancelButton: true
    })

    if (!resultado.isConfirmed) {
      return
    }

    try {

      await api.post('/agendamentos', {

        titulo: resultado.value,

        data_agendamento: info.dateStr
      })

      carregarAgendamentos()

      Swal.fire({

        icon: 'success',

        title: 'Sucesso',

        text: 'Agendamento criado com sucesso.'
      })

    } catch (erro) {

      console.error('Erro ao criar agendamento:', erro)

      Swal.fire({

        icon: 'error',

        title: 'Erro',

        text: 'Erro ao criar agendamento.'
      })
    }
  },

  // CLICAR NO EVENTO
  eventClick: async (info) => {

    const titulo = info.event.title

    const data = info.event.startStr

    const status =
      info.event.backgroundColor === 'green'
        ? 'FINALIZADO'
        : 'PENDENTE'

    const resultado = await Swal.fire({

      icon: 'info',

      title: titulo,

      html: `
        <div style="text-align:left">

          <p>
            <b>Data:</b> ${data}
          </p>

          <p>
            <b>Status:</b> ${status}
          </p>

          <hr>

          <p>
            Deseja finalizar este agendamento?
          </p>

        </div>
      `,

      confirmButtonText: 'Finalizar',

      cancelButtonText: 'Cancelar',

      showCancelButton: true
      
    })

    if (!resultado.isConfirmed) {
      return
    }

    const observacao = await Swal.fire({

      title: 'Observação final',

      input: 'textarea',

      inputPlaceholder: 'Digite a observação final',

      confirmButtonText: 'Salvar',

      cancelButtonText: 'Cancelar',

      showCancelButton: true
    })

    if (!observacao.isConfirmed) {
      return
    }

    try {

      await api.put(
        `/agendamentos/${info.event.id}/finalizar`,
        {
          observacao: observacao.value
        }
      )

      carregarAgendamentos()

      Swal.fire({

        icon: 'success',

        title: 'Finalizado',

        text: 'Agendamento finalizado com sucesso.'
      })

    } catch (erro) {

      console.error('Erro ao finalizar:', erro)

      Swal.fire({

        icon: 'error',

        title: 'Erro',

        text: 'Erro ao finalizar agendamento.'
      })
    }
  }
})

// CARREGAR AGENDAMENTOS
async function carregarAgendamentos() {

  try {

    const resposta = await api.get('/agendamentos')

    console.log(resposta.data)

    eventos.value = resposta.data.map(item => ({

      id: item.id,

      title: item.titulo,

      start: item.data_agendamento,

      allDay: true,

      color:
        item.status === 'FINALIZADO'
          ? 'green'
          : '#2c3e50'
    }))

    calendarOptions.value.events = [...eventos.value]

  } catch (erro) {

    console.error('Erro ao carregar agendamentos:', erro)

    Swal.fire({

      icon: 'error',

      title: 'Erro',

      text: 'Erro ao carregar agendamentos.'
    })
  }
}

// AO ABRIR A PÁGINA
onMounted(() => {

  carregarAgendamentos()
})

</script>

<template>

  <div class="pagina">

    <div class="card">

      <h1>Agendamento</h1>

      <FullCalendar :options="calendarOptions" />

    </div>

  </div>

</template>

<style scoped>

.pagina {

  padding: 15px;
}

.card {

  background: white;

  padding: 15px;

  border-radius: 10px;

  box-shadow: 0 2px 10px rgba(0,0,0,0.08);
}

/* CÉLULAS */

:deep(.fc-daygrid-day) {

  min-height: 90px;
}

/* EVENTOS */

:deep(.fc-daygrid-event) {

  width: 98%;

  display: block;

  font-size: 12px;

  padding: 6px 10px;

  border-radius: 8px;

  margin-top: 4px;

  border: none;

  white-space: nowrap;

  overflow: hidden;

  text-overflow: ellipsis;
}

/* TEXTO */

:deep(.fc-event-title) {

  font-weight: 500;

  overflow: hidden;

  text-overflow: ellipsis;
}

/* HEADER */

:deep(.fc-col-header-cell) {

  padding: 12px;

  font-size: 15px;
}

/* NÚMEROS DOS DIAS */

:deep(.fc-daygrid-day-number) {

  font-size: 12px;

  padding: 8px;
}

/* CALENDÁRIO */

:deep(.fc) {

  font-size: 14px;
}

</style>
```
