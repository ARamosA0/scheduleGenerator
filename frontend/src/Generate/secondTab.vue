<template>
    <div class="grid grid-cols-6 gap-4">
        <Card class="col-span-3">
            <template #title>
                <div class="flex justify-between items-start">
                    <div>
                        <p class="font-bold text-xl">
                            Configuración del Algoritmo
                        </p>
                        <p class="text-sm">
                            Ajusta los parámetros del algoritmo genético
                        </p>
                    </div>
                </div>
            </template>
            <template #content>
                <p>Nombre del Proceso</p>
                <InputText
                    type="text"
                    v-model="configurationData.processName"
                    variant="filled"
                    class="w-full mt-3"
                />
                <div class="mt-5">
                    <div class="mt-3">
                        <p class="mt-3">
                            Tamaño de Población:
                            {{ configurationData.population }}
                        </p>
                        <Slider
                            v-model="configurationData.population"
                            :min="100"
                            :max="5000"
                            class="w-full mt-3"
                        />
                        <p class="mt-3">
                            Mayor población = mejor exploración, pero más lento
                        </p>
                    </div>
                    <div class="mt-3">
                        <p class="mt-3">
                            Número de Generaciones:
                            {{ configurationData.generations }}
                        </p>
                        <Slider
                            v-model="configurationData.generations"
                            :min="0"
                            :max="100"
                            class="w-full mt-3"
                        />
                        <p class="mt-3">
                            Más generaciones = mejor convergencia
                        </p>
                    </div>
                    <div class="mt-3">
                        <p class="mt-3">
                            Tasa de Mutación:
                            {{ configurationData.mutation }}
                        </p>
                        <Slider
                            v-model="configurationData.mutation"
                            :min="0"
                            :max="1"
                            :step="0.01"
                            class="w-full mt-3"
                        />
                        <p class="mt-3">Controla la diversidad genética</p>
                    </div>
                    <div class="mt-3">
                        <p class="mt-3">
                            Tasa de Cruce:
                            {{ configurationData.crossOver }}
                        </p>
                        <Slider
                            v-model="configurationData.crossOver"
                            :min="0"
                            :max="1"
                            :step="0.01"
                            class="w-full mt-3"
                        />
                        <p class="mt-3">
                            Probabilidad de cruce entre individuos
                        </p>
                    </div>
                    <div class="mt-3">
                        <p class="mt-3">
                            Elitismo:
                            {{ configurationData.elitism }}
                        </p>
                        <Slider
                            v-model="configurationData.elitism"
                            :min="0"
                            :max="1"
                            :step="0.01"
                            class="w-full mt-3"
                        />
                        <p class="mt-3">
                            Porcentaje de mejores individuos que se preservan
                        </p>
                    </div>
                </div>
                <div class="mt-3">
                    <Button
                        label="Iniciar Generacion"
                        :disabled="validacionDatos"
                        icon="pi pi-play"
                        @click="runProcess"
                    />
                </div>
            </template>
        </Card>
        <Card class="col-span-3">
            <template #title>
                <div class="flex justify-between items-start">
                    <div>
                        <p class="font-bold text-xl">Monitor de Ejecución</p>
                        <p class="text-sm">
                            Seguimiento en tiempo real del algoritmo
                        </p>
                    </div>
                </div>
            </template>
            <template #content>
                <Card
                    v-if="!startProcess"
                    class="flex justify-between items-start"
                >
                    <template #content>
                        Configura los parámetros y presiona "Iniciar Generación"
                        para comenzar
                    </template>
                </Card>
                <div v-else>
                    <p class="mt-3 font-bold text-lg">Progreso general</p>
                    <ProgressBar
                        :value="processStatus.progrecion"
                        class="mt-3"
                    ></ProgressBar>

                    <div class="grid grid-cols-6 gap-4 mt-3">
                        <Card class="col-span-3">
                            <template #content>
                                <div class="flex flex-col justify-center">
                                    <p class="text-3xl font-bold text-center">
                                        {{ processStatus.generacion }}
                                    </p>
                                    <p class="text-lg text-center">
                                        Generacion Actual
                                    </p>
                                </div>
                            </template>
                        </Card>
                        <Card class="col-span-3">
                            <template #content>
                                <div class="flex flex-col justify-center">
                                    <p class="text-3xl font-bold text-center">
                                        {{ processStatus.fitness }}
                                    </p>
                                    <p class="text-lg text-center">
                                        Fitness Actual
                                    </p>
                                </div>
                            </template>
                        </Card>
                    </div>
                </div>
            </template>
        </Card>
    </div>
</template>
<script setup lang="ts">
import { ref, computed } from "vue";
import { Button, Card, InputText, Slider } from "primevue";
import ProgressBar from "primevue/progressbar";
const props = defineProps({
    processName: {
        type: String,
        default: "",
    },
    processData: {
        type: Object,
        default: null,
    },
});

const configurationData = ref({
    processName: "",
    population: 0,
    generations: 0,
    mutation: 0,
    crossOver: 0,
    elitism: 0,
});

const processStatus = ref({
    progrecion: 0,
    generacion: 0,
    fitness: 0,
});

const startProcess = ref(false);

const emits = defineEmits(["startProcess"]);

const runProcess = () => {
    startProcess.value = true;
    emits("startProcess", configurationData.value);
};

const validacionDatos = computed(() => {
    return (
        configurationData.value.processName === "" ||
        configurationData.value.poblacion === 0 ||
        configurationData.value.generaciones === 0 ||
        configurationData.value.mutacion === 0 ||
        configurationData.value.cruce === 0 ||
        configurationData.value.elitismo === 0
    );
});
</script>
