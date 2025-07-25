<template>
    <div class="grid grid-cols-12 gap-3">
        <Button
            icon="pi pi-angle-left"
            class="col-span-1"
            severity="secondary"
            @click="router.push({ name: 'home' })"
        ></Button>
        <div class="col-span-11">
            <p class="text-2xl font-bold">Gestion de Cursos</p>
            <p class="text-lg">Administra las materias y asignaturas</p>
        </div>
    </div>

    <Card>
        <template #title>
            <div class="flex justify-between items-start">
                <div>
                    <p class="text-xl font-semibold">Lista de Cursos</p>
                    <p class="text-base text-gray-600">
                        Gestiona los cursos del sistema
                    </p>
                </div>
                <Button label="Agregar Curso" @click="open" />
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
                <Column field="code" header="Codigo" />
                <Column field="name" header="Nombre" />
                <Column field="creadits" header="Creditos" />
                <Column field="hours" header="Horas/Sem" />
                <Column field="semester" header="Semestre" />
                <Column field="career" header="Carrera" />
                <Column field="acciones" header="Acciones">
                    <template #body="slotProps" severity="secondary" rounded>
                        <Button
                            icon="pi pi-trash"
                            @click="delCourse(slotProps.data)"
                        ></Button>
                    </template>
                </Column>
            </DataTable>
        </template>
    </Card>
    <addCursos
        v-model:visible="openDialog"
        @save="saveCourse"
        @update="update"
        :course="selectedData"
    />
</template>

<script setup lang="ts">
import { ref, watch } from "vue";
import DataTable from "primevue/datatable";
import Column from "primevue/column";
import Card from "primevue/card";
import Button from "primevue/button";

import addCursos from "./addCursos.vue";
const selectedData = ref(null);

import {
    getAllCourses,
    createCourse,
    updateCourse,
    deleteCourse,
} from "../../api/cursosApi";

import { useRouter } from "vue-router";
const router = useRouter();

const openDialog = ref(false);
const saved = ref(false);

const data = ref([]);

const getCourse = async () => {
    const response = await getAllCourses();
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

const saveCourse = async (value: any) => {
    console.log("SAVE TEACHER", value);
    await createCourse(value);
    saved.value = true;
    await getCourse();
};

const update = async (value: any) => {
    console.log("UPDATE TEACHER", value);
    await updateCourse(value);
    saved.value = true;
    // await getTeachers();
};

const delCourse = async (value: any) => {
    await deleteCourse(value);
    saved.value = true;
};

const onRowSelect = async (value: any) => {
    console.log("VALUE SELECTED", value);
    openDialog.value = true;
};
</script>
