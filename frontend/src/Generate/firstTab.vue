<template>
    <div>
        <label for="template">Selecciotar template de horas</label>
        <Select
            id="template"
            v-model="selectedData.selectedTemplate"
            :options="templates"
            optionLabel="name"
            optionValue="ID"
            placeholder="Seleccionar template"
            class="w-full mt-3"
        />
    </div>
    <div class="grid grid-cols-4 gap-2 mt-5">
        <Card>
            <template #title>
                <div class="flex justify-between items-start">
                    <div>
                        <p>Grupos</p>
                    </div>
                    <Button label="Todos" @click="selectAllGroups" />
                </div>
            </template>
            <template #content>
                <itemSelect :data="props.grupos" @toggle="onToggleGroup">
                    <template #item="{ item, onToggle }">
                        <Checkbox
                            v-model="item.status"
                            binary
                            @change="onToggle"
                        />
                        <div class="flex flex-col">
                            <p class="text-lg font-bold">
                                {{ item.name }}
                            </p>
                            <p class="text-sm">{{ item.size }} alumnos</p>
                        </div>
                    </template>
                </itemSelect>
            </template>
        </Card>
        <Card>
            <template #title>
                <div class="flex justify-between items-start">
                    <div>
                        <p class="font-bold">Profesores</p>
                    </div>
                    <Button label="Todos" @click="selectAllTeachers" />
                </div>
            </template>
            <template #content>
                <itemSelect :data="props.profesores" @toggle="onToggleTeacher">
                    <template #item="{ item, onToggle }">
                        <Checkbox
                            v-model="item.status"
                            binary
                            @change="onToggle"
                        />
                        <div class="flex flex-col">
                            <p class="text-lg font-bold">
                                {{ item.name }}
                            </p>
                            <p>{{ item.lastName }}</p>
                            <p class="text-sm">
                                Disponible: {{ item.available_days }}
                            </p>
                        </div>
                    </template>
                </itemSelect>
            </template>
        </Card>
        <Card>
            <template #title>
                <div class="flex justify-between items-start">
                    <div>
                        <p>Cursos</p>
                    </div>
                    <Button label="Todos" @click="selectAllSubjects" />
                </div>
            </template>
            <template #content>
                <itemSelect :data="props.cursos" @toggle="onToggleSubjects">
                    <template #item="{ item, onToggle }">
                        <Checkbox
                            v-model="item.status"
                            binary
                            @change="onToggle"
                        />
                        <div class="flex flex-col">
                            <p class="text-lg font-bold">
                                {{ item.code }} - {{ item.name }}
                            </p>
                            <p>
                                {{ item.career }} - Semester {{ item.semester }}
                            </p>
                            <p class="text-sm">
                                {{ item.credits }} créditos •
                                {{ item.hours }}h/semana
                            </p>
                        </div>
                    </template>
                </itemSelect>
            </template>
        </Card>
        <Card>
            <template #title>
                <div class="flex justify-between items-start">
                    <div>
                        <p>Salones</p>
                    </div>
                    <Button label="Todos" @click="selectAllRooms" />
                </div>
            </template>
            <template #content>
                <itemSelect :data="props.salones" @toggle="onToggleRooms">
                    <template #item="{ item, onToggle }">
                        <Checkbox
                            v-model="item.status"
                            binary
                            @change="onToggle"
                        />
                        <div class="flex flex-col">
                            <p class="text-lg font-bold">
                                {{ item.code }} - {{ item.name }}
                            </p>
                            <p>
                                {{ item.room_type }} - Semester
                                {{ item.building }} {{ item.floor }}
                            </p>
                            <p class="text-sm">
                                Capacidad: {{ item.capacity }} personas
                            </p>
                        </div>
                    </template>
                </itemSelect>
            </template>
        </Card>
    </div>
    <Card class="mt-4">
        <template #title>
            <div class="flex justify-between items-start">
                <div>
                    <p>Resumen de Selección</p>
                    <p class="text-sm">
                        Elementos seleccionados para la generación del horario
                    </p>
                </div>
            </div>
        </template>
        <template #content>
            <div class="grid grid-cols-4 gap-3 mt-2">
                <div class="text-center">
                    <p class="text-2xl">
                        {{ selectedData.selectedGroups.length }}
                    </p>
                    <p class="text-sm">Grupos Seleccionados</p>
                </div>
                <div class="text-center">
                    <p class="text-2xl">
                        {{ selectedData.selectedTeachers.length }}
                    </p>
                    <p class="text-sm">Profesores Seleccionados</p>
                </div>
                <div class="text-center">
                    <p class="text-2xl">
                        {{ selectedData.selectedSubjects.length }}
                    </p>
                    <p class="text-sm">Cursos Seleccionados</p>
                </div>
                <div class="text-center">
                    <p class="text-2xl">
                        {{ selectedData.selectedRooms.length }}
                    </p>
                    <p class="text-sm">Salones Seleccionados</p>
                </div>
            </div>
            <div class="mt-3 flex justify-center">
                <Button @click="ChangeTab" :disabled="allowContinue">
                    Continuar a Configuracion
                    <i class="pi pi-arrow-right" />
                </Button>
            </div>
        </template>
    </Card>
