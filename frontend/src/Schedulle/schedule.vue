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
            <Button label="Exportar" class="ml-3" />
            <Button label="Imprimir" class="ml-3" />
            <Button label="Regenerar" class="ml-3" />
        </div>
    </div>

    <div class="grid grid-cols-4 gap-3">
        <Card>
            <template #content>
                <i class="pi pi-calendar" />
                <div class="flex flex-col">
                    <p>Fitness</p>
                    <p>95.0%</p>
                </div>
            </template>
        </Card>
        <Card>
            <template #content>
                <i class="pi pi-clock" />
                <div class="flex flex-col">
                    <p>Conflictos</p>
                    <p>2</p>
                </div>
            </template>
        </Card>
        <Card>
            <template #content>
                <i class="pi pi-replay" />
                <div class="flex flex-col">
                    <p>Generaciones</p>
                    <p>150</p>
                </div>
            </template>
        </Card>
        <Card>
            <template #content>
                <i class="pi pi-clock" />
                <div class="flex flex-col">
                    <p>Duracion</p>
                    <p>15m 23s</p>
                </div>
            </template>
        </Card>
    </div>

    <div>
        <Select
            v-model="data.displayPeriodUom"
            :options="periodsOptions"
            optionLabel="name"
            optionValue="value"
        />
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
            <CalendarView
                :displayPeriodUom="data.displayPeriodUom"
                :items="data.items"
            />
        </template>
    </Card>
</template>

<script setup lang="js">
import { ref, onMounted } from "vue";
import { DataTable, Column, Card, Button, Select } from "primevue";
import { getScheduleById } from "../../api/scheduleApi";

import CalendarView from "../common/calendarView.vue";

import { useRouter, useRoute } from "vue-router";
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
});

const periodsOptions = ref([
    { name: "Vista Semanal", value: "week" },
    { name: "Vista Mensual", value: "month" },
    { name: "Vista Anual", value: "year" },
]);

const getSchedule = async () => {
    console.log("ASSIGMENTIDS", route.params.id);
    const response = await getScheduleById(route.params.id.toString());
    data.value.items = response.schedule_response;
    console.log("ITEMS", response.schedule_response);
};

onMounted(() => {
    getSchedule();
});
</script>
