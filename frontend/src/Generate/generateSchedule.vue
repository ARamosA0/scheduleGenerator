<template>
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
                    @change-tab="changeTab"
                />
            </TabPanel>
            <TabPanel value="1">
                <secondTab
                    :process-name="data.processName"
                    :process-data="data.processData"
                    @start-process="runProcess"
                    @tab1Data="tab1Data"
                />
            </TabPanel>
        </TabPanels>
    </Tabs>
</template>
<script setup lang="ts">
import { Button, Tabs, TabList, Tab, TabPanels, TabPanel } from "primevue";
import firstTab from "./firstTab.vue";
import secondTab from "./secondTab.vue";
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
const router = useRouter();
import { getAllCourses } from "../../api/cursosApi";
import { getAllRooms } from "../../api/salonesApi";
import { getAllTeachers } from "../../api/profesoresApi";
import { getAllTemplates } from "../../api/templateApi";
import { createAssigment } from "../../api/assigmentApi";

const data = ref({
    processName: "",
    processData: {
        poblacion: 0,
        generaciones: 0,
        mutacion: 0,
        cruce: 0,
        elitismo: 0,
    },
    selectedData: {
        selectedTemplate: "",
        selectedSubjects: [],
        selectedRooms: [],
        selectedTeachers: [],
    },
    start: false,
    tab: "0",
});

const tab1Data = (value: any) => {
    data.value.selectedData = value;
};

const profesores = ref();
const cursos = ref();
const salones = ref();
const templates = ref();

const changeTab = (value: any) => (data.value.tab = value);

onMounted(async () => {
    profesores.value = await getAllTeachers();
    cursos.value = await getAllCourses();
    salones.value = await getAllRooms();
    templates.value = await getAllTemplates();

    console.log("TEMPLATES", templates.value);
});

const runProcess = async (value: any) => {
    await createAssigment(value);
};
</script>
