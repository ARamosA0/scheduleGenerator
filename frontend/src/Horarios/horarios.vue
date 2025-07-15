<template>
    <div class="grid grid-cols-12 gap-3">
        <Button
            icon="pi pi-angle-left"
            class="col-span-1"
            severity="secondary"
            @click="router.push({ name: 'home' })"
        ></Button>
        <div class="col-span-8">
            <p class="text-2xl font-bold">Configuración de Horarios</p>
            <p class="text-lg">
                Define los días y franjas horarias disponibles
            </p>
        </div>

        <div class="col-span-3">
            <Button
                icon="pi pi-plus"
                severity="secondary"
                label="Agregar template"
                class="w-full md:w-56"
                @click="open"
            />
            <Select
                v-model="selectedCity"
                :options="cities"
                showClear
                optionLabel="name"
                placeholder="Select a City"
                class="w-full md:w-56 mt-3"
            />
        </div>
    </div>
    <div v-for="(d, index) in data" :key="index">
        <Card class="mt-3">
            <template #title>
                <div class="flex justify-between items-center">
                    <div class="flex items-center">
                        <i class="pi pi-clock mr-2" />
                        <span class="text-2xl font-bold">{{ d.title }}</span>
                        <ToggleSwitch class="ml-2" v-model="d.state" />
                        <Tag
                            class="ml-2"
                            :severity="d.state ? 'contrast' : 'secondary'"
                        >
                            {{ d.state ? "Activo" : "Inactivo" }}
                        </Tag>
                    </div>
                    <Button label="Agregar Franja" icon="pi pi-plus" />
                </div>
            </template>
            <template #content>
                <DataTable>
                    <template #loading>
                        <p>
                            No hay franjas horarias configuradas para este día
                        </p>
                    </template>
                    <Column field="codigo" header="Nombre" />
                    <Column field="nombre" header="Hora Inicio" />
                    <Column field="creditos" header="Hora Fin" />
                    <Column field="horas" header="Estado">
                        <template #body>
                            <ToggleSwitch class="ml-2" v-model="d.state" />
                            <Tag
                                class="ml-2"
                                :severity="d.state ? 'contrast' : 'secondary'"
                            >
                                {{ d.state ? "Activo" : "Inactivo" }}
                            </Tag>
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
    </div>
    <AddTemplate v-model:visible="openDialog" />
</template>
<script setup lang="ts">
import { ref } from "vue";
import {
    Button,
    DataTable,
    Column,
    Card,
    ToggleSwitch,
    Select,
} from "primevue";
import { useRouter } from "vue-router";
import AddTemplate from "./addTemplate.vue";
const router = useRouter();
const openDialog = ref(false);

const data = ref([
    {
        title: "Lunes",
        state: true,
        dayData: [],
    },
    {
        title: "Martes",
        state: true,
        dayData: [],
    },
    {
        title: "Miercoles",
        state: true,
        dayData: [],
    },
    {
        title: "Jueves",
        state: true,
        dayData: [],
    },
    {
        title: "Viernes",
        state: true,
        dayData: [],
    },
    {
        title: "Sabado",
        state: true,
        dayData: [],
    },
]);

const open = () => {
    openDialog.value = true;
    console.log("openDIALGO", openDialog.value);
};
</script>