</template>
<script setup lang="ts">
import { ref, computed, watch, toRaw } from "vue";
import { Button, Card, Select, Checkbox } from "primevue";
import itemSelect from "../common/itemSelect.vue";

const props = defineProps({
    profesores: {
        type: Array,
        default: null,
    },
    cursos: {
        type: Array,
        default: null,
    },
    salones: {
        type: Array,
        default: null,
    },
    templates: {
        type: Array,
        default: null,
    },
    grupos: {
        type: Array,
        default: null,
    },
});

const selectedData = ref({
    selectedTemplate: "",
    selectedSubjects: [],
    selectedRooms: [],
    selectedTeachers: [],
    selectedGroups: [],
});

const emits = defineEmits(["changeTab", "tab1Data"]);
const ChangeTab = () => {
    console.log("TAB1DATA", selectedData.value);
    emits("changeTab", "1");
    emits("tab1Data", selectedData.value);
};

const onToggleGroup = (item: any) => {
    console.log('GROUP ITEM', item)
    const index = selectedData.value.selectedGroups.findIndex(
        (t) => t.ID === item.ID,
    );
    if (item.status && index === -1) {
        selectedData.value.selectedGroups.push(item);
    } else if (!item.status && index !== -1) {
        selectedData.value.selectedGroups.splice(index, 1);
    }
};

watch(() => selectedData.value.selectedGroups, (newValue: any) => {
    console.log('GROUP SELECTED', newValue)
    if (newValue !== undefined && newValue !== null){
        const rawValue = toRaw(newValue)
        rawValue.forEach((group : any) => {
            group.subjects.forEach((id: any) => {
                const selectedSubject = props.cursos.find((curso: any) => curso.ID === id )
                selectedSubject.status = true
                selectedData.value.selectedSubjects.push(selectedSubject)
                console.log('SUBJECT', selectedSubject)
            })
        })
    }
    console.log('SUBJECT SELECTED', selectedData.value.selectedSubjects)
},{deep: true})

const onToggleTeacher = (item: any) => {
    const index = selectedData.value.selectedTeachers.findIndex(
        (t) => t.ID === item.ID,
    );
    if (item.status && index === -1) {
        selectedData.value.selectedTeachers.push(item);
    } else if (!item.status && index !== -1) {
        selectedData.value.selectedTeachers.splice(index, 1);
    }
};

const onToggleSubjects = (item: any) => {
    const index = selectedData.value.selectedSubjects.findIndex(
        (t) => t.ID === item.ID,
    );
    if (item.status && index === -1) {
        selectedData.value.selectedSubjects.push(item);
    } else if (!item.status && index !== -1) {
        selectedData.value.selectedSubjects.splice(index, 1);
    }

    console.log('SUBJECTS SELECTED', selectedData.value.selectedSubjects)
};

const onToggleRooms = (item: any) => {
    const index = selectedData.value.selectedRooms.findIndex(
        (t) => t.ID === item.ID,
    );
    if (item.status && index === -1) {
        selectedData.value.selectedRooms.push(item);
    } else if (!item.status && index !== -1) {
        selectedData.value.selectedRooms.splice(index, 1);
    }
};

const allowContinue = computed(() => {
    return (
        selectedData.value.selectedTemplate === "" ||
        selectedData.value.selectedSubjects.length === 0 ||
        selectedData.value.selectedRooms.length === 0 ||
        selectedData.value.selectedTeachers.length === 0
    );
});

const selectAllTeachers = () => {
    props.profesores.forEach((profesor: any) => {
        profesor.status = true;
        onToggleTeacher(profesor);
    });
};

const selectAllSubjects = () => {
    props.cursos.forEach((curso: any) => {
        curso.status = true;
        onToggleSubjects(curso);
    });
};

const selectAllRooms = () => {
    props.salones.forEach((salon: any) => {
        salon.status = true;
        onToggleRooms(salon);
    });
};

const selectAllGroups = () => {
    props.grupos.forEach((group: any) => {
        group.status = true;
        onToggleGroup(group);
    });
};
</script>
