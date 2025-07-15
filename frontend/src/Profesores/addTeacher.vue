<template>
    <Dialog
        v-model:visible="props.visible"
        modal
        header="Agregar Nuevo
Profesor"
        :style="{ width: '25rem' }"
    >
        <span class="text-surface-500 dark:text-surface-400 block mb-8">
            Completa la información del profesor
        </span>
        <div class="grid grid-cols-2 gap-3">
            <div class="flex flex-col gap">
                <label for="name" class="font-semibold w-24">Nombre</label>
                <InputText
                    id="name"
                    v-model="data.teacher.name"
                    class="flex-auto"
                    autocomplete="off"
                />
            </div>
            <div class="flex flex-col">
                <label for="name" class="font-semibold w-24">Apellido</label>
                <InputText
                    id="name"
                    v-model="data.teacher.lastName"
                    class="flex-auto"
                    autocomplete="off"
                />
            </div>
        </div>
        <div class="flex flex-col mt-3">
            <label for="name" class="font-semibold w-24">Email</label>
            <InputText
                id="name"
                v-model="data.teacher.email"
                class="flex-auto"
                autocomplete="off"
            />
        </div>
        <div class="flex flex-col mt-3">
            <label for="name" class="font-semibold w-24">Telefono</label>
            <InputText
                id="name"
                v-model="data.teacher.phone"
                class="flex-auto"
                autocomplete="off"
            />
        </div>
        <div class="flex flex-col mt-3">
            <label for="name" class="font-semibold w-24">Especialidad</label>
            <InputText
                id="name"
                v-model="data.teacher.specialty"
                class="flex-auto"
                autocomplete="off"
            />
        </div>
        <div class="grid grid-cols-2 gap-3 mt-5">
            <div class="flex items-center gap-2">
                <Checkbox
                    v-model="data.teacher.available_days"
                    inputId="ingredient1"
                    name="Lunes"
                    value="LUNES"
                />
                <label for="ingredient1"> Lunes </label>
            </div>
            <div class="flex items-center gap-2">
                <Checkbox
                    v-model="data.teacher.available_days"
                    inputId="ingredient2"
                    name="Martes"
                    value="MARTES"
                />
                <label for="ingredient2"> Martes </label>
            </div>
            <div class="flex items-center gap-2">
                <Checkbox
                    v-model="data.teacher.available_days"
                    inputId="ingredient3"
                    name="Miercoles"
                    value="MIERCOLES"
                />
                <label for="ingredient3"> Miercoles </label>
            </div>
            <div class="flex items-center gap-2">
                <Checkbox
                    v-model="data.teacher.available_days"
                    inputId="ingredient4"
                    name="Jueves"
                    value="JUEVES"
                />
                <label for="ingredient4"> Jueves </label>
            </div>
            <div class="flex items-center gap-2">
                <Checkbox
                    v-model="data.teacher.availableDays"
                    inputId="ingredient4"
                    name="Viernes"
                    value="VIERNES"
                />
                <label for="ingredient4"> Viernes </label>
            </div>
            <div class="flex items-center gap-2">
                <Checkbox
                    v-model="data.teacher.available_days"
                    inputId="ingredient4"
                    name="Sabado"
                    value="SABADO"
                />
                <label for="ingredient4"> Sabado </label>
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
                    Object.keys(props.teacher).length > 0 ? 'Update' : 'Save'
                "
                @click="saveTeacher"
            ></Button>
        </div>
    </Dialog>
</template>
<script setup lang="ts">
import { ref, watch, toRef } from "vue";
import { Dialog, Button, InputText, Checkbox } from "primevue";

const props = defineProps({
    teacher: {
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
    teacher: {
        name: "",
        lastName: "",
        email: "",
        phone: "",
        specialty: "",
        available_days: [],
    },
});
const teacherRef = toRef(props, "teacher");
watch(teacherRef, (newValue: any) => {
    data.value.teacher = newValue;
});

const closeDialog = () => {
    emit("update:visible", false);
};

const saveTeacher = () => {
    if (props.teacher !== null) {
        console.log("UPDATE");
        emit("update:visible", false);
        emit("update", data.value.teacher);
    } else {
        emit("update:visible", false);
        console.log("EMIT TREACHER", data.value.teacher);
        emit("save", data.value.teacher);
    }
    data.value.teacher.name = "";
    data.value.teacher.lastName = "";
    data.value.teacher.email = "";
    data.value.teacher.phone = "";
    data.value.teacher.specialty = "";
    data.value.teacher.available_days = [];
};
</script>
