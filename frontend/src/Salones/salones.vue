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
                <Button label="Agregar Salon" @click="open" />
            </div>
        </template>
        <template #content>
            <DataTable
                :value="data"
                paginator
                :rows="5"
                :rowsPerPageOptions="[5, 10, 20, 50]"
                tableStyle="min-width:50rem"
                @rowSelect="onRowSelect"
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
                        <Button icon="pi pi-trash" @click="delRoom"></Button>
                    </template>
                </Column>
            </DataTable>
        </template>
    </Card>
    <addRoom
        v-model:visible="openDialog"
        :room="selectedData"
        @save="saveRoom"
        @update="update"
    />
</template>

<script setup lang="ts">
import { ref, watch } from "vue";
import DataTable from "primevue/datatable";
import Column from "primevue/column";
import Card from "primevue/card";
import Button from "primevue/button";

import addRoom from "../Salones/addRoom.vue";

import {
    getAllRooms,
    createRoom,
    updateRoom,
    deleteRoom,
} from "../../api/salonesApi";

import { useRouter } from "vue-router";
const router = useRouter();

const data = ref([]);

const selectedData = ref({});

const openDialog = ref(false);
const saved = ref(false);

const getRooms = async () => {
    const response = await getAllRooms();
    console.log("RESPONSE", response);
    data.value = response;
};

watch(
    () => saved.value,
    async (newVal: any) => {
        console.log("newVAL", newVal);
        await getRooms();
    },
    { immediate: true },
);

const open = () => {
    openDialog.value = true;
    console.log("openDIALGO", openDialog.value);
};

const saveRoom = async (value: any) => {
    console.log("SAVE TEACHER", value);
    await createRoom(value);
    saved.value = true;
    await getRooms();
};

const update = async (value: any) => {
    console.log("UPDATE TEACHER", value);
    await updateRoom(value);
    saved.value = true;
    // await getTeachers();
};

const delRoom = async (value: any) => {
    await deleteRoom(value);
    saved.value = true;
};

const onRowSelect = async (value: any) => {
    console.log("VALUE SELECTED", value);
    openDialog.value = true;
};
</script>
