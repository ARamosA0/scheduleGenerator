<template>
    {{ scheduleId }}
    <div class="grid grid-cols-12 gap-3">
        <Button
            icon="pi pi-angle-left"
            class="col-span-1"
            severity="secondary"
            @click="router.push({ name: 'home' })"
        ></Button>
        <div class="col-span-11">
            <p class="text-2xl font-bold">Generar Horario</p>
            <p class="text-lg">
                Selecciona los elementos y configura el algoritmo genético
            </p>
        </div>
    </div>

    <Tabs :value="data.tab" v-model="data.tab" class="w-full">
        <TabList>
            <Tab value="0">Seleccion de Datos</Tab>
            <Tab value="1" disabled>Configuracion y Ejecucion</Tab>
        </TabList>
        <TabPanels>
            <TabPanel value="0">
                <firstTab
                    :profesores="profesores"
                    :cursos="cursos"
                    :salones="salones"
                    :templates="templates"
                    :grupos="grupos"
                    @change-tab="changeTab"
                    @tab1Data="tab1Data"
                />
            </TabPanel>
            <TabPanel value="1">
                <secondTab @start-process="runProcess" :execution-result="executionResult" :execution-progress="executionProgress" :schedule-id="scheduleId"/>
            </TabPanel>
        </TabPanels>
    </Tabs>
</template>
<script setup lang="ts">
import { Button, Tabs, TabList, Tab, TabPanels, TabPanel } from "primevue";
import firstTab from "./firstTab.vue";
import secondTab from "./secondTab.vue";
import { ref, onMounted, watch } from "vue";
import { useRouter } from "vue-router";
const router = useRouter();
import { getAllCourses } from "../../api/cursosApi";
import { getAllRooms } from "../../api/salonesApi";
import { getAllTeachers } from "../../api/profesoresApi";
import { getAllTemplates } from "../../api/templateApi";
import { getAllGroups } from "../../api/grupoApi";
import { createAssigment } from "../../api/assigmentApi";
import { useSseStore } from '../../store/serverEvents'

const sse = useSseStore()

const executionResult = ref()
const executionProgress = ref()
const scheduleId = ref()

const data = ref({
    start: false,
    tab: "0",
});

const runProcessData = ref({
    processData: {
        processName: "",
        population: 0,
        generations: 0,
        mutation: 0,
        crossOver: 0,
        selection: 0,
        elitism: 0,
    },
    selectedData: {
        selectedTemplate: "",
        selectedSubjects: [],
        selectedRooms: [],
        selectedTeachers: [],
    },
});

const tab1Data = (value: any) => {
    runProcessData.value.selectedData = value;
};

const profesores = ref();
const cursos = ref();
const salones = ref();
const templates = ref();
const grupos = ref();

const changeTab = (value: any) => (data.value.tab = value);

onMounted(async () => {
    sse.connect(1)
    profesores.value = await getAllTeachers();
    cursos.value = await getAllCourses();
    salones.value = await getAllRooms();
    templates.value = await getAllTemplates();
    grupos.value = await getAllGroups();

    console.log("GROUPS", grupos.value);
    console.log('SSE', sse)
    console.log('SSE', sse.messages)
});

watch(sse.messages, (msg) => {
  if (msg) {
    console.log('MSG', msg)
    executionProgress.value = sse.progress
    if (sse.scheduleId !== ""){
        scheduleId.value = sse.scheduleId
    }
  }
})

const runProcess = async (value: any) => {
    runProcessData.value.processData = value;
    const result = await createAssigment(runProcessData.value);
    console.log('RESULT EXECUTION', result.data)
    executionResult.value = result.data
};
</script>
