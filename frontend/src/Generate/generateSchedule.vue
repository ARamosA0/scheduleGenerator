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
            <Tab value="1">Configuracion y Ejecucion</Tab>
        </TabList>
        <TabPanels>
            <TabPanel value="0">
                <firstTab
                    :profesores="profesores"
                    :cursos="cursos"
                    :salones="salones"
                    @change-tab="changeTab"
                />
            </TabPanel>
            <TabPanel value="1">
                <secondTab
                    :process-name="data.processName"
                    :process-data="data.processData"
                />
            </TabPanel>
        </TabPanels>
    </Tabs>
</template>
<script setup lang="ts">
import { Button, Tabs, TabList, Tab, TabPanels, TabPanel } from "primevue";
import firstTab from "./firstTab.vue";
import secondTab from "./secondTab.vue";
import { ref } from "vue";
import { useRouter } from "vue-router";
const router = useRouter();

const data = ref({
    processName: "",
    processData: {
        poblacion: 0,
        generaciones: 0,
        mutacion: 0,
        cruce: 0,
        elitismo: 0,
    },
    start: false,
    tab: "0",
});

const profesores = ref([
    {
        status: false,
        title: "Juan Perez",
        subtitle: "Matematica",
        detalle: "Lunes, Miercoles, Viernes",
    },
]);
const cursos = ref([
    {
        status: false,
        title: "MAT101 - Cálculo I",
        subtitle: "Ingeniería - Semestre 1",
        detalle: "4 créditos • 6h/semana",
    },
]);
const salones = ref([
    {
        status: false,
        title: "A101 - Aula Magna",
        subtitle: "Auditorio - Edificio A",
        detalle: "Capacidad: 120 personas",
    },
]);

const changeTab = (value: any) => (data.value.tab = value);
</script>
