<template>
    <h1 class="text-2xl font-bold">Sistema de horarios Universitarios</h1>
    <h4 class="text-lg font-medium mt-3">
        Generación automática de horarios usando algoritmos genéticos
    </h4>

    <div class="grid grid-cols-2 gap-3 mt-5">
        <div v-for="(cardData, index) in cardsData" :key="index">
            <Card class="mt-1">
                <template #title>
                    <div class="text-center">{{ cardData.title }}</div>
                </template>
                <template #content>
                    <div
                        class="flex flex-col items-center justify-center gap-2"
                    >
                        <h5>{{ cardData.subtitle }}</h5>
                        <Button
                            @click="router.push({ name: cardData.btnRedirect })"
                        >
                            {{ cardData.btnName }}
                        </Button>
                    </div>
                </template>
            </Card>
        </div>
    </div>
    <Card class="mt-5">
        <template #title>
            <i class="pi pi-file-arrow-up" />
            Resumen del Sistema
        </template>
        <template #content>
            <div class="grid grid-cols-4 gap-4 mt-2">
                <div
                    class="text-center"
                    v-for="(data, index) in resumeData"
                    :key="index"
                >
                    <p class="text-2x1 font-semibold">{{ data.value }}</p>
                    <p class="text-base">{{ data.title }}</p>
                </div>
            </div>
        </template>
    </Card>
</template>
<script lang="ts" setup>
import { ref } from "vue";
import Card from "primevue/card";
import Button from "primevue/button";

import { useRouter } from "vue-router";
const router = useRouter();

const cardsData = ref([
    {
        title: "Grupos",
        subtitle: "Gestionar información de Secciones o Carreras",
        btnName: "Administrar Secciones o Carreras",
        btnRedirect: "grupo",
    },
    {
        title: "Profesores",
        subtitle: "Gestionar información de docentes",
        btnName: "Administrar Profesores",
        btnRedirect: "profesores",
    },
    {
        title: "Cursos",
        subtitle: "Gestionar materias y asignaturas",
        btnName: "Administrar Cursos",
        btnRedirect: "cursos",
    },
    {
        title: "Salones",
        subtitle: "Gestionar aulas y espacios",
        btnName: "Administrar Salones",
        btnRedirect: "salones",
    },
    {
        title: "Template de Horas",
        subtitle: "Configurar días y franjas horarias",
        btnName: "Configurar Template de Horarios",
        btnRedirect: "template",
    },
    {
        title: "Generar Horarios",
        subtitle: "Ejecutar algoritmo genético",
        btnName: "Generar Nuevo Horario",
        btnRedirect: "generar",
    },
    {
        title: "Historial",
        subtitle: "Ver procesos ejecutados",
        btnName: "Ver Historial",
        btnRedirect: "historial",
    },
]);

const resumeData = ref([
    {
        value: 0,
        title: "Profesores",
    },
    {
        value: 0,
        title: "Cursos",
    },
    {
        value: 0,
        title: "Salones",
    },
    {
        value: 0,
        title: "Procesos",
    },
]);
</script>
