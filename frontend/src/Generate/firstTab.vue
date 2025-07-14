<template>
    <div class="grid grid-cols-3 gap-3">
        <Card>
            <template #title>
                <div class="flex justify-between items-start">
                    <div>
                        <p class="font-bold">Profesores</p>
                        <p class="text-sm">
                            Selecciona los profesores disponibles
                        </p>
                    </div>
                    <Button label="Seleccioar Todos" />
                </div>
            </template>
            <template #content>
                <itemSelect :data="props.profesores" />
            </template>
        </Card>
        <Card>
            <template #title>
                <div class="flex justify-between items-start">
                    <div>
                        <p>Profesores</p>
                        <p class="text-sm">
                            Selecciona los profesores disponibles
                        </p>
                    </div>
                    <Button label="Seleccioar Todos" />
                </div>
            </template>
            <template #content>
                <itemSelect :data="props.cursos" />
            </template>
        </Card>
        <Card>
            <template #title>
                <div class="flex justify-between items-start">
                    <div>
                        <p>Profesores</p>
                        <p class="text-sm">
                            Selecciona los profesores disponibles
                        </p>
                    </div>
                    <Button label="Seleccioar Todos" />
                </div>
            </template>
            <template #content>
                <itemSelect :data="props.salones" />
            </template>
        </Card>
    </div>
    <Card class="mt-4">
        <template #title>
            <div class="flex justify-between items-start">
                <div>
                    <p>Resumen de Selección</p>
                    <p class="text-sm">
                        Elementos seleccionados para la generación del horario
                    </p>
                </div>
            </div>
        </template>
        <template #content>
            <div class="grid grid-cols-3 gap-3 mt-2">
                <div class="text-center">
                    <p class="text-2xl">0</p>
                    <p class="text-sm">Profesores Seleccionados</p>
                </div>
                <div class="text-center">
                    <p class="text-2xl">0</p>
                    <p class="text-sm">Cursos Seleccionados</p>
                </div>
                <div class="text-center">
                    <p class="text-2xl">0</p>
                    <p class="text-sm">Salones Seleccionados</p>
                </div>
            </div>
            <div class="mt-3 flex justify-center">
                <Button @click="ChangeTab">
                    Continuar a Configuracion
                    <i class="pi pi-arrow-right" />
                </Button>
            </div>
        </template>
    </Card>
</template>
<script setup lang="ts">
import { Button, Card } from "primevue";
import itemSelect from "../common/itemSelect.vue";

const props = defineProps({
    profesores: {
        type: Array,
        default: null,
    },
    cursos: {
        type: Array,
        default: null,
    },
    salones: {
        type: Array,
        default: null,
    },
});

const emits = defineEmits(["changeTab"]);

const ChangeTab = () => emits("changeTab", "1");
</script>
