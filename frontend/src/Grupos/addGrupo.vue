<template>
    <Dialog
        v-model:visible="props.visible"
        modal
        header="Agregar Nuevo Grupo"
        :style="{ width: '50rem' }"
    >
        <span class="text-surface-500 dark:text-surface-400 block mb-8">
            Completa la información del grupo
        </span>
        <div class="grid grid-cols-2 gap-3">
            <div class="flex flex-col gap">
                <label for="name" class="font-semibold">Nombre</label>
                <InputText
                    id="name"
                    v-model="data.group.name"
                    autocomplete="off"
                />
            </div>
            <div class="flex flex-col">
                <label for="name" class="font-semibold">Tamano de Grupo</label>
                <InputNumber
                    id="size"
                    v-model="data.group.size"
                    autocomplete="off"
                    variant="outlined"
                />
            </div>
        </div>
        <div class="flex flex-col mt-3">
            <label for="name" class="font-semibold">Salon Requerido</label>

            <MultiSelect
                v-model="data.group.subjects"
                :options="data.subjects"
                optionLabel="name"
                optionValue="ID"
                filter
                placeholder="Seleccionar Cursos"
                display="chip"
                class="w-full md:w-80"
            >
                <template #option="slotProps">
                    <div class="flex items-center">
                        <div>{{ slotProps.option.name }}</div>
                    </div>
                </template>
                <template #dropdownicon>
                    <i class="pi pi-map" />
                </template>

                <template #header>
                    <div class="font-medium px-3 py-2">Cursos Disponibles</div>
                </template>
                <template #footer>
                    <div class="p-3 flex justify-between">
                        <Button
                            label="Add New"
                            severity="secondary"
                            variant="text"
                            size="small"
                            icon="pi pi-plus"
                        />
                        <Button
                            label="Remove All"
                            severity="danger"
                            variant="text"
                            size="small"
                            icon="pi pi-times"
                        />
                    </div>
                </template>
            </MultiSelect>
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
                :label="props.group !== null ? 'Update' : 'Save'"
                @click="saveGroup"
            ></Button>
        </div>
    </Dialog>
</template>
<script setup lang="ts">
import { ref, watch, toRef, onMounted } from "vue";
import { Dialog, Button, InputText, MultiSelect, InputNumber } from "primevue";
import { getAllCourses } from "../../api/cursosApi";

const props = defineProps({
    group: {
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
    getCourse();
});

const data = ref({
    visible: false,
    group: {
        name: "",
        size: 0,
        subjects: [],
    },
    subjects: [],
});
const groupRef = toRef(props, "group");
watch(groupRef, (newValue: any) => {
    if (newValue !== null && newValue !== undefined) {
        data.value.group = newValue;
    } else {
        data.value.group = {};
    }
});

const getCourse = async () => {
    const response = await getAllCourses();
    console.log("RESPONSE", response);
    data.value.subjects = response;
};

const closeDialog = () => {
    emit("update:visible", false);
};

const saveGroup = () => {
    if (props.group !== null) {
        console.log("UPDATE");
        emit("update:visible", false);
        emit("update", data.value.group);
    } else {
        emit("update:visible", false);
        console.log("EMIT GROUP", data.value.group);
        emit("save", data.value.group);
    }
    data.value.group.name = "";
    data.value.group.size = "";
};
</script>
