<template>
    <div class="grid grid-cols-12 gap-3">
        <Button
            icon="pi pi-angle-left"
            class="col-span-1"
            severity="secondary"
            @click="router.push({ name: 'historial' })"
        ></Button>
        <div class="col-span-7">
            <p class="text-2xl font-bold">Horario Semestre 2024-1</p>
            <p class="text-lg">Horario generado del 2024-01-15</p>
        </div>
        <div class="flex flex-row col-span-4">
            <Button label="Exportar" class="ml-3" @click="exportToPDF"/>
        </div>
    </div>

    <div class="grid grid-cols-3 gap-3" v-if="data.bestFitness !== null && data.bestFitness !== undefined">
        <Card>
            <template #content>
                <i class="pi pi-calendar" />
                <div class="flex flex-col">
                    <p>Fitness</p>
                    <p>{{ data.bestFitness }}</p>
                </div>
            </template>
        </Card>
        <Card>
            <template #content>
                <i class="pi pi-replay" />
                <div class="flex flex-col">
                    <p>Generaciones</p>
                    <p>{{ data.iteration }}</p>
                </div>
            </template>
        </Card>
        <Card>
            <template #content>
                <i class="pi pi-clock" />
                <div class="flex flex-col">
                    <p>Duracion</p>
                    <p>{{ data.duration }} s</p>
                </div>
            </template>
        </Card>
    </div>

    <Card class="mb-50">
        <template #title>
            <div class="flex justify-between items-start">
                <div>
                    <p class="text-xl">Horario Semanal</p>
                    <p class="text-base">
                        Distribucion de clases por dia y hora
                    </p>
                </div>
            </div>
        </template>
        <template #content>
            <SchoolSchedule
                v-if="data.items"
                ref="scheduleRef"
                :days="days"
                startHour="07:00"
                endHour="22:00"
                :slotMinutes="30"
                :rowHeight="45"
                :events="data.items"
                :legend="legend"
            />
            <!-- <CalendarView
                :displayPeriodUom="data.displayPeriodUom"
                :items="data.items"
            /> -->
        </template>
    </Card>
</template>

<script setup lang="js">
import { ref, onMounted } from "vue";
import { DataTable, Column, Card, Button, Select } from "primevue";
import { getScheduleById } from "../../api/scheduleApi";

import CalendarView from "../common/calendarView.vue";

import SchoolSchedule from "../common/SchoolSchedule.vue";

import { useRouter, useRoute } from "vue-router";
// import html2canvas from "html2canvas";
// import domtoimage from "dom-to-image-more";
import html2pdf from "html2pdf.js";

const scheduleRef =ref(null)

const days = ['Lunes', 'Martes', 'Miércoles', 'Jueves', 'Viernes', 'Sabado'];

const events = [
  { id: 1, dayIndex: 0, title: 'Matemática', startDate: '08:00', endDate: '09:30', room: '101', teacher: 'García', color: '#60a5fa' },
  { id: 2, dayIndex: 0, title: 'Comunicación', startDate: '09:00', endDate: '10:00', room: '102', teacher: 'Pérez', color: '#f59e0b' }, // se superpone
  { id: 3, dayIndex: 2, title: 'Historia', startDate: '11:00', endDate: '12:30', room: '201', teacher: 'Ramos', color: '#34d399' },
  { id: 4, dayIndex: 4, title: 'Inglés', startDate: '13:00', endDate: '14:00', room: '305', teacher: 'Smith', color: '#f472b6' },
];

const legend = [
  { label: 'Matemática', color: '#60a5fa' },
  { label: 'Comunicación', color: '#f59e0b' },
  { label: 'Historia', color: '#34d399' },
  { label: 'Inglés', color: '#f472b6' },
];


const router = useRouter();
const route = useRoute();

const props = defineProps({
    assigmentId: {
        type: String,
        default: "",
    },
});

const data = ref({
    displayPeriodUom: "month",
    items: null,
    bestFitness: null,
    iteration: null,
    duration: null,
});

const periodsOptions = ref([
    { name: "Vista Semanal", value: "week" },
    { name: "Vista Mensual", value: "month" },
    { name: "Vista Anual", value: "year" },
]);

const getSchedule = async () => {
    const response = await getScheduleById(route.params.id.toString());
    console.log("ITEMS", response);
    data.value.items = response.schedule_response.bestGeneration;
    data.value.bestFitness = response.schedule_response.bestFitness
    data.value.iteration = response.schedule_response.iteration
    data.value.duration = response.schedule_response.time
    console.log("DATA.VALUES", data.value)
};

onMounted( async() => {
    await getSchedule();
});

const exportToPDF = () => {
  if (!scheduleRef.value) return;

  const element = scheduleRef.value.$el ?? scheduleRef.value;

  const opt = {
    margin:       0.5,
    filename:     "horario.pdf",
    image:        { type: "jpeg", quality: 0.98 },
    html2canvas:  { scale: 2 },   // mayor resolución
    jsPDF:        { unit: "in", format: "a4", orientation: "landscape" }
  };

  html2pdf().set(opt).from(element).save();
};

// const exportSchedule = async () => {
//   if (!scheduleRef.value) return;

//   const canvas = await html2canvas(scheduleRef.value.$el ?? scheduleRef.value, {
//     scale: 2, 
//     backgroundColor: "#ffffff" 
//   });

//   const dataUrl = canvas.toDataURL("image/png");
//   const link = document.createElement("a");
//   link.href = dataUrl;
//   link.download = "horario.png";
//   link.click();
// };

// const exportSchedule = async () => {
//       const node = scheduleRef.value.$el ?? scheduleRef.value;
  
//   const scale = 3; // 2 o 3 veces más grande
//   const style = {
//     transform: `scale(${scale})`,
//     transformOrigin: "top left",
//     width: `${node.offsetWidth}px`,
//     height: `${node.offsetHeight}px`,
//   };

//   const param = {
//     width: node.offsetWidth * scale,
//     height: node.offsetHeight * scale,
//     style,
//     quality: 1, // máxima calidad (JPEG)
//     bgcolor: "#ffffff", // fondo blanco
//   };

//   domtoimage.toPng(node, param).then((dataUrl) => {
//     const link = document.createElement("a");
//     link.href = dataUrl;
//     link.download = "horario.png";
//     link.click();
//   });
// };
</script>
