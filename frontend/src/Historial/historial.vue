<template>
    <div class="grid grid-cols-12 gap-3">
        <Button
            icon="pi pi-angle-left"
            class="col-span-1"
            severity="secondary"
            @click="router.push({ name: 'home' })"
        ></Button>
        <div class="col-span-11">
            <p class="text-2xl font-bold">Historial de Procesos</p>
            <p class="text-lg">
                Revisa todos los procesos de generación ejecutados
            </p>
        </div>
    </div>
    <Card>
        <template #title>
            <div class="flex justify-between items-start">
                <div>
                    <p class="text-xl">Procesos Ejecutados</p>
                    <p class="text-base">
                        Historial completo de ejecuciones del algoritmo genético
                    </p>
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
                :loading="loading"
            >
                <Column field="processName" header="Proceso" />
                <Column field="CreatedAt" header="Fecha">
                    <template #body="{ data }">
                        {{
                            new Date(data.CreatedAt).toLocaleDateString("es-PE")
                        }}
                    </template>
                </Column>
                <Column field="" header="Duracion" />
                <Column field="" header="Estado" />
                <Column field="" header="Fitness" />
                <Column field="horas" header="Generaciones" />
                <Column field="acciones" header="Acciones">
                    <template #body="{ data }" severity="secondary" rounded>
                        <Button
                            icon="pi pi-trash"
                            @click="delAssigment(data)"
                        ></Button>
                        <Button
                            icon="pi pi-eye"
                            class="ml-3"
                            @click="
                                router.push({
                                    name: 'calendario',
                                    params: { id: data.ID },
                                })
                            "
                        />
                    </template>
                </Column>
            </DataTable>
        </template>
    </Card>
    <!-- <div class="grid grid-cols-3 gap-3 mt-3">
        <Card>
            <template #title>
                <p class="text-large">Estadisticas Generales</p>
            </template>
            <template #content>
                <div>
                    <span>Total Procesos: </span>
                    <span v-if="!loading">{{ data.length }}</span>
                </div>
                <div>
                    <span>Completados: </span>
                    <span v-if="!loading">{{ data.length }}</span>
                </div>
                <div>
                    <span>En Progreso: </span>
                    <span v-if="!loading">0</span>
                </div>
                <div>
                    <span>Fallidos: </span>
                    <span v-if="!loading">0</span>
                </div>
            </template>
        </Card>
        <Card>
            <template #title>
                <p class="text-large">Mejor Resultado</p>
            </template>
            <template #content>
                <div>
                    <p class="font-bold">Horario Semestre</p>
                    <p>fecha</p>
                </div>
                <div>
                    <span>Fitness: </span>
                    <span>3</span>
                </div>
                <div>
                    <span>Concflicto: </span>
                    <span>1</span>
                </div>
                <div>
                    <span>Generaciones: </span>
                    <span>1</span>
                </div>
            </template>
        </Card>
        <Card>
            <template #title>
                <p class="text-large">Tiempo Promedio</p>
            </template>
            <template #content>
                <div class="text-center">
                    <p class="font-bold text-3xl">17m 12s</p>
                    <p>Tiempo promedio de ejecucion</p>
                </div>
            </template>
        </Card>
    </div> -->
</template>
<script setup lang="ts">
import { ref, onMounted } from "vue";
import { Button, Card, DataTable, Column } from "primevue";
import { useRouter } from "vue-router";
import { getAllAssigment, deleteAssigment } from "../../api/assigmentApi";
const router = useRouter();

const data = ref([]);

const loading = ref(false);

onMounted(async () => {
    await getAssigments();
});

const getAssigments = async () => {
    const response = await getAllAssigment();
    data.value = response;
    // loading.value = true
};

const delAssigment = async (data: object) => {
    await deleteAssigment(data);
    await getAllAssigment();
    loading.value = true;
};
</script>
