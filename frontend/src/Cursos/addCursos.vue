<template>
    <Dialog
        v-model:visible="props.visible"
        modal
        header="Agregar Nuevo Curso"
        :style="{ width: '50rem' }"
    >
        <span class="text-surface-500 dark:text-surface-400 block mb-8">
            Completa la información del curso
        </span>
        <div class="grid grid-cols-2 gap-3">
            <div class="flex flex-col gap">
                <label for="name" class="font-semibold">Codigo</label>
                <InputText
                    id="code"
                    v-model="data.course.code"
                    autocomplete="off"
                />
            </div>
            <div class="flex flex-col">
                <label for="name" class="font-semibold">Nombre del Curso</label>
                <InputText
                    id="name"
                    v-model="data.course.name"
                    autocomplete="off"
                    variant="outlined"
                />
            </div>
        </div>
        <div class="grid grid-cols-3 gap-3 mt-3">
            <div class="flex flex-col gap">
                <label for="name" class="font-semibold">Creditos</label>
                <InputNumber
                    id="credits"
                    type="number"
                    v-model="data.course.credits"
                    autocomplete="off"
                />
            </div>
            <div class="flex flex-col gap">
                <label for="name" class="font-semibold">Horas/Semanas</label>
                <InputNumber
                    id="hours"
                    type="number"
                    v-model="data.course.hours"
                    autocomplete="off"
                />
            </div>
            <div class="flex flex-col gap">
                <label for="name" class="font-semibold">Semestre</label>
                <InputNumber
                    id="semester"
                    type="number"
                    v-model="data.course.semester"
                    autocomplete="off"
                />
            </div>
        </div>
        <div class="flex flex-col mt-3">
            <label for="name" class="font-semibold">Carrera</label>
            <InputText
                id="career"
                v-model="data.course.career"
                autocomplete="off"
            />
        </div>
        <div class="flex flex-col mt-3">
            <label for="name" class="font-semibold w-24">Prerequisitos</label>
            <InputText
                id="requirements"
                v-model="data.course.requirements"
                autocomplete="off"
            />
        </div>
        <div class="flex flex-col mt-3">
            <label for="name" class="font-semibold">Descripción</label>
            <InputText
                id="description"
                v-model="data.course.description"
                autocomplete="off"
            />
        </div>
        <div class="flex flex-col mt-3">
            <label for="name" class="font-semibold">Salon Requerido</label>

            <Select
                v-model="data.course.required_room_type"
                :options="rooms"
                optionLabel="label"
                optionValue="value"
                placeholder="Seleccionar salon"
                class="w-full"
            />
        </div>
        <div class="flex flex-col mt-3">
            <label for="name" class="font-semibold">Especialidad Requerida</label>

            <Select
                v-model="data.course.specialty"
                :options="specialtis"
                optionLabel="label"
                optionValue="value"
                placeholder="Seleccionar especialidad"
                class="w-full"
            />
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
                :label="props.course !== null ? 'Update' : 'Save'"
                @click="saveTeacher"
            ></Button>
        </div>
    </Dialog>
</template>
<script setup lang="ts">
import { ref, watch, toRef, onMounted } from "vue";
import { Dialog, Button, InputText, InputNumber, Select } from "primevue";
import { RoomTypes } from "../common/enums";

const props = defineProps({
    course: {
        type: Object,
        default: null,
    },
    visible: {
        type: Boolean,
        default: null,
    },
});

const emit = defineEmits(["update:visible", "save", "update"]);

onMounted(async () => {
    // const allRooms = await getAllRooms();
    // data.value.rooms = allRooms;
});

const data = ref({
    visible: false,
    course: {
        code: "",
        name: "",
        credits: 0,
        hours: 0,
        semester: 0,
        career: "",
        requirements: "",
        description: "",
        required_room_type: 1,
        specialty: 1
    },
});
const courseRef = toRef(props, "course");
watch(courseRef, (newValue: any) => {
    if (newValue !== null && newValue !== undefined) {
        data.value.course = newValue;
    } else {
        data.value.course = {};
    }
});

const rooms = ref([
    { label: "Aula", value: RoomTypes.Classroom },
    { label: "Laboratorio", value: RoomTypes.Laboratory },
    { label: "Auditorio", value: RoomTypes.Auditorium },
    { label: "Oficina", value: RoomTypes.Office },
]);

const specialtis = ref([
    { label: "Ciencias Exactas", value: 1 },
    { label: "Ciencias Naturales", value: 2 },
    { label: "Ingenieria y Tecnologia", value: 3 },
    { label: "Ciencias Sociales", value: 4 },
    { label: "Lengua y Humanidades", value: 5 },
    { label: "Economia y Administracion", value: 6 },
    { label: "Salud", value: 7 },
]);

const closeDialog = () => {
    emit("update:visible", false);
};

const saveTeacher = () => {
    if (props.course !== null) {
        console.log("UPDATE");
        emit("update:visible", false);
        emit("update", data.value.course);
    } else {
        emit("update:visible", false);
        console.log("EMIT TREACHER", data.value.course);
        emit("save", data.value.course);
    }
    data.value.course.code = "";
    data.value.course.name = "";
    data.value.course.credits = 0;
    data.value.course.hours = 0;
    data.value.course.semester = 0;
    data.value.course.career = "";
    data.value.course.requirements = "";
    data.value.course.description = "";
};
</script>
