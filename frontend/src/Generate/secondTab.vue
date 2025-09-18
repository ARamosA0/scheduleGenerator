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
                            {{ configurationData.cross_over }}
                        </p>
                        <Slider
                            v-model="configurationData.cross_over"
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
                            Ratio de Seleccion:
                            {{ configurationData.selection }}
                        </p>
                        <Slider
                            v-model="configurationData.selection"
                            :min="0"
                            :max="1"
                            :step="0.01"
                            class="w-full mt-3"
                        />
                        <p class="mt-3">
                            Porcentaje de mejores individuos que se preservan
                        </p>
                    </div>
                    <div class="mt-3">
                        <p class="mt-3">
                            Ratio de reinsercion:
                            {{ configurationData.reinsertion }}
                        </p>
                        <Slider
                            v-model="configurationData.reinsertion"
                            :min="0"
                            :max="5"
                            :step="0.1"
                            class="w-full mt-3"
                        />
                        <p class="mt-3">
                            Porcentaje de individuos que se reinsertan
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
                    <span> 0% </span>
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
                    <div class="mt-5">
                        <div class="grid grid-cols-2 gap-4">
                            <Button 
                            :disabled="processStatus.finish"
                            label="Ver Historial" 
                            @click="router.push({ name: 'historial' })" 
                            severity="secondary" 
                            class="w-full"
                            />
                            <Button 
                            :disabled="processStatus.finish"
                            label="Ver Horario" 
                            @click="router.push({ name: 'calendario', params: { id: processStatus.scheduleId } })" 
                            class="w-full"
                            />
                        </div>
                    </div>
                </div>
            </template>
        </Card>
    </div>
</template>
<script setup lang="ts">
import { ref, computed, watch } from "vue";
import { Button, Card, InputText, Slider } from "primevue";
import ProgressBar from "primevue/progressbar";
import { useRouter } from "vue-router";

const router = useRouter();

const props = defineProps({
    processName: {
        type: String,
        default: "",
    },
    processData: {
        type: Object,
        default: null,
    },
    executionResult: {
        type: Object,
        default: null
    }
});

const configurationData = ref({
    processName: "",
    population: 100,
    generations: 200,
    mutation: 0.05,
    cross_over: 0.8,
    selection: 0.5,
    reinsertion: 2,
});

const processStatus = ref({
    progrecion: 0,
    generacion: 0,
    fitness: 0,
    scheduleId: "",
    finish: true
});

const animateProgress = (target: any) => {
  const interval = setInterval(() => {
    if (processStatus.value.progrecion < target) {
      processStatus.value.progrecion += 1   // velocidad = +1
    } else {
      clearInterval(interval)
    }
  }, 500)}

watch(() => props.executionResult, (newVal) => {
    processStatus.value.progrecion = 100
    processStatus.value.generacion = newVal.iteration   
    processStatus.value.fitness = newVal.bestFitness
    processStatus.value.scheduleId = newVal.scheduleId
    processStatus.value.finish = false
})

const startProcess = ref(false);

const emits = defineEmits(["startProcess"]);

const runProcess = () => {
    startProcess.value = true;
    emits("startProcess", configurationData.value);
    animateProgress(100)
};

const validacionDatos = computed(() => {
    return (
        configurationData.value.processName === "" ||
        configurationData.value.population === 0 ||
        configurationData.value.generations === 0 ||
        configurationData.value.mutation === 0 ||
        configurationData.value.cross_over === 0 ||
        configurationData.value.reinsertion === 0 ||
        configurationData.value.selection === 0
    );
});
</script>
