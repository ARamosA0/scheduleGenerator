<template>
    <div class="grid grid-cols-12 gap-3">
        <Button
            icon="pi pi-angle-left"
            class="col-span-1"
            severity="secondary"
            @click="router.push({ name: 'home' })"
        ></Button>
        <div class="col-span-11">
            <p class="text-2xl font-bold">Horario Semestre 2024-1</p>
            <p class="text-lg">Horario generado del 2024-01-15</p>
        </div>
        <div class="flex flex-row">
            <Button label="Exportar" />
            <Button label="Imprimir" />
            <Button label="Regenerar" />
        </div>
    </div>

    <div>
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

    <Card>
        <template #title>
            <div class="flex justify-between items-start">
                <div>
                    <p class="text-xl">Lista de Profesores</p>
                    <p class="text-base">Gestiona los profesores del sistema</p>
                </div>
                <Button label="Agregar profesor" />
            </div>
        </template>
        <template #content>
            <DataTable
                :value="data"
                paginator
                :rows="5"
                :rowsPerPageOptions="[5, 10, 20, 50]"
                tableStyle="min-width:50rem"
            >
                <Column field="codigo" header="Nombre" />
                <Column field="nombre" header="Email" />
                <Column field="creditos" header="Especialidad" />
                <Column field="horas" header="Disponibilidad" />
                <Column field="acciones" header="Acciones">
                    <template #body severity="secondary" rounded>
                        <Button icon="pi pi-trash"></Button>
                    </template>
                </Column>
            </DataTable>
        </template>
    </Card>
</template>

<script setup lang="ts">
import { ref } from "vue";
import DataTable from "primevue/datatable";
import Column from "primevue/column";
import Card from "primevue/card";
import Button from "primevue/button";

import { useRouter } from "vue-router";
const router = useRouter();

const data = ref([
    {
        codigo: "CS101",
        nombre: "Introducción a la Computación",
        creditos: 3,
        horas: 4,
        semestre: "I",
        carrera: "Ingeniería de Software",
        acciones: "Editar",
    },
]);
</script>
