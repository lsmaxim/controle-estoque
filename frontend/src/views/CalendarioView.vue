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

  height:700,
  eventColor: '#2c3e50',


  // EVENTOS
  events: eventos.value,

  // CLICAR NA DATA
  dateClick: async (info) => {

    // POPUP
    const resultado = await Swal.fire({

      title: 'Novo agendamento',

      input: 'text',

      inputPlaceholder: 'Digite a tarefa',

      confirmButtonText: 'Salvar',

      cancelButtonText: 'Cancelar',

      showCancelButton: true
    })

    // CANCELADO
    if (!resultado.isConfirmed) {
      return
    }

    try {

      // SALVAR NO BACKEND
      await api.post('/agendamentos', {

        titulo: resultado.value,

        data_agendamento: info.dateStr
      })

      // RECARREGA
      carregarAgendamentos()

      // SUCESSO
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

    // DADOS
    const titulo = info.event.title

    const data = info.event.startStr

    const status =
      info.event.backgroundColor === 'green'
        ? 'FINALIZADO'
        : 'PENDENTE'

    // POPUP INICIAL
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

    // CANCELADO
    if (!resultado.isConfirmed) {
      return
    }

    // OBSERVAÇÃO
    const observacao = await Swal.fire({

      title: 'Observação final',

      input: 'textarea',

      inputPlaceholder: 'Digite a observação final',

      confirmButtonText: 'Salvar',

      cancelButtonText: 'Cancelar',

      showCancelButton: true
    })

    // CANCELADO
    if (!observacao.isConfirmed) {
      return
    }

    try {

      // FINALIZA
      await api.put(
        `/agendamentos/${info.event.id}/finalizar`,
        {
          observacao: observacao.value
        }
      )

      // RECARREGA
      carregarAgendamentos()

      // SUCESSO
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

    // API
    const resposta = await api.get('/agendamentos')

    console.log(resposta.data)

    // CONVERTE PARA FULLCALENDAR
    eventos.value = resposta.data.map(item => ({

      id: item.id,

      title: item.titulo,

      date: item.data_agendamento,

      color:
        item.status === 'FINALIZADO'
          ? 'green'
          : '#2c3e50'
    }))

    // ATUALIZA CALENDÁRIO
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

  padding: 17px;
}

.card {

  background: white;

  padding: 15px;

  border-radius: 10px;

  box-shadow: 0 2px 10px rgba(0,0,0,0.08);
}
:deep(.fc-daygrid-day) {

  height: 40px;

}

:deep(.fc-daygrid-event) {

  font-size: 12px;

  padding: 4px;

  border-radius: 9px;
}

:deep(.fc) {

  font-size: 12px;
}

:deep(.fc-col-header-cell) {
  

  padding: 15px;

  font-size: 14px;
}
</style>