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
            <InputText
                id="career"
                v-model="data.template.name"
                autocomplete="off"
            />
        </div>
        <div v-for="timeRange in data.template.timeRange">
            <div class="flex flex-col mt-3">
                <label for="name" class="font-semibold">Dias</label>
                <Select
                    id="career"
                    editable
                    v-model="timeRange.day"
                    optionLabel="day"
                    placeholder="Seleccionar dia"
                    class="w-full"
                />
            </div>
            <div class="grid grid-cols-2 gap-3">
                <div class="flex flex-col gap">
                    <label for="name" class="font-semibold"
                        >Hora de Inicio</label
                    >
                    <InputText
                        id="code"
                        v-model="timeRange.startHour"
                        autocomplete="off"
                    />
                </div>
                <div class="flex flex-col">
                    <label for="name" class="font-semibold">Hora de Fin</label>
                    <InputText
                        id="name"
                        v-model="timeRange.endHour"
                        autocomplete="off"
                        variant="outlined"
                    />
                </div>
                <div class="flex flex-col">
                    <label for="name" class="font-semibold">Hora de Fin</label>
                    <InputNumber
                        id="name"
                        v-model="timeRange.period"
                        inputId="integeronly"
                        fluid
                    />
                </div>
            </div>
        </div>
        <div class="flex justify-end gap-2 mt-5">
            <Button
                type="button"
                label="Cancel"
                severity="secondary"
                @click="closeDialog"
            ></Button>
            <Button
                type="button"
                :label="
                    Object.keys(props.template).length > 0 ? 'Update' : 'Save'
                "
                @click="saveTeacher"
            ></Button>
        </div>
    </Dialog>
</template>
<script setup lang="ts">
import { ref, watch, toRef } from "vue";
import { Dialog, Button, InputText, Select, InputNumber } from "primevue";

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

const emit = defineEmits(["update:visible", "save", "update"]);

const data = ref({
    visible: false,
    template: {
        name: "",
        timeRange: [
            {
                day: "",
                startHour: "",
                endHour: "",
                period: 0,
                status: false,
            },
        ],
    },
});
const courseRef = toRef(props, "template");
watch(courseRef, (newValue: any) => {
    data.value.template = newValue;
});

const closeDialog = () => {
    emit("update:visible", false);
};

const saveTeacher = () => {
    // if (props.course !== null) {
    //     console.log("UPDATE");
    //     emit("update:visible", false);
    //     emit("update", data.value.course);
    // } else {
    //     emit("update:visible", false);
    //     console.log("EMIT TREACHER", data.value.course);
    //     emit("save", data.value.course);
    // }
    // data.value.course.code = "";
    // data.value.course.name = "";
    // data.value.course.credits = 0;
    // data.value.course.hours = 0;
    // data.value.course.semester = 0;
    // data.value.course.career = "";
    // data.value.course.requirements = "";
    // data.value.course.description = "";
};
</script>
