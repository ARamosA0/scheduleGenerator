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
                label="Agregar Template"
                class="w-full md:w-56"
                @click="open"
            />
            <Select
                v-model="selectedTemplate"
                :options="templates"
                showClear
                optionLabel="name"
                placeholder="Selecciona un Template"
                class="w-full md:w-56 mt-3"
            />
            <Button
                icon="pi pi-pencil"
                severity="secondary"
                label="Actualizar Template"
                class="w-full md:w-56 mt-3"
                @click="open"
            />
        </div>
    </div>
    <div v-if="selectedTemplate === null">
        <p class="text-center text-2xl mt-15">
            Seleccione un template o cree uno nuevo
        </p>
    </div>
    <div v-else>
        <p class="my-10 text-4xl font-bold uppercase">
            {{ selectedTemplate.name }}
        </p>
        <div
            v-for="(d, index) in selectedTemplate.daysRangeParsed"
            :key="index"
        >
            <Card class="mt-3">
                <template #title>
                    <div class="flex justify-between items-center">
                        <div class="flex items-center">
                            <i class="pi pi-clock mr-2" />
                            <span class="text-2xl font-bold">{{ d.day }}</span>
                            <ToggleSwitch class="ml-2" v-model="d.status" />
                            <Tag
                                class="ml-2"
                                :severity="d.state ? 'contrast' : 'secondary'"
                            >
                                {{ d.status ? "Activo" : "Inactivo" }}
                            </Tag>
                        </div>
                        <Button label="Agregar Franja" icon="pi pi-plus" />
                    </div>
                </template>
                <template #content>
                    <DataTable :value="d.periods">
                        <template #loading>
                            <p>
                                No hay franjas horarias configuradas para este
                                día
                            </p>
                        </template>
                        <Column field="name" header="Nombre" />
                        <Column field="startHour" header="Hora Inicio">
                            <template #body="{ data }">
                                {{ formatHour(data.startHour) }}
                            </template>
                        </Column>
                        <Column field="endHour" header="Hora Fin">
                            <template #body="{ data }">
                                {{ formatHour(data.endHour) }}
                            </template>
                        </Column>
                        <Column field="acciones" header="Acciones">
                            <template #body severity="secondary" rounded>
                                <Button
                                    icon="pi pi-trash"
                                    :disabled="!updateEnabled"
                                />
                            </template>
                        </Column>
                    </DataTable>
                </template>
            </Card>
        </div>
    </div>
    <AddTemplate
        v-model:visible="openDialog"
        :template="selectedTemplate"
        @save="save"
    />
</template>
<script setup lang="ts">
import { ref, onMounted } from "vue";
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

import {
    getAllTemplates,
    createTemplate,
    updateTemplate,
    deleteTemplate,
} from "../../api/templateApi";

const router = useRouter();
const openDialog = ref(false);

const templates = ref([]);
const selectedTemplate = ref(null);
const updateEnabled = ref(false);

onMounted(async () => {
    await getTemplates();
});

const formatHour = (date: Date | string): string => {
    const d = new Date(date);
    return d.toLocaleTimeString("es-PE", {
        hour: "2-digit",
        minute: "2-digit",
        hour12: false,
    });
};

const open = () => {
    openDialog.value = true;
    console.log("openDIALGO", openDialog.value);
};

const getTemplates = async () => {
    const allTemplates = await getAllTemplates();
    allTemplates.forEach((template: any) => {
        const periodos = parseDaysRange(template.daysRange);
        console.log("PER", periodos);
        template.daysRangeParsed = periodos;
    });
    console.log("ALL TEMPLATES", allTemplates);
    templates.value = allTemplates;
};

const parseDaysRange = (jsonString: string) => {
    try {
        const raw = JSON.parse(jsonString);
        return raw.map((item: any) => {
            console.log("ITEM", item);
            const periodos = dividirEnPeriodos(
                new Date(item.startHour),
                new Date(item.endHour),
                item.period,
            );

            return {
                day: item.day,
                periods: periodos,
                status: item.status,
            };
        });
    } catch (e) {
        console.error("Error al parsear daysRange:", e);
        return [];
    }
};

const dividirEnPeriodos = (
    startHour: Date,
    endHour: Date,
    period: number,
): { name: string; startHour: Date; endHour: Date }[] => {
    const result = [];
    let currentStart = startHour;
    let index = 1;
    while (currentStart < endHour) {
        const currentEnd = new Date(
            currentStart.getTime() + period * 60 * 1000,
        );

        result.push({
            name: `periodo ${index}`,
            startHour: new Date(currentStart),
            endHour: new Date(currentEnd > endHour ? endHour : currentEnd),
        });

        currentStart = currentEnd;
        index++;
    }

    return result;
};

const save = async (value: any) => {
    console.log("VALUE TEMPLATE", value);
    await createTemplate(value);
};
</script>
