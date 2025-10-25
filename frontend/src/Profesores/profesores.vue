<template>
    <div class="grid grid-cols-12 gap-3">
        <Button
            icon="pi pi-angle-left"
            class="col-span-1"
            severity="secondary"
            @click="router.push({ name: 'home' })"
        ></Button>
        <div class="col-span-11">
            <p class="text-2xl font-bold">Gestión de Profesores</p>
            <p class="text-lg">Administra la información de los docentes</p>
        </div>
    </div>

    <Card>
        <template #title>
            <div class="flex justify-between items-start">
                <div>
                    <p class="text-xl">Lista de Profesores</p>
                    <p class="text-base">Gestiona los profesores del sistema</p>
                </div>
                <div>
                    <Button
                        class="m-1"
                        icon="pi-file-excel"
                        @click="openExcelDialog = true"
                    />
                    <Button
                        class="m-1"
                        label="Agregar profesor"
                        @click="open"
                    />
                </div>
            </div>
        </template>
        <template #content>
            <DataTable
                :value="data"
                paginator
                :rows="5"
                :rowsPerPageOptions="[5, 10, 20, 50]"
                tableStyle="min-width:50rem"
                v-model:selection="selectedData"
                selectionMode="single"
                :metakeySelection="false"
                @rowSelect="onRowSelect"
            >
                <Column field="name" header="Nombre" />
                <Column field="email" header="Email" />
                <Column field="specialty" header="Especialidad" />
                <Column field="acciones" header="Acciones">
                    <template #body="slotProps" severity="secondary" rounded>
                        <Button
                            icon="pi pi-trash"
                            @click="delTeacher(slotProps.data)"
                        ></Button>
                    </template>
                </Column>
            </DataTable>
        </template>
    </Card>
    <AddTeacher
        v-model:visible="openDialog"
        @save="saveTeacher"
        @update="update"
        :teacher="selectedData"
    />
    <Dialog v-model:visible="openExcelDialog" modal header="Edit Profile">
        <uploadFile
            :items="columns"
            :table-column="columns"
            :type="'Teacher'"
        />
    </Dialog>
</template>

<script setup lang="ts">
import { ref, watch } from "vue";
import { DataTable, Column, Card, Button, Dialog } from "primevue";
import uploadFile from "../common/upload-file.vue";

import {
    getAllTeachers,
    createTeacher,
    updateTeacher,
    deleteTeacher,
} from "../../api/profesoresApi";

import { useRouter } from "vue-router";
import AddTeacher from "./addTeacher.vue";
const router = useRouter();

const data = ref([]);

const selectedData = ref(null);

const openDialog = ref(false);
const saved = ref(false);
const openExcelDialog = ref(false);

const columns = ref([
    {
        name: "Nombre",
        value: "name",
    },
    {
        name: "Apellido",
        value: "lastName",
    },
    {
        name: "Email",
        value: "email",
    },
    {
        name: "Telefono",
        value: "phone",
    },
    {
        name: "Especialidad",
        value: "specialty",
    },
    {
        name: "Disponibilidad",
        value: "available_days",
    },
]);

const getTeachers = async () => {
    const response = await getAllTeachers();
    console.log("RESPONSE", response);
    data.value = response;
};

watch(
    () => saved.value,
    async (newVal: any) => {
        console.log("newVAL", newVal);
        await getTeachers();
    },
    { immediate: true },
);

const open = () => {
    openDialog.value = true;
    console.log("openDIALGO", openDialog.value);
};

const saveTeacher = async (value: any) => {
    console.log("SAVE TEACHER", value);
    await createTeacher(value);
    saved.value = true;
    await getTeachers();
};

const update = async (value: any) => {
    console.log("UPDATE TEACHER", value);
    await updateTeacher(value);
    saved.value = true;
    // await getTeachers();
};

const delTeacher = async (value: any) => {
    await deleteTeacher(value);
    saved.value = true;
};

const onRowSelect = async (value: any) => {
    console.log("VALUE SELECTED", value);
    openDialog.value = true;
};
</script>
