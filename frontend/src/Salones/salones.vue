<template>
    <div class="grid grid-cols-12 gap-3">
        <Button
            icon="pi pi-angle-left"
            class="col-span-1"
            severity="secondary"
            @click="router.push({ name: 'home' })"
        ></Button>
        <div class="col-span-11">
            <p class="text-2xl font-bold">Gestión de Salones</p>
            <p class="text-lg">Administra las aulas y espacios físicos</p>
        </div>
    </div>

    <Card>
        <template #title>
            <div class="flex justify-between items-start">
                <div>
                    <p class="text-xl">Lista de Salones</p>
                    <p class="text-base">Gestiona los salones del sistema</p>
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
                <Column field="codigo" header="Codigo" />
                <Column field="nombre" header="Nombre" />
                <Column field="creditos" header="Capacidad" />
                <Column field="horas" header="Tipo" />
                <Column field="horas" header="Ubicacion" />
                <Column field="horas" header="Equipamiento" />
                <Column field="horas" header="Estado">
                    <template #body="data">
                        <Tag :value="data" severity="secondary" />
                    </template>
                </Column>
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
