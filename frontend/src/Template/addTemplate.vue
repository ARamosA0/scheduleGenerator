<template>
    <Dialog
        v-model:visible="props.visible"
        modal
        header="Agregar Template"
        :style="{ width: '50rem' }"
    >
        <span class="text-surface-500 dark:text-surface-400 block mb-8">
            Completa la información del template
        </span>
        <div class="flex flex-col mt-3">
            <label for="name" class="font-semibold">Nombre del Template</label>
            <InputText id="career" v-model="template.name" autocomplete="off" />
        </div>
        <div class="flex justify-end mt-4">
            <Button
                icon="pi pi-plus"
                class="mr-3"
                @click="addTimeRange"
                severity="contrast"
            />
        </div>
        <Card v-for="(tr, key) in timeRange" :key="key" class="mt-4">
            <template #header>
                <Button icon="pi pi-trash" severity="secondary" @click="" />
            </template>
            <template #content>
                <div class="flex flex-col mt-3">
                    <label for="name" class="font-semibold">Dias</label>
                    <Select
                        id="day"
                        :options="days"
                        v-model="tr.day"
                        optionLabel="label"
                        optionValue="value"
                        placeholder="Seleccionar dia"
                        class="w-full"
                    />
                </div>
                <div class="grid grid-cols-2 gap-3">
                    <div class="flex flex-col gap">
                        <label for="name" class="font-semibold"
                            >Hora de Inicio</label
                        >
                        <DatePicker
                            id="datepicker-timeonly"
                            timeOnly
                            fluid
                            v-model="tr.startHour"
                            autocomplete="off"
                        />
                    </div>
                    <div class="flex flex-col">
                        <label for="name" class="font-semibold"
                            >Hora de Fin</label
                        >
                        <DatePicker
                            id="datepicker-timeonly"
                            timeOnly
                            fluid
                            v-model="tr.endHour"
                            autocomplete="off"
                            variant="outlined"
                        />
                    </div>
                    <div class="flex flex-col">
                        <label for="name" class="font-semibold"
                            >Periodo entre horas en minutos</label
                        >
                        <Select
                            id="period"
                            :options="periods"
                            v-model="tr.period"
                            optionLabel="label"
                            optionValue="value"
                            placeholder="Seleccionar periodo de horas"
                            class="w-full"
                        />
                    </div>
                </div>
            </template>
        </Card>
        <div class="flex justify-end gap-2 mt-5">
            <Button
                type="button"
                label="Cancel"
                severity="secondary"
                @click="closeDialog"
            ></Button>
            <Button type="button" label="Save" @click="saveTemplate"></Button>
        </div>
    </Dialog>
</template>
<script setup lang="ts">
import { ref, watch, toRef } from "vue";
import { Dialog, Button, InputText, Select, Card, DatePicker } from "primevue";

interface TimeRange {
    day: string;
    startHour: Date;
    endHour: Date;
    period: number;
    status: boolean;
}

const props = defineProps({
    template: {
        type: Object,
        default: null,
    },
    visible: {
        type: Boolean,
        default: null,
    },
});

const days = [
    { label: "Lunes", value: 1 },
    { label: "Martes", value: 2 },
    { label: "Miércoles", value: 3 },
    { label: "Jueves", value: 4 },
    { label: "Viernes", value: 5 },
    { label: "Sabado", value: 6 },
    { label: "Domingo", value: 7 },
];

const periods = [
    { label: "15 min", value: 15 },
    { label: "30 min", value: 30 },
    { label: "45 min", value: 45 },
    { label: "60 min", value: 60 },
];

const emit = defineEmits(["update:visible", "save", "update"]);

const template = ref({
    name: "",
    daysRange: "",
});

const timeRange = ref<TimeRange[]>([
    {
        day: "",
        startHour: new Date(2000, 0, 1, 8, 0),
        endHour: new Date(2000, 0, 1, 10, 0),
        period: 0,
        status: true,
    },
]);

const courseRef = toRef(props, "template");

watch(courseRef, (newValue: any) => {
    template.value = newValue;
    const parsedDates = newValue.daysRangeParsed.map(
        (item: object, key: any) => {
            return {
                ...item,
                startHour: new Date(item.startHour),
                endHour: new Date(item.endHour),
            };
        },
    );
    timeRange.value = parsedDates;
    console.log("NEWVALUE.TEMPLATE", newValue.daysRangeParsed);
});

function addTimeRange() {
    timeRange.value.push({
        day: "",
        startHour: new Date(2000, 0, 1, 8, 0),
        endHour: new Date(2000, 0, 1, 10, 0),
        period: 0,
        status: true,
    });
}

const closeDialog = () => {
    emit("update:visible", false);
};

const saveTemplate = () => {
    console.log("props.template", props.template);
    if (props.template !== null) {
        emit("update:visible", false);
        template.value.daysRange = JSON.stringify(timeRange.value);
        emit("update", template.value);
    } else {
        emit("update:visible", false);
        template.value.daysRange = JSON.stringify(timeRange.value);
        emit("save", template.value);
    }
};
</script>
