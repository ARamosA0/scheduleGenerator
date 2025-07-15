<template>
    <Dialog
        v-model:visible="props.visible"
        modal
        header="Agregar Nuevo Salón"
        :style="{ width: '50rem' }"
    >
        <span class="text-surface-500 dark:text-surface-400 block mb-8">
            Completa la información del salón
        </span>
        <div class="grid grid-cols-2 gap-3">
            <div class="flex flex-col gap">
                <label for="name" class="font-semibold">Codigo</label>
                <InputText
                    id="name"
                    v-model="data.room.code"
                    autocomplete="off"
                />
            </div>
            <div class="flex flex-col">
                <label for="name" class="font-semibold">Nombre del Salón</label>
                <InputText
                    id="name"
                    v-model="data.room.name"
                    autocomplete="off"
                />
            </div>
        </div>
        <div class="grid grid-cols-3 gap-3 mt-3">
            <div class="flex flex-col">
                <label for="name" class="font-semibold">Capacidad</label>
                <InputNumber
                    id="credits"
                    type="number"
                    v-model="data.room.capacity"
                    autocomplete="off"
                />
            </div>
            <div class="flex flex-col gap">
                <label for="name" class="font-semibold">Tipo</label>
                <InputNumber
                    id="hours"
                    type="number"
                    v-model="data.room.room_type_id"
                    autocomplete="off"
                />
            </div>
            <div class="flex flex-col gap">
                <label for="name" class="font-semibold">Piso</label>
                <InputNumber
                    id="semester"
                    type="number"
                    v-model="data.room.floor"
                    autocomplete="off"
                />
            </div>
        </div>
        <div class="flex flex-col mt-3">
            <label for="name" class="font-semibold w-24">Edificio</label>
            <InputText
                id="name"
                v-model="data.room.building"
                autocomplete="off"
            />
        </div>
        <div class="flex flex-col mt-3">
            <label for="name" class="font-semibold w-24">Equipamiento</label>
            <Textarea
                id="name"
                v-model="data.room.observations"
                autocomplete="off"
                rows="5"
                cols="30"
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
                :label="Object.keys(props.room).length > 0 ? 'Update' : 'Save'"
                @click="save"
            ></Button>
        </div>
    </Dialog>
</template>
<script setup lang="ts">
import { ref, watch, toRef } from "vue";
import { Dialog, Button, InputText, Textarea, InputNumber } from "primevue";

const props = defineProps({
    room: {
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
    room: {
        code: "",
        name: "",
        capacity: 0,
        room_type_id: 0,
        floor: 0,
        building: "",
        observations: "",
    },
});
const teacherRef = toRef(props, "room");
watch(teacherRef, (newValue: any) => {
    data.value.room = newValue;
});

const closeDialog = () => {
    emit("update:visible", false);
};

const save = () => {
    if (props.room !== null) {
        console.log("UPDATE");
        emit("update:visible", false);
        emit("update", data.value.room);
    } else {
        emit("update:visible", false);
        console.log("EMIT TREACHER", data.value.room);
        emit("save", data.value.room);
    }
    data.value.room.code = "";
    data.value.room.name = "";
    data.value.room.capacity = 0;
    data.value.room.room_type_id = 0;
    data.value.room.floor = R0;
    data.value.room.building = "";
    data.value.room.observations = "";
};
</script>
