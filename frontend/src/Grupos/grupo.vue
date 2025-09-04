<template>
    <div class="grid grid-cols-12 gap-3">
        <Button
            icon="pi pi-angle-left"
            class="col-span-1"
            severity="secondary"
            @click="router.push({ name: 'home' })"
        ></Button>
        <div class="col-span-11">
            <p class="text-2xl font-bold">Gestion de Grupos</p>
            <p class="text-lg">Administra las secciones o carreras</p>
        </div>
    </div>

    <Card>
        <template #title>
            <div class="flex justify-between items-start">
                <div>
                    <p class="text-xl font-semibold">
                        Lista de Secciones o Carreras
                    </p>
                    <p class="text-base text-gray-600">
                        Gestiona las secciones o carreras del sistema
                    </p>
                </div>
                <div>
                    <Button
                        class="m-1"
                        icon="pi-file-excel"
                        @click="openExcelDialog = true"
                    />
                    <Button class="m-1" label="Agregar grupo" @click="open" />
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
                <Column field="size" header="Tamano" />
                <Column field="acciones" header="Acciones">
                    <template #body="slotProps" severity="secondary" rounded>
                        <Button
                            icon="pi pi-trash"
                            @click="delGroup(slotProps.data)"
                        ></Button>
                    </template>
                </Column>
            </DataTable>
        </template>
    </Card>
    <addGroup
        v-model:visible="openDialog"
        @save="saveGroup"
        @update="update"
        :group="selectedData"
    />
    <Dialog v-model:visible="openExcelDialog" modal header="Edit Profile">
        <uploadFile :items="columns" :table-column="columns" :type="'Group'" />
    </Dialog>
</template>

<script setup lang="ts">
import { ref, watch } from "vue";
import { DataTable, Column, Card, Button, Dialog } from "primevue";
import uploadFile from "../common/upload-file.vue";
import addGroup from "./addGrupo.vue";

// import addCursos from "./addCursos.vue";
const selectedData = ref(null);

import {
    getAllGroups,
    createGroup,
    updateGroup,
    deleteGroup,
} from "../../api/grupoApi";

import { useRouter } from "vue-router";
const router = useRouter();

const openDialog = ref(false);
const saved = ref(false);

const data = ref([]);

const openExcelDialog = ref(false);
const columns = ref([
    {
        name: "Nombre",
        value: "name",
    },
    {
        name: "Tamano",
        value: "size",
    },
    {
        name: "Cursos",
        value: "",
    },
]);

const getCourse = async () => {
    const response = await getAllGroups();
    console.log("RESPONSE", response);
    data.value = response;
};

watch(
    () => saved.value,
    async (newVal: any) => {
        await getCourse();
    },
    { immediate: true },
);

const open = () => {
    openDialog.value = true;
    console.log("openDIALGO", openDialog.value);
};

const saveGroup = async (value: any) => {
    console.log("SAVE GROUP", value);
    await createGroup(value);
    saved.value = true;
    await getCourse();
};

const update = async (value: any) => {
    console.log("UPDATE GROUP", value);
    await updateGroup(value);
    saved.value = true;
    // await getTeachers();
};

const delGroup = async (value: any) => {
    await deleteGroup(value);
    saved.value = true;
};

const onRowSelect = async (value: any) => {
    console.log("VALUE SELECTED", value);
    openDialog.value = true;
};
</script>
